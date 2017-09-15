package middleware

import (
	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/dtrinh100/Music-Playlist/src/api/model"
	"github.com/dgrijalva/jwt-go"

	"time"
	"net/http"
	"context"
	"crypto/rsa"
	"io/ioutil"
	"path/filepath"
	"errors"
)

const JWTTokenString = "jwtTokenString"

const mpClaimsKey = "musicPlaylistClaimsKey"
const mpJWTIssuer = "musicPlaylistIssuer"
const mpJWTCookieName = "musicPlaylistJWTAuth"

const jwtExpireMinutes = 30

const (
	// TODO: change these paths to a private path in production
	// NOTE: do NOT use these keys in production. These are for demo only.
	privKeyPath = "./src/mp.rsa"     // openssl genrsa -out mp.rsa 1024
	pubKeyPath  = "./src/mp.rsa.pub" // openssl rsa -in mp.rsa -pubout > mp.rsa.pub
)

// TokenFunc helps JWTMiddleware get a copy of the JWT Public Key.
type TokenFunc struct {
	RSAPublicKey *rsa.PublicKey
}

// Verify returns the public RSA key to compare
func (tf TokenFunc) Verify(tkn *jwt.Token) (interface{}, error) {
	// Ensure signing method
	if _, ok := tkn.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, errors.New("unexpected signing method")
	}

	return tf.RSAPublicKey, nil
}

/**
	JWTMiddleware helps manage JSON Web Tokens (JWT) and authorization to URL-paths.
*/
func JWTMiddleware(rw http.ResponseWriter, req *http.Request, next http.Handler, env *common.Env) {
	tknCookie, cookieErr := req.Cookie(mpJWTCookieName)

	if cookieErr != nil {
		common.JSONErrorResponse(rw, common.ErrMap{
			"Login": "User Needs To Log In"}, http.StatusUnauthorized)
		return
	}

	tknString := tknCookie.Value
	var claims model.AppClaims

	// Parse JWT from token-string, also get AppClaims
	tkn, jwtParseErr := jwt.ParseWithClaims(tknString, &claims,
		TokenFunc{RSAPublicKey: env.RSAKeys.Public}.Verify)

	if jwtParseErr != nil {
		switch jwtParseErr.(type) {
		case *jwt.ValidationError:
			validationErr := jwtParseErr.(*jwt.ValidationError)
			switch validationErr.Errors {
			case jwt.ValidationErrorExpired:
				common.JSONErrorResponse(rw, common.ErrMap{
					"Token": "Token Expired. Request A New One"}, http.StatusUnauthorized)
			default:
				common.GenericJSONErrorResponse(rw)
			}
		default:
			common.GenericJSONErrorResponse(rw)
		}

		return
	}

	if tkn.Valid {
		if claims.RefreshAt < time.Now().Unix() {
			refreshCookie(rw, claims, env)
		}

		ctx := context.WithValue(req.Context(), mpClaimsKey, claims)
		ctx = context.WithValue(ctx, JWTTokenString, tknString)

		next.ServeHTTP(rw, req.WithContext(ctx))
	} else {
		common.JSONErrorResponse(rw, common.ErrMap{
			"Token": "Invalid JWT"}, http.StatusInternalServerError)
	}
}

/**
	refreshCookie refreshes the cookie in the middleware if the cookie is near
	jwt-expiration time: jwt-refresh time < current time < jwt-expiration time.
*/
func refreshCookie(rw http.ResponseWriter, claims model.AppClaims, env *common.Env) {
		updatedTokenStr, expirationTime, jwtErr := GetJWT(env.RSAKeys.Private, claims.UserEmail, jwtExpireMinutes)

	if jwtErr != nil {
		common.JSONErrorResponse(rw, common.ErrMap{
			"Token": "Failed to sign"}, http.StatusInternalServerError)
	}

	SetSecuredCookie(rw, updatedTokenStr, expirationTime)
}

/**
	InitRSAKeyPair uses initRSAKeyPairHelper to init public/private keys for JWT.
*/
func InitRSAKeyPair() (*rsa.PublicKey, *rsa.PrivateKey) {
	return initRSAKeyPairHelper(pubKeyPath, privKeyPath)
}

/**
	initRSAKeyPairHelper initializes the public & private keys used for JWT.
*/
func initRSAKeyPairHelper(pubRelPath, privRelPath string) (*rsa.PublicKey, *rsa.PrivateKey) {
	getDataFromFile := func(path string) []byte {
		absPath, pathErr := filepath.Abs(path)
		common.Fatal(pathErr, "Could not read file path "+path)

		fileBytes, fileErr := ioutil.ReadFile(absPath)
		common.Fatal(fileErr, "Error reading from "+absPath)

		return fileBytes
	}

	publicKey, pubKeyErr := jwt.ParseRSAPublicKeyFromPEM(getDataFromFile(pubRelPath))
	common.Fatal(pubKeyErr, "Could not parse public key")
	privateKey, privKeyErr := jwt.ParseRSAPrivateKeyFromPEM(getDataFromFile(privRelPath))
	common.Fatal(privKeyErr, "Could not parse private key")

	return publicKey, privateKey
}

/**
	GetJWT returns a valid JWT string
*/
func GetJWT(rsaPrivateKey *rsa.PrivateKey, email string, minutes time.Duration) (string, time.Time, error) {
	if minutes < 2 {
		common.Fatal(errors.New("JWT Expiration Time Duration Error"),
			"JWT needs to have an expiration time >= 2 minutes")
	}

	return getJWTHelper(rsaPrivateKey, email, minutes, minutes/2)
}

func getJWTHelper(rsaPrivateKey *rsa.PrivateKey, email string, expMins, refMins time.Duration) (string, time.Time, error) {
	nowTime := time.Now()
	expirationTime := nowTime.Add(expMins * time.Minute)

	// Set claims
	claims := &model.AppClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expirationTime.Unix(),
			Issuer:    mpJWTIssuer,
		},
		RefreshAt: nowTime.Add(refMins * time.Minute).Unix(),
		UserEmail: email,
	}
	
	// Create a signer for RSA 256
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, signErr := tkn.SignedString(rsaPrivateKey)

	return tokenString, expirationTime, signErr
}

/**
	SetSecuredCookie helps set a cookie into 'http'
*/
func SetSecuredCookie(rw http.ResponseWriter, signedTokenStr string, tknExpiration time.Time) {
	// More info: https://tools.ietf.org/html/rfc6265
	cookie := http.Cookie{
		Name:     mpJWTCookieName,
		Path:     "/",
		Value:    signedTokenStr,
		Expires:  tknExpiration,
		HttpOnly: true, // Limits the scope of the cookie to HTTP requests.
		// TODO: set Secure to 'true' once HTTPS is being used.
		Secure: false, // User agent will include the cookie in an HTTP request
		//   only if the request is transmitted over a secure channel
		//   (typically HTTP over Transport Layer Security (TLS)
	}

	http.SetCookie(rw, &cookie)
}

package interfaces

import (
	"time"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"crypto/rsa"
	"errors"
	"path/filepath"
	"log"
	"io/ioutil"
)

const (
	// TODO: change these paths to a private path in production
	// NOTE: do NOT use these keys in production. These are for demo only.
	privKeyPath = "./src/mp.rsa"     // openssl genrsa -out mp.rsa 1024
	pubKeyPath  = "./src/mp.rsa.pub" // openssl rsa -in mp.rsa -pubout > mp.rsa.pub
)
const mpJWTCookieName = "musicPlaylistJWTAuth"

type AppClaims struct {
	jwt.StandardClaims
	RefreshAt int64  `json:"rat,omitempty"`
	UserEmail string `json:"email"`
}

type JWTHandler struct {
	rsaPrivateKey *rsa.PrivateKey
	rsaPublicKey  *rsa.PublicKey
}

func (handler *JWTHandler) ValidateUserEmail(rw http.ResponseWriter, userEmail string) error {
	nowTime := time.Now()
	expirationTime := nowTime.Add(15 * time.Minute) // TODO: change 15 minutes to something else

	claims := &AppClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "MPIssuer", // TODO: rename issuer
		},
		RefreshAt: nowTime.Add(7 * time.Minute).Unix(), // TODO: change 7 minutes to something else
		UserEmail: userEmail,
	}

	jwtSignedStr, signErr := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(handler.rsaPrivateKey)

	if signErr != nil {
		return signErr
	}

	handler.setSecureCookie(rw, jwtSignedStr, expirationTime)

	return nil
}

func (handler *JWTHandler) InvalidateUserEmail(rw http.ResponseWriter) {
	handler.unsetSecureCookie(rw)
}

func (handler *JWTHandler) ExtractJWTFromRequest(req *http.Request) (*jwt.Token, *AppClaims, error) {
	jwtCookie, cookieErr := req.Cookie(mpJWTCookieName)

	if cookieErr != nil {
		return nil, nil, cookieErr
	}

	jwtStr := jwtCookie.Value
	var claims AppClaims

	tkn, parseErr := jwt.ParseWithClaims(jwtStr, &claims, func(jwtPointer *jwt.Token) (interface{}, error) {
		if _, ok := jwtPointer.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return handler.rsaPublicKey, nil
	})

	return tkn, &claims, parseErr
}

func (handler *JWTHandler) setSecureCookie(rw http.ResponseWriter, signedJWTStr string, jwtExpiration time.Time) {
	secureCookie := http.Cookie{
		Name:     mpJWTCookieName,
		Path:     "/",             // TODO: change this to appropriate path
		Value:    signedJWTStr,
		Expires:  jwtExpiration,
		HttpOnly: true,
		Secure:   false, // TODO: change to true when using https
	}

	http.SetCookie(rw, &secureCookie)
}

func (handler *JWTHandler) unsetSecureCookie(rw http.ResponseWriter) {
	handler.setSecureCookie(rw, "", time.Now())
}

func NewJWTHandler() JWTHandler {
	fatalFn := func(err error, msg string) {
		if err != nil {
			log.Fatal(msg)
		}
	}

	getDataFromFileFn := func(path string) []byte {
		absPath, absErr := filepath.Abs(path)
		fatalFn(absErr, "filepath does not exist: " + path)

		fileBytes, readErr := ioutil.ReadFile(absPath)
		fatalFn(readErr, "failed to read file-path: " + absPath)

		return fileBytes
	}

	publicKey, pubKeyErr := jwt.ParseRSAPublicKeyFromPEM(getDataFromFileFn(pubKeyPath))
	fatalFn(pubKeyErr, "failed to parse public key from PEM")
	privKey, privKeyErr := jwt.ParseRSAPrivateKeyFromPEM(getDataFromFileFn(privKeyPath))
	fatalFn(privKeyErr, "failed to parse private key from PEM")

	jwtHandler := new(JWTHandler)
	jwtHandler.rsaPublicKey = publicKey
	jwtHandler.rsaPrivateKey = privKey

	return *jwtHandler
}
package middleware

import (
	"net/http"
	"testing"
	"net/http/httptest"
	"encoding/json"
	"crypto/rsa"
	"strconv"

	"github.com/stretchr/testify/assert"
	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/dtrinh100/Music-Playlist/src/api/model"
	"github.com/dgrijalva/jwt-go"
)

const pubTestRelPath = "../src/mp.rsa.pub"
const privTestRelPath = "../src/mp.rsa"

/**
	TestJWTMiddlewareWithInvalidJWTCookie tests the server's response if an invalid
	cookie is given (e.g. no cookie). If no cookie is passed to the server, it
	should be assumed that the user has not logged in yet.
	Testing Expectations:
		An error list with the following error:
		{"Login": "User Needs To Log In"}
*/
func TestJWTMiddlewareWithInvalidJWTCookie(t *testing.T) {
	asrt := assert.New(t)

	jwtMwareAH := AliceMiddlewareEnvHandler{Env: nil, AliceEnvFn: JWTMiddleware}
	server := httptest.NewServer(jwtMwareAH.Handle(GetTestHandler()))
	defer server.Close()

	resp, respErr := http.Get(server.URL)
	asrt.NoError(respErr)

	expected := common.ErrorList{Errors: common.ErrMap{"Login": "User Needs To Log In"}}
	result := common.ErrorList{}

	decodeErr := json.NewDecoder(resp.Body).Decode(&result)
	asrt.NoError(decodeErr)

	asrt.Equal(expected, result)
}

/**
	TestJWTMiddlewareWithExpiredJWT tests the response if an expired JWT is given.
	If an expired JWT is given, it should be assumed that the user's logged-in
	session has expired and the user needs to log in once again.
	Testing Expectations:
		An error list with the following error:
		{"Token": "Token Expired. Request A New One"}
*/
func TestJWTMiddlewareWithExpiredJWT(t *testing.T) {
	asrt := assert.New(t)
	pubKey, privKey := initRSAKeyPairHelper(pubTestRelPath, privTestRelPath)

	expJWTStr, _, jwtErr := getJWTHelper(privKey, "test@email.com", -1, -1)
	asrt.NoError(jwtErr)

	resp := getServerResponseWithJWTInfo(&common.RSAKeys{
		Public: pubKey, Private: privKey}, expJWTStr)

	expected := common.ErrorList{
		Errors: common.ErrMap{
			"Token": "Token Expired. Request A New One"},
	}
	result := common.ErrorList{}
	json.NewDecoder(resp.Body).Decode(&result)

	asrt.Equal(expected, result)
}

/**
	TestJWTMiddlewareWithInvalidJWT tests what happens if an malformed JWT is given.
	Testing Expectations:
		An error list with the following error:
		{"Internal Server Error": "Something Went Wrong In The API"}
*/
func TestJWTMiddlewareWithInvalidJWT(t *testing.T) {
	asrt := assert.New(t)
	pubKey, privKey := initRSAKeyPairHelper(pubTestRelPath, privTestRelPath)

	invalidTknString := "eyJhbGciOiJSPrL1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDUzMjk" +
		"0NzAsImlhdCI6MTUwNTMyOTQxMCwiaXNzIjoibXVzaWNQbGF5bGlz.oUrwTKTP-H0MjGOZKf"

	resp := getServerResponseWithJWTInfo(&common.RSAKeys{
		Public: pubKey, Private: privKey}, invalidTknString)

	expected := common.ErrorList{
		Errors: common.ErrMap{
			"Internal Server Error": "Something Went Wrong In The API"},
	}
	result := common.ErrorList{}
	json.NewDecoder(resp.Body).Decode(&result)

	asrt.Equal(expected, result)
}

/**
	TestJWTMiddlewareRefresh tests to see if the user is provided a new JWT once
	theirs is about to expire. This occurs when the following is true:
	jwt-refreshTime < now < jwt-expirationTime.
	Testing Expectations:
		Receive a new JWT in a cookie where the new JWT has an updated time for
		the jwt -refreshTime & -expirationTime.
*/
func TestJWTMiddlewareRefresh(t *testing.T) {
	asrt := assert.New(t)
	pubKey, privKey := initRSAKeyPairHelper(pubTestRelPath, privTestRelPath)
	jwtStr, _, jwtErr := getJWTHelper(privKey, "test@email.com", 1, -1)
	asrt.NoError(jwtErr)

	resp := getServerResponseWithJWTInfo(&common.RSAKeys{
		Public: pubKey, Private: privKey}, jwtStr)

	jwtNewStr := (resp.Cookies()[0]).Value

	// getClaims makes it convenient/cleaner to parse data into model.AppClaims
	getClaims := func(jwtString string) model.AppClaims {
		var claims model.AppClaims

		_, jwtParseErr := jwt.ParseWithClaims(jwtString, &claims,
			TokenFunc{RSAPublicKey: pubKey}.Verify)
		asrt.NoError(jwtParseErr)

		return claims
	}

	claims0 := getClaims(jwtStr)
	claims1 := getClaims(jwtNewStr)

	refresh0 := claims0.RefreshAt + 60 // offset: refMins is 1 minute in the past in jwtStr
	expire0 := claims0.ExpiresAt - 60  // offset: expMins is 1 minute in the future in jwtStr

	refLen := claims1.RefreshAt - refresh0
	expectedRefLen := int64(60 * jwtExpireMinutes / 2)
	refAtLeast := refLen >= expectedRefLen
	refAtMost := refLen < expectedRefLen+2

	expLen := claims1.ExpiresAt - expire0
	expectedExpLen := int64(60 * jwtExpireMinutes)
	expAtLeast := expLen >= expectedExpLen
	expAtMost := expLen < expectedExpLen+2

	errMsg := func(expectedSeconds, actualSeconds int64) string {
		return "len should be about " + strconv.FormatInt(expectedSeconds, 10) + " to " +
			strconv.FormatInt(expectedSeconds+2, 10) + ", not " +
			strconv.FormatInt(actualSeconds, 10) + " seconds"
	}

	asrt.True(refAtLeast && refAtMost, "Refresh "+errMsg(refLen, refLen))
	asrt.True(expAtLeast && expAtMost, "Expire "+errMsg(expLen, expLen))
}

/**
	TestJWTMiddlewareWithValidJWT tests what happens when a cookie containing a valid JWT
	is given.
	Testing Expectations:
		The call should reach the desired function and return a successful response:
		{"Response": "Success"}
*/
func TestJWTMiddlewareWithValidJWT(t *testing.T) {
	asrt := assert.New(t)

	pubKey, privKey := initRSAKeyPairHelper(pubTestRelPath, privTestRelPath)
	jwtStr, _, jwtErr := getJWTHelper(privKey, "test@email.com", 2, 1)
	asrt.NoError(jwtErr)

	resp := getServerResponseWithJWTInfo(&common.RSAKeys{
		Public: pubKey, Private: privKey}, jwtStr)

	expected := map[string]string{"Response": "Success"}
	result := map[string]string{}

	json.NewDecoder(resp.Body).Decode(&result)

	asrt.Equal(expected, result)
}

/**
	TestInitRSAKeyPairHelper tests whether the public/private RSA keys get loaded appropriately.
	Testing Expectations:
		Ensure that a *rsa.PublicKey & *rsa.PrivateKey are returned.
*/
func TestInitRSAKeyPairHelper(t *testing.T) {
	asrt := assert.New(t)
	pubKey, privKey := initRSAKeyPairHelper(pubTestRelPath, privTestRelPath)

	asrt.IsType(&rsa.PublicKey{}, pubKey)
	asrt.IsType(&rsa.PrivateKey{}, privKey)
}

/**
	TestInitRSAKeyPairHelper tests whether the public/private RSA keys get loaded appropriately.
	Testing Expectations:
		Ensure that signErr is nil.
*/
func TestGetJWTValid(t *testing.T) {
	asrt := assert.New(t)
	_, privKey := initRSAKeyPairHelper(pubTestRelPath, privTestRelPath)

	_, _, signErr := GetJWT(privKey, "test@email.com", 2)
	asrt.NoError(signErr)
}

/**
	getServerResponseWithJWTInfo is a helper function that allows for
	DRY/Clean -er test-code. It makes a simple call to GetTestHandler() while
	using the JWTMiddleware.
*/
func getServerResponseWithJWTInfo(rsKeys *common.RSAKeys, jwtStr string) *http.Response {
	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: mpJWTCookieName, Value: jwtStr})

	rw := httptest.NewRecorder()

	jwtMwareAH := AliceMiddlewareEnvHandler{Env: &common.Env{
		DB:      nil,
		RSAKeys: *rsKeys,
	}, AliceEnvFn: JWTMiddleware}

	jwtMwareAH.Handle(GetTestHandler()).ServeHTTP(rw, req)

	return rw.Result()
}

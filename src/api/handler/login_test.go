package handler

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"io/ioutil"
	"strings"
	"encoding/json"
	"github.com/dtrinh100/Music-Playlist/src/api/common"
)

/**
	TestLoginValid tests Login.
	Testing for valid response, given valid input.

	Testing Expectations:
		response.Status = "200 OK"
		responseWriter.Header().Get("Content-Type") = "application/json"
		response.Body = (see variable 'expectedBody' below)
*/
func TestLoginValid(t *testing.T) {
	asrt := assert.New(t)
	credentialsJson := `{
		"email": "user.one@email.com",
		"username": "user_one",
		"password": "somepassword"
	}`
	req := httptest.NewRequest("POST", "/api/auth", strings.NewReader(credentialsJson))
	req.Header.Set("Content-Type", "application/json")
	rw := httptest.NewRecorder()

	loginH := Handler{nil, Login}
	loginH.ServeHTTP(rw, req)

	resp := rw.Result()

	body, bodyErr := ioutil.ReadAll(resp.Body)
	asrt.NoError(bodyErr)

	expectedStatus := "200 OK"
	resultStatus := resp.Status

	expectedHeader := "application/json"
	resultHeader := rw.Header().Get("Content-Type")

	// Testing Expections: response.Body
	expectedBody := map[string]string{"email": "user.one@email.com", "username": "user_one"}
	resultBody := map[string]string{}
	unmErr := json.Unmarshal(body, &resultBody)

	asrt.NoError(unmErr)
	asrt.Equal(expectedHeader, resultHeader)
	asrt.Equal(expectedStatus, resultStatus)
	asrt.Equal(expectedBody, resultBody)
}

/**
	TestLoginInvalid tests Login.
	Testing for invalid response, given invalid input.

	Testing Expectations:
		response.Status = "500 Internal Server Error"
		responseWriter.Header().Get("Content-Type") = "application/json"
		response.Body = (see variable 'expectedBody' below)
*/
func TestLoginInvalid(t *testing.T) {
	asrt := assert.New(t)
	credentialsJson := ``
	req := httptest.NewRequest("POST", "/api/auth", strings.NewReader(credentialsJson))
	req.Header.Set("Content-Type", "application/json")
	rw := httptest.NewRecorder()

	loginH := Handler{nil, Login}
	loginH.ServeHTTP(rw, req)

	resp := rw.Result()

	body, bodyErr := ioutil.ReadAll(resp.Body)
	asrt.NoError(bodyErr)

	expectedStatus := "500 Internal Server Error"
	resultStatus := resp.Status

	expectedHeader := "application/json"
	resultHeader := rw.Header().Get("Content-Type")

	// Testing Expections: response.Body
	expectedBody := common.Str2MapStr{
		"errors": common.ErrMap{
			"Internal Server Error": "Something Went Wrong In The API",
		},
	}

	resultBody := common.Str2MapStr{}

	unmErr := json.Unmarshal(body, &resultBody)

	asrt.NoError(unmErr)
	asrt.Equal(expectedHeader, resultHeader)
	asrt.Equal(expectedStatus, resultStatus)
	asrt.Equal(expectedBody, resultBody)
}

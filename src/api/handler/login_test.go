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
	TestLogin_Valid tests Login.
	Testing for valid response, given valid input.

	Testing Expectations:
		response.Status = "200 OK"
		responseWriter.Header().Get("Content-Type") = "application/json"
		response.Body = (see variable 'expected_body' below)
*/
func TestLogin_Valid(t *testing.T) {
	assert := assert.New(t)
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

	body, body_err := ioutil.ReadAll(resp.Body)
	assert.NoError(body_err)

	expected_status := "200 OK"
	result_status := resp.Status

	expected_header := "application/json"
	result_header := rw.Header().Get("Content-Type")

	// Testing Expections: response.Body
	expected_body := map[string]string{"email": "user.one@email.com", "username": "user_one"}
	result_body := map[string]string{}
	unm_err := json.Unmarshal(body, &result_body)

	assert.NoError(unm_err)
	assert.Equal(expected_header, result_header)
	assert.Equal(expected_status, result_status)
	assert.Equal(expected_body, result_body)
}

/**
	TestLogin_Invalid tests Login.
	Testing for invalid response, given invalid input.

	Testing Expectations:
		response.Status = "500 Internal Server Error"
		responseWriter.Header().Get("Content-Type") = "application/json"
		response.Body = (see variable 'expected_body' below)
*/
func TestLogin_Invalid(t *testing.T) {
	assert := assert.New(t)
	credentialsJson := ``
	req := httptest.NewRequest("POST", "/api/auth", strings.NewReader(credentialsJson))
	req.Header.Set("Content-Type", "application/json")
	rw := httptest.NewRecorder()

	loginH := Handler{nil, Login}
	loginH.ServeHTTP(rw, req)

	resp := rw.Result()

	body, body_err := ioutil.ReadAll(resp.Body)
	assert.NoError(body_err)

	expected_status := "500 Internal Server Error"
	result_status := resp.Status

	expected_header := "application/json"
	result_header := rw.Header().Get("Content-Type")

	// Testing Expections: response.Body
	expected_body := common.Str2mapstr{
		"errors": common.ErrMap{
			"Internal Server Error": "Failed To Decode JSON",
		},
	}

	result_body := common.Str2mapstr{}

	unm_err := json.Unmarshal(body, &result_body)

	assert.NoError(unm_err)
	assert.Equal(expected_header, result_header)
	assert.Equal(expected_status, result_status)
	assert.Equal(expected_body, result_body)
}

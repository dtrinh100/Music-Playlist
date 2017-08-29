package common

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

/**
	MockMethodHelper utilizes JsonStdResponse & HandleErrorWithMap so that
	they may be tested indirectly.
*/
func MockMethodHelper(rw http.ResponseWriter, req *http.Request) {
	var data map[string]string

	dec_err := json.NewDecoder(req.Body).Decode(&data)
	if HandleErrorWithMap(rw, dec_err, ErrMap{
		"Internal Server Error": "Failed to decode JSON",
	}, http.StatusInternalServerError) {
		return
	}

	JsonStdResponse(data, rw)
}

/**
	TestJsonStdResponse tests JsonStdResponse through the MockMethodHelper method.
	Testing for a valid response, given valid JSON input.

	Testing Expectations:
		response.Status = "200 OK"
		responseWriter.Header().Get("Content-Type") = "application/json"
		response.Body = (see variable 'expected_body' below)
*/
func TestJsonStdResponse(t *testing.T) {
	assert := assert.New(t)
	dataJson := `{
		"field_one": "val_one",
		"field_two": "val_two"
	}`
	req := httptest.NewRequest("GET", "/fake/path", strings.NewReader(dataJson))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	MockMethodHelper(w, req)

	resp := w.Result()

	body, body_err := ioutil.ReadAll(resp.Body)
	assert.NoError(body_err)

	expected_status := "200 OK"
	result_status := resp.Status

	expected_header := "application/json"
	result_header := w.Header().Get("Content-Type")

	// Testing Expections: response.Body
	expected_body := map[string]string{"field_one": "val_one", "field_two": "val_two"}
	result_body := map[string]string{}

	unm_err := json.Unmarshal(body, &result_body)

	assert.NoError(unm_err)
	assert.Equal(expected_header, result_header)
	assert.Equal(expected_status, result_status)
	assert.Equal(expected_body, result_body)
}

/**
	TestHandleErrorWithMap tests HandleErrorWithMap through the MockMethodHelper method.
	Testing for an invalid response, given invalid input.

	Testing Expectations:
		response.Status = "500 Internal Server Error"
		responseWriter.Header().Get("Content-Type") = "application/json"
		response.Body = (see variable 'expected_body' below)
*/
func TestHandleErrorWithMap(t *testing.T) {
	assert := assert.New(t)
	dataString := `not json`
	req := httptest.NewRequest("GET", "/fake/path", strings.NewReader(dataString))
	req.Header.Set("Content-Type", "text/plain")

	w := httptest.NewRecorder()

	MockMethodHelper(w, req)

	resp := w.Result()

	body, body_err := ioutil.ReadAll(resp.Body)
	assert.NoError(body_err)

	expected_status := "500 Internal Server Error"
	result_status := resp.Status

	expected_header := "application/json"
	result_header := w.Header().Get("Content-Type")

	// Testing Expections: response.Body
	expected_body := Str2mapstr{
		"errors": ErrMap{
			"Internal Server Error": "Failed to decode JSON",
		},
	}

	result_body := Str2mapstr{}

	unm_err := json.Unmarshal(body, &result_body)

	assert.NoError(unm_err)
	assert.Equal(expected_header, result_header)
	assert.Equal(expected_status, result_status)
	assert.Equal(expected_body, result_body)
}

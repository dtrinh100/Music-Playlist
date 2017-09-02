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
	MockMethodHelper utilizes JSONStdResponse & JSONErrorResponse so that
	they may be tested indirectly.
*/
func MockMethodHelper(rw http.ResponseWriter, req *http.Request) {
	var data map[string]string

	decErr := json.NewDecoder(req.Body).Decode(&data)
	if decErr != nil {
		JSONErrorResponse(rw, ErrMap{
			"Internal Server Error": "Failed to decode JSON",
		}, http.StatusInternalServerError)
		return
	}

	JSONStdResponse(rw, data)
}

/**
	TestJSONStdResponse tests JSONStdResponse through the MockMethodHelper method.
	Testing for a valid response, given valid JSON input.

	Testing Expectations:
		response.Status = "200 OK"
		responseWriter.Header().Get("Content-Type") = "application/json"
		response.Body = (see variable 'expectedBody' below)
*/
func TestJSONStdResponse(t *testing.T) {
	asrt := assert.New(t)
	dataJSON := `{
		"field_one": "val_one",
		"field_two": "val_two"
	}`
	req := httptest.NewRequest("GET", "/fake/path", strings.NewReader(dataJSON))
	req.Header.Set("Content-Type", "application/json")

	rw := httptest.NewRecorder()

	MockMethodHelper(rw, req)

	resp := rw.Result()

	body, bodyErr := ioutil.ReadAll(resp.Body)
	asrt.NoError(bodyErr)

	expectedStatus := "200 OK"
	resultStatus := resp.Status

	expectedHeader := "application/json"
	resultHeader := rw.Header().Get("Content-Type")

	// Testing Expections: response.Body
	expectedBody := map[string]string{"field_one": "val_one", "field_two": "val_two"}
	resultBody := map[string]string{}

	unmErr := json.Unmarshal(body, &resultBody)

	asrt.NoError(unmErr)
	asrt.Equal(expectedHeader, resultHeader)
	asrt.Equal(expectedStatus, resultStatus)
	asrt.Equal(expectedBody, resultBody)
}

/**
	TestJSONErrorResponse tests JSONErrorResponse through the MockMethodHelper method.
	Testing for an invalid response, given invalid input.

	Testing Expectations:
		response.Status = "500 Internal Server Error"
		responseWriter.Header().Get("Content-Type") = "application/json"
		response.Body = (see variable 'expectedBody' below)
*/
func TestJSONErrorResponse(t *testing.T) {
	assert := assert.New(t)
	dataString := `not json`
	req := httptest.NewRequest("GET", "/fake/path", strings.NewReader(dataString))
	req.Header.Set("Content-Type", "text/plain")

	rw := httptest.NewRecorder()

	MockMethodHelper(rw, req)

	resp := rw.Result()

	body, bodyErr := ioutil.ReadAll(resp.Body)
	assert.NoError(bodyErr)

	expectedStatus := "500 Internal Server Error"
	resultStatus := resp.Status

	expectedHeader := "application/json"
	resultHeader := rw.Header().Get("Content-Type")

	// Testing Expections: response.Body
	expectedBody := Str2MapStr{
		"errors": ErrMap{
			"Internal Server Error": "Failed to decode JSON",
		},
	}

	resultBody := Str2MapStr{}

	unmErr := json.Unmarshal(body, &resultBody)

	assert.NoError(unmErr)
	assert.Equal(expectedHeader, resultHeader)
	assert.Equal(expectedStatus, resultStatus)
	assert.Equal(expectedBody, resultBody)
}

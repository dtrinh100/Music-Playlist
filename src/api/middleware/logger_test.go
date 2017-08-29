package middleware

import (
	"net/http"
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"

	"fmt"
)

/**
	GetTestHandler returns a http.HandlerFunc for testing http middleware.
*/
func GetTestHandler() http.HandlerFunc {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		// Do nothing
	}

	return http.HandlerFunc(fn)
}

/**
	TestLoggerMiddlewareSlashPath tests the Logger middleware & helper functions.

	Testing Expectations:
		Ensure that the log contains the string: '| [GET] "/" |'
*/
func TestLoggerMiddlewareSlashPath(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(Logger(GetTestHandler()))
	defer ts.Close()

	fn := func() *http.Response {
		res, res_err := http.Get(ts.URL)
		assert.NoError(res_err)
		return res
	}

	_, str := captureOutputExpectResponse(fn)
	fmt.Println(str)
	assert.Contains(str, "| [GET] \"/\" |")
}

/**
	TestLoggerMiddlewareLongPath tests the Logger middleware & helper functions.

	Testing Expectations:
		Ensure that the log contains the string: '| [POST] "/api/users/1" |'
*/
func TestLoggerMiddlewareLongPath(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(Logger(GetTestHandler()))
	defer ts.Close()

	fn := func() *http.Response {
		res, res_err := http.Post(ts.URL + "/api/users/1", "application/json", nil)
		assert.NoError(res_err)
		return res
	}

	_, str := captureOutputExpectResponse(fn)
	fmt.Println(str)
	assert.Contains(str, "| [POST] \"/api/users/1\" |")
}

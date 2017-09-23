package middleware

import (
	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
GetTestHandler returns a http.HandlerFunc for testing http middleware.
*/
func GetTestHandler() http.HandlerFunc {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		common.JSONStdResponse(rw, map[string]string{"Response": "Success"})
	}

	return http.HandlerFunc(fn)
}

/**
TestLoggerMiddlewareSlashPath tests the Logger middleware & helper functions.
Testing Expectations:
	Ensure that the log contains the string: '| [GET] "/" |'
*/
func TestLoggerMiddlewareSlashPath(t *testing.T) {
	asrt := assert.New(t)

	loggerMiddlewareAH := AliceMiddlewareHandler{AliceFn: LoggerMiddleware}

	server := httptest.NewServer(loggerMiddlewareAH.Handle(GetTestHandler()))
	defer server.Close()

	fn := func() *http.Response {
		resp, respErr := http.Get(server.URL)
		asrt.NoError(respErr)
		return resp
	}

	_, capturedStr := captureOutputExpectResponse(fn)

	asrt.Contains(capturedStr, "| [GET] \"/\" |")
}

/**
TestLoggerMiddlewareLongPath tests the Logger middleware & helper functions.
Testing Expectations:
	Ensure that the log contains the string: '| [POST] "/api/users/1" |'
*/
func TestLoggerMiddlewareLongPath(t *testing.T) {
	asrt := assert.New(t)

	LoggerMiddlewareAH := AliceMiddlewareHandler{AliceFn: LoggerMiddleware}

	server := httptest.NewServer(LoggerMiddlewareAH.Handle(GetTestHandler()))
	defer server.Close()

	fn := func() *http.Response {
		resp, respErr := http.Post(server.URL+"/api/users/1", "application/json", nil)
		asrt.NoError(respErr)
		return resp
	}

	_, capturedStr := captureOutputExpectResponse(fn)

	asrt.Contains(capturedStr, "| [POST] \"/api/users/1\" |")
}

package middleware

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"log"
	"net/http"
	"bytes"
	"os"
)

/**
	captureOutputExpectResponse captures log's output to a buffer and returns
	the output + an *http.Response. This will be used to test methods that
	output to the log.
*/
func captureOutputExpectResponse(f (func() *http.Response)) (*http.Response, string) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	res := f()
	log.SetOutput(os.Stderr)
	return res, buf.String()
}

/**
	TestCaptureOutputExpectResponse tests captureOutputExpectResponse by making sure
	that anything that is logged goes into a buffer and returned as a string.

	Testing Expectations:
		string-result: "Printing should be directed to a buffer by captureOutputExpectResponse."
*/
func TestCaptureOutputExpectResponse(t *testing.T) {
	assert := assert.New(t)

	expected := "Printing should be directed to a buffer by captureOutputExpectResponse."
	fn := func() *http.Response {
		log.Println(expected)
		return nil
	}

	_, result := captureOutputExpectResponse(fn)
	assert.Contains(result, expected)
}

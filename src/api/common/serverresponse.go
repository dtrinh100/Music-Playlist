package common

import (
	"encoding/json"
	"net/http"
)

/**
JSONResponse accepts some data-struct and converts it into JSON then sends it
to the client.
*/
func JSONResponse(rw http.ResponseWriter, dataMap interface{}, status int) {
	resp, respErr := json.Marshal(dataMap)

	if respErr != nil {
		// If you end up here, something really went wrong.
		http.Error(rw, "Error: Something Went Wrong In The API", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(status)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(resp)
}

/**
GenericJSONErrorResponse sends a generic, json-formatted error-response to the client.
*/
func JSONErrorResponse(rw http.ResponseWriter) {
	errMap := ErrMap{"Internal Server Error": "Something Went Wrong In The API"}
	JSONResponse(rw, ErrorList{Errors: errMap}, http.StatusInternalServerError)
}

/**
JSONStdResponse helps send json-formated 'standard' responses to the client.
*/
func JSONStdResponse(rw http.ResponseWriter, response interface{}) {
	JSONResponse(rw, map[string]interface{}{"data": response}, http.StatusOK)
}

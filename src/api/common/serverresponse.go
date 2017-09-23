package common

import (
	"encoding/json"
	"net/http"
)

/**
GenericJSONErrorResponse sends a generic, json-formatted error-response to the client.
*/
func GenericJSONErrorResponse(rw http.ResponseWriter) {
	JSONErrorResponse(rw, ErrMap{
		"Internal Server Error": "Something Went Wrong In The API",
	}, http.StatusInternalServerError)
}

// TODO: Look into if these functions can become middleware.

/**
JSONErrorResponse helps send json-formatted error-responses to the client.
*/
func JSONErrorResponse(rw http.ResponseWriter, errMap ErrMap, status int) {
	resp, respErr := json.Marshal(ErrorList{
		Errors: errMap,
	})

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
JSONStdResponse helps send json-formated standard-responses to the client.
*/
func JSONStdResponse(rw http.ResponseWriter, response interface{}) {
	json, jsonErr := json.Marshal(response)

	if jsonErr != nil {
		GenericJSONErrorResponse(rw)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(json)
}

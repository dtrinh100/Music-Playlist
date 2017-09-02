package common

import (
	"net/http"
	"encoding/json"
)

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
		http.Error(rw, respErr.Error(), http.StatusInternalServerError)
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
		JSONErrorResponse(rw, ErrMap{
			"Internal Server Error": "Failed to Marshal JSON",
		}, http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(json)
}

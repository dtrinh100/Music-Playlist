package common

import (
	"net/http"
	"encoding/json"
)

// TODO: Look into if these functions can become middleware.

/**
	JsonErrorResponse helps send json-formatted error-responses to the client.
*/
func JsonErrorResponse(rw http.ResponseWriter, errMap map[string]string, status int) {
	resp, err := json.Marshal(ErrorList{
		Errors: errMap,
	})

	if err != nil {
		// If you end up here, something really went wrong.
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(status)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(resp)
}

/**
	JsonStdResponse helps send json-formated standard-responses to the client.
*/
func JsonStdResponse(rw http.ResponseWriter, response interface{}) {
	json, err := json.Marshal(response)

	if err != nil {
		JsonErrorResponse(rw, ErrMap{
			"Internal Server Error": "Failed to Marshal JSON",
		}, http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(json)
}

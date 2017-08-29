package common

import (
	"net/http"
	"encoding/json"
	"log"
)

func JsonErrorResponse(errMap map[string]string, rw http.ResponseWriter, status int) {
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

func JsonStdResponse(response interface{}, rw http.ResponseWriter) {
	json, err := json.Marshal(response)

	if HandleErrorWithMap(rw, err, ErrMap{
		"Internal Server Error": "Failed to Marshal JSON",
	}, http.StatusInternalServerError) {
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(json)
}

func HandleErrorWithMap(rw http.ResponseWriter, err error, errMap ErrMap, httpStatus int) bool {
	errorOccurred := (err != nil)

	if errorOccurred {
		for k, v := range errMap {
			log.Println(k, v, "--", err)
		}

		JsonErrorResponse(errMap, rw, httpStatus)
	}

	return errorOccurred
}

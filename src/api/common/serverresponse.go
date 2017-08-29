package common

import (
	"net/http"
	"encoding/json"
	"log"
)

func JsonErrorResponse(errMap map[string]string, rw http.ResponseWriter, status int) {
	resp, err := json.Marshal(&ErrorList{
		Errors: errMap,
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(status)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(resp)
}

func JsonStdResponse(response interface{}, rw http.ResponseWriter) {
	json, err := json.Marshal(response)
	if HandleErrorWithMessage(rw, err, ErrMap{
		"JSONResponse": "Failed to parse"}, http.StatusInternalServerError) {
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(json)
}

func HandleErrorWithMessage(rw http.ResponseWriter, err error, msg ErrMap, httpStatus int) bool {
	errorOccurred := (err != nil)

	if errorOccurred {
		for k, v := range msg {
			log.Println(k, v, "--", err)
		}

		JsonErrorResponse(msg, rw, httpStatus)
	}

	return errorOccurred
}

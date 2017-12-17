package infrastructure

import (
	"encoding/json"
	"net/http"
	"github.com/dtrinh100/Music-Playlist/src/api/usecases"
)

type JSONWebResponder struct {
	WebResponder
}

func (responder *JSONWebResponder) jsonResponse(rw http.ResponseWriter, respMap usecases.M, statusCode int) {
	resp, respErr := json.Marshal(respMap)

	if respErr != nil {
		responder.InternalServerError(rw)
		return
	}

	rw.WriteHeader(statusCode)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(resp)
}

func (responder *JSONWebResponder) Success(rw http.ResponseWriter, respMap usecases.M) {
	responder.jsonResponse(rw, respMap, http.StatusOK)
}

func (responder *JSONWebResponder) Created(rw http.ResponseWriter, respMap usecases.M) {
	responder.jsonResponse(rw, respMap, http.StatusCreated)
}

func (responder *JSONWebResponder) BadRequest(rw http.ResponseWriter, err usecases.MPError) {
	responder.jsonResponse(rw, usecases.M{"error": err.Error()}, http.StatusBadRequest)
}

package infrastructure

import (
	"net/http"
	"github.com/dtrinh100/Music-Playlist/src/api/usecases"
	"encoding/json"
)

type WebResponder struct{}

func (responder *WebResponder) Success(rw http.ResponseWriter, respMap usecases.M) {
	rw.WriteHeader(http.StatusOK)
}

func (responder *WebResponder) Created(rw http.ResponseWriter, respMap usecases.M) {
	rw.WriteHeader(http.StatusCreated)
}

func (responder *WebResponder) BadRequest(rw http.ResponseWriter, err usecases.MPError) {
	rw.WriteHeader(http.StatusBadRequest)
}

func (responder *WebResponder) NoContent(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusNoContent)
}

func (responder *WebResponder) Redirection(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusPermanentRedirect)
}

func (responder *WebResponder) Unauthorized(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusUnauthorized)
}

func (responder *WebResponder) Forbidden(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusForbidden)
}

func (responder *WebResponder) NotFound(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusNotFound)
}

func (responder *WebResponder) Gone(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusGone)
}

func (responder *WebResponder) InternalServerError(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusInternalServerError)
}

func (responder *WebResponder) ServiceUnavailable(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusServiceUnavailable)
}

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

package handler

import "github.com/dtrinh100/Music-Playlist/src/api/common"

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Error gets the error string of the error
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Status gets the status code of the error
func (se StatusError) Status() int {
	return se.Code
}

// JSONError represents a StatusError with an ErrMap.
type JSONError struct {
	StatusError
	ErrMap common.ErrMap
}

// Set helps write cleaner code by allowing values to be passed through it's parameters.
func (je JSONError) Set(err error, errMap common.ErrMap, code int) JSONError {
	je.StatusError = StatusError{
		Code: code,
		Err:  err,
	}
	je.ErrMap = errMap
	return je
}

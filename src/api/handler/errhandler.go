package handler

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// JSONError represents a StatusError with a DataMap.
type JSONError struct {
	Code    int
	Err     error
	DataMap interface{}
}

// Error gets the error string of the error
func (je JSONError) Error() string {
	return je.Err.Error()
}

// Status gets the status code of the error
func (je JSONError) Status() int {
	return je.Code
}

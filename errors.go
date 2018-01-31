package grest

import "log"

// CustomError with custom text
func CustomError(text string) error {
	return &customErrorStruct{text}
}

type customErrorStruct struct {
	err string
}

func (e *customErrorStruct) Error() string {
	return e.err
}

// ErrHandler wraps Stackdriver to provide automatic logging of errors
// Should use logger by default or user can pass in his/her chosen logger as a third option
// Example usage:
func ErrHandler(err error, location string) {
	log.Println(err, location)

	return
}

package util

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Init initialize utility functions. App will exit if an error is occurred while initializing utilities
func Init() {
	// Init error codes
	initErrorCodes()
}

// An ApiError represent standard API error that contains HTTP Status (Status) and API-scoped Error Code (Code).
type ApiError struct {
	Status  int    `yaml:"status" json:"-"`
	Code    string `yaml:"code" json:"code"`
	Message string `yaml:"message" json:"message"`
}

// Error is an implementation of built-in error type interface
func (e ApiError) Error() string {
	return e.Message
}

// Error Codes from error_codes.yml
var errorCodes = make(map[string]ApiError)

// Errors
var ErrBadRequest, ErrInternalServer, ErrUserNotFound ApiError

// initErrorCodes load error codes from file. App will exit when an error occurred.
func initErrorCodes() {
	// Read file
	bytes, err := ioutil.ReadFile("errors_codes.yml")
	if err != nil {
		fmt.Printf("Unable to read error_codes.yml file. Error: %s\n", err.Error())
		os.Exit(5)
	}
	// Parse error codes file
	err = yaml.Unmarshal(bytes, &errorCodes)
	if err != nil {
		fmt.Printf("Unable to parse error_codes.yml file. Error: %s\n", err.Error())
		os.Exit(6)
	}
	// Init bad request
	ErrBadRequest = NewError("400")
	ErrInternalServer = NewError("500")
	ErrUserNotFound = NewError("-1002")
}

// NewError returns ApiError that is defined in error_codes.yml
func NewError(code string) ApiError {
	return errorCodes[code]
}

// CastError cast error interface as an ApiError
func CastError(err error) ApiError {
	apiErr, ok := err.(ApiError)
	if !ok {
		// If assert type fail, create new internal error
		apiErr = NewError("500")
	}
	return apiErr
}

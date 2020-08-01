// Package handler provides handler functions to handle request.
package handler

import (
	"html"
	"time"

	"encoding/json"
	"net/http"

	"../flags"
	"../loggers"
	"../response"
	"../util"
)

// Logger
var log = loggers.Get()

// RESTFunc is a handler function that handles error and writes response in JSON
type RESTFunc func(*http.Request) (*response.Success, error)

// ServeHTTP implement http.Handler interface to write success or error response in JSON
func (h RESTFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get execution time
	start := time.Now()
	// Init http status
	var httpStatus int
	// Execute handler
	result, err := h(r)
	// If an error returned, return error
	if err != nil {
		httpStatus = sendErrorJSON(w, err)
	} else if result == nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		// if header exist, add header to response
		if result.Header != nil {
			for k, v := range result.Header {
				w.Header().Set(k, v)
			}
		}
		// send json success
		httpStatus = sendJSON(w, http.StatusOK, result)
	}
	// Log elapsed time
	log.Infof("HTTP Status: %d, Request: %s %s, Time elapsed: %s", httpStatus, r.Method, html.EscapeString(r.URL.Path), time.Since(start))
}

// sendJSON write response in JSON
func sendJSON(w http.ResponseWriter, httpStatus int, obj interface{}) int {
	// Add content type
	w.Header().Add(flags.HeaderKeyContentType, flags.ContentTypeJSON)
	// Write http status
	w.WriteHeader(httpStatus)
	// Send JSON response
	json.NewEncoder(w).Encode(obj)
	// Return httpStatus
	return httpStatus
}

// sendErrorJSON write error response in JSON
func sendErrorJSON(w http.ResponseWriter, err error) int {
	// Cast error to ApiError
	apiError := util.CastError(err)
	// Send error json
	return sendJSON(w, apiError.Status, apiError)
}

// parseJSON parse json request body to o (target) and returns error
func parseJSON(r *http.Request, o interface{}) error {
	d := json.NewDecoder(r.Body)
	if err := d.Decode(o); err != nil {
		return err
	}
	return nil
}

// NotFound Returns not found in json
func NotFound(_ *http.Request) (*response.Success, error) {
	return nil, util.NewError("404")
}

// Unauthorized Returns not found in json
func MethodNotAllowed(_ *http.Request) (*response.Success, error) {
	return nil, util.NewError("404")
}

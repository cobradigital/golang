package middleware

import (
	"net/http"

	"../handler"
	"../response"
	"../services"
)

type Middleware func(handler.RESTFunc, ...string) handler.RESTFunc

func Auth() Middleware {
	// Create middleware
	m := func(next handler.RESTFunc, args ...string) handler.RESTFunc {
		// Define new handler
		h := func(r *http.Request) (*response.Success, error) {
			// Get purpose
			var purpose string
			if len(args) > 0 {
				purpose = args[0]
			}
			// Call Validate token service
			err := services.Auth.ValidateAccessToken(r, purpose)
			if err != nil {
				return nil, err
			}
			return next(r)
		}
		return h
	}
	// Return middleware
	return m
}

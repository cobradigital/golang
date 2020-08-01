package handler

import (
	"net/http"

	"../response"
)

func Home(r *http.Request) (*response.Success, error) {
	result := "test"
	return response.NewSuccess(result, nil), nil
}

package handler

import (
	"net/http"

	"../request"
	"../response"
	"../services"
)

func PostToken(r *http.Request) (*response.Success, error) {

	var req request.AuthRequest
	err := parseJSON(r, &req)
	if err != nil {
		return nil, err
	}

	result, err := services.Auth.NewToken(req.SecretId, req.SecretKey)
	if err != nil {
		return nil, err
	}
	return result, err
}

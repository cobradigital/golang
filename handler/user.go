package handler

import (
	"net/http"

	"../request"
	"../response"
	"../services"
)

func GetBalance(r *http.Request) (*response.Success, error) {
	userId := request.JWTSubject(r)

	result, _ := services.Account.Balance(userId)

	return response.NewSuccess(result, nil), nil
}

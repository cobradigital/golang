package services

import (
	"fmt"

	"../repositories"
)

type AccountService interface {
	// ExtractToken takes value of Authorization Header, validate expected token type and extract token value
	Balance(user_id string) (map[string]string, bool)
}

type accountService struct {
	userRepo repositories.UserRepository
}

func (s *accountService) Balance(user_id string) (map[string]string, bool) {

	response := make(map[string]string)

	status, _, _ := repositories.User.FindById("id = ?", user_id)

	if status.Premium > 0 {
		response["my_balance"] = "Unlimited"
		return response, true
	}

	balance, _, _ := repositories.User.BalanceById("user_id = ?", user_id)
	response["my_balance"] = fmt.Sprintf("%f", balance.Deposit)
	return response, true
}

package services

import (
	"../repositories"
)

// Services
var Auth AuthService
var Account AccountService

func Init() {
	Auth = &authService{repositories.User}
	Account = &accountService{repositories.User}
}

package models

import "time"

type User struct {
	Id        int64
	Username  string
	Email     string
	Password  string
	SecretId  string
	SecretKey string
	Status    int8
	Premium   int8
	CreatedAt time.Time
	UpdatedAt time.Time
}

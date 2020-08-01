package models

import "time"

type Deposit struct {
	Id       int
	User_Id  string
	Nominal  float64
	Datetime time.Time
}

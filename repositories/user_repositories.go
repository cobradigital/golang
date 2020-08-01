package repositories

import (
	"../models"
)

type UserRepository interface {
	Find(where string, args ...interface{}) ([]models.User, int, error)
	FindById(where string, args ...interface{}) (models.User, int, error)
	BalanceById(where string, args ...interface{}) (DepositSum, int, error)
}

type userRepository struct{}

type DepositSum struct {
	User_Id string
	Deposit float64
}

func initUserRepository() UserRepository {
	// Prepare statements
	var r userRepository
	return &r
}

func (s *userRepository) Find(where string, args ...interface{}) ([]models.User, int, error) {
	var data []models.User

	count := 0

	err := DBConnect.Where(where, args...).Find(&data).Count(&count).Error

	return data, count, err
}

func (s *userRepository) FindById(where string, args ...interface{}) (models.User, int, error) {
	var data models.User

	count := 0

	err := DBConnect.Where(where, args...).Find(&data).Count(&count).Error

	return data, count, err
}

func (s *userRepository) BalanceById(where string, args ...interface{}) (DepositSum, int, error) {
	var data DepositSum

	err := DBConnect.Where(where, args...).Table("deposit").Select("user_id, sum(nominal) as deposit").Group("user_id").Scan(&data).Error
	if err != nil {
		return data, 0, err
	}
	return data, 0, err
}

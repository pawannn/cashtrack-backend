package services

import (
	"fmt"
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/ports"
	"github.com/pawannn/cashtrack/internal/utils"
)

type UserService struct {
	UserRepo ports.UserRepo
}

func InitUserService(userRepo ports.UserRepo) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (uS *UserService) ValidatePhone(country string, phone string) error {
	return uS.UserRepo.ValidatePhone(country, phone)
}

func (uS *UserService) VerifyPhone(phone string, OTP string) (*models.User, error) {
	ok, err := uS.UserRepo.VerifyPhone(phone, OTP)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("invalid otp")
	}
	user := models.User{
		Id:         utils.NewUUID(),
		Phone:      phone,
		Name:       "",
		Currency:   "INR",
		IsVerified: true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	return &user, nil
}

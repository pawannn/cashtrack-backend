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

func (uS *UserService) ValidatePhone(phone string) error {
	return uS.UserRepo.ValidatePhone(phone)
}

func (uS *UserService) VerifyPhone(userDetails *models.User, OTP string) (*models.User, error) {
	ok, err := uS.UserRepo.VerifyPhone(userDetails.Phone, OTP)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("invalid otp")
	}

	user, err := uS.UserRepo.GetUserByPhone(userDetails.Phone)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, err
	}

	userDetails.Id = utils.NewUUID()
	userDetails.IsVerified = true
	userDetails.CreatedAt = time.Now()
	userDetails.UpdatedAt = time.Now()

	user, err = uS.UserRepo.Create(userDetails)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uS *UserService) UpdateUser(userData *models.User) (*models.User, error) {
	userData.UpdatedAt = time.Now()
	return uS.UserRepo.Update(userData)
}

func (uS *UserService) GetUserByID(id string) (*models.User, error) {
	return uS.UserRepo.GetUserByID(id)
}

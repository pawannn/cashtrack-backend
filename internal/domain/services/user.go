package services

import (
	"net/http"
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/ports"
	"github.com/pawannn/cashtrack/internal/utils"
)

type UserService struct {
	UserRepo ports.UserRepo
}

func (uS *UserService) GetUserByID(id string) (*models.User, utils.CashTrackError) {
	return uS.UserRepo.GetUserByID(id)
}

func InitUserService(userRepo ports.UserRepo) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (uS *UserService) ValidatePhone(phone string, country string) utils.CashTrackError {
	return uS.UserRepo.ValidatePhone(phone, country)
}

func (uS *UserService) VerifyPhone(userDetails *models.User, OTP string) (*models.User, utils.CashTrackError) {
	ok, err := uS.UserRepo.VerifyPhone(userDetails.Phone, userDetails.Country, OTP)
	if err != utils.NoErr {
		return nil, err
	}
	if !ok {
		return nil, utils.CashTrackError{
			Code:    http.StatusUnauthorized,
			Message: "Invalid OPT",
			Error:   nil,
		}
	}

	user, err := uS.UserRepo.GetUserByPhone(userDetails.Phone)
	if err != utils.NoErr {
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
	if err != utils.NoErr {
		return nil, err
	}

	return user, utils.NoErr
}

func (uS *UserService) UpdateUser(userData *models.User) (*models.User, utils.CashTrackError) {
	userData.UpdatedAt = time.Now()
	return uS.UserRepo.Update(userData)
}

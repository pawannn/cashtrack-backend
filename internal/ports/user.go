package ports

import (
	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

type UserRepo interface {
	ValidatePhone(phone string, country string) utils.CashTrackError
	VerifyPhone(phone string, country string, OTP string) (bool, utils.CashTrackError)
	Create(userDetails *models.User) (*models.User, utils.CashTrackError)
	Update(userDetails *models.User) (*models.User, utils.CashTrackError)
	GetUserByID(id string) (*models.User, utils.CashTrackError)
	GetUserByPhone(phone string) (*models.User, utils.CashTrackError)
}

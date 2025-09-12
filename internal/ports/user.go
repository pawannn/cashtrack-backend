package ports

import "github.com/pawannn/cashtrack/internal/domain/models"

type UserRepo interface {
	ValidatePhone(phone string) error
	VerifyPhone(phone string, OTP string) (bool, error)
	Create(userDetails *models.User) (*models.User, error)
	Update(userDetails *models.User) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUserByPhone(phone string) (*models.User, error)
}

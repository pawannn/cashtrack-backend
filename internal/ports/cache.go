package ports

import (
	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

type CacheRepo interface {
	GetUserInfo(userID string) (*models.User, utils.CashTrackError)
	StoreUserInfo(userDetails models.User) utils.CashTrackError
	GetCategories() ([]models.Category, utils.CashTrackError)
	SetCategories(categories []models.Category) utils.CashTrackError
}

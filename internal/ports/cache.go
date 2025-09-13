package ports

import "github.com/pawannn/cashtrack/internal/domain/models"

type CacheRepo interface {
	GetUserInfo(userID string) (*models.User, error)
	StoreUserInfo(userDetails models.User) error
	GetCategories() ([]models.Category, error)
	SetCategories(categories []models.Category) error
}

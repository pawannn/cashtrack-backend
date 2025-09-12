package ports

import "github.com/pawannn/cashtrack/internal/domain/models"

type CacheRepo interface {
	GetUserInfo(userID string) (*models.User, error)
	StoreUserInfo(userDetails models.User) error
	UpdateUserInfo(userDetails models.User) error
	DeleteUserInfo(userID string) error
	GetCategories() ([]models.Category, error)
	SetCategories() error
}

package ports

import "github.com/pawannn/cashtrack/internal/domain/models"

type CategoriesRepo interface {
	GetCategories() ([]models.Category, error)
	GetByID(cid string) (*models.Category, error)
	Create(category models.Category) (*models.Category, error)
	Update(category models.Category) (*models.Category, error)
	Delete(cid string) error
}

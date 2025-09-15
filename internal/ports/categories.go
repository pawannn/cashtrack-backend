package ports

import (
	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

type CategoriesRepo interface {
	GetCategories() ([]models.Category, utils.CashTrackError)
	// GetByID(cid string) (*models.Category, utils.CashTrackError)
	// Create(category models.Category) (*models.Category, utils.CashTrackError)
	// Update(category models.Category) (*models.Category, utils.CashTrackError)
	// Delete(cid string) utils.CashTrackError
}

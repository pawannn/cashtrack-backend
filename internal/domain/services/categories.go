package services

import (
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/ports"
	"github.com/pawannn/cashtrack/internal/utils"
)

type CategoriesService struct {
	CategoryRepo ports.CategoriesRepo
}

func InitCategoriesService(repo ports.CategoriesRepo) *CategoriesService {
	return &CategoriesService{
		CategoryRepo: repo,
	}
}

func (cS *CategoriesService) GetCategories() ([]models.Category, utils.CashTrackError) {
	return cS.CategoryRepo.GetCategories()
}

func (cS *CategoriesService) Create(category models.Category) (*models.Category, utils.CashTrackError) {
	category.Id = utils.NewUUID()
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	return cS.CategoryRepo.Create(category)
}

func (cS *CategoriesService) Update(category models.Category) (*models.Category, utils.CashTrackError) {
	category.UpdatedAt = time.Now()
	return cS.CategoryRepo.Update(category)
}

func (cS *CategoriesService) Delete(cid string) utils.CashTrackError {
	return cS.CategoryRepo.Delete(cid)
}

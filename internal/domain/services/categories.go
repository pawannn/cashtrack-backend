package services

import (
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

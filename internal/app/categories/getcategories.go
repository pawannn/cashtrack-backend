package categories

import (
	"net/http"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (cA *CategoriesApp) GetCategories() ([]models.Category, utils.CashTrackError) {
	categories, err := cA.cacheRepo.GetCategories()
	if err != utils.NoErr {
		return nil, err
	}

	if len(categories) > 0 {
		return categories, utils.NoErr
	}

	categories, err = cA.databaseRepo.GetCategories()
	if err != utils.NoErr {
		return nil, err
	}

	if len(categories) == 0 || categories == nil {
		return nil, utils.CashTrackError{
			Code:    http.StatusNotFound,
			Message: "Categories are empty",
			Error:   nil,
		}
	}

	if err := cA.cacheRepo.StoreCategories(categories); err != utils.NoErr {
		return nil, err
	}
	return categories, utils.NoErr
}

package cache

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
	"github.com/redis/go-redis/v9"
)

const categoriesKey = "cashtrack:categories"

func (rS *RedisService) GetCategories() ([]models.Category, utils.CashTrackError) {
	ctx := context.Background()

	res, err := rS.rClient.Get(ctx, categoriesKey).Result()
	if err == redis.Nil {
		return []models.Category{}, utils.NoErr
	} else if err != nil {
		return nil, utils.CashTrackError{
			Code:    http.StatusInternalServerError,
			Message: "Unable to get categories from cache",
			Error:   err,
		}
	}

	var categories []models.Category
	if err := json.Unmarshal([]byte(res), &categories); err != nil {
		return nil, utils.CashTrackError{
			Code:    http.StatusInternalServerError,
			Message: "Unable to parse categories from cache",
			Error:   err,
		}
	}

	return categories, utils.NoErr
}

func (rS *RedisService) SetCategories(categories []models.Category) utils.CashTrackError {
	ctx := context.Background()

	data, err := json.Marshal(categories)
	if err != nil {
		return utils.CashTrackError{
			Code:    http.StatusInternalServerError,
			Message: "Unable to encode categories",
			Error:   err,
		}
	}

	err = rS.rClient.Set(ctx, categoriesKey, data, time.Duration(time.Hour*10)).Err()
	if err != nil {
		return utils.CashTrackError{
			Code:    http.StatusInternalServerError,
			Message: "Unable to set categories in cache",
		}
	}

	return utils.NoErr
}

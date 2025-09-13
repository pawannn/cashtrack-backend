package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/redis/go-redis/v9"
)

const categoriesKey = "cashtrack:categories"

func (rS *RedisService) GetCategories() ([]models.Category, error) {
	ctx := context.Background()

	res, err := rS.rClient.Get(ctx, categoriesKey).Result()
	if err == redis.Nil {
		return []models.Category{}, nil
	} else if err != nil {
		return nil, err
	}

	var categories []models.Category
	if err := json.Unmarshal([]byte(res), &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func (rS *RedisService) SetCategories(categories []models.Category) error {
	ctx := context.Background()

	data, err := json.Marshal(categories)
	if err != nil {
		return err
	}

	return rS.rClient.Set(ctx, categoriesKey, data, time.Duration(time.Hour*10)).Err()
}

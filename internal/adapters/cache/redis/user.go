package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
)

func (rS *RedisService) GetUserInfo(userID string) (*models.User, error) {
	key := userID + ":cashTrack"
	res, err := rS.rClient.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	user := new(models.User)
	if err := json.Unmarshal([]byte(res), user); err != nil {
		return nil, err
	}
	return user, nil
}

func (rS *RedisService) StoreUserInfo(userDetails models.User) error {
	userID := userDetails.Id + ":cashTrack"
	userByte, err := json.Marshal(userDetails)
	if err != nil {
		return err
	}
	_, err = rS.rClient.Set(context.Background(), userID, userByte, time.Duration(time.Hour*1)).Result()
	return err
}

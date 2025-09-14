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

func (rS *RedisService) GetUserInfo(userID string) (*models.User, utils.CashTrackError) {
	key := userID + ":cashTrack"
	res, err := rS.rClient.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return nil, utils.NoErr
	}
	if err != nil {
		return nil, utils.CashTrackError{
			Code:    http.StatusInternalServerError,
			Message: "Unable to get user details from cache",
			Error:   err,
		}
	}
	user := new(models.User)
	if err := json.Unmarshal([]byte(res), user); err != nil {
		return nil, utils.CashTrackError{
			Code:    http.StatusInternalServerError,
			Message: "Unable to user details from cache",
			Error:   err,
		}
	}
	return user, utils.NoErr
}

func (rS *RedisService) StoreUserInfo(userDetails *models.User) utils.CashTrackError {
	userID := userDetails.Id + ":cashTrack"
	userByte, err := json.Marshal(userDetails)
	if err != nil {
		return utils.CashTrackError{
			Code:    http.StatusInternalServerError,
			Message: "Unable to encode user details in cache",
			Error:   err,
		}
	}
	_, err = rS.rClient.Set(context.Background(), userID, userByte, time.Duration(time.Hour*1)).Result()
	if err != nil {
		return utils.CashTrackError{
			Code:    http.StatusInternalServerError,
			Message: "Unable to set user details in cache",
		}
	}
	return utils.NoErr
}

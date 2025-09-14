package cache

import (
	"context"
	"time"

	"github.com/pawannn/cashtrack/internal/utils"
)

func (rS *RedisService) StoreOtpSentNumbers(phone string) utils.CashTrackError {
	ctx := context.Background()
	key := "otp_sent:" + phone

	err := rS.rClient.Set(ctx, key, true, 3*time.Minute).Err()
	if err != nil {
		return utils.CashTrackError{
			Code:    500,
			Message: "Failed to store phone in cache",
			Error:   err,
		}
	}
	return utils.NoErr
}

func (rS *RedisService) CheckOtpSentNumbers(phone string) (bool, utils.CashTrackError) {
	ctx := context.Background()
	key := "otp_sent:" + phone

	exists, err := rS.rClient.Exists(ctx, key).Result()
	if err != nil {
		return false, utils.CashTrackError{
			Code:    500,
			Message: "Failed to check phone in cache",
			Error:   err,
		}
	}

	return exists > 0, utils.NoErr
}

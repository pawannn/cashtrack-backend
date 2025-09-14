package user

import (
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (uA *UserApp) Update(userDetails *models.User) (*models.User, utils.CashTrackError) {
	userDetails.UpdatedAt = time.Now()
	user, err := uA.databaseRepo.UpdateUser(userDetails)
	if err != utils.NoErr {
		return nil, err
	}
	err = uA.cacheRepo.StoreUserInfo(user)
	if err != utils.NoErr {
		return nil, err
	}
	return user, utils.NoErr
}

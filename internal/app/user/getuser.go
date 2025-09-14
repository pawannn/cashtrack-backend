package user

import (
	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (uA *UserApp) GetUserByID(id string) (*models.User, utils.CashTrackError) {
	user, err := uA.databaseRepo.GetUserByID(id)
	if err != utils.NoErr {
		return nil, err
	}
	return user, utils.NoErr
}

func (uA *UserApp) GetUserByPhone(phone string) (*models.User, utils.CashTrackError) {
	user, err := uA.databaseRepo.GetUserByPhone(phone)
	if err != utils.NoErr {
		return nil, err
	}

	return user, utils.NoErr
}

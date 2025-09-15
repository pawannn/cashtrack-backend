package user

import (
	"net/http"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (uA *UserApp) GetUserByID(id string) (*models.User, utils.CashTrackError) {
	user, err := uA.cacheRepo.GetUserInfo(id)
	if err != utils.NoErr {
		return nil, err
	}
	if user != nil {
		return user, utils.NoErr
	}

	user, err = uA.databaseRepo.GetUserByID(id)
	if err != utils.NoErr {
		return nil, err
	}
	if user == nil {
		return nil, utils.CashTrackError{
			Code:    http.StatusNotFound,
			Message: "User is not registered with cashtrack",
			Error:   nil,
		}
	}

	err = uA.cacheRepo.StoreUserInfo(user)
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
	if user == nil {
		return nil, utils.CashTrackError{
			Code:    http.StatusNotFound,
			Message: "User is not registered with cashtrack",
			Error:   nil,
		}
	}

	return user, utils.NoErr
}

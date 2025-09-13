package user

import (
	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (uA *UserApp) GetUserByID(id string) (*models.User, utils.CashTrackError) {
	return nil, utils.NoErr
}

func (uA *UserApp) GetUserByPhone(phone string) (*models.User, utils.CashTrackError) {
	return nil, utils.NoErr
}

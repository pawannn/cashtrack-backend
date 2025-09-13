package user

import (
	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (uA *UserApp) Update(userDetails *models.User) (*models.User, utils.CashTrackError) {
	return nil, utils.NoErr
}

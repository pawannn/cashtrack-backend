package database

import (
	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (pS *PGService) CreateUser(user models.User) (*models.User, utils.CashTrackError) {
	return nil, utils.NoErr
}

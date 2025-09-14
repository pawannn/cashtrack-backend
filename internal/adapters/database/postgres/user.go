package database

import (
	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (pS *PGService) CreateUser(user models.User) (*models.User, utils.CashTrackError) {
	var insertedUser models.User
	query := `INSERT INTO users (id, phone, name, currency, country, is_verified, created_at, updated_at)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
			RETURNING id, phone, name, currency, country, is_verified, created_at, updated_at`
	err := pS.db.QueryRow(query, user.Id, user.Phone, user.Name, user.Currency, user.Country, user.IsVerified, user.CreatedAt, user.UpdatedAt).Scan(&insertedUser.Id, &insertedUser.Phone, &insertedUser.Name, &insertedUser.Currency, &insertedUser.Country, &insertedUser.IsVerified, &insertedUser.CreatedAt, &insertedUser.UpdatedAt)
	if err != nil {
		return nil, utils.CashTrackError{
			Code:    500,
			Message: "Unable to save user details",
			Error:   err,
		}
	}
	return &insertedUser, utils.NoErr
}

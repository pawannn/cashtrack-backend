package database

import (
	"fmt"
	"net/http"
	"strings"

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
		if strings.Contains(err.Error(), "duplicate") {
			return nil, utils.CashTrackError{
				Code:    http.StatusConflict,
				Message: "User already exist",
				Error:   err,
			}
		}
		return nil, utils.CashTrackError{
			Code:    500,
			Message: "Unable to save user details",
			Error:   err,
		}
	}
	return &insertedUser, utils.NoErr
}

func (pS *PGService) GetUserByID(userID string) (*models.User, utils.CashTrackError) {
	var user models.User
	query := `SELECT id, phone, name, currency, country, is_verified, created_at, updated_at FROM users where id = $1`
	err := pS.db.QueryRow(query, userID).Scan(&user.Id, &user.Phone, &user.Name, &user.Currency, &user.Country, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err.Error() == utils.ERR_NO_ROWS {
			return nil, utils.NoErr
		}
		return nil, utils.CashTrackError{
			Code:    500,
			Message: "Unable to get user details",
			Error:   err,
		}
	}
	return &user, utils.NoErr
}

func (pS *PGService) GetUserByPhone(phone string) (*models.User, utils.CashTrackError) {
	var user models.User
	query := `SELECT id, phone, name, currency, country, is_verified, created_at, updated_at FROM users where phone = $1`
	err := pS.db.QueryRow(query, phone).Scan(&user.Id, &user.Phone, &user.Name, &user.Currency, &user.Country, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err.Error() == utils.ERR_NO_ROWS {
			return nil, utils.NoErr
		}
		return nil, utils.CashTrackError{
			Code:    500,
			Message: "Unable to get user details",
			Error:   err,
		}
	}
	return &user, utils.NoErr
}

func (pS *PGService) UpdateUser(user *models.User) (*models.User, utils.CashTrackError) {
	fmt.Println(user)
	var updatedUser models.User
	query := `UPDATE users SET name = $1, updated_at = $2 WHERE id = $3  RETURNING id, phone, name, currency, country, is_verified, created_at, updated_at`
	err := pS.db.QueryRow(query, user.Name, user.UpdatedAt, user.Id).Scan(&updatedUser.Id, &updatedUser.Phone, &updatedUser.Name, &updatedUser.Currency, &updatedUser.Country, &updatedUser.IsVerified, &updatedUser.CreatedAt, &updatedUser.UpdatedAt)
	if err != nil {
		if err.Error() == utils.ERR_NO_ROWS {
			return nil, utils.CashTrackError{
				Code:    http.StatusNotFound,
				Message: "User does not exist",
				Error:   nil,
			}
		}
		return nil, utils.CashTrackError{
			Code:    500,
			Message: "Unable to update user details",
			Error:   err,
		}
	}
	return &updatedUser, utils.NoErr
}

package database

import (
	"database/sql"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (pS *PGService) RecordTransaction(transaction models.Transaction) (*models.Transaction, utils.CashTrackError) {
	query := `
		INSERT INTO transactions
		(id, user_id, category_id, amount, payment_method, date, note, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING id, user_id, category_id, amount, payment_method, date, note, created_at, updated_at
	`

	var updatedTransaction models.Transaction

	err := pS.db.QueryRow(
		query,
		transaction.Id,
		transaction.UserID,
		transaction.CategoryID,
		transaction.Amount,
		transaction.PaymentMethod,
		transaction.Date,
		transaction.Note,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	).Scan(
		&updatedTransaction.Id,
		&updatedTransaction.UserID,
		&updatedTransaction.CategoryID,
		&updatedTransaction.Amount,
		&updatedTransaction.PaymentMethod,
		&updatedTransaction.Date,
		&updatedTransaction.Note,
		&updatedTransaction.CreatedAt,
		&updatedTransaction.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.CashTrackError{
				Code:    404,
				Message: "Transaction not returned by database",
				Error:   err,
			}
		}
		return nil, utils.CashTrackError{
			Code:    500,
			Message: "Unable to store transaction in database",
			Error:   err,
		}
	}

	return &updatedTransaction, utils.CashTrackError{}
}

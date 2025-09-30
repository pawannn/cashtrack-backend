package database

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

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

	var tx models.Transaction

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
		&tx.Id,
		&tx.UserID,
		&tx.CategoryID,
		&tx.Amount,
		&tx.PaymentMethod,
		&tx.Date,
		&tx.Note,
		&tx.CreatedAt,
		&tx.UpdatedAt,
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

	return &tx, utils.CashTrackError{}
}

func (pS *PGService) FilterUserTransaction(userID string, from *time.Time, to *time.Time) ([]models.Transaction, utils.CashTrackError) {
	var transactions []models.Transaction
	var args []any
	var conditions []string

	conditions = append(conditions, "user_id = $1")
	args = append(args, userID)
	argPos := 2

	if from != nil {
		conditions = append(conditions, "date >= $"+strconv.Itoa(argPos))
		args = append(args, *from)
		argPos++
	}

	if to != nil {
		conditions = append(conditions, "date <= $"+strconv.Itoa(argPos))
		args = append(args, *to)
		argPos++
	}

	query := "SELECT id, user_id, category_id, amount, payment_method, date, note, created_at, updated_at FROM transactions"
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}
	query += " ORDER BY date DESC"

	rows, err := pS.db.Query(query, args...)
	if err != nil {
		return nil, utils.CashTrackError{
			Code:    500,
			Message: "Unable to fetch transactions",
			Error:   err,
		}
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(
			&t.Id,
			&t.UserID,
			&t.CategoryID,
			&t.Amount,
			&t.PaymentMethod,
			&t.Date,
			&t.Note,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		if err != nil {
			return nil, utils.CashTrackError{
				Code:    500,
				Message: "Error scanning transaction row",
				Error:   err,
			}
		}
		transactions = append(transactions, t)
	}

	return transactions, utils.CashTrackError{}
}

func (pS *PGService) UpdateTransaction(transaction models.Transaction) (*models.Transaction, utils.CashTrackError) {
	query := `
        UPDATE transactions
        SET category_id = $1, amount = $2, payment_method = $3, date = $4, note = $5, updated_at = $6
        WHERE id = $7 AND user_id = $8
        RETURNING id, user_id, category_id, amount, payment_method, date, note, created_at, updated_at
    `

	var tx models.Transaction

	err := pS.db.QueryRow(
		query,
		transaction.CategoryID,
		transaction.Amount,
		transaction.PaymentMethod,
		transaction.Date,
		transaction.Note,
		transaction.UpdatedAt,
		transaction.Id,
		transaction.UserID,
	).Scan(
		&tx.Id,
		&tx.UserID,
		&tx.CategoryID,
		&tx.Amount,
		&tx.PaymentMethod,
		&tx.Date,
		&tx.Note,
		&tx.CreatedAt,
		&tx.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.CashTrackError{
				Code:    404,
				Message: "Transaction not found or not updated",
				Error:   err,
			}
		}
		return nil, utils.CashTrackError{
			Code:    500,
			Message: "Unable to update transaction in database",
			Error:   err,
		}
	}

	return &tx, utils.NoErr
}

func (pS *PGService) DeleteTransaction(transaction models.Transaction) utils.CashTrackError {
	query := `
        DELETE FROM transactions
        WHERE id = $1 AND user_id = $2
    `

	result, err := pS.db.Exec(query, transaction.Id, transaction.UserID)
	if err != nil {
		return utils.CashTrackError{
			Code:    500,
			Message: "Unable to delete transaction from database",
			Error:   err,
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.CashTrackError{
			Code:    500,
			Message: "Unable to verify deletion of transaction",
			Error:   err,
		}
	}

	if rowsAffected == 0 {
		return utils.CashTrackError{
			Code:    404,
			Message: "Transaction not found or not deleted",
			Error:   sql.ErrNoRows,
		}
	}

	return utils.NoErr
}

func (pS *PGService) GetUserStats(userID string, from *time.Time, to *time.Time) ([]models.CategoryStat, utils.CashTrackError) {
	var stats []models.CategoryStat
	var args []any
	var conditions []string

	conditions = append(conditions, "t.user_id = $1")
	args = append(args, userID)
	argPos := 2

	if from != nil {
		conditions = append(conditions, "t.date >= $"+strconv.Itoa(argPos))
		args = append(args, *from)
		argPos++
	}

	if to != nil {
		conditions = append(conditions, "t.date <= $"+strconv.Itoa(argPos))
		args = append(args, *to)
		argPos++
	}

	query := `
        SELECT t.category_id, c.name, SUM(t.amount) as total_amount
        FROM transactions t
        JOIN categories c ON t.category_id = c.id
    `
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}
	query += " GROUP BY t.category_id, c.name ORDER BY total_amount DESC"

	rows, err := pS.db.Query(query, args...)
	if err != nil {
		return nil, utils.CashTrackError{
			Code:    500,
			Message: "Unable to fetch user statistics",
			Error:   err,
		}
	}
	defer rows.Close()

	for rows.Next() {
		var stat models.CategoryStat
		err := rows.Scan(
			&stat.CategoryID,
			&stat.CategoryName,
			&stat.TotalAmount,
		)
		if err != nil {
			return nil, utils.CashTrackError{
				Code:    500,
				Message: "Error scanning category statistics",
				Error:   err,
			}
		}
		stats = append(stats, stat)
	}

	if err = rows.Err(); err != nil {
		return nil, utils.CashTrackError{
			Code:    500,
			Message: "Error iterating over category statistics",
			Error:   err,
		}
	}

	return stats, utils.NoErr
}

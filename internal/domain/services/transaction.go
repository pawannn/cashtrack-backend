package services

import (
	"errors"
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/ports"
	"github.com/pawannn/cashtrack/internal/utils"
)

type TransactionService struct {
	TxRepo ports.TransactionRepo
}

func InitNewTransactionRepo(txRepo ports.TransactionRepo) *TransactionService {
	return &TransactionService{
		TxRepo: txRepo,
	}
}

func (tS *TransactionService) Record(tx models.Transaction) (*models.Transaction, error) {
	if tx.Amount == 0 {
		return nil, errors.New("amount cannot be zero")
	}
	tx.Id = utils.NewUUID()
	tx.Created_at = time.Now()
	tx.UpdatedAt = time.Now()
	return tS.TxRepo.Record(tx)
}

func (tS *TransactionService) FilterUserTransactions(userID string, from time.Time, to time.Time) ([]models.Transaction, error) {
	return tS.TxRepo.FilterUserTransactions(userID, from, to)
}

func (tS *TransactionService) GetByID(txID string) (*models.Transaction, error) {
	return tS.TxRepo.GetByID(txID)
}

func (tS *TransactionService) Update(tx *models.Transaction) (*models.Transaction, error) {
	tx.UpdatedAt = time.Now()
	return tS.TxRepo.Update(tx)
}

func (tS *TransactionService) Delete(txID string) error {
	return tS.TxRepo.Delete(txID)
}

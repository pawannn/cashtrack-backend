package services

import (
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

func (tS *TransactionService) Record(tx models.Transaction) (*models.Transaction, utils.CashTrackError) {
	tx.Id = utils.NewUUID()
	tx.CreatedAt = time.Now()
	tx.UpdatedAt = time.Now()
	return tS.TxRepo.Record(tx)
}

func (tS *TransactionService) FilterUserTransactions(userID string, from *time.Time, to *time.Time) ([]models.Transaction, utils.CashTrackError) {
	return tS.TxRepo.FilterUserTransactions(userID, from, to)
}

func (tS *TransactionService) UserStats(userID string, from *time.Time, to *time.Time) ([]models.CategoryStat, utils.CashTrackError) {
	return tS.TxRepo.UserStats(userID, from, to)
}

func (tS *TransactionService) Update(tx *models.Transaction) (*models.Transaction, utils.CashTrackError) {
	tx.UpdatedAt = time.Now()
	return tS.TxRepo.Update(tx)
}

func (tS *TransactionService) Delete(txID string) utils.CashTrackError {
	return tS.TxRepo.Delete(txID)
}

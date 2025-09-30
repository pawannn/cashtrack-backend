package transaction

import (
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (tA *TransactionApp) FilterUserTransactions(userID string, from *time.Time, to *time.Time) ([]models.Transaction, utils.CashTrackError) {
	tx, err := tA.databaseRepo.FilterUserTransaction(userID, from, to)
	if err != utils.NoErr {
		return nil, err
	}
	return tx, utils.NoErr
}

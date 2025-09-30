package transaction

import (
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (tA *TransactionApp) Update(tx *models.Transaction) (*models.Transaction, utils.CashTrackError) {
	tx.UpdatedAt = time.Now()

	txn, err := tA.databaseRepo.UpdateTransaction(*tx)
	if err != utils.NoErr {
		return nil, err
	}

	return txn, utils.NoErr
}

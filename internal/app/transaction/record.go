package transaction

import (
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (tA *TransactionApp) Record(tx *models.Transaction) (*models.Transaction, utils.CashTrackError) {
	tx.Id = utils.NewUUID()
	tx.CreatedAt = time.Now()
	tx.UpdatedAt = time.Now()

	txn, err := tA.databaseRepo.RecordTransaction(*tx)
	if err != utils.NoErr {
		return nil, err
	}

	return txn, utils.NoErr
}

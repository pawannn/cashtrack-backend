package transaction

import (
	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (tA *TransactionApp) Record(tx models.Transaction) (*models.Transaction, utils.CashTrackError) {
	return nil, utils.NoErr
}

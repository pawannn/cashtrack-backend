package transaction

import (
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (tA *TransactionApp) FilterUserTransactions(userID string, from *time.Time, to *time.Time) ([]models.Transaction, utils.CashTrackError) {
	return nil, utils.NoErr
}

package transaction

import (
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (tA *TransactionApp) UserStats(userID string, from *time.Time, to *time.Time) ([]models.CategoryStat, utils.CashTrackError) {
	return nil, utils.NoErr
}

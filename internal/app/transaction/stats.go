package transaction

import (
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (tA *TransactionApp) UserStats(userID string, from *time.Time, to *time.Time) ([]models.CategoryStat, utils.CashTrackError) {
	stats, err := tA.databaseRepo.GetUserStats(userID, from, to)
	if err != utils.NoErr {
		return nil, err
	}
	return stats, utils.NoErr
}

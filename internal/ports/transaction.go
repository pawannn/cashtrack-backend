package ports

import (
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

type TransactionRepo interface {
	Record(tx *models.Transaction) (*models.Transaction, utils.CashTrackError)
	Update(tx *models.Transaction) (*models.Transaction, utils.CashTrackError)
	Delete(tx *models.Transaction) utils.CashTrackError
	FilterUserTransactions(userID string, from *time.Time, to *time.Time) ([]models.Transaction, utils.CashTrackError)
	UserStats(userID string, from *time.Time, to *time.Time) ([]models.CategoryStat, utils.CashTrackError)
}

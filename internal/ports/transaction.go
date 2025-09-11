package ports

import (
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
)

type TransactionRepo interface {
	Record(tx models.Transaction) (*models.Transaction, error)
	FilterUserTransactions(userID string, from time.Time, to time.Time) ([]models.Transaction, error)
	GetByID(txID string) (*models.Transaction, error)
	Update(tx *models.Transaction) (*models.Transaction, error)
	Delete(txID string) error
}

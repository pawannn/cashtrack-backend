package ports

import (
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

type DatabaseRepo interface {
	// --- User operations ---
	CreateUser(user models.User) (*models.User, utils.CashTrackError)
	GetUserByID(userID string) (*models.User, utils.CashTrackError)
	GetUserByPhone(phone string) (*models.User, utils.CashTrackError)
	UpdateUser(user *models.User) (*models.User, utils.CashTrackError)

	// --- Category operations ---
	GetCategories() ([]models.Category, utils.CashTrackError)

	// --- Transaction operations ---
	RecordTransaction(transaction models.Transaction) (*models.Transaction, utils.CashTrackError)
	FilterUserTransaction(userID string, from *time.Time, to *time.Time) ([]models.Transaction, utils.CashTrackError)
	UpdateTransaction(transaction models.Transaction) (*models.Transaction, utils.CashTrackError)
	DeleteTransaction(transactionID models.Transaction) utils.CashTrackError

	// --- Stats & analytics ---
	GetUserStats(userID string, from *time.Time, to *time.Time) ([]models.CategoryStat, utils.CashTrackError)
}

package ports

import (
	"time"

	"github.com/pawannn/cashtrack/internal/domain/models"
)

type DatabaseRepo interface {
	// --- User operations ---
	CreateUser(user models.User) (*models.User, error)
	UpdateUser(user models.User) (*models.User, error)
	GetUserByID(userID string) (*models.User, error)
	GetUserByPhone(phone string) (*models.User, error)

	// --- Transaction operations ---
	RecordTransaction(transaction models.Transaction) (*models.Transaction, error)
	UpdateTransaction(transaction models.Transaction) (*models.Transaction, error)
	DeleteTransaction(transactionID string) error
	FilterUserTransaction(userID string, from *time.Time, to *time.Time) ([]models.Transaction, error)

	// --- Stats & analytics ---
	GetUserStats(userID string, from *time.Time, to *time.Time) ([]models.CategoryStat, error)

	// --- Category operations ---
	GetCategories() ([]models.Category, error)
	CreateCategory(category models.Category) (*models.Category, error)
	UpdateCategory(category models.Category) (*models.Category, error)
	DeleteCategory(categoryID string) error
}

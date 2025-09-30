package transaction

import "github.com/pawannn/cashtrack/internal/ports"

type TransactionApp struct {
	databaseRepo ports.DatabaseRepo
	cacheRepo    ports.CacheRepo
}

func InitTransactionApp(dbRepo ports.DatabaseRepo, cacheRepo ports.CacheRepo) ports.TransactionRepo {
	return &TransactionApp{
		databaseRepo: dbRepo,
		cacheRepo:    cacheRepo,
	}
}

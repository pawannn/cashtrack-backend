package user

import "github.com/pawannn/cashtrack/internal/ports"

type UserApp struct {
	databaseRepo ports.DatabaseRepo
	cacheRepo    ports.CacheRepo
	smsRepo      ports.SMSRepo
}

func InitUserApp(dbRepo ports.DatabaseRepo, cacheRepo ports.CacheRepo, smsRepo ports.SMSRepo) ports.UserRepo {
	return &UserApp{
		databaseRepo: dbRepo,
		cacheRepo:    cacheRepo,
		smsRepo:      smsRepo,
	}
}

package categories

import "github.com/pawannn/cashtrack/internal/ports"

type CategoriesApp struct {
	databaseRepo ports.DatabaseRepo
	cacheRepo    ports.CacheRepo
}

func InitCategoriesApp(dbRepo ports.DatabaseRepo, cacheRepo ports.CacheRepo) ports.CategoriesRepo {
	return &CategoriesApp{
		databaseRepo: dbRepo,
		cacheRepo:    cacheRepo,
	}
}

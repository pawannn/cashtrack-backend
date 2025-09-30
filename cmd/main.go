package main

import (
	"log"

	"github.com/pawannn/cashtrack/api/categories"
	"github.com/pawannn/cashtrack/api/transaction"
	"github.com/pawannn/cashtrack/api/user"
	auth "github.com/pawannn/cashtrack/internal/adapters/auth/jwt"
	cache "github.com/pawannn/cashtrack/internal/adapters/cache/redis"
	database "github.com/pawannn/cashtrack/internal/adapters/database/postgres"
	sms "github.com/pawannn/cashtrack/internal/adapters/sms/twillo"
	categoryApp "github.com/pawannn/cashtrack/internal/app/categories"
	transactionApp "github.com/pawannn/cashtrack/internal/app/transaction"
	userApp "github.com/pawannn/cashtrack/internal/app/user"
	"github.com/pawannn/cashtrack/internal/domain/services"
	"github.com/pawannn/cashtrack/internal/pkg/config"
	"github.com/pawannn/cashtrack/internal/pkg/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Unable to load config, %s", err.Error())
	}

	// Initialize Repositories
	dbRepo, err := database.InitPGService(cfg)
	if err != nil {
		log.Fatal(err)
	}
	cacheRepo := cache.InitRedisService(cfg)
	smsRepo := sms.InitTwilloClient(cfg)
	authRepo := auth.InitJWTService(cfg)
	userRepo := userApp.InitUserApp(dbRepo, cacheRepo, smsRepo)
	categoryRepo := categoryApp.InitCategoriesApp(dbRepo, cacheRepo)
	transactionRepo := transactionApp.InitTransactionApp(dbRepo, cacheRepo)

	// Initialize Services
	userService := services.InitUserService(userRepo)
	categoriesService := services.InitCategoriesService(categoryRepo)
	transactionService := services.InitNewTransactionRepo(transactionRepo)

	// Initialize http engine
	server := http.InitCashtrackEngine(cfg, authRepo)

	// Initialize Api Services
	categoriesApi := categories.InitCategoriesApi(server, categoriesService)
	userApi := user.InitUserApi(server, userService)
	transactionApi := transaction.InitTransactionApi(server, transactionService)

	// Add Routes to server
	userApi.AddRoutes()
	categoriesApi.AddRoutes()
	transactionApi.AddRoutes()
	// transactionApi.

	// Start the server
	if err := server.StartServer(); err != nil {
		log.Fatal("Unable to start server", err)
	}
}

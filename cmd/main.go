package main

import (
	"log"

	"github.com/pawannn/cashtrack/api/user"
	auth "github.com/pawannn/cashtrack/internal/adapters/auth/jwt"
	cache "github.com/pawannn/cashtrack/internal/adapters/cache/redis"
	database "github.com/pawannn/cashtrack/internal/adapters/database/postgres"
	sms "github.com/pawannn/cashtrack/internal/adapters/sms/twillo"
	useApp "github.com/pawannn/cashtrack/internal/app/user"
	"github.com/pawannn/cashtrack/internal/domain/services"
	"github.com/pawannn/cashtrack/internal/pkg/config"
	"github.com/pawannn/cashtrack/internal/pkg/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Unable to load config, %s", err.Error())
	}

	server := http.InitCashtrackEngine(cfg)

	cacheRepo := cache.InitRedisService(cfg)
	smsRepo := sms.InitTwilloClient(cfg)
	authRepo := auth.InitJWTService(cfg)
	dbRepo, err := database.InitPGService(cfg)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := useApp.InitUserApp(dbRepo, cacheRepo, smsRepo)
	userService := services.InitUserService(userRepo)
	userApi := user.InitUserApi(server, userService, &authRepo)
	userApi.AddRoutes()

	if err := server.StartServer(); err != nil {
		log.Fatal("Unable to start server", err)
	}
}

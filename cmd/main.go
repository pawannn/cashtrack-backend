package main

import (
	"log"

	"github.com/pawannn/cashtrack/internal/pkg/config"
	"github.com/pawannn/cashtrack/internal/pkg/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Unable to load config, %s", err.Error())
	}
	server := http.InitCashtrackEngine(cfg)
	if err := server.StartServer(); err != nil {
		log.Fatal("Unable to start server", err)
	}
}

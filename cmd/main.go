package main

import (
	"fmt"
	"log"

	"github.com/pawannn/cashtrack/internal/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Unable to load config, %s", err.Error())
	}
	fmt.Printf("%+v", cfg)
}

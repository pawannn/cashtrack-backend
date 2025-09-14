package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pawannn/cashtrack/internal/pkg/config"
	"github.com/pawannn/cashtrack/internal/ports"
)

type PGService struct {
	db *sql.DB
}

func InitPGService(config *config.CashTrackCfg) (ports.DatabaseRepo, error) {
	dbConnUri := fmt.Sprintf("dbname=%s host=%s port=%d user=%s password=%s sslmode=%s ", config.DBName, config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBSsl)
	db, err := sql.Open("postgres", dbConnUri)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PGService{
		db: db,
	}, nil
}

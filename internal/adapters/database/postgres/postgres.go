package database

import (
	_ "github.com/lib/pq"
)

// type PGService struct {
// 	db *sql.DB
// }

// func InitPGService(config config.CashTrackCfg) (ports.DatabaseRepo, error) {
// 	dbConnUri := fmt.Sprintf("dname=%s dbhost=%s dbport=%s user=%s pass=%s sslmode=%s", config.DBName, config.DBHost, config.CachePort, config.DBUser, config.CachePass, config.DBSsl)
// 	db, err := sql.Open("postgres", dbConnUri)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &PGService{
// 		db: db,
// 	}, nil
// }

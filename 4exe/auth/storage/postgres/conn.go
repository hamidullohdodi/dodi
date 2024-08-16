package postgres

import (
	"auth_service/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// ConnectDB function in postgres package
func ConnectDB() (*sql.DB, error) {
	cfg := config.Load()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME)
	
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

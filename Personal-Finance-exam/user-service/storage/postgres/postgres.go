package postgres

import (
	"database/sql"
	"fmt"

	"user/config"
	"user/storage"

	_ "github.com/lib/pq"
	"golang.org/x/exp/slog"
)

type Storage struct {
	Db    *sql.DB
	AuthS storage.AuthI
	UserS storage.UserI
}

func NewPostgresStorage(config config.Config) (*Storage, error) {
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		config.Database.Host,
		config.Database.User,
		config.Database.Name,
		config.Database.Password,
		config.Database.Port,
	)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	slog.Info("connected to db")

	return &Storage{
		Db:    db,
		AuthS: NewAuthRepo(db),
		UserS: NewUserRepo(db),
	}, nil
}

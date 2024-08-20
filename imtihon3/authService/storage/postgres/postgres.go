package postgres

import (
	"database/sql"
	"fmt"

	"auth_service/config"
	pb "auth_service/genproto"
	_ "github.com/lib/pq"
)

type Storagem struct {
	db *sql.DB
	pb.UnimplementedAuthServiceServer
}

func ConnectDB() (*sql.DB, error) {
	cfg := config.Load()
	conn := fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB_PORT, cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()

	return db, err
}

//func (s *Storagem) Auth() pb.AuthServiceServer {
//	if s.db== nil {
//		s.db = &UserAuthRepo{
//			db: s.db,
//		}
//	}
//	return s.db
//}
//
//func (s *Storagem) User() pb.AuthServiceServer {
//	if s.auth == nil {
//		s.auth = &UserAuthRepo{
//			db: s.db,
//		}
//	}
//	return s.auth
//}

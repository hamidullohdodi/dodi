package postgres

import (
	"database/sql"
	"fmt"

	auth "user/genproto/auth"
	_ "github.com/lib/pq"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) Register(req *auth.RegisterUserReq) (*auth.RegisterUserResp, error) {
	res := &auth.RegisterUserResp{}

	var id string
	query := `INSERT INTO users (first_name, email, password, last_name, date_of_birth, role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := r.db.QueryRow(query, req.FirstName, req.Email, req.Password, req.LastName, req.DateOfBirth, req.Role).Scan(&id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *AuthRepo) Login(req *auth.LoginUserReq) (*auth.User, error) {
	res := &auth.User{}

	var passwordHash string
	query := `SELECT id, first_name, email, password, role FROM users WHERE first_name = $1`
	err := r.db.QueryRow(query, req.FirstName).Scan(
		&res.Id,
		&res.FirstName,
		&res.Email,
		&passwordHash,
		&res.Role,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	} else if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *AuthRepo) RefreshTokenInsert(req *auth.RefreshTokenReq) (*auth.Empty, error) {
	res := &auth.Empty{}

	query := `INSERT INTO tokens (user_id, token) VALUES ($1, $2)`

	_, err := r.db.Exec(query, req.UserId, req.Token)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (r *AuthRepo) RefreshToken(req *auth.RefreshTokenReq) (*auth.RefreshTokenResp, error) {

	res := &auth.RefreshTokenResp{}

	query := `UPDATE tokens SET token = $2 WHERE user_id = $1`

	_, err := r.db.Exec(query, req.UserId, req.Token)
	if err != nil {
		return nil, err
	}

	return res, nil
}

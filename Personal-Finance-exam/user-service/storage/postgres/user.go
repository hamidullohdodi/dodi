package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	user "user/genproto/user"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) UpdateProfile(req *user.EditProfileReq) (*user.UserResp, error) {
	res := &user.UserResp{}

	query := `UPDATE users SET updated_at = NOW()`

	var arg []interface{}
	var conditions []string

	if req.FirstName != "" && req.FirstName != "string" {
		arg = append(arg, req.FirstName)
		conditions = append(conditions, fmt.Sprintf("first_name = $%d", len(arg)))
	}

	if req.Email != "" && req.Email != "string" {
		arg = append(arg, req.Email)
		conditions = append(conditions, fmt.Sprintf("email = $%d", len(arg)))
	}

	if req.LastName != "" && req.LastName != "string" {
		arg = append(arg, req.LastName)
		conditions = append(conditions, fmt.Sprintf("last_name = $%d", len(arg)))
	}

	if req.DateOfBirth != "" && req.DateOfBirth != "string" {
		arg = append(arg, req.DateOfBirth)
		conditions = append(conditions, fmt.Sprintf("date_of_birth = $%d", len(arg)))
	}

	if len(conditions) > 0 {
		query += ", " + strings.Join(conditions, ", ")
	}

	query += fmt.Sprintf(" WHERE id = $%d", len(arg)+1)
	arg = append(arg, req.Id)

	_, err := r.db.Exec(query, arg...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *UserRepo) GetProfile(req *user.UserId) (*user.UserResp, error) {
	res := &user.UserResp{}

	var date string
	query := `SELECT id, first_name, email, last_name, date_of_birth FROM users WHERE id = $1`
	err := r.db.QueryRow(query, req.Id).
		Scan(
			&res.Id,
			&res.FirstName,
			&res.Email,
			&res.LastName,
			&date,
		)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	res.DateOfBirth = date[:10]

	return res, nil
}

func (r *UserRepo) DeleteUser(req *user.UserId) (*user.Void, error) {
	res := &user.Void{}

	query := `UPDATE users SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1`
	_, err := r.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

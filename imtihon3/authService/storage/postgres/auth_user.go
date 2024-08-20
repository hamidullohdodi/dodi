package postgres

import (
	pb "auth_service/genproto"
	"auth_service/token"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
)

type UserAuthRepo struct {
	db *sql.DB
}

func NewUserAuthRepo(db *sql.DB) *UserAuthRepo {
	return &UserAuthRepo{
		db: db,
	}
}

func (r *UserAuthRepo) ResetPassword(req *pb.ByEmail) (*pb.Void, error) {
	return nil, nil
}

func (r *UserAuthRepo) RefreshToken(req *pb.RefreshTokenReq) (*pb.Token, error) {
	return nil, nil
}

func (r *UserAuthRepo) RegisterUser(req *pb.RegisterUserReq) (*pb.Register, error) {
	Id := uuid.New().String()
	query := `INSERT INTO users (
		id, 
		username,
		email, 
		password, 
		full_name,
		user_type,
		bio
	) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(
		query,
		Id,
		req.Username,
		req.Email,
		req.Password,
		req.FullName,
		req.UserType,
		req.Bio,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &pb.Register{}, nil
}

func (r *UserAuthRepo) LoginUser(req *pb.LoginReq) (*pb.Token, error) {
	var emailDB, passwordDB, user_id string
	query := "SELECT id, email, password FROM users WHERE email = $1"
	err := r.db.QueryRow(query, req.Email).Scan(&user_id, &emailDB, &passwordDB)
	if err != nil {
		return nil, err
	}
	fmt.Println(emailDB, passwordDB, user_id)
	qualify := true

	if passwordDB != req.Password || emailDB != req.Email {
		qualify = false
	}
	if !qualify {
		return nil, errors.New("email or password is incorrect")
	}
	tokens, err := token.GenereteJWTToken(user_id, req.GetEmail())
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (r *UserAuthRepo) GetUser(req *pb.ById) (*pb.Register, error) {
	query := "SELECT id, username, email, password, full_name, user_type, bio  FROM users WHERE id = $1"
	row := r.db.QueryRow(query, req.Id)
	res := &pb.Register{}
	err := row.Scan(&res.Id, &res.Username, &res.Email, &res.Password, &res.FullName, &res.UserType, &res.Bio)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *UserAuthRepo) UpdateUser(req *pb.UpdateUserReq) (*pb.Register, error) {
	var updates []string
	var args []interface{}
	argIdx := 1

	if req.Username != "" {
		updates = append(updates, fmt.Sprintf("username = $%d", argIdx))
		args = append(args, req.Username)
		argIdx++
	}
	if req.Email != "" {
		updates = append(updates, fmt.Sprintf("email = $%d", argIdx))
		args = append(args, req.Email)
		argIdx++
	}
	if req.Password != "" {
		updates = append(updates, fmt.Sprintf("password = $%d", argIdx))
		args = append(args, req.Password)
		argIdx++
	}
	if req.FullName != "" {
		updates = append(updates, fmt.Sprintf("full_name = $%d", argIdx))
		args = append(args, req.FullName)
		argIdx++
	}
	if req.UserType != "" {
		updates = append(updates, fmt.Sprintf("user_type = $%d", argIdx))
		args = append(args, req.UserType)
		argIdx++
	}
	if req.Bio != "" {
		updates = append(updates, fmt.Sprintf("bio = $%d", argIdx))
		args = append(args, req.Bio)
		argIdx++
	}

	if len(updates) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}
	UserId := "ed414140-dab8-4815-bd71-0b740ae1414e"

	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d AND deleted_at = 0", strings.Join(updates, ", "), argIdx)
	args = append(args, UserId)

	fmt.Println(query)
	fmt.Println(query)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return &pb.Register{}, nil
}

func (r *UserAuthRepo) UpdateUserType(req *pb.UserTypeReq) (*pb.UserTypeRes, error) {
	_, err := r.db.Exec(`UPDATE users SET user_type = $2 WHERE id = $1`, req.Id, req.UserType)
	if err != nil {
		return nil, err
	}

	var res pb.UserTypeRes
	err = r.db.QueryRow(`SELECT id, username, user_type, updated_at FROM users WHERE id = $1`, req.Id).Scan(
		&res.Id, &res.Username, &res.UserType, &res.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &res, nil
}

func (r *UserAuthRepo) GetAllUsers(req *pb.PageLimit) (*pb.GetAllUsersRes, error) {
	rows, err := r.db.Query(
		"SELECT id, username, full_name, user_type FROM users LIMIT $1 OFFSET $2;",
		req.Limit, req.Page,
	)
	fmt.Println(rows)
	if err != nil {
		return nil, err
	}

	var user pb.GetAllUsersRes
	err = rows.Scan(user.Id, user.Username, user.FullName, user.UserType)
	fmt.Println(rows)
	if err != nil {
		return nil, err
	}

	fmt.Println(rows)
	var total int32
	err = r.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&total)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *UserAuthRepo) DeleteUser(req *pb.ById) (*pb.Void, error) {
	query := `UPDATE users SET deleted_at = date_part('epoch', current_timestamp)::INT WHERE id = $1 AND deleted_at = 0`

	_, err := r.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

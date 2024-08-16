package postgres

import (
	pb "auth_service/genproto/user"
	"auth_service/model"
	"auth_service/storage"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type AuthService struct {
	Db *sql.DB
}

func NewAuthService(db *sql.DB) storage.StorageI {
	return &AuthService{
		Db: db,
	}
}

func (s *AuthService) Logout(ctx context.Context, req *model.LogoutRequest) (*model.LogoutResponse, error) {
	// Implement the Logout method here
	return &model.LogoutResponse{}, nil
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	query := `
	INSERT INTO users(first_name, last_name, email, password_hash, phone_number, role, created_at)
	values($1, $2, $3, $4, $5, $6, $7)
	returning id, first_name, last_name, phone_number, created_at
	`
	res := pb.RegisterResponse{}
	row := s.Db.QueryRow(query, req.FirstName, req.LastName, req.Email, req.Password, req.PhoneNumber, req.Role, time.Now()).Scan(&res.Id, &res.FirstName, &res.LastName, &res.PhoneNumber, &res.CreatedAt)

	if row != nil {
		return nil, row
	}

	return &res, nil

}

func (s *AuthService) Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error) {

	query := `select id, role, first_name from users where email = $1 and password_hash = $2 and deleted_at is null`

	row := s.Db.QueryRow(query, req.Email, req.Password)
	res := model.LoginResponse{}
	fmt.Println(req.Email, req.Password)

	err := row.Scan(&res.Id, &res.Role, &res.FirstName)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *AuthService) GetByIdProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {

	query := `SELECT id, first_name, last_name, email, phone_number, role, created_at FROM users WHERE id = $1 and deleted_at is null`

	res := pb.GetProfileResponse{}
	row := s.Db.QueryRow(query, req.Id)
	err := row.Scan(&res.Id, &res.FirstName, &res.LastName, &res.Email, &res.PhoneNumber, &res.Role, &res.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *AuthService) UpdateUserProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {

	query := `UPDATE users SET first_name = $1, phone_number = $2, role = $3, updated_at = $4 WHERE id = $5 and deleted_at is null`

	fmt.Println(req)
	_, err := s.Db.Exec(query, req.NewFirstName, req.NewPhoneNumber, req.NewRole, time.Now(), req.Id)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateProfileResponse{
		
	}, nil

}

func (s *AuthService) DeleteUserProfile(ctx context.Context, req *pb.DeleteProfileRequest) (*pb.DeleteProfileResponse, error) {

	query := `DELETE FROM users WHERE id = $1 and deleted_at is null`

	_, err := s.Db.ExecContext(ctx, query, req.Id)

	if err != nil {
		return nil, err
	}

	return &pb.DeleteProfileResponse{
		Message: "User deleted successfully",
	}, nil
}

func (s *AuthService) GetAllProfile(ctx context.Context, req *pb.GetProfilesRequest) (*pb.GetProfilesResponse, error) {

	query := `SELECT id, first_name, last_name, email, phone_number, role, created_at FROM users WHERE deleted_at is null`
	res := pb.GetProfilesResponse{}
	rows, err := s.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := pb.GetProfileResponse{}
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber, &user.Role, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		res.AllProfile = append(res.AllProfile, &user)
	}
	return &res, nil

}

func (s *AuthService) LogOut(ctx context.Context, req *model.LogoutRequest) (*model.LogoutResponse, error) {

	query := `UPDATE users SET refresh_token = $1, updated_at = $2 WHERE id = $3 and deleted_at is null`

	res := model.LogoutResponse{}
	err := s.Db.QueryRowContext(ctx, query, req.RefreshToken, time.Now()).Scan(&res.Message)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

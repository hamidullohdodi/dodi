package service

import (
	"auth_service/storage/postgres"
	"context"
	"database/sql"

	pb "auth_service/genproto"
)

type UserService struct {
	Repo *postgres.UserAuthRepo
	pb.UnimplementedAuthServiceServer
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		Repo: postgres.NewUserAuthRepo(db),
	}
}

func (s *UserService) RegisterUser(ctx context.Context, req *pb.RegisterUserReq) (*pb.Register, error) {
	return s.Repo.RegisterUser(req)
}

func (s *UserService) LoginUser(ctx context.Context, req *pb.LoginReq) (*pb.Token, error) {
	return s.Repo.LoginUser(req)
}

func (s *UserService) GetUser(ctx context.Context, req *pb.ById) (*pb.Register, error) {
	return s.Repo.GetUser(req)
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.Register, error) {
	return s.Repo.UpdateUser(req)
}

func (s *UserService) UpdateUserType(ctx context.Context, req *pb.UserTypeReq) (*pb.UserTypeRes, error) {
	return s.Repo.UpdateUserType(req)
}

func (s *UserService) GetAllUsers(ctx context.Context, req *pb.PageLimit) (*pb.GetAllUsersRes, error) {
	return s.Repo.GetAllUsers(req)
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	return s.Repo.DeleteUser(req)
}

func (s *UserService) ResetPassword(ctx context.Context, req *pb.ByEmail) (*pb.Void, error) {
	return s.Repo.ResetPassword(req)
}

func (s *UserService) RefreshToken(ctx context.Context, req *pb.RefreshTokenReq) (*pb.Token, error) {
	return s.Repo.RefreshToken(req)
}

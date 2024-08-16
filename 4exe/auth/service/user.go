package service

import (
	pb "auth_service/genproto/user"
	"auth_service/model"
	"auth_service/storage"
	"context"
	"log/slog"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	userRepo storage.StorageI
	log      *slog.Logger
}

type AuthServiceI interface {
	Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error)
	GetByIdProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error)
	UpdateUserProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error)
	DeleteUserProfile(ctx context.Context, req *pb.DeleteProfileRequest) (*pb.DeleteProfileResponse, error)
	GetAllProfile(ctx context.Context, req *pb.GetProfilesRequest) (*pb.GetProfilesResponse, error)
	LogOut(ctx context.Context, req *model.LogoutRequest) (*model.LogoutResponse, error)
}

func NewAuthService(userRepo storage.StorageI, log *slog.Logger) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		log:      log,
	}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	res, err := s.userRepo.Register(ctx, req)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	return res, nil
}

func (s *AuthService) Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error) {
	res, err := s.userRepo.Login(ctx, req)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	return res, nil
}

func (s *AuthService) GetByIdProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {

	res, err := s.userRepo.GetByIdProfile(ctx, req)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	return res, nil
}

func (s *AuthService) UpdateUserProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {

	res, err := s.userRepo.UpdateUserProfile(ctx, req)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	return res, nil
}

func (s *AuthService) DeleteUserProfile(ctx context.Context, req *pb.DeleteProfileRequest) (*pb.DeleteProfileResponse, error) {

	res, err := s.userRepo.DeleteUserProfile(ctx, req)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	return res, nil
}

func (s *AuthService) GetAllProfile(ctx context.Context, req *pb.GetProfilesRequest) (*pb.GetProfilesResponse, error) {

	res, err := s.userRepo.GetAllProfile(ctx, req)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	return res, nil
}

func (s *AuthService) LogOut(ctx context.Context, req *model.LogoutRequest) (*model.LogoutResponse, error) {

	res, err := s.userRepo.Logout(ctx, req)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	return res, nil
}

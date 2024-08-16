package storage

import (
	pb "auth_service/genproto/user"
	"auth_service/model"
	"context"
)

type StorageI interface {
	Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
	UpdateUserProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error)
	DeleteUserProfile(ctx context.Context, req *pb.DeleteProfileRequest) (*pb.DeleteProfileResponse, error)
	GetByIdProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error)
	GetAllProfile(ctx context.Context, req *pb.GetProfilesRequest) (*pb.GetProfilesResponse, error)
	Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error)
	Logout(ctx context.Context, req *model.LogoutRequest) (*model.LogoutResponse, error)
}

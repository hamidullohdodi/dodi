package storage

import (
	auth "user/genproto/auth"
	user "user/genproto/user"
)

type StorageI interface {
	Auth() AuthI
	User() UserI
}

type UserI interface {
	GetProfile(*user.UserId) (*user.UserResp, error)
	UpdateProfile(req *user.EditProfileReq) (*user.UserResp, error)
	DeleteUser(*user.UserId) (*user.Void, error)
}

type AuthI interface {
	Register(*auth.RegisterUserReq) (*auth.RegisterUserResp, error)
	Login(*auth.LoginUserReq) (*auth.User, error)
	RefreshToken(*auth.RefreshTokenReq)(*auth.RefreshTokenResp, error)
	RefreshTokenInsert(*auth.RefreshTokenReq)(*auth.Empty, error)
}

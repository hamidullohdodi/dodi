package service

import (
	auth "user/genproto/auth"
	st "user/storage/postgres"
	"context"
)

type AuthService struct {
	storage st.Storage
	auth.UnimplementedAuthServiceServer
}

func NewAuthService(storage *st.Storage) *AuthService {
	return &AuthService{
		storage: *storage,
	}
}

func (s *AuthService) Register(ctx context.Context, req *auth.RegisterUserReq) (*auth.RegisterUserResp, error) {
	res, err := s.storage.AuthS.Register(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AuthService) Login(ctx context.Context, req *auth.LoginUserReq) (*auth.User, error) {
	res, err := s.storage.AuthS.Login(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AuthService) RefreshTokenCreate(ctx context.Context, req *auth.RefreshTokenReq) (*auth.Empty, error) {
	res, err := s.storage.AuthS.RefreshTokenInsert(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// func (s *AuthService) ValidateToken(cxt context.Context, req *auth.ValidateTokenReq) (*auth.ValidateTokenRes, error) {
// 	res, err := s.storage.AuthS.ValidateToken(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

func (s *AuthService) RefreshToken(cxt context.Context, req *auth.RefreshTokenReq) (*auth.RefreshTokenResp, error) {
	res, err := s.storage.AuthS.RefreshToken(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

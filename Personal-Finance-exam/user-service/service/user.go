package service

import (
	"context"

	user "user/genproto/user"
	st "user/storage/postgres"
)

type UserService struct {
	storage st.Storage
	user.UnimplementedUserServiceServer
}

func NewUserService(storage *st.Storage) *UserService {
	return &UserService{
		storage: *storage,
	}
}

func (s *UserService) GetProfile(ctx context.Context, req *user.UserId) (*user.UserResp, error) {
	res, err := s.storage.UserS.GetProfile(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, req *user.EditProfileReq) (*user.UserResp, error) {
	res, err := s.storage.UserS.UpdateProfile(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *user.UserId) (*user.Void, error) {
	res, err := s.storage.UserS.DeleteUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

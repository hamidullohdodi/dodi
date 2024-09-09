package service

import (
	"auth-service/pkg/hashing"
	"auth-service/pkg/models"
	"auth-service/pkg/token"
	"auth-service/storage"
	"errors"
	"log"
	"log/slog"
)

type AuthService interface {
	Register(in models.RegisterRequest) (models.RegisterResponse, error)
	LoginEmail(in models.LoginEmailRequest) (models.Tokens, error)
	LoginUsername(in models.LoginUsernameRequest) (models.Tokens, error)
}

func NewAuthService(st storage.AuthStorage, logger *slog.Logger) AuthService {
	return &authService{st, logger}
}

type authService struct {
	st  storage.AuthStorage
	log *slog.Logger
}

func (a *authService) Register(in models.RegisterRequest) (models.RegisterResponse, error) {
	hash, err := hashing.HashPassword(in.Password)
	if err != nil {
		a.log.Error("Failed to hash password", "error", err)
		return models.RegisterResponse{}, err
	}

	in.Password = hash

	res, err := a.st.Register(in)
	if err != nil {
		a.log.Error("Failed to register user", "error", err)
		return models.RegisterResponse{}, err
	}

	return res, nil
}

func (a *authService) LoginEmail(in models.LoginEmailRequest) (models.Tokens, error) {
	res, err := a.st.LoginEmail(in)
	if err != nil {
		a.log.Error("Failed to login", "error", err)
		return models.Tokens{}, err
	}

	check := hashing.CheckPasswordHash(res.Password, in.Password)
	if !check {
		a.log.Error("Invalid password")
		return models.Tokens{}, errors.New("Invalid password")
	}

	refreshToken, err := token.GenerateRefreshToken(res)
	if err != nil {
		a.log.Error("Failed to generate refresh token", "error", err)
		return models.Tokens{}, err
	}

	accessToken, err := token.GenerateAccessToken(res)
	if err != nil {
		a.log.Error("Failed to generate access token", "error", err)
		return models.Tokens{}, err
	}

	response := models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (a *authService) LoginUsername(in models.LoginUsernameRequest) (models.Tokens, error) {
	res, err := a.st.LoginUsername(in)
	if err != nil {
		a.log.Error("Failed to login", "error", err)
		return models.Tokens{}, err
	}

	check := hashing.CheckPasswordHash(res.Password, in.Password)
	log.Println(check)
	if !check {
		a.log.Error("Invalid password")
		log.Println("\n----------", in.Password, res.Password, "\n---------")
		return models.Tokens{}, errors.New("Invalid password")
	}

	refreshToken, err := token.GenerateRefreshToken(res)
	if err != nil {
		a.log.Error("Failed to generate refresh token", "error", err)
		return models.Tokens{}, err
	}

	accessToken, err := token.GenerateAccessToken(res)
	if err != nil {
		a.log.Error("Failed to generate access token", "error", err)
		return models.Tokens{}, err
	}

	response := models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

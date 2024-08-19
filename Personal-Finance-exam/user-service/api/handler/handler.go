package handler

import (
	"user/service"
)

type Handlers struct {
	Auth *service.AuthService
}

func NewHandler(auth *service.AuthService) *Handlers {
	return &Handlers{
		Auth: auth,
	}
}

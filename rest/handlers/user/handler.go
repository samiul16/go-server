package user

import (
	"go-server/repo"
	"go-server/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	userRepo repo.UserRepo
}

func NewHandler(middlewares *middleware.Middlewares, userRepo repo.UserRepo) *Handler {
	return &Handler{
		middlewares: middlewares,
		userRepo: userRepo,
	}
}
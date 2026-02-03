package product

import (
	"go-server/repo"
	"go-server/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	porductRepo repo.ProductRepo
}

func NewHandler(middlewares *middleware.Middlewares, porductRepo repo.ProductRepo) *Handler {

	return &Handler{
		middlewares: middlewares,
		porductRepo: porductRepo,
	}
}
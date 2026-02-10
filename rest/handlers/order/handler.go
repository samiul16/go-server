package order

import (
	"fmt"
	"go-server/rest/middleware"
	"net/http"
)

type Handler struct {
	middleware *middleware.Middlewares
	svc Service
}

func NewHandler(middleware *middleware.Middlewares, svc Service) *Handler {
	return &Handler{
		middleware: middleware,
		svc: svc,
	}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Order")
}
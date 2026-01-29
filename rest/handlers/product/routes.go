package product

import (
	"go-server/rest/middleware"
	"net/http"
)


func  (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProducts), h.middlewares.Authenticate))
	mux.Handle("GET /product/{id}", manager.With(http.HandlerFunc(h.GetProductById)))
}
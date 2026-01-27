package rest

import (
	"go-server/rest/handlers"
	"go-server/rest/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /test", manager.With(http.HandlerFunc(handlers.Test)) )
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProducts)))
	mux.Handle("GET /product/{id}", manager.With(http.HandlerFunc(handlers.GetProductById)))

	mux.Handle("POST /users", manager.With(http.HandlerFunc(handlers.CreateUser)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(handlers.CreateProducts)))
}
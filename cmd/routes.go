package cmd

import (
	"go-server/handlers"
	"go-server/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /test", manager.With(http.HandlerFunc(handlers.Test)) )
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProducts)))
	mux.Handle("GET /product/{id}", manager.With(http.HandlerFunc(handlers.GetProductById)))
}
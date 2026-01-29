package cmd

import (
	"go-server/config"
	"go-server/rest"
	"go-server/rest/handlers/product"
	"go-server/rest/handlers/user"
	"go-server/rest/middleware"
)

func Server() {
	configs := config.GetConfig()
	
	middlewares := middleware.NewMiddleware(configs)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler(middlewares)
	
	server := rest.NewServer(configs, productHandler, userHandler)

	server.Start()
}
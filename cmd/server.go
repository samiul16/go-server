package cmd

import (
	"fmt"
	"go-server/config"
	"go-server/infra/db"
	"go-server/product"
	"go-server/repo"
	"go-server/rest"
	productHandler "go-server/rest/handlers/product"
	userHandler "go-server/rest/handlers/user"
	"go-server/rest/middleware"
	"go-server/user"
	"os"
)

func Server() {
	configs := config.GetConfig()

	db, err := db.NewConnection(configs.DBConfig)

	if err != nil {
		fmt.Println("Error in db connection", err)
		os.Exit(1)
	}
	
	middlewares := middleware.NewMiddleware(configs)

   productRepo := repo.NewProductRepo(db)
   userRepo := repo.NewUserRepo(db)

   userService := user.NewUserService(userRepo)
	productService := product.NewService(productRepo)

	productHandler := productHandler.NewHandler(middlewares, productService)
	userHandler := userHandler.NewHandler(middlewares, userService)
	
	server := rest.NewServer(configs, productHandler, userHandler)

	server.Start()
}
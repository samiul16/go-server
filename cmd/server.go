package cmd

import (
	"fmt"
	"go-server/config"
	"go-server/infra/db"
	"go-server/repo"
	"go-server/rest"
	"go-server/rest/handlers/product"
	"go-server/rest/handlers/user"
	"go-server/rest/middleware"
	"os"
)

func Server() {
	configs := config.GetConfig()

	db, err := db.NewConnection()

	if err != nil {
		fmt.Println("Error in db connection", err)
		os.Exit(1)
	}
	
	middlewares := middleware.NewMiddleware(configs)

   productRepo := repo.NewProductRepo(db)
   userRepo := repo.NewUserRepo(db)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(middlewares, userRepo)
	
	server := rest.NewServer(configs, productHandler, userHandler)

	server.Start()
}
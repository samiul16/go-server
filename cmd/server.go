package cmd

import (
	"fmt"
	"go-server/middleware"
	"net/http"
)

func Server() {
	fmt.Println("Server starting")
	manager := middleware.NewManger()

	mux := http.NewServeMux()

	var globalMiddlewares []middleware.Middleware

	globalMiddlewares = append(globalMiddlewares, middleware.Logger, middleware.Preflight, middleware.Cors)

	WrapedMux := manager.WrapMux(globalMiddlewares, mux)

	initRoutes(mux, manager)

	err := http.ListenAndServe(":8080", WrapedMux)

	if err != nil {
		fmt.Println("Error on starting server", err)
	}
}
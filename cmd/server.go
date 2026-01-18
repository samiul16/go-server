package cmd

import (
	"fmt"
	"go-server/middleware"
	"net/http"
)

func Server() {
	fmt.Println("Server starting")
	manager := middleware.NewManger()

	manager.Use(middleware.Logger, middleware.Hudai, middleware.CorsWithPreflight)
	mux := http.NewServeMux()

	initRoutes(mux, manager)

	muxWithMiddleware := manager.With(mux)

	err := http.ListenAndServe(":8080", muxWithMiddleware)

	if err != nil {
		fmt.Println("Error on starting server", err)
	}
}
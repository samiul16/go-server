package cmd

import (
	"fmt"
	"go-server/config"
	"go-server/middleware"
	"net/http"
	"strconv"
)

func Server() {
	configs := config.GetConfig()
	manager := middleware.NewManger()

	mux := http.NewServeMux()

	var globalMiddlewares []middleware.Middleware

	globalMiddlewares = append(globalMiddlewares, middleware.Logger, middleware.Preflight, middleware.Cors)

	WrapedMux := manager.WrapMux(globalMiddlewares, mux)

	initRoutes(mux, manager)

	port := ":" + strconv.Itoa(configs.HttpPort)

	err := http.ListenAndServe(port, WrapedMux)

	if err != nil {
		fmt.Println("Error on starting server", err)
	}
}
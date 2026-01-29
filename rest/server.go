package rest

import (
	"fmt"
	"go-server/config"
	"go-server/rest/handlers/product"
	"go-server/rest/handlers/user"
	"go-server/rest/middleware"
	"net/http"
	"strconv"
)

type Server struct {
	configs *config.Config
	productHandler *product.Handler // Dependency
	userHandler *user.Handler // Dependency
}

func NewServer(configs *config.Config, productHandler *product.Handler, userHandler *user.Handler) *Server {
	return &Server{
		configs: config.GetConfig(),
		productHandler: productHandler,
		userHandler: userHandler,
	}
}

func (server *Server) Start() {
	
	manager := middleware.NewManger()

	mux := http.NewServeMux()

	var globalMiddlewares []middleware.Middleware

	globalMiddlewares = append(globalMiddlewares, middleware.Logger, middleware.Preflight, middleware.Cors)

	WrapedMux := manager.WrapMux(globalMiddlewares, mux)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	port := ":" + strconv.Itoa(server.configs.HttpPort)

	err := http.ListenAndServe(port, WrapedMux)

	if err != nil {
		fmt.Println("Error on starting server", err)
	}
}
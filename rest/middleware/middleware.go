package middleware

import "go-server/config"

type Middlewares struct {
	configs *config.Config //dependency
}

func NewMiddleware(configs *config.Config) *Middlewares {
	return &Middlewares{
		configs: configs,
	}
}

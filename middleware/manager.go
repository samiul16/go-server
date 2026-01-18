package middleware

import (
	"log"
	"net/http"
)

type Middleware func(next http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManger() *Manager {
 manager := Manager{
	globalMiddlewares: make([]Middleware, 0),
 }
 return &manager
}

func (mangr *Manager) Use(middlewares ...Middleware) {
	mangr.globalMiddlewares = append(mangr.globalMiddlewares, middlewares...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler{
	
		n := next

		for i := len(middlewares) - 1; i >=0; i -- {
			middleware := middlewares[i]
			n = middleware(n) // calling this func, returns another func -> handler func
		}

		log.Println("global middleware:", mngr.globalMiddlewares)

		for _, middleware := range mngr.globalMiddlewares {
			n = middleware(n)
		}

		log.Println("n",n)

		return n
}
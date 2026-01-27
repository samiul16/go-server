package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println("Startof Logger middleware")
		startTime := time.Now()
		log.Println("Middle of Logger middleware")
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time.Since(startTime))

	})
}
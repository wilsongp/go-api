package middleware

import (
	"log"
	"net/http"
	"time"
)

//Logger middleware to log http requests
func Logger(nextHandler http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		nextHandler.ServeHTTP(w, r)

		log.Printf(
			"\n%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

//AsyncLogger will execute the given function in a go routine and immediately call next.ServeHTTP
func AsyncLogger(nextHandler http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		go nextHandler.ServeHTTP(w, r)

		log.Printf(
			"\n%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
		)

	})
}

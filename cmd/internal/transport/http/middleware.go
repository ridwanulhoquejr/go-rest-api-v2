package http

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// set the header to application/json
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		// call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log the request

		log.WithFields(log.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
		}).Info("handled request")

		// call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// TimeoutMiddleware is a middleware that times out the request after 15 seconds
// if it takes longer than that to process, it will return a 500
func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

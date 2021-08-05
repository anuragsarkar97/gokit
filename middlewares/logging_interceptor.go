package middlewares

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msgf("Requested %s with method %s", r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
	})
}

package middlewares

import (
	"github.com/google/uuid"
	"net/http"
)

func RequestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Request-Id") == "" {
			r.Header.Set("X-Request-Id", generateRequestId())
		}
		next.ServeHTTP(w, r)
	})
}

func generateRequestId() string {
	return uuid.New().String()
}

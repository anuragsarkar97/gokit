package middlewares

import "net/http"

func GenericMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// does nothing
		// can be used for pre-flight debugging
		next.ServeHTTP(w, r)
	})
}

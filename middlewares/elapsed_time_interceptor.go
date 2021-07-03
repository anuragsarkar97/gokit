package middlewares

import (
	"net/http"
	"time"
)

func ElapsedTimeInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		w.Header().Add("x-start-time", startTime.String())
		defer func() {
			elapsed := time.Now().Sub(startTime)
			w.Header().Add("x-time-used", elapsed.String())
		}()
		next.ServeHTTP(w, r)
	})
}

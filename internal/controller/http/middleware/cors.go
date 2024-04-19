package middleware

import "net/http"

func GetCorsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")

			next.ServeHTTP(w, r)
		})
}

func GetCorsMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return GetCorsHandler(next)
	}
}

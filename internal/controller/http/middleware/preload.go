package middleware

import "net/http"

func GretPreloadMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			//w.Header().Set("Link", "</static/shared.bundle.js>; rel=preload; as=script")

			next.ServeHTTP(w, r)
		})
}

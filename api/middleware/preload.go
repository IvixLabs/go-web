package middleware

import "net/http"

func GretPreloadMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Link", "</static/main.js>; rel=preload; as=script")

		next(w, r)
	}
}

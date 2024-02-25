package middleware

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	context3 "ivixlabs.com/proj5/internal/http/context"
	"net/http"
)

func GetContextMiddleware(sessionStore sessions.Store) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, context3.SetApp(sessionStore, r, w))
			})
	}
}

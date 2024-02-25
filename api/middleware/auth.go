package middleware

import (
	"github.com/gorilla/sessions"
	"ivixlabs.com/proj5/internal/http/context"
	"net/http"
)

func GretAuthMiddleware(next http.HandlerFunc, cookieStore sessions.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		appContext := context.GetApp(r.Context())

		if !appContext.IsAuth() {
			appContext.SaveRedirectUrl(r.URL.String())
			http.Redirect(w, r, "/auth", http.StatusTemporaryRedirect)
			return
		}

		next(w, r)
	}
}

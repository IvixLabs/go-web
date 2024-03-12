package middleware

import (
	"net/http"

	"ivixlabs.com/goweb/internal/http/context"
)

func GretAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		appContext := context.GetApp(r.Context())

		if !appContext.IsAuth() {
			appContext.SaveRedirectUrl(r.URL.String())
			http.Redirect(w, r, "/auth", http.StatusTemporaryRedirect)
			return
		}

		next.ServeHTTP(w, r)
	})
}

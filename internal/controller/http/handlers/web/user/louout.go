package user

import (
	"net/http"

	"ivixlabs.com/goweb/internal/http/context"
)

func GetLogoutHandler() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		appContext := context.GetApp(r.Context())
		appContext.Logout()

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})
}

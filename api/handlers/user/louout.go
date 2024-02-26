package user

import (
	"net/http"

	"ivixlabs.com/goweb/internal/http/context"
)

func GetLogoutHandlerFunc() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		appContext := context.GetApp(r.Context())
		appContext.Logout()

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

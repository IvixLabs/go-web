package user

import (
	"ivixlabs.com/proj5/internal/http/context"
	"net/http"
)

func GetLogoutHandlerFunc() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		appContext := context.GetApp(r.Context())
		appContext.Logout()

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

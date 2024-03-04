package user

import (
	"net/http"

	"github.com/gorilla/sessions"
	userComponents "ivixlabs.com/goweb/api/web/components/user"
	"ivixlabs.com/goweb/internal/http/context"
	"ivixlabs.com/goweb/internal/user"
	"ivixlabs.com/goweb/internal/user/auth"
	"ivixlabs.com/goweb/internal/validation/form"
)

func GetAuthHandlerFunc(userService user.Service, formValidator *form.Validator, cookieStore sessions.Store) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		authForm := auth.Form{}

		formErrors := &form.Errors{}

		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				panic(err)
			}

			authForm = auth.Form{
				Email:    r.PostFormValue("email"),
				Password: r.PostFormValue("password"),
			}

			ok := false
			if formErrors, ok = formValidator.ValidateForm(&authForm); ok {

				userObj := userService.FindByEmail(authForm.Email)

				appContext := context.GetApp(r.Context())
				appContext.Login(userObj.Id())

				w.Header().Set("HX-Redirect", appContext.GetRedirectUrl())
			}
		}

		err := userComponents.Auth(&authForm, formErrors).Render(r.Context(), w)
		if err != nil {
			panic(err)
		}
	}
}

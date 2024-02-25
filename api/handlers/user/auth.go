package user

import (
	"github.com/gorilla/sessions"
	userComponents "ivixlabs.com/proj5/api/web/components/user"
	"ivixlabs.com/proj5/internal/http/context"
	"ivixlabs.com/proj5/internal/user"
	"ivixlabs.com/proj5/internal/user/auth"
	"ivixlabs.com/proj5/internal/validation/form"
	"net/http"
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
				appContext.Login(userObj.GetId())

				w.Header().Set("HX-Redirect", appContext.GetRedirectUrl())
			}
		}

		err := userComponents.Auth(&authForm, formErrors).Render(r.Context(), w)
		if err != nil {
			panic(err)
		}
	}
}

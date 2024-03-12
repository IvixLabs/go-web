package user

import (
	"net/http"

	user2 "ivixlabs.com/goweb/internal/controller/http/template/user"
	"ivixlabs.com/goweb/internal/user"
	"ivixlabs.com/goweb/internal/user/registration"
	"ivixlabs.com/goweb/internal/validation/form"
)

func GetSignupHandler(userService user.Service, formValidator *form.Validator) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		creationForm := registration.Form{}

		formErrors := &form.Errors{}

		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				panic(err)
			}

			creationForm = registration.Form{
				Email:    r.PostFormValue("email"),
				Password: r.PostFormValue("password"),
				Address:  r.PostFormValue("address"),
			}

			ok := false
			if formErrors, ok = formValidator.ValidateForm(&creationForm); ok {
				_, err = userService.CreateNewUser(&creationForm)
				if err != nil {
					panic(err)
				}
				w.Header().Set("HX-Redirect", "/products")
			}
		}

		err := user2.Signup(&creationForm, formErrors).Render(r.Context(), w)
		if err != nil {
			panic(err)
		}
	})
}

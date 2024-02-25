package user

import (
	user2 "ivixlabs.com/proj5/api/web/components/user"
	"ivixlabs.com/proj5/internal/user"
	"ivixlabs.com/proj5/internal/user/registration"
	"ivixlabs.com/proj5/internal/validation/form"
	"net/http"
)

func GetSignupHandlerFunc(userService user.Service, formValidator *form.Validator) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}

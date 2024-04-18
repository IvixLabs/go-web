package user

import (
	"encoding/json"
	"ivixlabs.com/goweb/internal/model/user"
	"ivixlabs.com/goweb/internal/model/user/dto"
	"ivixlabs.com/goweb/internal/validation/form"
	"net/http"
)

func GetCreateHandler(userRepository user.Repository, formValidator *form.Validator) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var createForm dto.CreateForm

		err := json.NewDecoder(r.Body).Decode(&createForm)
		if err != nil {
			panic(err)
		}

		if formErrors, ok := formValidator.ValidateForm(&createForm); !ok {
			w.WriteHeader(http.StatusBadRequest)

			err = json.NewEncoder(w).Encode(formErrors.GetErrorsDto())
			if err != nil {
				panic(err)
			}

			return
		}

		userEntity := user.New(createForm.Email, createForm.Password, createForm.Address)
		userRepository.SaveUser(userEntity)

		err = json.NewEncoder(w).Encode(GetUserDto(userEntity))

		if err != nil {
			panic(err)
		}
	})
}

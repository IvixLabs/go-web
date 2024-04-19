package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"ivixlabs.com/goweb/internal/model/user"
	"ivixlabs.com/goweb/internal/model/user/dto"
	"ivixlabs.com/goweb/internal/validation/form"
	"net/http"
)

func GetUpdateHandler(userRepository user.Repository, formValidator *form.Validator) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var updateForm dto.UpdateForm

		err := json.NewDecoder(r.Body).Decode(&updateForm)
		if err != nil {
			panic(err)
		}

		if formErrors, ok := formValidator.ValidateForm(&updateForm); !ok {
			w.WriteHeader(http.StatusBadRequest)

			err = json.NewEncoder(w).Encode(formErrors.GetErrorsDto())
			if err != nil {
				panic(err)
			}

			return
		}

		params := mux.Vars(r)
		userId := params["id"]

		userEntity, err := userRepository.GetUserById(userId)
		if err != nil {
			panic(err)
		}

		updateDto, _ := updateForm.GetUpdateDto()
		userEntity.Update(updateDto)

		userRepository.SaveUser(userEntity)

		err = json.NewEncoder(w).Encode(GetUserDto(userEntity))

		if err != nil {
			panic(err)
		}
	})
}

package product

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"ivixlabs.com/goweb/internal/model/user"
	"ivixlabs.com/goweb/internal/validation/form"
)

func InitUserValidation(formValidator *form.Validator, userRepository user.Repository) {

	err := formValidator.Validate.RegisterValidation("user_not_exists", isUserNotExists(userRepository))
	if err != nil {
		panic(err)
	}

	err = formValidator.Validate.RegisterTranslation("user_not_exists", formValidator.Translator, func(ut ut.Translator) error {
		return ut.Add("user_not_exists", "{0} does not exist", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("user_not_exists", (fe.Value()).(string))

		return t
	})
	if err != nil {
		panic(err)
	}

}

func isUserNotExists(userRepository user.Repository) func(fl validator.FieldLevel) bool {

	return func(fl validator.FieldLevel) bool {
		userObj, _ := userRepository.GetUserById(fl.Field().String())

		return userObj != nil
	}
}

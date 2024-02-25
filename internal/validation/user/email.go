package user

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	user2 "ivixlabs.com/proj5/internal/user"
	"ivixlabs.com/proj5/internal/user/auth"
	"ivixlabs.com/proj5/internal/validation/form"
)

func InitEmailValidation(formValidator *form.Validator, userService user2.Service) {

	err := formValidator.Validate.RegisterValidation("user_email_exists", isUserEmailExists(userService))
	if err != nil {
		panic(err)
	}

	err = formValidator.Validate.RegisterTranslation("user_email_exists", formValidator.Translator, func(ut ut.Translator) error {
		return ut.Add("user_email_exists", "{0} is already registered", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("user_email_exists", (fe.Value()).(string))

		return t
	})
	if err != nil {
		panic(err)
	}

	err = formValidator.Validate.RegisterValidation("user_email_not_exists", isUserEmailNotExists(userService))
	if err != nil {
		panic(err)
	}

	err = formValidator.Validate.RegisterTranslation("user_email_not_exists", formValidator.Translator, func(ut ut.Translator) error {
		return ut.Add("user_email_not_exists", "{0} is not found", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("user_email_not_exists", (fe.Value()).(string))

		return t
	})
	if err != nil {
		panic(err)
	}

	formValidator.Validate.RegisterStructValidation(GetAuthLevelValidator(userService), auth.Form{})

	err = formValidator.Validate.RegisterTranslation("user_wrong_password", formValidator.Translator, func(ut ut.Translator) error {
		return ut.Add("user_wrong_password", "Wrong credentials", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("user_wrong_password", (fe.Value()).(string))

		return t
	})
	if err != nil {
		panic(err)
	}
}

func isUserEmailExists(userService user2.Service) func(fl validator.FieldLevel) bool {

	return func(fl validator.FieldLevel) bool {
		userObj := userService.FindByEmail(fl.Field().String())

		return userObj == nil
	}
}

func isUserEmailNotExists(userService user2.Service) func(fl validator.FieldLevel) bool {

	return func(fl validator.FieldLevel) bool {
		userObj := userService.FindByEmail(fl.Field().String())

		return userObj != nil
	}
}

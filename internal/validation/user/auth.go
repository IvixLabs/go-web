package user

import (
	"github.com/go-playground/validator/v10"
	"ivixlabs.com/goweb/internal/user"
	"ivixlabs.com/goweb/internal/user/auth"
)

func GetAuthLevelValidator(userService user.Service) validator.StructLevelFunc {

	return func(sl validator.StructLevel) {
		authForm := sl.Current().Interface().(auth.Form)

		userObj := userService.FindByEmail(authForm.Email)

		if userObj != nil {
			if userObj.GetPassword() != authForm.Password {
				sl.ReportError(authForm.Email, "Email", "Email", "user_wrong_password", "")
			}
		} else {
			sl.ReportError(authForm.Email, "Email", "Email", "user_wrong_password", "")
		}
	}
}

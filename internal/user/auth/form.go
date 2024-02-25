package auth

type Form struct {
	Email    string `validate:"required,user_email_not_exists"`
	Password string `validate:"required"`
}

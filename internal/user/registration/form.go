package registration

type Form struct {
	Email    string `validate:"required,user_email_exists"`
	Password string `validate:"required"`
	Address  string `validate:"required"`
}

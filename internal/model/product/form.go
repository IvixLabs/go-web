package product

type Form struct {
	Title string `validate:"required"`
	Brand string `validate:"required"`
	Price string `validate:"required,number"`
}

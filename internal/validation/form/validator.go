package form

import (
	"errors"
	enlocale "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	Validate   *validator.Validate
	Translator ut.Translator
}

func NewValidator() *Validator {
	validate := validator.New(validator.WithRequiredStructEnabled())

	enLocale := enlocale.New()
	uni := ut.New(enLocale, enLocale)
	translator, _ := uni.FindTranslator("en")

	err := en.RegisterDefaultTranslations(validate, translator)
	if err != nil {
		panic(err)
	}

	return &Validator{Validate: validate, Translator: translator}
}

func (formValidator *Validator) ValidateForm(form interface{}) (*Errors, bool) {

	validate := formValidator.Validate

	err := validate.Struct(form)

	if err != nil {
		var invalidValidationError validator.InvalidValidationError
		if errors.Is(err, &invalidValidationError) {
			panic(err)

		}

		var validationErrors validator.ValidationErrors
		if !errors.As(err, &validationErrors) {
			panic(err)
		}

		return NewErrors(&validationErrors, formValidator.Translator), false
	}

	return &Errors{}, true
}

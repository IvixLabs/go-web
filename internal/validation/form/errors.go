package form

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Errors struct {
	errors     *validator.ValidationErrors
	translator ut.Translator
}

func NewErrors(errors *validator.ValidationErrors, translator ut.Translator) *Errors {

	return &Errors{errors: errors, translator: translator}
}

func (form *Errors) GetStringError(name string) (string, bool) {

	if form.errors == nil {
		return "", false
	}

	for _, err := range *form.errors {
		if err.Field() == name {
			return err.Translate(form.translator), true
		}
	}

	return "", false
}

func (form *Errors) IsError(name string) bool {

	if form.errors == nil {
		return false
	}

	for _, err := range *form.errors {
		if err.Field() == name {
			return true
		}
	}

	return false
}

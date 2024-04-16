package form

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
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

func (form *Errors) GetErrorsDto() map[string][]string {

	finalErrors := make(map[string][]string)

	if form.errors == nil {
		return finalErrors
	}

	for _, err := range *form.errors {
		field := err.Field()

		finalErrors[strings.ToLower(string(field[0]))+field[1:]] = append(finalErrors[field], err.Translate(form.translator))
	}

	return finalErrors
}

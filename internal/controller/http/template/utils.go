package template

import (
	"context"
	"strings"

	appContext "ivixlabs.com/goweb/internal/http/context"
	"ivixlabs.com/goweb/internal/validation/form"
)

func App(ctx context.Context) *appContext.App {
	return appContext.GetApp(ctx)
}

func GetErrorMessage(formErrors *form.Errors, fieldName string) string {
	if errorString, isError := formErrors.GetStringError(fieldName); isError {
		return "<div class='invalid-feedback'>" + errorString + "</div>"
	}
	return ""
}

func GetClass(isError bool) string {

	parts := []string{"form-control"}

	if isError {
		parts = append(parts, "is-invalid")
	}

	return strings.Join(parts, " ")
}

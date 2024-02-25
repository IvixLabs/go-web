package components

import (
	"context"
	appContext "ivixlabs.com/proj5/internal/http/context"
	"ivixlabs.com/proj5/internal/validation/form"
	"strings"
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

package property

import (
	"encoding/json"
	"ivixlabs.com/goweb/internal/model/property"
	property2 "ivixlabs.com/goweb/internal/property"
	"ivixlabs.com/goweb/internal/validation/form"
	"net/http"
)

type createProperty struct {
	Name string
}

func GetCreateHandler(propertyRepository property.Repository, formValidator *form.Validator) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var createForm property2.CreateForm

		err := json.NewDecoder(r.Body).Decode(&createForm)
		if err != nil {
			panic(err)
		}

		if formErrors, ok := formValidator.ValidateForm(&createForm); !ok {
			w.WriteHeader(http.StatusBadRequest)

			err = json.NewEncoder(w).Encode(formErrors.GetErrorsDto())
			if err != nil {
				panic(err)
			}

			return
		}

		propertyEntity := property.New(
			"bb920be7-166e-4dc7-8a6f-696f0ec7a963",
			createForm.Name,
		)

		propertyRepository.Create(propertyEntity)

		err = json.NewEncoder(w).Encode(property.GetPropertyDto(propertyEntity))
		if err != nil {
			panic(err)
		}
	})
}

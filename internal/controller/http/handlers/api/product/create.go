package product

import (
	"encoding/json"
	"ivixlabs.com/goweb/internal/model/product"
	"ivixlabs.com/goweb/internal/model/product/dto"
	"ivixlabs.com/goweb/internal/validation/form"
	"net/http"
)

func GetCreateHandler(productRepository product.Repository, formValidator *form.Validator) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var createForm dto.CreateForm

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

		prodObj := product.New(createForm.UserId, createForm.Title, createForm.Price, createForm.Brand)
		productRepository.SaveProduct(prodObj)

		err = json.NewEncoder(w).Encode(GetDto(prodObj))

		if err != nil {
			panic(err)
		}
	})
}

package product

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"ivixlabs.com/goweb/internal/model/product"
	"ivixlabs.com/goweb/internal/model/product/dto"
	"ivixlabs.com/goweb/internal/validation/form"
	"net/http"
)

func GetUpdateHandler(productRepository product.Repository, formValidator *form.Validator) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var updateForm dto.UpdateForm

		err := json.NewDecoder(r.Body).Decode(&updateForm)
		if err != nil {
			panic(err)
		}

		if formErrors, ok := formValidator.ValidateForm(&updateForm); !ok {
			w.WriteHeader(http.StatusBadRequest)

			err = json.NewEncoder(w).Encode(formErrors.GetErrorsDto())
			if err != nil {
				panic(err)
			}

			return
		}

		params := mux.Vars(r)
		id := params["id"]

		entity, err := productRepository.GetProductById(id)
		if err != nil {
			panic(err)
		}

		updateDto, _ := updateForm.GetUpdateDto()
		entity.Update(updateDto)

		productRepository.SaveProduct(entity)

		err = json.NewEncoder(w).Encode(GetDto(entity))

		if err != nil {
			panic(err)
		}
	})
}

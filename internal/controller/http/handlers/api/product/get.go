package product

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"ivixlabs.com/goweb/internal/model/product"
	"net/http"
)

func GetGetHandler(productRepository product.Repository) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		userId := params["id"]

		userObj, err := productRepository.GetProductById(userId)

		if err != nil {
			panic(err)
		}

		err = json.NewEncoder(w).Encode(GetDto(userObj))

		if err != nil {
			panic(err)
		}
	})
}

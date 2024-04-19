package product

import (
	"fmt"
	"github.com/gorilla/mux"
	"ivixlabs.com/goweb/internal/model/product"
	"net/http"
)

func GetDeleteHandler(productRepository product.Repository) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		userId := params["id"]

		productRepository.DeleteProductById(userId)

		_, err := fmt.Fprint(w, "ok")

		if err != nil {
			panic(err)
		}
	})
}

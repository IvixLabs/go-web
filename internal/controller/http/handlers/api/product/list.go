package product

import (
	"encoding/json"
	"ivixlabs.com/goweb/internal/model/product"
	"net/http"
)

func GetListHandler(productService product.Service) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prods := productService.FindAll()

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		ret := make([]map[string]string, len(prods))

		for userIndex, userItem := range prods {
			ret[userIndex] = map[string]string{
				"id":    userItem.Id(),
				"title": userItem.Title(),
				"brand": userItem.Brand(),
			}
		}

		err := json.NewEncoder(w).Encode(ret)

		if err != nil {
			panic(err)
		}
	})
}

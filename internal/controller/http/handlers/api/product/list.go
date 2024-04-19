package product

import (
	"encoding/json"
	"ivixlabs.com/goweb/internal/model/product"
	"net/http"
)

func GetListHandler(productService product.Service) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prods := productService.FindAll()

		ret := make([]*Dto, len(prods))

		for prodIndex, prodItem := range prods {
			ret[prodIndex] = GetDto(prodItem)
		}

		err := json.NewEncoder(w).Encode(ret)

		if err != nil {
			panic(err)
		}
	})
}

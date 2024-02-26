package product

import (
	"net/http"

	"ivixlabs.com/goweb/api/web/components/product"
	"ivixlabs.com/goweb/internal/http/context"
	product3 "ivixlabs.com/goweb/internal/product"
)

func GetListHandler(productService product3.Service) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		appContext := context.GetApp(r.Context())
		userId := appContext.GetUserId()

		products := productService.FindByUserId(userId)

		if r.URL.Query().Has("tableOnly") {
			err := product.ProductTableView(products).Render(r.Context(), w)
			if err != nil {
				panic(err)
			}
		} else {
			err := product.ListView(products).Render(r.Context(), w)
			if err != nil {
				panic(err)
			}
		}
	}
}

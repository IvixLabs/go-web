package product

import (
	"net/http"

	"ivixlabs.com/goweb/internal/controller/http/template/product"
	"ivixlabs.com/goweb/internal/http/context"
	product2 "ivixlabs.com/goweb/internal/model/product"
)

func GetDeleteHandler(productRepository product2.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		productId := r.URL.Query().Get("productId")

		if productId == "" {
			return
		}

		if r.Method == http.MethodDelete {

			appContext := context.GetApp(r.Context())
			userId := appContext.GetUserId()

			num := productRepository.DeleteProduct(productId, userId)
			if num > 0 {
				w.Header().Set("HX-Trigger", "UpdateProductsTable,CloseProductDeleteModal")
			}

			err := product.DisabledButton().Render(r.Context(), w)
			if err != nil {
				panic(err)
			}
		}

		if r.Method == http.MethodGet {
			err := product.ModalDeleteView(productId).Render(r.Context(), w)
			if err != nil {
				panic(err)
			}
		}
	}
}

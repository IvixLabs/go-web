package product

import (
	"net/http"
	"strconv"

	"ivixlabs.com/proj5/api/web/components/product"
	"ivixlabs.com/proj5/internal/http/context"
	"ivixlabs.com/proj5/internal/model"
	product2 "ivixlabs.com/proj5/internal/product"
	"ivixlabs.com/proj5/internal/validation/form"
)

func GetSaveHandler(formValidator *form.Validator, productService product2.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		productForm := product2.Form{}
		formErrors := &form.Errors{}
		productId := r.URL.Query().Get("productId")

		var productObj model.Product

		if productId != "" {
			productObj = productService.FindById(productId)
			appContext := context.GetApp(r.Context())
			userId := appContext.GetUserId()

			if productObj == nil || productObj.GetUserId() != userId {
				panic("Product is not found")
			}

			productForm.Title = productObj.GetTitle()
			productForm.Brand = productObj.GetBrand()
			productForm.Price = strconv.Itoa(productObj.GetPrice())
		}

		var err error
		if r.Method == http.MethodPost {
			err = r.ParseForm()
			if err != nil {
				panic(err)
			}

			productForm = product2.Form{
				Title: r.PostFormValue("title"),
				Brand: r.PostFormValue("brand"),
				Price: r.PostFormValue("price"),
			}

			formDisabled := false
			ok := false
			if formErrors, ok = formValidator.ValidateForm(&productForm); ok {

				if productObj != nil {
					productService.UpdateProduct(&productForm, productObj)
				} else {
					app := context.GetApp(r.Context())

					productObj = productService.CreateNewProduct(&productForm, app.GetUserId())
					productId = productObj.GetId()
				}

				w.Header().Set("HX-Trigger", "UpdateProductsTable,CloseProductFormModal")
				formDisabled = true
			}

			err = product.ProductFormView(&productForm, formErrors, productId, formDisabled).Render(r.Context(), w)
		} else {
			err = product.ModalFormView(&productForm, formErrors, productId).Render(r.Context(), w)
		}

		if err != nil {
			panic(err)
		}
	}
}

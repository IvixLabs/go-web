package product

import (
	"ivixlabs.com/goweb/internal/controller/http/template/product"
	productModel "ivixlabs.com/goweb/internal/model/product"
	"ivixlabs.com/goweb/internal/model/product/usecase"
	"net/http"
	"strconv"

	"ivixlabs.com/goweb/internal/http/context"
	product2 "ivixlabs.com/goweb/internal/model/product"
	"ivixlabs.com/goweb/internal/validation/form"
)

func GetSaveHandler(formValidator *form.Validator,
	productRepository product2.Repository,
	productUpdating usecase.ProductUpdating,
	productCreation usecase.ProductCreation) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		productForm := product2.Form{}
		formErrors := &form.Errors{}
		productId := r.URL.Query().Get("productId")

		var productObj productModel.Product

		if productId != "" {
			productObj = productRepository.FindProductById(productId)
			appContext := context.GetApp(r.Context())
			userId := appContext.GetUserId()

			if productObj == nil || productObj.UserId() != userId {
				panic("Product is not found")
			}

			productForm.Title = productObj.Title()
			productForm.Brand = productObj.Brand()
			productForm.Price = strconv.Itoa(productObj.Price())
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
					productUpdating.UpdateProduct(&productForm, productObj)
				} else {
					app := context.GetApp(r.Context())

					productObj = productCreation.CreateNewProduct(&productForm, app.GetUserId())
					productId = productObj.Id()
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
	})
}

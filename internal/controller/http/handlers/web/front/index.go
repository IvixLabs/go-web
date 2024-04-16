package front

import (
	"ivixlabs.com/goweb/internal/controller/http/template"
	"net/http"
)

func GetFrontHandler() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := template.Front().Render(r.Context(), w)
		if err != nil {
			panic(err)
		}
	})
}

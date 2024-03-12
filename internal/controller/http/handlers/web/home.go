package web

import (
	"ivixlabs.com/goweb/internal/controller/http/template"
	"net/http"
)

func GetHomeHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		err := template.Home().Render(r.Context(), w)
		if err != nil {
			panic(err)
		}
	}

}

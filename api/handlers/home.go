package handlers

import (
	"net/http"

	"ivixlabs.com/goweb/api/web/components"
)

func GetHomeHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		err := components.Home().Render(r.Context(), w)
		if err != nil {
			panic(err)
		}
	}

}

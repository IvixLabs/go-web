package handlers

import (
	"ivixlabs.com/proj5/api/web/components"
	"net/http"
)

func GetHomeHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		err := components.Home().Render(r.Context(), w)
		if err != nil {
			panic(err)
		}
	}

}

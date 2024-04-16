package dashboard

import (
	"ivixlabs.com/goweb/internal/controller/http/template"
	"net/http"
)

func GetDashboardHandler() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := template.Dashboard().Render(r.Context(), w)
		if err != nil {
			panic(err)
		}
	})
}

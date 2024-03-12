package video

import (
	"ivixlabs.com/goweb/internal/controller/http/template/video"
	"net/http"
)

func GetIndexHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		err := video.IndexView().Render(r.Context(), w)
		if err != nil {
			panic(err)
		}
	}
}

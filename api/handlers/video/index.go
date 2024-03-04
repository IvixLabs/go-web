package video

import (
	"net/http"

	"ivixlabs.com/goweb/api/web/components/video"
)

func GetIndexHandlerFunc() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		err := video.IndexView().Render(r.Context(), w)
		if err != nil {
			panic(err)
		}
	}
}

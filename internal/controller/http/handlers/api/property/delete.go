package property

import (
	"encoding/json"
	"ivixlabs.com/goweb/internal/model/property"
	"log"
	"net/http"
)

func GetDeleteHandler(propertyRepository property.Repository) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println(r.URL.Query())

		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if !r.URL.Query().Has("id") {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		id := r.URL.Query().Get("id")

		propertyRepository.Delete(id, "bb920be7-166e-4dc7-8a6f-696f0ec7a963")

		err := json.NewEncoder(w).Encode(struct{ status string }{status: "ok"})
		if err != nil {
			panic(err)
		}
	})
}

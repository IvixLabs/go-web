package property

import (
	"encoding/json"
	"ivixlabs.com/goweb/internal/model/property"
	"net/http"
)

func GetListHandler(propertyRepository property.Repository) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		properties := propertyRepository.FindAll()

		ret := make([]property.PropertyDto, len(properties))

		for index, item := range properties {
			ret[index] = property.GetPropertyDto(item)
		}

		err := json.NewEncoder(w).Encode(ret)

		if err != nil {
			panic(err)
		}
	})
}

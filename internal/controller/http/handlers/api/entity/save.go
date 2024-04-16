package entity

import (
	"encoding/json"
	"ivixlabs.com/goweb/internal/model"
	"log"
	"net/http"
)

func GetSaveHandler(entityRepo model.EntityRepository) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var entityDto model.EntityDto

		err := json.NewDecoder(r.Body).Decode(&entityDto)
		if err != nil {
			panic(err)
		}

		log.Println("===========")
		log.Println(entityDto)

		var en model.Entity
		if entityDto.Id != "" {
			ens, _ := entityRepo.Find([]string{"EntityId='" + entityDto.Id + "'"}, 0, 1, "", "", false)
			if len(ens) > 0 {
				en = ens[0]
			}
		} else {
			en = model.NewEntity()
		}

		for _, propDto := range entityDto.Properties {
			if propDto.Disabled {
				en.SetProperty(propDto.Name, propDto.Type, nil)
			} else {
				en.SetProperty(propDto.Name, propDto.Type, propDto.Value)
			}
		}

		log.Println(en)

		entityRepo.Save(en)

		err = json.NewEncoder(w).Encode("OK")

		if err != nil {
			panic(err)
		}
	})
}

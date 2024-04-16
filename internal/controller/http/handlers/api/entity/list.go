package entity

import (
	"cmp"
	"encoding/json"
	"ivixlabs.com/goweb/internal/model"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

type property struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func GetListHandler(entityRepo model.EntityRepository, entityPropertiesRepo model.EntityPropertyRepository) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sortField := r.URL.Query().Get("sortField")
		if sortField == "" {
			sortField = "EntityId"
		}

		sortFieldType := r.URL.Query().Get("sortFieldType")
		if sortFieldType == "" {
			sortFieldType = "string"
		}

		sortOrder := r.URL.Query().Get("sortOrder") == "1"

		id := r.URL.Query().Get("id")

		var firstRows int64 = 0
		var totalRows int64 = 10

		var err error

		strFirstRows := r.URL.Query().Get("first")
		if strFirstRows != "" {
			firstRows, err = strconv.ParseInt(strFirstRows, 10, 64)
			if err != nil {
				panic(err)
			}

		}

		strTotalRows := r.URL.Query().Get("totalRows")
		if strTotalRows != "" {
			totalRows, err = strconv.ParseInt(strTotalRows, 10, 64)
			if err != nil {
				panic(err)
			}
		}

		var whereAndPart []string
		if id != "" {
			whereAndPart = []string{"EntityId='" + id + "'"}
		}

		log.Println(firstRows)

		entities, total := entityRepo.Find(whereAndPart, firstRows, totalRows, sortField, sortFieldType, sortOrder)

		ret := make(map[string]any)

		dtos := make([]model.EntityDto, len(entities))

		for index, item := range entities {
			dtos[index] = model.GetEntityDto(item)
		}

		ret["total"] = total
		ret["entities"] = dtos

		mapFilteredPropNames := make(map[string]bool)
		slicesFilteredPropNames := make([]*property, 0, 10)

		for _, dto := range dtos {
			for _, prop := range dto.Properties {

				_, ok := mapFilteredPropNames[prop.Name]
				if !ok {
					mapFilteredPropNames[prop.Name] = true
					slicesFilteredPropNames = append(
						slicesFilteredPropNames,
						&property{Name: prop.Name, Type: prop.Type},
					)
				}
			}
		}

		slices.SortFunc(slicesFilteredPropNames, func(a, b *property) int {
			return cmp.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
		})

		ret["properties"] = slicesFilteredPropNames

		err = json.NewEncoder(w).Encode(ret)

		if err != nil {
			panic(err)
		}
	})
}

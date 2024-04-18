package user

import (
	"encoding/json"
	"net/http"

	"ivixlabs.com/goweb/internal/user"
)

func GetListHandler(userService user.Service) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users := userService.FindAll()

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		ret := make([]*Dto, len(users))

		for userIndex, userItem := range users {
			ret[userIndex] = GetUserDto(userItem)
		}

		err := json.NewEncoder(w).Encode(ret)

		if err != nil {
			panic(err)
		}
	})
}

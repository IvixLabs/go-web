package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"ivixlabs.com/goweb/internal/model/user"
	"log"
	"net/http"
)

func GetGetHandler(userRepository user.Repository) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println(r.URL.Path)

		params := mux.Vars(r)
		userId := params["id"]

		userObj, err := userRepository.GetUserById(userId)

		if err != nil {
			panic(err)
		}

		err = json.NewEncoder(w).Encode(GetUserDto(userObj))

		if err != nil {
			panic(err)
		}
	})
}

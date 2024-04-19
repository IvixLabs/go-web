package user

import (
	"fmt"
	"github.com/gorilla/mux"
	"ivixlabs.com/goweb/internal/model/user"
	"net/http"
)

func GetDeleteHandler(userRepository user.Repository) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		userId := params["id"]

		userRepository.DeleteUserById(userId)

		_, err := fmt.Fprint(w, "ok")

		if err != nil {
			panic(err)
		}
	})
}

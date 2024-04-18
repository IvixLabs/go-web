package user

import (
	"fmt"
	"ivixlabs.com/goweb/internal/model/user"
	"net/http"
)

func GetDeleteHandler(userRepository user.Repository) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userId := r.URL.Query().Get("userId")

		userRepository.DeleteUserById(userId)

		_, err := fmt.Fprint(w, "ok")

		if err != nil {
			panic(err)
		}
	})
}

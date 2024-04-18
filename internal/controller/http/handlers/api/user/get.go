package user

import (
	"encoding/json"
	"ivixlabs.com/goweb/internal/model/user"
	"net/http"
)

func GetGetHandler(userRepository user.Repository) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userId := r.URL.Query().Get("userId")

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

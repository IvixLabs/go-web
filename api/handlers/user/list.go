package user

import (
	"net/http"

	user2 "ivixlabs.com/goweb/api/web/components/user"
	"ivixlabs.com/goweb/internal/user"
)

func GetListHandlerFunc(userService user.Service) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		users := userService.FindAll()

		err := user2.List(users).Render(r.Context(), w)
		if err != nil {
			panic(err)
		}
	}
}

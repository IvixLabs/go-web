package user

import (
	"net/http"

	user2 "ivixlabs.com/goweb/internal/controller/http/template/user"
	"ivixlabs.com/goweb/internal/user"
)

func GetListHandler(userService user.Service) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users := userService.FindAll()

		err := user2.List(users).Render(r.Context(), w)
		if err != nil {
			panic(err)
		}
	})
}

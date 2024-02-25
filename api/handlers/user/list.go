package user

import (
	user2 "ivixlabs.com/proj5/api/web/components/user"
	"ivixlabs.com/proj5/internal/user"
	"net/http"
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

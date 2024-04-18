package user

import "ivixlabs.com/goweb/internal/model/user"

type Dto struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func GetUserDto(userObj user.User) *Dto {
	return &Dto{
		Id:      userObj.Id(),
		Email:   userObj.Email(),
		Address: userObj.Address(),
	}
}

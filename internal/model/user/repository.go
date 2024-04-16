package user

type Repository interface {
	SaveUser(u User)
	FindAllUsers() []User
	FindUserByEmail(email string) User
}

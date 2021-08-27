package user

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	Login(User)(string,error)
	Signup(User)(int,error)
}
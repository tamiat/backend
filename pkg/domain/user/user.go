package user

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	Otp string `json:"otp"`
	EmailVerified bool `json:"email_verified"`
}

type UserRepository interface {
	Login(User)(string,error)
	Signup(User)(int,error)
	InsertOTP(string)error
}

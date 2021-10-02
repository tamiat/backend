package user

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	Otp string `json:"otp"`
	EmailVerified bool `json:"email_verified"`
	Role     string `json:"role"`
}

type UserRepository interface {
	Login(User)(string,error)
	Signup(User)(int,error)
	InsertOTP(User)error
	VerifyEmail(User) error
	ReadOTP(User)(string,error)
}

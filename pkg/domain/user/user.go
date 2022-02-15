package user

type User struct {
	ID            int    `json:"id"`
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password" binding:"required"`
	Otp           string `json:"otp"`
	EmailVerified bool   `json:"email_verified" binding:"required"`
	Role          string `json:"role"`
}

type UserRepository interface {
	Login(User) (string, error)
	Signup(User) (int, error)
	InsertOTP(User) error
	VerifyEmail(User) error
	ReadOTP(User) (string, error)
}

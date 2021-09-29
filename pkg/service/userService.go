package service

import (
	"github.com/tamiat/backend/pkg/domain/user"
)

type UserService interface {
	Login(user.User)(string,error)
	Signup(user.User)(int,error)
	InsertOTP(string)error

}

type DefaultUserService struct {
	repo user.UserRepository
}

func (s DefaultUserService )Login(user user.User) (string,error) {
	return s.repo.Login(user)
}

func (s DefaultUserService )Signup(user user.User) (int,error) {
	return s.repo.Signup(user)
}
func (s DefaultUserService )InsertOTP(otp string) error {
	return s.repo.InsertOTP(otp)
}
func NewUserService(repository user.UserRepository) DefaultUserService {
	return DefaultUserService{repo: repository}
}
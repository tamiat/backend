package service

import (
	"github.com/tamiat/backend/pkg/domain/user"
)

type UserService interface {
	Login(user.User)(string,error)
	Signup(user.User)(int,error)
	InsertOTP(user.User)error
	VerifyEmail(user.User)error
	ReadOTP(user.User)(string,error)

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
func (s DefaultUserService )InsertOTP(user user.User) error {
	return s.repo.InsertOTP(user)
}
func (s DefaultUserService )VerifyEmail(user user.User) error {
	return s.repo.VerifyEmail(user)
}
func (s DefaultUserService )ReadOTP(user user.User) (string,error) {
	return s.repo.ReadOTP(user)
}

func NewUserService(repository user.UserRepository) DefaultUserService {
	return DefaultUserService{repo: repository}
}

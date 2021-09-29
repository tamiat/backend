package user

import (
	"gorm.io/gorm"

	"github.com/tamiat/backend/pkg/errs"
)

type UserRepositoryDb struct {
	db *gorm.DB
}

func (r UserRepositoryDb) Login(userObj User) (string,error){
	var retrievedUsr User
	if err := r.db.Where("email = ?", userObj.Email).First(&retrievedUsr).Error; err!=nil{
		return "",errs.ErrDb
	}
	return retrievedUsr.Password,nil
}

func (r UserRepositoryDb) Signup(user User) (int,error){
	if err:= r.db.Select("email","password").Create(&user).Error; err!=nil{
		return -1,errs.ErrDb
	}
	return user.ID,nil
}

func (r UserRepositoryDb) InsertOTP(user User) error {
	if err:= r.db.Model(&User{}).Where("id = ?", user.ID).Update("otp", user.Otp).Error; err != nil {
		return err
	}
	return nil
}
func (r UserRepositoryDb) ReadOTP(user User) (string,error) {
	if err:= r.db.First(&user, user.ID).Error; err != nil {
		return "",err
	}
	return user.Otp,nil
}
func (r UserRepositoryDb) VerifyEmail(user User) error {
	if err:= r.db.Model(&User{}).Where("id = ?", user.ID).Update("email_verified", true).Error; err != nil {
		return err
	}
	return nil
}

func NewUserRepositoryDb(db *gorm.DB) UserRepositoryDb {
	return UserRepositoryDb{db}
}
package user

import "database/sql"

type UserRepositoryDb struct {
	db *sql.DB
}

func (r UserRepositoryDb) Login(user User) error{
	return nil
}

func (r UserRepositoryDb) Signup(user User) (User,error){
	return user,nil
}

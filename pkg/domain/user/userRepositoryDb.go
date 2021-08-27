package user

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type UserRepositoryDb struct {
	db *sql.DB
}

func (r UserRepositoryDb) Login(userObj User) (string,error){
	row := r.db.QueryRow("select * from users where email=$1", userObj.Email)
	err := row.Scan(&userObj.Id, &userObj.Email, &userObj.Password)
	if err != nil {
		return "",err
	}
	return userObj.Password,nil
}

func (r UserRepositoryDb) Signup(user User) (User,error){
	return user,nil
}

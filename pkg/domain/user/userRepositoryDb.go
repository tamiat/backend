package user

import (
	"database/sql"
)

type UserRepositoryDb struct {
	db *sql.DB
}

func (r UserRepositoryDb) Login(userObj User) (string,error){
	row := r.db.QueryRow("select password from users where email=$1", userObj.Email)
	err := row.Scan(&userObj.Password)
	if err != nil {
		return "",err
	}
	return userObj.Password,nil
}

func (r UserRepositoryDb) Signup(user User) (int,error){
	query := "insert into users (email, password) values($1, $2) RETURNING id;"
	err := r.db.QueryRow(query, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		return 0,err
	}
	return user.Id,nil
}

func NewUserRepositoryDb(db *sql.DB) UserRepositoryDb {
	return UserRepositoryDb{db}
}
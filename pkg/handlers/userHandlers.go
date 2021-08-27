package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/tamiat/backend/pkg/domain"
	"github.com/tamiat/backend/pkg/domain/user"
	"github.com/tamiat/backend/pkg/middleware"
	"github.com/tamiat/backend/pkg/service"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/mail"
)

type UserHandlers struct {
	service service.UserService
}

type JWT struct {
	Token string `json:"token"`
}

func (receiver UserHandlers) Signup(w http.ResponseWriter, r *http.Request){
	//extracting usr obj
	var userObj user.User
	json.NewDecoder(r.Body).Decode(&userObj)
	//validating email and password
	status,msg:=validateEmailAndPassword(userObj)
	if status==http.StatusBadRequest{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.Response400(msg))
		return
	}
	//encrypting password
	hash, err := bcrypt.GenerateFromPassword([]byte(userObj.Password), 10)
	if err != nil {
		//TODO check if this is right
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.Response500(err.Error()))
		return
	}
	userObj.Password = string(hash)
	//database connection
	userObj.Id,err = receiver.service.Signup(userObj)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.Response500("Internal server error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	userObj.Password = ""
	json.NewEncoder(w).Encode(userObj)
}
func (receiver UserHandlers) Login(w http.ResponseWriter, r *http.Request)  {
	var userObj user.User
	json.NewDecoder(r.Body).Decode(&userObj)
	//validating email and password
	status,msg:=validateEmailAndPassword(userObj)
	if status==http.StatusBadRequest{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.Response400(msg))
		return
	}
	hashedPassword,err:=receiver.service.Login(userObj)
	if err!=nil{
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(domain.Response500("This user does not exist"))
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(domain.Response500("Server error"))
			return
		}
	}
	//usr password before hashing
	password := userObj.Password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	token, err := middleware.GenerateToken(userObj)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(domain.Response401("can't generate token"))
		return
	}
	w.WriteHeader(http.StatusOK)
	jwtObj:=JWT{Token: token}
	json.NewEncoder(w).Encode(jwtObj)
}
func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func validateEmailAndPassword(userObj user.User)(int,string){
	//err error()
	if userObj.Email == "" {
		/*w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response400("Email is missing"))
		return*/
		return http.StatusBadRequest,"Email is missing"
	}
	if !valid(userObj.Email){
		/*w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response400("Invalid Email"))*/
		return http.StatusBadRequest,"Invalid Email"
	}
	if userObj.Password == "" {
		/*w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response400("Password is missing"))*/
		return http.StatusBadRequest,"Password is missing"
	}
	return -1,""
}
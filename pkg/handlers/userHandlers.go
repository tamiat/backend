package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tamiat/backend/pkg/emailVerification"
	"net/http"
	"net/mail"
	"regexp"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/tamiat/backend/pkg/domain/user"
	"github.com/tamiat/backend/pkg/errs"
	"github.com/tamiat/backend/pkg/middleware"
	"github.com/tamiat/backend/pkg/service"
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
	err:=validateEmailAndPassword(userObj)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.NewResponse(err.Error(),http.StatusBadRequest))
		return
	}
	//encrypting password
	hash, err := bcrypt.GenerateFromPassword([]byte(userObj.Password), 10)
	if err != nil {
		//TODO check if this is right
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrServerErr.Error(),http.StatusInternalServerError))
		return
	}
	userObj.Password = string(hash)
	//database connection
	userObj.ID,err = receiver.service.Signup(userObj)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrServerErr.Error(),http.StatusInternalServerError))
		return
	}
	code,err:=emailVerification.SendEmail(userObj.Email)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(err.Error(),http.StatusInternalServerError))
		return
	}
	//fmt.Println(code)

	hashOTP, err := bcrypt.GenerateFromPassword([]byte(code), 10)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrServerErr.Error(),http.StatusInternalServerError))
		return
	}
	userObj.Otp = string(hashOTP)
	err=receiver.service.InsertOTP(userObj)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(err.Error(),http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	userObj.Password = ""
	userObj.Otp=""
	json.NewEncoder(w).Encode(userObj)

}
func (receiver UserHandlers) Login(w http.ResponseWriter, r *http.Request)  {
	var userObj user.User
	json.NewDecoder(r.Body).Decode(&userObj)
	//validating email and password
	err:=validateEmailAndPassword(userObj)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.NewResponse(err.Error(),http.StatusBadRequest))
		return
	}
	hashedPassword,err:=receiver.service.Login(userObj)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDb.Error(),http.StatusInternalServerError))
		return
	}
	//usr password before hashing
	password := userObj.Password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrInvalidPassword.Error(),http.StatusUnauthorized))
		return
	}
	token, err := middleware.GenerateToken(userObj)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrTokenErr.Error(),http.StatusUnauthorized))
		return
	}
	w.WriteHeader(http.StatusOK)
	jwtObj:=JWT{Token: token}
	json.NewEncoder(w).Encode(jwtObj)
}
func (receiver UserHandlers) VerifyEmail(w http.ResponseWriter, r *http.Request){
	var opt string
	json.NewDecoder(r.Body).Decode(&opt)
	if len(opt)!=6{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrInvalidVerificationCode.Error(),http.StatusBadRequest))
		return
	}
	vars := mux.Vars(r)
	pattern1, _ := regexp.Match(`^[0-9]+$`, []byte(vars["id"]))
	if !pattern1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.ErrContentParams)
		return
	}
	id := vars["id"]
	var userObj user.User
	intId, err := strconv.Atoi(id)
	userObj.ID = intId
	hashedOPT,err:=receiver.service.ReadOTP(userObj)
	err = bcrypt.CompareHashAndPassword([]byte(hashedOPT), []byte(opt))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrInvalidVerificationCode.Error(),http.StatusUnauthorized))
		return
	}
	err=receiver.service.UpdateEmailStatus(userObj)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.ErrDb)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func validateEmailAndPassword(userObj user.User)error{
	//err error()
	if userObj.Email == "" {
		return errs.ErrEmailMissing
	}
	if !valid(userObj.Email){
		return errs.ErrInvalidEmail
	}
	if userObj.Password == "" {
		return errs.ErrInvalidPassword
	}
	return nil
}
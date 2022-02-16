package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/mail"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/tamiat/backend/pkg/domain/user"
	"github.com/tamiat/backend/pkg/emailVerification"
	"github.com/tamiat/backend/pkg/errs"
	"github.com/tamiat/backend/pkg/middleware"
	"github.com/tamiat/backend/pkg/response"
	"github.com/tamiat/backend/pkg/service"
)

type UserHandlers struct {
	Service service.UserService
}

type JWT struct {
	Token string `json:"token"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (receiver UserHandlers) Signup(ctx *gin.Context) (user.User, int, error) {
	//w.Header().Set("Content-Type", "application/json")
	//extracting usr obj
	var userObj user.User
	err := ctx.ShouldBindJSON(&userObj)
	//json.NewDecoder(r.Body).Decode(&userObj)
	//validating email and password
	if err != nil {
		return userObj, http.StatusBadRequest, err
	}
	/*err = validateEmailAndPassword(userObj)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.NewResponse(err.Error(), http.StatusBadRequest))
		return
	}*/
	//encrypting password
	hash, err := bcrypt.GenerateFromPassword([]byte(userObj.Password), 10)
	if err != nil {
		//TODO check if this is right
		//w.WriteHeader(http.StatusInternalServerError)
		//json.NewEncoder(w).Encode(response.NewResponse(errs.ErrServerErr.Error(), http.StatusInternalServerError))
		return userObj, http.StatusInternalServerError, err
	}
	userObj.Password = string(hash)
	//database connection
	userObj.ID, err = receiver.Service.Signup(userObj)
	if err != nil {
		fmt.Println("mn eel service ya rahmaaa")
		return userObj, http.StatusInternalServerError, err
	}
	code, err := emailVerification.SendEmail(userObj.Email)
	if err != nil {
		//w.WriteHeader(http.StatusInternalServerError)
		//json.NewEncoder(w).Encode(response.NewResponse(err.Error(), http.StatusInternalServerError))
		return userObj, http.StatusInternalServerError, err
	}

	hashOTP, err := bcrypt.GenerateFromPassword([]byte(code), 10)
	if err != nil {
		//w.WriteHeader(http.StatusInternalServerError)
		//json.NewEncoder(w).Encode(response.NewResponse(errs.ErrServerErr.Error(), http.StatusInternalServerError))
		return userObj, http.StatusInternalServerError, err

	}
	userObj.Otp = string(hashOTP)
	err = receiver.Service.InsertOTP(userObj)
	if err != nil {
		//w.WriteHeader(http.StatusInternalServerError)
		//json.NewEncoder(w).Encode(response.NewResponse(err.Error(), http.StatusInternalServerError))
		return userObj, http.StatusInternalServerError, err

	}
	//w.WriteHeader(http.StatusOK)
	userObj.Password = ""
	userObj.Otp = ""
	//json.NewEncoder(w).Encode(userObj)
	return userObj, http.StatusOK, nil
}

//
// @Summary Login endpoint
// @Description Provide email and password to login, response is JWT
// @Accept application/json
// @Produce application/json
// @Param login body handlers.Login true "Login"
// @Success 200 {object} handlers.JWT
// @Failure 400  {object}  errs.ErrResponse "Bad Request"
// @Failure 404  {object}  errs.ErrResponse "User not found"
// @Failure 401  {object}  errs.ErrResponse "Unauthorizes"
// @Failure 500  {object}  errs.ErrResponse "Internal server error"
// @Router /login [post]
func (receiver UserHandlers) Login(ctx *gin.Context) {
	var userObj user.User
	var loginRequestData Login
	//decoding request body
	if err := ctx.ShouldBindJSON(&loginRequestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userObj.Email = loginRequestData.Email
	userObj.Password = loginRequestData.Password

	// retrieving hashed password from database
	hashedPassword, err := receiver.Service.Login(userObj)
	if err != nil {
		if err == errs.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": errs.ErrRecordNotFound.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrDb.Error()})
		return
	}
	// authentication process
	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(loginRequestData.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errs.ErrInvalidPassword.Error()})
		return
	}
	//generating token
	token, err := middleware.GenerateToken(userObj)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrTokenErr.Error()})
		return
	}
	jwtObj := JWT{Token: token}
	ctx.JSON(http.StatusOK, jwtObj)
}
func (receiver UserHandlers) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	var userObj user.User
	json.NewDecoder(r.Body).Decode(&userObj)
	otp := userObj.Otp
	if len(otp) != 6 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.NewResponse(errs.ErrInvalidVerificationCode.Error(), http.StatusBadRequest))
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
	//var userObj user.User
	intId, err := strconv.Atoi(id)
	fmt.Println(intId)
	userObj.ID = intId
	hashedOTP, err := receiver.Service.ReadOTP(userObj)
	err = bcrypt.CompareHashAndPassword([]byte(hashedOTP), []byte(otp))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response.NewResponse(errs.ErrInvalidVerificationCode.Error(), http.StatusUnauthorized))
		return
	}
	err = receiver.Service.VerifyEmail(userObj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.ErrDb)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// valid used to check email validation
func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// validateEmailAndPassword validates email and pass
func validateEmailAndPassword(userObj user.User) error {
	//err error()
	if userObj.Email == "" {
		return errs.ErrEmailMissing
	}
	if !valid(userObj.Email) {
		return errs.ErrInvalidEmail
	}
	if userObj.Password == "" {
		return errs.ErrInvalidPassword
	}
	return nil
}

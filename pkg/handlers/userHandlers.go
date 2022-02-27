package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tamiat/backend/pkg/emailVerification"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/mail"
	"strconv"

	"github.com/tamiat/backend/pkg/domain/user"
	"github.com/tamiat/backend/pkg/errs"
	"github.com/tamiat/backend/pkg/middleware"
	"github.com/tamiat/backend/pkg/service"
)

type UserHandlers struct {
	Service service.UserService
}

type JWT struct {
	Token string `json:"token"`
}

type Login struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
type Signup struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
	Role     string `json:"role" form:"role" binding:"required"`
}

//
// @Summary Signup endpoint
// @Description Provide email and password to login, response is JWT
// @Consume application/x-www-form-urlencoded
// @Produce application/json
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Param role formData string true "Role"
// @Success 200 {object} user.User
// @Failure 400  {object}  errs.ErrResponse "Bad Request"
// @Failure 500  {object}  errs.ErrResponse "Internal server error"
// @Router /signup [post]
func (receiver UserHandlers) Signup(ctx *gin.Context) {
	var userObj user.User
	var signupRequestData Signup
	//decoding request body
	if err := ctx.ShouldBind(&signupRequestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userObj.Email = signupRequestData.Email
	userObj.Password = signupRequestData.Password
	userObj.Role = signupRequestData.Role
	//encrypting password
	hash, err := bcrypt.GenerateFromPassword([]byte(userObj.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userObj.Password = string(hash)
	//database connection

	userObj.ID, err = receiver.Service.Signup(userObj)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	code := emailVerification.CodeGenerator()
	hashOTP, err := bcrypt.GenerateFromPassword([]byte(code), 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	userObj.Otp = string(hashOTP)
	err = receiver.Service.InsertOTP(&userObj)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	err = emailVerification.SendEmail(userObj.Email, code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//w.WriteHeader(http.StatusOK)
	userObj.Password = ""
	userObj.Otp = ""
	//json.NewEncoder(w).Encode(userObj)
	ctx.JSON(http.StatusOK, userObj)
}

//
// @Summary Login endpoint
// @Description Provide email and password to login, response is JWT
// @Consume application/x-www-form-urlencoded
// @Produce application/json
// @Param email formData string true "Email"
// @Param password formData string true "Password"
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

	if err := ctx.ShouldBind(&loginRequestData); err != nil {
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
	password := loginRequestData.Password
	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
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

type otp struct {
	Otp string `json:"otp" form:"otp" binding:"required"`
}

// @Summary verify user email
// @Description provide user id and otp sent to his email, it consists of 6 characters
// @Consume application/x-www-form-urlencoded
// @Produce  application/json
// @Param  id path int true "User ID"
// @Param otp formData string true "OTP"
// @Success 200 {object} response.Response
// @Failure 401 {pbject} errs.ErrResponse
// @Failure 500 {object} errs.ErrResponse "Internal server error"
// @Failure 400 {object} errs.ErrResponse "Bad request"
// @Router /confirmEmail/{id} [post]
func (receiver UserHandlers) VerifyEmail(ctx *gin.Context) {
	var userObj user.User
	var otpObj otp
	if err := ctx.ShouldBind(&otpObj); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(otpObj.Otp) != 6 {
		fmt.Println(otpObj.Otp)
		fmt.Println("error 1")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errs.ErrInvalidVerificationCode.Error()})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errs.ErrParsingID)
		return
	}
	userObj.ID = id
	hashedOTP, err := receiver.Service.ReadOTP(&userObj)
	err = bcrypt.CompareHashAndPassword([]byte(hashedOTP), []byte(otpObj.Otp))
	if err != nil {
		fmt.Println("error 2")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errs.ErrInvalidVerificationCode.Error()})
		return
	}
	err = receiver.Service.VerifyEmail(&userObj)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrDb.Error()})
		return
	}
	ctx.JSON(http.StatusOK, otpObj)
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

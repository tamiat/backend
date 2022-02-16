package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tamiat/backend/pkg/handlers"
)

type UserApi struct {
	userHandler handlers.UserHandlers
}

func NewUserAPI(userHandler handlers.UserHandlers) *UserApi {
	return &UserApi{
		userHandler: userHandler,
	}
}

//
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400  {string}  string  "ok"
// @Failure 404 {string}  string  "ok"
// @Router /signup [post]
func (api *UserApi) SignUpAPI(ctx *gin.Context) {
	userObj, code, err := api.userHandler.Signup(ctx)
	if err != nil {
		ctx.JSON(code, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(code, gin.H{
			"message": userObj,
		})
	}
}

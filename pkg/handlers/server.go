package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tamiat/backend/pkg/domain/user"
	"github.com/tamiat/backend/pkg/driver"
	"github.com/tamiat/backend/pkg/service"
)

func TestStart() {

	//handlers.Start()
	dbConnection, _ := driver.GetDbConnection()
	auth := driver.InitAuthority(dbConnection)
	usertHandler := UserHandlers{service.NewUserService(user.NewUserRepositoryDb(dbConnection, auth))}

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())
	server.POST("/api/v1/signup", func(ctx *gin.Context) {
		userObj, code, err := usertHandler.Signup(ctx)
		if err != nil {
			ctx.JSON(code, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(code, gin.H{
				"message": userObj,
			})
		}
	})
	server.Run("localhost:8080")

}

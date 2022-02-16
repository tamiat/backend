package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tamiat/backend/api"
	"github.com/tamiat/backend/docs"
	"github.com/tamiat/backend/pkg/domain/user"
	"github.com/tamiat/backend/pkg/driver"
	"github.com/tamiat/backend/pkg/handlers"
	"github.com/tamiat/backend/pkg/service"

	"log"
)

// @title Swagger Example API
// @version 1.0
// @description This is a documentation of our cms.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	//load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbConnection, _ := driver.GetDbConnection()
	auth := driver.InitAuthority(dbConnection)
	usertHandler := handlers.UserHandlers{service.NewUserService(user.NewUserRepositoryDb(dbConnection, auth))}
	userAPI := api.NewUserAPI(usertHandler)

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())
	docs.SwaggerInfo_swagger.BasePath = "/api/v1"

	server.POST("/api/v1/signup", userAPI.SignUpAPI)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	server.Run("localhost:8080")
}

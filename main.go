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

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	//load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	docs.SwaggerInfo_swagger.Title = "Pragmatic Reviews - Video API"
	docs.SwaggerInfo_swagger.Description = "Pragmatic Reviews - Youtube Video API."
	docs.SwaggerInfo_swagger.Version = "1.0"
	docs.SwaggerInfo_swagger.Host = "localhost:8080"
	docs.SwaggerInfo_swagger.BasePath = "/api/v1"
	docs.SwaggerInfo_swagger.Schemes = []string{"http"}

	dbConnection, _ := driver.GetDbConnection()
	auth := driver.InitAuthority(dbConnection)
	usertHandler := handlers.UserHandlers{service.NewUserService(user.NewUserRepositoryDb(dbConnection, auth))}
	userAPI := api.NewUserAPI(usertHandler)

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())
	docs.SwaggerInfo_swagger.BasePath = "/api/v1"

	server.POST("/api/v1/signup", userAPI.SignUpAPI)
	server.POST("/api/v1/login", usertHandler.Login)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	server.Run("localhost:8080")
}

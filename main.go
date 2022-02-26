package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tamiat/backend/docs"
	"github.com/tamiat/backend/pkg/domain/contentType"
	"github.com/tamiat/backend/pkg/domain/role"
	"github.com/tamiat/backend/pkg/domain/user"
	"github.com/tamiat/backend/pkg/driver"
	"github.com/tamiat/backend/pkg/handlers"
	"github.com/tamiat/backend/pkg/middleware"
	"github.com/tamiat/backend/pkg/service"
	"os"

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
	log.Println(os.Getenv("DEV_MODE"))
	docs.SwaggerInfo_swagger.Title = "TAMIAT-CMS"
	docs.SwaggerInfo_swagger.Description = "Content management system"
	docs.SwaggerInfo_swagger.Version = "1.0"
	docs.SwaggerInfo_swagger.Host = "localhost:8080"
	docs.SwaggerInfo_swagger.BasePath = "/api/v1"

	docs.SwaggerInfo_swagger.Schemes = []string{"http"}

	dbConnection, sqlDBConnection := driver.GetDbConnection()
	auth := driver.InitAuthority(dbConnection)
	usertHandler := handlers.UserHandlers{service.NewUserService(user.NewUserRepositoryDb(dbConnection, auth))}
	roleHandler := handlers.RoleHandlers{service.NewRoleService(role.NewRoleRepositoryDb(sqlDBConnection, auth))}
	contentTypeHandler := handlers.ContentTypeHandlers{service.NewContentTypeService(contentType.NewContentTypeRepositoryDb(dbConnection, sqlDBConnection, auth))}

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())

	apiRoutes := server.Group("/api/v1")
	{
		apiRoutes.POST("/signup", usertHandler.Signup)
		apiRoutes.POST("/login", usertHandler.Login)

		apiRoutes.POST("/roles", roleHandler.Create)
		rolesRoutes := apiRoutes.Group("/roles", middleware.TokenVerifyMiddleWare())
		{
			rolesRoutes.GET("", roleHandler.Read)
			rolesRoutes.DELETE(":id", roleHandler.Delete)
		}
		contentTypeRoutes := apiRoutes.Group("/contentType", middleware.TokenVerifyMiddleWare())
		{
			contentTypeRoutes.POST("/:userId", contentTypeHandler.CreateContentType)
			contentTypeRoutes.DELETE("/:userId/:contentTypeId", contentTypeHandler.DeleteContentType)
			contentTypeRoutes.PUT("/renamecol/:userId/:contentTypeId", contentTypeHandler.UpdateColName)
			contentTypeRoutes.PUT("/addcol/:userId/:contentTypeId", contentTypeHandler.AddCol)
			contentTypeRoutes.PUT("/delcol/:userId/:contentTypeId", contentTypeHandler.DeleteCol)
		}
	}

	if os.Getenv("DEV_MODE") == "true" {
		server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	server.Run("localhost:8080")
}

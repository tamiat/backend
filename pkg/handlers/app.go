package handlers

//this file is used to handle all business logic

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/tamiat/backend/pkg/domain/content"
	"github.com/tamiat/backend/pkg/domain/contentType"
	"github.com/tamiat/backend/pkg/domain/user"
	"github.com/tamiat/backend/pkg/middleware"
	"github.com/tamiat/backend/pkg/service"
)

func Start() {
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"content-type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	dbConnection, sqlDBConnection := getDbConnetion()
	contentHandler := ContentHandlers{service.NewContentService(content.NewContentRepositoryDb(dbConnection))}
	usertHandler := UserHandlers{service.NewUserService(user.NewUserRepositoryDb(dbConnection))}
	ct := ContentTypeHandlers{service.NewContentTypeService(contentType.NewContentTypeRepositoryDb(dbConnection, sqlDBConnection))}

		router.Path("/api/v1/contentType").
			HandlerFunc(ct.createContentType).Methods(http.MethodPost)

		router.Path("/api/v1/contentType").Queries("id", "{id}").
			HandlerFunc(ct.deleteContentType).Methods(http.MethodDelete)

		router.Path("/api/v1/contentType/renamecol").Queries("id", "{id}").
			HandlerFunc(ct.updateColName).Methods(http.MethodPut)

		router.Path("/api/v1/contentType/addcol").Queries("id", "{id}").
			HandlerFunc(ct.addCol).Methods(http.MethodPut)

		router.Path("/api/v1/contentType/delcol").Queries("id", "{id}").
			HandlerFunc(ct.deleteCol).Methods(http.MethodPut)

	router.HandleFunc("/api/v1/contents/", middleware.TokenVerifyMiddleWare(contentHandler.readAllContents)).Methods(http.MethodGet)
	router.Path("/api/v1/content").Queries("id", "{id}").
		HandlerFunc(middleware.TokenVerifyMiddleWare(contentHandler.readContent)).Methods(http.MethodGet)
	router.Path("/api/v1/contents").Queries("id", "{id}").
		HandlerFunc(middleware.TokenVerifyMiddleWare(contentHandler.readRangeOfContents)).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/content/", middleware.TokenVerifyMiddleWare(contentHandler.createContent)).Methods(http.MethodPost)
	router.Path("/api/v1/content").Queries("id", "{id}").
		HandlerFunc(middleware.TokenVerifyMiddleWare(contentHandler.deleteContent)).Methods(http.MethodDelete)
	router.Path("/api/v1/content").Queries("id", "{id}").
		HandlerFunc(middleware.TokenVerifyMiddleWare(contentHandler.updateContent)).Methods(http.MethodPut)

	router.HandleFunc("/api/v1/login", usertHandler.Login).Methods("POST")
	router.HandleFunc("/api/v1/signup", usertHandler.Signup).Methods("POST")
	router.Path("/api/v1/confirmEmail").Queries("id", "{id}").HandlerFunc(usertHandler.VerifyEmail).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8000", handlers.CORS(headers, methods, origins)(router)))
}
func getDbConnetion() (*gorm.DB, *sql.DB) {
	dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		os.Getenv("HOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"),
		os.Getenv("USER"),
		os.Getenv("PASS"))
	sqlDB, err := sql.Open("pgx", dataSourceName)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to conect to db"))
		panic(err)
	}
	log.Println("connected to db ")

	//test connection
	/*err = db.Ping()
	if err != nil {
		log.Fatal("cannot ping db")
		panic(err)
	}*/
	log.Println("pinged db")
	return db, sqlDB
}

package handlers

//this file is used to handle all business logic

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tamiat/backend/pkg/domain/content"
	"github.com/tamiat/backend/pkg/domain/user"
	"github.com/tamiat/backend/pkg/middleware"
	"github.com/tamiat/backend/pkg/service"
	"log"
	"net/http"
	"os"
)

func Start() {
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"content-type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	dbConnection := getDbConnetion()
	contentHandler := ContentHandlers{service.NewContentService(content.NewContentRepositoryDb(dbConnection))}
	usertHandler := UserHandlers{service.NewUserService(user.NewUserRepositoryDb(dbConnection))}

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
  
	log.Fatal(http.ListenAndServe("localhost:8080", handlers.CORS(headers,methods,origins)(router)))
}
func getDbConnetion() *sql.DB{
	dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		os.Getenv("HOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"),
		os.Getenv("USER"),
		os.Getenv("PASS"))
	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to conect to db"))
		panic(err)
	}
	log.Println("connected to db ")

	//test connection
	err = db.Ping()
	if err != nil {
		log.Fatal("cannot ping db")
		panic(err)
	}
	log.Println("pinged db")
	return db
}

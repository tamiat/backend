package handlers

//this file is used to handle all business logic

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tamiat/backend/pkg/domain/content"
	"github.com/tamiat/backend/pkg/domain/contentType"
	"github.com/tamiat/backend/pkg/domain/user"
	"github.com/tamiat/backend/pkg/driver"
	"github.com/tamiat/backend/pkg/middleware"
	"github.com/tamiat/backend/pkg/service"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"content-type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	dbConnection, sqlDBConnection := driver.GetDbConnetion()
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
	//router.HandleFunc("/api/v1/confirmEmail/", homePage).Methods("POST")
	//router.Path("/api/v1/confirmEmail/").Queries("id", "{id}").HandlerFunc(homePage).Methods(http.MethodGet)
	//router.HandleFunc("/api/v1/confirmEmail/", homePage).Methods(http.MethodGet)
	router.Path("/api/v1/confirmEmail/{id}").
		HandlerFunc(usertHandler.VerifyEmail).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
}


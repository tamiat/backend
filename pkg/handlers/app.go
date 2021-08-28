package handlers

//this file is used to handle all business logic

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tamiat/backend/pkg/domain/content"
	"github.com/tamiat/backend/pkg/domain/contentType"
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
	ch := ContentHandlers{service.NewContentService(content.NewContentRepositoryDb(dbConnection))}
	ct := ContentTypeHandlers{service.NewContentTypeService(contentType.NewContentTypeRepositoryDb(dbConnection))}

	router.HandleFunc("/api/v1/contents/", ch.readAllContents).Methods(http.MethodGet)

	router.Path("/api/v1/content").Queries("id", "{id}").
		HandlerFunc(ch.readContent).Methods(http.MethodGet)

	router.Path("/api/v1/contents").Queries("id", "{id}").
		HandlerFunc(ch.readRangeOfContents).Methods(http.MethodGet)

	router.HandleFunc("/api/v1/content/", ch.createContent).Methods(http.MethodPost)

	router.Path("/api/v1/content").Queries("id", "{id}").
		HandlerFunc(ch.deleteContent).Methods(http.MethodDelete)

	router.Path("/api/v1/content").Queries("id", "{id}").
		HandlerFunc(ch.updateContent).Methods(http.MethodPut)

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

	log.Fatal(http.ListenAndServe("localhost:8080", handlers.CORS(headers, methods, origins)(router)))
}

func getDbConnetion() *sql.DB {
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

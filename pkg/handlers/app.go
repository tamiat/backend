package handlers

import (
	"github.com/gorilla/mux"
	"github.com/tamiat/backend/pkg/domain/content"
	"github.com/tamiat/backend/pkg/service"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	ch := ContentHandlers{service.NewContentService(content.NewContentRepositoryDb())}
	router.HandleFunc("/api/v1/contents", ch.readAllContents).Methods(http.MethodGet)
	//read a content by id
	router.Path("/api/v1/content").Queries("id", "{id}").
		HandlerFunc(ch.readContent).Methods(http.MethodGet)

	//read range of contents
	router.Path("/api/v1/contents").Queries("id", "{id}").
		HandlerFunc(ch.readRangeOfContents).Methods(http.MethodGet)

	//post a content
	router.HandleFunc("/api/v1/content", ch.createContent).Methods(http.MethodPost)

	//
	router.Path("/api/v1/content").Queries("id", "{id}").
		HandlerFunc(ch.deleteContent).Methods(http.MethodDelete)

	router.Path("/api/v1/content").Queries("id", "{id}").
		HandlerFunc(ch.updateContent).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

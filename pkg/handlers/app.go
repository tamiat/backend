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
	router.HandleFunc("/api/v1/contents", ch.getAllContents).Methods(http.MethodGet)
	//get a content by id
	router.Path("/api/v1/content").Queries("id", "{id}").
		HandlerFunc(ch.getContent).Methods(http.MethodGet)

	//get range of contents
	router.Path("/api/v1/contents").Queries("id", "{id}").
		HandlerFunc(ch.getRangeOfContents).Methods(http.MethodGet)

	//post a content
	router.HandleFunc("/api/v1/content", ch.postContent).Methods(http.MethodPost)

	//
	router.Path("/api/v1/content").Queries("id", "{id}").
		HandlerFunc(ch.deleteContent).Methods(http.MethodDelete)

	router.Path("/api/v1/content").Queries("id", "{id}").
		HandlerFunc(ch.updateContent).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

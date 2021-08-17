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
	ch:= ContentHandlers{service.NewContentService(content.NewContentRepositoryDb())}
	router.HandleFunc("/contents",ch.getAllContents).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/content/{content_id:[0-9]+}",ch.getContent).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

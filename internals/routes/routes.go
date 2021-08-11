package routes

import (
	"github.com/gorilla/mux"
	"github.com/tamiat/backend/internals/handlers"
	"log"
	"net/http"
)

func Routes() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/content/{content_id}", handlers.GetContentById).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

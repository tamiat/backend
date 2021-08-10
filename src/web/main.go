package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/marwangalal746/backend/src/pkg/content"
	"github.com/marwangalal746/backend/src/pkg/handlers"
	"log"
	"net/http"
)

func addSomeContents() {
	content.Contents = append(content.Contents, content.Content{
		ID:    "0",
		Title: "First content",
		Details: "blablablablabla",
	})

	content.Contents = append(content.Contents, content.Content{
		ID:    "1",
		Title: "Second content",
		Details: "jajajajajajajaja",
	})
}


func main() {
	r := mux.NewRouter()
	addSomeContents()
	r.HandleFunc("/content/{id}", handlers.GetContent).Methods("GET")
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

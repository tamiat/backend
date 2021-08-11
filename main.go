package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type apiTest struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

var data []apiTest

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range data {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func main() {
	r := mux.NewRouter()

	data = append(data, apiTest{ID: "0", Title: "Hello World", Details: "dummy details"})
	data = append(data, apiTest{ID: "1", Title: "Hello World2", Details: "dummy details2"})

	r.HandleFunc("/apiTest/{id}", getData).Methods("GET")
	fmt.Printf("Starting server at port 8001\n")
	log.Fatal(http.ListenAndServe(":8001", r))
}

package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/marwangalal746/backend/src/pkg/content"
	"net/http"
)

func GetContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range content.Contents {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tamiat/backend/internals/models"
	"net/http"
	"strconv"
)



func GetContentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["content_id"])
	invalidUrl := false
	if err != nil {
		invalidUrl = true
	}
	if id >= len(models.ContentArr){
		invalidUrl = true
	}
	if invalidUrl {
		fmt.Fprintf(w, "Invalid Url")
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ContentArr[id])
}

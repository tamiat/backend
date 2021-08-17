package handlers

import (
	"encoding/json"
	"github.com/tamiat/backend/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ContentHandlers struct {
	service service.ContentService
}

func (ch *ContentHandlers)getAllContents(w http.ResponseWriter, r *http.Request) {
	contents,_:= ch.service.GetAllContents()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contents)
}
func (ch *ContentHandlers)getContent(w http.ResponseWriter,r *http.Request)  {
	vars:= mux.Vars(r)
	id:=vars["content_id"]
	log.Println(id)
	content,_ := ch.service.GetContent(id)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(content)

}
package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tamiat/backend/pkg/domain/content"
	"github.com/tamiat/backend/pkg/service"
	"log"
	"net/http"
	"regexp"
)

type ContentHandlers struct {
	service service.ContentService
}

func (ch *ContentHandlers) getAllContents(w http.ResponseWriter, r *http.Request) {
	contents, err := ch.service.GetAllContents()
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contents)
}
func (ch *ContentHandlers) getContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["content_id"]
	log.Println(id)
	content, _ := ch.service.GetContent(id)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(content)
}

func (ch *ContentHandlers) getRangeOfContents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//regular expression to check if the string in the pattern of this examples ([1:2], [35:40])
	pattern, _ := regexp.Match(`^([)([0-9]+)[:]([0-9]+)(])$`, []byte(params["id"]))
	//if the string can't match with any RG, the response will be 400 (badrequest)
	if !pattern {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	idValues := parseNums(params["id"])
	items, _ := ch.service.GetRangeOfContents(idValues)
	json.NewEncoder(w).Encode(items)
}

func (ch *ContentHandlers) postContent(w http.ResponseWriter, r *http.Request) {
	var newContent content.Content
	err := json.NewDecoder(r.Body).Decode(&newContent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
		return
	}
	id, err := ch.service.PostContent(newContent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
		return
	}
	type ID struct {
		ID string `json:"id"`
	}
	var IDobj ID
	IDobj.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(IDobj)
}

func (ch *ContentHandlers) deleteContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Println(id)
	_ = ch.service.DeleteContent(id)
	w.Header().Add("Content-Type", "application/json")
}

func (ch *ContentHandlers) updateContent(w http.ResponseWriter, r *http.Request) {
	var newContent content.Content
	err := json.NewDecoder(r.Body).Decode(&newContent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
		return
	}
	err = ch.service.UpdateContent(mux.Vars(r)["id"], newContent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
		return
	}
}

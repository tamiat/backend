// Package handlers, This package is used to implement the rest API operations
// we use a file for each domain object, here is the content handler
//which is responsible for crud operations of content object
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

func (ch *ContentHandlers) readAllContents(w http.ResponseWriter, r *http.Request) {
	contents, err := ch.service.ReadAllContents()
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contents)
}
func (ch *ContentHandlers) readContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	//regular expression to check if the string has numbers only	example: 1234
	pattern1, _ := regexp.Match(`^[0-9]+$`, []byte(vars["id"]))
	//if the string can't match with any RG, the response will be 400 (badrequest)
	if !pattern1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := vars["id"]
	log.Println(id)
	content, err := ch.service.ReadContent(id)
	if err != nil {
		var res response
		if err.Error() == "content not found" {
			res = response{Message: "This id is not found", Status: "404"}
			w.WriteHeader(http.StatusNotFound)

		} else {
			res = response{Message: err.Error(), Status: "503"}
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		json.NewEncoder(w).Encode(res)
		return

	}
	json.NewEncoder(w).Encode(content)
}

func (ch *ContentHandlers) readRangeOfContents(w http.ResponseWriter, r *http.Request) {
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
	items, _ := ch.service.ReadRangeOfContents(idValues)
	json.NewEncoder(w).Encode(items)
}

func (ch *ContentHandlers) createContent(w http.ResponseWriter, r *http.Request) {
	var newContent content.Content
	err := json.NewDecoder(r.Body).Decode(&newContent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
		return
	}
	id, err := ch.service.CreateContent(newContent)
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
	//regular expression to check if the string has numbers only	example: 1234
	pattern1, _ := regexp.Match(`^[0-9]+$`, []byte(vars["id"]))
	if !pattern1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := vars["id"]
	log.Println(id)
	_ = ch.service.DeleteContent(id)
	w.Header().Add("Content-Type", "application/json")
}

func (ch *ContentHandlers) updateContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//regular expression to check if the string has numbers only	example: 1234
	pattern1, _ := regexp.Match(`^[0-9]+$`, []byte(vars["id"]))
	if !pattern1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
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

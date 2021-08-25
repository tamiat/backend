package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tamiat/backend/pkg/service"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type ContentTypeHandlers struct {
	service service.ContentTypeService
}

func (ch *ContentTypeHandlers) createContentType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newContentType interface{} // The interface where we will save the converted JSON data.
	buffer, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
		return
	}
	err = r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
		return
	} // Close this
	err = json.Unmarshal(buffer, &newContentType)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
		return
	} // Convert JSON data into interface{} type
	m := newContentType.(map[string]interface{}) // To use the converted data we will need to convert it
	// into a map[string]interface{}
	var name, cols string
	name = ""
	for key, element := range m {
		if key == "name" {
			name = strings.TrimSpace(m["name"].(string))
		} else {
			cols += key
			cols += " "
			cols += strings.TrimSpace(element.(string))
			cols += ","
		}
	}
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response400("There is no content type name"))
		return
	}
	cols = cols[0 : len(cols)-1]
	var id string
	id, err = ch.service.CreateContentType(name, cols)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response500(err.Error()))
		return
	}
	type ID struct {
		ID string `json:"id"`
	}
	var IDobj ID
	IDobj.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(IDobj)
	return
}

func (ch *ContentTypeHandlers) deleteContentType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	//regular expression to check if the string has numbers only	example: 1234
	pattern1, _ := regexp.Match(`^[0-9]+$`, []byte(vars["id"]))
	if !pattern1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response400("Parameter value is not valid"))
		return
	}
	id := vars["id"]
	log.Println(id)
	err := ch.service.DeleteContentType(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response500(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response200("This content has been deleted successfully"))
	return
}



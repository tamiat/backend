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
	err = r.Body.Close()
	err = json.Unmarshal(buffer, &newContentType)
	m := newContentType.(map[string]interface{}) // To use the converted data we will need to convert it
	// into a map[string]interface{}
	var name, col string
	name = ""
	for key, element := range m {
		if key == "name" {
			name = strings.TrimSpace(m["name"].(string))
		} else {
			col += key
			col += " "
			col += strings.TrimSpace(element.(string))
			col += ","
		}
	}
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response400("There is no content type name"))
		return
	}
	col = col[0 : len(col)-1]
	var id string
	id, err = ch.service.CreateContentType(name, col)
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
		if err.Error() == "content not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response404(err.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response500(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response200("This content has been deleted successfully"))
	return
}

func (ch *ContentTypeHandlers) updateColName(w http.ResponseWriter, r *http.Request) {
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
	var newContentType interface{} // The interface where we will save the converted JSON data.
	buffer, err := ioutil.ReadAll(r.Body)
	err = r.Body.Close()
	err = json.Unmarshal(buffer, &newContentType)
	m := newContentType.(map[string]interface{}) // To use the converted data we will need to convert it
	// into a map[string]interface{}
	i := 0
	var oldName, newName string
	for key, element := range m {
		i++
		oldName = key
		newName = strings.TrimSpace(element.(string))
	}
	if i != 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response400("no specific column name was sent"))
		return
	}
	err = ch.service.UpdateColName(id, oldName, newName)
	if err != nil {
		if err.Error() == "content type not found" || err.Error() == "column not found"{
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response404(err.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response500(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response200("This column has been renamed successfully"))
	return
}

func (ch *ContentTypeHandlers) addCol(w http.ResponseWriter, r *http.Request) {
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
	var newContentType interface{} // The interface where we will save the converted JSON data.
	buffer, err := ioutil.ReadAll(r.Body)
	err = r.Body.Close()
	err = json.Unmarshal(buffer, &newContentType)
	m := newContentType.(map[string]interface{}) // To use the converted data we will need to convert it
	// into a map[string]interface{}
	var col string
	i := 0
	for key, element := range m {
		i++
		col += key
		col += " "
		col += strings.TrimSpace(element.(string))
	}
	if i != 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response400("no specific column name was sent"))
		return
	}
	err = ch.service.AddCol(id, col)
	if err != nil {
		if err.Error() == "content type not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response404(err.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response500(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response200("This new column has been added successfully"))
	return
}

func (ch *ContentTypeHandlers) deleteCol(w http.ResponseWriter, r *http.Request) {
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
	var newContentType interface{} // The interface where we will save the converted JSON data.
	buffer, err := ioutil.ReadAll(r.Body)
	err = r.Body.Close()
	err = json.Unmarshal(buffer, &newContentType)
	m := newContentType.(map[string]interface{}) // To use the converted data we will need to convert it
	// into a map[string]interface{}
	var col string
	i := 0
	for _, element := range m {
		i++
		col += strings.TrimSpace(element.(string))
	}
	if i != 1 || m["column name"] != col {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response400("no specific column name was sent"))
		return
	}
	err = ch.service.DeleteCol(id, col)
	if err != nil {
		if err.Error() == "content type not found" || err.Error() == "column not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response404(err.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response500(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response200("This column has been deleted successfully"))
	return
}

package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tamiat/backend/pkg/errs"
	"github.com/tamiat/backend/pkg/domain/type"
	"github.com/tamiat/backend/pkg/service"
)

type TypeHandlers struct {
	service service.TypeService
}

func (_typeHandler TypeHandlers) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var newType _type.Type
	_ = json.NewDecoder(r.Body).Decode(&newType)
	id, err := _typeHandler.service.Create(newType)
	newType.ID = uint(id)
	//handling errors
	if err != nil && err.Error() == `ERROR: duplicate key value violates unique constraint "types_name_key" (SQLSTATE 23505)` {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDuplicateValue.Error(), http.StatusBadRequest))
		return
	} else if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newType)
}

func (_typeHandler TypeHandlers) Read(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")
	_types,err:=_typeHandler.service.Read()
	//handling errors
	if err == sql.ErrNoRows || len(_types) == 0{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrNoTypesFound.Error(),http.StatusOK))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDb.Error(),http.StatusInternalServerError))
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(_types)
}


func (_typeHandler TypeHandlers) Update(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	id :=params["id"]
	var _type _type.Type
	_ = json.NewDecoder(r.Body).Decode(&_type)
	// validate inputs
	err := _typeHandler.service.Update(_type,id)
	//handling errors
	if err != nil {
		if err.Error() == `sql: no rows in result set`{
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrNoTypesFound.Error(),http.StatusBadRequest))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDb.Error(),http.StatusInternalServerError))
		}
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errs.NewResponse("Type has been updated successfully",http.StatusOK))
}

func (_typeHandler TypeHandlers) Delete(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	id :=params["id"]
	err := _typeHandler.service.Delete(id)
	//handling errors
	if err != nil {
		if err.Error() == `sql: no rows in result set`{
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrNoRowsFound.Error(),http.StatusBadRequest))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDb.Error(),http.StatusInternalServerError))
		}
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errs.NewResponse("Type has been deleted successfully",http.StatusOK))
}

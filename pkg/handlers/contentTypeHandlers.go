package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"

	"github.com/tamiat/backend/pkg/errs"
	"github.com/tamiat/backend/pkg/service"
)

type ContentTypeHandlers struct {
	service service.ContentTypeService
}

func (ch *ContentTypeHandlers) createContentType(ctx *gin.Context) {
	var newContentType map[string]interface{}
	err := ctx.ShouldBind(&newContentType)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrParsingID.Error()})
		return
	}
	//userId := params["userId"]
	//var newContentType interface{} // The interface where we will save the converted JSON data.
	/*buffer, err := ioutil.ReadAll(r.Body)
	err = r.Body.Close()
	err = json.Unmarshal(buffer, &newContentType)
	m := newContentType.(map[string]interface{})*/// To use the converted data we will need to convert it
	// into a map[string]interface{}
	var name, col string
	name = ""
	for key, element := range newContentType {
		if key == "name" {
			name = strings.TrimSpace(newContentType["name"].(string))
		} else {
			col += key
			col += " "
			col += strings.TrimSpace(element.(string))
			col += ","
		}
	}
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrNoContentTypeName.Error()})
		return
	}
	col = col[0 : len(col)-1]
	var id string
	id, err = ch.service.CreateContentType(userId, name, col)
	if err != nil {
		if err == errs.ErrUnauthorized {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": errs.ErrUnauthorized.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrServerErr.Error()})
		return
	}
	type ID struct {
		ID string `json:"id"`
	}
	var IDobj ID
	IDobj.ID = id
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(IDobj)
	ctx.JSON(http.StatusOK, IDobj)
	return
}

/*func (ch *ContentTypeHandlers) deleteContentType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["userId"]
	contentTypeId := params["contentTypeId"]
	err := ch.service.DeleteContentType(userId, contentTypeId)
	if err != nil {
		if err == errs.ErrContentNotFound {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response.NewResponse(err.Error(), http.StatusNotFound))
			return
		} else if err == errs.ErrUnauthorized {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response.NewResponse(err.Error(), http.StatusUnauthorized))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.NewResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.NewResponse("This content has been deleted successfully", http.StatusOK))
	return
}

func (ch *ContentTypeHandlers) updateColName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["userId"]
	contentTypeId := params["contentTypeId"]
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
		json.NewEncoder(w).Encode(response.NewResponse(errs.ErrColumnName.Error(), http.StatusBadRequest))
		return
	}
	err = ch.service.UpdateColName(userId, contentTypeId, oldName, newName)
	if err != nil {
		if err == errs.ErrContentNotFound || err == errs.ErrColNotFound {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response.NewResponse(err.Error(), http.StatusNotFound))
			return
		} else if err == errs.ErrUnauthorized {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response.NewResponse(err.Error(), http.StatusUnauthorized))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.NewResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.NewResponse("This column has been renamed successfully", http.StatusOK))
	return
}

func (ch *ContentTypeHandlers) addCol(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["userId"]
	contentTypeId := params["contentTypeId"]
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
		json.NewEncoder(w).Encode(response.NewResponse(errs.ErrColumnName.Error(), http.StatusBadRequest))
		return
	}
	err = ch.service.AddCol(userId, contentTypeId, col)
	if err != nil {
		if err == errs.ErrContentTypeNotFound {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response.NewResponse(err.Error(), http.StatusNotFound))
			return
		} else if err == errs.ErrUnauthorized {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response.NewResponse(err.Error(), http.StatusUnauthorized))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.NewResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.NewResponse("This new column has been added successfully", http.StatusOK))
	return
}

func (ch *ContentTypeHandlers) deleteCol(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["userId"]
	contentTypeId := params["contentTypeId"]
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
		json.NewEncoder(w).Encode(response.NewResponse(errs.ErrColumnName.Error(), http.StatusBadRequest))
		return
	}
	err = ch.service.DeleteCol(userId, contentTypeId, col)
	if err != nil {
		if err == errs.ErrContentNotFound || err == errs.ErrColNotFound {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response.NewResponse(err.Error(), http.StatusNotFound))
			return
		} else if err == errs.ErrUnauthorized {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response.NewResponse(err.Error(), http.StatusUnauthorized))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.NewResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.NewResponse("This column has been deleted successfully", http.StatusOK))
	return
}
*/
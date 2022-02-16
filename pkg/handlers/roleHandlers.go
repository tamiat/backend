package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/tamiat/backend/pkg/domain/role"
	"github.com/tamiat/backend/pkg/errs"
	"github.com/tamiat/backend/pkg/response"
	"github.com/tamiat/backend/pkg/service"
)

type RoleHandlers struct {
	service service.RoleService
}

func (roleHandler RoleHandlers) Create(ctx *gin.Context) {
	var newRole role.Role
	//decoding request body
	if err := ctx.ShouldBindJSON(&newRole); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// creating role in db
	id, err := roleHandler.service.Create(newRole)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrDb.Error()})
		return
	}
	newRole.ID = id
	//sending the response
	ctx.JSON(http.StatusOK, newRole)

}

func (roleHandler RoleHandlers) Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var roles []role.Role
	roles, err := roleHandler.service.Read()
	//handling errors
	if err == sql.ErrNoRows || len(roles) == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.NewResponse(errs.ErrNoRolesFound.Error(), http.StatusOK))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(roles)
}

func (roleHandler RoleHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	id := params["id"]
	tempId, err := strconv.Atoi(id)
	err = roleHandler.service.Delete(tempId)
	//handling errors
	if err != nil {
		if err.Error() == `sql: no rows in result set` {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.NewResponse(errs.ErrNoRowsFound.Error(), http.StatusBadRequest))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		}
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.NewResponse("Role has been deleted successfully", http.StatusOK))
}

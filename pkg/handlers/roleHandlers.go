package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/tamiat/backend/pkg/domain/role"
	"github.com/tamiat/backend/pkg/errs"
	"github.com/tamiat/backend/pkg/response"
	"github.com/tamiat/backend/pkg/service"
)

type RoleHandlers struct {
	Service service.RoleService
}

func (roleHandler RoleHandlers) Create(ctx *gin.Context) {
	var newRole role.Role
	//decoding request body
	if err := ctx.ShouldBind(&newRole); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// creating role in db
	id, err := roleHandler.Service.Create(newRole)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrDb.Error()})
		return
	}
	newRole.ID = id
	//sending the response
	ctx.JSON(http.StatusOK, newRole)

}

// @Security bearerAuth
// @Summary read roles endpoint
// @Description returns all roles
// @Accept  application/json
// @Produce  application/json
// @Success 200 {array} role.Role
// @Failure 401
// @Failure 500 {object} errs.ErrResponse "Internal server error"
// @Router /roles [get]
func (roleHandler RoleHandlers) Read(ctx *gin.Context) {
	//w.Header().Add("Content-Type", "application/json")
	var roles []role.Role
	roles, err := roleHandler.Service.Read()
	//handling errors
	if err == sql.ErrNoRows || len(roles) == 0 {
		ctx.JSON(http.StatusOK, errs.ErrNoRolesFound)
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.ErrDb)
		return
	}
	//sending the response
	ctx.JSON(http.StatusOK, roles)
}

func (roleHandler RoleHandlers) Delete(ctx *gin.Context) {
	// read id from path
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errs.ErrParsingID)
		return
	}
	err = roleHandler.Service.Delete(id)
	//handling errors
	if err != nil {
		if err.Error() == `sql: no rows in result set` {
			ctx.JSON(http.StatusBadRequest, errs.ErrNoRolesFound)
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, errs.ErrDb)
			return
		}
	}
	//sending the response
	var responseObj response.Response
	responseObj.Message = "Role has been deleted successfully"
	responseObj.Status = http.StatusOK
	ctx.JSON(http.StatusOK, responseObj)
}

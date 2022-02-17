package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tamiat/backend/pkg/response"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/tamiat/backend/pkg/errs"
	"github.com/tamiat/backend/pkg/service"
)

type ContentTypeHandlers struct {
	Service service.ContentTypeService
}

func (ch *ContentTypeHandlers) CreateContentType(ctx *gin.Context) {
	var newContentType map[string]interface{}
	err := ctx.ShouldBind(&newContentType)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrParsingID.Error()})
		return
	}
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
	id, err = ch.Service.CreateContentType(userId, name, col)
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
	ctx.JSON(http.StatusOK, IDobj)
	return
}
func (ch *ContentTypeHandlers) deleteContentType(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrParsingID.Error()})
		return
	}
	contentTypeId := ctx.Param("contentTypeId")
	_, err = strconv.Atoi(contentTypeId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrParsingID.Error()})
		return
	}
	err = ch.Service.DeleteContentType(userId, contentTypeId)
	if err != nil {
		log.Println(err)
		if err == errs.ErrContentNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": errs.ErrContentTypeNotFound.Error()})
			return
		} else if err == errs.ErrUnauthorized {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": errs.ErrUnauthorized.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrServerErr.Error()})
		return
	}
	res := response.Response{
		Message: "This content has been deleted successfully",
		Status:  200,
	}
	ctx.JSON(http.StatusOK, res)
}
func (ch *ContentTypeHandlers) updateColName(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrParsingID.Error()})
		return
	}
	contentTypeId := ctx.Param("contentTypeId")
	_, err = strconv.Atoi(contentTypeId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrParsingID.Error()})
		return
	}
	var newContentType map[string]interface{}
	err = ctx.ShouldBind(&newContentType)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	i := 0
	var oldName, newName string
	for key, element := range newContentType {
		i++
		oldName = key
		newName = strings.TrimSpace(element.(string))
	}
	if i < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrColumnName.Error()})
		return
	}
	if i > 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrColumnNameMoreThanOne.Error()})
		return
	}
	err = ch.Service.UpdateColName(userId, contentTypeId, oldName, newName)
	if err != nil {
		if err == errs.ErrContentNotFound || err == errs.ErrColNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": errs.ErrColNotFound})
			return
		} else if err == errs.ErrUnauthorized {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": errs.ErrUnauthorized.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errs.ErrServerErr.Error())
		return
	}
	res := response.Response{
		Message: "This column has been renamed successfully",
		Status:  200,
	}
	ctx.JSON(http.StatusOK, res)
}

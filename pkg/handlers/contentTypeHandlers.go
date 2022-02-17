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
type ID struct {
	ID string `json:"table_id"`
}

//
// @Summary Create content type endpoint
// @Description takes user id as path param to check his role and see if he is authorized to do this action, name is a required attribute
// @Accept application/json
// @Produce application/json
// @Param contentType body contentType.ContentTypeExample true "Content Type body"
// @Success 200 {object} handlers.ID
// @Failure 500  {object}  errs.ErrResponse "Internal server error"
// @Failure 400  {object}  errs.ErrResponse "Bad request"
// @Failure 401  {object}  errs.ErrResponse "Unauthorized error"
// @Router /contentType/{userId} [post]
func (ch *ContentTypeHandlers) CreateContentType(ctx *gin.Context) {
	var newContentType map[string]interface{}
	err := ctx.ShouldBind(&newContentType)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, err := strconv.Atoi(ctx.Param("userId"))
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

	var IDobj ID
	IDobj.ID = id
	ctx.JSON(http.StatusOK, IDobj)
	return
}

// @Security bearerAuth
// @Summary delete content type endpoint
// @Description takes userId and content type Id in path to delete content type
// @Accept  application/json
// @Produce  application/json
// @Param  userId path int true "User ID"
// @Param  contentTypeId path int true "Content Type ID"
// @Success 200 {object} response.Response
// @Failure 401 object} errs.ErrResponse "Unauthorized"
// @Failure 500 {object} errs.ErrResponse "Internal server error"
// @Failure 404 {object} errs.ErrResponse "Content type not found"
// @Failure 400 {object} errs.ErrResponse "Bad request"
// @Router /contentType/{userId}/{contentTypeId} [delete]
func (ch *ContentTypeHandlers) DeleteContentType(ctx *gin.Context) {
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

// @Security bearerAuth
// @Summary updates column name endpoint
// @Description takes userId and content type Id in path to update name of content type column
// @Accept  application/json
// @Produce  application/json
// @Param  userId path int true "User ID"
// @Param  contentTypeId path int true "Content Type ID"
// @Success 200 {object} response.Response
// @Failure 401 object} errs.ErrResponse "Unauthorized"
// @Failure 500 {object} errs.ErrResponse "Internal server error"
// @Failure 404 {object} errs.ErrResponse "Content type not found"
// @Failure 400 {object} errs.ErrResponse "Bad request"
// @Router /contentType/renamecol/{userId}/{contentTypeId} [put]
func (ch *ContentTypeHandlers) UpdateColName(ctx *gin.Context) {
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
			ctx.JSON(http.StatusNotFound, gin.H{"error": errs.ErrColNotFound.Error()})
			return
		} else if err == errs.ErrUnauthorized {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": errs.ErrUnauthorized.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrServerErr.Error()})
		return
	}
	res := response.Response{
		Message: "This column has been renamed successfully",
		Status:  200,
	}
	ctx.JSON(http.StatusOK, res)
}

// @Security bearerAuth
// @Summary adds column endpoint
// @Description takes userId and content type Id in path to add new column
// @Accept  application/json
// @Produce  application/json
// @Param  userId path int true "User ID"
// @Param  contentTypeId path int true "Content Type ID"
// @Success 200 {object} response.Response
// @Failure 401 object} errs.ErrResponse "Unauthorized"
// @Failure 500 {object} errs.ErrResponse "Internal server error"
// @Failure 404 {object} errs.ErrResponse "Content type not found"
// @Failure 400 {object} errs.ErrResponse "Bad request"
// @Router /contentType/addcol/{userId}/{contentTypeId} [put]
func (ch *ContentTypeHandlers) AddCol(ctx *gin.Context) {
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
	var col string
	i := 0
	for key, element := range newContentType {
		i++
		col += key
		col += " "
		col += strings.TrimSpace(element.(string))
	}
	if i < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrColumnName.Error()})
		return
	}
	if i > 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrColumnNameMoreThanOne.Error()})
		return
	}
	err = ch.Service.AddCol(userId, contentTypeId, col)
	if err != nil {
		if err == errs.ErrContentTypeNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if err == errs.ErrUnauthorized {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrServerErr.Error()})
		return
	}
	res := response.Response{
		Message: "This new column has been added successfully",
		Status:  200,
	}
	ctx.JSON(http.StatusOK, res)
}

// @Security bearerAuth
// @Summary deletes column endpoint
// @Description takes userId and content type Id in path to delete a column
// @Accept  application/json
// @Produce  application/json
// @Param  userId path int true "User ID"
// @Param  contentTypeId path int true "Content Type ID"
// @Success 200 {object} response.Response
// @Failure 401 object} errs.ErrResponse "Unauthorized"
// @Failure 500 {object} errs.ErrResponse "Internal server error"
// @Failure 404 {object} errs.ErrResponse "Content type not found"
// @Failure 400 {object} errs.ErrResponse "Bad request"
// @Router /contentType/delcol/{userId}/{contentTypeId} [delete]
func (ch *ContentTypeHandlers) DeleteCol(ctx *gin.Context) {
	type ColumnName struct {
		ColumnName string `json:"column_name" binding:"required"`
	}
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
	var colNameObj = ColumnName{}
	err = ctx.ShouldBindJSON(&colNameObj)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = ch.Service.DeleteCol(userId, contentTypeId, colNameObj.ColumnName)
	if err != nil {
		if err == errs.ErrContentNotFound || err == errs.ErrColNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if err == errs.ErrUnauthorized {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := response.Response{
		Message: "This column has been deleted successfully",
		Status:  200,
	}
	ctx.JSON(http.StatusOK, res)
}

type ContentTypeExample struct {
	Name        string `json:"name"`
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

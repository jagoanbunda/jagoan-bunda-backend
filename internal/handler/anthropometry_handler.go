package handler

import (
	// "net/http"

	// "net/http"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/dto"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/service"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/utils"
	// "github.com/jagoanbunda/jagoanbunda-backend/internal/utils"
)

type AnthropometryHandler interface {
	GetRecordFromChildID(c *gin.Context)
	CreateWithChildID(c *gin.Context)
	GetRecordByIDWithChildID(c *gin.Context)
	UpdateWithChildID(c *gin.Context)
	Delete(c *gin.Context)
}

type anthropometryHandler struct {
	service service.AnthropometryService
}

// Delete implements [AnthropometryHandler].
func (a *anthropometryHandler) Delete(c *gin.Context) {
	childID, err := utils.ParseUUIDFromParamsID(c, "childID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : "child id is not provided"})
		return
	}

	anthropometryID := c.Param("anthropometryID")
	if anthropometryID == ""{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : "anthropometry id is not provided"})
		return
	}
	parsedAnthropometryID, err := strconv.Atoi(anthropometryID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : "fail to convert to int"})
		return
	}

	if err := a.service.Delete(c.Request.Context(), *childID, uint(parsedAnthropometryID)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"err" : "record deleted"})
	return
}

// UpdateWithChildID implements [AnthropometryHandler].
func (a *anthropometryHandler) UpdateWithChildID(c *gin.Context) {
	var request dto.UpdateAnthropometryRequest
	childID, err := utils.ParseUUIDFromParamsID(c, "childID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	anthropometryID := c.Param("anthropometryID")
	if anthropometryID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "anthropometry id is not provided"})
		return
	}

	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	parsedAnthropometryID, err := strconv.Atoi(anthropometryID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	request.ChildID = *childID
	request.ID = uint(parsedAnthropometryID)

	response, err := a.service.UpdateWithChildID(c.Request.Context(), &request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
	return
}

// GetRecordByIDWithChildID implements [AnthropometryHandler].
func (a *anthropometryHandler) GetRecordByIDWithChildID(c *gin.Context) {
	childID, err := utils.ParseUUIDFromParamsID(c, "childID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	anthropometryID := c.Param("anthropometryID")
	if anthropometryID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "anthropometry id is not provided"})
		return
	}

	response, err := a.service.GetRecordByIDWithChildID(c.Request.Context(), anthropometryID, *childID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
	return
}

// CreateWithChildID implements [AnthropometryHandler].
func (a *anthropometryHandler) CreateWithChildID(c *gin.Context) {
	var request dto.CreateAnthropometryRequest
	childID := c.Param("childID")
	if childID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "no child id provided"})
		return
	}
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "failed to bind"})
		return
	}
	parsedChildID, err := uuid.Parse(childID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": "failed to parse child id to uid"})
		return
	}
	userInfo, exist := c.Get("userInfo")
	if !exist {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "bad claims"})
		return
	}

	data, ok := userInfo.(*utils.AccessTokenClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "bad claims"})
		return
	}

	request.ChildID = parsedChildID
	request.UserID = data.UserID

	response, err := a.service.CreateRecordWithChildID(c.Request.Context(), &request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	c.JSON(http.StatusCreated, response)

}

// GetRecordFromChildID implements [AnthropometryHandler].
func (a *anthropometryHandler) GetRecordFromChildID(c *gin.Context) {
	childID := c.Param("childID")
	if childID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "no child id is provided"})
	}

	childIDParsed, err := uuid.Parse(childID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err})
	}

	response, err := a.service.GetRecordFromChildID(c.Request.Context(), childIDParsed)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err})
	}

	c.JSON(http.StatusOK, response)

}

func NewAnthropometryHandler(service service.AnthropometryService) AnthropometryHandler {
	return &anthropometryHandler{service: service}
}

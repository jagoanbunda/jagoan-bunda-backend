package handler

import (
	// "net/http"

	// "net/http"

	"net/http"

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
}

type anthropometryHandler struct {
	service service.AnthropometryService
}

// CreateWithChildID implements [AnthropometryHandler].
func (a *anthropometryHandler) CreateWithChildID(c *gin.Context) {
	var request dto.CreateAnthropometryRequest
	childID := c.Param("childID")
	if childID == ""{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : "no child id provided"})
		return
	}
	userInfo, exist := c.Get("userInfo")
	if !exist{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : "bad claims"})
		return
	}

	data, ok := userInfo.(*utils.AccessTokenClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : "bad claims"})
		return
	}

	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : "failed to bind"})
		return
	}

	response, err := a.service.CreateRecordWithChildID(c.Request.Context(), &request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err" : err.Error()})
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

package handler

import (
	// "net/http"

	// "net/http"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/service"
	// "github.com/jagoanbunda/jagoanbunda-backend/internal/utils"
)

type AnthropometryHandler interface {
	GetRecordFromChildID(c *gin.Context)
}

type anthropometryHandler struct {
	service service.AnthropometryService
}

// GetRecordFromChildID implements [AnthropometryHandler].
func (a *anthropometryHandler) GetRecordFromChildID(c *gin.Context) {
	childID := c.Param("childID")
	if childID == ""{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : "no child id is provided"})
	}

	childIDParsed, err := uuid.Parse(childID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err" : err})
	}

	response, err := a.service.GetRecordFromChildID(c.Request.Context(), childIDParsed)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : err})
	}

	c.JSON(http.StatusOK, response)

}

func NewAnthropometryHandler(service service.AnthropometryService) AnthropometryHandler {
	return &anthropometryHandler{service: service}
}

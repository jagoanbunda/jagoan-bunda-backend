package handler

import (
	// "net/http"

	// "net/http"

	"github.com/gin-gonic/gin"
	// "github.com/jagoanbunda/jagoanbunda-backend/internal/dto"
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
	// userInfo, exist := c.Get("userInfo")
	// if !exist {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : "bad token claims"})
	// }
	// data, ok := userInfo.(*utils.AccessTokenClaims)
	// if !ok {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : "bad claims"})
	// 	return
	// }

	// record := a.service.GetRecord(c.Request.Context())

}

func NewAnthropometryHandler(service service.AnthropometryService) AnthropometryHandler {
	return &anthropometryHandler{service: service}
}

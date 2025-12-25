package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/service"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/utils"
)

type ChildHandler interface {
	Get(c *gin.Context)
}

type childHandler struct {
	service service.ChildService
}

// Get implements [ChildHandler].
func (cH *childHandler) Get(c *gin.Context) {

	userInfo, exist := c.Get("userInfo")
	if !exist {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error" : "bad claims"})
		return
	}
	data, ok := userInfo.(*utils.AccessTokenClaims)
	if !ok{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error" : "bad claims"})
		return
	}

	childID, err := uuid.Parse(c.Param("childID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	userID := data.UserID
	role := data.Role

	result, err := cH.service.GetChildWithAccess(c.Request.Context(), childID, userID, role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
	return
}

func NewChildHandler(service service.ChildService) ChildHandler {
	return &childHandler{service: service}
}

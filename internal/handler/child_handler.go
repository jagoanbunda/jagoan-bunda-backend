package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/dto"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/service"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/utils"
)

type ChildHandler interface {
	Get(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
}

type childHandler struct {
	service service.ChildService
}

// Get implements [ChildHandler].
func (cH *childHandler) Get(c *gin.Context) {
	userInfo, exist := c.Get("userInfo")
	if !exist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "token not provided"})
		return
	}

	data, ok := userInfo.(*utils.AccessTokenClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "bad claims"})
		return
	}

	response, err := cH.service.GetChildWithAccess(c.Request.Context(), data.UserID, data.Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "unauthorized"})
		return
	}

	 c.JSON(http.StatusOK, response)
	 return
}

// Create implements [ChildHandler].
func (cH *childHandler) Create(c *gin.Context) {
	var child dto.CreateChildRequest

	if err := c.ShouldBind(&child); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}

	parentID, exist := c.Get("userInfo")
	if !exist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "token not provided"})
		return
	}

	if data, ok := parentID.(*utils.AccessTokenClaims); !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "bad claims"})
		return
	} else {
		child.ParentID = data.UserID
	}

	childResponse, err := cH.service.Create(c.Request.Context(), &child)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, childResponse)
	return
}

// GetByID implements [ChildHandler].
func (cH *childHandler) GetByID(c *gin.Context) {

	userInfo, exist := c.Get("userInfo")
	if !exist {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad claims"})
		return
	}
	data, ok := userInfo.(*utils.AccessTokenClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad claims"})
		return
	}

	childID, err := uuid.Parse(c.Param("childID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := data.UserID
	role := data.Role

	result, err := cH.service.GetChildByIDWithAccess(c.Request.Context(), childID, userID, role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
	return
}

func NewChildHandler(service service.ChildService) ChildHandler {
	return &childHandler{service: service}
}

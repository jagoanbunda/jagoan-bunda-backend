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

	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type childHandler struct {
	service service.ChildService
}

// Delete implements [ChildHandler].
func (cH *childHandler) Delete(c *gin.Context) {
	childID := c.Param("childID")
	if childID == ""{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : "child id is not given"})
		return
	}

	parsedChildID, err := uuid.Parse(childID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err" : "child id parsing is failed"})
		return
	}

	if err := cH.service.Delete(c.Request.Context(), parsedChildID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err" : err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"err" : ""})
	return
}

// Update implements [ChildHandler].
func (cH *childHandler) Update(c *gin.Context) {
	// userInfo, exist := c.Get("userInfo")
	// if !exist{
	// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "token not provided"})
	// return
	// }

	var request dto.UpdateChildRequest
	if c.Param("childID") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "child_id not given"})
		return
	}
	parsedChildID, err := uuid.Parse(c.Param("childID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "uuid parse failed"})
		return
	}

	request.ID = parsedChildID
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "bad request"})
		return
	}

	response, err := cH.service.Update(c.Request.Context(), &request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "bad request"})
		return
	}

	c.JSON(http.StatusOK, response)
	return

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

	request, err := cH.service.GetChildWithAccess(c.Request.Context(), data.UserID, data.Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, request)
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

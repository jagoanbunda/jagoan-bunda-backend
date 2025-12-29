package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/dto"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/service"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/utils"
)

type FoodHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Search(c *gin.Context)
}

type foodHandler struct {
	service service.FoodService
}

// Search implements [FoodHandler].
func (f *foodHandler) Search(c *gin.Context) {
	key := c.Param("key")
	if key == ""{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "no key is provided"})
		return
	}

	responses, err := f.service.Search(c.Request.Context(), key)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses)
}

// Create implements [FoodHandler].
func (f *foodHandler) Create(c *gin.Context) {
	var request dto.FoodRequest

	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	response, err := f.service.Create(c.Request.Context(), &request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)

}

// Delete implements [FoodHandler].
func (f *foodHandler) Delete(c *gin.Context) {
	var request dto.FoodRequest

	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	foodID, err := utils.ParseUintFromParamsID(c, "foodID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err})
	}

	request.ID = foodID

	if err := f.service.Delete(c.Request.Context(), &request); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

// Get implements [FoodHandler].
func (f *foodHandler) Get(c *gin.Context) {
	response, err := f.service.Get(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)

}

// Update implements [FoodHandler].
func (f *foodHandler) Update(c *gin.Context) {
	var request dto.FoodRequest

	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	foodID, err := utils.ParseUintFromParamsID(c, "foodID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err})
	}

	request.ID = foodID

	response, err := f.service.Update(c.Request.Context(), &request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func NewFoodHandler(service service.FoodService) FoodHandler {
	return &foodHandler{service: service}
}

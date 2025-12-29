// Package utils is mostly for helpers, like creating and assessing jwt tokens, converting string to pointer, hashing password, file upload,  etc
package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func StringToPtr(s string) *string {
	if s != "" {
		return &s
	}
	return nil
}

func ParseUUIDFromParamsID(c *gin.Context, keys string) (*uuid.UUID, error){
	id := c.Param(keys)
	if id == ""{
		return nil, fmt.Errorf("%v is not provided", keys)
	}

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return &parsedID, nil
}

func ParseUintFromParamsID(c *gin.Context, keys string) (*uint, error) {
	id := c.Param(keys)

	if id == ""{
		return nil, fmt.Errorf("%v is not provided" , keys)
	}
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}
	result := uint(parsedID)
	return &result, nil

}

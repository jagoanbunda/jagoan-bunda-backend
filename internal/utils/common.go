// Package utils is mostly for helpers, like creating and assessing jwt tokens, converting string to pointer, hashing password, file upload,  etc
package utils

import (
	"fmt"

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

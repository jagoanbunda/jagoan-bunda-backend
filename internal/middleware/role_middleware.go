package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/domain"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/utils"
)

func RequireRole(requiredRole domain.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo, exist := c.Get("userInfo")
		if !exist {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": "claims didn't exist"})
			return
		}

		data, ok := userInfo.(*utils.AccessTokenClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": "error claim type"})
			return
		}
		if data.Role != requiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"err": fmt.Sprintf(`user with the role %v are not authorized to access a route for %v role`, data.Role, requiredRole)})
			return
		}
		c.Next()
	}
}

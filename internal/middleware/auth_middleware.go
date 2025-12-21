package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/utils"
)

func AuthenticateAccessToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString, &utils.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}

		secret := utils.GetEnv("JWT_SECRET_KEY", "jagoanbundah")
		return []byte(secret), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": err})
		return
	}

	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return
	}

	accessTokenClaims, ok := token.Claims.(*utils.AccessTokenClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "Bad Claims"})
		return
	}

	c.Set("userInfo", accessTokenClaims)
	c.Next()
}

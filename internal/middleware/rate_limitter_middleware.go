package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimitter(limitter *rate.Limiter) gin.HandlerFunc{
	return func(c *gin.Context){
		if limitter.Allow(){
			c.Next()
		} else {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message" : "limit exceeded",
			})
		}
	}

}

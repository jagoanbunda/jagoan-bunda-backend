package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/utils"
	"github.com/sirupsen/logrus"
)

func CustomLogger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()

		endTime := time.Now()
		statusCode := c.Writer.Status()
		latency := endTime.Sub(startTime)

		userID := "Guess"
		if val, exist := c.Get("userInfo"); exist {
			if claims, ok := val.(*utils.AccessTokenClaims); ok {
				userID = claims.Email
			}
		}

		logEntry := logger.WithFields(logrus.Fields{
			"status":     statusCode,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"ip":         c.ClientIP(),
			"latency":    fmt.Sprintf("%v ms", latency.Milliseconds()),
			"user_id":    userID,
			"user_agent": c.Request.UserAgent(),
		})

		if statusCode >= 500 {
			logEntry.Error("Internal Server Error")
		} else if statusCode >= 400 {
			logEntry.Warn("Client Error / Bad Request")
		} else {
			logEntry.Info("Success Request")
		}
	}
}

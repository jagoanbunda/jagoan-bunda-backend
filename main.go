package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/handler"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/middleware"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/repository"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/service"
	"github.com/jagoanbunda/jagoanbunda-backend/pkg/database"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

var logger *logrus.Logger
var limitter *rate.Limiter
var DB *gorm.DB

func init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	limitter = rate.NewLimiter(1, 4)

	DB = database.InitDB()
}

func main() {
	userRepository := repository.NewUserRepository(DB)
	authService := service.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)

	router := gin.Default()
	router.Use(middleware.CustomLogger(logger))
	router.Use(gin.Recovery())
	router.Use(middleware.RateLimitter(limitter))

	// v1
	v1Group := router.Group("/api/v1")

	// // auth group
	authGroup := v1Group.Group("/auth")
	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/refresh", authHandler.RefreshToken)

	// //

	if err := router.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("ERROR : %v", err.Error()))
	} else {
		fmt.Println("Server running at : http://0.0.0.0:8080/api/v1")
	}

}

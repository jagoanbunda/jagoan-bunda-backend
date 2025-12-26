package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/domain"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/handler"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/middleware"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/repository"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/service"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/utils"
	"github.com/jagoanbunda/jagoanbunda-backend/pkg/database"
	"github.com/joho/godotenv"
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
	if err := godotenv.Load(".env"); err != nil {
		panic(fmt.Errorf(`error loading env : %v`, err))
	}

	DB = database.InitDB()
}

func main() {

	// repos
	userRepository := repository.NewUserRepository(DB)
	childRepository := repository.NewChildRepository(DB)
	anthropometryRepository := repository.NewAnthropometryRepository(DB)

	// services
	authService := service.NewAuthService(userRepository)
	userService := service.NewUserService(userRepository)
	childService := service.NewChildService(childRepository)
	anthropometryService := service.NewAnthropometryService(anthropometryRepository)

	// handlers
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	childHandler := handler.NewChildHandler(childService)
	anthropometryHandler := handler.NewAnthropometryHandler(anthropometryService)

	router := gin.Default()
	router.Use(middleware.CustomLogger(logger))
	router.Use(gin.Recovery())
	router.Use(middleware.RateLimitter(limitter))

	// Static file serving for uploads
	router.Static("/uploads", utils.GetUploadDir())

	// v1
	v1Group := router.Group("/api/v1")

	// // auth group
	authGroup := v1Group.Group("/auth")
	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/refresh", authHandler.RefreshToken)

	// // user group
	userGroup := v1Group.Group("/user").Use(middleware.AuthenticateAccessToken)
	userGroup.GET("/me", userHandler.Get)
	userGroup.PUT("/profile", userHandler.UpdateProfile)

	childGroup := v1Group.Group("/children").Use(middleware.AuthenticateAccessToken)
	childGroup.GET("", childHandler.Get)
	childGroup.POST("", childHandler.Create)
	childGroup.GET("/:childID", childHandler.GetByID)
	childGroup.PUT("/:childID", childHandler.Update)
	childGroup.DELETE("/:childID", childHandler.Delete)
	// anthropometry stuff
	childGroup.GET("/:childID/anthropometry", anthropometryHandler.GetRecordFromChildID)
	childGroup.POST("/:childID/anthropometry", anthropometryHandler.CreateWithChildID).Use(middleware.RequireRole(domain.RoleNakes))



	if err := router.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("ERROR : %v", err.Error()))
	}

}

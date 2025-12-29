package main

import (
	"fmt"
	"time"

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
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/time/rate"
	"gorm.io/gorm"

	_ "github.com/jagoanbunda/jagoanbunda-backend/docs"
)

// @title Jagoan Bunda API
// @version 1.0
// @description API Backend untuk aplikasi Jagoan Bunda - Sistem monitoring tumbuh kembang anak
// @termsOfService http://swagger.io/terms/

// @contact.name Jagoan Bunda Team
// @contact.email support@jagoanbunda.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

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
	postgreDB, err := DB.DB()
	if err != nil {
		panic(fmt.Errorf("ERROR ON INSTANCING DB : %w", err))
	}

	postgreDB.SetMaxOpenConns(25)
	postgreDB.SetMaxIdleConns(10)
	postgreDB.SetConnMaxLifetime(time.Duration(time.Minute * 30))
	postgreDB.SetConnMaxIdleTime(time.Duration(time.Minute * 5))

}

func main() {

	// repos
	userRepository := repository.NewUserRepository(DB)
	childRepository := repository.NewChildRepository(DB)
	anthropometryRepository := repository.NewAnthropometryRepository(DB)
	foodRepository := repository.NewFoodRepository(DB)

	// services
	authService := service.NewAuthService(userRepository)
	userService := service.NewUserService(userRepository)
	childService := service.NewChildService(childRepository)
	anthropometryService := service.NewAnthropometryService(anthropometryRepository)
	foodService := service.NewFoodService(foodRepository)

	// handlers
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	childHandler := handler.NewChildHandler(childService)
	anthropometryHandler := handler.NewAnthropometryHandler(anthropometryService)
	foodHandler := handler.NewFoodHandler(foodService)

	gin.ForceConsoleColor()

	router := gin.Default()

	// Swagger documentation - placed before rate limiter to avoid 429 errors

	router.Use(middleware.CustomLogger(logger))
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	childGroup.GET("/:childID/anthropometry/:anthropometryID", anthropometryHandler.GetRecordByIDWithChildID)
	childGroup.PUT("/:childID/anthropometry/:anthropometryID", anthropometryHandler.UpdateWithChildID)
	childGroup.DELETE("/:childID/anthropometry/:anthropometryID", anthropometryHandler.Delete)

	// foods
	foodGroup := v1Group.Group("/foods").Use(middleware.AuthenticateAccessToken)
	foodGroup.GET("", foodHandler.Get)
	foodGroup.POST("", foodHandler.Create)
	foodGroup.PUT("/:foodID", foodHandler.Update)
	foodGroup.DELETE("/:foodID", foodHandler.Delete)
	foodGroup.GET("/search/:key", foodHandler.Search)

	if err := router.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("ERROR : %v", err.Error()))
	}
}

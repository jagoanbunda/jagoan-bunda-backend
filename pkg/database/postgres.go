package database

import (
	"fmt"

	"github.com/jagoanbunda/jagoanbunda-backend/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbUser := utils.GetEnv("DB_USER", "root")
	dbPassword := utils.GetEnv("DB_PASSWORD", "")
	dbHost := utils.GetEnv("DB_HOST", "127.0.0.1")
	dbPort := utils.GetEnv("DB_PORT", "5432")
	dbName := utils.GetEnv("DB_NAME", "jagoanbunda")
	timeZone := utils.GetEnv("TIMEZONE", "Asia/Jakarta")
	dbSSLMode := utils.GetEnv("DB_SSL_MODE", "disable")

	dsn := fmt.Sprintf(`host='%s' user='%s' password='%s' dbname='%s' port='%s' sslmode='%s' TimeZone=%s`, dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode, timeZone)
	var err error

	fmt.Println(dsn)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database : %v", err))
	}

	return DB

	//if err := DB.AutoMigrate(); err != nil {
	//	panic(fmt.Sprintf("Automigrate failed : %v", err))
	//}
}

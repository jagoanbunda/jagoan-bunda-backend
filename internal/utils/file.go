package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

// AllowedImageExtensions contains list of allowed image file extensions
var AllowedImageExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".webp": true,
}

// IsAllowedImageType checks if the file extension is an allowed image type
func IsAllowedImageType(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return AllowedImageExtensions[ext]
}

// GenerateUniqueFilename generates a unique filename with UUID prefix
func GenerateUniqueFilename(originalName string) string {
	ext := filepath.Ext(originalName)
	return fmt.Sprintf("%s%s", uuid.New().String(), ext)
}

// GetUploadDir returns the upload directory from environment variable or default
func GetUploadDir() string {
	dir := os.Getenv("UPLOAD_DIR")
	if dir == "" {
		dir = "./uploads"
	}
	return dir
}

// GetMaxFileSizeMB returns the max file size in MB from environment variable or default
func GetMaxFileSizeMB() int64 {
	// Default 5MB
	return 5
}

// GetBaseURL returns the base URL from environment variable
func GetBaseURL() string {
	url := os.Getenv("BASE_URL")
	if url == "" {
		url = "http://0.0.0.0:8080"
	}
	return url
}

// EnsureDir creates directory if it doesn't exist
func EnsureDir(dirPath string) error {
	return os.MkdirAll(dirPath, 0755)
}

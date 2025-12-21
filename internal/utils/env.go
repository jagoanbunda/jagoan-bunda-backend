package utils

import "os"

func GetEnv(key, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist && value != "" {
		return value
	}
	return defaultVal
}

package constant

import (
	"log"
	"os"
	"slices"
)

const APP_ENV_NAME = "APP_ENV"

var APP_ENV string

func init() {
	APP_ENV = GetEnv(APP_ENV_NAME, "dev")
	envList := []string{"dev", "prod", "test"}
	if slices.Contains(envList, APP_ENV) {
		return
	}
	log.Printf("Invalid APP_ENV(%s), supported values are %v, set to default 'dev'", APP_ENV, envList)
	APP_ENV = "dev"
}

func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

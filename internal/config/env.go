package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ENV              string
	IsProduction     bool
	IsTest           bool
	IsDevelopment    bool
	AppName          string
	AppURL           string
	AppURLClient     string
	ConnectionString string
	DbDriver         string
	Port             string
	RetryCount       int
}

var Env EnvConfig

func LoadEnv() {

	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	if env != "production" {
		godotenv.Load(".env." + env + ".local")
	}

	Env = EnvConfig{
		ENV:              getEnv("GO_ENV", "development"),
		IsProduction:     getEnv("GO_ENV", "development") == "production",
		IsTest:           getEnv("GO_ENV", "development") == "test",
		IsDevelopment:    getEnv("GO_ENV", "development") == "development",
		AppName:          getEnv("APP_NAME", ""),
		AppURL:           getEnv("APP_URL", ""),
		AppURLClient:     getEnv("APP_URL_CLIENT", ""),
		DbDriver:         getEnv("DB_Driver", "postgres"),
		ConnectionString: getEnv("DB_CONNECTION_STRING", "postgresql://root:secret@localhost:5432/test_db?sslmode=disable"),
		Port:             getEnv("PORT", "8080"),
		RetryCount:       getEnvAsInt("RETRY_COUNT", 15),
	}

}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valStr := getEnv(name, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultVal
}

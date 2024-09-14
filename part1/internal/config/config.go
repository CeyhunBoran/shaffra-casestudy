package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT       int
	DbHost     string
	DbPort     int
	DbName     string
	DbUser     string
	DbPass     string
	DbSsl      string
	DbTestPort int
	DbTestName string
}

var Conf *Config

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func InitConfig() {
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory")
	}
	loadErr := godotenv.Load(curDir + "/.env")
	if loadErr != nil {
		log.Fatal("Error loading .env file")
	}
	Conf = &Config{
		PORT:       getEnvAsInt("PORT", 8080),
		DbHost:     getEnv("DB_HOST", "localhost"),
		DbPort:     getEnvAsInt("DB_PORT", 5444),
		DbName:     getEnv("DB_NAME", "shaffra"),
		DbUser:     getEnv("DB_USER", "shaffra"),
		DbPass:     getEnv("DB_PASSWORD", ""),
		DbSsl:      getEnv("DB_SSL", "disable"),
		DbTestPort: getEnvAsInt("DB_TEST_PORT", 5442),
		DbTestName: getEnv("DB_TEST_NAME", "shaffratest"),
	}
}

func InitConfigTest() {
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory")
	}
	loadErr := godotenv.Load(curDir + "/../.env")
	if loadErr != nil {
		log.Fatal("Error loading .env file")
	}
	Conf = &Config{
		PORT:       getEnvAsInt("PORT", 8080),
		DbHost:     getEnv("DB_HOST", "localhost"),
		DbPort:     getEnvAsInt("DB_PORT", 5444),
		DbName:     getEnv("DB_NAME", "shaffra"),
		DbUser:     getEnv("DB_USER", "shaffra"),
		DbPass:     getEnv("DB_PASSWORD", ""),
		DbSsl:      getEnv("DB_SSL", "disable"),
		DbTestPort: getEnvAsInt("DB_TEST_PORT", 5442),
		DbTestName: getEnv("DB_TEST_NAME", "shaffratest"),
	}
}

package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	DB_USER string
	DB_PASSWORD string
	DB_HOST string
	DB_PORT string
	DB_NAME string
	SSL_MODE bool
}

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	Jwt_Secret string
	DBConfig *DbConfig
}

var configurations *Config

func loadConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	httpPort := os.Getenv("HttpPort")
	port, err := strconv.ParseInt(httpPort, 10, 64)

	if err != nil {
		log.Fatal("PORT must be a number")
	}

	version := os.Getenv("Version")
	serviceName := os.Getenv("ServiceName")
	jwtSecret := os.Getenv("Jwt_Secret")

	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	ssl := os.Getenv("ENABLE_SSL_MODE")

	SSL_MODE, err := strconv.ParseBool(ssl)

	if err != nil {
		fmt.Println("Error in parsing string to bool")
		os.Exit(1)
	}

	dbconfig := &DbConfig{
		DB_USER: DB_USER,
		DB_PASSWORD: DB_PASSWORD,
		DB_HOST: DB_HOST,
		DB_PORT: DB_PORT,
		DB_NAME: DB_NAME,
		SSL_MODE: SSL_MODE,
	}



	configurations = &Config{
		Version: version,
		HttpPort: int(port),
		ServiceName: serviceName,
		Jwt_Secret: jwtSecret,
		DBConfig: dbconfig,
	}

}

func GetConfig () *Config {
	if configurations == nil {
		loadConfig()
	}
	
	return configurations
}
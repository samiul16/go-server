package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

var configurations Config

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

	configurations = Config{
		Version: version,
		HttpPort: int(port),
		ServiceName: serviceName,
	}

}

func GetConfig () Config {
	loadConfig()
	return configurations
}
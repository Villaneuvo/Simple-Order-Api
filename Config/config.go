package config

import (
	models "order-api/Models"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() models.Config {
	err := godotenv.Load()

	if err != nil {
		panic("Can't load .env file")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	serverPort := os.Getenv("SERVICE_PORT")

	config := models.Config{
		ServerPort: serverPort,
		Database: models.Database{
			Username: dbUsername,
			Password: dbPassword,
			Host:     dbHost,
			Port:     dbPort,
			Name:     dbName,
		},
	}

	return config
}

package lib

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App      App
	Database Database
}

type App struct {
	Port      string
	SecretKey string
}

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func NewConfig() Config {
	return Config{
		App:      NewApp(),
		Database: NewDatabase(),
	}
}

func NewApp() App {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return App{
		Port:      os.Getenv("APP_PORT"),
		SecretKey: os.Getenv("SECRET_KEY"),
	}
}

func NewDatabase() Database {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Database{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}
}

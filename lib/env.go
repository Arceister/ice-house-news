package lib

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string
	APPPort    string
	SecretKey  string
}

func NewEnv() Env {
	env := Env{}
	godotenv.Load()
	env.loadEnv()
	return env
}

func (env *Env) loadEnv() {
	env.DBHost = os.Getenv("DB_HOST")
	env.DBPort = os.Getenv("DB_PORT")
	env.DBUsername = os.Getenv("DB_USERNAME")
	env.DBPassword = os.Getenv("DB_PASSWORD")
	env.DBName = os.Getenv("DB_NAME")
	env.APPPort = os.Getenv("APP_PORT")
	env.SecretKey = os.Getenv("SECRET_KEY")
}

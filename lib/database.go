package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	DB *pgxpool.Pool
}

func NewDatabase(env Env) Database {
	DBHost := env.DBHost
	DBPort := env.DBPort
	DBUsername := env.DBUsername
	DBPassword := env.DBPassword
	DBName := env.DBName

	DBUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DBUsername, DBPassword, DBHost, DBPort, DBName)

	db, err := pgxpool.Connect(context.Background(), DBUrl)

	if err != nil {
		fmt.Println("Unable to connect to database!")
		os.Exit(1)
	} else {
		fmt.Println("Connected to Database!")
	}

	return Database{DB: db}
}

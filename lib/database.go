package lib

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	DB *sql.DB
}

func NewDB(env Database) DB {
	DBHost := env.Host
	DBPort := env.Port
	DBUsername := env.Username
	DBPassword := env.Password
	DBName := env.Name

	DBUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DBUsername, DBPassword, DBHost, DBPort, DBName)

	db, err := sql.Open("pgx", DBUrl)

	if err != nil {
		fmt.Println("Unable to connect to database!")
		os.Exit(1)
	} else {
		fmt.Println("Connected to Database!")
	}

	return DB{DB: db}
}

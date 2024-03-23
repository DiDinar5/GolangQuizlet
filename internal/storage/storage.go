package storage

import (
	"GolangQuizlet/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func DbConnection() (*sql.DB, error) {
	user, host, port, password, dbname := config.LoadEnv()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

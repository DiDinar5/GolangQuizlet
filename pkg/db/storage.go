package db

import (
	"GolangQuizlet/internal/config"
	"GolangQuizlet/internal/domain"
	"database/sql"
	"encoding/json"
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
	CreateTable(db)
	return db, nil
}
func CreateTable(db *sql.DB) *domain.User {
	createTableQuery := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name TEXT,
            email TEXT,
            grade TEXT
        )
    `
	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	return &domain.User{}
}
func SaveUser(user domain.User) error {
	const q = `INSERT INTO users(id,name,email,grade) 
			VALUES ($1, $2, $3, $4)`

	args := []interface{}{
		user.Id,
		user.Name,
		user.Email,
		user.Grade,
	}

	db, err := DbConnection()
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec(q, args...)
	if err != nil {
		log.Panic(err)
	}
	/*id, err := res.LastInsertId()
	if err != nil {
		log.Panic(err)---------------------------добавить возврат id
	}*/
	defer db.Close()
	return nil
}

func GetData(id int) (map[string]interface{}, error) {
	db, err := DbConnection()
	if err != nil {
		log.Fatal(err)
	}
	var user domain.User
	rows, err := db.Query(`SELECT * FROM user WHERE id=$1`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.Id,
			&user.Name,
			&user.Email,
			&user.Grade,
		)
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

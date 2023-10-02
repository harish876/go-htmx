package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Init() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "films.db")
	if err != nil {
		return nil, err
	}
	log.Println("Connected to the Sqlite DB")

	return db, nil
}

func Seed() {
	db, _ := Init()
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS films (id INTEGER PRIMARY KEY,title TEXT,director TEXT)")
	statement.Exec()
	statement, _ = db.Prepare("INSERT INTO films (title,director) VALUES(?,?)")
	statement.Exec("Oldboy", "Harish Gokul")
}

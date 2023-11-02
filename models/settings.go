package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDb() {
	var err error
	db, err = sql.Open("postgres", "user=user1 host=localhost port=5432 password=haslo dbname=db1 sslmode=disable")
	if err != nil {
		fmt.Println("error during open sql")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("error during open sql")
	}

}

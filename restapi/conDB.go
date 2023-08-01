package restapi

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "../userDB.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

}

package database

import (
	"database/sql"
	"fmt"
	"log"
)


func ConnectDB(user, password, host, port, dbname string) *sql.DB {
	log.Println("Connecting to database:")
	// Database connection logic will be implemented here
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("mysql open error:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("mysql ping error:", err)
	}

	log.Println("Connected to database successfully")

	return db
}

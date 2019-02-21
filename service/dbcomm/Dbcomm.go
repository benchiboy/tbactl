package dbcomm

import (
	"database/sql"
	"log"
)

var (
	db *sql.DB
)

func InitDB() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/tba2_db")
	if err != nil {
		log.Println("Open database error:", err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Println("Ping database error:", err)
		return
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	log.Println("Database Connected successful!")
}

func GetDB() *sql.DB {
	return db
}

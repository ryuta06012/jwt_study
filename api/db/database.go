package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func OpenDB(driver, dsn string, count uint) *sql.DB {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal("OpenDB failed:", err)
	}

	if err = db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("retry... count:%v\n", count)
		return OpenDB(driver, dsn, count)
	}
	fmt.Println("db connected!!")
	return db
}

func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatal("CloseDB failed:", err)
	}
}

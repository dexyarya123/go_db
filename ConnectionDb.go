package belajar_golang_db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func getConnectionDb() *sql.DB {

	db, err := sql.Open("mysql", "root:rifkiganteng@/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(50 * time.Minute)
	return db
}

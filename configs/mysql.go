package configs

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sekolahku")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

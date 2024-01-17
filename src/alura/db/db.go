package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectWithDatabase() *sql.DB {
	connect := "user=root dbname=alura_loja password=root host=0.0.0.0 sslmode=disable"
	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err.Error())
	}
	return db
}

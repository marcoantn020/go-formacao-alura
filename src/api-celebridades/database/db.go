package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	dsn := "host=localhost user=root password=root dbname=app port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Ocorreu um erro ao conectar com o banco ", err)
	}
}

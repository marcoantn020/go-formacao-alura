package database

import (
	"gin/model"
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
		log.Panicln("Error in connect database ", err)
	}
	err := DB.AutoMigrate(&model.Student{})
	if err != nil {
		log.Panicln("error in create table", err)
	}
}

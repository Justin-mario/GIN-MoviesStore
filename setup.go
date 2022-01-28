package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "host=localhost user=postgres password=HUNters007 dbname=movies.db port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	//db, err := gorm.Open(postgresOpen(dsn), &gorm.Config{})
	database, err := gorm.Open("postgres", dsn)

	if err != nil {
		panic("Failed To Connect To Database")
	}
	database.AutoMigrate(&Movie{})
	DB = database
}

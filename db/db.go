package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)


var DB *gorm.DB

func Connect() {

    dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    
    if err != nil {
        panic(err.Error())
    }

    DB = db

}

func GetDB() *gorm.DB {
    return DB 
}

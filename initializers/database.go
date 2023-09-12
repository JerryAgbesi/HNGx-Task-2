package initializers

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	var err error

	DB, err = gorm.Open(sqlite.Open("database.db"),&gorm.Config{})

	if err != nil{
		log.Fatal("Failed to connect to database",err)
	}


}
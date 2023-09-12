package initializers

import (
	"jerryagbesi/hngxtask2/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)



func connectDatabase(){
	db,err := gorm.Open(sqlite.Open("database.db"),&gorm.Config{})

	if err != nil{
		log.Fatal("Failed to connect to database",err)
	}

	db.AutoMigrate(&models.Person{})

}
package main

import (
	"jerryagbesi/hngxtask2/initializers"
	"jerryagbesi/hngxtask2/models"
)

func init() {
	initializers.ConnectDatabase()

}

func main() {
	initializers.DB.AutoMigrate(&models.Person{})

}

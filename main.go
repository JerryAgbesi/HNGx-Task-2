package main

import (
	Controllers "jerryagbesi/hngxtask2/Controllers"
	"jerryagbesi/hngxtask2/initializers"
	"jerryagbesi/hngxtask2/utilities"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectDatabase()

}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(utilities.CORS())

    router.GET("/",Controllers.Root)
	router.POST("api/", Controllers.CreatePerson)
	router.GET("api/:id",Controllers.ReadPerson)
	router.GET("api/",Controllers.ReadPersonList) //returns a paginated list of records
	router.PUT("api/:id",Controllers.UpdatePerson)
	router.DELETE("api/:id",Controllers.DeletePerson)

	router.Run()

}

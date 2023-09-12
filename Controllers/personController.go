package Controllers

import (
	"jerryagbesi/hngxtask2/initializers"
	"jerryagbesi/hngxtask2/models"

	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Welcome to the HNGX REST API"})

}

func CreatePerson(c *gin.Context) {

	var body struct {
		Name string `json:"name"`
	}

	c.Bind(&body)

	newPerson := models.Person{Name: body.Name}

	query := initializers.DB.Create(&newPerson)

	if query.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"id":   newPerson.ID,
		"name": newPerson.Name})
}

func ReadPerson(c *gin.Context) {
	var person models.Person

	result := initializers.DB.First(&person, c.Param("id"))

	if result.Error != nil {
		c.Status(404)
		return
	}

	c.JSON(200, gin.H{
		"id":   person.ID,
		"name": person.Name})

}

func ReadPersonList(c *gin.Context) {
	var people []models.Person

	initializers.DB.Find(&people)

	c.JSON(200, gin.H{"persons":people})


}

func UpdatePerson(c *gin.Context) {
	var person models.Person

	result := initializers.DB.First(&person, c.Param("id"))

	if result.Error != nil {
		c.Status(404)
		return
	}

	var body struct {
		Name string `json:"name"`
	}

	c.Bind(&body)

	person.Name = body.Name

	initializers.DB.Save(&person)

	c.JSON(200, gin.H{
		"id":   person.ID,
		"name": person.Name})

}

func DeletePerson(c *gin.Context){
	var person models.Person

	result := initializers.DB.First(&person, c.Param("id"))

	if result.Error != nil {
		c.Status(404)
		return
	}

	initializers.DB.Delete(&person)

	c.Status(204)

}

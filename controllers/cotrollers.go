package controllers

import (
	"fmt"
	"net/http"

	"github.com/OmnGabriel/go-api-rest.git/database"
	"github.com/OmnGabriel/go-api-rest.git/models"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	fmt.Fprint(c.Writer, "Home Page")

}

func AllCharacters(c *gin.Context) {
	var character []models.Character
	database.DB.Find(&character)

	c.JSON(http.StatusOK, gin.H{"character": character})

}

func MakeANewCharacter(c *gin.Context) {
	var character models.Character
	if err := c.ShouldBindJSON(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&character)
	c.JSON(http.StatusCreated, character)
}

func DeletACharacter(c *gin.Context) {
	var character models.Character
	id := c.Params.ByName("id")
	database.DB.Delete(&character, id)
	c.JSON(http.StatusOK, gin.H{"data": "Character has been deleted"})

}

func ReturnACharacter(c *gin.Context) {
	var character models.Character
	id := c.Params.ByName("id")
	database.DB.First(&character, id)

	if character.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Character not afound"})
		return
	}

	c.JSON(http.StatusOK, character)
}

func EditACharacter(c *gin.Context) {
	var character models.Character
	id := c.Params.ByName("id")
	database.DB.First(&character, id)

	var characterToUpdate models.Character
	if err := c.ShouldBindJSON(&characterToUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	database.DB.Model(&character).Updates(characterToUpdate)
	c.JSON(http.StatusOK, gin.H{"data": "Successfully changed character"})
	c.JSON(http.StatusOK, character)
}

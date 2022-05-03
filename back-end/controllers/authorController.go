package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/back-end/database"
	"github.com/raelnogpires/libraryapp/back-end/models"
)

func getAllAuthors(c *gin.Context) {
	db := database.GetDB()
	var a []models.Author

	err := db.Find(&a).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": "internal server error - " + err.Error(),
		})
		return
	}

	c.JSON(200, a)
}

func GetAuthorById(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id must be an integer",
		})
	}

	db := database.GetDB()
	var a models.Author
	err = db.First(&a, intid).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "no author found - " + err.Error(),
		})
	}

	c.JSON(200, a)
}

func CreateAuthor(c *gin.Context) {
	db := database.GetDB()
	var a models.Author

	err := c.ShouldBindJSON(&a)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't bind json - " + err.Error(),
		})
	}

	err = db.Create(&a).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't create author - " + err.Error(),
		})
	}

	c.JSON(200, a)
}

func EditAuthor(c *gin.Context) {
	db := database.GetDB()
	var a models.Author

	err := c.ShouldBindJSON(&a)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't bind json - " + err.Error(),
		})
	}

	err = db.Save(&a).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't edit author - " + err.Error(),
		})
	}
}

func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id must be an integer",
		})
	}

	db := database.GetDB()
	err = db.Delete(&models.Author{}, intid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't delete author - " + err.Error(),
		})
	}

	c.Status(204)
}

package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/back-end/database"
	"github.com/raelnogpires/libraryapp/back-end/models"
)

func GetAllAuthors(c *gin.Context) {
	db := database.GetDB()
	var a []models.Author

	err := db.Find(&a).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": "internal server error",
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
		return
	}

	db := database.GetDB()
	var a models.Author
	err = db.First(&a, intid).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "author doesn't exist",
		})
		return
	}

	c.JSON(200, a)
}

func CreateAuthor(c *gin.Context) {
	db := database.GetDB()
	var a models.Author

	err := c.ShouldBindJSON(&a)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't bind json",
		})
		return
	}

	err = db.Create(&a).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't register author",
		})
		return
	}

	c.JSON(201, a)
}

func EditAuthor(c *gin.Context) {
	id := c.Param("id")
	// https://it-qa.com/how-to-convert-string-to-uint-in-golang/
	n, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id must be an integer",
		})
		return
	}

	db := database.GetDB()
	var a models.Author
	a.ID = uint(n)

	err = c.ShouldBindJSON(&a)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't bind json",
		})
		return
	}

	err = db.Save(&a).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "author doesn't exist",
		})
		return
	}

	c.JSON(200, a)
}

func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id must be an integer",
		})
		return
	}

	db := database.GetDB()
	err = db.Delete(&models.Author{}, intid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "author doesn't exist",
		})
		return
	}

	c.Status(204)
}

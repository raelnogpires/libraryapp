package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/back-end/database"
	"github.com/raelnogpires/libraryapp/back-end/models"
)

func GetAllBooks(c *gin.Context) {
	db := database.GetDB()
	var b []models.Book

	err := db.Find(&b).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": "internal server error - " + err.Error(),
		})
		return
	}

	c.JSON(200, b)
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id must be an integer",
		})
	}

	db := database.GetDB()
	var b models.Book
	err = db.First(&b, intid).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "no book found - " + err.Error(),
		})
	}

	c.JSON(200, b)
}

func CreateBook(c *gin.Context) {
	db := database.GetDB()
	var b models.Book

	err := c.ShouldBindJSON(&b)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't bind json - " + err.Error(),
		})
	}

	err = db.Create(&b).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't register book - " + err.Error(),
		})
	}

	c.JSON(201, b)
}

func EditBook(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id must be an integer",
		})
	}

	db := database.GetDB()
	var b models.Book
	b.Id = intid

	err = c.ShouldBindJSON(&b)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't bind json - " + err.Error(),
		})
	}

	err = db.Save(&b).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't edit book - " + err.Error(),
		})
	}

	c.JSON(200, b)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id must be an integer",
		})
	}

	db := database.GetDB()
	err = db.Delete(&models.Book{}, intid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't delete book - " + err.Error(),
		})
	}

	c.Status(204)
}

package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/back-end/database"
	"github.com/raelnogpires/libraryapp/back-end/models"
)

// ugly? kinda, but returns everything's necessary
var allQuery string = `SELECT
	b.id, b.name, b.description,
	c.id AS category_id, c.name AS category_name,
	a.id AS author_id, a.name AS author_name, b.img_url
	FROM books AS b
	INNER JOIN categories AS c
		ON b.category_id = c.id
	INNER JOIN authors AS a
		ON b.author_id = a.id
	ORDER BY b.id;`

func GetAllBooks(c *gin.Context) {
	db := database.GetDB()
	var b []models.FullBook

	err := db.Raw(allQuery).Scan(&b).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": "internal server error - " + err.Error(),
		})
		return
	}

	c.JSON(200, b)
}

// also ugly
var idQuery string = `SELECT
	b.id, b.name, b.description,
	c.id AS category_id, c.name AS category_name,
	a.id AS author_id, a.name AS author_name, b.img_url
	FROM books AS b
	INNER JOIN categories AS c
		ON b.category_id = c.id
	INNER JOIN authors AS a
		ON b.author_id = a.id
	WHERE b.id = ?
	ORDER BY b.id;`

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id must be an integer",
		})
		return
	}

	db := database.GetDB()
	var b models.Book

	err = db.Raw(idQuery, intid).Scan(&b).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "no book found - " + err.Error(),
		})
		return
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
		return
	}

	err = db.Create(&b).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't register book - " + err.Error(),
		})
		return
	}

	c.JSON(201, b)
}

func EditBook(c *gin.Context) {
	id := c.Param("id")
	n, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id must be an integer",
		})
		return
	}

	db := database.GetDB()
	var b models.Book
	b.ID = uint(n)

	err = c.ShouldBindJSON(&b)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't bind json - " + err.Error(),
		})
		return
	}

	err = db.Save(&b).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't edit book - " + err.Error(),
		})
		return
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
		return
	}

	db := database.GetDB()
	err = db.Delete(&models.Book{}, intid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't delete book - " + err.Error(),
		})
		return
	}

	c.Status(204)
}

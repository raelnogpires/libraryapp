package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/src/models"
	"github.com/raelnogpires/libraryapp/src/services"
)

func GetAllBooks(c *gin.Context) {
	books, err := services.GetAllBooks()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, books)
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id must be an integer",
		})
		return
	}

	book, err := services.GetBookById(intid)
	if err != nil {
		c.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request body",
		})
		return
	}

	err = services.CreateBook(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(201, book)
}

func EditBook(c *gin.Context) {
	id := c.Param("id")
	n, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id must be an integer number",
		})
		return
	}

	var book models.Book
	book.ID = uint(n)

	err = c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request body",
		})
		return
	}

	err = services.EditBook(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id must be an integer",
		})
		return
	}

	err = services.DeleteBook(intid)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "book doesn't exist",
		})
		return
	}

	c.Status(204)
}

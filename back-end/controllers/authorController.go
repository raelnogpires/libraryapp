package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/back-end/models"
	"github.com/raelnogpires/libraryapp/back-end/services"
)

func GetAllAuthors(c *gin.Context) {
	authors, err := services.GetAllAuthors()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, authors)
}

func GetAuthorById(c *gin.Context) {
	id := c.Param("id")
	// Atoi method converts string to int type
	intid, err := strconv.Atoi(id)
	if err != nil {
		// if id is alphabetical (a, x, s, d) returns error
		c.JSON(400, gin.H{
			"error": "id must be an integer",
		})
		return
	}

	author, err := services.GetAuthorById(intid)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, author)
}

func CreateAuthor(c *gin.Context) {
	var a models.Author

	err := c.ShouldBindJSON(&a)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid json data",
		})
		return
	}

	err = services.CreateAuthor(&a)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, a)
}

func EditAuthor(c *gin.Context) {
	id := c.Param("id")
	// https://it-qa.com/how-to-convert-string-to-uint-in-golang/
	n, err := strconv.ParseUint(id, 10, 64)
	// when id is receveid from params, comes as a string
	// is converted to uint because needs to match the model
	// for service layer
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id must be an integer",
		})
		return
	}

	var author models.Author
	author.ID = uint(n)

	err = c.ShouldBindJSON(&author)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid json data",
		})
		return
	}

	err = services.EditAuthor(&author)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, author)
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

	err = services.DeleteAuthor(intid)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(204)
}

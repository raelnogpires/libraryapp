package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/back-end/models"
	"github.com/raelnogpires/libraryapp/back-end/services"
)

func GetAllAuthors(c *gin.Context) {
	a, err := services.GetAllAuthors()
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

	a, err := services.GetAuthorById(intid)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "author doesn't exist",
		})
		return
	}

	c.JSON(200, a)
}

func CreateAuthor(c *gin.Context) {
	// needs middlewares for request validation
	var a models.Author

	err := c.ShouldBindJSON(&a)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't bind json",
		})
		return
	}

	err = services.CreateAuthor(&a)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't register author",
		})
		return
	}

	c.JSON(201, a)
}

func EditAuthor(c *gin.Context) {
	// needs middlewares for request validation
	id := c.Param("id")
	// https://it-qa.com/how-to-convert-string-to-uint-in-golang/
	n, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id must be an integer",
		})
		return
	}

	var a models.Author
	a.ID = uint(n)

	err = c.ShouldBindJSON(&a)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't bind json",
		})
		return
	}

	err = services.EditAuthor(&a)
	// returns the json of the edited author even if don't exist
	// but don't alters the db - kinda good thing.
	if err != nil {
		c.JSON(404, gin.H{
			"error": "author doesn't exist" + err.Error(),
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

	err = services.DeleteAuthor(intid)
	// returns 204 even if the author don't exist
	if err != nil {
		c.JSON(404, gin.H{
			"error": "author doesn't exist",
		})
		return
	}

	c.Status(204)
}

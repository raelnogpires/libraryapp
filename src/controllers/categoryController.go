package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/src/models"
	"github.com/raelnogpires/libraryapp/src/services"
)

func GetAllCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, categories)
}

func GetCategoryById(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id must be an integer",
		})
		return
	}

	category, err := services.GetCategoryById(intid)
	if err != nil {
		c.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, category)
}

func CreateCategory(c *gin.Context) {
	var ctg models.Category

	err := c.ShouldBindJSON(&ctg)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request body",
		})
		return
	}

	err = services.CreateCategory(&ctg)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(201, ctg)
}

func EditCategory(c *gin.Context) {
	id := c.Param("id")
	n, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id must be an integer",
		})
		return
	}

	var ctg models.Category
	ctg.ID = uint(n)

	err = c.ShouldBindJSON(&ctg)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request body",
		})
		return
	}

	err = services.EditCategory(&ctg)
	if err != nil {
		c.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, ctg)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id must be an integer",
		})
		return
	}

	err = services.DeleteCategory(intid)
	if err != nil {
		c.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(204)
}

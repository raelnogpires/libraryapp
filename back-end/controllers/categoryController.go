package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/back-end/database"
	"github.com/raelnogpires/libraryapp/back-end/models"
)

func GetAllCategories(c *gin.Context) {
	db := database.GetDB()
	var ctg []models.Category

	err := db.Find(&ctg).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(200, ctg)
}

func GetCategoryById(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id must be an integer",
		})
		return
	}

	db := database.GetDB()
	var ctg models.Category

	err = db.First(&ctg, intid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "category doesn't exist",
		})
		return
	}

	c.JSON(200, ctg)
}

func CreateCategory(c *gin.Context) {
	db := database.GetDB()
	var ctg models.Category

	err := c.ShouldBindJSON(&ctg)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't bind json",
		})
		return
	}

	err = db.Create(&ctg).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't register category",
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
			"error": "id must be an integer",
		})
		return
	}

	db := database.GetDB()
	var ctg models.Category
	ctg.ID = uint(n)

	err = c.ShouldBindJSON(&ctg)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't bind json",
		})
		return
	}

	err = db.Save(&ctg).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "category doesn't exist",
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
			"error": "id must be an integer",
		})
		return
	}

	db := database.GetDB()
	err = db.Delete(&models.Author{}, intid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "category doesn't exist",
		})
		return
	}

	c.Status(204)
}

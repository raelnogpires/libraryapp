package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/back-end/database"
	"github.com/raelnogpires/libraryapp/back-end/models"
)

func RegisterUser(c *gin.Context) {
	db := database.GetDB()
	var u models.User

	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't bind json",
		})
		return
	}

	err = db.Create(&u).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't register user",
		})
		return
	}

	c.Status(200)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id must be an integer",
		})
		return
	}

	db := database.GetDB()
	var u models.User

	err = db.First(&u, intid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "user doesn't exist",
		})
		return
	}

	c.JSON(200, u)
}

// needs to implement func DeleteMe

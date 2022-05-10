package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/back-end/auth"
	"github.com/raelnogpires/libraryapp/back-end/database"
	"github.com/raelnogpires/libraryapp/back-end/models"
	"github.com/raelnogpires/libraryapp/back-end/services"
)

func RegisterUser(c *gin.Context) {
	var u models.User

	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid json data",
		})
		return
	}

	u.Password = auth.SHA256Encoder(u.Password)

	err = services.RegisterUser(&u)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	n := uint64(u.ID)
	headerId := strconv.FormatUint(n, 10)
	c.Header("user_id", headerId)

	c.Status(201)
}

func Login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid json data",
		})
		return
	}

	err = services.Login(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	token, err := auth.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "internal server error",
		})
	}

	c.JSON(200, gin.H{
		"token": token,
	})
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

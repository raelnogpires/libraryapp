package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/back-end/auth"
	"github.com/raelnogpires/libraryapp/back-end/models"
	"github.com/raelnogpires/libraryapp/back-end/services"
)

func RegisterUser(c *gin.Context) {
	var u models.User

	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request body",
		})
		return
	}

	u.Password = auth.SHA256Encoder(u.Password)

	err = services.RegisterUser(&u)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(201)
}

func Login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request body",
		})
		return
	}

	err = services.Login(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	u, err := services.GetUser(user.Email)
	if err != nil {
		c.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := auth.NewJWTService().GenerateToken(u.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"token":   token,
		"user_id": u.ID,
	})
}

func DeleteMe(c *gin.Context) {
	n := c.GetHeader("user_id")
	id, err := strconv.Atoi(n)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "user_id must be a integer",
		})
		return
	}

	err = services.DeleteMe(id)
	if err != nil {
		c.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(204)
}

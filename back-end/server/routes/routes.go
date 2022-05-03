package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/back-end/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		authors := main.Group("authors")
		{
			authors.GET("/", controllers.GetAllAuthors)
			authors.GET("/:id", controllers.GetAuthorById)
			authors.POST("/", controllers.CreateAuthor)
			authors.PUT("/:id", controllers.EditAuthor)
			authors.DELETE("/:id", controllers.DeleteAuthor)
		}
	}

	return router
}

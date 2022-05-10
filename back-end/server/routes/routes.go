package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/back-end/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		register := main.Group("register")
		{
			register.POST("/", controllers.RegisterUser)
		}

		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}

		user := main.Group("user")
		{
			user.DELETE("/me", controllers.DeleteMe)
		}

		authors := main.Group("authors")
		{
			authors.GET("/", controllers.GetAllAuthors)
			authors.GET("/:id", controllers.GetAuthorById)
			authors.POST("/", controllers.CreateAuthor)
			authors.PUT("/:id", controllers.EditAuthor)
			authors.DELETE("/:id", controllers.DeleteAuthor)
		}

		books := main.Group("books")
		{
			books.GET("/", controllers.GetAllBooks)
			books.GET("/:id", controllers.GetBookById)
			books.POST("/", controllers.CreateBook)
			books.PUT("/:id", controllers.EditBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}

		categories := main.Group("categories")
		{
			categories.GET("/", controllers.GetAllCategories)
			categories.GET("/:id", controllers.GetCategoryById)
			categories.POST("/", controllers.CreateCategory)
			categories.PUT("/:id", controllers.EditCategory)
			categories.DELETE("/:id", controllers.DeleteCategory)
		}
	}

	return router
}

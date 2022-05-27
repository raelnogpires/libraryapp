package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/src/controllers"
	"github.com/raelnogpires/libraryapp/src/server/middlewares"
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

		user := main.Group("user", middlewares.Auth())
		{
			user.DELETE("/me", controllers.DeleteMe)
		}

		authors := main.Group("authors", middlewares.Auth())
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
			books.POST("/", middlewares.Auth(), controllers.CreateBook)
			books.PUT("/:id", middlewares.Auth(), controllers.EditBook)
			books.DELETE("/:id", middlewares.Auth(), controllers.DeleteBook)
		}

		categories := main.Group("categories", middlewares.Auth())
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

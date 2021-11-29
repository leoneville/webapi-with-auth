package routes

import (
	"webapi-with-go/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1") //Seta como path após a url principal (localhost:5000/api/v1)
	{
		user := main.Group("user")
		{
			user.POST("/", controllers.CreateUser)
		}

		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}

		books := main.Group("books") //localhost:5000/api/v1/books
		{
			books.GET("/:id", controllers.ShowBook)
			books.GET("/", controllers.ShowBooks)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}
	return router
}

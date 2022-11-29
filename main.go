package main

import (
	"github.com/crud-app/controllers"
	"github.com/crud-app/initializers"
	"github.com/crud-app/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	protected := r.Group("").Use(middlewares.Authz())
	{
		protected.POST("/createPost", controllers.PostCreate)
		protected.GET("/getPosts", controllers.PostsIndex)
	}

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

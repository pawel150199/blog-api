package main

import (
	"example/library_system/main/controllers"
	"example/library_system/main/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	// Posts
	router.POST("/posts", controllers.PostsCreate)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)
	router.PUT("/posts/:id", controllers.PostUpdate)
	router.DELETE("/posts/:id", controllers.PostDelete)

	// Users
	router.POST("/users", controllers.UserCreate)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)

	router.Run()
}

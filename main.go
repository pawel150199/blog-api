package main

import (
	"example/library_system/controllers"
	"example/library_system/initializers"
	"example/library_system/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	// Login
	router.POST("/login", controllers.Login)

	// Posts
	router.POST("/posts", middleware.RequireAuth, controllers.CreatePost)
	router.GET("/posts", controllers.GetPosts)
	router.GET("/posts/:id", controllers.GetPost)
	router.PUT("/posts/:id", middleware.RequireAuth, controllers.UpdatePost)
	router.DELETE("/posts/:id", middleware.RequireAuth, controllers.DeletePost)

	// Users
	router.POST("/users", controllers.CreateUser)
	router.GET("/me", middleware.RequireAuth, controllers.GetMe)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.PUT("/users/:id", middleware.RequireAuth, controllers.UpdateUser)
	router.DELETE("/users/:id", middleware.RequireAuth, controllers.DeleteUser)

	// Topics
	router.POST("/topics", middleware.RequireAuth, controllers.CreateTopic)
	router.GET("/topics", controllers.GetTopics)
	router.GET("/topics/:id", controllers.GetTopic)
	router.PUT("/topics/:id", middleware.RequireAuth, controllers.UpdateTopic)
	router.DELETE("/topics/:id", middleware.RequireAuth, controllers.DeleteTopic)

	router.Run()
}

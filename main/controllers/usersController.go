package controllers

import (
	"example/library_system/main/initializers"
	"example/library_system/main/models"

	"github.com/gin-gonic/gin"
)

// Create user
func UserCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Username  string
		Firstname string
		Lastname  string
		Password  string
	}

	c.Bind(&body)

	// Create a user
	user := models.User{
		Username:  body.Username,
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Password:  body.Password,
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return response
	c.JSON(200, gin.H{
		"user": user,
	})
}

// Return all users
func GetUsers(c *gin.Context) {
	// Get the users
	var users []models.User
	initializers.DB.Find(&users)

	// Respond data
	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetUser(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get the users
	var user models.User
	initializers.DB.First(&user, id)

	// Respond data
	c.JSON(200, gin.H{
		"user": user,
	})

}

package controllers

import (
	"example/library_system/initializers"
	"example/library_system/models"

	"github.com/gin-gonic/gin"
)

// Create user
func CreateUser(c *gin.Context) {
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

func UpdateUser(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get the data off the req body
	var body struct {
		Username  string
		Firstname string
		Lastname  string
		Password  string
	}

	c.Bind(&body)

	// Find the user were updating
	var user models.User
	initializers.DB.First(&user, id)

	// Update it
	initializers.DB.Model(&user).Updates(models.User{
		Username:  body.Username,
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Password:  body.Password,
	})

	// Respond with updates data
	c.JSON(200, gin.H{
		"user": user,
	})
}

func DeleteUser(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get deleting user
	var user models.User
	initializers.DB.First(&user, id)

	if user == (models.User{}) {
		c.JSON(404, gin.H{
			"message": "User does not exist.",
		})
	} else {
		// Delete user
		initializers.DB.Delete(&models.Post{}, id)

		// Respond
		c.JSON(200, gin.H{
			"user": user,
		})
	}

}

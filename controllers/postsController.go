package controllers

import (
	"example/library_system/initializers"
	"example/library_system/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Title string
		Body  string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})

}

func DeletePost(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get deleting post
	var post models.Post
	initializers.DB.First(&post, id)

	if post == (models.Post{}) {
		c.JSON(404, gin.H{
			"message": "Post does not exist.",
		})
	} else {
		// Delete post
		initializers.DB.Delete(&models.Post{}, id)

		// Respond
		c.JSON(200, gin.H{
			"post": post,
		})

	}
}

package controllers

import (
	"example/library_system/main/initializers"
	"example/library_system/main/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
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

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

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
		"posts": post,
	})

}

func PostDelete(c *gin.Context) {
	// Get id off url
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

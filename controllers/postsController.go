package controllers

import (
	"example/blog/initializers"
	"example/blog/models"

	"github.com/gin-gonic/gin"
)

// CreatePost godoc
// @Summary                    Create Post
// @Description                Add a new Post
// @Tags                       posts
// @Accept                     json
// @Produce                    json
// @Param                      post body     models.Post true "Post Data"
// @Success                    200  {object} models.Post
// @Router                     /posts [post]
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

// GetPosts godoc
// @Summary                 Get all Posts
// @Description             Get all Posts
// @Tags                    posts
// @Produce                 json
// @Router                  /posts [get]
func GetPosts(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// GetPostByID godoc
// @Summary                    Get Post by ID
// @Description                Get a Post by ID
// @Tags                       posts
// @Produce                    json
// @Param                      id path string true "Post ID"
// @Success      200  {object}  models.Post
// @Router                     /post/{id} [get]
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

// UpdatePostByID godoc
// @Summary                    Update Post
// @Description                Update Post by ID
// @Tags                       posts
// @Produce                    json
// @Param                      id path string      true "Post ID"
// @Param                      Post   body models.Post true "Post Data"
// @Router                     /posts/{id} [patch]
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

// DeletePostByID godoc
// @Summary                    Delete Post
// @Description                Delete Post by ID
// @Tags                       posts
// @Produce                    json
// @Param                      id path string true "Post ID"
// @Router                     /posts/{id} [delete]
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

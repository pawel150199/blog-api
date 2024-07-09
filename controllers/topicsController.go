package controllers

import (
	"example/blog/initializers"
	"example/blog/models"

	"github.com/gin-gonic/gin"
)

// CreateTopic godoc
// @Summary                    Create Topic
// @Description                Add a new Topic
// @Tags                       topics
// @Accept                     json
// @Produce                    json
// @Param                      user body     models.Topic true "Topic Data"
// @Success                    200  {object} models.Topic
// @Router                     /topics [post]
func CreateTopic(c *gin.Context) {
	// Get data off req body
	var body struct {
		Name  string
		Owner string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Create a topic
	topic := models.Topic{
		Name:  body.Name,
		Owner: body.Owner,
	}

	result := initializers.DB.Create(&topic)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"topic": topic,
	})
}

// GetTopics godoc
// @Summary                 Get all Topics
// @Description             Get all Topics
// @Tags                    topics
// @Produce                 json
// @Router                  /topics [get]
func GetTopics(c *gin.Context) {
	// Get the topics
	var topics []models.Topic
	initializers.DB.Find(&topics)

	//Respond with them
	c.JSON(200, gin.H{
		"topics": topics,
	})
}

// GetTopicByID godoc
// @Summary                    Get Topic
// @Description                Get a Topic by ID
// @Tags                       topics
// @Produce                    json
// @Param                      id path string true "Topic ID"
// @Success      200  {object}  models.Topic
// @Router                     /topics/{id} [get]
func GetTopic(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get the topic
	var topic models.Topic
	initializers.DB.First(&topic, id)

	// Respond with topic
	c.JSON(200, gin.H{
		"topic": topic,
	})
}

// UpdateTopicByID godoc
// @Summary                    Update Topic
// @Description                Update Topic by ID
// @Tags                       topics
// @Produce                    json
// @Param                      id path string      true "Topic ID"
// @Param                      Topic   body models.Topic true "Topic Data"
// @Router                     /topics/{id} [patch]
func UpdateTopic(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Name  string
		Owner string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Find the topic were updating
	var topic models.Topic
	initializers.DB.First(&topic, id)

	// Update it
	initializers.DB.Model(&topic).Updates(models.Topic{
		Name:  body.Name,
		Owner: body.Owner,
	})

	// Respond woth them
	c.JSON(200, gin.H{
		"topic": topic,
	})
}

// DeleteTopicByID godoc
// @Summary                    Delete Topic
// @Description                Delete Topic by ID
// @Tags                       topics
// @Produce                    json
// @Param                      id path string true "Topic ID"
// @Router                     /topics/{id} [delete]
func DeleteTopic(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get deleting topic
	var topic models.Topic
	initializers.DB.First(&topic, id)

	/*
		if topic == (models.Topic{}) {
			c.JSON(404, gin.H{
				"message": "Topic does not exist.",
			})
		}
	*/

	// Delete topic
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.JSON(200, gin.H{
		"topic": topic,
	})
}

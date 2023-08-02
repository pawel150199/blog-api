package controllers

import (
	"example/library_system/initializers"
	"example/library_system/models"

	"github.com/gin-gonic/gin"
)

func CreateTopic(c *gin.Context) {
	// Get data off req body
	var body struct {
		Name  string
		Owner string
	}

	c.Bind(&body)

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

func GetTopics(c *gin.Context) {
	// Get the topics
	var topics []models.Topic
	initializers.DB.Find(&topics)

	//Respond with them
	c.JSON(200, gin.H{
		"topics": topics,
	})
}

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

func UpdateTopic(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Name  string
		Owner string
	}

	c.Bind(&body)

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

func DeleteTopic(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get deleting topic
	var topic models.Topic
	initializers.DB.First(&topic, id)

	if topic == (models.Topic{}) {
		c.JSON(404, gin.H{
			"message": "Topic does not exist.",
		})
	} else {
		// Delete topic
		initializers.DB.Delete(&models.Post{}, id)

		// Respond
		c.JSON(200, gin.H{
			"topic": topic,
		})
	}
}

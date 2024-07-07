package controllers

import (
	"github.com/gin-gonic/gin"
)

// GetInfo godoc
// @Summary                    Get Information about service
// @Description                Get Information about service
// @Tags                       info
// @Produce                    json
// @Router                     /info [get]
func GetInfo(c *gin.Context) {
	// Respond with Body
	c.JSON(200, gin.H{
		"info": "Blog API v1",
	})
}

// GetHealthcheck godoc
// @Summary                    Get healthcheck info
// @Description                Get healthcheck info
// @Tags                       info
// @Produce                    json
// @Router                     /heathcheck [get]
func GetHealthcheck(c *gin.Context) {
	// Respond with Body
	c.JSON(200, gin.H{
		"status": "success",
	})
}

package controllers

import (
	"example/blog/helpers"
	"example/blog/initializers"
	"example/blog/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// GetLoggedUser godoc
// @Summary                    Get logged User
// @Description                Get a logged User
// @Tags                       users
// @Produce                    json
// @Success      200  {object}  models.User
// @Router                     /me [get]
func GetMe(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(200, gin.H{
		"user": user,
	})
}

func Login(c *gin.Context) {
	// Get the email and pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(400, gin.H{
			"error": "Invalid email or password.",
		})
		return
	}

	// Compare sent password with saved one
	isMatches := helpers.ChecksPasswordHash(user.Password, body.Password)

	// Generate jwt token
	if !isMatches {
		c.JSON(400, gin.H{
			"error": "Password is incorrect.",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(os.Getenv("SECRET"))

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid token.",
		})
		return
	}
	// Send jwt token back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(200, gin.H{
		"token": tokenString,
	})

}

// CreateUser godoc
// @Summary                    Create User
// @Description                Add a new User
// @Tags                       users
// @Accept                     json
// @Produce                    json
// @Param                      user body     models.User true "User Data"
// @Success                    200  {object} models.User
// @Router                     /users [post]
func CreateUser(c *gin.Context) {
	// Get the email and password off req body
	var body struct {
		Email     string
		Firstname string
		Lastname  string
		Password  string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Hash the password
	hashedPassword, err := helpers.CalculatePasswordHash(body.Password)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to hash password.",
		})
	}

	// Create the user
	user := models.User{
		Email:     body.Email,
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Password:  hashedPassword,
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Failed to create user.",
		})
	}

}

// GetUsers godoc
// @Summary                 Get all Users
// @Description             Get all Users
// @Tags                    users
// @Produce                 json
// @Router                  /users [get]
func GetUsers(c *gin.Context) {
	// Get the users
	var users []models.User
	initializers.DB.Find(&users)

	// Respond data
	c.JSON(200, gin.H{
		"users": users,
	})
}

// GetUserByID godoc
// @Summary                    Get User
// @Description                Get a User by ID
// @Tags                       users
// @Produce                    json
// @Param                      id path string true "User ID"
// @Success      200  {object}  models.User
// @Router                     /users/{id} [get]
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

// UpdateUserByID godoc
// @Summary                    Update User
// @Description                Update User by ID
// @Tags                       users
// @Produce                    json
// @Param                      id path string      true "User ID"
// @Param                      User   body models.User true "User Data"
// @Router                     /users/{id} [patch]
func UpdateUser(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get the data off the req body
	var body struct {
		Email     string
		Firstname string
		Lastname  string
		Password  string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Find the user were updating
	var user models.User
	initializers.DB.First(&user, id)

	// Update it
	initializers.DB.Model(&user).Updates(models.User{
		Email:     body.Email,
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Password:  body.Password,
	})

	// Respond with updates data
	c.JSON(200, gin.H{
		"user": user,
	})
}

// DeleteUserByID godoc
// @Summary                    Delete User
// @Description                Delete User by ID
// @Tags                       users
// @Produce                    json
// @Param                      id path string true "User ID"
// @Router                     /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get deleting user
	var user models.User
	initializers.DB.First(&user, id)

	/*
		if user == (models.User{}) {
			c.JSON(404, gin.H{
				"message": "User does not exist.",
			})
		}
	*/

	// Delete user
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.JSON(200, gin.H{
		"user": user,
	})
}

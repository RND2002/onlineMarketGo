package controllers

import (
	"net/http"

	"github.com/RND2002/onlineMarketGo/internals/config"
	"github.com/RND2002/onlineMarketGo/internals/model"
	"github.com/gin-gonic/gin"
)

func AuthController(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	db := config.GetDb()

	var user model.User

	// Check if the user exists
	if err := db.Where("username=?", requestBody.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Verify password
	if requestBody.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Authentication successful
	c.JSON(http.StatusOK, gin.H{"message": "Authentication successful"})
}

func BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract basic authentication credentials from the request header
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Perform authentication check
		if !checkCredentials(username, password) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		user := model.User{
			Username: username,
			Password: password,
		}
		c.Set("user", user)

		// If authentication succeeds, continue to the next handler
		c.Next()
	}
}

func checkCredentials(username, password string) bool {
	db := config.GetDb()

	var user model.User

	// Check if the user exists
	if err := db.Where("username=?", username).First(&user).Error; err != nil {
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return false
	}

	// Verify password
	if password != user.Password {
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return false
	}
	return true
}

//	func ProfileController(c * gin.Context){
//		BasicChecker(c);
//	}
func ProfileController(c *gin.Context) {
	// Extract the authenticated user from the context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found in context"})
		return
	}

	// Type assertion to convert the user to the appropriate type (assuming it's a model.User)
	authUser, ok := user.(model.User)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user type in context"})
		return
	}

	// Return the user's profile information
	c.JSON(http.StatusOK, gin.H{
		"id":       authUser.ID,
		"username": authUser.Username,
		// Add other profile information as needed
	})
}

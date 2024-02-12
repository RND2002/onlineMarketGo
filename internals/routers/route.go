package routers

import (
	"github.com/RND2002/onlineMarketGo/internals/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", controllers.AuthController)
	r.POST("/register/customer", controllers.RegisterCustomer)
	r.POST("/register/provider", controllers.RegisterProvider)

	authenticated := r.Group("/auth")
	authenticated.Use(controllers.BasicAuthMiddleware()) // Middleware to protect authenticated routes
	{
		// Define your authenticated routes here
		// Example:
		authenticated.GET("/profile", controllers.ProfileController)
	}

	return r
}

package controllers

import (
	"net/http"

	"github.com/RND2002/onlineMarketGo/internals/config"
	"github.com/RND2002/onlineMarketGo/internals/model"
	"github.com/gin-gonic/gin"
)

func RegisterCustomer(c *gin.Context) {
	// Define a struct to represent the customer data
	var customer model.Customer

	// Bind the request body to the customer struct
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Start a database transaction
	tx := config.GetDb().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := customer.Create(tx); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	username := customer.Username
	password := customer.Password
	role := "customer"

	var user model.User

	user = model.User{

		Username: username,
		Password: password,
		Role:     role,
	}
	if err := user.Create(tx); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Some internal error try again"})
		return
	}

	// Commit the transaction if all operations are successful
	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{"message": "Customer created successfully"})
}

package controllers

import (
	"net/http"

	"github.com/RND2002/onlineMarketGo/internals/config"
	"github.com/RND2002/onlineMarketGo/internals/model"
	"github.com/gin-gonic/gin"
)

func RegisterProvider(c *gin.Context) {

	var provider model.Provider

	if err := c.ShouldBindJSON(&provider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "err.Error()"})
		return
	}
	tx := config.GetDb().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := provider.Create(tx); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Provider"})
		return
	}

	username := provider.UserName
	password := provider.Password
	role := "provider"

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

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{"message": "Provider created successfully"})

}

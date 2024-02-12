package repository

import (
	"net/http"

	"github.com/RND2002/onlineMarketGo/internals/config"
	"github.com/RND2002/onlineMarketGo/internals/model"
	"github.com/gin-gonic/gin"
)

func AuthUser(username string, password string) {
	db := config.GetDb()

	var user model.User

	if err := db.Where("username=?", "username").Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

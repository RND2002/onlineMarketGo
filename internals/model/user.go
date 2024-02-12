package model

import (
	"github.com/RND2002/onlineMarketGo/internals/config"
	"gorm.io/gorm"
)

type User struct {
	ID uint `json:"-" gorm:"primary_key"`

	Username string `json:"username" gorm:"unique_index"`
	Password string `json:"password"`
	Role     string
}

func init() {
	config.Connect()
	db := config.GetDb()
	db.AutoMigrate(&User{})
}

func (user *User) Create(tx *gorm.DB) error {
	db := config.GetDb()
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

package model

import (
	"time"

	"github.com/RND2002/onlineMarketGo/internals/config"
	"gorm.io/gorm"
)

type Customer struct {
	ID        uint      `json:"-" gorm:"primary_key"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Username  string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
}

func init() {
	config.Connect()
	db := config.GetDb()
	db.AutoMigrate(&Customer{})

}

func (customer *Customer) Create(tx *gorm.DB) error {

	db := config.GetDb()

	if err := db.Create(&customer).Error; err != nil {
		return err
	}

	return nil
}

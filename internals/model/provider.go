package model

import (
	"time"

	"github.com/RND2002/onlineMarketGo/internals/config"
	"gorm.io/gorm"
)

type Provider struct {
	ID        uint      `json:"-" gorm:"primary_key"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	UserName  string    `json:"name"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	City      string    `json:"city"`
	Street    string    `json:"street"`
}

func init() {
	config.Connect()
	db := config.GetDb()
	db.AutoMigrate(&Provider{})
}
func (p *Provider) Create(tx *gorm.DB) error {
	db := config.GetDb()
	if err := db.Create(&p).Error; err != nil {
		return err
	}
	return nil
}

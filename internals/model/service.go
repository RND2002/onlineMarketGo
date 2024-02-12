package model

import "time"

type Service struct {
	ID          uint      `json:"-" gorm:"primary_key"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	Description string    `json:"description"`
	Charge      string    `json:"charge"`
	ProviderId  string    `json:"providerid"`
}

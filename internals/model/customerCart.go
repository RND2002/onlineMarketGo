package model

import "time"

type CustomerCart struct {
	ID        uint      `json:"-" gorm:"primary_key"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	ServId    []uint
}

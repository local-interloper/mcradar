package db

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Id        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

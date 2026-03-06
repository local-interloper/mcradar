package db

import (
	"time"

	"gorm.io/gorm"
)

type ScanResult struct {
	Ip            string `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	OnlinePlayers int
	MaxPlayers    int
	Version       string
	Type          string
	Players       []Player `gorm:"many2many:server_players"`
}

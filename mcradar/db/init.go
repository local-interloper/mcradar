package db

import (
	"context"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Ctx context.Context

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("data.db?_journal_mode=WAL&_busy_timeout=5000"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Ctx = context.Background()

	DB.AutoMigrate(&ScanResult{}, &Player{})

	db, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Hour)
}

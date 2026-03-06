package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Ctx context.Context

func Init() {
	var err error

	dsn := fmt.Sprintf("host=%s user=postgres password=%s port=5432 database=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	Ctx = context.Background()

	DB.AutoMigrate(&ScanResult{}, &Player{})

	db, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(512)
	db.SetMaxIdleConns(512)
	db.SetConnMaxLifetime(time.Hour)
}

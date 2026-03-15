package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/local-interloper/mcradar/mcradar/internal/settings"
	"github.com/local-interloper/mcradar/mcradar/internal/types/knownserverstore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Ctx context.Context

var KnownServers *knownserverstore.KnownServerStore = knownserverstore.New()

func Init() {
	var err error

	dsn := fmt.Sprintf("host=%s user=postgres password=%s port=5432 database=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}

	Ctx = context.Background()

	DB.AutoMigrate(&Server{}, &Player{})

	db, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(settings.Splits)
	db.SetMaxIdleConns(settings.Splits)
	db.SetConnMaxLifetime(time.Hour)

	servers, err := gorm.G[Server](DB).Select("ip").Find(Ctx)
	if err != nil {
		log.Fatal(err)
	}

	KnownServers.Mutex.Lock()
	defer KnownServers.Mutex.Unlock()
	for _, server := range servers {
		KnownServers.Store[server.Ip] = struct{}{}
	}
}

package main

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/local-interloper/mcradar/mcradar/internal/db"
	"github.com/local-interloper/mcradar/mcradar/internal/scanning"
	"github.com/local-interloper/mcradar/mcradar/internal/settings"
)

func main() {
	wg := &sync.WaitGroup{}

	godotenv.Load()

	settings.Init()

	db.Init()

	log.Println("Begginging a full IPv4 range scan using", settings.Splits, "goroutines")
	scanning.BeginFullRangeScan(wg)

	wg.Wait()
}

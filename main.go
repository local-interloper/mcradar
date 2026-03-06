package main

import (
	"github.com/joho/godotenv"
	"github.com/local-interloper/mc-radar/mcradar/db"
	"github.com/local-interloper/mc-radar/mcradar/scanning"
)

func main() {
	godotenv.Load()

	db.Init()

	scanning.BeginFullRangeScan()
}

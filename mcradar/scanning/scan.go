package scanning

import (
	"math"
	"strings"
	"sync"

	"github.com/local-interloper/mc-radar/mcradar/consts"
	"github.com/local-interloper/mc-radar/mcradar/db"
	"github.com/local-interloper/mc-radar/mcradar/types/mcconnection"
	"github.com/local-interloper/mc-radar/mcradar/types/servertype"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func BeginFullRangeScan() {
	var wg sync.WaitGroup

	for i := uint32(0); i < consts.Splits-1; i++ {
		segment := math.MaxUint32 / consts.Splits

		wg.Go(func() { ScanRange(db.DB.Session(&gorm.Session{}), segment*i, segment*(i+1)) })
	}

	wg.Wait()
}

func ScanRange(dbs *gorm.DB, from uint32, to uint32) {
	for n := from; n < to; n++ {
		addr := NumericIpToString(n)

		result := ScanServer(dbs, addr)
		if result == nil {
			continue
		}

		gorm.G[db.ScanResult](dbs, clause.OnConflict{DoNothing: true}).Create(db.Ctx, result)
	}
}

func ScanServer(dbs *gorm.DB, addr string) *db.ScanResult {
	if strings.HasPrefix(addr, "127") {
		return nil
	}

	params := mcconnection.Params{
		Address: addr,
		Port:    25565,
	}

	con, err := mcconnection.Connect(params)
	if err != nil {
		return nil
	}

	serverType, err := con.GetServerType()
	if serverType == servertype.Unknown || err != nil {
		return nil
	}

	con.Close()

	con, err = mcconnection.Connect(params)
	if err != nil {
		return nil
	}

	defer con.Close()

	status := con.Status()

	if status.Players.Max == 0 {
		return nil
	}

	var players []db.Player
	for _, player := range status.Players.Sample {
		players = append(players, db.Player{
			Id:   player.Id,
			Name: player.Name,
		})
	}

	return &db.ScanResult{
		Ip:            addr,
		Version:       status.Version.Name,
		OnlinePlayers: status.Players.Online,
		MaxPlayers:    status.Players.Max,
		Players:       players,
	}
}

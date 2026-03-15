package scanning

import (
	"math"
	"sync"

	"github.com/local-interloper/mcradar/mcradar/internal/db"
	"github.com/local-interloper/mcradar/mcradar/internal/settings"
	"github.com/local-interloper/mcradar/mcradar/internal/types/mcconnection"
	"github.com/local-interloper/mcradar/mcradar/internal/types/servertype"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func BeginFullRangeScan(wg *sync.WaitGroup) {
	for i := uint32(0); i < uint32(settings.Splits)-1; i++ {
		segment := uint32(math.MaxUint32 / settings.Splits)

		wg.Go(func() { ScanAndAddToDatabase(db.DB.Session(&gorm.Session{}), segment*i, segment*(i+1)) })
	}
}

func ScanAndAddToDatabase(dbs *gorm.DB, from uint32, to uint32) {
	for n := from; n < to; n++ {
		addr := NumericIpToString(n)

		db.KnownServers.Mutex.RLock()
		_, ok := db.KnownServers.Store[addr]
		db.KnownServers.Mutex.RUnlock()

		if ok {
			continue
		}

		result := ScanServer(dbs, addr)
		if result == nil {
			continue
		}

		err := gorm.G[db.Server](dbs, clause.OnConflict{DoNothing: true}).Create(db.Ctx, result)
		if err != nil {
			continue
		}

		db.KnownServers.Mutex.Lock()
		defer db.KnownServers.Mutex.Unlock()

		db.KnownServers.Store[addr] = struct{}{}
	}
}

func ScanServer(dbs *gorm.DB, addr string) *db.Server {
	if IsReserved(addr) {
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

	return &db.Server{
		Ip:            addr,
		Version:       status.Version.Name,
		OnlinePlayers: status.Players.Online,
		MaxPlayers:    status.Players.Max,
		Type:          servertype.Map[serverType],
		Players:       players,
	}
}

package mcstatus

type VersionModel struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type PlayerInfo struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type PlayersModel struct {
	Max    int          `json:"max"`
	Online int          `json:"online"`
	Sample []PlayerInfo `json:"sample"`
}

type McStatus struct {
	Version VersionModel `json:"version"`
	Players PlayersModel `json:"players"`
}

func (m McStatus) Print() {
	println("Version:", m.Version.Name)
	println("Players:", m.Players.Online, "/", m.Players.Max)
	println("Player list:")
	for _, player := range m.Players.Sample {
		println("\t", player.Name)
	}
}

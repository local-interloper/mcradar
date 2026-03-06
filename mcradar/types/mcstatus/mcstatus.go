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

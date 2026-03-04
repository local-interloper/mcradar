package servertype

const (
	Unknown ServerType = iota
	Legit
	Cracked
)

type ServerType int

var Map map[ServerType]string = map[ServerType]string{
	Unknown: "Unknown",
	Legit:   "Legit",
	Cracked: "Cracked",
}

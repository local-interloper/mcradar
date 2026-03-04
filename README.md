# mc-radar

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?style=flat-square&logo=go&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite-WAL-003B57?style=flat-square&logo=sqlite&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS-lightgrey?style=flat-square)

> A concurrent Minecraft server scanner that sweeps the entire IPv4 space and catalogs every public server it finds.

---

## What it does

mc-radar splits the full 32-bit IPv4 address space into 256 parallel goroutines, connects to port `25565` on each address using a hand-rolled Minecraft Java Edition protocol implementation, and stores every live server it finds into a local SQLite database.

For each discovered server it records:

- IP address
- Server version
- Online / max player counts
- Online player sample (names + UUIDs)
- Whether the server runs in online mode (legit) or offline mode (cracked)

## Requirements

- Go 1.25+
- CGO enabled (required by `go-sqlite3`)

## Getting started

```bash
git clone https://github.com/local-interloper/mc-radar
cd mc-radar
go build -o mc-radar .
./mc-radar
```

Results are written to `data.db` (SQLite, WAL mode) in the current directory.

## Configuration

The number of parallel scan goroutines is controlled by `mcradar/consts/splits.go`:

```go
const Splits uint32 = 256
```

The connection timeout per host is `500 ms` (`mcradar/types/mcconnection/mcconnection.go`).

## Tech stack

- **[GORM](https://gorm.io/)** — ORM with SQLite driver
- **Minecraft Java Edition protocol** — implemented from scratch (VarInt, VarLong, McString, packets)
- **`sync.WaitGroup`** — concurrent range scanning

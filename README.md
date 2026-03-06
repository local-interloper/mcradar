# mc-radar

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?style=flat-square&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-18-336791?style=flat-square&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=flat-square&logo=docker&logoColor=white)
![License](https://img.shields.io/badge/License-GPL--3.0-blue?style=flat-square)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS-lightgrey?style=flat-square)

> A concurrent Minecraft server scanner that sweeps the entire IPv4 space and catalogs every public server it finds.

---

## What it does

mc-radar splits the full 32-bit IPv4 address space into 256 parallel goroutines, connects to port `25565` on each address using a hand-rolled Minecraft Java Edition protocol implementation, and stores every live server it finds into a PostgreSQL database.

For each discovered server it records:

- IP address
- Server version
- Online / max player counts
- Online player sample (names + UUIDs)
- Whether the server runs in online mode (legit) or offline mode (cracked)

## Requirements

- Go 1.25+
- PostgreSQL instance (or use Docker Compose)

CGO is not required.

## Getting started

### With Docker Compose (recommended)

```bash
git clone https://github.com/local-interloper/mc-radar
cd mc-radar
cp .example.env .env
# Edit .env with your desired credentials
docker compose up --build
```

This starts both the scanner and a PostgreSQL 18 database.

### Manually

```bash
git clone https://github.com/local-interloper/mc-radar
cd mc-radar
cp .example.env .env
# Edit .env with your PostgreSQL connection details
go build -o mc-radar .
./mc-radar
```

## Configuration

All configuration is done via environment variables. Copy `.example.env` to `.env` and adjust as needed:

| Variable            | Description                        | Default           |
|---------------------|------------------------------------|-------------------|
| `POSTGRES_HOST`     | PostgreSQL host                    | `mc-radar-db`     |
| `POSTGRES_PASSWORD` | PostgreSQL password                | —                 |
| `POSTGRES_PORT`     | PostgreSQL port                    | `5432`            |
| `POSTGRES_DB`       | PostgreSQL database name           | `postgres`        |

The number of parallel scan goroutines is controlled by `mcradar/consts/splits.go`:

```go
const Splits uint32 = 256
```

The connection timeout per host is `500 ms` (`mcradar/types/mcconnection/mcconnection.go`).

## Tech stack

- **[GORM](https://gorm.io/)** — ORM with PostgreSQL driver (`pgx`)
- **[godotenv](https://github.com/joho/godotenv)** — `.env` file loading
- **Minecraft Java Edition protocol** — implemented from scratch (VarInt, VarLong, McString, packets)
- **`sync.WaitGroup`** — concurrent range scanning
- **Docker Compose** — containerized deployment with health-checked PostgreSQL

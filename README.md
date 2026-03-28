# mc-radar

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat-square&logo=go&logoColor=white)
![Python](https://img.shields.io/badge/Python-3.14-3776AB?style=flat-square&logo=python&logoColor=white)
![TypeScript](https://img.shields.io/badge/TypeScript-5.9-3178C6?style=flat-square&logo=typescript&logoColor=white)
![React](https://img.shields.io/badge/React-19-61DAFB?style=flat-square&logo=react&logoColor=black)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-18-336791?style=flat-square&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=flat-square&logo=docker&logoColor=white)
![License](https://img.shields.io/badge/License-GPL--3.0-blue?style=flat-square)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS-lightgrey?style=flat-square)

> A concurrent Minecraft server scanner that sweeps the entire IPv4 space, catalogs every public server it finds, and exposes the data through a REST API and web dashboard.

---

## What it does

mc-radar is a monorepo containing three services:

### mcradar (scanner)

Splits the full 32-bit IPv4 address space into a configurable number of parallel goroutines, connects to port `25565` on each address using a hand-rolled Minecraft Java Edition protocol implementation, and stores every live server it finds into a PostgreSQL database.

For each discovered server it records:

- IP address
- Server version
- Online / max player counts
- Online player sample (names + UUIDs)
- Whether the server runs in online mode (legit) or offline mode (cracked)

### mcradar-api

A FastAPI service that exposes the collected data over a REST API with bearer token authentication, pagination, and filtering.

### mcradar-gui

A React Router web dashboard that provides a visual interface for browsing the collected data. Includes an overview page with aggregate stats, and filterable/paginated tables for servers and players.

## Requirements

- Docker Compose (recommended)

Or, to run manually:

- Go 1.26+ (scanner)
- Python 3.14+ with [uv](https://github.com/astral-sh/uv) (API)
- Node.js 22+ (GUI)
- PostgreSQL instance

## Getting started

### With Docker Compose (recommended)

```bash
git clone https://github.com/local-interloper/mc-radar
cd mc-radar
cp .example.env .env
# Edit .env with your desired credentials
docker compose up --build
```

This starts the scanner, the API, the GUI, and a PostgreSQL 18 database.

### Manually

**Scanner:**

```bash
cd mcradar
cp ../.example.env .env
# Edit .env with your PostgreSQL connection details
go build -o mcradar ./cmd/mcradar/main.go
./mcradar
```

**API:**

```bash
cd mcradar-api
# Set environment variables (see Configuration below)
uv sync
uv run fastapi run app/main.py
```

The API listens on port `8000`.

**GUI:**

```bash
cd mcradar-gui
npm install
npm run dev
```

The GUI listens on port `3000`.

## Configuration

All configuration is done via environment variables. Copy `.example.env` to `.env` and adjust as needed:

| Variable                   | Description                              | Default       |
|----------------------------|------------------------------------------|---------------|
| `API_KEY`                  | Bearer token for API auth                | —             |
| `APP_WORKERS`              | Number of parallel scan goroutines       | —             |
| `APP_TIMEOUT_MS`           | Connection timeout per host (ms)         | —             |
| `APP_GUI_PORT`             | Host port for the GUI                    | `3000`        |
| `POSTGRES_HOST`            | PostgreSQL host                          | `mcradar-db`  |
| `POSTGRES_PASSWORD`        | PostgreSQL password                      | —             |
| `POSTGRES_DB`              | PostgreSQL database name                 | `postgres`    |
| `POSTGRES_MAX_CONNECTIONS` | Max PostgreSQL connections               | —             |

## API

All endpoints (except `/api/health`) require an `Authorization: Bearer <API_KEY>` header.

API endpoints are documented as a [Bruno](https://www.usebruno.com/) collection in the [`bruno-collection/`](bruno-collection/) directory. Import it into Bruno to explore and test the API interactively.

## Tech stack

### mcradar
- **[GORM](https://gorm.io/)** — ORM with PostgreSQL driver (`pgx`)
- **[godotenv](https://github.com/joho/godotenv)** — `.env` file loading
- **Minecraft Java Edition protocol** — implemented from scratch (VarInt, VarLong, McString, packets)
- **`sync.WaitGroup`** — concurrent range scanning with known-server deduplication cache

### mcradar-api
- **[FastAPI](https://fastapi.tiangolo.com/)** — async web framework
- **[psycopg v3](https://www.psycopg.org/psycopg3/)** — PostgreSQL driver with connection pooling
- **[pypika](https://github.com/kayak/pypika)** — SQL query builder
- **[uv](https://github.com/astral-sh/uv)** — package and project manager

### mcradar-gui
- **[React Router v7](https://reactrouter.com/)** — full-stack React framework (SSR)
- **[HeroUI](https://www.heroui.com/)** — component library
- **[Tailwind CSS v4](https://tailwindcss.com/)** — utility-first CSS framework
- **[Vite](https://vite.dev/)** — build tool

### Infrastructure
- **Docker Compose** — containerized deployment with health-checked PostgreSQL 18
- **[Bruno](https://www.usebruno.com/)** — API collection for testing and documentation

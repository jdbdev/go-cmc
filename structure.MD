# Project Structure (Monorepo)

```
.
├── ARCHITECTURE.md
├── docker-compose.yml
├── .env
├── .env.example
├── LICENSE
├── migrations
│   ├── coins
│   ├── collector
│   │   ├── 001_create_coin_info_table.down.sql
│   │   ├── 001_create_coin_info_table.up.sql
│   │   ├── 002_create_coin_quote_table.down.sql
│   │   └── 002_create_coin_quote_table.up.sql
│   └── website
├── README.md
├── services
│   ├── collector
│   │   ├── cmd
│   │   │   └── main.go
│   │   ├── config
│   │   │   └── config.go
│   │   ├── db
│   │   │   ├── manager.go
│   │   │   └── postgres.go
│   │   ├── Dockerfile
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── internal
│   │   │   ├── coins
│   │   │   │   ├── service.go
│   │   │   │   └── types.go
│   │   │   ├── mapper
│   │   │   │   ├── service.go
│   │   │   │   └── types.go
│   │   │   └── ticker
│   │   │       ├── service.go
│   │   │       ├── service_test.go
│   │   │       ├── types.go
│   │   │       └── types_reference.go
│   │   ├── sample_response.json
│   │   └── utils
│   │       ├── reference.go
│   │       └── write_files.go
│   └── website
│       ├── cmd
│       │   └── main.go
│       ├── config
│       │   └── config.go
│       ├── go.mod
│       └── internal
└── structure.MD

```
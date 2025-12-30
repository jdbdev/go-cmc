# Project Structure (Monorepo)

```
go-cmc/
├── .git/
├── .gitignore
├── .env
├── .env.example
├── docker-compose.yml
├── README.md
│
└── services/
    ├── collector/
    │   ├── cmd/
    │   ├── internal/
    │   ├── db/
    │   ├── config/
    │   ├── migrations/
    │   ├── utils/
    │   ├── Dockerfile
    │   ├── go.mod
    │   └── go.sum
    │
    └── website/
        ├── cmd/
        ├── internal/
        ├── templates/
        ├── static/
        ├── Dockerfile
        ├── go.mod
        └── go.sum
```
# Starting
You must create a `.env` file with the existing keys in `.env.example` in the `./infrastructure` directory

### 📋 Prerequisites
Tools: 

- [Docker](https://docs.docker.com/desktop/)
- [Golang](https://golang.org/doc/install)
- [Nodemon](https://nodemon.io/)

### How to run
```bash
first run 'make compose-up' to run the ecosystem
then ('make run' / 'make run-watch') or 'go run main.go -port=3005 -debug=true' to run the application
```

```bash
"args": ["-port=9000", "-debug=true"]
```

### Architecture
```
├── adapter
|   ├── primary
|   |   └── api
|   |       └── api.go
|   |       └── routes.go
|   └── secondary
|       └── postgres
|           └── postgres.go
|           └── sql.go
├── cmd
|   └── main.go
├── infrastructure
|   ├── docker-compose.yaml
|   ├── .env
|   ├── cloud-build-staging.yaml
|   └── cloud-build-prod.yaml   
└── internal
    ├── domain
    │   └── user
    |       └── user.go
    |       └── interface.go 
    ├── entity
    │   └── entity.go
    ├── hanlder
    |   └── user.go
    └── script
        └── migration
            └── ....sql
```

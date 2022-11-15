# Go Crud
This application it's a simple go crud using Gin and Gorm with postgres

## Setup
For this application you need to have a postgres instance running, for that you can use those docker commands.

```bash
docker pull postgres
```

```bash
docker run --name myPostgresDb -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DB=gocrud -d postgres
```

install the packages for this application with 

```bash
go mod tidy
```

## Running
for a live-reload run you can execute the command

```bash
CompileDaemon -command="./go-crud"
```


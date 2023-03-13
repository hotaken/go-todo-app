# REST API TODO lists

## USED
- Framework <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>
- Clean Architecture
- Dependency injection
- Docker
- DataBase: Postgres
- Migration files generation with migrate
- Logging: <a href="https://github.com/sirupsen/logrus">logrus</a>
- JWT: <a href="https://github.com/dgrijalva/jwt-go">jwt-go</a>
- ...
## Run app with docker-compose
Server will wait for postgres
```shell
docker-compose up --build go-todo-app
```
Stop
```shell
docker-compose down
```

## Some my notes
### Docker
Run postgres
```shell
docker pull postgres
docker run --name=todo-db -e POSTGRES_PASSWORD='postgres' -p 5436:5432 -d --rm postgres
docker ps
```

### Migrate
Installation
```shell
irm get.scoop.sh | iex
scoop install migrate
```
Migration files generation
```shell
migrate create -ext sql -dir ./schema -seq init
```
Apply migration to db
```shell
migrate -path ./schema -database 'postgres://postgres:postgres@0.0.0.0:5436/postgres?sslmode=disable' up
```
Check migration
```shell
docker ps
docker exec -it <id> powershell # /bin/bash
    psql -U postgres
        \d
```
Migrate down
```shell
migrate -path ./schema -database 'postgres://postgres:postgres@0.0.0.0:5436/postgres?sslmode=disable' down
```



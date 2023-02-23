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

### Docker
Run postgres
```shell
doker pull postgres
doker run --name=todo-db -e POSTGRES_PASSWORD='postgres' -p 5436:5432 -d --rm postgres
doker ps
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
doker ps
doker exec -it <id> powershell # /bin/bash
    psql -U postgres
        \d
```
Migrate down
```shell
migrate -path ./schema -database 'postgres://postgres:postgres@0.0.0.0:5436/postgres?sslmode=disable' down
```



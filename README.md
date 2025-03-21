# Guide
```
    Bahasa Pemrograman : Golang
    Library : github.com/gin-gonic/gin v1.10.0
	github.com/go-sql-driver/mysql v1.8.1
	github.com/gocql/gocql v1.7.0
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/google/uuid v1.6.0
	github.com/sirupsen/logrus v1.9.3
	gopkg.in/yaml.v2 v2.4.0
```

```
    Arsitekur : Depedency Injection dengan pendekatan DDD
```
## Run in local

```
go run main.go
```
### Run in docker

#### run mysql
```
docker run -d -p 3307:3306 -p 33060:33060 -e MYSQL_ROOT_PASSWORD=root --name db mysql:8.0

```
#### Build image docker 
```
docker build -f docker -t ahfrd/grpc-boiler-plate-go:v1.0 .
```

#### Run grpc-boiler-plate-go image on docker
```
docker run -d -p 9018:9018 -v config:/app/config --name grpc-boiler-plate-go-v1.0 ahfrd/grpc-boiler-plate-go:v1.0
```


#### Export Database
```
cd infra/database/migration/ 
goose mysql "root:root@tcp(127.0.0.1:3306)/loyalty?parseTime=true" up
```
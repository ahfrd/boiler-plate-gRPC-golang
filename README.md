# Guide

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
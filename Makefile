run:
	go run main.go


## docker 
up:
	docker-compose up -d

build:
	docker-compose up -d --build

kill:
	docker-compose kill

reload: kill up

## docker-clean: docker remove all containers in stack
clean:
	docker-compose rm -fv --all
	docker-compose down --rmi local --remove-orphans

net:
	docker network create my-network


## mysql: workspace container bash
db-bash:
	docker-compose exec sample-db bash

init:
	go get -u github.com/gin-gonic/gin
	go get -u github.com/go-sql-driver/mysql

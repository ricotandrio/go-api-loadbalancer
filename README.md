### 2602050620 - RICO TANDRIO
# Command that used in this API

## init go-workspace
go mod init cloud-api

## build docker image
docker build . -t go-service:latest

## show all docker image 
docker image ls

## run docker image
docker run go-service:latest

## run docker api in local machine at port 8080
docker run -p 8080:8080 [image_tag]

## tag image 
docker tag go-service:latest ricotandrio/dockerized-go-service:latest

## push docker image to docker-hub
docker push ricotandrio/dockerized-go-service:latest
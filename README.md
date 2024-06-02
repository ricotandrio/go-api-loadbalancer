# go-loadbalancer-service

This repository contains the codebase for images that I have pushed to Docker Hub as my submission for a cloud computing course.

Docker Hub: https://hub.docker.com/r/ricotandrio/go-api-loadbalancer

## Docker Commands
Commands that I used to build and push each image to Docker Hub.

### Main Image

```bash
cd main

docker build -t main_img .

docker tag main_img ricotandrio/go-api-loadbalancer:main_img

docker push ricotandrio/go-api-loadbalancer:main_img
```

### Alternative Image

```bash
cd alter

docker build -t alter_img .

docker tag alter_img ricotandrio/go-api-loadbalancer:alter_img

docker push ricotandrio/go-api-loadbalancer:alter_img
```

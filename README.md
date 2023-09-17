
# Kafka with golang


## Requirement

- [docker](https://www.docker.com/products/docker-desktop/)
- [golang](https://go.dev/dl/)
- [kafka](https://kafka.apache.org/downloads)
## Get Started
***To init the server following this step*** 
 
go to server directory

    cd server

start docker
    
    docker-compose up

to stop the server press Ctrl+c

***start the consumer***

open another terminal

    cd consumer

to execute go

    go run main.go

***start the producer***

open another terminal

    cd producer

to execute go

    go run main.go
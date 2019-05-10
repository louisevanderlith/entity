# entity
Entity.API holds simply identifiable information about a person or company

## Run with Docker
* $ docker build -t avosa/entity:dev .
* $ docker rm EntityDEV
* $ docker run -d -e RUNMODE=DEV -p 8097:8097 --network mango_net --name EntityDEV avosa/entity:dev
* $ docker logs EntityDEV
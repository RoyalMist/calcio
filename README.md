# Calcio

Table Football Manager

## External Dependencies

### GO

It is a go project make sure to have go installed locally in case you want to develop.

### Node & NPM

The frontend uses React make sure to have node with npm installed locally in case you want to develop.

### PostgreSQL

You must have a PostgreSQL server up and running.  
By default, the application will use the following connection string:  
`host=localhost port=5432 user=postgres dbname=calcio password=postgres sslmode=disable`
This can be overridden by setting the `DB_URL` env variable.

### Docker and Docker-Compose

If you want to use the docker-compose setup Docker and docker-compose are required.

### Make

The project provides a Makefile that can make the development experience easier.  
This is not a mandatory requirement but the current Readme assumes that you got make installed.

## Deployment

In order to log in a default admin user is created a first boot (admin / admin123).  
You can find a running deployment [here](https://calcio.alpchemist.ch)  
You can find the api documentation [here](https://calcio.alpchemist.ch/doc/index.html)  
The image is available on [Dockerhub](https://hub.docker.com/r/royalmist/calcio)  
If you want to run the project locally without any dependencies use the provided docker-compose file.  
It will set up a PostgreSQL server for you along with the application.  
Run `$ docker-compose up -d` and then go to http://localhost:4000

## Development

A development server and the go api are launched by the following commands:  
`npm i --prefix web`  
`$ make dev`  
Then go to http://localhost:3000 to benefit hot code reload on frontend.  
Api are available on http://localhost:4000.

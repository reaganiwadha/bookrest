# Setting up bookrest and bookest mongo using Docker

## Docker Topology
![Docker topology](https://i.ibb.co/cwTNbZq/Bookrest-Diagram.jpg)


## Prerequisites
Docker Installed

## Pulling both images
First step is to pull the docker images from this repo

[reaganiwadha/bookrest](https://hub.docker.com/repository/docker/reaganiwadha/bookrest)

[reaganiwadha/bookrest_mongo](https://hub.docker.com/repository/docker/reaganiwadha/bookrest_mongo)

```bash
docker pull reaganiwadha/bookrest
docker pull reaganiwadha/bookrest_mongo
```

## Running the database container

The database is mongoDB, bookrest_mongo is the image we are using to replicate the collections inside it.

```
docker run -d --name bookrest_mongo reaganiwadha/bookrest_mongo
```

You can also expose the port so you can manage the database or populate it

```
docker run -d --name bookrest_mongo -p 27017:27017 reaganiwadha/bookrest_mongo
```

This will run a container based on the bookrest_mongo image

## Running the logic container

Bookrest is an api that is build with golang, to run it we need to also link it to the database (NOTE: If you are running the container in a different name than bookrest_mongo, you will have to set a host environment variable that is BOOKREST_HOST)

And we need to also expose the port to this container.

```bash
docker run --name bookrest -p 8000:8000 --link bookrest_mongo reaganiwadha/bookrest
```

Now, bookrest should be available
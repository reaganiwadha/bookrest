FROM golang:onbuild


ENV BOOKREST_PORT 8000
ENV BOOKREST_DATABASE bookrest
ENV BOOKREST_DB_PORT 27017

# Change this to whatever your bookrest_mongo container name is
ENV BOOKREST_DB_HOST bookrest_mongo
# Remember to link it afterwards

ENV BOOKREST_DB_USERNAME ""
ENV BOOKREST_DB_PASSWORD ""


# ENV BOOKREST_CONNECTION_STRING = 

EXPOSE 8000

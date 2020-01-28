FROM golang:1.13-alpine

LABEL maintainer="Victor Nielsen"


ARG app_env="local"
ENV APP_ENV ${app_env}

ARG mongo_connection_string="mongodb://mongodb:20712"
ENV MONGO_CONNECTION_STRING ${mongo_connection_string}}


RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o server .

CMD ["/app/server"]

EXPOSE 1323
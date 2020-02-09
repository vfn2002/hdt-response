FROM golang:1.13-alpine

LABEL maintainer="Victor Nielsen"

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o server .

CMD ["/app/server"]

EXPOSE 1323
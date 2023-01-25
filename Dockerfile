# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN mkdir /certs
COPY /certs/* /certs/

RUN go build -o /docker-mtls-api

EXPOSE 8443
EXPOSE 8080

CMD [ "/docker-mtls-api" ]
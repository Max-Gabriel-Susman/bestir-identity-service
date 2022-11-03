# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ARG PROJECT_BINARY=rs-identity-service
ARG PROJECT_BUILD_DIR=./build/bin

WORKDIR /app

RUN ls

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./cmd/bestir-identity-service
COPY *.go ./internal/
#  cmd/bestir-identity-service
RUN go build -o /docker-gs-ping

EXPOSE 8080

 CMD [ "/docker-gs-ping" ]
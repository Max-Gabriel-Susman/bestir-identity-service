# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ARG PROJECT_BINARY=bestir-identity-service
ARG PROJECT_BUILD_DIR=./build/bin

WORKDIR /app

RUN ls

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# these are previous failed attempts to copy subdirectories of the working directory of docker invocation into the working 
# directory of 
# COPY *.go ./cmd/bestir-identity-service
# COPY *.go ./internal/
# cmd/bestir-identity-service

# temporary measure
COPY *.go ./

RUN go build -o /bestir-identity-service

EXPOSE 8080

 CMD [ "/bestir-identity-service" ]
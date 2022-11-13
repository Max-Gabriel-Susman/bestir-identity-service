FROM ubuntu:20.04


ENV GOLANG_VERSION 1.15.7

# Install wget
RUN apt update && apt install -y build-essential wget
# Install Go
RUN wget https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz
RUN rm -f go${GOLANG_VERSION}.linux-amd64.tar.gz

ENV PATH "$PATH:/usr/local/go/bin"

ARG PROJECT_BINARY=bestir-identity-service
ARG PROJECT_BUILD_DIR=./build/bin

WORKDIR /app

RUN ls

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /bestir-identity-service

EXPOSE 8080

 CMD [ "/bestir-identity-service" ]
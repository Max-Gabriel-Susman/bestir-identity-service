# syntax=docker/dockerfile:1

# set the base image
FROM golang:alpine as builder

# turn on go modules
ENV GO111MODULE=on

# git is required for dependcies
RUN apk update && apk add --no-cache git

# set curren working directory inside the container
WORKDIR /app

# copy go mod and sum files 
COPY go.mod go.sum ./

# download all dependencies
RUN go mod download

# copy source from the current working directory to the containers working directory
COPY . . 

# build the go app 
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN CGO_ENABLED=0 GOOS=linux go build .

EXPOSE 80

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .       

# Expose port 8080 to the outside world
EXPOSE 8080

# CMD [ "/bestir-identity-service" ]
#Command to run the executable
CMD ["./main"]
 
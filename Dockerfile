# syntax=docker/dockerfile:1

FROM golang:latest
RUN apt-get install git
EXPOSE 8080
WORKDIR /go/src/app
COPY /api .
RUN go mod download -x
RUN ["go", "install", "github.com/githubnemo/CompileDaemon@latest"]
ENTRYPOINT CompileDaemon -build="go build main.go" -command="/go/src/app/main" -directory=. -polling=true
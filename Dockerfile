# syntax=docker/dockerfile:1

FROM golang as builder

RUN mkdir -p /api
WORKDIR /api

COPY api/go.mod .
COPY api/go.sum .
RUN go mod download

COPY /api .

RUN go build -o /docker-gs-ping

EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o /bin/server

CMD ["/bin/server", "/api"]
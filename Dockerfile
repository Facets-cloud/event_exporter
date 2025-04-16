# Dockerfile
FROM golang:1.24 AS build

WORKDIR /src

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o event_exporter ./cmd/main.go

FROM debian:stable-slim

USER nobody

COPY --from=build /src/event_exporter /event_exporter

ENTRYPOINT ["/event_exporter"]

EXPOSE 9102

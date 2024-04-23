# Dockerfile
FROM golang:1.17 AS build

ARG TARGETARCH
WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -o /event_exporter .

FROM debian:stretch-slim

USER nobody

COPY --from=build /event_exporter /event_exporter

ENTRYPOINT ["/event_exporter"]

EXPOSE 9102
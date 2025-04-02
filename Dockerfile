# Dockerfile
FROM golang:1.17 AS build

ARG TARGETARCH
WORKDIR /src

COPY . .

COPY ./bin/event_exporter_${TARGETARCH} /event_exporter

FROM debian:stable-slim

USER nobody

COPY --from=build /event_exporter /event_exporter

ENTRYPOINT ["/event_exporter"]

EXPOSE 9102

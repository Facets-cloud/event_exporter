FROM debian:stretch-slim

COPY bin/event_exporter /

ENTRYPOINT ["/event_exporter"]

EXPOSE 9102

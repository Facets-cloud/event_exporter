FROM debian:stretch-slim

USER nobody

COPY bin/event_exporter /

ENTRYPOINT ["/event_exporter"]

EXPOSE 9102

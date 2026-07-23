FROM debian:bookworm-slim

ENV DEBIAN_FRONTEND=noninteractive
ENV WINEDEBUG=-all

RUN dpkg --add-architecture i386 \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
        ca-certificates \
        wine \
        wine64 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /work

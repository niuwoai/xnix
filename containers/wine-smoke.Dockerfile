ARG XNIX_WINE_BASE_PLATFORM=linux/amd64
FROM --platform=${XNIX_WINE_BASE_PLATFORM} debian:bookworm-slim

ENV DEBIAN_FRONTEND=noninteractive
ENV WINEDEBUG=-all

RUN if [ "$(dpkg --print-architecture)" = "amd64" ]; then dpkg --add-architecture i386; fi \
    && apt-get update \
    && if [ "$(dpkg --print-architecture)" = "amd64" ]; then wine_packages="wine wine32 wine64"; else wine_packages="wine wine64"; fi \
    && apt-get install -y --no-install-recommends \
        ca-certificates \
        procps \
        ${wine_packages} \
        x11-utils \
        xvfb \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /work

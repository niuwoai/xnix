FROM debian:bookworm-slim

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
    && apt-get install --yes --no-install-recommends \
        bc \
        bison \
        build-essential \
        ca-certificates \
        cpio \
        file \
        flex \
        git \
        libelf-dev \
        libssl-dev \
        qemu-system-x86 \
        qemu-utils \
        ruby \
        rsync \
        unzip \
        wget \
        xz-utils \
    && rm -rf /var/lib/apt/lists/*

RUN useradd --create-home --shell /bin/bash xnix

USER xnix
WORKDIR /workspace

ENV BR2_DL_DIR=/workspace/.cache/buildroot-dl

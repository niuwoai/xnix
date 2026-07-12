FROM debian:bookworm-slim

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
    && apt-get install --yes --no-install-recommends \
        bc \
        bison \
        build-essential \
        ca-certificates \
        cpio \
        dbus \
        file \
        flex \
        git \
        libglib2.0-bin \
        libelf-dev \
        libssl-dev \
        openssh-client \
        qemu-system-x86 \
        qemu-utils \
        ruby \
        rsync \
        unzip \
        wget \
        xz-utils \
    && rm -rf /var/lib/apt/lists/*

RUN useradd --create-home --shell /bin/bash xnix

COPY --chown=xnix:xnix . /workspace
RUN mkdir --parents /workspace/.cache/buildroot \
    && chown --recursive xnix:xnix /workspace/.cache

USER xnix
WORKDIR /workspace
ENV BR2_DL_DIR=/workspace/.cache/buildroot-dl

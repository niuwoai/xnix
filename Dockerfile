ARG XNIX_TOOLS_BASE_IMAGE=debian:bookworm-slim
FROM ${XNIX_TOOLS_BASE_IMAGE} AS dbus-tools

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
    && apt-get install --yes --no-install-recommends \
        build-essential \
        ca-certificates \
        dbus \
        libglib2.0-bin \
        libglib2.0-dev \
        pkg-config \
        ruby \
    && rm -rf /var/lib/apt/lists/*

RUN useradd --create-home --shell /bin/bash xnix

WORKDIR /workspace
USER xnix

FROM dbus-tools AS tools

USER root
ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
    && apt-get install --yes --no-install-recommends \
        bc \
        bison \
        clang \
        cpio \
        file \
        flex \
        git \
        golang-go \
        libelf-dev \
        libssl-dev \
        lld \
        llvm \
        openssh-client \
        qemu-system-x86 \
        qemu-utils \
        rsync \
        unzip \
        wget \
        xz-utils \
    && rm -rf /var/lib/apt/lists/*

COPY --chown=xnix:xnix . /workspace
RUN mkdir --parents /workspace/.cache/buildroot \
    && chown --recursive xnix:xnix /workspace/.cache
WORKDIR /workspace
USER xnix
WORKDIR /workspace
ENV BR2_DL_DIR=/workspace/.cache/buildroot-dl

FROM tools AS tested-runtime

USER root
WORKDIR /workspace
RUN go test -timeout 90m ./...
RUN go build -o /usr/local/bin/xnix-runtime-go ./cmd/xnix-runtime-go
RUN go build -o /usr/local/bin/xnix-runtime-owner ./cmd/xnix-runtime-owner
RUN go build -o /usr/local/bin/xnix-compat-launch ./cmd/xnix-compat-launch
RUN gcc /workspace/runtime/dbus/xnix_compatd_smoke.c \
        -o /usr/local/bin/xnix-dbus-smoke \
        $(pkg-config --cflags --libs gio-2.0)
RUN gcc -std=c11 -Wall -Wextra -Werror \
        /workspace/runtime/core/xnix_runtime_core.c \
        /workspace/runtime/core/xnix_runtime_core_cli.c \
        -o /usr/local/bin/xnix-runtime-core

USER xnix
WORKDIR /workspace

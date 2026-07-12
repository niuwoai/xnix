# Xnix Development Guide

This file is for Claude Code and other automated contributors. Follow [AGENTS.md](AGENTS.md) for the shared project rules.

## Current Phase

- Version: `0.1.0`
- Target architecture: x86_64
- Virtual machine: QEMU running with software emulation inside a restricted Docker container
- System stack: Linux LTS kernel, Buildroot, BusyBox, and initramfs
- Service target: OpenSSH `sshd`
- Container host: Colima Docker, limited to 1 GiB RAM and 1 CPU

## Safety Constraints

- Do not use `--privileged`, `--network host`, a Docker socket mount, or writable mounts outside the repository.
- Bind test ports only to `127.0.0.1`.
- Do not install QEMU or build dependencies on the macOS host; keep them in the build container.

## Planned Commands

```text
yarn system:build
yarn qemu:run
yarn ssh:verify
```

These commands are planned, not implemented. Update this file and `README.md` when an executable command is added.

## Delivery Sequence

1. Reproducible build environment and source layout.
2. Minimal bootable image and serial logging.
3. BusyBox shell and user-mode networking.
4. Init-managed OpenSSH `sshd` with loopback-only port forwarding.
5. Logging, users, key authentication, and hardening.

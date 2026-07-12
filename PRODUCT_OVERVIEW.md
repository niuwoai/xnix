# Xnix Product Overview

> Last updated: 2026-07-12 | Current version: v0.1.1

## Summary

Xnix is a Linux-compatible system for learning operating-system principles. It reuses the Linux kernel for reliable hardware, network, and user-space ABI support while keeping the surrounding system small and observable. The first runnable target is a QEMU x86_64 image that starts and manages OpenSSH `sshd`.

## Core Goals

- Produce an x86_64 image that boots in QEMU with persistent serial logs.
- Provide Linux user-space compatibility through an LTS Linux kernel.
- Start a small root filesystem with a shell, networking, and a remote-maintenance service.
- Start, stop, restart, and inspect OpenSSH `sshd` through the init system.
- Run builds and tests inside constrained Colima Docker infrastructure.

## Non-Goals for This Phase

- Reimplementing the Linux kernel, its syscall ABI, or driver ecosystem.
- Supporting physical hardware, graphical desktops, public network exposure, or package repositories.
- Adding localization or non-English project-facing content.

## Technology Choices

| Layer | Initial choice | Purpose |
| --- | --- | --- |
| Kernel | Linux LTS | Stable hardware, networking, and Linux ABI support |
| System builder | Buildroot | Reproducible toolchain, kernel, root filesystem, and image generation |
| User space | BusyBox with initramfs | Small, understandable boot path |
| Service | OpenSSH `sshd` | Validate networking and service management |
| Virtualization | QEMU x86_64 with TCG | Safe, portable software-emulated test environment |
| Container host | Colima Docker | Constrained build and test isolation |

## Buildroot Baseline

- Buildroot version: `2025.02.15` LTS.
- Buildroot source is intentionally not committed. A later build command will retrieve the pinned release into an ignored cache directory.
- The external tree lives in `buildroot/` and provides the `xnix_x86_64_defconfig` configuration.

## Test and Release Policy

- Each small code version receives focused unit tests and a Git commit.
- Every tenth code version requires a complete build and QEMU smoke test.
- If the full milestone test exposes a defect, it must be fixed and the full test rerun before later work starts.
- Colima is limited to 1 GiB memory and 1 CPU. Runtime containers receive explicit resource limits and no elevated privileges.

## Milestones

- [x] M0: Reproducible source layout, constrained container definition, and test policy.
- [ ] M1: Linux kernel boot in QEMU with serial output.
- [ ] M2: BusyBox shell and initramfs.
- [ ] M3: User-mode virtual NIC and network reachability.
- [ ] M4: Init-managed OpenSSH `sshd` and loopback-only host SSH verification.
- [ ] M5: Service logs, user management, key authentication, and baseline hardening.

## Current Decisions

- “ssdh” is interpreted as OpenSSH `sshd` unless clarified otherwise.
- BusyBox init is used before evaluating systemd.
- initramfs is used before introducing a persistent disk image.

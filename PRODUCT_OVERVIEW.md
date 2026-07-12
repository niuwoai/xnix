# Xnix Product Overview

> Last updated: 2026-07-12 | Current version: v0.2.6

## Summary

Xnix is an atomic Linux desktop designed to make existing Windows applications feel native. KDE Plasma provides the familiar desktop shell; the independent Xnix Compatibility Runtime owns backend selection, recipes, diagnostics, permissions, snapshots, and rollback. The existing Buildroot/QEMU image remains a learning and low-level verification baseline, not the flagship product base.

## Core Goals

- Ship one official flagship desktop: KDE Plasma 6.
- Make Windows applications appear as ordinary Linux applications in the launcher, task manager, file associations, notifications, and settings.
- Keep Wine/Proton and Windows VM management behind an independent D-Bus Runtime.
- Use application recipes for backend selection, dependencies, diagnostics, permissions, snapshots, and rollback.
- Use XDG Desktop Portal for user-controlled file, URI, print, clipboard, screen, camera, and remote-desktop access.
- Preserve reproducible low-level QEMU verification and constrained Colima development infrastructure.

## Non-Goals for This Phase

- Forking KDE Plasma, KWin, Wine, Proton, or a Linux distribution.
- Supporting GNOME or XFCE as an officially integrated desktop in the first product release.
- Exposing Wine prefixes, raw executable paths, or backend implementation terminology in normal user entry points.
- Adding localization or non-English project-facing content during the current phase.

## Technology Choices

| Layer | Initial choice | Purpose |
| --- | --- | --- |
| Atomic base | Fedora Kinoite-compatible image | KDE Plasma desktop and transactional operating-system rollback |
| Desktop shell | KDE Plasma 6 | Familiar launcher, task manager, tray, desktop icons, file manager, and multi-window behavior |
| Runtime | Xnix Compatibility Runtime | Independent D-Bus service for recipes, Wine/VM backends, permissions, diagnostics, snapshots, and rollback |
| Compatibility backends | Wine/Proton and Windows VM | Selectable implementation behind a stable Runtime API |
| Permissions | XDG Desktop Portal | User-mediated file, URI, print, clipboard, screen, and remote-desktop access |
| Desktop integration | Plasmoid, KRunner, KWin, Dolphin, and system services | Presentation and interaction only; no backend policy in KDE code |
| Learning baseline | Buildroot plus QEMU | Reproducible boot, initramfs, networking, and SSH verification |
| Container host | Colima Docker | Constrained developer build and test isolation |

## Compatibility Runtime Boundary

- The Runtime contract reserves `org.xnix.Compatibility1` for application discovery, recipe installation, launch, diagnostics, snapshot, and restore operations. The local daemon core is executable and can probe itself, list managed recipes, report diagnostics, dispatch read-only Runtime methods, stage activation files under an unprivileged root, and emit a KDE-safe read model for the Compatibility Center. The D-Bus session smoke adapter owns the Runtime bus name and answers read-only calls in the constrained Linux container; the production daemon binding is the next implementation step.
- KDE packages consume Runtime state and submit user actions over D-Bus. They do not invoke Wine or a VM directly.
- Application recipes generate normal `.desktop` launchers that call `xnix-compat-launch --app <id>`. They never display a prefix path or a raw Windows executable command.
- Runtime operations that need desktop resources must use XDG Desktop Portal request/response flows; backend-specific direct access is not a public UI contract.

## Desktop Integration Scope

| Entry point | First implementation responsibility |
| --- | --- |
| Launcher | Generated standard desktop files unify Windows and Linux applications |
| Task manager | KWin rules map Wine and VM windows to Runtime application identities |
| File manager | Dolphin action resolves a Runtime application recipe and asks for Portal-granted files |
| System tray | Plasma applet presents Runtime activity and compatible tray applications |
| Notifications | Runtime requests KDE notifications for install, repair, backend, and approval events |
| Compatibility Center | Plasma package displays the Runtime-backed read model for compatibility, diagnostics, snapshots, and actions |
| Settings | KDE settings module presents user concepts such as automatic mode, performance, compatibility, and allowed resources |

## Learning Baseline

- Buildroot version: `2025.02.15` LTS.
- Buildroot source is intentionally not committed. `buildroot/sources.lock` pins its upstream URL and SHA-256; a build command retrieves it into an ignored cache directory.
- The external tree lives in `buildroot/` and provides the `xnix_x86_64_defconfig` configuration.

## Verified Runtime Baseline

- A complete constrained Buildroot build exits successfully and produces the x86_64 kernel and initramfs.
- The QEMU serial console shows Linux booting, running `/init`, and reaching the `xnix login:` prompt.
- The QEMU user-mode NIC obtains `10.0.2.15` through DHCP.
- OpenSSH `sshd` accepts a non-interactive key-authenticated connection through the container-only loopback forwarding rule.

## Test and Release Policy

- Each small code version receives focused unit tests and a Git commit.
- Every tenth code version requires a complete build and QEMU smoke test.
- If the full milestone test exposes a defect, it must be fixed and the full test rerun before later work starts.
- Colima is limited to 6 GiB memory and 1 CPU. Runtime containers receive explicit resource limits and no elevated privileges.
- Buildroot is configured with a single job to prevent toolchain build memory spikes in the 4 GiB container limit.
- Xnix runtime containers are limited to 4 GiB, one CPU, and 256 processes. They use a read-only root filesystem, a small temporary filesystem, no Linux capabilities, no host-directory mounts, and no network unless a dedicated source-retrieval or Buildroot dependency-download command requires it. These commands receive bridge networking and a Docker-managed named volume only; they never mount a host directory. Image builds run inside the 6 GiB and one-CPU Colima VM limit.
- SSH verification creates an ephemeral Ed25519 keypair only in the Docker-managed cache volume, installs its public key during image assembly, forwards only container loopback `127.0.0.1:2222` to guest port 22, and deletes the private key after the test.

## Milestones

- [x] B0: Reproducible constrained Buildroot/QEMU learning baseline.
- [x] B1: Kernel boot, initramfs, DHCP, and loopback `sshd` verification.
- [x] C0: Compatibility Runtime D-Bus contract, safe recipe model, desktop launcher generator, and Plasma package skeleton.
- [ ] C1: Activated Runtime implementation and signed recipe storage. The local daemon core, managed recipe-store loader, read-only method dispatcher, packaged activation wrapper, activation-file installer, Linux D-Bus session smoke adapter, and KDE-safe Compatibility Center model are in place; production daemon binding and signature validation are still pending.
- [ ] C2: Wine/Proton backend with snapshots, diagnostics, and Portal-mediated permissions.
- [ ] C3: Windows VM backend with file, clipboard, print, and window bridging.
- [ ] C4: KDE launcher, task manager, Dolphin, tray, notification, KRunner, KWin, Compatibility Center, and settings integrations.
- [ ] C5: Reproducible atomic KDE desktop image and graphical QEMU smoke test.

## Current Decisions

- KDE Plasma is the only official desktop for the first product release; GNOME and XFCE remain future optional shells.
- The flagship base is Fedora Kinoite-compatible because an atomic deployment model matches system and application rollback requirements.
- Buildroot is retained for kernel, boot, initramfs, and QEMU learning experiments, not for packaging the flagship desktop.
- The Runtime is the product core; desktop packages remain replaceable adapters.

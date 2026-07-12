# Xnix Product Overview

> Last updated: 2026-07-13 | Current version: v0.2.22

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

- The Runtime contract reserves `org.xnix.Compatibility1` for application discovery, recipe installation, launch, diagnostics, snapshot, and restore operations. The local daemon core is executable and can probe itself, list managed recipes, report diagnostics, dispatch read-only Runtime methods, stage activation files under an unprivileged root, and emit a KDE-safe read model for the Compatibility Center. The D-Bus session smoke adapter owns the Runtime bus name and answers read-only calls in the constrained Linux container; the KDE read model can now consume that session-bus source. The production daemon binding is the next implementation step.
- Recipe registry verification delegates to `xnix-recipe-registry`, which validates registry schema version, recipe identifiers, relative recipe paths, SHA-256 digests, and signature status without claiming production signed-recipe validation yet. The Runtime daemon uses registry-backed recipe loading when `registry.json` is present, so the default application list is digest-verified before exposure. The Runtime probe reports recipe trust status for KDE surfaces and diagnostics, while `xnix-recipe-trust-policy` evaluates whether the registry is production-trusted, development-only, or untrusted. Recipe install gates delegate to `xnix-recipe-install-gate`, which blocks production activation of development-only registries while still allowing digest-verified development staging.
- KDE packages consume Runtime state and submit user actions over D-Bus. They do not invoke Wine or a VM directly.
- Application recipes generate normal `.desktop` launchers that call `xnix-compat-launch --app <id>`. That entry point validates the Runtime application id, preserves optional file URIs, and models a `Launch` request without displaying a prefix path, a backend command, or a raw Windows executable command.
- Dolphin service menus delegate selected files to `xnix-compat-open`, which validates file URIs, resolves a Runtime application recipe by file extension, and models a portal-required `Launch` request without directly invoking Wine or a VM.
- Runtime events delegate to `xnix-compat-notify`, which models KDE notification payloads for install failures, automatic repairs, mode changes, and approval requests without exposing backend implementation terms.
- Compatibility settings delegate to `xnix-compat-settings`, which models user-facing controls for automatic mode, performance or compatibility priority, documents and downloads access, camera access, network access, and snapshots without backend terminology.
- System tray status delegates to `xnix-compat-tray-status`, which models Runtime activity, attention state, and bridged tray application counts for a KDE tray surface.
- Task manager identity delegates to `xnix-compat-window-identity`, which models desktop file mapping, grouping, pinning, restore behavior, and KWin identity-only metadata for compatibility windows.
- Desktop integration manifests delegate to `xnix-desktop-integration-manifest`, which bundles a recipe's launcher, task manager, Dolphin, tray, notification, Compatibility Center, and settings artifacts into one activation plan without exposing backend commands.
- Desktop activation delegates to `xnix-install-desktop-integration`, which stages generated application launchers, Dolphin service menus, and the desktop integration manifest under a target root without modifying the host root.
- Desktop activation rollback delegates to `xnix-rollback-desktop-integration`, which reads activation receipts and removes only unchanged staged files after SHA-256 verification.
- Runtime operations that need desktop resources must use XDG Desktop Portal request/response flows; backend-specific direct access is not a public UI contract.
- `xnix-kde-integration-status` tracks the seven first-release KDE entry points and records which have an initial Runtime-backed integration versus planned work.

## Desktop Integration Scope

| Entry point | First implementation responsibility |
| --- | --- |
| Launcher | Generated standard desktop files call the managed Runtime launch request path |
| Task manager | Window identity model maps compatibility windows to Runtime application identities for grouping, pinning, and restore |
| File manager | Dolphin action delegates file URIs to the Runtime file-open request path and requires Portal-mediated access |
| System tray | Tray status model presents Runtime activity, attention state, and bridged tray applications |
| Notifications | Runtime events produce KDE notification request models for install, repair, mode-change, and approval events |
| Compatibility Center | Plasma package displays the Runtime-backed read model for compatibility, diagnostics, snapshots, and actions; the read path is smoke-tested over D-Bus |
| Settings | KDE settings model presents user concepts such as automatic mode, performance, compatibility, allowed resources, devices, network, and snapshots |
| Activation manifest | Runtime manifest groups all seven KDE artifacts for a recipe activation without exposing backend commands |
| Activation staging | Runtime installer writes desktop entry, Dolphin service menu, and manifest files under a staging root |
| Activation rollback | Runtime rollback removes only receipt-tracked, checksum-matching staged files |
| Recipe registry | Registry verifier checks recipe metadata, paths, digests, and development signature status |
| Registry-backed loading | Runtime loads registered recipes after digest verification and keeps no-registry fallback only for development fixtures |
| Recipe trust probe | Runtime probe reports registry-backed loading, digest verification, and signed-recipe validation status |
| Recipe trust policy | Policy model turns registry trust signals into production, development-only, or untrusted decisions |
| Recipe install gate | Install gate blocks production activation until recipe trust requirements are satisfied |

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
- [ ] C1: Activated Runtime implementation and signed recipe storage. The local daemon core, registry-backed recipe-store loader, recipe registry verifier, recipe trust probe, recipe trust policy, recipe install gate, read-only method dispatcher, packaged activation wrapper, activation-file installer, desktop activation installer, activation rollback receipt, Linux D-Bus session smoke adapter, KDE-safe Compatibility Center model, KDE D-Bus read smoke, launcher request model, Dolphin file-open request model, notification request model, settings model, tray status model, task manager identity model, desktop integration manifest, and seven-entry-point KDE integration status are in place; production daemon binding and production signature validation are still pending.
- [ ] C2: Wine/Proton backend with snapshots, diagnostics, and Portal-mediated permissions.
- [ ] C3: Windows VM backend with file, clipboard, print, and window bridging.
- [ ] C4: KDE launcher, task manager, Dolphin, tray, notification, KRunner, KWin, Compatibility Center, and settings integrations.
- [ ] C5: Reproducible atomic KDE desktop image and graphical QEMU smoke test.

## Current Decisions

- KDE Plasma is the only official desktop for the first product release; GNOME and XFCE remain future optional shells.
- The flagship base is Fedora Kinoite-compatible because an atomic deployment model matches system and application rollback requirements.
- Buildroot is retained for kernel, boot, initramfs, and QEMU learning experiments, not for packaging the flagship desktop.
- The Runtime is the product core; desktop packages remain replaceable adapters.

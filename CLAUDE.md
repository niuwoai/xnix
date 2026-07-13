# Xnix Development Guide

This file is for Claude Code and other automated contributors. Follow [AGENTS.md](AGENTS.md) for the shared project rules.

## Current Phase

- Version: `0.2.35`
- Target architecture: x86_64
- Virtual machine: QEMU running with software emulation inside a restricted Docker container
- Flagship stack: Fedora Kinoite-compatible atomic base, KDE Plasma 6, XDG Desktop Portal, and the Xnix Compatibility Runtime
- Learning baseline: Linux LTS kernel, Buildroot, BusyBox, initramfs, and OpenSSH `sshd`
- Container host: Colima Docker, limited to 6 GiB RAM and 1 CPU

## Safety Constraints

- Do not use `--privileged`, `--network host`, a Docker socket mount, or host-directory mounts at runtime.
- Bind test ports only to `127.0.0.1`.
- Do not install QEMU or build dependencies on the macOS host; keep them in the build container.

## Commands

```text
ruby scripts/verify_layout.rb
ruby scripts/fetch_buildroot.rb --verify-lock
ruby -Ilib test/test_container.rb
ruby scripts/container.rb fetch-sources
ruby -Ilib test/test_buildroot.rb
ruby -Ilib test/test_qemu.rb
ruby -Ilib test/test_serial_log.rb
ruby -Ilib test/test_milestone.rb
ruby scripts/full_smoke.rb
ruby -Ilib test/test_ssh_probe.rb
ruby scripts/container.rb prepare-ssh-test-key
ruby scripts/container.rb build-ssh-test-system
ruby scripts/container.rb ssh-smoke
ruby -Ilib test/test_compatibility_engine_catalog.rb
ruby -Ilib test/test_application_recipe.rb
ruby -Ilib test/test_compatibility_repair_plan.rb
ruby -Ilib test/test_compatibility_snapshot_plan.rb
ruby -Ilib test/test_compatibility_run_plan.rb
ruby -Ilib test/test_compatibility_test_plan.rb
ruby -Ilib test/test_recipe_store.rb
ruby -Ilib test/test_recipe_registry.rb
ruby -Ilib test/test_recipe_trust_policy.rb
ruby -Ilib test/test_recipe_install_gate.rb
ruby -Ilib test/test_registry_backed_recipe_store.rb
ruby -Ilib test/test_desktop_entry.rb
ruby -Ilib test/test_dbus_runtime_client.rb
ruby -Ilib test/test_dolphin_service_menu.rb
ruby -Ilib test/test_file_open_request.rb
ruby -Ilib test/test_launch_request.rb
ruby -Ilib test/test_notification_request.rb
ruby -Ilib test/test_portal_access_policy.rb
ruby -Ilib test/test_settings_model.rb
ruby -Ilib test/test_tray_status_model.rb
ruby -Ilib test/test_task_manager_identity.rb
ruby -Ilib test/test_desktop_activation_installer.rb
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby -Ilib test/test_desktop_integration_manifest.rb
ruby -Ilib test/test_runtime_contract.rb
ruby -Ilib test/test_runtime_daemon.rb
ruby -Ilib test/test_runtime_dispatch.rb
ruby -Ilib test/test_runtime_activation.rb
ruby -Ilib test/test_runtime_activation_install.rb
ruby -Ilib test/test_runtime_dbus_smoke_script.rb
ruby -Ilib test/test_kde_center_model.rb
ruby -Ilib test/test_kde_integration_status.rb
ruby scripts/container.rb runtime-dbus-smoke
ruby scripts/container.rb kde-center-dbus-smoke
```

Buildroot commands are restricted to the managed Docker volume. The SSH smoke command creates a disposable test key there, forwards only container loopback to the guest, and removes the private key afterward. The Compatibility Runtime owns Wine/VM policy and state; KDE packages must use its D-Bus contract rather than embedding compatibility logic. The current daemon core can probe itself, list registry-verified recipes, catalog compatibility engines, plan compatibility runs, plan compatibility repairs, plan compatibility snapshots, plan compatibility tests, report recipe trust status, evaluate recipe trust policy, evaluate recipe install gates, evaluate Portal access policy, model Portal requests, report diagnostics with repair, snapshot, and test-plan summaries, dispatch read-only Runtime planning methods, verify recipe registry digests and development signature status, stage activation files under an unprivileged root after install-gate preflight, provide a KDE-safe read model for the Compatibility Center, consume read-only planning methods through the D-Bus client, verify that model over a session-bus Runtime smoke adapter, model KRunner query results, model KWin window rules, model file associations, model Dolphin file-open requests, model managed launcher requests with run-plan summaries, model Runtime notification requests, model user-facing compatibility settings, model system tray status, model task manager window identity, build a desktop integration manifest for the seven KDE entry points, stage desktop activation files from that manifest under a target root, write rollback receipts, rollback unchanged staged files with checksum verification, and report KDE seven-entry-point integration status. The production daemon binding is still pending.

## Delivery Sequence

1. Atomic KDE Plasma desktop image and standard desktop services.
2. Independent Compatibility Runtime with D-Bus service activation and signed recipe registry.
3. Wine/Proton backend, snapshots, diagnostics, and managed desktop entry generation.
4. VM backend with file, clipboard, print, and application-window bridging.
5. KDE Compatibility Center, KRunner, KWin, Dolphin, tray, notification, and settings integration.

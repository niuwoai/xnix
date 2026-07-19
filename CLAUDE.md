# Xnix Development Guide

This file is for Claude Code and other automated contributors. Follow [AGENTS.md](AGENTS.md) for the shared project rules.

## Current Phase

- Version: `0.2.373`
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
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/merge_readiness_packet.rb --format json
ruby scripts/fetch_buildroot.rb --verify-lock
ruby -Ilib test/test_action_review_receipt.rb
ruby -Ilib test/test_implementation_evidence_report.rb
ruby -Ilib test/test_merge_readiness_packet.rb
ruby -Ilib test/test_container.rb
ruby -Ilib test/test_compatibility_artifact_manifest.rb
ruby -Ilib test/test_compatibility_install_plan.rb
ruby -Ilib test/test_compatibility_acquisition_preflight.rb
ruby -Ilib test/test_compatibility_action_queue.rb
ruby -Ilib test/test_compatibility_package_source.rb
ruby -Ilib test/test_application_state_root.rb
ruby -Ilib test/test_compatibility_backend_binding.rb
ruby -Ilib test/test_ai_diagnostic_input.rb
ruby -Ilib test/test_ai_diagnostic_recommendation.rb
ruby -Ilib test/test_ai_repair_approval_gate.rb
ruby -Ilib test/test_runtime_live_owner_gate.rb
ruby -Ilib test/test_runtime_method_parity_manifest.rb
ruby -Ilib test/test_runtime_owner_smoke_plan.rb
ruby -Ilib test/test_runtime_service_binding.rb
ruby -Ilib test/test_runtime_write_gate.rb
ruby -Ilib test/test_runtime_core.rb
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
ruby -Ilib test/test_compatibility_test_result.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
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
ruby -Ilib test/test_settings_change_plan.rb
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
ruby -Ilib test/test_runtime_activation_smoke_script.rb
ruby -Ilib test/test_runtime_dbus_smoke_script.rb
ruby -Ilib test/test_kde_center_model.rb
ruby -Ilib test/test_kde_integration_status.rb
ruby scripts/container.rb runtime-dbus-smoke
ruby scripts/container.rb kde-center-dbus-smoke
```

Buildroot commands are restricted to the managed Docker volume. The SSH smoke command creates a disposable test key there, forwards only container loopback to the guest, and removes the private key afterward. The Compatibility Runtime owns Wine/VM policy and state; KDE packages must use its D-Bus contract rather than embedding compatibility logic. Important Runtime product logic should move into Go first, with C kept for low-level or already-owned Runtime policy surfaces, while Ruby remains focused on tests, scaffolding, and development tools. The current daemon core can probe itself, list registry-verified recipes, catalog compatibility engines, expose a registry-backed Go Runtime desktop identity planner, activation-staging renderer, activation-transaction renderer, and activation-status renderer, expose a C Runtime core for stable identity, safe application catalog metadata, compatibility engine selection metadata, Portal access policy metadata, Portal request plan metadata, desktop entry plan metadata, task manager identity plan metadata, KWin window rule metadata, file association plan metadata, notification plan metadata, tray status plan metadata, KRunner query plan metadata, AI diagnostic input metadata, AI diagnostic recommendation metadata, AI repair approval gate metadata, Compatibility Center summary metadata, snapshot policy metadata, application state-root policy metadata, install readiness metadata, compatibility install plan metadata, Compatibility Center action queue metadata, Compatibility Center action review receipt metadata, compatibility run plan metadata, compatibility repair plan metadata, compatibility test plan metadata, compatibility test result metadata, execution readiness metadata, launch intent metadata, desktop activation manifest metadata, artifact manifest metadata, acquisition preflight metadata, package source metadata, backend binding metadata, settings metadata, settings change metadata, service binding metadata, live owner gate metadata, owner smoke plan metadata, method parity manifest metadata, recipe trust policy metadata, recipe install gate metadata, and write-gate decisions, model managed desktop activation manifests, model managed artifact manifests, model managed install plans, model managed acquisition preflight, model managed package source selection, model managed application state roots, model managed backend binding status, plan Compatibility Center summaries, plan C-owned compatibility repairs, plan compatibility snapshots, plan C-owned compatibility tests, summarize compatibility test results, model C-owned execution readiness, model C-owned launch intent, build AI diagnostic inputs, build AI diagnostic recommendations, gate AI repair approval, model Runtime service binding status, model Runtime live owner gates, model Runtime owner smoke plans, model Runtime method parity manifests, model Runtime write gates, report recipe trust status, evaluate recipe trust policy, evaluate recipe install gates, evaluate Portal access policy, model C-owned Portal request plans, model C-owned desktop entry plans, model C-owned task manager identity plans, model C-owned KWin window rule plans, model C-owned file association plans, model C-owned notification plans, model C-owned tray status plans, model C-owned KRunner query plans, model C-owned compatibility repair plans, model C-owned compatibility snapshot plans, model C-owned compatibility test plans, model C-owned compatibility test results, model C-owned execution readiness, model C-owned AI diagnostic inputs, model C-owned AI diagnostic recommendations, model C-owned AI repair approval gates, model C-owned Compatibility Center summaries, model Portal requests, report diagnostics with action-queue, action-review-receipt, compatibility-center-summary, desktop-activation-manifest, desktop-entry-plan, task-manager-identity-plan, kwin-window-rule-plan, file-association-plan, notification-plan, tray-status-plan, krunner-query-plan, portal-request-plan, artifact-manifest, install-plan, acquisition-preflight, package-source, state-root, settings, settings-change-plan, repair, snapshot, test-plan, test-result, execution-readiness, launch-intent, backend binding, AI diagnostic input, AI diagnostic recommendation, AI repair approval gate, Runtime live owner gate, Runtime owner smoke plan, Runtime method parity manifest, Runtime write gate, and Runtime service binding summaries, dispatch read-only Runtime planning methods, reject reserved write methods with explicit Runtime gate errors, verify recipe registry digests and development signature status, install Runtime libraries and assets beside the packaged activation wrapper, stage activation files under an unprivileged root after install-gate preflight, smoke the staged packaged wrapper, provide a KDE-safe read model for the Compatibility Center, consume read-only planning methods through the D-Bus client, verify that model over a split session-bus Runtime smoke adapter with introspection XML in `runtime/dbus/xnix_compatd_introspection.inc`, model KRunner query results, model KWin window rules, model file associations, model Dolphin file-open requests, model managed launcher requests and C-owned launch intents with run-plan summaries, model Runtime notification requests, model Runtime-backed user-facing compatibility settings and settings change plans, model system tray status, model task manager identity, build a desktop integration manifest for the seven KDE entry points, stage desktop activation files from that manifest under a target root, write rollback receipts, rollback unchanged staged files with checksum verification, and report KDE seven-entry-point integration status. Live production D-Bus ownership and enabled compatibility backend launch remain pending.

## Delivery Sequence

1. Atomic KDE Plasma desktop image and standard desktop services.
2. Independent Compatibility Runtime with D-Bus service activation and signed recipe registry.
3. Wine/Proton backend, snapshots, diagnostics, and managed desktop entry generation.
4. VM backend with file, clipboard, print, and application-window bridging.
5. KDE Compatibility Center, KRunner, KWin, Dolphin, tray, notification, and settings integration.

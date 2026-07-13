# Xnix

Xnix is an atomic KDE Plasma desktop project focused on making existing Windows applications feel native on Linux.

The project is currently at `v0.2.53`. KDE Plasma is the only supported flagship desktop. An independent Compatibility Runtime manages application recipes, Wine/VM backends, Compatibility Center action queues and review receipts, managed artifact manifests, managed install plans, managed acquisition preflight, managed package source selection, managed application state roots, managed backend binding status, Runtime-backed compatibility settings and settings change plans, diagnostics, AI diagnostic inputs and recommendations, AI repair approval gates, Runtime service binding status, Runtime live owner gates, Runtime owner smoke plans, Runtime method parity manifests, compatibility tests, snapshots, and rollback without exposing implementation details in normal desktop entry points.

See [PRODUCT_OVERVIEW.md](PRODUCT_OVERVIEW.md) for the roadmap and technical choices. See [AGENTS.md](AGENTS.md) for contribution and safety rules.

## Current Verification

The first end-to-end Buildroot run has completed successfully: Xnix boots its Linux kernel and initramfs in constrained QEMU, acquires `10.0.2.15` through user-mode DHCP, and accepts a loopback-only key-authenticated SSH probe.

The v0.2.53 Runtime foundation defines application recipes, compatibility engine cataloging, Compatibility Center action queue and review receipt modeling, managed artifact manifest modeling, managed install plan modeling, managed acquisition preflight, managed package source selection, managed application state root modeling, managed backend binding status modeling, Runtime-backed compatibility settings, Runtime-owned settings change planning, compatibility run planning, compatibility repair planning, compatibility snapshot planning, compatibility test planning, compatibility test result summarization, AI diagnostic input modeling, AI diagnostic recommendation modeling, AI repair approval gate modeling, Runtime service binding status modeling, Runtime live owner gate modeling, Runtime owner smoke planning, Runtime method parity manifests, recipe registry verification, registry-backed Runtime recipe loading, recipe trust probe reporting, recipe trust policy evaluation, recipe install gate evaluation, preflight-gated desktop activation, Portal access policy evaluation, Portal request modeling, standard desktop entries, file association generation, managed launcher, Dolphin file-open, notification, settings, tray status, task-manager identity, KWin window rule, desktop integration manifest models, a staged desktop activation installer, and activation receipt rollback, a D-Bus service contract with read-only planning methods, a hardened systemd unit, a Plasma 6 Compatibility Center package skeleton, a Dolphin service menu, a KRunner query model, a local daemon core, read-only Runtime method dispatch, a KDE-facing D-Bus client for read-only planning calls, a packaged libexec activation wrapper, an unprivileged activation-file installer, a Linux session-bus smoke adapter for read-only calls, a KDE-safe Compatibility Center read model that can read through D-Bus, and a seven-entry-point KDE integration status model. It does not yet ship a full KDE image, live production D-Bus daemon ownership, enabled Wine backend, enabled VM backend, or AI service.

Run the focused scaffold test:

```text
ruby scripts/verify_layout.rb
ruby scripts/fetch_buildroot.rb --verify-lock
ruby -Ilib test/test_action_review_receipt.rb
ruby -Ilib test/test_compatibility_acquisition_preflight.rb
ruby -Ilib test/test_compatibility_action_queue.rb
ruby -Ilib test/test_compatibility_artifact_manifest.rb
ruby -Ilib test/test_compatibility_install_plan.rb
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
ruby -Ilib test/test_container.rb
ruby scripts/container.rb fetch-sources
ruby -Ilib test/test_buildroot.rb
ruby -Ilib test/test_qemu.rb
ruby -Ilib test/test_serial_log.rb
ruby -Ilib test/test_milestone.rb
ruby scripts/full_smoke.rb
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
ruby -Ilib test/test_recipe_store.rb
ruby -Ilib test/test_recipe_registry.rb
ruby -Ilib test/test_recipe_trust_policy.rb
ruby -Ilib test/test_recipe_install_gate.rb
ruby -Ilib test/test_registry_backed_recipe_store.rb
ruby -Ilib test/test_desktop_entry.rb
ruby -Ilib test/test_file_association_model.rb
ruby -Ilib test/test_dbus_runtime_client.rb
ruby -Ilib test/test_dolphin_service_menu.rb
ruby -Ilib test/test_file_open_request.rb
ruby -Ilib test/test_launch_request.rb
ruby -Ilib test/test_notification_request.rb
ruby -Ilib test/test_portal_access_policy.rb
ruby -Ilib test/test_portal_request_model.rb
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
ruby -Ilib test/test_runtime_dbus_smoke_script.rb
ruby -Ilib test/test_kde_center_model.rb
ruby -Ilib test/test_kde_integration_status.rb
ruby -Ilib test/test_kwin_window_rule.rb
ruby -Ilib test/test_krunner_model.rb
ruby scripts/container.rb runtime-dbus-smoke
ruby scripts/container.rb kde-center-dbus-smoke
```

# Xnix

Xnix is an atomic KDE Plasma desktop project focused on making existing Windows applications feel native on Linux.

The project is currently at `v0.2.195`. KDE Plasma is the only supported flagship desktop. An independent Compatibility Runtime manages application recipes, Wine/VM backends, Runtime application catalog previews, Runtime engine catalog previews, Runtime run plan previews, desktop activation manifests, KDE integration status, KDE shell integration plans, KDE application surface plans, KDE entrypoint previews, KDE entrypoint action previews, KDE action queue previews, KDE action review previews, KDE action preflight previews, KDE action receipt previews, KDE action status previews, KDE action card previews, KDE action card deck previews, KDE Compatibility Center page previews, KDE Compatibility Center page section previews, KDE Compatibility Center page section detail previews, desktop activation bundle previews, desktop activation preflight previews, desktop activation staging previews, desktop activation transaction previews, desktop activation status previews, backend environment previews, backend binding previews, Dolphin drag-and-drop previews, Dolphin AI analysis previews, desktop resource bridge plans, desktop entry plans, desktop icon plans, launch intents, execution request previews, execution review previews, execution decision previews, execution preflight previews, execution resource grant previews, execution transaction previews, execution session previews, execution session status previews, task manager identity plans, KWin window rule plans, file association plans, notification plans, tray status, KRunner query plans, Portal request plans, Compatibility Center action queues, review receipts, and center summaries, managed artifact manifests, managed install plans, managed acquisition preflight, managed package source selection, managed application state roots, managed backend binding, backend capability matrices, backend selection plans, environment, and lifecycle status, execution readiness, Runtime-backed compatibility settings, settings change plans, compatibility mode switch plans, compatibility permission review plans, compatibility review flow plans, diagnostics, Runtime diagnostics previews, AI diagnostic inputs and recommendations, AI repair approval gates, Runtime service binding status, Runtime live owner gates, Runtime owner smoke plans, Runtime method parity manifests, Runtime owner process previews, Runtime owner route manifests, Runtime owner recipe trust previews, Runtime owner readiness views, Runtime write gates, compatibility repair and test previews, snapshots, and rollback without exposing implementation details in normal desktop entry points. Important Runtime product logic is expected to move into Go first, with C kept for low-level or already-owned Runtime policy surfaces; Ruby remains the test and development-tool harness.

See [PRODUCT_OVERVIEW.md](PRODUCT_OVERVIEW.md) for the roadmap and technical choices. See [AGENTS.md](AGENTS.md) for contribution and safety rules.

## Current Verification

The first end-to-end Buildroot run has completed successfully: Xnix boots its Linux kernel and initramfs in constrained QEMU, acquires `10.0.2.15` through user-mode DHCP, and accepts a loopback-only key-authenticated SSH probe.

The v0.2.195 Runtime foundation defines a registry-backed Go Runtime desktop identity planner, Runtime application catalog renderer, Runtime engine catalog renderer, Runtime run plan renderer, Compatibility Center renderer, backend-selection renderer, backend-environment renderer, backend-binding renderer, backend capability and lifecycle renderers, Runtime repair-plan renderer, Runtime test-plan renderer, Runtime test-result renderer, Runtime write-gate renderer, Runtime diagnostics renderer, Runtime service binding renderer, Runtime live owner gate renderer, Runtime owner smoke plan renderer, Runtime method parity manifest renderer, Runtime owner process renderer, Runtime owner route manifest renderer, Runtime owner recipe trust renderer, Runtime owner readiness renderer, and the KDE-facing preview renderers needed for the seven official entry points. The remaining C Runtime core still provides stable identity, ABI-shaped policy surfaces, D-Bus smoke coverage, and lower-level compatibility records while those routes continue migrating into Go. Live production D-Bus ownership and enabled compatibility backend launch remain pending.

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
ruby -Ilib test/test_runtime_write_gate.rb
ruby -Ilib test/test_runtime_core.rb
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
ruby -Ilib test/test_compatibility_execution_readiness.rb
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
ruby -Ilib test/test_runtime_activation_smoke_script.rb
ruby -Ilib test/test_runtime_dbus_smoke_script.rb
ruby -Ilib test/test_kde_center_model.rb
ruby -Ilib test/test_kde_integration_status.rb
ruby -Ilib test/test_kwin_window_rule.rb
ruby -Ilib test/test_krunner_model.rb
ruby scripts/container.rb runtime-dbus-smoke
ruby scripts/container.rb runtime-activation-smoke
ruby scripts/container.rb kde-center-dbus-smoke
```

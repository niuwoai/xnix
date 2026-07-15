# Xnix

Xnix is an atomic KDE Plasma desktop project focused on making existing Windows applications feel native on Linux.

The project is currently at `v0.2.170`. KDE Plasma is the only supported flagship desktop. An independent Compatibility Runtime manages application recipes, Wine/VM backends, desktop activation manifests, KDE integration status, KDE shell integration plans, KDE application surface plans, KDE entrypoint previews, KDE entrypoint action previews, KDE action queue previews, KDE action review previews, KDE action preflight previews, KDE action receipt previews, KDE action status previews, KDE action card previews, KDE action card deck previews, KDE Compatibility Center page previews, KDE Compatibility Center page section previews, KDE Compatibility Center page section detail previews, desktop activation bundle previews, desktop activation preflight previews, desktop activation staging previews, desktop activation transaction previews, desktop activation status previews, backend environment previews, backend binding previews, Dolphin drag-and-drop previews, Dolphin AI analysis previews, desktop resource bridge plans, desktop entry plans, desktop icon plans, launch intents, execution request previews, execution review previews, execution decision previews, execution preflight previews, execution resource grant previews, execution transaction previews, execution session previews, execution session status previews, task manager identity plans, KWin window rule plans, file association plans, notification plans, tray status, KRunner query plans, Portal request plans, Compatibility Center action queues, review receipts, and center summaries, managed artifact manifests, managed install plans, managed acquisition preflight, managed package source selection, managed application state roots, managed backend binding, backend capability matrices, backend selection plans, environment, and lifecycle status, execution readiness, Runtime-backed compatibility settings, settings change plans, compatibility mode switch plans, compatibility permission review plans, compatibility review flow plans, diagnostics, AI diagnostic inputs and recommendations, AI repair approval gates, Runtime service binding status, Runtime live owner gates, Runtime owner smoke plans, Runtime method parity manifests, Runtime write gates, compatibility tests, snapshots, and rollback without exposing implementation details in normal desktop entry points. Important Runtime product logic is expected to move into Go first, with C kept for low-level or already-owned Runtime policy surfaces; Ruby remains the test and development-tool harness.

See [PRODUCT_OVERVIEW.md](PRODUCT_OVERVIEW.md) for the roadmap and technical choices. See [AGENTS.md](AGENTS.md) for contribution and safety rules.

## Current Verification

The first end-to-end Buildroot run has completed successfully: Xnix boots its Linux kernel and initramfs in constrained QEMU, acquires `10.0.2.15` through user-mode DHCP, and accepts a loopback-only key-authenticated SSH probe.

The v0.2.170 Runtime foundation defines a registry-backed Go Runtime desktop identity planner, Compatibility Center renderer, backend-selection renderer, backend-environment renderer, backend-binding renderer, desktop-activation-preflight renderer, desktop-activation-staging renderer, desktop-activation-transaction renderer, desktop-activation-status renderer, execution-readiness renderer, launch-intent renderer, execution-request renderer, execution-review renderer, execution-decision renderer, execution-preflight renderer, execution-resource-grant renderer, execution-transaction renderer, execution-session renderer, execution-session-status renderer, KDE entrypoints renderer, KDE entrypoint-action renderer, KDE action-queue renderer, KDE action-review renderer, KDE action-preflight renderer, KDE action-receipt renderer, KDE action-status renderer, KDE action-card renderer, KDE action-card-deck renderer, KDE Compatibility Center page renderer, KDE Compatibility Center page-section renderer, KDE Compatibility Center page-section detail renderer, KRunner query renderer, Dolphin file-open renderer, Dolphin drag-and-drop renderer, Dolphin AI analysis renderer, KDE desktop activation bundle renderer, KDE desktop resource bridge renderer, KDE Portal request renderer, KDE mode-switch renderer, KDE settings-change renderer, KDE review-flow renderer, desktop entry renderer, desktop icon renderer, MIME association renderer, KDE notification renderer, KDE unified-settings renderer, KDE permission review renderer, KDE window identity renderer, KDE tray status renderer, and activation-staging desktop entry and file association sources plus a C Runtime core for stable Runtime identity, application catalog metadata, engine catalog selection, Portal access policy, Portal request planning, KDE shell integration planning, desktop resource bridge planning, compatibility mode switch planning, compatibility permission review planning, compatibility review flow planning, snapshot policy, compatibility snapshot plan policy, compatibility test plan policy, compatibility test result policy, execution readiness policy, launch intent policy, backend capability matrix policy, backend selection plan policy, backend lifecycle policy, KDE application surface policy, application state-root policy, install readiness policy, compatibility install plan policy, Compatibility Center action queue policy, Compatibility Center action review receipt policy, Compatibility Center summary policy, compatibility run plan policy, compatibility repair plan policy, desktop activation manifest policy, KDE integration status policy, desktop entry plan policy, task manager identity plan policy, KWin window rule plan policy, file association plan policy, notification plan policy, tray status policy, KRunner query plan policy, AI diagnostic input policy, AI diagnostic recommendation policy, AI repair approval gate policy, artifact manifest policy, acquisition preflight policy, package source policy, backend binding policy, settings policy, settings change policy, service binding policy, live owner gate policy, owner smoke plan policy, method parity manifest policy, recipe trust policy, recipe install gate policy, and write-gate decisions, application recipes, compatibility engine cataloging, managed artifact manifest modeling, managed install plan modeling, managed acquisition preflight, managed package source selection, managed application state root modeling, managed backend binding, backend capability matrix, backend selection planning, and lifecycle status modeling, Runtime-owned execution readiness, C-owned launch intent modeling, Runtime-backed compatibility settings, Runtime-owned settings change planning, Runtime-owned compatibility mode switch planning, Runtime-owned compatibility permission review planning, Runtime-owned compatibility review flow planning, Runtime-owned compatibility repair planning, Runtime-owned compatibility snapshot planning, Runtime-owned compatibility test planning, Runtime-owned compatibility test result summarization, Compatibility Center summary modeling, AI diagnostic input modeling, AI diagnostic recommendation modeling, AI repair approval gate modeling, Runtime service binding status modeling, Runtime live owner gate modeling, Runtime owner smoke planning, Runtime method parity manifests, Runtime write gates, recipe registry verification, registry-backed Runtime recipe loading, recipe trust probe reporting, recipe trust policy evaluation, recipe install gate evaluation, preflight-gated desktop activation, Portal access policy evaluation, C-owned Portal request planning, C-owned KDE shell integration planning, C-owned desktop resource bridge planning, C-owned compatibility mode switch planning, C-owned compatibility permission review planning, C-owned compatibility review flow planning, C-owned backend capability matrix planning, C-owned backend selection planning, C-owned KDE integration status planning, C-owned KDE application surface planning, C-owned desktop entry planning, C-owned task manager identity planning, C-owned KWin window rule planning, C-owned file association planning, C-owned notification planning, C-owned tray status planning, C-owned KRunner query planning, C-owned compatibility repair planning, C-owned compatibility snapshot planning, C-owned compatibility test planning, C-owned compatibility test result summarization, C-owned execution readiness summarization, C-owned launch intent summarization, C-owned backend lifecycle summarization, C-owned AI diagnostic input planning, C-owned AI diagnostic recommendation planning, C-owned AI repair approval gate planning, C-owned Compatibility Center summary planning, Portal request modeling, standard desktop entries, file association generation, managed launcher, Dolphin file-open, notification, settings, mode switching, tray status, task-manager identity, KWin window rule, desktop integration manifest models, C-owned desktop activation manifests, a staged desktop activation installer, and activation receipt rollback, a D-Bus service contract with read-only planning methods and explicitly gated write methods, a hardened systemd unit, a Plasma 6 Compatibility Center package skeleton, a Dolphin service menu, a Runtime-backed KDE integration status model, a Runtime-backed KDE shell integration model, a Runtime-backed KDE application surface model, a Runtime-backed desktop resource bridge model, a Runtime-backed compatibility mode switch model, a Runtime-backed compatibility permission review model, a Runtime-backed compatibility review flow model, a Runtime-backed backend capability matrix model, a Runtime-backed backend selection model, a Runtime-backed KRunner query model, a local daemon core, read-only Runtime method dispatch, a KDE-facing D-Bus client for read-only planning calls, a packaged libexec activation wrapper with installed Runtime libraries and assets, an unprivileged activation-file installer, a constrained packaged activation smoke, a Linux session-bus smoke adapter for read-only calls and write-gate errors in the constrained Linux container; the KDE read model and D-Bus client can consume that session-bus source for planning reads. Live production D-Bus ownership is explicitly gated and still pending.

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

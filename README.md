# Xnix

Xnix is an atomic KDE Plasma desktop project focused on making existing Windows applications feel native on Linux.

The project is currently at `v0.2.15`. KDE Plasma is the only supported flagship desktop. An independent Compatibility Runtime manages application recipes, Wine/VM backends, diagnostics, snapshots, and rollback without exposing implementation details in normal desktop entry points.

See [PRODUCT_OVERVIEW.md](PRODUCT_OVERVIEW.md) for the roadmap and technical choices. See [AGENTS.md](AGENTS.md) for contribution and safety rules.

## Current Verification

The first end-to-end Buildroot run has completed successfully: Xnix boots its Linux kernel and initramfs in constrained QEMU, acquires `10.0.2.15` through user-mode DHCP, and accepts a loopback-only key-authenticated SSH probe.

The v0.2.15 Runtime foundation defines application recipes, standard desktop entries, managed launcher, Dolphin file-open, notification, settings, tray status, task-manager identity, and desktop integration manifest models, a D-Bus service contract, a hardened systemd unit, a Plasma 6 Compatibility Center package skeleton, a Dolphin service menu, a local daemon core, read-only Runtime method dispatch, a packaged libexec activation wrapper, an unprivileged activation-file installer, a Linux session-bus smoke adapter for read-only calls, a KDE-safe Compatibility Center read model that can read through D-Bus, and a seven-entry-point KDE integration status model. It does not yet ship a full KDE image, production D-Bus daemon binding, Wine backend, VM backend, or AI service.

Run the focused scaffold test:

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
ruby scripts/container.rb prepare-ssh-test-key
ruby scripts/container.rb build-ssh-test-system
ruby scripts/container.rb ssh-smoke
ruby -Ilib test/test_application_recipe.rb
ruby -Ilib test/test_recipe_store.rb
ruby -Ilib test/test_desktop_entry.rb
ruby -Ilib test/test_dbus_runtime_client.rb
ruby -Ilib test/test_dolphin_service_menu.rb
ruby -Ilib test/test_file_open_request.rb
ruby -Ilib test/test_launch_request.rb
ruby -Ilib test/test_notification_request.rb
ruby -Ilib test/test_settings_model.rb
ruby -Ilib test/test_tray_status_model.rb
ruby -Ilib test/test_task_manager_identity.rb
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

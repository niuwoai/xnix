# Xnix

Xnix is an atomic KDE Plasma desktop project focused on making existing Windows applications feel native on Linux.

The project is currently at `v0.2.0`. KDE Plasma is the only supported flagship desktop. An independent Compatibility Runtime manages application recipes, Wine/VM backends, diagnostics, snapshots, and rollback without exposing implementation details in normal desktop entry points.

See [PRODUCT_OVERVIEW.md](PRODUCT_OVERVIEW.md) for the roadmap and technical choices. See [AGENTS.md](AGENTS.md) for contribution and safety rules.

## Current Verification

The first end-to-end Buildroot run has completed successfully: Xnix boots its Linux kernel and initramfs in constrained QEMU, acquires `10.0.2.15` through user-mode DHCP, and accepts a loopback-only key-authenticated SSH probe.

The v0.2.0 Runtime foundation defines application recipes, standard desktop entries, a D-Bus service contract, a hardened systemd unit, and a Plasma 6 Compatibility Center package skeleton. It does not yet ship a full KDE image, Wine backend, VM backend, or AI service.

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
ruby -Ilib test/test_desktop_entry.rb
ruby -Ilib test/test_runtime_contract.rb
```

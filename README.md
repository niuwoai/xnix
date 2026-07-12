# Xnix

Xnix is a learning-oriented Linux-compatible system project.

The project is currently at `v0.1.6`. Its first target is a minimal x86_64 Linux system that boots in QEMU and manages OpenSSH `sshd` through its init process.

See [PRODUCT_OVERVIEW.md](PRODUCT_OVERVIEW.md) for the roadmap and technical choices. See [AGENTS.md](AGENTS.md) for contribution and safety rules.

## Current Verification

Run the focused scaffold test:

```text
ruby scripts/verify_layout.rb
ruby scripts/fetch_buildroot.rb --verify-lock
ruby -Ilib test/test_container.rb
ruby scripts/container.rb fetch-sources
ruby -Ilib test/test_buildroot.rb
ruby -Ilib test/test_qemu.rb
```

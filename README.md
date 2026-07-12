# Xnix

Xnix is a learning-oriented Linux-compatible system project.

The project is currently at `v0.1.1`. Its first target is a minimal x86_64 Linux system that boots in QEMU and manages OpenSSH `sshd` through its init process.

See [PRODUCT_OVERVIEW.md](PRODUCT_OVERVIEW.md) for the roadmap and technical choices. See [AGENTS.md](AGENTS.md) for contribution and safety rules.

## Current Verification

Run the focused scaffold test:

```text
ruby scripts/verify_layout.rb
```

# Xnix Contribution Rules

## Project Scope

Xnix is a learning-oriented Linux-compatible system. Its first deliverable is an x86_64 image that boots in QEMU and whose init process reliably runs and manages the OpenSSH `sshd` service.

## Language

- All project-facing text, including source comments, documentation, CLI output, and test descriptions, must be English only.
- Do not introduce internationalization infrastructure during the current phase.

## Engineering Principles

- Make the system bootable, observable, and reproducible before adding features.
- Treat the Linux LTS kernel as the Linux ABI provider. Do not claim to reimplement the Linux kernel ABI.
- Prefer small, replaceable components: Linux LTS, Buildroot, BusyBox, initramfs, and QEMU.
- Keep host-machine impact minimal: no privileged containers, host networking, Docker socket mounts, or broad host-directory mounts.

## Versioning, Testing, and Commits

- Use semantic versions. Each code change must update the version, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md`.
- Commit every small version to Git after its targeted unit tests pass.
- Run targeted tests for each small version; do not substitute a narrow test for the full milestone test.
- At every tenth code version (for example, `0.1.10`, `0.1.20`), run a full build and QEMU smoke test. Fix all discovered project defects, rerun the complete test, and commit the repair before continuing.
- Do not commit secrets, private keys, passwords, tokens, local configuration, or build artifacts.

## File Discipline

- Use patch-based edits for project files; do not generate or overwrite files with shell redirection, `cat`, `echo`, pipelines, or heredocs.
- Reconsider splitting a source file above 2,000 lines; split it before it reaches 3,000 lines.

## Expected Layout

```text
boot/             Kernel and boot configuration
buildroot/        Buildroot external tree and configuration
overlay/          Root filesystem overlay
scripts/          Build, run, verification, and cleanup scripts
docs/             Design notes and learning material
experiments/      Future isolated kernel experiments
```

## Acceptance Baseline

- QEMU reaches an interactive shell or init log on the serial console.
- Serial logs are persisted for diagnosis.
- SSH is available only through a loopback-bound forwarded port and is managed by init.
- A clean checkout can reproduce the build.

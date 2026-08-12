# q4 Remote Build Policy

This policy keeps the macOS development host responsive while preserving Xnix's restricted build and smoke-test safety boundary.

## Default Rule

Compile-heavy work runs on q4 by default.

The macOS host is reserved for lightweight checks, source editing, Git operations, and dry-run planning output. Do not run local Go compilation, QEMU boot tests, Wine GUI tests, Docker heavy smoke, or full milestone smoke on macOS unless a human explicitly approves that local run.

## q4 Build Host

- Default SSH target: `root@q4`
- Override target: `XNIX_REMOTE_HOST`
- Remote source roots must stay under `/home/xnix-*` or `/tmp/xnix-*`.
- Remote build caches must stay under `/home/xnix-*` or `/tmp/xnix-*`.
- Do not sync secrets, private keys, host-only private material, broad host directories, `.git`, local caches, or build artifacts to q4.

## Local Commands That Are Allowed by Default

These commands are intentionally lightweight and may run on the macOS host:

```text
ruby scripts/verify_layout.rb
git diff --check
ruby scripts/remote_go_build.rb
ruby scripts/remote_go_test.rb
ruby scripts/q4_full_smoke.rb
```

The remote scripts without `--execute` emit plans only. They do not compile, boot QEMU, run Wine, or mutate q4.

## Remote Commands for Compilation and Heavy Validation

Use these commands when a task requires compilation or heavy smoke coverage:

```text
ruby scripts/remote_go_build.rb --execute
ruby scripts/remote_go_test.rb --execute
ruby scripts/q4_full_smoke.rb --execute
ruby scripts/q4_staged_desktop_external_winapp_smoke.rb --execute
ruby scripts/q4_runtime_run_plan_execution_smoke.rb --execute
```

The q4 scripts sync a constrained source set, avoid macOS host compilation, and keep report artifacts under the checkout's `output/` tree or another explicitly scoped Xnix path.

## Local Override

Local compilation is blocked by policy unless a human explicitly approves it. If an emergency local run is approved, use a narrow override such as:

```text
XNIX_ALLOW_LOCAL_GO_COMPILE=1
```

The override must be scoped to the one approved command. Do not persist it in shell profiles, committed scripts, or shared project configuration.

## Safety Expectations

- No privileged containers.
- No host networking.
- No Docker socket mounts.
- No broad host-directory mounts.
- No host-root mutation.
- No secret transfer to q4.
- No full-smoke substitution with a narrow test at twentieth-version gates.

## Contributor Checklist

- Prefer q4 for any command that invokes `go build`, `go test`, `go run`, QEMU, Wine, or full smoke.
- Review the remote plan before first execution when command scope is unclear.
- Keep q4 outputs redacted and fetch only structured JSON, Markdown reports, serial logs, and explicitly requested evidence.
- Record q4 validation in `CHANGELOG.md` and `PRODUCT_OVERVIEW.md` when it gates a release checkpoint.

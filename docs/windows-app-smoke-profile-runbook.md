# Windows App Smoke Profile Runbook

> Last updated: 2026-07-24 | Current version: v0.2.640-rc40

This runbook is the shortest path from an existing Windows executable to repeatable Xnix smoke evidence.

Use a smoke profile when an application needs sidecar DLLs, config files, resource folders, runner pre-arguments, a CrossOver bottle selector, or a non-marker success mode. The profile is an operator-local input. Do not commit profiles that contain private host paths, private bottle names, credentials, customer data, or proprietary application locations.

## Template

Start from [docs/examples/windows-app-smoke-profile.template.json](examples/windows-app-smoke-profile.template.json).

Copy it to an ignored local path, then replace the placeholder values:

```text
cp docs/examples/windows-app-smoke-profile.template.json .local/xnix/winapp-smoke/my-app.profile.json
```

Profile fields:

| Field | Purpose | Report exposure |
| --- | --- | --- |
| `executable_path` | Existing Windows `.exe` path | Basename only |
| `working_directory` | Directory used as the process working directory | Mode only |
| `runner_path` | Wine-compatible runner path | Not exposed |
| `runner_bottle` | Optional CrossOver-style bottle selector | Count only |
| `runner_arguments` | Runner arguments before the executable path | Count only |
| `arguments` | Application arguments after the executable path | Not exposed |
| `state_root` | Isolated runtime state root | Not exposed |
| `timeout` | Smoke timeout such as `30s` or `2m` | Not exposed |
| `expected_marker` | Strong marker proof for console fixtures | Marker value |
| `success_mode` | `marker`, `exit-code`, or `startup-window` | Mode only |
| `redact_output` | Omit raw stdout and stderr from report payloads | Boolean |

## Local Smoke

Run through the report script:

```text
ruby scripts/winapp_smoke.rb --format json --profile .local/xnix/winapp-smoke/my-app.profile.json
```

Or call the Go Runtime command directly:

```text
go run ./cmd/xnix-runtime-go windows-app-run-smoke --profile .local/xnix/winapp-smoke/my-app.profile.json
```

Both paths keep Docker, QEMU, Colima, network checks, package managers, privileged containers, host networking, Docker socket mounts, broad host mounts, and host-root mutation disabled.

## Choosing Success Mode

- Use `marker` when the app or fixture can print `expected_marker`.
- Use `exit-code` for console or setup helpers that start and exit cleanly.
- Use `startup-window` for GUI-style apps that are expected to remain alive until the timeout window.

## Interpreting Results

- `passed` means the configured success mode observed its proof.
- `skipped` with `windows compatibility runner unavailable` means no Wine-compatible runner was found or supplied.
- `profile_supplied: true` proves the report used a profile input.
- `working_directory_mode: operator-supplied` proves the profile or CLI supplied an explicit working directory.
- `runner_argument_count` proves runner argument forwarding without exposing the raw values.

The report must never expose raw executable paths, runner paths, working directories, private bottle names, runner arguments, Docker commands, QEMU commands, or backend internals.

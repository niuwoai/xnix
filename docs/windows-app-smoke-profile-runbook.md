# Windows App Smoke Profile Runbook

> Last updated: 2026-07-24 | Current version: v0.2.640-rc48

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
| `skip_bootstrap` | Optional local-runner bypass for Wine prefix bootstrap when first-run bootstrap hangs or is already complete | Boolean |
| `redact_output` | Omit raw stdout and stderr from report payloads | Boolean |

## Preflight

Validate the profile before launch:

```text
go run ./cmd/xnix-runtime-go windows-app-smoke-profile-preflight --profile .local/xnix/winapp-smoke/my-app.profile.json
```

Or use the report script without launching the Windows app:

```text
ruby scripts/winapp_smoke.rb --format json --profile .local/xnix/winapp-smoke/my-app.profile.json --preflight-only
```

Preflight checks the schema, executable file, Windows `MZ` signature, PE machine architecture, working directory, state-root setting, success mode, runner argument count, app argument count, and runner diagnostics. It does not start Wine or the Windows application.

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

Direct `--exe` smoke runs also validate the Windows `MZ` executable signature and PE machine architecture before runner resolution, working-directory setup, or app execution.

When diagnosing an existing Wine prefix or a runner whose `wineboot --init` hangs, use `--skip-bootstrap` or set `skip_bootstrap: true` in a local profile. This still prepares the architecture-scoped prefix directory and launches through the managed runner environment, but it does not invoke companion `wineboot` before the app.

## Choosing Success Mode

- Use `marker` when the app or fixture can print `expected_marker`.
- Use `exit-code` for console or setup helpers that start and exit cleanly.
- Use `startup-window` for GUI-style apps that are expected to remain alive until the timeout window.

## Interpreting Results

- `passed` means the configured success mode observed its proof.
- `skipped` with `windows compatibility runner unavailable` means no Wine-compatible runner was found or supplied.
- `profile_supplied: true` proves the report used a profile input.
- `windows_executable_signature_observed: true` proves the profile points to a Windows PE-style executable header.
- Direct smoke reports also expose `executable_format: pe-mz` when the supplied `--exe` target is a Windows PE-style executable.
- `executable_architecture: x86_64` or `x86` proves the Runtime parsed the PE machine field for Wine prefix policy.
- `executable_architecture_supported: true` means the current smoke path accepts the executable architecture before invoking a runner.
- `wine_architecture: win64` or `win32` proves the Runtime selected the Wine prefix architecture from the PE machine field.
- `wine_prefix_mode: architecture-scoped` proves the Runtime will keep win64 and win32 prefixes separate under the managed state root.
- `wine_prefix_prepared: true` appears only after a direct smoke run creates the architecture-scoped prefix; profile preflight reports the intended mode without creating it.
- `wine_bootstrap_skipped: true` proves the operator requested the local smoke to bypass companion `wineboot` before app execution.
- `working_directory_mode: operator-supplied` proves the profile or CLI supplied an explicit working directory.
- `runner_argument_count` proves runner argument forwarding without exposing the raw values.

The report must never expose raw executable paths, runner paths, working directories, private bottle names, runner arguments, Docker commands, QEMU commands, or backend internals.

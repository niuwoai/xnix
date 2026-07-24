# Windows App Smoke Profile Runbook

> Last updated: 2026-07-24 | Current version: v0.2.640-rc59

This runbook is the shortest path from an existing Windows executable to repeatable Xnix smoke evidence.

Use a smoke profile when an application needs sidecar DLLs, config files, resource folders, runner pre-arguments, a CrossOver bottle selector, or a non-marker success mode. The profile is an operator-local input. Do not commit profiles that contain private host paths, private bottle names, credentials, customer data, or proprietary application locations.

## Template

Start from [docs/examples/windows-app-smoke-profile.template.json](examples/windows-app-smoke-profile.template.json).

Copy it to an ignored local path, then replace the placeholder values:

```text
cp docs/examples/windows-app-smoke-profile.template.json .local/xnix/winapp-smoke/my-app.profile.json
```

Or render a reusable profile from CLI settings:

```text
ruby scripts/winapp_smoke.rb --format json --exe path/to/application.exe --state-root .local/xnix/winapp-smoke/my-app-state --stage-app-dir --write-profile .local/xnix/winapp-smoke/my-app.profile.json
```

The script asks the Go Runtime to render the profile through `windows-app-smoke-profile-render`, then writes the JSON to the requested operator-local path. Reports expose `profile_write_invoked` and `profile_written`; they do not expose the profile path.

To also create a managed launcher script and `.desktop` file under the profile state root:

```text
ruby scripts/winapp_smoke.rb --format json --exe path/to/application.exe --state-root .local/xnix/winapp-smoke/my-app-state --stage-app-dir --write-profile .local/xnix/winapp-smoke/my-app.profile.json --write-launcher-bundle --app-id org.xnix.myapp --app-name "My Windows App"
```

The launcher bundle is written by `windows-app-launcher-bundle-record`. Reports expose `launcher_bundle_write_invoked`, `launcher_bundle_written`, safe desktop file names, and safe launcher script names; they do not expose profile paths, state-root paths, or runner paths.

For development checkouts where `xnix-runtime-go` is not installed into `PATH`, add runtime argv pieces:

```text
ruby scripts/winapp_smoke.rb --format json --exe path/to/application.exe --state-root .local/xnix/winapp-smoke/my-app-state --stage-app-dir --write-profile .local/xnix/winapp-smoke/my-app.profile.json --write-launcher-bundle --app-id org.xnix.myapp --runtime-bin go --runtime-arg run --runtime-arg ./cmd/xnix-runtime-go
```

Reports expose only `runtime_argument_count`; they do not expose the runtime argv values.

To create a product-style desktop launcher that preflights first and launches only when ready, use launch mode:

```text
ruby scripts/winapp_smoke.rb --format json --exe path/to/application.exe --state-root .local/xnix/winapp-smoke/my-app-state --stage-app-dir --write-profile .local/xnix/winapp-smoke/my-app.profile.json --write-launcher-bundle --app-id org.xnix.myapp.launch --launcher-mode launch
```

`--launcher-mode launch` calls `windows-app-launch-profile`, which returns blocked readiness evidence when the runner is unavailable and delegates to redacted execution only after preflight reports ready.

To create a desktop launcher that only validates the profile without starting Wine or the Windows application, use preflight mode:

```text
ruby scripts/winapp_smoke.rb --format json --exe path/to/application.exe --state-root .local/xnix/winapp-smoke/my-app-state --stage-app-dir --write-profile .local/xnix/winapp-smoke/my-app.profile.json --write-launcher-bundle --app-id org.xnix.myapp.preflight --launcher-mode preflight
```

`--launcher-mode execute` is the compatibility default and calls `windows-app-run-smoke`. `--launcher-mode preflight` calls `windows-app-smoke-profile-preflight`, preserving a KDE-friendly click target for readiness checks before the real runner is available. Reports expose the safe `launcher_mode` and `launcher_command` values.

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
| `stage_app_dir` | Optional copy of the application directory into the isolated Runtime state root before launch | Boolean and counts |
| `redact_output` | Omit raw stdout and stderr from report payloads | Boolean |

## Preflight

Validate the profile before launch:

```text
go run ./cmd/xnix-runtime-go windows-app-smoke-profile-preflight --profile .local/xnix/winapp-smoke/my-app.profile.json
```

Product-style launch with a preflight gate:

```text
go run ./cmd/xnix-runtime-go windows-app-launch-profile --profile .local/xnix/winapp-smoke/my-app.profile.json
```

## Known App Materialization

For a Runtime-owned known app such as the pinned 7-Zip standalone executable, first fetch the artifact through the existing known-app acquisition path when network access is explicitly allowed by the operator. Once the artifact is cached and checksum-verified, materialize a product launch profile and managed launcher without downloading or running the app:

```text
go run ./cmd/xnix-runtime-go windows-known-app-launch-profile-materialize --app 7zr --cache-root .cache/xnix/known-winapps --state-root .local/xnix/known-apps/7zr-state --runtime-bin go --runtime-arg run --runtime-arg ./cmd/xnix-runtime-go --runner path/to/wine --runner-bottle bottle-name --runner-arg "--private-runner-option" --skip-bootstrap
```

If the artifact is missing or fails checksum verification, the command returns safe `skipped` evidence and writes no profile. If the artifact is verified, it writes a reusable profile and a `launch` mode launcher bundle that calls `windows-app-launch-profile`. The optional runner settings are stored in the local profile, but reports expose only `runner_configured`, `runner_bottle_configured`, `runner_argument_count`, and `skip_bootstrap`; they do not expose raw runner paths, bottle names, or runner argv values.

To combine offline cache checking, optional explicit acquisition, and materialization in one Runtime-owned command:

```text
go run ./cmd/xnix-runtime-go windows-known-app-prepare-launch-profile --app 7zr --cache-root .cache/xnix/known-winapps --state-root .local/xnix/known-apps/7zr-state --runtime-bin go --runtime-arg run --runtime-arg ./cmd/xnix-runtime-go --runner path/to/wine --runner-bottle bottle-name --runner-arg "--private-runner-option" --skip-bootstrap
```

The command is offline by default. Add `--allow-download` only when the operator intends to fetch the pinned artifact from the catalog. After checksum verification, it writes the same launch profile and managed launcher bundle as the materializer, including any explicit runner and bootstrap settings supplied by the operator.

To prepare and immediately attempt a Runtime-owned launch in one step:

```text
go run ./cmd/xnix-runtime-go windows-known-app-prepare-and-launch-profile --app 7zr --cache-root .cache/xnix/known-winapps --state-root .local/xnix/known-apps/7zr-state --runtime-bin go --runtime-arg run --runtime-arg ./cmd/xnix-runtime-go --runner path/to/wine --runner-bottle bottle-name --runner-arg "--private-runner-option" --skip-bootstrap
```

The command emits combined `prepare_payload` and `launch_payload` evidence. It still checks the cache offline by default and still requires `--allow-download` for acquisition. If the artifact is verified but the runner is unavailable, the launch payload reports `blocked` with `launch_attempted=false`. If the runner is ready, the same command launches through `windows-app-launch-profile`, forces redacted output, and reports Runtime smoke evidence without exposing raw local paths or runner details.

For the unified Runtime-owned known app execution entrypoint, select a backend explicitly:

```text
go run ./cmd/xnix-runtime-go windows-known-app-run --backend local --app 7zr --cache-root .cache/xnix/known-winapps --state-root .local/xnix/known-apps/7zr-state --runner path/to/wine
go run ./cmd/xnix-runtime-go windows-known-app-run --backend guest-wine --app 7zr --cache-root .cache/xnix/known-winapps --key .cache/xnix/ssh-test-key/id_ed25519 --timeout 90s
```

`--backend local` reuses `windows-known-app-prepare-and-launch-profile`. `--backend guest-wine` reuses the loopback SSH Wine guest smoke path and assumes the Wine-capable guest is already running; it does not start QEMU or Docker by itself. The unified result reports backend selection, backend readiness, checksum verification, launch attempt, runner availability, marker observation, and either `local_payload` or `guest_payload`.

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

To render the profile JSON without the Ruby wrapper:

```text
go run ./cmd/xnix-runtime-go windows-app-smoke-profile-render --exe path/to/application.exe --state-root .local/xnix/winapp-smoke/my-app-state --stage-app-dir
```

To write a launcher bundle directly from an existing profile:

```text
go run ./cmd/xnix-runtime-go windows-app-launcher-bundle-record --profile .local/xnix/winapp-smoke/my-app.profile.json --app-id org.xnix.myapp --name "My Windows App"
```

Both paths keep Docker, QEMU, Colima, network checks, package managers, privileged containers, host networking, Docker socket mounts, broad host mounts, and host-root mutation disabled.

Direct `--exe` smoke runs also validate the Windows `MZ` executable signature and PE machine architecture before runner resolution, working-directory setup, or app execution.

When diagnosing an existing Wine prefix or a runner whose `wineboot --init` hangs, use `--skip-bootstrap` or set `skip_bootstrap: true` in a local profile. This still prepares the architecture-scoped prefix directory and launches through the managed runner environment, but it does not invoke companion `wineboot` before the app.

When an app depends on sidecar DLLs, config files, or resource folders next to the executable, use `--stage-app-dir` or set `stage_app_dir: true` in a local profile. The Runtime copies the executable's directory, or the explicit `working_directory` when supplied, into `state_root/app-workspace` and launches the staged executable from that managed workspace. Reports expose only `application_workspace_mode`, `application_staged`, file counts, and byte counts, not source or staged paths.

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
- `application_workspace_mode: staged-application-directory` proves the Runtime copied the application directory into its managed workspace before launch.
- `application_staged: true` proves direct smoke completed that copy; profile preflight reports only the intended workspace mode.
- `working_directory_mode: operator-supplied` proves the profile or CLI supplied an explicit working directory.
- `runner_argument_count` proves runner argument forwarding without exposing the raw values.

The report must never expose raw executable paths, runner paths, working directories, private bottle names, runner arguments, Docker commands, QEMU commands, or backend internals.

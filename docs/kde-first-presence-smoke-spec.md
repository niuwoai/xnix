# KDE-First Presence Smoke Specification

> Last updated: 2026-07-15 | Proposed command: `ruby scripts/kde_first_presence_smoke.rb`

This specification defines the first KDE-focused product smoke for Xnix. The smoke should prove that one recipe can appear across the first-release KDE surface as a normal Linux application while the Xnix Compatibility Runtime keeps all compatibility policy and launch authority.

The smoke is intentionally non-executing. It must not start compatibility backends, claim production D-Bus ownership, mutate the host root, require network access, use privileged containers, or expose backend implementation details.

## Purpose

The current repository has many focused previews and tests. The missing acceptance proof is a single smoke that checks the whole first milestone in one pass:

> A digest-verified Windows application recipe appears as a normal KDE application across the seven first-release entry points, while Runtime gates keep execution and host mutation disabled.

This smoke is the bridge between isolated contract tests and product-level confidence.

## Inputs

Default inputs:

- Registry: `runtime/recipes/registry.json`
- Application id: `org.xnix.sample.notepad`
- Decision: `approved`
- File URI fixture: `file:///home/xnix/Documents/sample.txt`
- Mode: `development`

The smoke should accept flags for these values but must keep safe defaults.

Suggested CLI:

```text
ruby scripts/kde_first_presence_smoke.rb \
  --registry runtime/recipes/registry.json \
  --app org.xnix.sample.notepad \
  --decision approved \
  --mode development \
  file:///home/xnix/Documents/sample.txt
```

## Required Runtime Preview Commands

The smoke should call `xnix-runtime-go` through the local Go command or packaged test command and inspect JSON payloads from these previews:

```text
desktop-identity-plan
desktop-entry-preview
mimeapps-preview
desktop-activation-bundle-preview
desktop-activation-staging-preview
desktop-activation-transaction-preview
desktop-activation-status-preview
kde-entrypoints-preview
kde-action-card-deck-preview
kde-center-page-preview
kde-center-page-sections-preview
kde-center-page-section-detail-preview
file-open-preview
dolphin-drop-preview
dolphin-ai-analysis-preview
window-identity-preview
tray-status-preview
notification-preview
settings-preview
runtime-owner-route-manifest-preview
runtime-method-parity-manifest-preview
runtime-write-gate-preview or GetRuntimeWriteGate equivalent
```

If a preview is only available through an existing Ruby or C-backed command, the smoke may use that command temporarily, but it must report the source as a migration gap.

## Seven Entry-Point Assertions

### 1. Start Menu

Evidence sources:

- `desktop-identity-plan`
- `desktop-entry-preview`
- `desktop-activation-bundle-preview`
- `kde-entrypoints-preview`

Required assertions:

- The application has a normal user-visible name.
- The generated desktop file id is stable.
- The launcher action routes through the Runtime.
- The launcher is visible in the KDE entrypoint payload.
- Launch remains disabled or gated by Runtime write gates.
- No raw executable path, raw backend command, or internal state-root path appears.

### 2. Task Manager

Evidence sources:

- `window-identity-preview`
- `kde-entrypoints-preview`
- `kde-center-page-preview`

Required assertions:

- The task-manager entrypoint is present.
- Pinning or restore intent is represented as identity metadata only.
- KWin or task-manager policy remains Runtime-owned.
- Backend launch, live window observation, and task-manager activation remain disabled.

### 3. File Manager

Evidence sources:

- `file-open-preview`
- `dolphin-drop-preview`
- `dolphin-ai-analysis-preview`
- `mimeapps-preview`
- `desktop-activation-staging-preview`
- `kde-entrypoints-preview`

Required assertions:

- Dolphin file-open routes through the Runtime.
- Drag-and-drop routes through the same Portal-mediated model.
- AI analysis exposes only file count, extension, and safe disclosure metadata.
- MIME association output targets the generated desktop file.
- The Dolphin service menu is represented in staging.
- No direct file read, file contents, host path disclosure, or permission grant occurs.

### 4. System Tray

Evidence sources:

- `tray-status-preview`
- `desktop-activation-bundle-preview`
- `kde-center-page-preview`

Required assertions:

- The tray entrypoint is represented.
- Tray status is user-facing and Runtime-owned.
- Live tray bridge enablement remains disabled.
- No backend process is started from tray status.

### 5. Notification Center

Evidence sources:

- `notification-preview`
- `kde-action-card-deck-preview`
- `kde-center-page-preview`

Required assertions:

- Notification plan payloads use desktop-safe copy.
- Notification actions route to review or navigation, not execution.
- Repair execution, settings persistence, launch, and host mutation remain disabled.

### 6. AI Compatibility Center

Evidence sources:

- `kde-center-page-preview`
- `kde-center-page-sections-preview`
- `kde-center-page-section-detail-preview`
- `kde-action-card-deck-preview`
- `dolphin-ai-analysis-preview`

Required assertions:

- The Compatibility Center page composes Runtime read models.
- The action deck contains entrypoint cards.
- Diagnostics or AI sections expose privacy-safe inputs only.
- AI provider calls, file contents, host paths, raw commands, automatic repair, and backend launch remain disabled.

### 7. Unified Settings

Evidence sources:

- `settings-preview`
- `kde-center-page-preview`
- `kde-entrypoints-preview`

Required assertions:

- Settings are user-facing: mode, priority, file access, camera, network, and snapshots.
- Settings changes remain Runtime-owned and gated.
- No developer-only environment terminology is exposed.
- Settings persistence remains disabled unless a later package explicitly enables it.

## Global Safety Assertions

Every inspected payload must satisfy:

- `host_root_modified` is absent or false.
- `network_required` is absent or false.
- `privileged_container_required` is absent or false.
- `backend_details_exposed` is absent or false.
- `raw_command_exposed` is absent or false.
- `raw_windows_executable_exposed` is absent or false.
- `execution_started` is absent or false.
- `backend_launch_enabled` is absent or false.
- `launch_enabled` is absent or false unless explicitly described as a disabled launcher surface.
- No field value contains raw backend commands, internal state-root paths, or direct user file paths.

The smoke should use a shared recursive JSON scanner for these assertions.

## Route and Contract Assertions

The smoke should inspect owner route and method parity previews:

- Method parity must pass.
- Ruby legacy route count must be zero.
- Write route gate must pass.
- Host safety boundary must pass.
- Any C-backed read-only routes should be reported as migration gaps, not smoke failures, unless the route is required by this smoke and missing.

## Staging Assertions

The smoke should inspect desktop activation staging and transaction previews:

- Planned files include:
  - desktop entry
  - Dolphin service menu
  - MIME apps list
  - desktop integration manifest
  - desktop activation receipt
- Paths are relative to the staging target model and do not expose a host root.
- Transaction commit remains disabled.
- Rollback steps exist for staged files.

## Output

The smoke should print a concise human-readable summary and optionally write JSON/Markdown reports later.

Suggested successful output:

```text
PASS: KDE-first presence smoke
application: org.xnix.sample.notepad
entrypoints: launcher, task-manager, file-manager, system-tray, notifications, compatibility-center, settings
route-baseline: 57 total, 40 Go-routed, 17 C-backed, 0 Ruby legacy
execution: disabled by Runtime gates
host-root: unchanged
```

Suggested failure output:

```text
FAIL: KDE-first presence smoke
failed_check: file-manager.portal_boundary
reason: file-open preview exposed a direct host path
```

## Failure Conditions

The smoke must fail if:

- Any required preview command fails.
- Any required entrypoint is missing.
- A payload exposes raw backend commands or raw executable paths.
- A payload exposes direct user file paths beyond the supplied safe URI fixture.
- A payload reports host-root mutation, network requirement, privileged container requirement, backend launch, or execution start.
- Write methods appear enabled.
- Method parity fails.
- Ruby legacy route count is non-zero.

The smoke should warn, not fail, for C-backed read-only routes that are already known migration gaps.

## Implementation Notes

- Prefer Ruby for the smoke harness because it is test orchestration, not durable Runtime product logic.
- Use Go previews for product data wherever available.
- Keep all temporary output under a test-controlled temporary directory.
- Do not shell out through pipelines or host-mutating commands.
- Do not run Docker or QEMU from this smoke.
- Do not modify `docs/claude-code-implementation-packages.md`.

## Relationship to Existing Tests

This smoke does not replace focused tests. It aggregates evidence from existing Runtime preview surfaces into one product-level first-milestone check.

Focused tests should continue to own detailed unit coverage. This smoke should answer the product question:

> Can KDE see one compatibility application as a normal Linux application across the seven agreed entry points while Runtime gates keep the system safe?

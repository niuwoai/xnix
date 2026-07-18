# Go Stability Release Train

> Last updated: 2026-07-18 | Baseline: v0.2.300 | Scope: KDE-first Windows application compatibility MVP

This document defines the stabilization plan for moving Xnix from preview-heavy evidence toward a runnable KDE-first Windows application MVP. It is intentionally operational: each small version must be independently reviewable, locally committed, and tested with a targeted command set. The original `v0.2.301`-`v0.2.320` train closed under its historical twentieth-version gate; beginning with `v0.2.321`, full builds, complete test suites, product smoke, and remote pushes run every thirty small versions, with the next gate at `v0.2.350`.

## Policy

- Use Go for durable Runtime product logic, long-running services, state machines, package acquisition, execution planning, and stable APIs.
- Use C only for low-level policy surfaces, ABI-shaped boundaries, and already-owned Runtime records.
- Use Ruby for tests, smoke scripts, reports, and lightweight developer tooling.
- Keep KDE as the shell and presentation layer; the Runtime owns policy, state, permissions, compatibility decisions, and execution gates.
- Keep real Wine, Proton, Windows VM launch, real Portal transport, production D-Bus ownership, host package-manager calls, privileged containers, host networking, Docker socket mounts, broad host mounts, and host-root mutation disabled until a specific gated train step enables a reviewed test path.

## Version Cadence

- Each small version is one coherent patch-level step.
- Each small version must update `VERSION`, `CHANGELOG.md`, `PRODUCT_OVERVIEW.md`, and `scripts/verify_layout.rb` when code or required layout changes.
- Each small version must be committed locally after its targeted tests pass.
- Do not push every small version. Beginning with `v0.2.321`, push only at each thirtieth-version train gate after the full build and bug-fix loop passes.
- If a small version needs a tag, create it locally with `v<version>` and push accumulated tags only at the configured full gate.
- Do not combine unrelated workstreams in one small version just to reach the train gate faster.

## Test Cadence

Every small version runs targeted tests only:

- `go test` for the touched Go packages and their CLI package when applicable.
- Focused Ruby tests for any touched Ruby model, report, or smoke script.
- `ruby scripts/verify_layout.rb`.
- `git diff --check`.

The historical first train ran this full gate at `v0.2.320`. Beginning with `v0.2.321`, every thirtieth small version runs it; the next full gate is `v0.2.350`:

- `go test ./...`.
- `ruby scripts/verify_layout.rb`.
- `ruby scripts/runtime_contract_drift_report.rb --format json`.
- `ruby scripts/implementation_evidence_report.rb --format json`.
- `ruby scripts/mainline_integration_review.rb --format json`.
- `ruby scripts/release_evidence_index.rb --format json`.
- `ruby scripts/merge_readiness_packet.rb --format json`.
- The full product build and constrained QEMU/product smoke for the train, with persisted logs and bug fixes before push.

If the full gate exposes project defects, fix them in the same full-gate release candidate before pushing. If the host environment cannot safely run the full product build or QEMU smoke, stop the push and record the missing authorization or environment blocker.

## Train Goal

The first train is `v0.2.301` through `v0.2.320`.

Target outcome by `v0.2.320`:

- KDE can consume one digest-verified application identity across launcher, task manager, file manager, tray, notifications, Compatibility Center, and settings.
- The Runtime has Go-owned readiness, trust, backend lifecycle, Portal, snapshot, execution, diagnostic, and support evidence for that application.
- A controlled test path can prove the product image and QEMU smoke readiness without weakening host safety.
- Real Windows application launch remains disabled unless a gated test-only launch step has explicit evidence and user authorization.

## Version Slices

| Version band | Primary objective | Expected evidence |
| --- | --- | --- |
| `0.2.301`-`0.2.304` | Close review gaps around KDE notification digest, signed recipe verifier evidence, and restricted product smoke packet preparation. | Go Runtime read models, CLI previews, focused Go/Ruby tests, layout verification. |
| `0.2.305`-`0.2.308` | Move more Runtime owner read routes from preview or bridge ownership into Go-owned service paths. | Route manifest diffs, method parity evidence, disabled-write tests, no production ownership. |
| `0.2.309`-`0.2.312` | Make one offline application fixture behave like a normal KDE application identity without launch. | Desktop entry, MIME, KRunner, tray, notification, settings, and Compatibility Center evidence for one fixture. |
| `0.2.313`-`0.2.316` | Tighten fake-mode execution, Portal, snapshot, diagnostics, and backend lifecycle joins around the same fixture. | State-root receipts, readiness graph, blocked launch reasons, redaction tests, no backend process start. |
| `0.2.317`-`0.2.319` | Prepare the restricted test-only launch boundary and product-image smoke packet. | Explicit launch gate state, blocked unsafe paths, product smoke metadata, no host mutation. |
| `0.2.320` | Full train gate: build, smoke, bug-fix loop, local version commit, and remote push. | Full test/build/QEMU or authorized smoke evidence, persisted logs, release evidence, pushed branch and tags. |

## Stop Conditions

Stop and report before continuing if a task requires:

- A production backend launch instead of a gated test path.
- Production D-Bus ownership.
- Real Portal transport.
- Host package installation.
- Privileged containers, host networking, Docker socket mounts, or broad host mounts.
- Host-root mutation.
- Secrets, signing keys, credentials, private webhook URLs, or private user data.
- Pushing before the configured full-version train gate.

## Per-Version Handoff Checklist

Use this checklist before committing each small version:

- The version implements one coherent slice.
- Durable product logic is in Go unless the change is explicitly test or tooling.
- KDE-facing output hides raw executable paths, state-root paths, host paths, backend commands, backend names, and backend details.
- Unsafe gates stay disabled unless the slice explicitly owns a reviewed test-only gate.
- Targeted tests passed.
- `ruby scripts/verify_layout.rb` passed.
- `git diff --check` passed.
- The local commit message names the version and the slice.
- The branch is not pushed unless this is the configured full-version train gate.

## Suggested Commit Message Shape

```text
feat(runtime): advance v0.2.301 kde notification digest evidence

- Add Go Runtime notification digest preview evidence
- Keep notification sending, backend launch, and host mutation disabled
- Run targeted Go/Ruby tests and layout verification
```

Use `fix`, `test`, `docs`, or `chore` instead of `feat` when the slice is not a new product capability.

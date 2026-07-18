# Claude Code Current Dispatch Picks

> Last updated: 2026-07-19 | Baseline: v0.2.371

This is the current short dispatch sheet for Claude Code. It is intentionally smaller than the full wave documents and should be used when choosing the next independent branch.

Use `docs/claude-code-stability-release-train.md` as the cadence authority where it does not conflict with the current goal: one coherent small version per local commit, targeted tests on every small version, and a full build plus complete tests every twenty small versions. The `v0.2.360` implementation checkpoint is present locally, but the full build, QEMU, and push gate still require explicit operator authorization.

Local work already contains implementation evidence for all seventh-wave tasks (`S7W1` through `S7W8`) plus `S6W1`, `S6W2`, `S6W3`, `S6W4`, `S6W5`, `S6W6`, `S6W7`, `S6W8`, `F5W7`, the `v0.2.321` offline fixture artifact-receipt read model, the `v0.2.322` offline fixture snapshot-baseline read model, the `v0.2.323` production Runtime service activation preflight read model, the `v0.2.324` test-only launch materialization plan receipt, the `v0.2.325` backend adapter no-op contract preview, the `v0.2.326` restricted owner smoke execution receipt, the `v0.2.327` test-only materialization fan-out to KDE surfaces, the `v0.2.328` backend adapter fixture binding audit, the `v0.2.329` restricted owner smoke receipt fan-out to readiness and support surfaces, the `v0.2.330` materialization fan-out owner-route audit, the `v0.2.331` adapter contract owner-route audit, the `v0.2.332` restricted owner smoke receipt fan-out owner-route audit, the `v0.2.333` materialization fan-out receipt-consumption split, the `v0.2.334` backend adapter redacted profile owner-route split, the `v0.2.335` restricted owner smoke receipt opaque lookup split, the `v0.2.336` materialization fan-out opaque receipt route audit, the `v0.2.337` restricted owner smoke fan-out owner-local read route, the `v0.2.338` materialization fan-out owner-managed opaque receipt lookup, the `v0.2.339` materialization fan-out owner-local read route, the `v0.2.340` materialization fan-out owner smoke coverage read model, the `v0.2.341` restricted owner smoke fan-out owner smoke coverage read model, the `v0.2.342` redacted adapter profile owner smoke coverage read model, the `v0.2.343` materialization smoke coverage audit refresh, the `v0.2.344` restricted owner smoke fan-out coverage audit refresh, the `v0.2.345` redacted adapter profile audit smoke coverage refresh, the `v0.2.346` production D-Bus gate review packet, the `v0.2.347` production D-Bus human authorization preflight, the `v0.2.348` production service activation gate-consumption preflight refresh, the `v0.2.349` Runtime write-gate production-consumption review, the `v0.2.350` route-by-route production D-Bus method review, the `v0.2.351` production rollback diagnostics review, the `v0.2.352` production desktop side-effect review, the `v0.2.353` production human authorization receipt consolidation, the `v0.2.354` production authorization consumption audit, the `v0.2.355` production receipt acceptance propagation preflight, the `v0.2.356` production receipt writer authorization review, the `v0.2.357` production receipt persistence threat review, the `v0.2.358` production receipt revocation visibility audit, the `v0.2.359` production receipt notification delivery gate audit, the `v0.2.360` production receipt notification action safety audit for review, renew, open Compatibility Center, dismiss, support-info, navigation, and support requests, the `v0.2.361` production receipt notification action request-object audit for review, renew, open Compatibility Center, dismiss, support-info, navigation, request creation, and support requests, the `v0.2.362` production receipt notification action dispatch authorization audit for dispatch authorization without dispatch, the `v0.2.363` production receipt notification action dispatch dry-run audit for dry-run only without real dispatch, request creation, or side effects, the `v0.2.364` production receipt notification action dry-run result visibility audit for redacted KDE and Runtime diagnostics visibility without result persistence or execution, the `v0.2.365` production receipt notification action dry-run result persistence authorization audit for result storage authorization without persistence, dry-run execution, or side effects, the `v0.2.366` production receipt notification action dry-run result retention redaction policy audit for retention and redaction rules without policy enforcement, persistence, or execution, the `v0.2.367` production receipt notification action dry-run result opaque lookup audit for owner-managed opaque identifiers without lookup enablement, lookup persistence, result persistence, execution, or side effects, the `v0.2.368` production receipt notification action dry-run result lookup route authorization audit for owner-local route authorization boundaries without route authorization grants, route enablement, lookup enablement, persistence, execution, or side effects, the `v0.2.369` production receipt notification action dry-run result lookup consumer redaction audit for KDE and Runtime consumer redaction boundaries without consumer authorization, consumer enablement, lookup route enablement, persistence, execution, raw result exposure, or side effects, the `v0.2.370` production receipt notification action dry-run result lookup consumer enablement gate audit for a fail-closed gate that combines route authorization and consumer redaction before any consumer can be enabled while keeping authorization grants, consumer enablement, lookup route enablement, persistence, execution, raw result exposure, and side effects disabled, and the `v0.2.371` production receipt notification action dry-run result lookup consumer enablement authorization receipt audit for an opaque receipt boundary over opaque result identifiers without receipt presence, receipt acceptance, consumer authorization, consumer enablement, lookup route enablement, persistence, execution, raw result exposure, or side effects. Do not dispatch those again unless their local branches are discarded or a reviewer explicitly asks for a repair branch.

Do not modify `docs/claude-code-implementation-packages.md` from any task listed here.

## Good Claude Code Fit

A task is suitable for Claude Code now if it is:

- Bounded to one Runtime or KDE-facing product lane.
- Mostly Go Runtime read-model work, with Ruby only for tests, reports, or smoke scripts.
- Reviewable offline with local fixtures.
- Useful to the KDE-first Windows application compatibility mainline.
- Safe without Docker, QEMU, network fetch, package-manager calls, real Portal transport, compatibility backend launch, privileged containers, broad host mounts, or host-root mutation.
- Mergeable one branch at a time through `scripts/mainline_integration_review.rb`.

## Recommended Dispatch Order

| Order | Task | Source document | Suggested branch | Why it is suitable now |
| --- | --- | --- | --- | --- |
| 1 | Production receipt notification action dry-run result lookup consumer enablement receipt writer authorization audit | `docs/windows-app-compatibility-implementation-brief.md` | `codex/production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit` | Model the writer authorization boundary required before any consumer enablement authorization receipt can be written while keeping receipt writes, receipt acceptance, persistence, replay, consumer authorization, route enablement, lookup enablement, lookup persistence, result persistence, dry-run execution, dispatch, request creation, actions, writes, production ownership, and side effects disabled. |

## Best First Pick

Start with production receipt notification action dry-run result lookup consumer enablement receipt writer authorization audit.

The v0.2.320 offline and constrained-host gates pass, including the full Buildroot build and QEMU serial smoke. Authorized q3 and q4 composes produced the Fedora 44 container, clean qcow2, persisted KVM graphical-login evidence, and closed host-boundary proof. The v0.2.321 and v0.2.322 matrix evidence can now consume controlled artifact and snapshot receipts, v0.2.323 turns production Runtime service activation into a fail-closed preflight, v0.2.324 materializes only a test-only review plan from restricted launch evidence, v0.2.325 defines the no-op backend adapter boundary before real Wine/Proton/VM launch paths are allowed, v0.2.326 records restricted owner smoke evidence without claiming production ownership, v0.2.327 fans the materialized review plan out to KDE-facing surfaces without enabling launch, v0.2.328 lets fixture-backed readiness rows consume no-op adapter contract profiles without invoking a backend, v0.2.329 lets readiness and support surfaces consume the restricted owner smoke receipt without claiming production ownership, v0.2.330 audits that the materialization fan-out should remain CLI-only until test receipt creation is split from read-only receipt consumption, v0.2.331 audits that the adapter contract should remain fixture-local until a redacted owner-local adapter profile audit route exists, v0.2.332 audits that the restricted owner smoke receipt fan-out should remain CLI-only until owner-managed opaque receipt lookup exists, v0.2.333 splits materialization receipt creation from read-only fan-out consumption, v0.2.334 creates the redacted adapter profile owner route while keeping the full adapter contract fixture-local, v0.2.335 introduces owner-managed opaque lookup for the restricted owner smoke receipt, v0.2.336 re-audits materialization fan-out to prove the read-only consume split exists while owner-managed opaque materialization receipt lookup is still missing, v0.2.337 routes restricted owner smoke fan-out through owner-managed opaque receipt lookup without caller state-root paths, v0.2.338 adds the matching owner-managed opaque lookup for materialization receipts, v0.2.339 lets materialization fan-out consume that owner lookup through an owner-local read route while staying non-launching and write-free, v0.2.340 proves that route appears in restricted owner smoke evidence through `Service.Call`, v0.2.341 gives the restricted owner smoke fan-out route the same explicit smoke coverage, v0.2.342 covers the redacted adapter profile route in restricted owner smoke evidence while keeping the full adapter contract fixture-local, v0.2.343 refreshes the materialization audit so it recognizes smoke coverage, v0.2.344 refreshes the restricted owner smoke fan-out audit so it recognizes smoke coverage, v0.2.345 refreshes the redacted adapter profile audit so it recognizes smoke coverage, v0.2.346 adds the read-only production D-Bus gate review packet, v0.2.347 adds the human authorization preflight consumed by that gate, v0.2.348 lets production service activation preflight consume both gates while still refusing service start and bus ownership, v0.2.349 lets the Runtime write gate consume the same production gate state while keeping write methods disabled, v0.2.350 inventories route-by-route production D-Bus method exposure decisions while registering no production methods, v0.2.351 adds rollback and diagnostics review before any production ownership commit, v0.2.352 adds desktop side-effect review for all seven KDE first-release surfaces, v0.2.353 consolidates the remaining human authorization receipt boundary without granting production authorization, v0.2.354 audits that each production gate consumes that consolidated boundary while still refusing ownership and side effects, v0.2.355 models future receipt acceptance propagation without writing or accepting a real receipt, v0.2.356 reviews the separate receipt writer authorization boundary while keeping persistence and replay disabled, v0.2.357 models persistence, expiry, revocation, replay, and audit-visibility threats without enabling storage, v0.2.358 models revocation and expiry visibility for production gates and KDE-safe status views without enabling revocation writes or notifications, v0.2.359 models notification delivery gating without enabling notification delivery, v0.2.360 models notification action safety without enabling notification actions or requests, v0.2.361 maps those actions to future Runtime request-object kinds without creating request objects, v0.2.362 models dispatch authorization without granting or executing dispatch, v0.2.363 models dispatch dry-run without execution, v0.2.364 models dry-run result visibility for KDE and Runtime diagnostics without result persistence, v0.2.365 models dry-run result persistence authorization without granting persistence, v0.2.366 models retention and redaction policy without enforcement, v0.2.367 models owner-managed opaque lookup without enabling lookup routes or persistence, v0.2.368 models lookup route authorization without granting or enabling routes, v0.2.369 models KDE and Runtime consumer redaction without enabling consumers or exposing raw results, v0.2.370 models the fail-closed consumer enablement gate without enabling consumers, and v0.2.371 models the opaque authorization receipt boundary without accepting or writing a receipt. The next useful blocker is a separate receipt writer authorization audit before any consumer enablement authorization receipt can be written. Windows application execution remains unproven.

The next branch must stay non-launching and must not claim the production bus, enable write methods, write KDE configuration, start a backend, call a real Portal, run QEMU, or mutate the host root.

## Good Branch Shape

Each Claude branch should:

- Implement exactly one listed task.
- Add a Go read model and a CLI preview command when the source task asks for one.
- Add success, malformed-input, blocked-state, redaction, and no-side-effect tests.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for code changes.
- Wire new files into `scripts/verify_layout.rb`.
- Wire new evidence into `scripts/implementation_evidence_report.rb` and `scripts/mainline_integration_review.rb` when the task creates a new product surface.
- Report exact verification commands run and exact verification commands skipped.

## Do Not Dispatch Now

Avoid these prompts:

- "Finish Windows compatibility."
- "Make KDE integration real."
- "Enable Wine launch."
- "Install Proton or Wine."
- "Run QEMU acceptance."
- "Wire production D-Bus."
- "Call the real Portal."
- "Add real cleanup."
- "Delete stale state."
- "Renew real permissions."
- "Revoke real permissions."
- "Send notifications."
- "Auto-fix permissions."
- "Download packages."
- Any task that would require editing `docs/claude-code-implementation-packages.md`.

## Copyable Prompt Prefix

Use this prefix before the selected task prompt:

```text
You are implementing one bounded Xnix task from docs/claude-code-current-dispatch-picks.md.

Use the referenced source task document as the authority for detailed requirements.

Hard boundaries:
- Implement exactly one task.
- Do not modify docs/claude-code-implementation-packages.md.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Prefer Go for durable Runtime product behavior.
- Use Ruby only for tests, smoke scripts, reports, and lightweight developer tooling.
- Keep KDE as shell and presentation only.
- Keep the Xnix Compatibility Runtime as the policy and product-logic owner.
- Do not enable production D-Bus ownership, Runtime writes, real Wine/Proton/VM launch, real Portal calls, network fetch, host package-manager calls, privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation.
- Do not expose raw executable paths, Runtime state-root paths, host paths, file contents, secrets, tokens, credentials, usernames, environment variables, raw backend commands, backend names, or backend details in KDE-facing output.

Before finishing:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md if code changes.
- Run the task-specific verification commands.
- Run ruby scripts/verify_layout.rb.
- Run git diff --check.
- Report exact commands run and any commands not run.
```

## Intake Reminder

When a Claude branch returns:

1. Run `ruby scripts/mainline_integration_review.rb --format json`.
2. Confirm `protected_claude_file_modified` is `false`.
3. Confirm `unclassified_file_count` is `0`.
4. Confirm the branch only touches the expected lane plus required shared plumbing and version metadata.
5. Run the task-specific tests listed in the source task document.
6. Exclude `.gocache/`, `.cache/`, `tmp/`, logs, build artifacts, and local config from staging.
7. Stage by lane, never with `git add .`.

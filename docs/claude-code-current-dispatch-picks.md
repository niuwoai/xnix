# Claude Code Current Dispatch Picks

> Last updated: 2026-07-17 | Baseline: v0.2.318

This is the current short dispatch sheet for Claude Code. It is intentionally smaller than the full wave documents and should be used when choosing the next independent branch.

Use `docs/claude-code-stability-release-train.md` as the cadence authority for the next Go-first stabilization train: one coherent small version per local commit, targeted tests on every small version, and full build plus remote push only at each twentieth-version train gate.

Local work already contains implementation evidence for all seventh-wave tasks (`S7W1` through `S7W8`) plus `S6W1`, `S6W2`, `S6W3`, `S6W4`, `S6W5`, `S6W6`, `S6W7`, `S6W8`, and `F5W7`. Do not dispatch those again unless their local branches are discarded or a reviewer explicitly asks for a repair branch.

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
| 1 | Restricted product-image smoke checkpoint | `docs/claude-code-stability-release-train.md` | `codex/stability-train-v0.2.319` | Join the restricted launch preflight packet to fixed product-image and smoke metadata without running Docker, QEMU, or a compatibility backend. |

## Best First Pick

Start with the restricted product-image smoke checkpoint.

The v0.2.318 checkpoint persists and verifies a fail-closed restricted launch preflight packet, names exactly the production recipe-trust and Runtime write-gate blockers, and permits only evidence-packet assembly. The next handoff should join that packet to fixed product-image and restricted smoke metadata without claiming that Docker, QEMU, KDE, Wine, Proton, a VM, or a Windows application was launched.

The offline identity must stay non-launching and must not write KDE configuration, start a backend, call a real Portal, or mutate the host root.

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

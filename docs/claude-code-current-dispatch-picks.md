# Claude Code Current Dispatch Picks

> Last updated: 2026-07-16 | Baseline: v0.2.294

This is the current short dispatch sheet for Claude Code. It is intentionally smaller than the full wave documents and should be used when choosing the next independent branch.

Local work already contains implementation evidence for `S7W1`, `S7W2`, and `S7W8`. Do not dispatch those again unless their local branches are discarded or a reviewer explicitly asks for a repair branch.

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
| 1 | `S7W3` State-root quota and retention preview | `docs/claude-code-seventh-wave-task-batch.md` | `codex/s7w3-state-root-retention-preview` | State-root, snapshot, diagnostic, execution, and Portal receipt models now exist. A dry-run storage pressure view is useful and can stay strictly read-only. |
| 2 | `S7W4` Crash and hang signal summary preview | `docs/claude-code-seventh-wave-task-batch.md` | `codex/s7w4-crash-hang-summary` | Diagnostic history, test results, repair plans, and support bundle evidence exist. Claude can add a privacy-safe failure summary without reading private logs. |
| 3 | `S7W7` Permission evidence audit preview | `docs/claude-code-seventh-wave-task-batch.md` | `codex/s7w7-permission-evidence-audit` | Portal, settings, execution, and desktop-safety evidence exist but still need one offline audit surface before renewal or revocation work becomes real. |
| 4 | `S7W5` Compatibility backend fallback preview | `docs/claude-code-seventh-wave-task-batch.md` | `codex/s7w5-backend-fallback-preview` | Runtime backend readiness, lifecycle, safety policy, and explanation cards are in place. Automatic-mode fallback can be explained without starting Wine, Proton, or a VM. |
| 5 | `S7W6` KDE search visibility plan | `docs/claude-code-seventh-wave-task-batch.md` | `codex/s7w6-kde-search-visibility` | KDE launcher, KRunner, Dolphin, file association, and Compatibility Center surfaces need one Runtime-owned visibility model before UI integration grows. |
| 6 | `S6W2` Desktop deactivation dry-run plan | `docs/claude-code-sixth-wave-task-batch.md` | `codex/s6w2-desktop-deactivation-dry-run` | Activation staging and rollback receipts exist. The matching safe removal story is valuable before real activation advances. |
| 7 | `S6W3` Recipe conflict and pin audit | `docs/claude-code-sixth-wave-task-batch.md` | `codex/s6w3-recipe-conflict-audit` | Recipe trust, upgrade impact, package-source, and install-planning lanes are active. Conflict evidence can prevent later install and upgrade ambiguity. |
| 8 | `S6W5` Snapshot restore candidate ranking | `docs/claude-code-sixth-wave-task-batch.md` | `codex/s6w5-snapshot-restore-ranking` | Snapshot and rollback records exist. A ranked restore preview can be implemented without executing restore or modifying application state. |

## Best First Pick

Start with `S7W3`.

`S7W3` is the best next handoff because it turns existing state-root and receipt evidence into a user-visible storage and retention explanation. It is also low-risk: the required output is a dry-run preview, and the branch must prove it never deletes files, creates directories, truncates logs, rewrites receipts, exposes state-root paths, or mutates the host.

If `S7W3` is already assigned, dispatch `S7W4` next. It builds directly on diagnostic evidence and improves the support story without enabling AI calls, backend launch, private log reads, or repair execution.

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

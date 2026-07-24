# v0.2.640 Post-Checkpoint Promotion Checklist

> Last updated: 2026-07-24 | Active candidate: v0.2.640-rc15 | Target formal release: v0.2.640

This checklist is the human-operator runbook for promoting the active `v0.2.640` release candidate to the formal `v0.2.640` tag.

Promotion is forbidden until `ruby scripts/full_smoke.rb` passes and the generated promotion packet allows the release. This document does not claim that full smoke has passed.

## Ownership Boundary

Codex and Claude Code may update review-only tooling and documentation, but the final promotion is operator-owned.

Claude Code must not:

- Run `ruby scripts/full_smoke.rb` without task-local human approval.
- Run Docker, QEMU, Wine, Proton, Colima, or backend execution without task-local human approval.
- Pull images, fetch packages, repair host networking, restart host services, or use host package managers.
- Use privileged containers, host networking, Docker socket mounts, broad host-directory mounts, or host-root mutation.
- Create, move, delete, or push the formal `v0.2.640` tag unless explicitly instructed by the human operator.
- Claim `formal_release_ready=true` from targeted tests, historical product smoke evidence, fixture-only reports, or a missing promotion packet.
- Modify `docs/claude-code-implementation-packages.md`.

## Required PASS Evidence

The promotion review must verify all of these items from current evidence, not from memory:

| Evidence | Required proof | Review command or source |
| --- | --- | --- |
| Full smoke | `ruby scripts/full_smoke.rb` exits successfully. | Human-authorized terminal output. |
| Full-smoke JSON report | `output/full-smoke-report.json` is present, parseable, and has `formal_release_ready=true`. | `ruby scripts/full_checkpoint_promotion_packet.rb --format json` |
| Full-smoke Markdown report | `output/full-smoke-report.md` exists for reviewer-readable evidence. | `ruby scripts/full_checkpoint_promotion_packet.rb --format markdown` |
| Restricted container lanes | Report evidence shows restricted container execution without privileged containers, Docker socket mounts, host networking, broad host mounts, or host-root mutation. | `output/full-smoke-report.json` and promotion packet flags. |
| QEMU serial smoke | Report evidence shows `qemu_booted=true`, persisted serial log, and no missing boot markers. | `output/full-smoke-report.json`, `output/serial.log` |
| `sshd` baseline | Full smoke includes the boot and SSH test baseline before Windows app lanes. | `ruby scripts/full_smoke.rb` output and report steps. |
| Known Windows app smoke | Report evidence shows `known_app_smoke_passed=true` for the staged known app lane. | `output/full-smoke-report.json` |
| Fixture Windows app smoke | Report evidence shows `fixture_app_smoke_passed=true`. | `output/full-smoke-report.json` |
| KDE controlled-launch action smoke | Report evidence shows `kde_action_smoke_passed=true`. | `output/full-smoke-report.json` |
| D-Bus fixture smoke | Full smoke includes the `kde-controlled-launch-action-dbus-fixture-smoke` lane. | `output/full-smoke-report.json` steps. |
| Promotion packet | `promotion_allowed=true`, `promotion_decision=promote`, and `formal_release_ready=true`. | `ruby scripts/full_checkpoint_promotion_packet.rb --format json` |
| Release evidence index | `full-checkpoint-promotion` and `product-image-qemu-acceptance` are not blocked by `full-checkpoint-promotion-not-allowed`. | `ruby scripts/release_evidence_index.rb --format json` |
| Merge readiness | `merge_ready=true`, no protected Claude file changes, no unclassified files, and no unsafe-operation findings. | `ruby scripts/merge_readiness_packet.rb --format json --offline-only` |
| Layout verification | Layout is consistent for the candidate version. | `ruby scripts/verify_layout.rb` |
| Mainline review | Current changes are classified and safe to stage by lane. | `ruby scripts/mainline_integration_review.rb --format json` |
| Diff hygiene | No whitespace errors and no protected Claude implementation package diff. | `git diff --check`; `git diff -- docs/claude-code-implementation-packages.md` |

If any row fails, stop. Do not create the formal tag.

## Exact Promotion Procedure

1. Confirm the worktree is clean before promotion work:

   ```text
   git status --short --branch
   ```

2. Confirm the current candidate version is the active `v0.2.640` release-candidate line:

   ```text
   ruby scripts/verify_layout.rb
   ruby scripts/full_checkpoint_promotion_packet.rb --format json
   ```

   Before the human-authorized full smoke passes, the expected decision is a blocker such as `blocked-incomplete-full-smoke-report`.

3. The human operator explicitly authorizes and runs the full checkpoint:

   ```text
   ruby scripts/full_smoke.rb
   ```

4. If full smoke fails:

   - Do not tag.
   - Read `output/full-smoke-report.json` and `output/full-smoke-report.md`.
   - Preserve `output/serial.log`.
   - Fix only the concrete project defect or operator-owned environment blocker shown by the report.
   - Rerun the complete full smoke after the fix.

5. After full smoke passes, run the promotion evidence chain:

   ```text
   ruby scripts/full_checkpoint_promotion_packet.rb --format json
   ruby scripts/full_checkpoint_promotion_packet.rb --format markdown
   ruby scripts/release_evidence_index.rb --format json
   ruby scripts/merge_readiness_packet.rb --format json --offline-only
   ruby scripts/mainline_integration_review.rb --format json
   ruby scripts/verify_layout.rb
   git diff --check
   git diff -- docs/claude-code-implementation-packages.md
   ```

6. Confirm the required release gate values:

   - Promotion packet:
     - `promotion_allowed=true`
     - `promotion_decision=promote`
     - `formal_release_ready=true`
     - `full_smoke_executed_by_packet=false`
     - `docker_executed_by_packet=false`
     - `qemu_executed_by_packet=false`
     - `wine_executed_by_packet=false`
     - `host_root_modified=false`
   - Release evidence index:
     - `full-checkpoint-promotion` is `implemented`.
     - `product-image-qemu-acceptance` is not blocked by `full-checkpoint-promotion-not-allowed`.
     - `release_ready=false` is still acceptable because production Windows execution remains separately gated.
   - Merge readiness:
     - `merge_ready=true`
     - `protected_file_status.status=clean`
     - `unsafe_operation_status.detected_unsafe_findings=[]`

7. Prepare the formal release commit:

   - Change `VERSION` from the active release candidate to `0.2.640`.
   - Change `scripts/verify_layout.rb` expected version to `0.2.640`.
   - Change `runtime/core/xnix_runtime_core.h` to `0.2.640`.
   - Change KDE metadata version to `0.2.640`.
   - Update `CHANGELOG.md` with a formal `0.2.640` entry.
   - Update `PRODUCT_OVERVIEW.md`, `README.md`, `CLAUDE.md`, and `docs/xnix-current-mainline.md`.

8. Run release-candidate verification again after the version bump:

   ```text
   ruby scripts/full_checkpoint_promotion_packet.rb --format json
   ruby scripts/release_evidence_index.rb --format json
   ruby scripts/merge_readiness_packet.rb --format json --offline-only
   ruby scripts/verify_layout.rb
   git diff --check
   git diff -- docs/claude-code-implementation-packages.md
   ```

9. Commit the formal release version bump:

   ```text
   git add CHANGELOG.md CLAUDE.md PRODUCT_OVERVIEW.md README.md VERSION docs/xnix-current-mainline.md kde/plasmoids/org.xnix.compatibilitycenter/metadata.json runtime/core/xnix_runtime_core.h scripts/verify_layout.rb
   git commit -m "chore(release): promote v0.2.640"
   ```

10. Create the formal annotated tag only after the commit succeeds:

    ```text
    git tag -a v0.2.640 -m "Release v0.2.640"
    ```

11. Push only when the human operator explicitly requests publication:

    ```text
    git push
    git push origin v0.2.640
    ```

## Rollback Notes

If promotion review fails before the formal commit:

- Do not tag.
- Keep the generated `output/full-smoke-report.json`, `output/full-smoke-report.md`, and `output/serial.log` for diagnosis.
- Revert only the failed promotion preparation changes, not unrelated user or Claude Code work.
- Return to the latest release-candidate tag and continue with a new `v0.2.640-rcN` repair version.

If the formal commit was created but the tag was not:

- Add a new release-candidate repair commit instead of rewriting published history.
- Do not reuse `v0.2.640` until the evidence chain passes again.

If the formal tag was created locally by mistake:

```text
git tag -d v0.2.640
```

Only delete a remote tag when the human operator explicitly authorizes it.

## Current Local Status

As of this checklist, `v0.2.640` is not promoted. The current local promotion packet is expected to remain blocked until a human-authorized full smoke produces current PASS evidence.

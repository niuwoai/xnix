# KDE-First Compatibility Current Gap Audit

> Last updated: 2026-07-15 | Evidence command: `ruby scripts/kde_first_presence_smoke.rb`

This audit records the current evidence for the KDE-first Windows application compatibility goal. It is intentionally scoped to current repository evidence and does not claim production readiness.

The product direction is clear: KDE Plasma is the first official shell; the Xnix Compatibility Runtime owns compatibility policy; Go should own durable Runtime product logic; C remains valid for low-level and already implemented Runtime policy surfaces; Ruby should stay focused on tests and development tooling.

## Current Evidence Summary

| Area | Current evidence | Status |
| --- | --- | --- |
| Official desktop strategy | Product documentation and Runtime docs name KDE Plasma as the flagship shell and GNOME/XFCE as later shells | Aligned |
| Runtime boundary | `runtime/dbus/org.xnix.Compatibility1.xml` exposes read-only planning methods and gated write methods | Aligned |
| KDE seven entry points | Runtime docs and models cover launcher, task manager, file manager, tray, notifications, Compatibility Center, and settings | Contracted |
| Go migration | Runtime owner route manifest reports 57 read-only routes, 40 Go-routed, 17 C-backed, and 0 Ruby legacy | In progress |
| Production Runtime owner | Owner readiness, service binding, live owner gate, process preview, and smoke plan exist, but production bus ownership remains disabled | Pending |
| Package acquisition | Package source, artifact manifest, acquisition preflight, and install plan contracts exist; current route evidence shows package source, acquisition preflight, and artifact manifest are Go-routed | In progress |
| Backend lifecycle | Backend selection, environment, binding, capability matrix, and lifecycle previews exist; capability matrix and lifecycle are now Go-routed but remain non-executing | In progress |
| Portal permissions | Portal access policy and request plan contracts exist; request plan is Go-routed, access policy remains C-backed | Contracted |
| Snapshot and rollback | Snapshot plan and rollback concepts exist, but production snapshot creation and restore remain unavailable | Contracted |
| Execution | Launch intent, execution readiness, request, review, decision, preflight, resource grant, transaction, session, and session status previews exist; actual launch remains disabled | Contracted |
| AI diagnostics | AI diagnostic input, recommendation, and approval gate contracts exist; provider calls and automatic repair remain disabled | Contracted |
| KDE materialization | Activation bundle, staging, transaction, status, installer, and rollback models exist; writes must remain target-root scoped | In progress |

## Runtime Owner Route Baseline

The current route manifest reports:

- Total read-only routes: 57.
- Go-routed routes: 40.
- C Runtime-backed routes: 17.
- Ruby legacy routes: 0.
- Read-only method parity: pass.
- Go route coverage: pending.
- C core adapter boundary: pending.
- Write route gate: pass.
- Host safety boundary: pass.

This is strong evidence that the Ruby-to-Go/C migration has crossed an important threshold: there are no remaining legacy Ruby dispatch routes in the owner route manifest. The next useful migration work is not "remove Ruby everywhere"; it is to add Go owner handlers or Go owner adapter boundaries for the remaining C-backed read-only policy routes.

## Remaining C-Backed Read-Only Routes

These methods still need native Go owner routes or explicit Go owner adapter boundaries:

- `GetDesktopActivationManifest`
- `GetTaskManagerIdentityPlan`
- `GetKDEIntegrationStatus`
- `GetKDEShellIntegrationPlan`
- `GetKDEApplicationSurfacePlan`
- `GetKWinWindowRulePlan`
- `GetApplicationStateRoot`
- `GetCompatibilityInstallPlan`
- `GetRepairPlan`
- `GetTestPlan`
- `GetTestResult`
- `GetAIDiagnosticInput`
- `GetAIDiagnosticRecommendation`
- `GetAIRepairApprovalGate`
- `GetSnapshotPlan`
- `GetPortalAccessPolicy`
- `GetRuntimeWriteGate`

Recommended migration waves:

1. KDE surface wave: `GetKDEIntegrationStatus`, `GetKDEShellIntegrationPlan`, `GetKDEApplicationSurfacePlan`, `GetTaskManagerIdentityPlan`, `GetKWinWindowRulePlan`.
2. Safety gate wave: `GetRuntimeWriteGate`, `GetPortalAccessPolicy`, `GetSnapshotPlan`, `GetApplicationStateRoot`.
3. Compatibility validation wave: `GetTestPlan`, `GetTestResult`, `GetRepairPlan`.
4. AI safety wave: `GetAIDiagnosticInput`, `GetAIDiagnosticRecommendation`, `GetAIRepairApprovalGate`.
5. Install wave: `GetDesktopActivationManifest`, `GetCompatibilityInstallPlan`.

## First Milestone Gap

Target: safe KDE-native application presence.

Current evidence supports the shape of this milestone:

- KDE-first product direction exists.
- The Runtime contract exposes the KDE-facing read-only methods.
- The seven KDE entry points are represented in Runtime docs and models.
- Standard desktop identity, desktop entries, MIME associations, tray, notifications, Compatibility Center, and settings previews exist.
- Write methods stay disabled.
- Backend execution remains blocked.

Remaining gaps:

- Several first-milestone KDE surface methods still route through C policy records rather than Go owner routes.
- Production owner smoke is still a plan, not a passing production owner.
- KDE materialization is still mostly preview/staging oriented.
- The evidence is now aggregated by `scripts/kde_first_presence_smoke.rb`, but that smoke should continue growing with future real state evidence.

Recommended next implementation:

- Add the next C-backed KDE surface owner route migration wave, then keep `scripts/kde_first_presence_smoke.rb` passing as the first-milestone regression proof.

## Second Milestone Gap

Target: safe preparation for execution.

Current evidence supports the contracts:

- Recipe trust and install gates exist.
- Package source, artifact manifest, acquisition preflight, and install plans exist.
- Backend selection, environment, binding, lifecycle, and capability matrix models exist.
- Portal policy and request planning exist.
- Snapshot, repair, execution preflight, and transaction models exist.

Remaining gaps:

- Production-shaped recipe signatures are still not proven.
- Artifact acquisition is not yet a real verified local staging pipeline.
- Backend lifecycle is still mostly policy/status modeling rather than a persistent state machine.
- Portal request planning does not yet prove request object lifecycle through a fake broker.
- Snapshot planning does not yet prove a real test-root snapshot store.
- Execution transactions remain blocked-by-default, which is correct, but they need stronger evidence from backend lifecycle, Portal, and snapshot state.

Recommended next implementation:

- Build the fake/test mode foundations in this order: recipe trust evidence, local artifact staging, backend lifecycle state root, fake Portal request broker, test-root snapshot store, then execution transaction readiness.

## Third Milestone Gap

Target: gated real compatibility execution.

Current evidence does not prove this milestone yet.

Known missing pieces:

- A production-shaped long-running Go Runtime owner.
- Stable D-Bus bus-name ownership in constrained smoke.
- Real backend lifecycle start in an isolated test environment.
- Launch permission only after recipe trust, package readiness, Portal permission, snapshot baseline, and user review.
- Observable live session status.
- Repair and rollback before risky changes.

Do not start real compatibility execution until the second milestone evidence is strong. This is not timidity; this is the difference between a product and a delightful foot-gun wearing a hat.

## Highest-Value Next Work

The next work should prefer concrete evidence over more contracts.

Recommended order:

1. Finish a small Go owner route migration wave for C-backed KDE surface methods.
2. Keep the first-milestone KDE presence smoke passing as new route migrations land.
3. Add a persistent test-root state model for backend lifecycle.
4. Add fake Portal request object lifecycle.
5. Add test-root snapshot create/list/verify/rollback.
6. Connect execution transaction readiness to the real lifecycle, Portal, and snapshot evidence.

## Evidence Caveats

- This audit was produced against a dirty worktree with parallel implementation changes present.
- The KDE-first presence smoke ran successfully and reported 57 total read-only routes, 40 Go-routed routes, 17 C-backed routes, and 0 Ruby legacy routes.
- Full `go test ./...`, constrained Docker smokes, and full smoke were not run for this audit.
- `docs/claude-code-implementation-packages.md` is intentionally excluded from this audit because another agent owns that workstream.

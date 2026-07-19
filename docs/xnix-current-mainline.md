# Xnix Current Mainline

> Last updated: 2026-07-19 | Baseline: v0.2.387

This is the Codex-owned current mainline for Xnix. It replaces ad-hoc external-agent dispatch as the default planning source for new implementation work. Historical `claude-code-*` documents remain repository evidence, but new mainline work should use this document unless a user explicitly asks for a different handoff artifact.

## Product Goal

Xnix aims to be the Linux desktop that best integrates existing Windows applications without making users learn compatibility internals. KDE Plasma is the only flagship desktop for the first product line. GNOME and XFCE can run later as alternate shells, but they are not first-release integration targets.

The product stack is:

- KDE Plasma as the replaceable official shell.
- KDE integration plugins for presentation and user interaction.
- Xnix AI Compatibility Runtime as an independent Go-owned system service with D-Bus APIs.
- Wine, Proton, and Windows VM backends behind Runtime-owned policies.
- Linux as the host system layer.

## Ownership Rules

- Runtime business logic is Go-first.
- Ruby is for tests, reports, fixtures, and low-frequency development tooling.
- C remains for low-level or already-owned Runtime policy surfaces.
- KDE plugins display state, request user interaction, and forward actions; they do not own compatibility decisions.
- Runtime owns recipes, backend selection, permissions, diagnostics, snapshots, rollback, launch planning, and future execution gates.
- The official desktop remains KDE Plasma; do not fork Plasma or KWin for core product behavior.
- Do not expose raw backend commands, host paths, state-root paths, or compatibility internals in normal desktop output.

## First-Release KDE Entrypoints

The first KDE release line must converge on these seven user-facing entrypoints:

1. Start menu integration for Linux and managed Windows applications.
2. Task manager identity for compatibility-backed application windows.
3. Dolphin file-manager actions for opening files with managed Windows applications and AI analysis.
4. System tray status for compatibility environments and tray-heavy applications.
5. Notification Center events for install failure, repair suggestions, environment switching, and approvals.
6. AI Compatibility Center pages for compatibility status, run mode, known issues, and repair history.
7. Unified settings that avoid user-facing compatibility jargon and expose safe choices such as run mode, performance/compatibility preference, document access, camera policy, network policy, and snapshots.

## Current Safe Next Task

The next Codex-owned mainline task is:

Production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed writer implementation preview.

This follows the redacted status writer enablement audit, the redacted status writer grant audit, the redacted status writer accepted receipt gate audit, the redacted status writer authorization receipt acceptance audit, the redacted status writer authorization receipt audit, and the redacted status writer authorization gate audit. Continuity evidence from the writer authorization gate audit must continue to consume the v0.2.380 redacted status persistence write-model audit and this Xnix current mainline document. The closed writer implementation preview must continue to consume the writer enablement audit and this Xnix current mainline document as evidence. The next closed writer implementation preview should model the callable writer shape while keeping real persistence writes, status persistence, consumer enablement, route enablement, lookup enablement, dry-run execution, dispatch, request creation, production ownership, backend launch, and host mutation disabled.

## Safety Gates

Every small version must keep these gates closed unless the user explicitly authorizes a broader milestone:

- No production D-Bus ownership.
- No Runtime write methods.
- No KDE configuration writes.
- No real Portal calls.
- No backend launch.
- No package manager or network fetch.
- No Docker, QEMU, or full build unless explicitly authorized.
- No privileged containers, broad host mounts, or host-root mutation.

At every twentieth small version, full build and QEMU or container smoke checks require explicit operator authorization before execution.

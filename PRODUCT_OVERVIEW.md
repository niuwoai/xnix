# Xnix Product Overview

> Last updated: 2026-07-13 | Current version: v0.2.39

## Summary

Xnix is an atomic Linux desktop designed to make existing Windows applications feel native. KDE Plasma provides the familiar desktop shell; the independent Xnix Compatibility Runtime owns backend selection, recipes, diagnostics, permissions, snapshots, and rollback. The existing Buildroot/QEMU image remains a learning and low-level verification baseline, not the flagship product base.

## Core Goals

- Ship one official flagship desktop: KDE Plasma 6.
- Make Windows applications appear as ordinary Linux applications in the launcher, task manager, file associations, notifications, and settings.
- Keep Wine/Proton and Windows VM management behind an independent D-Bus Runtime.
- Use application recipes for backend selection, dependencies, diagnostics, permissions, snapshots, and rollback.
- Use XDG Desktop Portal for user-controlled file, URI, print, clipboard, screen, camera, and remote-desktop access.
- Preserve reproducible low-level QEMU verification and constrained Colima development infrastructure.

## Non-Goals for This Phase

- Forking KDE Plasma, KWin, Wine, Proton, or a Linux distribution.
- Supporting GNOME or XFCE as an officially integrated desktop in the first product release.
- Exposing Wine prefixes, raw executable paths, or backend implementation terminology in normal user entry points.
- Adding localization or non-English project-facing content during the current phase.

## Technology Choices

| Layer | Initial choice | Purpose |
| --- | --- | --- |
| Atomic base | Fedora Kinoite-compatible image | KDE Plasma desktop and transactional operating-system rollback |
| Desktop shell | KDE Plasma 6 | Familiar launcher, task manager, tray, desktop icons, file manager, and multi-window behavior |
| Runtime | Xnix Compatibility Runtime | Independent D-Bus service for recipes, Wine/VM backends, permissions, diagnostics, snapshots, and rollback |
| Compatibility backends | Wine/Proton and Windows VM | Selectable implementation behind a stable Runtime API |
| Permissions | XDG Desktop Portal | User-mediated file, URI, print, clipboard, screen, and remote-desktop access |
| Desktop integration | Plasmoid, KRunner, KWin, Dolphin, and system services | Presentation and interaction only; no backend policy in KDE code |
| Learning baseline | Buildroot plus QEMU | Reproducible boot, initramfs, networking, and SSH verification |
| Container host | Colima Docker | Constrained developer build and test isolation |

## Compatibility Runtime Boundary

- The Runtime contract reserves `org.xnix.Compatibility1` for application discovery, recipe installation, launch, diagnostics, snapshot, and restore operations. The local daemon core is executable and can probe itself, list managed recipes, report diagnostics with compatibility test plan, test result, AI diagnostic input, AI diagnostic recommendation, and AI repair approval gate summaries, dispatch read-only Runtime methods, stage activation files under an unprivileged root, and emit a KDE-safe read model for the Compatibility Center. The D-Bus session smoke adapter owns the Runtime bus name and answers read-only calls in the constrained Linux container; the KDE read model and D-Bus client can consume that session-bus source for planning reads. The production daemon binding is the next implementation step.
- The Runtime D-Bus contract now exposes read-only planning methods for engine catalog, run plans, repair plans, test plans, test results, AI diagnostic inputs, AI diagnostic recommendations, AI repair approval gates, snapshot plans, and Portal access policy. `DBusRuntimeClient` wraps those reads for KDE-facing code and parses D-Bus boolean variants into native booleans. Write operations such as launch and restore remain unsupported until production backends exist.
- Recipe registry verification delegates to `xnix-recipe-registry`, which validates registry schema version, recipe identifiers, relative recipe paths, SHA-256 digests, and signature status without claiming production signed-recipe validation yet. The Runtime daemon uses registry-backed recipe loading when `registry.json` is present, so the default application list is digest-verified before exposure. The Runtime probe reports recipe trust status for KDE surfaces and diagnostics, while `xnix-recipe-trust-policy` evaluates whether the registry is production-trusted, development-only, or untrusted. Recipe install gates delegate to `xnix-recipe-install-gate`, which blocks production activation of development-only registries while still allowing digest-verified development staging.
- KDE packages consume Runtime state and submit user actions over D-Bus. They do not invoke Wine or a VM directly.
- Compatibility engine cataloging delegates to `xnix-compat-engine-catalog`, which keeps automatic, local, and isolated engine selection in the Runtime and reports that backend launch bindings are still pending.
- Compatibility run planning delegates to `xnix-compat-run-plan`, which maps recipe intent to automatic, local, or isolated execution strategies, keeps backend details out of desktop-facing JSON, and reports that launch backend binding is still pending.
- Compatibility repair planning delegates to `xnix-compat-repair-plan`, which turns diagnostics issues into approval, snapshot, rollback, and notification-ready repair actions for the Compatibility Center.
- Compatibility snapshot planning delegates to `xnix-compat-snapshot-plan`, which scopes restore points to application state, Runtime metadata, and desktop activation receipts while excluding user documents and the host system.
- Compatibility test planning delegates to `xnix-compat-test-plan`, which combines recipe validation, Portal preflight, snapshot preflight, and managed launch binding checks into Runtime-owned preflight, smoke, and repair-readiness plans for the Compatibility Center.
- Compatibility test results delegate to `xnix-compat-test-result`, which converts plan steps into desktop-safe result summaries for AI diagnostics and Compatibility Center cards without claiming backend execution before launch binding exists.
- AI diagnostic inputs delegate to `xnix-ai-diagnostic-input`, which packages recipe, run-plan, test-result, repair, signal, and privacy-boundary context without calling an AI provider or requiring network access.
- AI diagnostic recommendations delegate to `xnix-ai-diagnostic-recommendation`, which converts diagnostic inputs into user-visible, review-first recommendations without executing repairs or calling an AI provider.
- AI repair approval gates delegate to `xnix-ai-repair-approval-gate`, which blocks AI-driven repair execution until Compatibility Center review, Runtime approval, and restore-point preflight gates pass.
- Application recipes generate normal `.desktop` launchers that call `xnix-compat-launch --app <id>`. That entry point validates the Runtime application id, preserves optional file URIs, and models a `Launch` request without displaying a prefix path, a backend command, or a raw Windows executable command.
- File association generation delegates to `xnix-file-association-model`, which maps recipe MIME types to generated desktop files and standard `mimeapps.list` defaults inside a staging root without overwriting existing associations.
- Dolphin service menus delegate selected files to `xnix-compat-open`, which validates file URIs, resolves a Runtime application recipe by file extension, and models a portal-required `Launch` request without directly invoking Wine or a VM.
- Runtime events delegate to `xnix-compat-notify`, which models KDE notification payloads for install failures, automatic repairs, mode changes, and approval requests without exposing backend implementation terms.
- Compatibility settings delegate to `xnix-compat-settings`, which models user-facing controls for automatic mode, performance or compatibility priority, documents and downloads access, camera access, network access, and snapshots without backend terminology.
- System tray status delegates to `xnix-compat-tray-status`, which models Runtime activity, attention state, and bridged tray application counts for a KDE tray surface.
- Task manager identity delegates to `xnix-compat-window-identity`, which models desktop file mapping, grouping, pinning, restore behavior, and KWin identity-only metadata for compatibility windows.
- KWin window rules delegate to `xnix-kwin-window-rule`, which gives KWin scripts identity and layout hints for compatibility windows while leaving backend policy in the Runtime.
- Desktop integration manifests delegate to `xnix-desktop-integration-manifest`, which bundles a recipe's launcher, task manager, Dolphin, tray, notification, Compatibility Center, and settings artifacts into one activation plan without exposing backend commands.
- KRunner query integration delegates to `xnix-krunner-model`, which maps application names and file-oriented natural queries to Runtime application identities and managed launcher actions without exposing backend details.
- Desktop activation delegates to `xnix-install-desktop-integration`, which stages generated application launchers, Dolphin service menus, and the desktop integration manifest under a target root without modifying the host root. The CLI path now enforces recipe install-gate preflight, blocks development-only registries in production mode, and keeps explicit development staging available for digest-verified local tests.
- Desktop activation rollback delegates to `xnix-rollback-desktop-integration`, which reads activation receipts and removes only unchanged staged files after SHA-256 verification.
- Runtime operations that need desktop resources must use XDG Desktop Portal request/response flows; backend-specific direct access is not a public UI contract.
- Portal access policy delegates to `xnix-portal-access-policy`, which keeps file, URI, print, screenshot, clipboard, camera, and remote-desktop decisions in the Runtime while requiring user-mediated XDG Desktop Portal requests and denying direct desktop access.
- Portal request modeling delegates to `xnix-portal-request-model`, which describes the XDG Desktop Portal destination, interface, method, request handle, completion signal, denied-state guidance, and Runtime-owned result handling for sensitive desktop operations.
- `xnix-kde-integration-status` tracks the seven first-release KDE entry points and records which have an initial Runtime-backed integration versus planned work.

## Desktop Integration Scope

| Entry point | First implementation responsibility |
| --- | --- |
| Launcher | Generated standard desktop files call the managed Runtime launch request path |
| Task manager | Window identity and KWin rule models map compatibility windows to Runtime application identities for grouping, pinning, taskbar visibility, switcher visibility, and restore |
| File manager | Dolphin action and MIME association model map file types to Runtime-managed applications and require Portal-mediated access |
| System tray | Tray status model presents Runtime activity, attention state, and bridged tray applications |
| Notifications | Runtime events produce KDE notification request models for install, repair, mode-change, and approval events |
| Compatibility Center | Plasma package displays the Runtime-backed read model for compatibility, diagnostics, snapshots, and actions; the read path is smoke-tested over D-Bus |
| Settings | KDE settings model presents user concepts such as automatic mode, performance, compatibility, allowed resources, devices, network, and snapshots |
| KRunner query | KDE natural-language entry maps queries to Runtime application identities and managed launcher actions |
| Activation manifest | Runtime manifest groups all seven KDE artifacts for a recipe activation without exposing backend commands |
| Activation staging | Runtime installer writes desktop entry, MIME association, Dolphin service menu, and manifest files under a staging root after recipe install-gate preflight |
| Activation rollback | Runtime rollback removes only receipt-tracked, checksum-matching staged files |
| Recipe registry | Registry verifier checks recipe metadata, paths, digests, and development signature status |
| Registry-backed loading | Runtime loads registered recipes after digest verification and keeps no-registry fallback only for development fixtures |
| Recipe trust probe | Runtime probe reports registry-backed loading, digest verification, and signed-recipe validation status |
| Recipe trust policy | Policy model turns registry trust signals into production, development-only, or untrusted decisions |
| Recipe install gate | Install gate blocks production activation until recipe trust requirements are satisfied |
| Portal access policy | Runtime policy requires user-mediated XDG Desktop Portal requests for sensitive desktop operations |
| Portal request model | Runtime request model describes XDG Desktop Portal calls and completion handling without granting direct access |
| Compatibility engine catalog | Runtime owns automatic, local, and isolated engine choices without exposing backend details |
| Compatibility run plan | Runtime maps recipe intent to a desktop-safe execution strategy and marks backend binding as pending |
| Compatibility repair plan | Runtime maps diagnostic issues to approval, snapshot, rollback, and notification-ready repair actions |
| Compatibility snapshot plan | Runtime scopes restore points before repairs without touching user documents or the host system |
| Compatibility test plan | Runtime maps preflight, smoke, and repair-readiness checks to Compatibility Center summaries |
| Compatibility test result | Runtime summarizes passed, pending, and blocked compatibility test outcomes for AI diagnostics |
| AI diagnostic input | Runtime packages safe diagnostic context for future AI analysis without exposing user documents or backend details |
| AI diagnostic recommendation | Runtime prepares user-visible review-first recommendations without auto-executing repairs |
| AI repair approval gate | Runtime blocks AI repair execution until review, approval, and restore-point preflight gates pass |

## Learning Baseline

- Buildroot version: `2025.02.15` LTS.
- Buildroot source is intentionally not committed. `buildroot/sources.lock` pins its upstream URL and SHA-256; a build command retrieves it into an ignored cache directory.
- The external tree lives in `buildroot/` and provides the `xnix_x86_64_defconfig` configuration.

## Verified Runtime Baseline

- A complete constrained Buildroot build exits successfully and produces the x86_64 kernel and initramfs.
- The QEMU serial console shows Linux booting, running `/init`, and reaching the `xnix login:` prompt.
- The QEMU user-mode NIC obtains `10.0.2.15` through DHCP.
- OpenSSH `sshd` accepts a non-interactive key-authenticated connection through the container-only loopback forwarding rule.

## Test and Release Policy

- Each small code version receives focused unit tests and a Git commit.
- Every tenth code version requires a complete build and QEMU smoke test.
- If the full milestone test exposes a defect, it must be fixed and the full test rerun before later work starts.
- Colima is limited to 6 GiB memory and 1 CPU. Runtime containers receive explicit resource limits and no elevated privileges.
- Buildroot is configured with a single job to prevent toolchain build memory spikes in the 4 GiB container limit.
- Xnix runtime containers are limited to 4 GiB, one CPU, and 256 processes. They use a read-only root filesystem, a small temporary filesystem, no Linux capabilities, no host-directory mounts, and no network unless a dedicated source-retrieval or Buildroot dependency-download command requires it. These commands receive bridge networking and a Docker-managed named volume only; they never mount a host directory. Image builds run inside the 6 GiB and one-CPU Colima VM limit.
- SSH verification creates an ephemeral Ed25519 keypair only in the Docker-managed cache volume, installs its public key during image assembly, forwards only container loopback `127.0.0.1:2222` to guest port 22, and deletes the private key after the test.

## Milestones

- [x] B0: Reproducible constrained Buildroot/QEMU learning baseline.
- [x] B1: Kernel boot, initramfs, DHCP, and loopback `sshd` verification.
- [x] C0: Compatibility Runtime D-Bus contract, safe recipe model, desktop launcher generator, and Plasma package skeleton.
- [ ] C1: Activated Runtime implementation and signed recipe storage. The local daemon core, registry-backed recipe-store loader, recipe registry verifier, recipe trust probe, recipe trust policy, recipe install gate, Portal access policy, Portal request model, compatibility engine catalog, compatibility run plan, compatibility repair plan, compatibility snapshot plan, compatibility test plan, compatibility test result model, AI diagnostic input model, AI diagnostic recommendation model, AI repair approval gate model, read-only planning D-Bus methods, D-Bus planning client reads, install-gated desktop activation, read-only method dispatcher, packaged activation wrapper, activation-file installer, desktop activation installer, activation rollback receipt, Linux D-Bus session smoke adapter, KDE-safe Compatibility Center model, KDE D-Bus read smoke, KRunner query model, KWin window rule model, file association model, launcher request model, Dolphin file-open request model, notification request model, settings model, tray status model, task manager identity model, desktop integration manifest, and seven-entry-point KDE integration status are in place; production daemon binding and production signature validation are still pending.
- [ ] C2: Wine/Proton backend with snapshots, diagnostics, and Portal-mediated permissions.
- [ ] C3: Windows VM backend with file, clipboard, print, and window bridging.
- [ ] C4: KDE launcher, task manager, Dolphin, tray, notification, KRunner, KWin, Compatibility Center, and settings integrations.
- [ ] C5: Reproducible atomic KDE desktop image and graphical QEMU smoke test.

## Current Decisions

- KDE Plasma is the only official desktop for the first product release; GNOME and XFCE remain future optional shells.
- The flagship base is Fedora Kinoite-compatible because an atomic deployment model matches system and application rollback requirements.
- Buildroot is retained for kernel, boot, initramfs, and QEMU learning experiments, not for packaging the flagship desktop.
- The Runtime is the product core; desktop packages remain replaceable adapters.

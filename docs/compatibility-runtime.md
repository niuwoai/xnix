# Xnix Compatibility Runtime

## Purpose

The Xnix Compatibility Runtime is the product core. It selects and manages Wine/Proton and Windows VM backends while presenting applications as normal Linux desktop applications. KDE Plasma is the flagship shell, but the Runtime must not depend on Plasma implementation details.

## Base System Decision

The flagship image targets a Fedora Kinoite-compatible atomic KDE desktop instead of the Buildroot learning image. Kinoite supplies an immutable KDE Plasma desktop base, while an atomic deployment model supports a system rollback path. Buildroot remains useful for the low-level boot laboratory but is not a practical package base for the complete desktop compatibility stack.

## Component Boundaries

```text
KDE Plasma packages
  Plasmoid / KRunner / KWin / Dolphin / settings
                 |
                 | D-Bus
                 v
Xnix Compatibility Runtime
  recipes / diagnostics / permissions / snapshots / rollback
                 |
          Wine/Proton and Windows VM backends
                 |
       atomic Linux desktop base and XDG Desktop Portal
```

KDE packages display status and submit user decisions. They do not create a Wine environment, start a virtual machine, parse backend logs, or select a backend. The Runtime owns those decisions and will implement the stable interface defined in `runtime/dbus/org.xnix.Compatibility1.xml`.

The KDE Compatibility Center consumes a read-only presentation model from `xnix-kde-center-model`. That model prefers the Runtime D-Bus service when a session source is available and falls back to the local Runtime read model for offline development. In both cases, it filters backend storage paths and implementation terminology before anything reaches the Plasma shell.

The Runtime D-Bus contract exposes read-only planning methods for engine catalog, Compatibility Center action queues, Compatibility Center action review receipts, compatibility artifact manifests, compatibility install plans, compatibility acquisition preflight, compatibility package sources, application state roots, managed backend binding, compatibility settings, compatibility settings change plans, run plans, repair plans, test plans, test results, AI diagnostic inputs, AI diagnostic recommendations, AI repair approval gates, Runtime service binding status, Runtime live owner gates, Runtime owner smoke plans, Runtime method parity manifests, snapshot plans, and Portal access policy. `DBusRuntimeClient` wraps these methods for KDE-facing code and parses D-Bus boolean variants into native booleans. These methods give KDE surfaces a stable integration path without giving KDE ownership of backend decisions. Write methods remain asynchronous and unsupported until production backend binding exists.
The Runtime D-Bus contract also exposes Compatibility Center action queues, action review receipts, compatibility artifact manifests, compatibility install plans, compatibility acquisition preflight, compatibility package sources, application state roots, compatibility settings, compatibility settings change plans, compatibility test plans, test results, managed backend binding, AI diagnostic inputs, AI diagnostic recommendations, AI repair approval gates, Runtime service binding status, Runtime live owner gates, Runtime owner smoke plans, and Runtime method parity manifests as read-only planning data. KDE can show queued user-review cards, recorded review intent, artifact manifest readiness, install readiness, acquisition readiness, package source state, state ownership, settings state, pending settings changes, preflight, smoke, backend binding readiness, repair-readiness, AI-ready diagnostic summaries, review-first recommendations, blocked repair-gate summaries, activation-binding readiness, production owner transition gates, production owner smoke readiness, and read-only method parity in the Compatibility Center, but the Runtime owns action execution, artifact manifest resolution, install readiness gates, artifact acquisition, package source selection, state allocation, settings policy, settings persistence, recipe validation, Portal preflight, snapshot preflight, managed launch-binding checks, result status, AI diagnostic boundaries, recommendation safety policy, repair execution approval, and daemon ownership readiness.

Generated application launchers delegate to `xnix-compat-launch --app <id>`. That entry point validates the Runtime application id, accepts optional `file://` URIs from desktop file associations, and emits a Runtime `Launch` request model. Plain application launches do not need file portal access; file launches are marked as portal-mediated.

File association generation delegates to `xnix-file-association-model`. The model maps recipe MIME types to the generated desktop file and emits standard `mimeapps.list` content for a staging root. Activation refuses to overwrite an existing `mimeapps.list` until merge support exists, so local tests and future installers do not destroy user defaults.

The Dolphin service menu delegates selected files to `xnix-compat-open`. That entry point accepts only `file://` URIs, resolves a Runtime recipe by extension unless an explicit application id is supplied, and emits a portal-required Runtime `Launch` request model. It does not start Wine, a virtual machine, or a backend-specific executable.

`xnix-kde-integration-status` records the first-release KDE entry-point scope. Launcher, task manager, file manager, system tray, notifications, Compatibility Center, and settings are marked as initial because each entry point now has a Runtime-backed request, read-model, or identity model.

KRunner query integration delegates to `xnix-krunner-model`. The model resolves application names and file-oriented natural queries to Runtime application identities, standard desktop entry ids, and `xnix-compat-launch` actions. It is a presentation model only; matching results do not expose backend details and do not launch compatibility engines directly.

Runtime events delegate to `xnix-compat-notify`. That entry point models KDE notification payloads for install failures, automatic repairs, compatibility mode changes, and approval-required events. It chooses urgency and actions for the desktop shell, but it does not inspect backend logs or expose backend implementation terms.

Compatibility settings delegate to `xnix-compat-settings` and `GetCompatibilitySettings`. That model exposes Runtime-owned user-facing controls for run mode, performance or compatibility priority, documents and downloads access, camera access, network access, and snapshots. It intentionally avoids backend implementation terminology and does not let KDE own backend policy.

Compatibility settings change planning delegates to `xnix-compat-settings-change` and `GetCompatibilitySettingsChangePlan`. A setting change request validates the user-facing section, field, and value, then returns confirmation, Portal policy review, restore-point, and Runtime persistence gates. It is a planning model only: KDE must not persist compatibility settings, grant desktop resources, mutate the host root, or expose backend implementation settings while the Runtime reports `apply_enabled` as false.

Compatibility Center action queues delegate to `xnix-compat-action-queue` and `GetCompatibilityActionQueue`. The queue groups install readiness, settings changes, AI repair review, Runtime service ownership, and Portal policy review into KDE task cards. It is non-executing: KDE can present and collect review intent, but queued actions cannot start backends, persist settings, execute repairs, grant resources, or mutate the host root until Runtime gates are implemented.

Compatibility Center action review receipts delegate to `xnix-compat-action-review` and `GetCompatibilityActionReviewReceipt`. A receipt records KDE review intent for a queued action, but it is not an execution token. Review receipts cannot start backends, persist settings, execute repairs, grant resources, mutate the host root, or expose backend implementation details.

System tray status delegates to `xnix-compat-tray-status`. That model exposes Runtime activity, attention state, and bridged tray application counts for the KDE shell. It does not own backend policy or create backend tray bridges.

Task manager identity delegates to `xnix-compat-window-identity`. That model exposes desktop file mapping, Runtime application grouping, pinning, restore behavior, and KWin identity-only matching metadata. It does not decide backend policy or manage windows directly.

KWin window rules delegate to `xnix-kwin-window-rule`. The rule model gives a future KWin script enough information to match compatibility windows, attach generated desktop files, keep windows visible in the taskbar and switcher, and restore an existing window by Runtime application id. It is explicitly window-manager policy only; backend selection remains in the Runtime.

Desktop activation planning delegates to `xnix-desktop-integration-manifest`. That manifest groups a recipe's launcher, task manager, Dolphin, tray, notification, Compatibility Center, and settings artifacts into one KDE activation plan. It is ordered, portal-aware, does not require host privilege, and does not expose backend commands.

Desktop activation staging delegates to `xnix-install-desktop-integration`. That installer writes the generated application launcher, Dolphin service menu, and persisted desktop integration manifest under a supplied root after recipe install-gate preflight. It rejects host-root installation, blocks production activation of development-only registries, and gives tests an explicit development mode for verifying KDE desktop artifacts without modifying a developer workstation.

Desktop activation rollback delegates to `xnix-rollback-desktop-integration`. The installer writes a SHA-256 activation receipt, and rollback removes only receipt-tracked files whose current digest still matches the receipt. Changed files are preserved for diagnosis instead of being deleted.

## Application Recipes

A recipe has a stable reverse-DNS application identifier, display name, icon, requested run mode, and supported file extensions. The Runtime generates an ordinary `.desktop` file that calls `xnix-compat-launch --app <id>`. The desktop file intentionally omits backend commands, prefix paths, and Windows executable paths.

Recipe registry verification delegates to `xnix-recipe-registry`. The registry records schema version, recipe id, safe relative path, SHA-256 digest, and signature status for each recipe. The current development registry verifies digests and reports that production signed-recipe validation is still disabled; it does not yet trust or execute external recipe files.

Runtime recipe loading uses `RegistryBackedRecipeStore` when `registry.json` is present. That store verifies the registry before loading recipe files, so default application discovery is digest-checked. A no-registry fallback remains only for isolated development fixtures.

Runtime probes report recipe trust status. KDE surfaces and diagnostics can tell whether the application list came from registry-backed loading, whether digests were verified, and whether production signed-recipe validation is enabled.

Recipe trust policy delegates to `xnix-recipe-trust-policy`. The policy converts raw registry trust signals into desktop-safe decisions: production-trusted, development-only, or untrusted. Development registries remain blocked from production trust until signed recipe validation is enabled.

Recipe install gates delegate to `xnix-recipe-install-gate`. The gate evaluates a registry report, application id, and production or development mode before activation. Production installation blocks development-only registries until a production signed source and verified recipe signatures are available; development staging remains available for digest-verified local recipes.

Compatibility engine cataloging delegates to `xnix-compat-engine-catalog`. The catalog keeps engine selection in the Runtime, exposes automatic, local, and isolated choices, and reports pending readiness without revealing backend implementation details to KDE.

Compatibility artifact manifests delegate to `xnix-compat-artifact-manifest`. The manifest model plans signed manifests, artifact groups, digest verification, cache namespaces, and rollback references before artifact acquisition. It does not create network requests, download artifacts, install packages, expose cache paths to KDE, mutate the host root, require privileged containers, or reveal backend implementation details.

Compatibility install plans delegate to `xnix-compat-install-plan`. The install plan model joins artifact manifest, acquisition preflight, package source, state root, and recipe install gate readiness before desktop activation. It does not create network requests, download artifacts, install packages, stage desktop files, mutate the host root, require privileged containers, expose shell commands to KDE, or reveal backend implementation details.

Compatibility acquisition preflight delegates to `xnix-compat-acquisition-preflight`. The preflight model checks package-source readiness, signed artifact manifests, Runtime cache capacity, network policy review, and rollback markers before any acquisition is enabled. It does not create network requests, download artifacts, install packages, mutate the host root, require privileged containers, expose commands to KDE, or reveal backend implementation details.

Compatibility package source selection delegates to `xnix-compat-package-source`. The package-source model plans signed source policy, source channels, cache preflight, and offline fallback before launch binding. It does not install packages, invoke a host package manager, mutate the host root, require privileged containers, expose package manager commands to KDE, or reveal backend implementation details.

Managed backend binding delegates to `xnix-compat-backend-binding`. The binding model connects an application recipe to the Runtime-selected execution strategy, lists required preflight for package source selection, application state ownership, Portal policy, and baseline snapshots, and blocks unsafe launch actions until the Runtime owns the binding. It does not create an execution request, mutate the host root, require a privileged container, or expose implementation details.

Application state roots delegate to `xnix-compat-state-root`. The state-root model plans per-application Runtime state ownership, stable namespaces, managed scopes, snapshot eligibility, bounded restore-point retention, Portal-required user file access, and restore confirmation. It does not create directories, mutate the host root, include user documents in snapshots, expose host storage paths to KDE, or reveal backend implementation details.

Compatibility run planning delegates to `xnix-compat-run-plan`. The plan maps recipe intent to automatic, local, or isolated execution strategies, adds Portal and snapshot preflight requirements, and keeps desktop-facing output free of backend implementation details. Backend binding remains pending until the Runtime can safely launch real compatibility engines.

Compatibility repair planning delegates to `xnix-compat-repair-plan`. Diagnostics can attach repair summaries for the Compatibility Center, including approval requirements, snapshot requirements, rollback availability, and notification mapping. The plan is advisory until the Runtime has production repair execution.

Compatibility snapshot planning delegates to `xnix-compat-snapshot-plan`. Snapshot plans scope restore points to application state, Runtime metadata, and desktop activation receipts, exclude user documents and the host system, require confirmation before restore, and use bounded retention.

Compatibility test planning delegates to `xnix-compat-test-plan`. Test plans combine recipe validation, Portal preflight, snapshot preflight, and Runtime launch-binding checks into desktop-safe preflight, smoke, and repair-readiness plans. They can feed Compatibility Center cards and repair planning without exposing backend implementation details.

Compatibility test results delegate to `xnix-compat-test-result`. Test results summarize passed, pending, and blocked outcomes from the Runtime-owned plan so AI diagnostics can reason about current compatibility state without claiming backend execution before a real launch backend exists.

AI diagnostic inputs delegate to `xnix-ai-diagnostic-input`. The input model packages recipe facts, run-plan status, test-result status, repair context, diagnostic signals, allowed AI tasks, blocked AI tasks, and privacy boundaries. It does not call an AI provider, does not require network access, and excludes user documents, host paths, raw backend logs, secrets, and backend implementation details.

AI diagnostic recommendations delegate to `xnix-ai-diagnostic-recommendation`. The recommendation model converts Runtime-safe diagnostic input into user-visible recommendations for explanation, restore-point preparation, and test-progress presentation. Recommendations are review-first, do not auto-execute repair actions, do not call an AI provider, and do not require network access.

AI repair approval gates delegate to `xnix-ai-repair-approval-gate`. The gate model converts review-first recommendations into an explicit Runtime approval decision, blocks repair execution until Compatibility Center review, Runtime approval, and restore-point preflight gates pass, and keeps AI diagnostics non-executing. It does not create restore points, change compatibility modes, call an AI provider, or require network access.

Runtime service binding status delegates to `xnix-runtime-service-binding`. The binding model verifies that D-Bus activation, systemd hardening, the packaged libexec wrapper, and the D-Bus contract agree on the Runtime service boundary. It distinguishes activation binding readiness from live production D-Bus ownership, does not mutate the host root, does not require a privileged container, and does not claim backend readiness.

Runtime activation smoke delegates to `scripts/runtime_activation_smoke.rb`. The smoke installs activation files into a temporary root, including the packaged libexec wrapper, Runtime Ruby libraries, version metadata, recipe assets, and D-Bus reference files. It then runs the staged `/usr/libexec/xnix/compatd` through probe, read-only dispatch, write-gate dispatch, and gated write rejection. The smoke uses only the staging root and keeps live production D-Bus ownership pending.

Runtime live owner gates delegate to `xnix-runtime-live-owner-gate` and `GetRuntimeLiveOwnerGate`. The gate model keeps production bus ownership separate from the Linux session-bus smoke adapter, blocks KDE from claiming Runtime ownership, and records the remaining transition gates: long-running owner packaging, stable bus-name acquisition, read-only method parity, and production recipe trust. It is read-only and does not enable launch, install, repair, restore, settings persistence, host-root mutation, or backend-specific details.

Runtime owner smoke planning delegates to `xnix-runtime-owner-smoke-plan` and `GetRuntimeOwnerSmokePlan`. The smoke plan defines how a future packaged owner must prove activation-file alignment, stable bus-name ownership, read-only method parity, write-method rejection, smoke-adapter separation, and KDE-safe summary output. It is read-only, runs in a restricted-session model, does not start a host system service, does not claim the production bus, and does not enable backend launch or settings persistence.

Runtime method parity manifests delegate to `xnix-runtime-method-parity-manifest` and `GetRuntimeMethodParityManifest`. The manifest checks that the D-Bus XML contract, Runtime dispatch, D-Bus client wrapper, C smoke adapter, and session smoke script all expose the same KDE-facing read-only methods. Passing parity is required before production owner smoke can be trusted, but it does not enable write methods such as install, launch, snapshot, or restore.

Runtime write gates delegate to `xnix-runtime-write-gate` and `GetRuntimeWriteGate`. The gate model gives KDE a desktop-safe explanation for blocked `InstallRecipe`, `Launch`, `CreateSnapshot`, and `RestoreSnapshot` requests. Until production Runtime ownership, backend binding, recipe trust, user review, Portal approval, and snapshot preflight gates pass, write methods return `org.xnix.Compatibility1.Error.WriteMethodDisabled`; they do not create request objects, start execution, mutate the host root, or expose backend details.

## Permissions and Asynchronous Requests

Desktop-sensitive actions must use XDG Desktop Portal. Portal operations return request objects and complete with signals, so Runtime methods that need user approval return an object path and complete through `RequestCompleted`. The Runtime must not promise direct access to a user's files, clipboard, camera, printer, display, or screen capture.

Portal access policy delegates to `xnix-portal-access-policy`. The policy covers file open, URI open, print, screenshot, clipboard, camera, and remote-desktop operations. KDE settings may display the decision, but the Runtime owns the policy; the desktop shell requests authorization and receives only the resulting user-mediated grant.

Portal request modeling delegates to `xnix-portal-request-model`. The model describes the XDG Desktop Portal D-Bus destination, interface, method, request handle token, completion `Response` signal, denied-state guidance, and Runtime-owned result handling. It does not call the host portal during model generation and does not mutate host permissions.

## KDE Integration Rules

- The Compatibility Center Plasmoid is a Runtime client, not an alternative Runtime.
- Runtime planning methods must be available through the D-Bus contract as read-only calls.
- Runtime D-Bus clients must wrap planning reads directly and parse boolean variants as booleans.
- The Compatibility Center model may summarize applications, compatibility state, and pending actions, but must not expose backend storage paths or implementation details.
- The Compatibility Center may show compatibility test summaries, but Runtime code must own test planning and diagnostics policy.
- The Compatibility Center may show compatibility test result summaries, but Runtime code must own result status and diagnostic evidence.
- The Compatibility Center may show AI diagnostic readiness, but Runtime code must own AI input construction, privacy boundaries, and allowed task scope.
- The Compatibility Center may show AI diagnostic recommendations, but Runtime code must own recommendation policy and must not auto-execute repair actions.
- The Compatibility Center may show AI repair approval gates, but Runtime code must own approval tokens, restore-point preflight, and repair execution decisions.
- The Compatibility Center may show Runtime service binding status, live owner gates, owner smoke plans, and method parity manifests, but Runtime code must own D-Bus activation, systemd hardening, production bus ownership, smoke-adapter boundaries, and write-method rejection.
- KRunner query models must resolve to Runtime application identities and managed launcher actions, not backend commands.
- Launcher entries must call `xnix-compat-launch --app <id> %U` and must not expose backend commands, storage paths, or Windows executable paths.
- File association models must use generated desktop files, standard `mimeapps.list` syntax, and Portal-mediated file-open requests.
- The first-release KDE integration status must include exactly the seven agreed entry points.
- Notification requests must come from Runtime events and must keep backend details out of KDE-facing payloads.
- Settings models must expose user concepts and must not mention backend implementation names or storage details.
- Tray status models must summarize Runtime activity and compatible tray bridge state without owning backend policy.
- Task manager identity models must map windows to Runtime application identities without owning KWin policy decisions.
- KWin window rule models must stay limited to identity and layout hints and must not own backend policy.
- Desktop integration manifests must group all seven first-release KDE artifacts for a recipe activation without exposing backend commands.
- Desktop activation installers must write only under an explicit staging root and must reject direct host-root installation.
- Desktop activation installers must enforce recipe install-gate preflight before writing staged files.
- Desktop activation rollback must use activation receipts, verify SHA-256 digests before removal, and preserve changed files.
- Recipe registries must verify schema version, safe relative paths, SHA-256 digests, and signature status before recipes are treated as managed inputs.
- Runtime application listing must use registry-backed loading when a registry is available.
- Runtime probes must expose recipe trust status for desktop diagnostics.
- Recipe trust policy must explain why recipes are production-trusted, development-only, or untrusted.
- Recipe install gates must block production activation of development-only registries while preserving digest-verified development staging.
- Compatibility engine catalogs must keep engine selection inside the Runtime and must not claim backend readiness before implementation.
- Compatibility artifact manifests must keep signed manifest and artifact cache details inside the Runtime and must not expose cache paths to KDE.
- Compatibility install plans must keep artifact, acquisition, package source, state, and recipe gate readiness inside the Runtime and must not let KDE stage desktop files or launch backends directly.
- Compatibility acquisition preflight must keep artifact download enablement inside the Runtime and must not create network requests or download artifacts during inspection.
- Compatibility package sources must keep signed source policy and package installation enablement inside the Runtime and must not expose package manager commands to KDE.
- Managed backend bindings must keep launch enablement inside the Runtime and must not create execution requests before package source, state root, Portal policy, and snapshot preflight pass.
- Compatibility run plans must keep backend details out of desktop-facing output and must not claim launch backends are ready before implementation.
- Compatibility repair plans must expose approval, snapshot, rollback, and notification requirements before a repair is executed.
- Compatibility snapshot plans must preserve user documents, avoid host-system snapshots, and provide bounded restore-point retention.
- Compatibility test plans must keep recipe, Portal, snapshot, and launch-binding checks Runtime-owned and desktop-safe.
- Compatibility test results must distinguish passed, pending, and blocked checks without exposing backend implementation details.
- AI diagnostic inputs must not call an AI provider, require network access, include user documents, include host paths, include raw backend logs, include secrets, or expose backend implementation details.
- AI diagnostic recommendations must be user-visible, review-first, non-executing, and approval-aware before any Runtime repair action occurs.
- AI repair approval gates must be review-first, non-executing, and blocked until Compatibility Center review, Runtime approval, and restore-point preflight pass.
- Runtime service binding status must distinguish activation binding readiness from live D-Bus ownership and must not mutate the host root during inspection.
- Runtime activation smoke must run the staged packaged wrapper against installed Runtime libraries and assets before claiming activation wrapper readiness.
- Runtime live owner gates must keep the smoke adapter non-production and must not let KDE claim Runtime bus ownership.
- Runtime owner smoke plans must not start host services, claim the production bus, enable write methods, or expose backend details.
- Runtime method parity manifests must verify read-only method coverage without enabling write methods.
- Compatibility settings must keep run mode, resource access, device, network, and snapshot policy under Runtime ownership while exposing only user-facing labels to KDE.
- Application state roots must keep user documents excluded, require Portal grants for user files, and never expose host storage paths.
- Portal access policy must require user-mediated XDG Desktop Portal requests and must deny direct desktop access for sensitive operations.
- Portal request models must describe XDG Desktop Portal requests and completion handling without mutating host permissions.
- KRunner resolves a natural-language query to an application identity and requests a Runtime launch.
- KWin scripts attach window identity and layout metadata; they do not make backend policy decisions.
- Dolphin actions ask the Runtime to select an application and obtain portal-granted documents.
- Dolphin service menus must call `xnix-compat-open %U` and must not expose backend commands, storage paths, or Windows executable paths.
- The tray and notification integrations report Runtime events without exposing backend storage terminology.

## References

- [Fedora Kinoite documentation](https://docs.fedoraproject.org/en-US/fedora/f39/getting-started/)
- [KDE Plasma Widget setup](https://develop.kde.org/docs/plasma/widget/setup/)
- [KWin Scripting API](https://develop.kde.org/docs/plasma/kwin/api/)
- [XDG Desktop Portal API](https://flatpak.github.io/xdg-desktop-portal/docs/api-reference.html)

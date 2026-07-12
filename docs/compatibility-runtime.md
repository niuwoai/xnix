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

Generated application launchers delegate to `xnix-compat-launch --app <id>`. That entry point validates the Runtime application id, accepts optional `file://` URIs from desktop file associations, and emits a Runtime `Launch` request model. Plain application launches do not need file portal access; file launches are marked as portal-mediated.

The Dolphin service menu delegates selected files to `xnix-compat-open`. That entry point accepts only `file://` URIs, resolves a Runtime recipe by extension unless an explicit application id is supplied, and emits a portal-required Runtime `Launch` request model. It does not start Wine, a virtual machine, or a backend-specific executable.

`xnix-kde-integration-status` records the first-release KDE entry-point scope. Launcher, task manager, file manager, system tray, notifications, Compatibility Center, and settings are marked as initial because each entry point now has a Runtime-backed request, read-model, or identity model.

Runtime events delegate to `xnix-compat-notify`. That entry point models KDE notification payloads for install failures, automatic repairs, compatibility mode changes, and approval-required events. It chooses urgency and actions for the desktop shell, but it does not inspect backend logs or expose backend implementation terms.

Compatibility settings delegate to `xnix-compat-settings`. That model exposes user-facing controls for run mode, performance or compatibility priority, documents and downloads access, camera access, network access, and snapshots. It intentionally avoids backend implementation terminology.

System tray status delegates to `xnix-compat-tray-status`. That model exposes Runtime activity, attention state, and bridged tray application counts for the KDE shell. It does not own backend policy or create backend tray bridges.

Task manager identity delegates to `xnix-compat-window-identity`. That model exposes desktop file mapping, Runtime application grouping, pinning, restore behavior, and KWin identity-only matching metadata. It does not decide backend policy or manage windows directly.

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

## Permissions and Asynchronous Requests

Desktop-sensitive actions must use XDG Desktop Portal. Portal operations return request objects and complete with signals, so Runtime methods that need user approval return an object path and complete through `RequestCompleted`. The Runtime must not promise direct access to a user's files, clipboard, camera, printer, display, or screen capture.

Portal access policy delegates to `xnix-portal-access-policy`. The policy covers file open, URI open, print, screenshot, clipboard, camera, and remote-desktop operations. KDE settings may display the decision, but the Runtime owns the policy; the desktop shell requests authorization and receives only the resulting user-mediated grant.

## KDE Integration Rules

- The Compatibility Center Plasmoid is a Runtime client, not an alternative Runtime.
- The Compatibility Center model may summarize applications, compatibility state, and pending actions, but must not expose backend storage paths or implementation details.
- Launcher entries must call `xnix-compat-launch --app <id> %U` and must not expose backend commands, storage paths, or Windows executable paths.
- The first-release KDE integration status must include exactly the seven agreed entry points.
- Notification requests must come from Runtime events and must keep backend details out of KDE-facing payloads.
- Settings models must expose user concepts and must not mention backend implementation names or storage details.
- Tray status models must summarize Runtime activity and compatible tray bridge state without owning backend policy.
- Task manager identity models must map windows to Runtime application identities without owning KWin policy decisions.
- Desktop integration manifests must group all seven first-release KDE artifacts for a recipe activation without exposing backend commands.
- Desktop activation installers must write only under an explicit staging root and must reject direct host-root installation.
- Desktop activation installers must enforce recipe install-gate preflight before writing staged files.
- Desktop activation rollback must use activation receipts, verify SHA-256 digests before removal, and preserve changed files.
- Recipe registries must verify schema version, safe relative paths, SHA-256 digests, and signature status before recipes are treated as managed inputs.
- Runtime application listing must use registry-backed loading when a registry is available.
- Runtime probes must expose recipe trust status for desktop diagnostics.
- Recipe trust policy must explain why recipes are production-trusted, development-only, or untrusted.
- Recipe install gates must block production activation of development-only registries while preserving digest-verified development staging.
- Portal access policy must require user-mediated XDG Desktop Portal requests and must deny direct desktop access for sensitive operations.
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

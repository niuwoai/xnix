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

The KDE Compatibility Center consumes a read-only presentation model from `xnix-kde-center-model`. That model is derived from Runtime application and diagnostics data, but it filters backend storage paths and implementation terminology before anything reaches the Plasma shell.

## Application Recipes

A recipe has a stable reverse-DNS application identifier, display name, icon, requested run mode, and supported file extensions. The Runtime generates an ordinary `.desktop` file that calls `xnix-compat-launch --app <id>`. The desktop file intentionally omits backend commands, prefix paths, and Windows executable paths.

Recipe installation must later require a signed source and a schema version. The current model only establishes validation and output invariants; it does not yet trust or execute external recipe files.

## Permissions and Asynchronous Requests

Desktop-sensitive actions must use XDG Desktop Portal. Portal operations return request objects and complete with signals, so Runtime methods that need user approval return an object path and complete through `RequestCompleted`. The Runtime must not promise direct access to a user's files, clipboard, camera, printer, display, or screen capture.

## KDE Integration Rules

- The Compatibility Center Plasmoid is a Runtime client, not an alternative Runtime.
- The Compatibility Center model may summarize applications, compatibility state, and pending actions, but must not expose backend storage paths or implementation details.
- KRunner resolves a natural-language query to an application identity and requests a Runtime launch.
- KWin scripts attach window identity and layout metadata; they do not make backend policy decisions.
- Dolphin actions ask the Runtime to select an application and obtain portal-granted documents.
- The tray and notification integrations report Runtime events without exposing backend storage terminology.

## References

- [Fedora Kinoite documentation](https://docs.fedoraproject.org/en-US/fedora/f39/getting-started/)
- [KDE Plasma Widget setup](https://develop.kde.org/docs/plasma/widget/setup/)
- [KWin Scripting API](https://develop.kde.org/docs/plasma/kwin/api/)
- [XDG Desktop Portal API](https://flatpak.github.io/xdg-desktop-portal/docs/api-reference.html)

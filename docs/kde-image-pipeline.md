# KDE Plasma Atomic Image Pipeline

This subsystem delivers **step 1 of the Delivery Sequence**: an atomic KDE
Plasma desktop image and its standard desktop services. It is independent
of the Compatibility Runtime owner work — it produces the OS image that the
Runtime later runs inside.

## Design: one manifest, one rendered recipe

`image/kinoite/manifest.json` is the **single source of truth**. It declares
the Fedora Kinoite base, the KDE Plasma 6 / portal / runtime support package
set, the KDE entry-point assets layered on top, the image config overlays
(portal backend, branding, unit presets), the units enabled at build time,
and the systemd units the authenticated serial smoke probes. Fedora 44's
Plasma Login Manager owns the graphical login. The development Runtime CLI
wrapper is deliberately not installed or enabled as a production D-Bus owner.

`lib/xnix/image/kde_image.rb` loads that manifest, validates it against the
files actually present in the repository, and **renders the Containerfile**
from it. The checked-in `image/kinoite/Containerfile` is a snapshot of that
render; a test and the build driver both refuse to proceed if the two drift.
This keeps the definition and the build recipe from ever disagreeing.

The **disk-image step** turns the built container into a bootable disk.
`image/kinoite/disk-config.json` is the source of truth for
`bootc-image-builder` (its `source_image` is cross-checked against the image
manifest name), and `lib/xnix/image/disk_build.rb` renders the exact
privileged podman command per output type (qcow2/raw/iso). Its blueprint
sets `console=ttyS0` and supplies the development smoke account. The boot
driver signs in on the serial console, queries the graphical target and login
manager, and terminates QEMU after an explicit pass/fail result. This closes
the chain:

```
manifest.json ─render─> Containerfile ─podman build─> container image
   └─ disk-config.json ─bootc-image-builder─> disk.qcow2 ─QEMU─> boot smoke
```

## What runs where

| Stage | Where it runs | Tooling |
| --- | --- | --- |
| Manifest validation, Containerfile render, drift guard, serial-probe logic | Anywhere Ruby runs, including the constrained CI container | Ruby stdlib only |
| The image compose (multi-GB Fedora Kinoite ostree build) | A **privileged podman/bootc build host** | `podman build` on the rendered Containerfile |
| Disk-image production (qcow2/raw/iso) | A **privileged podman host** (bib reads the container store) | `scripts/build_kde_disk.rb` → `bootc-image-builder` |
| Boot smoke of the produced disk image | A host with `qemu-system-x86_64` + UEFI firmware | `scripts/boot_kde_image.rb` |

The compose cannot run in the project's 1 CPU / 6 GiB unprivileged Debian
build container. `scripts/build_kde_image.rb` detects this and exits with a
clear "toolchain unavailable" message plus the exact command to run on a
capable host — it never reports a build it did not perform.

## Commands

```text
ruby -Ilib test/test_kde_image.rb                                   # image unit tests
ruby -Ilib test/test_kde_disk.rb                                    # disk-build unit tests
ruby -Ilib lib/xnix/image/kde_image.rb validate                    # manifest + on-disk consistency
ruby -Ilib lib/xnix/image/kde_image.rb report                      # JSON summary
ruby -Ilib lib/xnix/image/kde_image.rb containerfile               # render the Containerfile
ruby -Ilib lib/xnix/image/disk_build.rb validate                   # disk config + manifest cross-check
ruby scripts/build_kde_image.rb --check                            # validate + drift + toolchain detect
ruby scripts/build_kde_image.rb                                    # build container (needs podman/buildah)
ruby scripts/build_kde_image.rb --storage-root DIR --runroot DIR --network slirp4netns --add-host HOST:IP --no-proxy
ruby scripts/build_kde_disk.rb --check                             # validate + toolchain detect
ruby scripts/build_kde_disk.rb --type qcow2                        # build disk (needs privileged podman)
ruby scripts/build_kde_disk.rb --type qcow2 --storage-root DIR --runroot DIR
ruby scripts/boot_kde_image.rb --firmware OVMF_CODE.fd             # boot smoke (defaults --disk to the qcow2 output)
```

## Policy ownership

The manifest keeps `runtime_owned: true` / `kde_policy_owner: false`: the
image ships the KDE entry points, but compatibility policy stays with the
Runtime. The image does not enable the development Ruby CLI wrapper as a
system service. Production D-Bus activation remains gated until the Go owner
is implemented and passes its live-owner checks.

## Regenerating the Containerfile

The Containerfile is generated. After any manifest change:

Render with `ruby -Ilib lib/xnix/image/kde_image.rb containerfile`, inspect the
result, and update the checked-in snapshot with a patch-based edit.

The drift guard (`test/test_kde_image.rb` and `scripts/build_kde_image.rb`)
will fail until the snapshot is regenerated.

## Verified release-train evidence

- Authorized q4 builds produced the Fedora 44 container and qcow2. The rc7 verifier reached `graphical.target` and `plasmalogin.service`, persisted the serial log, emitted `XNIX_BOOT_PROBE_PASS`, and exited without a remaining QEMU process. The structured evidence is stored in `docs/release-evidence/v0.2.320-rc7-kde-product-smoke.json`.

## Not yet covered

- Published ostree/bootc artifacts. The authorized q4 artifacts remain local to the build host.
- Interactive Plasma desktop evidence beyond the active `graphical.target` and `plasmalogin.service` serial checks.
- A production Runtime D-Bus owner and Windows application execution evidence.
- A CI definition that chains container build → disk build → boot smoke on a privileged runner.

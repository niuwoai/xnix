// Package image implements a product-image-root smoke: it verifies that a
// built Xnix atomic image would contain the Compatibility Runtime files and
// service definitions declared by an image manifest. It reads the manifest and
// checks the on-disk source files exist; it never runs podman or QEMU, never
// mutates the host root, and requires no network. This product smoke is
// separate from the Buildroot/QEMU learning baseline.
package image

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Artifact is a file laid into the image at build time.
type Artifact struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Dest   string `json:"dest"`
	Role   string `json:"role"`
}

// BootSmoke declares the serial markers a booted image must emit.
type BootSmoke struct {
	ExpectSerialMarkers []string `json:"expect_serial_markers"`
	TimeoutSeconds      int      `json:"timeout_seconds"`
}

// Manifest is the subset of the image manifest the product smoke needs.
type Manifest struct {
	Schema           string     `json:"schema"`
	ImageName        string     `json:"image_name"`
	BaseImage        string     `json:"base_image"`
	Architecture     string     `json:"architecture"`
	BuildMode        string     `json:"build_mode"`
	LayeredArtifacts []Artifact `json:"layered_artifacts"`
	ConfigOverlays   []Artifact `json:"config_overlays"`
	EnabledUnits     []string   `json:"enabled_units"`
	BootSmoke        BootSmoke  `json:"boot_smoke"`
	KDEEntryPoints   []string   `json:"kde_entry_points"`
}

// Roles the product smoke asserts the image provides.
const (
	RoleRuntimeServiceUnit = "runtime-service-unit"
	RoleRuntimeDBus        = "runtime-dbus-activation"
	RoleKDEEntryPoint      = "kde-entry-point"
)

// runtimeServiceUnit is the systemd unit the image must enable at first boot.
const runtimeServiceUnit = "xnix-compatd.service"

// expectedKDEEntryPoints are the KDE-first integration surfaces the flagship
// image must declare so a compatibility application is reachable across the
// desktop: the Compatibility Center, KRunner, KWin window identity, Dolphin,
// the system tray, notifications, and system settings.
var expectedKDEEntryPoints = []string{
	"compatibility-center",
	"krunner",
	"kwin",
	"dolphin",
	"system-tray",
	"notifications",
	"system-settings",
}

// missingEntryPoints returns the expected KDE entry points a manifest does not
// declare, sorted, so the smoke can report coverage gaps deterministically.
func missingEntryPoints(declared []string) []string {
	have := make(map[string]bool, len(declared))
	for _, e := range declared {
		have[e] = true
	}
	var missing []string
	for _, want := range expectedKDEEntryPoints {
		if !have[want] {
			missing = append(missing, want)
		}
	}
	sort.Strings(missing)
	return missing
}

// ParseManifest parses and validates an image manifest.
func ParseManifest(data []byte) (Manifest, error) {
	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return Manifest{}, fmt.Errorf("parse image manifest: %w", err)
	}
	if err := manifest.validate(); err != nil {
		return Manifest{}, err
	}
	return manifest, nil
}

func (m Manifest) validate() error {
	if !strings.HasPrefix(m.Schema, "xnix.image.") {
		return fmt.Errorf("unexpected image manifest schema %q", m.Schema)
	}
	if m.ImageName == "" {
		return errors.New("image manifest must declare an image_name")
	}
	if m.Architecture == "" {
		return errors.New("image manifest must declare an architecture")
	}
	if len(m.LayeredArtifacts) == 0 {
		return errors.New("image manifest must declare layered_artifacts")
	}
	seen := map[string]bool{}
	for _, artifact := range append(append([]Artifact{}, m.LayeredArtifacts...), m.ConfigOverlays...) {
		if artifact.ID == "" || artifact.Source == "" || artifact.Dest == "" {
			return errors.New("image artifact requires id, source, and dest")
		}
		if seen[artifact.ID] {
			return fmt.Errorf("duplicate image artifact id %q", artifact.ID)
		}
		seen[artifact.ID] = true
	}
	return nil
}

func (m Manifest) enablesRuntimeService() bool {
	for _, unit := range m.EnabledUnits {
		if unit == runtimeServiceUnit {
			return true
		}
	}
	return false
}

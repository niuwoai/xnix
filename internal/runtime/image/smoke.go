package image

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// FileCheck is the presence result for one manifest artifact source.
type FileCheck struct {
	ID      string `json:"id"`
	Source  string `json:"source"`
	Role    string `json:"role"`
	Overlay bool   `json:"overlay"`
	Present bool   `json:"present"`
}

// SmokeReport is the deterministic result of the product-image-root smoke.
type SmokeReport struct {
	ImageName                 string      `json:"image_name"`
	BaseImage                 string      `json:"base_image"`
	Architecture              string      `json:"architecture"`
	Checks                    []FileCheck `json:"checks"`
	MissingSources            []string    `json:"missing_sources"`
	RuntimeServiceUnitPresent bool        `json:"runtime_service_unit_present"`
	DBusActivationPresent     bool        `json:"dbus_activation_present"`
	RuntimeServiceEnabled     bool        `json:"runtime_service_enabled"`
	KDEEntryPointCount        int         `json:"kde_entry_point_count"`
	SerialMarkers             []string    `json:"serial_markers"`
	AllSourcesPresent         bool        `json:"all_sources_present"`
	RuntimeReady              bool        `json:"runtime_ready"`
	HostRootModified          bool        `json:"host_root_modified"`
	QEMULaunched              bool        `json:"qemu_launched"`
	NetworkRequired           bool        `json:"network_required"`
	PrivilegedRequired        bool        `json:"privileged_required"`
	Summary                   string      `json:"summary"`
}

// LoadManifest reads and parses an image manifest from a file.
func LoadManifest(path string) (Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Manifest{}, fmt.Errorf("read image manifest: %w", err)
	}
	return ParseManifest(data)
}

// safeSource resolves a manifest source path inside repoRoot and refuses any
// path that escapes it. Sources must be relative.
func safeSource(repoRoot, source string) (string, error) {
	if source == "" || filepath.IsAbs(source) {
		return "", fmt.Errorf("image source must be a relative path: %q", source)
	}
	joined := filepath.Join(repoRoot, filepath.FromSlash(source))
	rel, err := filepath.Rel(repoRoot, joined)
	if err != nil {
		return "", err
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
		return "", fmt.Errorf("image source escapes repo root: %q", source)
	}
	return joined, nil
}

// sourceExists reports whether a manifest source file or directory exists.
func sourceExists(repoRoot, source string) (bool, error) {
	path, err := safeSource(repoRoot, source)
	if err != nil {
		return false, err
	}
	if _, statErr := os.Stat(path); statErr != nil {
		if os.IsNotExist(statErr) {
			return false, nil
		}
		return false, statErr
	}
	return true, nil
}

// Verify runs the product-image-root smoke against a manifest, checking that
// every declared source exists under repoRoot and that the Runtime service and
// D-Bus activation are provided and enabled. It performs no builds, launches no
// QEMU, and never mutates the host root.
func Verify(repoRoot string, manifest Manifest) (SmokeReport, error) {
	if repoRoot == "" {
		return SmokeReport{}, errors.New("image smoke requires a repo root")
	}

	report := SmokeReport{
		ImageName:             manifest.ImageName,
		BaseImage:             manifest.BaseImage,
		Architecture:          manifest.Architecture,
		RuntimeServiceEnabled: manifest.enablesRuntimeService(),
		KDEEntryPointCount:    len(manifest.KDEEntryPoints),
		SerialMarkers:         append([]string(nil), manifest.BootSmoke.ExpectSerialMarkers...),
	}

	check := func(artifact Artifact, overlay bool) error {
		present, err := sourceExists(repoRoot, artifact.Source)
		if err != nil {
			return err
		}
		report.Checks = append(report.Checks, FileCheck{
			ID:      artifact.ID,
			Source:  artifact.Source,
			Role:    artifact.Role,
			Overlay: overlay,
			Present: present,
		})
		if !present {
			report.MissingSources = append(report.MissingSources, artifact.Source)
		}
		if present {
			switch artifact.Role {
			case RoleRuntimeServiceUnit:
				report.RuntimeServiceUnitPresent = true
			case RoleRuntimeDBus:
				report.DBusActivationPresent = true
			}
		}
		return nil
	}

	for _, artifact := range manifest.LayeredArtifacts {
		if err := check(artifact, false); err != nil {
			return SmokeReport{}, err
		}
	}
	for _, artifact := range manifest.ConfigOverlays {
		if err := check(artifact, true); err != nil {
			return SmokeReport{}, err
		}
	}

	sort.Strings(report.MissingSources)
	report.AllSourcesPresent = len(report.MissingSources) == 0
	report.RuntimeReady = report.AllSourcesPresent &&
		report.RuntimeServiceUnitPresent &&
		report.DBusActivationPresent &&
		report.RuntimeServiceEnabled
	report.HostRootModified = false
	report.QEMULaunched = false
	report.NetworkRequired = false
	report.PrivilegedRequired = false

	if report.RuntimeReady {
		report.Summary = fmt.Sprintf("Image %s provides the Compatibility Runtime service, D-Bus activation, and all %d declared sources.", manifest.ImageName, len(report.Checks))
	} else {
		report.Summary = fmt.Sprintf("Image %s is not Runtime-ready: %d missing source(s); runtime service present=%t enabled=%t.", manifest.ImageName, len(report.MissingSources), report.RuntimeServiceUnitPresent, report.RuntimeServiceEnabled)
	}
	return report, nil
}

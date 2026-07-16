package image

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func fixtureManifest() Manifest {
	return Manifest{
		Schema:       "xnix.image.kinoite/v1",
		ImageName:    "xnix-kinoite",
		BaseImage:    "quay.io/fedora/fedora-kinoite:41",
		Architecture: "x86_64",
		BuildMode:    "bootc-image",
		LayeredArtifacts: []Artifact{
			{ID: "compatd-system-unit", Source: "runtime/systemd/xnix-compatd.service", Dest: "/usr/lib/systemd/system/xnix-compatd.service", Role: RoleRuntimeServiceUnit},
			{ID: "compatd-dbus-service", Source: "runtime/dbus/org.xnix.Compatibility1.service", Dest: "/usr/share/dbus-1/services/org.xnix.Compatibility1.service", Role: RoleRuntimeDBus},
			{ID: "dolphin-service-menu", Source: "kde/dolphin.desktop", Dest: "/usr/share/kio/x.desktop", Role: RoleKDEEntryPoint},
		},
		ConfigOverlays: []Artifact{
			{ID: "portal-backend", Source: "image/config/portal.conf", Dest: "/usr/share/xdg-desktop-portal/x.conf", Role: "portal-backend"},
		},
		EnabledUnits:   []string{"sddm.service", "xnix-compatd.service"},
		BootSmoke:      BootSmoke{ExpectSerialMarkers: []string{"Reached target Graphical Interface"}, TimeoutSeconds: 180},
		KDEEntryPoints: []string{"compatibility-center", "krunner", "dolphin"},
	}
}

func buildImageRoot(t *testing.T, manifest Manifest, omit map[string]bool) string {
	t.Helper()
	root := t.TempDir()
	for _, a := range append(append([]Artifact{}, manifest.LayeredArtifacts...), manifest.ConfigOverlays...) {
		if omit[a.ID] {
			continue
		}
		abs := filepath.Join(root, filepath.FromSlash(a.Source))
		if err := os.MkdirAll(filepath.Dir(abs), 0o700); err != nil {
			t.Fatalf("MkdirAll: %v", err)
		}
		if err := os.WriteFile(abs, []byte("fixture"), 0o600); err != nil {
			t.Fatalf("WriteFile: %v", err)
		}
	}
	return root
}

func TestVerifyRuntimeReadyWhenAllSourcesPresent(t *testing.T) {
	manifest := fixtureManifest()
	root := buildImageRoot(t, manifest, nil)

	report, err := Verify(root, manifest)
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}
	if !report.AllSourcesPresent || !report.RuntimeReady {
		t.Fatalf("image should be runtime-ready: %#v", report)
	}
	if !report.RuntimeServiceUnitPresent || !report.DBusActivationPresent || !report.RuntimeServiceEnabled {
		t.Fatalf("runtime service/dbus/enable flags wrong: %#v", report)
	}
	if report.KDEEntryPointCount != 3 || len(report.Checks) != 4 {
		t.Fatalf("unexpected counts: %#v", report)
	}
	if report.HostRootModified || report.QEMULaunched || report.NetworkRequired || report.PrivilegedRequired {
		t.Fatalf("product smoke must not launch QEMU, use network, or mutate host: %#v", report)
	}
}

func TestVerifyReportsEntryPointCoverage(t *testing.T) {
	manifest := fixtureManifest()
	// The fixture declares only three entry points, so coverage is incomplete.
	manifest.KDEEntryPoints = []string{"compatibility-center", "krunner", "dolphin"}
	root := buildImageRoot(t, manifest, nil)

	report, err := Verify(root, manifest)
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}
	if report.KDEEntryPointsCovered {
		t.Fatalf("partial entry points must not be covered: %#v", report)
	}
	want := []string{"kwin", "notifications", "system-settings", "system-tray"}
	if len(report.MissingEntryPoints) != len(want) {
		t.Fatalf("unexpected missing entry points: %#v", report.MissingEntryPoints)
	}
	for i := range want {
		if report.MissingEntryPoints[i] != want[i] {
			t.Fatalf("missing entry points not sorted/expected: %#v", report.MissingEntryPoints)
		}
	}

	// A manifest declaring all seven is fully covered.
	full := fixtureManifest()
	full.KDEEntryPoints = []string{"compatibility-center", "krunner", "kwin", "dolphin", "system-tray", "notifications", "system-settings"}
	fullReport, _ := Verify(buildImageRoot(t, full, nil), full)
	if !fullReport.KDEEntryPointsCovered || len(fullReport.MissingEntryPoints) != 0 {
		t.Fatalf("all seven entry points must be covered: %#v", fullReport)
	}
}

func TestVerifyReportsMissingRuntimeSource(t *testing.T) {
	manifest := fixtureManifest()
	// Omit the D-Bus activation source.
	root := buildImageRoot(t, manifest, map[string]bool{"compatd-dbus-service": true})

	report, err := Verify(root, manifest)
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}
	if report.AllSourcesPresent || report.RuntimeReady || report.DBusActivationPresent {
		t.Fatalf("missing dbus source must not be runtime-ready: %#v", report)
	}
	if len(report.MissingSources) != 1 || report.MissingSources[0] != "runtime/dbus/org.xnix.Compatibility1.service" {
		t.Fatalf("unexpected missing sources: %#v", report.MissingSources)
	}
}

func TestVerifyRequiresRuntimeServiceEnabled(t *testing.T) {
	manifest := fixtureManifest()
	manifest.EnabledUnits = []string{"sddm.service"} // compatd not enabled
	root := buildImageRoot(t, manifest, nil)

	report, err := Verify(root, manifest)
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}
	if report.RuntimeReady || report.RuntimeServiceEnabled {
		t.Fatalf("runtime must not be ready when compatd is not enabled: %#v", report)
	}
}

func TestParseManifestValidation(t *testing.T) {
	if _, err := ParseManifest([]byte(`{"schema":"other/v1","image_name":"x","architecture":"x86_64","layered_artifacts":[{"id":"a","source":"s","dest":"d"}]}`)); err == nil {
		t.Fatalf("bad schema must fail")
	}
	if _, err := ParseManifest([]byte(`{"schema":"xnix.image.k/v1","architecture":"x86_64","layered_artifacts":[{"id":"a","source":"s","dest":"d"}]}`)); err == nil {
		t.Fatalf("missing image_name must fail")
	}
	if _, err := ParseManifest([]byte(`{"schema":"xnix.image.k/v1","image_name":"x","architecture":"x86_64","layered_artifacts":[]}`)); err == nil {
		t.Fatalf("empty layered_artifacts must fail")
	}
	dup := `{"schema":"xnix.image.k/v1","image_name":"x","architecture":"x86_64","layered_artifacts":[{"id":"a","source":"s","dest":"d"},{"id":"a","source":"s2","dest":"d2"}]}`
	if _, err := ParseManifest([]byte(dup)); err == nil {
		t.Fatalf("duplicate id must fail")
	}
}

func TestVerifyRefusesSourceEscape(t *testing.T) {
	manifest := fixtureManifest()
	manifest.LayeredArtifacts[0].Source = "../escape.service"
	root := buildImageRoot(t, fixtureManifest(), nil)
	if _, err := Verify(root, manifest); err == nil {
		t.Fatalf("source escaping repo root must error")
	}
}

// TestRealRepoManifestIsRuntimeReady runs the product smoke against the actual
// committed image manifest, so a clean checkout reproduces a Runtime-ready
// result and drift is caught.
func TestRealRepoManifestIsRuntimeReady(t *testing.T) {
	repoRoot := filepath.Join("..", "..", "..")
	manifestPath := filepath.Join(repoRoot, "image", "kinoite", "manifest.json")
	if _, err := os.Stat(manifestPath); err != nil {
		t.Skipf("image manifest not present: %v", err)
	}
	manifest, err := LoadManifest(manifestPath)
	if err != nil {
		t.Fatalf("LoadManifest: %v", err)
	}
	report, err := Verify(repoRoot, manifest)
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}
	if !report.RuntimeReady {
		out, _ := json.MarshalIndent(report, "", "  ")
		t.Fatalf("committed image manifest is not Runtime-ready:\n%s", out)
	}
}

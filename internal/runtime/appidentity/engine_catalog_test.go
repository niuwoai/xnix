package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestEngineCatalogPreviewReportsDesktopSafeCompatibilityChoices(t *testing.T) {
	preview, err := NewEngineCatalogPreview()
	if err != nil {
		t.Fatalf("NewEngineCatalogPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.engine_catalog.v1" ||
		preview.RequestType != "engine-catalog-preview" ||
		preview.CatalogType != "compatibility-engine-catalog" ||
		preview.Source != "go-runtime-engine-catalog" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetEngineCatalog" ||
		preview.ReadMethod != "GetEngineCatalogPreview" {
		t.Fatalf("unexpected engine catalog schema: %#v", preview)
	}
	if preview.EngineCount != 3 ||
		!sameStrings(preview.EngineIDs, []string{"automatic", "local-compatible", "isolated-compatible"}) ||
		preview.DefaultEngineID != "automatic" {
		t.Fatalf("unexpected engine catalog IDs: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.UserVisible ||
		!preview.BackendTerminologyHidden ||
		preview.SelectionPersisted ||
		preview.BackendInstalled ||
		preview.BackendLaunchEnabled ||
		preview.ExecutionStarted ||
		preview.HostRootModified ||
		preview.NetworkRequired ||
		preview.BackendDetailsExposed ||
		preview.RawCommandExposed {
		t.Fatalf("unexpected engine catalog safety flags: %#v", preview)
	}
	if len(preview.BlockedActions) != 5 ||
		preview.BlockedActions[0] != "persist compatibility engine selection from preview" {
		t.Fatalf("unexpected engine catalog blocked actions: %#v", preview.BlockedActions)
	}

	for _, engine := range preview.Engines {
		if engine.ID == "" ||
			engine.Label == "" ||
			engine.Strategy == "" ||
			engine.Isolation == "" ||
			engine.Availability != "planned" ||
			engine.Status != "ready-for-selection" ||
			!engine.UserSelectable ||
			!engine.RequiresInstall ||
			!engine.RequiresRuntimeReview ||
			engine.BackendInstalled ||
			engine.BackendLaunchEnabled ||
			engine.BackendDetailsExposed ||
			engine.RawCommandExposed {
			t.Fatalf("unexpected engine entry: %#v", engine)
		}
	}
	if !preview.Engines[0].Recommended || preview.Engines[1].Recommended || preview.Engines[2].Recommended {
		t.Fatalf("automatic engine should be the only recommended choice: %#v", preview.Engines)
	}
	if preview.DesktopSafeSummary != "Engine catalog preview is Go-owned and exposes user-facing compatibility choices while installation, launch, persistence, and backend details remain disabled." {
		t.Fatalf("unexpected engine catalog summary: %q", preview.DesktopSafeSummary)
	}
	if err := validateNoBackendTerms(preview, "Engine catalog preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	serialized := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("engine catalog preview exposes forbidden term %q: %s", forbidden, string(encoded))
		}
	}
}

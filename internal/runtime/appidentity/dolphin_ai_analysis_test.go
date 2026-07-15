package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestDolphinAIAnalysisPreviewHidesFilePathsAndContent(t *testing.T) {
	preview, err := NewDolphinAIAnalysisPreview([]Recipe{{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".abc", ".xls"},
	}}, Provenance{Source: "registry", RegistryName: "test-registry"}, []string{"file:///home/test/Documents/book.xls", "file:///home/test/Documents/tax.xls"}, "")
	if err != nil {
		t.Fatalf("NewDolphinAIAnalysisPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.dolphin_ai_analysis.v1" ||
		preview.RequestType != "dolphin-ai-analysis-preview" ||
		preview.Source != "dolphin-ai-action" ||
		preview.AnalysisSurface != "Dolphin" {
		t.Fatalf("unexpected Dolphin AI schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.SelectedExtension != ".xls" ||
		preview.SelectionMode != "extension-match" ||
		preview.FileCount != 2 {
		t.Fatalf("unexpected Dolphin AI identity: %#v", preview)
	}
	if preview.RuntimeMethod != "GetAIDiagnosticInput" ||
		preview.AnalysisTask != "compatibility-file-review" ||
		preview.SelectedFileDisclosure != "count-and-extension-only" {
		t.Fatalf("unexpected Dolphin AI task metadata: %#v", preview)
	}
	if !preview.PortalRequired ||
		preview.PortalInterface != "org.freedesktop.portal.FileChooser" ||
		preview.PortalMethod != "OpenFile" {
		t.Fatalf("unexpected Dolphin AI Portal metadata: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.SafeForAIDiagnostics ||
		!preview.UserReviewRequired ||
		preview.AIProviderCallEnabled ||
		preview.NetworkRequired ||
		preview.FileContentRead ||
		preview.FilePathsExposed ||
		preview.RequestObjectCreated ||
		preview.PermissionGranted ||
		preview.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected Dolphin AI safety flags: %#v", preview)
	}
	if !sameStrings(preview.AllowedAITasks, []string{
		"explain compatibility risk",
		"summarize required user review",
		"suggest safe next steps",
	}) {
		t.Fatalf("unexpected Dolphin AI allowed tasks: %#v", preview.AllowedAITasks)
	}
	if !sameStrings(preview.BlockedActions, []string{
		"read selected file contents before Portal approval",
		"send user documents to an AI provider",
		"expose selected file paths to AI diagnostics",
		"create Portal request objects from an AI preview",
		"grant file permissions from an AI preview",
		"start compatibility backends from an AI preview",
		"mutate the host root from an AI preview",
		"expose backend implementation details in AI prompts",
	}) {
		t.Fatalf("unexpected Dolphin AI blocked actions: %#v", preview.BlockedActions)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"file://", "/home/test", "documents/book.xls", "documents/tax.xls", "prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("Dolphin AI analysis preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

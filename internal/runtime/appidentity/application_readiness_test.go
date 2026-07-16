package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestApplicationReadinessPreviewJoinsRuntimeEvidenceWithoutLaunch(t *testing.T) {
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	})
	if err != nil {
		t.Fatalf("NewPlan returned error: %v", err)
	}

	preview, err := plan.ApplicationReadinessPreview(ApplicationReadinessOptions{RuntimeRoot: "../../.."})
	if err != nil {
		t.Fatalf("ApplicationReadinessPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.application_readiness.v1" ||
		preview.RequestType != "application-readiness-preview" ||
		preview.GraphType != "runtime-application-readiness-evidence-graph" ||
		preview.RuntimeMethod != "GetApplicationReadiness" ||
		preview.ReadMethod != "GetApplicationReadinessPreview" ||
		preview.Application.ID != "org.example.ledger" ||
		preview.Environment != "development" ||
		preview.OverallStatus != "not-ready" ||
		preview.Ready {
		t.Fatalf("unexpected application readiness identity: %#v", preview)
	}
	if strings.Join(preview.NodeIDs, ",") != "recipe-trust,artifact-stage-receipt,backend-lifecycle,portal-review,snapshot-baseline,execution-readiness,runtime-write-gate" ||
		preview.NodeCount != 7 ||
		preview.RequiredNodeCount != 7 ||
		preview.BlockedNodeCount != 2 {
		t.Fatalf("unexpected readiness graph nodes: %#v", preview)
	}
	if preview.RecipeTrustDecision != "block" ||
		preview.InstallReadiness.RecipeInstallAllowed ||
		preview.InstallReadiness.ArtifactStageReceiptReady ||
		preview.BackendLifecycleState != "blocked" ||
		preview.PortalDecision != "ask" ||
		preview.SnapshotReason != "before-repair" ||
		preview.ExecutionState != "blocked" ||
		preview.WriteGateDecision != "blocked-until-production-backend" {
		t.Fatalf("unexpected joined evidence state: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.UserVisible ||
		!preview.CompatibilityCenterCard ||
		preview.LaunchAllowed ||
		preview.LaunchEnabled ||
		preview.WriteMethodsEnabled ||
		preview.ExecutionRequestCreated ||
		preview.ExecutionStarted ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.RealPortalTransportEnabled ||
		preview.RequestObjectCreated ||
		preview.PermissionGranted ||
		preview.SnapshotCreated ||
		preview.RestoreExecuted ||
		preview.HostRootModified ||
		preview.NetworkRequired ||
		preview.PrivilegedContainerRequired ||
		preview.StateRootPathExposed ||
		preview.BackendDetailsExposed ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed {
		t.Fatalf("unexpected readiness safety flags: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("application readiness preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestApplicationReadinessPreviewRejectsUnsupportedWriteMethod(t *testing.T) {
	plan, err := NewPlan(Recipe{ID: "org.example.ledger", Name: "Example Ledger", Icon: "office-chart-area", Mode: "automatic"})
	if err != nil {
		t.Fatalf("NewPlan returned error: %v", err)
	}
	if _, err := plan.ApplicationReadinessPreview(ApplicationReadinessOptions{WriteMethod: "InstallRecipe"}); err == nil {
		t.Fatalf("ApplicationReadinessPreview accepted unsupported write method")
	}
}

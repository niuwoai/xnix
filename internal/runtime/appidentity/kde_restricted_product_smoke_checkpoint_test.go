package appidentity

import (
	"path/filepath"
	"testing"

	"xnix.local/xnix/internal/runtime/execution"
)

func TestKDERestrictedProductSmokeCheckpointJoinsMetadataWithoutExecution(t *testing.T) {
	recipeRecord := Recipe{ID: "org.xnix.sample.notepad", Name: "Sample Notepad", Version: "1.0.0", Icon: "accessories-text-editor", Mode: "automatic", SupportedExtensions: []string{".txt"}}
	provenance := Provenance{Source: "registry", RegistryName: "xnix-offline-samples", DigestVerified: true, SignatureStatus: "development-only"}
	record, err := NewKDERestrictedProductSmokeCheckpointRecord(recipeRecord, provenance, KDERestrictedProductSmokeCheckpointOptions{
		KDERestrictedLaunchAuthorizationOptions: KDERestrictedLaunchAuthorizationOptions{StateRoot: t.TempDir(), Mode: execution.RestrictedTestMode, Directive: execution.RestrictedTestPreparationDirective},
		RepositoryRoot:                          filepath.Join("..", "..", ".."),
		ManifestSource:                          "image/kinoite/manifest.json",
	})
	if err != nil {
		t.Fatalf("NewKDERestrictedProductSmokeCheckpointRecord returned error: %v", err)
	}
	if record.SchemaVersion != "xnix.runtime.kde_restricted_product_smoke_checkpoint.v1" || !record.CheckpointReady || !record.ReadyForTrainGate || record.CoreReceiptCount != 10 || !record.AllChecksPassed || record.CheckCount != 8 || record.PassedCheckCount != 8 {
		t.Fatalf("unexpected checkpoint state: %+v", record)
	}
	if !record.ProductImage.ImageManifestReady || !record.ProductImage.KDEEntryPointsCovered || !record.ProductImage.ProductImageMetadataReady || !record.ProductImage.PacketPrepared || !record.ProductImage.ReadyForAuthorizedSmoke || record.ProductImage.EvidenceCount != 5 || record.ProductImage.ReadyEvidenceCount != 5 || record.ProductImage.MissingEvidenceCount != 0 {
		t.Fatalf("unexpected product-image evidence: %+v", record.ProductImage)
	}
	if record.Preflight.Status != "blocked" || record.Preflight.BlockerCount != 2 || !containsString(record.Preflight.BlockerIDs, "recipe-trust") || !containsString(record.Preflight.BlockerIDs, "runtime-write-gate") || record.Execution.State != "blocked" || record.Session.State != "blocked" {
		t.Fatalf("restricted launch state changed: %+v", record)
	}
	assertKDERestrictedProductSmokeCheckpointDisabled(t, record)
}

func assertKDERestrictedProductSmokeCheckpointDisabled(t *testing.T, record KDERestrictedProductSmokeCheckpointRecord) {
	t.Helper()
	if !record.HumanAuthorizationRequired || record.ExecutionAuthorized || record.DockerExecuted || record.QEMUExecuted || record.ProductSmokeExecuted || record.SerialLogPersisted || record.ReleaseReady || record.LaunchPreflightPassed || record.LaunchAuthorized || record.ExecutionApproved || record.ProcessStartAuthorized || record.CommandMaterialized || record.ExecutablePathResolved || record.BackendSelectedForLaunch || record.BackendLaunchEnabled || record.BackendProcessStarted || record.ProductionBusOwnership || record.NetworkRequired || record.DockerSocketMounted || record.HostNetworkEnabled || record.BroadHostMountEnabled || record.PrivilegedContainerRequired || record.HostRootModified || record.StateRootPathExposed || record.RepositoryRootPathExposed || record.ManifestSourcePathExposed || record.RawCommandExposed || record.BackendDetailsExposed || record.ProductImage.DockerExecuted || record.ProductImage.QEMUExecuted || record.ProductImage.ProductSmokeExecuted || record.ProductImage.SerialLogPersisted || record.ProductImage.ReleaseReady {
		t.Fatalf("restricted product smoke checkpoint enabled an unsafe capability: %+v", record)
	}
}

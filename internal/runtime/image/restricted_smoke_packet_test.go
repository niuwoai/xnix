package image

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPrepareRestrictedProductSmokePacketRealRepo(t *testing.T) {
	repoRoot := filepath.Join("..", "..", "..")
	packet, err := PrepareRestrictedProductSmokePacket(repoRoot, "image/kinoite/manifest.json")
	if err != nil {
		t.Fatalf("PrepareRestrictedProductSmokePacket: %v", err)
	}
	if packet.SchemaVersion != "xnix.runtime.restricted_product_smoke_packet.v1" ||
		packet.RequestType != "restricted-product-smoke-packet-preview" || !packet.PacketPrepared ||
		!packet.ReadyForAuthorizedSmoke || packet.EvidenceCount != 5 || packet.ReadyEvidenceCount != 5 || packet.MissingEvidenceCount != 0 {
		t.Fatalf("unexpected packet readiness: %+v", packet)
	}
	if !packet.RuntimeOwnerEvidenceReady || !packet.ArtifactTrustEvidenceReady || !packet.BackendLifecycleEvidenceReady || !packet.PortalSafetyEvidenceReady || !packet.KDEEntryPointEvidenceReady {
		t.Fatalf("required evidence was not ready: %+v", packet)
	}
	if packet.ProductionRuntimeReady {
		t.Fatalf("restricted smoke readiness must not claim production Runtime activation: %+v", packet)
	}
	assertRestrictedProductSmokePacketDisabled(t, packet)
}

func TestPrepareRestrictedProductSmokePacketReportsMissingEvidence(t *testing.T) {
	repoRoot := filepath.Join("..", "..", "..")
	fixtureRoot := t.TempDir()
	manifestSource := "image/kinoite/manifest.json"
	manifestTarget := filepath.Join(fixtureRoot, manifestSource)
	if err := os.MkdirAll(filepath.Dir(manifestTarget), 0o700); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	manifestData, err := os.ReadFile(filepath.Join(repoRoot, manifestSource))
	if err != nil {
		t.Fatalf("ReadFile manifest: %v", err)
	}
	if err := os.WriteFile(manifestTarget, manifestData, 0o600); err != nil {
		t.Fatalf("WriteFile manifest: %v", err)
	}
	packet, err := PrepareRestrictedProductSmokePacket(fixtureRoot, manifestSource)
	if err != nil {
		t.Fatalf("PrepareRestrictedProductSmokePacket: %v", err)
	}
	if packet.ReadyForAuthorizedSmoke || packet.MissingEvidenceCount != 5 || packet.ReleaseReady {
		t.Fatalf("missing evidence must block readiness: %+v", packet)
	}
	assertRestrictedProductSmokePacketDisabled(t, packet)
}

func TestPrepareRestrictedProductSmokePacketRejectsManifestEscape(t *testing.T) {
	if _, err := PrepareRestrictedProductSmokePacket(t.TempDir(), "../manifest.json"); err == nil {
		t.Fatal("expected manifest escape to fail")
	}
}

func assertRestrictedProductSmokePacketDisabled(t *testing.T, packet RestrictedProductSmokePacket) {
	t.Helper()
	if !packet.HumanAuthorizationRequired || packet.ExecutionAuthorized || packet.DockerExecuted || packet.QEMUExecuted ||
		packet.ProductSmokeExecuted || packet.SerialLogPersisted || !packet.LoopbackOnlyNetworking ||
		packet.DockerSocketMounted || packet.HostNetworkEnabled || packet.BroadHostMountEnabled ||
		packet.PrivilegedContainerRequired || packet.BackendLaunchEnabled || packet.HostRootModified || packet.ReleaseReady {
		t.Fatalf("restricted smoke packet enabled an unsafe capability: %+v", packet)
	}
}

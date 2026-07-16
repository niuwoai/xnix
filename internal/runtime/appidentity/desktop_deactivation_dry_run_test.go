package appidentity

import (
	"errors"
	"testing"

	"xnix.local/xnix/internal/runtime/safety"
)

func newDeactivationPlan(t *testing.T) Plan {
	t.Helper()
	plan, err := NewPlanWithProvenance(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".abc"},
	}, Provenance{Source: "registry", RegistryName: "test-registry", DigestVerified: true, SignatureStatus: "development-only"})
	if err != nil {
		t.Fatalf("NewPlanWithProvenance: %v", err)
	}
	return plan
}

func deactivationSurface(t *testing.T, preview DesktopDeactivationDryRunPreview, id string) DesktopDeactivationSurface {
	t.Helper()
	for _, surface := range preview.Surfaces {
		if surface.ID == id {
			return surface
		}
	}
	t.Fatalf("deactivation surface %q not present; ids=%v", id, preview.SurfaceIDs)
	return DesktopDeactivationSurface{}
}

func assertDeactivationSafe(t *testing.T, preview DesktopDeactivationDryRunPreview) {
	t.Helper()
	if preview.FileDeletionEnabled || preview.MIMEDefaultsWritten || preview.KDECacheRefreshed ||
		preview.ReceiptsRewritten || preview.SessionTerminated || preview.BackendLaunchEnabled ||
		preview.TargetPathExposed || preview.HostRootModified || preview.BackendDetailsExposed {
		t.Fatalf("deactivation preview must keep every unsafe capability disabled: %+v", preview)
	}
	if err := safety.ValidatePayload("desktop deactivation dry run preview", preview); err != nil {
		t.Fatalf("deactivation preview leaks forbidden content: %v", err)
	}
}

func TestDesktopDeactivationDryRunMissingReceiptIsBlocked(t *testing.T) {
	plan := newDeactivationPlan(t)
	preview, err := plan.DesktopDeactivationDryRunPreview(DesktopDeactivationDryRunOptions{})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.OverallState != "blocked" || preview.ReceiptState != "missing" || preview.SurfaceCount != 9 {
		t.Fatalf("no-receipt deactivation should be blocked across 9 surfaces: %+v", preview.OverallState)
	}
	for _, id := range []string{"launcher", "desktop-icon", "mime-association", "dolphin-service-menu"} {
		surface := deactivationSurface(t, preview, id)
		if surface.DeactivationState != "blocked" || !containsDeactivationReason(surface.BlockedReasons, "missing-activation-receipt") {
			t.Fatalf("artifact surface %q should be blocked on a missing receipt: %+v", id, surface)
		}
	}
	for _, id := range []string{"kwin", "system-tray", "notification", "compatibility-center", "settings"} {
		if surface := deactivationSurface(t, preview, id); surface.DeactivationState != "no-host-artifact" {
			t.Fatalf("non-artifact surface %q should report no host artifact: %+v", id, surface)
		}
	}
	if len(preview.RemovalSteps) == 0 {
		t.Fatalf("deactivation preview should surface the rollback removal steps")
	}
	assertDeactivationSafe(t, preview)
}

func TestDesktopDeactivationDryRunClassifiesReceiptErrors(t *testing.T) {
	cases := map[string]string{
		"desktop activation receipt application mismatch: org.other": "unknown-file-owner",
		"desktop activation receipt entry lacks sha256 digest: x":    "installed-file-digest-mismatch",
		"read desktop activation receipt: open /x: no such file":     "missing-activation-receipt",
		"desktop activation receipt is not safe for KDE status":      "activation-receipt-unsafe",
		"some other failure":                                         "activation-receipt-unverified",
	}
	for message, want := range cases {
		if got := classifyDeactivationReceiptError(errors.New(message)); got != want {
			t.Fatalf("classify(%q) = %q, want %q", message, got, want)
		}
	}
}

func TestDesktopDeactivationDryRunRejectsUnsafePlan(t *testing.T) {
	plan := newDeactivationPlan(t)
	plan.BackendDetailsExposed = true
	if _, err := plan.DesktopDeactivationDryRunPreview(DesktopDeactivationDryRunOptions{}); err == nil {
		t.Fatalf("expected an unsafe plan to be rejected")
	}
}

func containsDeactivationReason(reasons []string, want string) bool {
	for _, reason := range reasons {
		if reason == want {
			return true
		}
	}
	return false
}

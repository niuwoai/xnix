package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEControlledLaunchActionSurfaceAuditAcceptsSafeMetadata(t *testing.T) {
	preview, err := PreviewKDEControlledLaunchActionSurfaceAudit(KDEControlledLaunchActionSurfaceAuditRequest{
		DesktopEntryContent: safeKDEControlledLaunchActionSurfaceMetadata(),
		EvidenceHandle:      "runtime/kde-runtime-status-launch-evidence/fixture.json",
	})
	if err != nil {
		t.Fatalf("PreviewKDEControlledLaunchActionSurfaceAudit returned error: %v", err)
	}
	if preview.SchemaVersion != KDEControlledLaunchActionSurfaceAuditSchemaVersion ||
		preview.RequestType != KDEControlledLaunchActionSurfaceAuditRequestType ||
		preview.AuditState != "safe" ||
		!preview.SurfaceSafeForHumanSmoke ||
		!preview.RequiredMetadataPresent ||
		preview.MetadataMalformed ||
		preview.UnsafeOwnerArgumentsPresent ||
		preview.UnsafeBackendTermsPresent ||
		!preview.EvidenceOnlyArgumentShape ||
		!preview.EvidenceHandleShapeAccepted ||
		preview.KDEActionID != "xnix.runtime-status.controlled-launch" ||
		preview.KDEActionLabel != "Xnix Runtime controlled launch" ||
		preview.PublicDBusMethod != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		preview.ForwardedArgumentKind != "evidence-relative-path" ||
		!preview.KDEForwardsOnlyEvidenceHandle ||
		preview.KDEPolicyOwner ||
		preview.OwnerServiceArgsExposedToKDE ||
		preview.OwnerServiceArgumentsRecreated ||
		preview.StateRootAccess ||
		preview.ReceiptReconstruction ||
		preview.RawExecutablePathExposed ||
		preview.BackendCommandExposed ||
		preview.BackendDetailsExposed ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.ExecutionStarted ||
		preview.DBusCalled ||
		preview.KDEConfigurationWritten ||
		preview.RuntimeStateWritten ||
		preview.HostRootModified ||
		preview.DockerSocketMounted ||
		preview.PrivilegedContainerRequired ||
		preview.HostNetworkRequired ||
		preview.NetworkRequired {
		t.Fatalf("unexpected safe surface audit: %#v", preview)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	for _, forbidden := range []string{"--state-root", "owner_service_call_args", ".exe", "wine ", "qemu-system"} {
		if strings.Contains(strings.ToLower(string(encoded)), forbidden) {
			t.Fatalf("surface audit leaked forbidden detail %q: %s", forbidden, string(encoded))
		}
	}
}

func TestKDEControlledLaunchActionSurfaceAuditBlocksUnsafeOwnerArgsBackendTermsMalformedAndMissingEvidence(t *testing.T) {
	cases := []struct {
		name    string
		content string
		handle  string
		want    string
	}{
		{
			name:    "owner-args",
			content: strings.Replace(safeKDEControlledLaunchActionSurfaceMetadata(), "X-Xnix-State-Root-Access=false", "X-Xnix-State-Root-Access=true\nX-Xnix-Unsafe-Example=--state-root /private/runtime", 1),
			handle:  "runtime/kde-runtime-status-launch-evidence/fixture.json",
			want:    "unsafe-owner-args",
		},
		{
			name:    "backend-terms",
			content: safeKDEControlledLaunchActionSurfaceMetadata() + "X-Xnix-Unsafe-Example=run wine C:/private/app.exe\n",
			handle:  "runtime/kde-runtime-status-launch-evidence/fixture.json",
			want:    "unsafe-backend-terms",
		},
		{
			name:    "malformed",
			content: "[Desktop Entry]\nName=Xnix Runtime controlled launch\n",
			handle:  "runtime/kde-runtime-status-launch-evidence/fixture.json",
			want:    "malformed",
		},
		{
			name:    "missing-handle",
			content: safeKDEControlledLaunchActionSurfaceMetadata(),
			handle:  "",
			want:    "missing-evidence-handle",
		},
		{
			name:    "unsafe-handle",
			content: safeKDEControlledLaunchActionSurfaceMetadata(),
			handle:  "../escape.json",
			want:    "missing-evidence-handle",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			preview, err := PreviewKDEControlledLaunchActionSurfaceAudit(KDEControlledLaunchActionSurfaceAuditRequest{
				DesktopEntryContent: tc.content,
				EvidenceHandle:      tc.handle,
			})
			if err != nil {
				t.Fatalf("PreviewKDEControlledLaunchActionSurfaceAudit returned error: %v", err)
			}
			if preview.AuditState != tc.want ||
				preview.SurfaceSafeForHumanSmoke ||
				preview.DBusCalled ||
				preview.KDEConfigurationWritten ||
				preview.RuntimeStateWritten ||
				preview.DesktopLaunchEnabled ||
				preview.BackendLaunchEnabled ||
				preview.ExecutionStarted ||
				preview.HostRootModified ||
				preview.NetworkRequired {
				t.Fatalf("unexpected blocked surface audit: %#v", preview)
			}
			encoded, err := json.Marshal(preview)
			if err != nil {
				t.Fatalf("Marshal returned error: %v", err)
			}
			for _, unsafe := range []string{"/private/runtime", "C:/private/app.exe", "run wine"} {
				if strings.Contains(string(encoded), unsafe) {
					t.Fatalf("surface audit leaked unsafe value %q: %s", unsafe, string(encoded))
				}
			}
		})
	}
}

func safeKDEControlledLaunchActionSurfaceMetadata() string {
	return `[Desktop Entry]
Type=Service
Name=Xnix Runtime controlled launch
Comment=Forward a Runtime-status evidence handle to the Xnix Runtime D-Bus controlled-launch action.
Icon=media-playback-start
X-Xnix-KDE-Action-ID=xnix.runtime-status.controlled-launch
X-Xnix-Runtime-Preview=xnix-runtime-go kde-controlled-launch-action-preview
X-Xnix-Restricted-Smoke-Plan=xnix-runtime-go kde-controlled-launch-session-bus-smoke-plan-preview
X-Xnix-DBus-Service=org.xnix.Compatibility1
X-Xnix-DBus-Object-Path=/org/xnix/Compatibility1
X-Xnix-DBus-Method=org.xnix.Compatibility1.ShowRuntimeControlledLaunch
X-Xnix-Forwarded-Argument=evidence-relative-path
X-Xnix-Forwards-Only-Evidence-Handle=true
X-Xnix-KDE-Policy-Owner=false
X-Xnix-Owner-Service-Args-Exposed-To-KDE=false
X-Xnix-State-Root-Access=false
X-Xnix-Receipt-Reconstruction=false
X-Xnix-Backend-Launch-Enabled=false
X-Xnix-Execution-Started=false
X-Xnix-Host-Root-Modified=false
X-Xnix-Docker-Socket-Mounted=false
X-Xnix-Privileged-Container-Required=false
X-Xnix-Host-Network-Required=false
`
}

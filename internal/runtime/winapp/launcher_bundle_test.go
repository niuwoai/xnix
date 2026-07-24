package winapp

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRecordLauncherBundleWritesManagedLauncherAndDesktopEntry(t *testing.T) {
	tempDir := t.TempDir()
	stateRoot := filepath.Join(tempDir, "state")
	profilePath := filepath.Join(tempDir, "real-app.profile.json")
	writeSmokeProfile(t, profilePath, map[string]any{
		"schema_version":  SmokeProfileSchemaVersion,
		"executable_path": filepath.Join(tempDir, "app", "hello.exe"),
		"state_root":      stateRoot,
		"runner_path":     filepath.Join(tempDir, "private-wine"),
		"timeout":         "30s",
		"stage_app_dir":   true,
	})

	record, err := RecordLauncherBundle(LauncherBundleRequest{
		ProfilePath:      profilePath,
		ApplicationID:    "org.xnix.realapp",
		DisplayName:      "Real Windows App",
		RuntimeBinary:    "go",
		RuntimeArguments: []string{"run", "./cmd/xnix-runtime-go"},
	})
	if err != nil {
		t.Fatalf("RecordLauncherBundle returned error: %v", err)
	}
	if record.Status != PassedStatus ||
		record.SchemaVersion != LauncherBundleSchemaVersion ||
		record.RequestType != LauncherBundleRequestType ||
		record.ApplicationID != "org.xnix.realapp" ||
		record.DisplayName != "Real Windows App" ||
		record.DesktopFileName != "org.xnix.realapp.desktop" ||
		record.LauncherScriptName != "org.xnix.realapp.sh" ||
		record.ReceiptFileName != "org.xnix.realapp.launcher-bundle.json" ||
		record.RuntimeArgumentCount != 2 ||
		!record.FilesWritten ||
		!record.LauncherScriptWritten ||
		!record.DesktopEntryWritten ||
		!record.ReceiptWritten ||
		!record.DesktopEntryExecUsesManagedLauncher ||
		!record.DesktopEntryTerminalDisabled ||
		record.RawProfilePathExposed ||
		record.RawStateRootPathExposed ||
		record.RawRunnerPathExposed ||
		record.RawRuntimeArgvExposed ||
		record.HostRootModified {
		t.Fatalf("unexpected launcher bundle record: %#v", record)
	}

	launcherPath := filepath.Join(stateRoot, "launcher-bundle", "launchers", "org.xnix.realapp.sh")
	launcherText, err := os.ReadFile(launcherPath)
	if err != nil {
		t.Fatalf("ReadFile launcher returned error: %v", err)
	}
	if !strings.Contains(string(launcherText), "exec 'go' 'run' './cmd/xnix-runtime-go' windows-app-run-smoke --profile") ||
		!strings.Contains(string(launcherText), shellQuote(profilePath)) {
		t.Fatalf("launcher script did not preserve managed profile launch: %s", string(launcherText))
	}
	desktopPath := filepath.Join(stateRoot, "launcher-bundle", "applications", "org.xnix.realapp.desktop")
	desktopText, err := os.ReadFile(desktopPath)
	if err != nil {
		t.Fatalf("ReadFile desktop entry returned error: %v", err)
	}
	if !strings.Contains(string(desktopText), "[Desktop Entry]") ||
		!strings.Contains(string(desktopText), "Name=Real Windows App") ||
		!strings.Contains(string(desktopText), "Exec="+desktopEntryExec(launcherPath)) ||
		!strings.Contains(string(desktopText), "Terminal=false") ||
		!strings.Contains(string(desktopText), "X-Xnix-RuntimeOwned=true") {
		t.Fatalf("desktop entry did not preserve managed launcher metadata: %s", string(desktopText))
	}
	receiptPath := filepath.Join(stateRoot, "launcher-bundle", "receipts", "org.xnix.realapp.launcher-bundle.json")
	receiptData, err := os.ReadFile(receiptPath)
	if err != nil {
		t.Fatalf("ReadFile receipt returned error: %v", err)
	}
	var receipt LauncherBundleRecord
	if err := json.Unmarshal(receiptData, &receipt); err != nil {
		t.Fatalf("Unmarshal receipt returned error: %v", err)
	}
	if !receipt.ReceiptWritten || receipt.RawProfilePathExposed || receipt.RawRunnerPathExposed || receipt.RawRuntimeArgvExposed {
		t.Fatalf("unexpected receipt: %#v", receipt)
	}
	output, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("Marshal record returned error: %v", err)
	}
	for _, leaked := range []string{profilePath, stateRoot, filepath.Join(tempDir, "private-wine")} {
		if strings.Contains(string(output), leaked) {
			t.Fatalf("launcher bundle record leaked private value %q: %s", leaked, string(output))
		}
	}
}

func TestRecordLauncherBundleBlocksUnsafeApplicationID(t *testing.T) {
	tempDir := t.TempDir()
	profilePath := filepath.Join(tempDir, "real-app.profile.json")
	writeSmokeProfile(t, profilePath, map[string]any{
		"schema_version":  SmokeProfileSchemaVersion,
		"executable_path": filepath.Join(tempDir, "app", "hello.exe"),
		"state_root":      filepath.Join(tempDir, "state"),
	})

	record, err := RecordLauncherBundle(LauncherBundleRequest{
		ProfilePath:   profilePath,
		ApplicationID: "../bad",
		DisplayName:   "Bad",
	})
	if err != nil {
		t.Fatalf("RecordLauncherBundle returned error: %v", err)
	}
	if record.Status != FailedStatus ||
		record.FilesWritten ||
		!strings.Contains(record.FailureReason, "application id") {
		t.Fatalf("unexpected unsafe application id record: %#v", record)
	}
}

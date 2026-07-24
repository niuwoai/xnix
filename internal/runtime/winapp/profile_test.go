package winapp

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSmokeProfileParsesReusableRealAppSettings(t *testing.T) {
	tempDir := t.TempDir()
	profilePath := filepath.Join(tempDir, "app.profile.json")
	profile := []byte(`{
  "schema_version": "xnix.runtime.windows_app_smoke_profile.v1",
  "executable_path": "path/to/app.exe",
  "working_directory": "path/to/app",
  "runner_path": "path/to/wine",
  "runner_bottle": "operator-bottle",
  "runner_arguments": ["--shim"],
  "arguments": ["--app-flag"],
  "state_root": ".local/xnix/winapp-smoke/profile-state",
  "timeout": "45s",
  "expected_marker": "APP_OK",
  "success_mode": "marker",
  "redact_output": true
}`)
	if err := os.WriteFile(profilePath, profile, 0o600); err != nil {
		t.Fatalf("WriteFile profile returned error: %v", err)
	}

	request, err := LoadSmokeProfile(profilePath)
	if err != nil {
		t.Fatalf("LoadSmokeProfile returned error: %v", err)
	}
	if request.ExecutablePath != "path/to/app.exe" ||
		request.WorkingDirectory != "path/to/app" ||
		request.RunnerPath != "path/to/wine" ||
		request.RunnerBottle != "operator-bottle" ||
		request.StateRoot != ".local/xnix/winapp-smoke/profile-state" ||
		request.Timeout.String() != "45s" ||
		request.ExpectedMarker != "APP_OK" ||
		request.SuccessMode != "marker" ||
		!request.RedactOutput {
		t.Fatalf("unexpected profile request: %#v", request)
	}
	if len(request.RunnerArguments) != 1 ||
		request.RunnerArguments[0] != "--shim" ||
		len(request.Arguments) != 1 ||
		request.Arguments[0] != "--app-flag" {
		t.Fatalf("unexpected profile arguments: %#v", request)
	}
}

func TestLoadSmokeProfileRejectsUnsupportedSchema(t *testing.T) {
	tempDir := t.TempDir()
	profilePath := filepath.Join(tempDir, "bad.profile.json")
	if err := os.WriteFile(profilePath, []byte(`{"schema_version":"invalid"}`), 0o600); err != nil {
		t.Fatalf("WriteFile profile returned error: %v", err)
	}

	if _, err := LoadSmokeProfile(profilePath); err == nil {
		t.Fatalf("LoadSmokeProfile must reject unsupported schemas")
	}
}

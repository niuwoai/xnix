package winapp

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
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
  "redact_output": true,
  "skip_bootstrap": true,
  "stage_app_dir": true
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
		!request.RedactOutput ||
		!request.SkipBootstrap ||
		!request.StageAppDir {
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

func TestSmokeProfileFromRequestRendersReusableSettings(t *testing.T) {
	profile := SmokeProfileFromRequest(Request{
		ExecutablePath:   "path/to/app.exe",
		WorkingDirectory: "path/to/app",
		RunnerPath:       "path/to/wine",
		RunnerBottle:     "operator-bottle",
		RunnerArguments:  []string{"--runner-shim"},
		Arguments:        []string{"--open"},
		StateRoot:        ".local/xnix/winapp-smoke/profile-state",
		Timeout:          45 * time.Second,
		ExpectedMarker:   "APP_OK",
		SuccessMode:      SuccessModeExitCode,
		RedactOutput:     true,
		SkipBootstrap:    true,
		StageAppDir:      true,
	})
	if profile.SchemaVersion != SmokeProfileSchemaVersion ||
		profile.ExecutablePath != "path/to/app.exe" ||
		profile.WorkingDirectory != "path/to/app" ||
		profile.RunnerPath != "path/to/wine" ||
		profile.RunnerBottle != "operator-bottle" ||
		profile.StateRoot != ".local/xnix/winapp-smoke/profile-state" ||
		profile.Timeout != "45s" ||
		profile.ExpectedMarker != "APP_OK" ||
		profile.SuccessMode != SuccessModeExitCode ||
		!profile.RedactOutput ||
		!profile.SkipBootstrap ||
		!profile.StageAppDir ||
		len(profile.RunnerArguments) != 1 ||
		len(profile.Arguments) != 1 {
		t.Fatalf("unexpected rendered profile: %#v", profile)
	}
	if _, err := json.Marshal(profile); err != nil {
		t.Fatalf("Marshal rendered profile returned error: %v", err)
	}
}

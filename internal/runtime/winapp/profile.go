package winapp

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

const SmokeProfileSchemaVersion = "xnix.runtime.windows_app_smoke_profile.v1"

type SmokeProfile struct {
	SchemaVersion    string   `json:"schema_version"`
	ExecutablePath   string   `json:"executable_path"`
	Arguments        []string `json:"arguments"`
	RunnerArguments  []string `json:"runner_arguments"`
	RunnerBottle     string   `json:"runner_bottle"`
	StateRoot        string   `json:"state_root"`
	WorkingDirectory string   `json:"working_directory"`
	RunnerPath       string   `json:"runner_path"`
	Timeout          string   `json:"timeout"`
	ExpectedMarker   string   `json:"expected_marker"`
	SuccessMode      string   `json:"success_mode"`
	RedactOutput     bool     `json:"redact_output"`
	SkipBootstrap    bool     `json:"skip_bootstrap"`
	StageAppDir      bool     `json:"stage_app_dir"`
}

func LoadSmokeProfile(path string) (Request, error) {
	path = strings.TrimSpace(path)
	if path == "" {
		return Request{}, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return Request{}, fmt.Errorf("read Windows app smoke profile: %w", err)
	}
	var profile SmokeProfile
	if err := json.Unmarshal(data, &profile); err != nil {
		return Request{}, fmt.Errorf("decode Windows app smoke profile: %w", err)
	}
	if profile.SchemaVersion != SmokeProfileSchemaVersion {
		return Request{}, fmt.Errorf("unsupported Windows app smoke profile schema %q", profile.SchemaVersion)
	}
	timeout, err := parseProfileDuration(profile.Timeout)
	if err != nil {
		return Request{}, err
	}
	return Request{
		ExecutablePath:   profile.ExecutablePath,
		Arguments:        append([]string{}, profile.Arguments...),
		RunnerArguments:  append([]string{}, profile.RunnerArguments...),
		RunnerBottle:     profile.RunnerBottle,
		StateRoot:        profile.StateRoot,
		WorkingDirectory: profile.WorkingDirectory,
		RunnerPath:       profile.RunnerPath,
		Timeout:          timeout,
		ExpectedMarker:   profile.ExpectedMarker,
		SuccessMode:      profile.SuccessMode,
		RedactOutput:     profile.RedactOutput,
		SkipBootstrap:    profile.SkipBootstrap,
		StageAppDir:      profile.StageAppDir,
	}, nil
}

func parseProfileDuration(value string) (time.Duration, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, nil
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("parse profile timeout: %w", err)
	}
	return duration, nil
}

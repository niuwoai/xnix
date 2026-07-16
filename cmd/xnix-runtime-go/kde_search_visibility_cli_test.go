package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestKDESearchVisibilityPlanPreviewCLI(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	if err := run([]string{"kde-search-visibility-plan-preview", "--registry", registryPath, "--app", app}, &output); err != nil {
		t.Fatalf("run search visibility preview: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("parse output: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.kde_search_visibility.v1" ||
		payload["request_type"] != "kde-search-visibility-plan-preview" {
		t.Fatalf("unexpected payload identity: %+v", payload)
	}
	for _, key := range []string{"desktop_files_written", "mime_defaults_written", "kde_cache_refreshed", "host_files_indexed", "search_index_persisted", "backend_launch_enabled", "host_root_modified"} {
		if payload[key] != false {
			t.Fatalf("preview must keep %q disabled: %+v", key, payload)
		}
	}
}

func TestKDESearchVisibilityPlanPreviewCLIRejectsMissingApp(t *testing.T) {
	registryPath, _ := writeTestRepairGroupRegistry(t)
	tests := [][]string{
		{"kde-search-visibility-plan-preview"},
		{"kde-search-visibility-plan-preview", "--registry", registryPath},
		{"kde-search-visibility-plan-preview", "--registry", registryPath, "--app", "org.example.missing"},
	}
	for _, args := range tests {
		var output bytes.Buffer
		if err := run(args, &output); err == nil {
			t.Fatalf("expected CLI args to fail: %+v", args)
		}
	}
}

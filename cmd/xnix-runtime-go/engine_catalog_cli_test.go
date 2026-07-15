package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestEngineCatalogPreviewCommandRendersGoEngineCatalog(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"engine-catalog-preview"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.engine_catalog.v1" ||
		payload["request_type"] != "engine-catalog-preview" ||
		payload["catalog_type"] != "compatibility-engine-catalog" ||
		payload["source"] != "go-runtime-engine-catalog" ||
		payload["desktop"] != "KDE Plasma" ||
		payload["runtime_method"] != "GetEngineCatalog" ||
		payload["read_method"] != "GetEngineCatalogPreview" {
		t.Fatalf("unexpected engine catalog CLI schema: %#v", payload)
	}
	if payload["engine_count"] != float64(3) ||
		payload["default_engine_id"] != "automatic" {
		t.Fatalf("unexpected engine catalog counts: %#v", payload)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["user_visible"] != true ||
		payload["backend_terminology_hidden"] != true ||
		payload["selection_persisted"] != false ||
		payload["backend_installed"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["raw_command_exposed"] != false {
		t.Fatalf("unexpected engine catalog safety flags: %#v", payload)
	}

	engines := payload["engines"].([]any)
	if len(engines) != 3 {
		t.Fatalf("unexpected engine entries: %#v", engines)
	}
	first := engines[0].(map[string]any)
	if first["id"] != "automatic" ||
		first["label"] != "Automatic" ||
		first["recommended"] != true ||
		first["user_selectable"] != true ||
		first["backend_launch_enabled"] != false ||
		first["backend_details_exposed"] != false ||
		first["raw_command_exposed"] != false {
		t.Fatalf("unexpected automatic engine entry: %#v", first)
	}

	serialized := strings.ToLower(output.String())
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("engine catalog CLI exposes forbidden term %q: %s", forbidden, output.String())
		}
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExternalWinAppBundleImportPlanPreviewCommand(t *testing.T) {
	root := t.TempDir()
	writeExternalWinAppBundleCLIFile(t, root, filepath.Join("Notepad++Portable", "notepad++.exe"), []byte("MZportable"))
	writeExternalWinAppBundleCLIFile(t, root, filepath.Join("Notepad++Portable", "plugins", "mimeTools.dll"), []byte("plugin"))
	writeExternalWinAppBundleCLIFile(t, root, filepath.Join("Notepad++Portable", "localization", "english.xml"), []byte("en"))
	version := currentProjectVersion(t)

	var output bytes.Buffer
	err := run([]string{
		"external-winapp-bundle-import-plan-preview",
		"--bundle-root", root,
		"--executable-relative-path", filepath.ToSlash(filepath.Join("Notepad++Portable", "notepad++.exe")),
		"--app-id", "org.xnix.external.notepadplusplus",
		"--display-name", "Notepad++ Portable",
		"--app-version", "8.9.7",
		"--version", version,
	}, &output)
	if err != nil {
		t.Fatalf("external-winapp-bundle-import-plan-preview returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("parse bundle import plan output: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.external_winapp_bundle_import_plan.v1" ||
		payload["request_type"] != "external-winapp-bundle-import-plan-preview" ||
		payload["plan_type"] != "external-windows-app-portable-bundle-import-plan" ||
		payload["source"] != "go-runtime-external-winapp-bundle-import" ||
		payload["runtime_method"] != "PreviewExternalWinAppBundleImportPlan" ||
		payload["application_id"] != "org.xnix.external.notepadplusplus" ||
		payload["display_name"] != "Notepad++ Portable" ||
		payload["app_version"] != "8.9.7" ||
		payload["bundle_kind"] != "portable-directory" {
		t.Fatalf("unexpected bundle import plan payload: %#v", payload)
	}
	if payload["bundle_file_count"] != float64(3) ||
		payload["sidecar_file_count"] != float64(2) ||
		payload["bundle_directory_count"] != float64(3) ||
		payload["sidecar_directory_count"] != float64(3) ||
		payload["windows_executable_validated"] != true ||
		payload["windows_executable_mz_header_verified"] != true ||
		payload["portable_bundle_import_ready"] != true ||
		payload["sidecar_directory_supported"] != true ||
		payload["single_executable_import_compatible"] != false {
		t.Fatalf("portable bundle sidecar evidence was not emitted: %#v", payload)
	}
	if len(payload["bundle_manifest_sha256"].(string)) != 64 ||
		len(payload["executable_sha256"].(string)) != 64 {
		t.Fatalf("portable bundle output must carry sha256 digests: %#v", payload)
	}
	if strings.Contains(output.String(), root) ||
		payload["bundle_root_path_exposed"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["privileged_container_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false {
		t.Fatalf("bundle import plan output exposed unsafe local details or gates: %s", output.String())
	}
}

func TestExternalWinAppBundleImportPlanPreviewCommandRejectsMissingBundleRoot(t *testing.T) {
	var output bytes.Buffer
	version := currentProjectVersion(t)
	err := run([]string{
		"external-winapp-bundle-import-plan-preview",
		"--executable-relative-path", "app.exe",
		"--app-id", "org.xnix.external.missing",
		"--display-name", "Missing",
		"--version", version,
	}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --bundle-root") {
		t.Fatalf("expected missing bundle root to fail, got %v", err)
	}
}

func writeExternalWinAppBundleCLIFile(t *testing.T, root string, relativePath string, content []byte) {
	t.Helper()
	path := filepath.Join(root, relativePath)
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatalf("create CLI bundle fixture directory: %v", err)
	}
	if err := os.WriteFile(path, content, 0o600); err != nil {
		t.Fatalf("write CLI bundle fixture file: %v", err)
	}
}

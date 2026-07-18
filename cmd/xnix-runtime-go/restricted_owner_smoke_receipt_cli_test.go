package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/owner"
)

func TestRestrictedOwnerSmokeReceiptRecordCommandPersistsReceipt(t *testing.T) {
	stateRoot := t.TempDir()
	var output bytes.Buffer
	err := run([]string{
		"restricted-owner-smoke-receipt-record",
		"--root", projectRootForRuntimeServiceBindingCommandTest(t),
		"--state-root", stateRoot,
		"--mode", owner.RestrictedOwnerSmokeMode,
		"--authorize", owner.RestrictedOwnerSmokeDirective,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.restricted_owner_smoke_receipt.v1" ||
		payload["record_type"] != "restricted-owner-smoke-execution-receipt" ||
		payload["source"] != "runtime-service-activation-preflight+runtime-owner-smoke-batch" ||
		payload["mode"] != owner.RestrictedOwnerSmokeMode ||
		payload["directive"] != owner.RestrictedOwnerSmokeDirective ||
		payload["relative_path"] != "owner-smoke/restricted-owner-smoke-receipt.json" ||
		payload["restricted_smoke_ready"] != true ||
		payload["restricted_smoke_authorized"] != true ||
		payload["production_activation_ready"] != false ||
		payload["system_service_started"] != false ||
		payload["session_bus_claimed"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["write_methods_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["state_root_path_exposed"] != false {
		t.Fatalf("unexpected restricted owner smoke receipt payload: %#v", payload)
	}
	if len(payload["sha256"].(string)) != 64 ||
		payload["check_count"] != float64(7) ||
		payload["passed_check_count"] != float64(7) ||
		payload["all_checks_passed"] != true ||
		payload["receipt_persisted"] != true ||
		payload["receipt_read_back"] != true {
		t.Fatalf("unexpected restricted owner smoke receipt checks: %#v", payload)
	}
	preflight := payload["activation_preflight"].(map[string]any)
	if preflight["preflight_decision"] != "restricted-owner-smoke-ready" ||
		preflight["restricted_smoke_ready"] != true ||
		preflight["blocked_check_count"] != float64(0) {
		t.Fatalf("unexpected restricted owner smoke preflight: %#v", preflight)
	}
	batch := payload["smoke_batch"].(map[string]any)
	if batch["request_type"] != "runtime-owner-smoke-batch-record" ||
		batch["all_read_dispatch_ready"] != true ||
		batch["all_write_denials_ready"] != true ||
		batch["session_bus_claimed"] != false ||
		batch["production_bus_claimed"] != false ||
		batch["system_service_started"] != false ||
		batch["host_root_modified"] != false {
		t.Fatalf("unexpected restricted owner smoke batch summary: %#v", batch)
	}
	if _, err := os.Stat(filepath.Join(stateRoot, "owner-smoke", "restricted-owner-smoke-receipt.json")); err != nil {
		t.Fatalf("restricted owner smoke receipt was not persisted: %v", err)
	}
}

func TestRestrictedOwnerSmokeReceiptRecordCommandRequiresAuthorization(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"restricted-owner-smoke-receipt-record", "--state-root", t.TempDir()}, &output); err == nil {
		t.Fatalf("restricted-owner-smoke-receipt-record must require mode and authorization")
	}
}

func TestRestrictedOwnerSmokeReceiptRecordCommandRequiresStateRoot(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"restricted-owner-smoke-receipt-record",
		"--root", projectRootForRuntimeServiceBindingCommandTest(t),
		"--mode", owner.RestrictedOwnerSmokeMode,
		"--authorize", owner.RestrictedOwnerSmokeDirective,
	}, &output); err == nil {
		t.Fatalf("restricted-owner-smoke-receipt-record must require --state-root")
	}
}

func TestRestrictedOwnerSmokeReceiptFanOutPreviewCommandCoversReadinessAndSupportSurfaces(t *testing.T) {
	stateRoot := t.TempDir()
	var recordOutput bytes.Buffer
	if err := run([]string{
		"restricted-owner-smoke-receipt-record",
		"--root", projectRootForRuntimeServiceBindingCommandTest(t),
		"--state-root", stateRoot,
		"--mode", owner.RestrictedOwnerSmokeMode,
		"--authorize", owner.RestrictedOwnerSmokeDirective,
	}, &recordOutput); err != nil {
		t.Fatalf("restricted-owner-smoke-receipt-record returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{"restricted-owner-smoke-receipt-fanout-preview", "--state-root", stateRoot}, &output); err != nil {
		t.Fatalf("restricted-owner-smoke-receipt-fanout-preview returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.restricted_owner_smoke_receipt_fanout.v1" ||
		payload["request_type"] != "restricted-owner-smoke-receipt-fanout-preview" ||
		payload["runtime_method"] != "GetRestrictedOwnerSmokeReceiptFanOut" ||
		payload["read_method"] != "GetRestrictedOwnerSmokeReceiptFanOutPreview" ||
		payload["receipt_relative_path"] != "owner-smoke/restricted-owner-smoke-receipt.json" ||
		payload["receipt_record_type"] != "restricted-owner-smoke-execution-receipt" ||
		payload["surface_count"] != float64(5) ||
		payload["check_count"] != float64(8) ||
		payload["passed_check_count"] != float64(8) ||
		payload["all_checks_passed"] != true ||
		payload["receipt_consumed"] != true ||
		payload["receipt_all_checks_passed"] != true ||
		payload["read_only_fan_out"] != true ||
		payload["readiness_surfaces_satisfied"] != true ||
		payload["support_surfaces_satisfied"] != true {
		t.Fatalf("unexpected restricted owner smoke receipt fan-out payload: %s", output.String())
	}
	if strings.Contains(output.String(), stateRoot) {
		t.Fatalf("fan-out command output must not expose state-root path: %s", output.String())
	}
	surfaces := payload["surfaces"].([]any)
	seen := map[string]bool{}
	for _, entry := range surfaces {
		surface := entry.(map[string]any)
		seen[surface["id"].(string)] = true
		for _, key := range []string{"mutates_runtime", "starts_service", "claims_session_bus", "claims_production_bus", "enables_write_methods", "starts_backend", "exports_support_bundle", "creates_support_case", "sends_notification", "exposes_state_root_path", "exposes_backend_details"} {
			if surface[key] != false {
				t.Fatalf("surface gate %s must remain false: %s", key, output.String())
			}
		}
	}
	for _, id := range []string{"runtime-owner-readiness", "service-activation-preflight", "compatibility-onboarding", "support-bundle-manifest", "support-case-timeline"} {
		if !seen[id] {
			t.Fatalf("missing restricted owner smoke receipt fan-out surface %s: %s", id, output.String())
		}
	}
	for _, key := range []string{"state_root_path_exposed", "state_root_writes_enabled", "fan_out_writes_enabled", "production_activation_ready", "production_owner_enabled", "system_service_started", "session_bus_claimed", "production_bus_claimed", "write_methods_enabled", "support_bundle_exported", "support_case_created", "notification_sent", "backend_launch_enabled", "network_required", "host_root_modified", "privileged_container_required", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestRestrictedOwnerSmokeReceiptFanOutPreviewCommandRequiresExistingReceipt(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"restricted-owner-smoke-receipt-fanout-preview"}, &output); err == nil {
		t.Fatalf("restricted-owner-smoke-receipt-fanout-preview must require --state-root")
	}
	if err := run([]string{"restricted-owner-smoke-receipt-fanout-preview", "--state-root", t.TempDir()}, &output); err == nil {
		t.Fatalf("restricted-owner-smoke-receipt-fanout-preview must require an existing receipt")
	}
}

func TestRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"restricted-owner-smoke-receipt-fanout-owner-route-preview", "--root", root, "--receipt-id", owner.RestrictedOwnerSmokeOpaqueReceiptID}, &output); err != nil {
		t.Fatalf("restricted-owner-smoke-receipt-fanout-owner-route-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.restricted_owner_smoke_receipt_fanout_owner_route.v1" ||
		payload["request_type"] != "restricted-owner-smoke-receipt-fanout-owner-route-preview" ||
		payload["runtime_method"] != "GetRestrictedOwnerSmokeReceiptFanOut" ||
		payload["read_method"] != "GetRestrictedOwnerSmokeReceiptFanOutPreview" ||
		payload["opaque_receipt_id"] != owner.RestrictedOwnerSmokeOpaqueReceiptID ||
		payload["receipt_lookup_state"] != "missing-receipt" ||
		payload["receipt_available"] != false ||
		payload["receipt_consumed"] != false ||
		payload["missing_receipt_safe"] != true ||
		payload["owner_managed_lookup"] != true ||
		payload["caller_state_root_required"] != false ||
		payload["fan_out_result_state"] != "missing-receipt-fail-closed" ||
		payload["owner_local_route_candidate_ready"] != true ||
		payload["production_dbus_exposure_ready"] != false ||
		payload["surface_count"] != float64(5) ||
		payload["check_count"] != float64(7) ||
		payload["passed_check_count"] != float64(7) ||
		payload["all_checks_passed"] != true {
		t.Fatalf("unexpected restricted owner smoke fan-out owner-route payload: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("owner-route fan-out command output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"state_root_path_exposed", "state_root_writes_enabled", "runtime_writes_enabled", "fan_out_writes_enabled", "production_activation_ready", "production_owner_enabled", "system_service_started", "session_bus_claimed", "production_bus_claimed", "write_methods_enabled", "support_bundle_exported", "support_case_created", "notification_sent", "backend_launch_enabled", "backend_process_started", "network_required", "host_root_modified", "privileged_container_required", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreviewCommandRejectsBadInputs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"restricted-owner-smoke-receipt-fanout-owner-route-preview", "extra"}, &output); err == nil {
		t.Fatalf("restricted-owner-smoke-receipt-fanout-owner-route-preview must reject positional arguments")
	}
	if err := run([]string{"restricted-owner-smoke-receipt-fanout-owner-route-preview", "--receipt-id", "unknown-receipt"}, &output); err == nil {
		t.Fatalf("restricted-owner-smoke-receipt-fanout-owner-route-preview must reject unknown opaque receipt ids")
	}
}

func TestRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"restricted-owner-smoke-receipt-fanout-owner-route-audit-preview", "--root", root}, &output); err != nil {
		t.Fatalf("restricted-owner-smoke-receipt-fanout-owner-route-audit-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.restricted_owner_smoke_receipt_fanout_owner_route_audit.v1" ||
		payload["request_type"] != "restricted-owner-smoke-receipt-fanout-owner-route-audit-preview" ||
		payload["runtime_method"] != "GetRestrictedOwnerSmokeReceiptFanOutOwnerRouteAudit" ||
		payload["read_method"] != "GetRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview" ||
		payload["subject_command"] != "restricted-owner-smoke-receipt-fanout-preview" ||
		payload["proposed_owner_method"] != "GetRestrictedOwnerSmokeReceiptFanOut" ||
		payload["route_decision"] != "owner-local-route-ready" ||
		payload["current_route_status"] != "owner-local-read-route-ready-production-dbus-blocked" {
		t.Fatalf("unexpected restricted owner smoke fan-out owner-route audit command payload: %s", output.String())
	}
	if payload["cli_command_registered"] != true ||
		payload["go_read_model_present"] != true ||
		payload["owner_dispatch_route_present"] != true ||
		payload["production_dbus_method_present"] != false ||
		payload["consumes_existing_receipt"] != true ||
		payload["requires_caller_state_root"] != false ||
		payload["receipt_lookup_owner_managed"] != true ||
		payload["opaque_receipt_id_supported"] != true ||
		payload["fan_out_writes_enabled"] != false ||
		payload["state_root_writes_enabled"] != false ||
		payload["support_bundle_exported"] != false ||
		payload["support_case_created"] != false ||
		payload["notification_sent"] != false ||
		payload["owner_local_route_candidate_ready"] != true ||
		payload["production_dbus_exposure_ready"] != false {
		t.Fatalf("unexpected restricted owner smoke fan-out owner-route audit command decision: %s", output.String())
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(9) ||
		counts["passed"] != float64(9) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected restricted owner smoke fan-out owner-route audit counts: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("audit command output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "write_methods_enabled", "runtime_writes_enabled", "backend_launch_enabled", "backend_process_started", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"restricted-owner-smoke-receipt-fanout-owner-route-audit-preview", "extra"}, &output); err == nil {
		t.Fatalf("restricted-owner-smoke-receipt-fanout-owner-route-audit-preview must reject positional arguments")
	}
}

func TestRestrictedOwnerSmokeReceiptLookupPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"restricted-owner-smoke-receipt-lookup-preview", "--root", root, "--receipt-id", owner.RestrictedOwnerSmokeOpaqueReceiptID}, &output); err != nil {
		t.Fatalf("restricted-owner-smoke-receipt-lookup-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.restricted_owner_smoke_receipt_lookup.v1" ||
		payload["request_type"] != "restricted-owner-smoke-receipt-lookup-preview" ||
		payload["lookup_type"] != "owner-managed-receipt-lookup" ||
		payload["runtime_method"] != "GetRestrictedOwnerSmokeReceiptLookup" ||
		payload["read_method"] != "GetRestrictedOwnerSmokeReceiptLookupPreview" ||
		payload["opaque_receipt_id"] != owner.RestrictedOwnerSmokeOpaqueReceiptID ||
		payload["receipt_relative_path"] != "owner-smoke/restricted-owner-smoke-receipt.json" ||
		payload["receipt_lookup_state"] != "missing-receipt" ||
		payload["receipt_available"] != false ||
		payload["receipt_consumed"] != false ||
		payload["missing_receipt_safe"] != true ||
		payload["owner_managed_lookup"] != true ||
		payload["caller_state_root_required"] != false ||
		payload["opaque_receipt_id_supported"] != true ||
		payload["read_only_lookup"] != true ||
		payload["check_count"] != float64(6) ||
		payload["passed_check_count"] != float64(6) ||
		payload["all_checks_passed"] != true {
		t.Fatalf("unexpected restricted owner smoke receipt lookup payload: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("lookup command output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"state_root_path_exposed", "state_root_writes_enabled", "runtime_writes_enabled", "fan_out_writes_enabled", "production_owner_enabled", "system_service_started", "session_bus_claimed", "production_bus_claimed", "write_methods_enabled", "support_bundle_exported", "support_case_created", "notification_sent", "backend_launch_enabled", "backend_process_started", "network_required", "host_root_modified", "privileged_container_required", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestRestrictedOwnerSmokeReceiptLookupPreviewCommandRejectsBadInputs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"restricted-owner-smoke-receipt-lookup-preview", "extra"}, &output); err == nil {
		t.Fatalf("restricted-owner-smoke-receipt-lookup-preview must reject positional arguments")
	}
	if err := run([]string{"restricted-owner-smoke-receipt-lookup-preview", "--receipt-id", "unknown-receipt"}, &output); err == nil {
		t.Fatalf("restricted-owner-smoke-receipt-lookup-preview must reject unknown opaque receipt ids")
	}
}

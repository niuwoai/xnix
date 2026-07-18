package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
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

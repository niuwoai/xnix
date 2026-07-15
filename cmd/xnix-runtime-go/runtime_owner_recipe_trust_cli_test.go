package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestRuntimeOwnerRecipeTrustPreviewCommandRendersGoReadModel(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"runtime-owner-recipe-trust-preview", "--root", projectRootForRuntimeServiceBindingCommandTest(t)}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != "0.2.232" ||
		payload["schema_version"] != "xnix.runtime.owner_recipe_trust.v1" ||
		payload["request_type"] != "runtime-owner-recipe-trust-preview" ||
		payload["trust_type"] != "runtime-owner-recipe-trust" ||
		payload["source"] != "registry-digests+recipe-signature-status" ||
		payload["runtime_method"] != "GetRuntimeOwnerRecipeTrust" ||
		payload["read_method"] != "GetRuntimeOwnerRecipeTrustPreview" ||
		payload["registry_path"] != "runtime/recipes/registry.json" {
		t.Fatalf("unexpected Runtime owner recipe trust CLI schema: %#v", payload)
	}
	if payload["registry_name"] != "xnix-local-development" ||
		payload["recipe_count"] != float64(1) ||
		payload["digest_verified"] != true ||
		payload["signed_recipe_validation"] != false ||
		payload["development_registry"] != true ||
		payload["unsigned_recipes_present"] != false ||
		payload["production_recipe_trust_ready"] != false {
		t.Fatalf("unexpected recipe trust flags: %#v", payload)
	}
	statuses := payload["signature_statuses"].([]any)
	if len(statuses) != 1 {
		t.Fatalf("unexpected signature status count: %#v", statuses)
	}
	status := statuses[0].(map[string]any)
	if status["status"] != "development-only" || status["count"] != float64(1) {
		t.Fatalf("unexpected signature status: %#v", status)
	}
	checks := payload["checks"].([]any)
	checkIDs := payload["check_ids"].([]any)
	expectedIDs := []string{"registry-present", "recipe-digests", "signed-recipe-validation", "development-registry", "unsigned-recipes"}
	expectedStatuses := []string{"pass", "pass", "pending", "pending", "pass"}
	if len(checks) != len(expectedIDs) || len(checkIDs) != len(expectedIDs) {
		t.Fatalf("unexpected checks: %#v ids=%#v", checks, checkIDs)
	}
	for index, id := range expectedIDs {
		check := checks[index].(map[string]any)
		if check["id"] != id || checkIDs[index] != id || check["status"] != expectedStatuses[index] {
			t.Fatalf("unexpected check at %d: %#v ids=%#v", index, checks, checkIDs)
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(5) ||
		counts["passed"] != float64(3) ||
		counts["pending"] != float64(2) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected recipe trust counts: %#v", counts)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected recipe trust safety flags: %#v", payload)
	}
	blockingReasons := payload["blocking_reasons"].([]any)
	if len(blockingReasons) != 2 ||
		blockingReasons[0] != "production signed recipe validation is not enabled" ||
		blockingReasons[1] != "registry contains development-only recipes" {
		t.Fatalf("unexpected blocking reasons: %#v", blockingReasons)
	}
	if strings.Contains(strings.ToLower(output.String()), "prefix") ||
		strings.Contains(strings.ToLower(output.String()), ".exe") ||
		strings.Contains(strings.ToLower(output.String()), "program files") {
		t.Fatalf("Runtime owner recipe trust CLI exposed backend terms: %s", output.String())
	}
}

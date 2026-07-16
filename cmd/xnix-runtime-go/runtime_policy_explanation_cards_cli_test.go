package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestRuntimePolicyExplanationCardsPreviewCommand(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	err := run([]string{
		"runtime-policy-explanation-cards-preview",
		"--registry", registryPath,
		"--app", app,
		"--runtime-root", "../..",
		"--mode", "development",
		"--issue", "portal-approval-required",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != "0.2.301" ||
		payload["schema_version"] != "xnix.runtime.policy_explanation_cards.v1" ||
		payload["request_type"] != "runtime-policy-explanation-cards-preview" ||
		payload["card_deck_type"] != "kde-runtime-policy-explanation-card-deck" ||
		payload["runtime_method"] != "GetRuntimePolicyExplanationCards" ||
		payload["read_method"] != "GetRuntimePolicyExplanationCardsPreview" ||
		payload["card_count"] != float64(11) ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["review_only"] != true ||
		payload["cards_persisted"] != false ||
		payload["action_enablement_changed"] != false ||
		payload["request_objects_created"] != false ||
		payload["permission_grants_created"] != false ||
		payload["settings_persisted"] != false ||
		payload["ai_provider_called"] != false ||
		payload["backend_process_started"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected policy explanation payload: %#v", payload)
	}
	ids := payload["card_ids"].([]any)
	cards := payload["cards"].([]any)
	counts := payload["counts"].(map[string]any)
	if len(ids) != 11 ||
		len(cards) != 11 ||
		counts["blocked"] == float64(0) ||
		counts["missing_evidence"] == float64(0) ||
		counts["review_only"] == float64(0) ||
		counts["not_yet_implemented"] == float64(0) {
		t.Fatalf("unexpected card ids or counts: ids=%#v counts=%#v", ids, counts)
	}
	for _, raw := range cards {
		card := raw.(map[string]any)
		if card["action_enabled"] != false ||
			card["card_persisted"] != false ||
			card["request_object_created"] != false ||
			card["permission_granted"] != false ||
			card["settings_persisted"] != false ||
			card["ai_provider_called"] != false ||
			card["backend_process_started"] != false ||
			card["launch_enabled"] != false ||
			card["execution_started"] != false ||
			card["host_root_modified"] != false ||
			card["raw_command_exposed"] != false ||
			card["backend_details_exposed"] != false {
			t.Fatalf("card enabled a side effect: %#v", card)
		}
		assertRuntimePolicyExplanationCLIUserSummarySafe(t, card["user_facing_summary"].(string))
	}
	assertRuntimePolicyExplanationCLISafe(t, output.String())
}

func TestRuntimePolicyExplanationCardsPreviewCommandRejectsAmbiguousSource(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	err := run([]string{"runtime-policy-explanation-cards-preview", "--registry", registryPath, "--app", app, "--recipe", registryPath}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires exactly one source") {
		t.Fatalf("expected source rejection, got %v", err)
	}
}

func TestRuntimePolicyExplanationCardsPreviewCommandRejectsPositionalArguments(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	err := run([]string{"runtime-policy-explanation-cards-preview", "--registry", registryPath, "--app", app, "extra"}, &output)
	if err == nil || !strings.Contains(err.Error(), "does not accept positional arguments") {
		t.Fatalf("expected positional rejection, got %v", err)
	}
}

func assertRuntimePolicyExplanationCLISafe(t *testing.T, text string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "virtual machine", "/users/", "/private/", "docker.sock"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("policy explanation CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}

func assertRuntimePolicyExplanationCLIUserSummarySafe(t *testing.T, summary string) {
	t.Helper()
	serialized := strings.ToLower(summary)
	for _, forbidden := range []string{"backend", "prefix", "wine", "proton", ".exe", "virtual machine", "qemu", "program files"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("policy explanation user summary exposed %q: %s", forbidden, summary)
		}
	}
}

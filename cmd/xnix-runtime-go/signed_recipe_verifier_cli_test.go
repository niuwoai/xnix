package main

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/recipe"
)

func writeSignedRecipeVerifierFixture(t *testing.T, corruptSignature bool) (string, string, string) {
	t.Helper()
	root := t.TempDir()
	recipeData := []byte(`{"id":"org.example.signed"}`)
	digest := sha256.Sum256(recipeData)
	metadata := recipe.SignedMetadata{SchemaVersion: "xnix.recipe.signed_metadata.v1", ApplicationID: "org.example.signed", RecipeSHA256: hex.EncodeToString(digest[:]), Algorithm: "ed25519"}
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	payload, err := metadata.SigningPayload()
	if err != nil {
		t.Fatalf("SigningPayload: %v", err)
	}
	signature := ed25519.Sign(privateKey, payload)
	if corruptSignature {
		signature[0] ^= 0xff
	}
	metadata.SignatureBase64 = base64.StdEncoding.EncodeToString(signature)
	metadataData, _ := json.Marshal(metadata)
	recipePath := filepath.Join(root, "recipe.json")
	metadataPath := filepath.Join(root, "metadata.json")
	publicKeyPath := filepath.Join(root, "public.key")
	for path, data := range map[string][]byte{recipePath: recipeData, metadataPath: metadataData, publicKeyPath: []byte(base64.StdEncoding.EncodeToString(publicKey))} {
		if err := os.WriteFile(path, data, 0o600); err != nil {
			t.Fatalf("WriteFile: %v", err)
		}
	}
	return recipePath, metadataPath, publicKeyPath
}

func TestSignedRecipeVerifierPreviewCLI(t *testing.T) {
	recipePath, metadataPath, publicKeyPath := writeSignedRecipeVerifierFixture(t, false)
	var output bytes.Buffer
	if err := run([]string{"signed-recipe-verifier-preview", "--recipe", recipePath, "--metadata", metadataPath, "--public-key", publicKeyPath}, &output); err != nil {
		t.Fatalf("run verifier preview: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("parse output: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.signed_recipe_verification.v1" || payload["verification_state"] != "fixture-verified" || payload["signature_verified"] != true || payload["production_trust_ready"] != false {
		t.Fatalf("unexpected verifier payload: %+v", payload)
	}
	for _, key := range []string{"production_key_configured", "private_key_loaded", "network_required", "package_manager_invoked", "recipe_written", "registry_migrated", "backend_launch_enabled", "host_root_modified", "recipe_path_exposed", "public_key_path_exposed", "signature_material_exposed"} {
		if payload[key] != false {
			t.Fatalf("verifier preview must keep %q disabled: %+v", key, payload)
		}
	}
	if strings.Contains(output.String(), recipePath) || strings.Contains(output.String(), metadataPath) || strings.Contains(output.String(), publicKeyPath) {
		t.Fatal("verifier preview exposed a fixture path")
	}
}

func TestSignedRecipeVerifierPreviewCLIReportsInvalidSignature(t *testing.T) {
	recipePath, metadataPath, publicKeyPath := writeSignedRecipeVerifierFixture(t, true)
	var output bytes.Buffer
	if err := run([]string{"signed-recipe-verifier-preview", "--recipe", recipePath, "--metadata", metadataPath, "--public-key", publicKeyPath}, &output); err != nil {
		t.Fatalf("run verifier preview: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("parse output: %v", err)
	}
	if payload["verification_state"] != "invalid-signature" || payload["signature_verified"] != false || payload["production_trust_ready"] != false {
		t.Fatalf("invalid signature did not fail closed: %+v", payload)
	}
}

func TestSignedRecipeVerifierPreviewCLIRejectsMissingArguments(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"signed-recipe-verifier-preview"}, &output); err == nil {
		t.Fatal("expected missing arguments to fail")
	}
}

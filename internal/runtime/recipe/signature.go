package recipe

import (
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"xnix.local/xnix/internal/runtime/appid"
)

const signedMetadataSchemaVersion = "xnix.recipe.signed_metadata.v1"

type SignedMetadata struct {
	SchemaVersion   string `json:"schema_version"`
	ApplicationID   string `json:"application_id"`
	RecipeSHA256    string `json:"recipe_sha256"`
	Algorithm       string `json:"algorithm"`
	SignatureBase64 string `json:"signature_base64"`
}

func (metadata SignedMetadata) SigningPayload() ([]byte, error) {
	if metadata.SchemaVersion != signedMetadataSchemaVersion {
		return nil, fmt.Errorf("unsupported signed recipe metadata schema %q", metadata.SchemaVersion)
	}
	if !appid.Valid(metadata.ApplicationID) {
		return nil, errors.New("signed recipe application id must be a reverse-DNS identifier")
	}
	if metadata.Algorithm != "ed25519" {
		return nil, fmt.Errorf("unsupported signed recipe algorithm %q", metadata.Algorithm)
	}
	if len(metadata.RecipeSHA256) != 64 {
		return nil, errors.New("signed recipe digest must be 64 hex characters")
	}
	return []byte(strings.Join([]string{
		metadata.SchemaVersion,
		metadata.ApplicationID,
		metadata.RecipeSHA256,
		metadata.Algorithm,
		"",
	}, "\n")), nil
}

type SignedVerifier struct {
	Metadata   map[string]SignedMetadata
	PublicKeys map[string]ed25519.PublicKey
}

func (verifier SignedVerifier) Verify(entry Entry, recipeData []byte) TrustState {
	state := TrustState{ID: entry.ID, SignatureStatus: entry.SignatureStatus}
	if entry.SignatureStatus != SignatureSigned {
		state.FailedClosed = true
		state.Reason = "production signature metadata is required"
		return state
	}
	if err := verifyDigest(recipeData, entry.SHA256); err != nil {
		state.FailedClosed = true
		state.Reason = err.Error()
		return state
	}
	state.DigestVerified = true

	metadata, ok := verifier.Metadata[entry.ID]
	if !ok {
		state.FailedClosed = true
		state.Reason = "signed recipe metadata is missing"
		return state
	}
	if metadata.ApplicationID != entry.ID {
		state.FailedClosed = true
		state.Reason = "signed recipe application id does not match registry entry"
		return state
	}
	if metadata.RecipeSHA256 != entry.SHA256 {
		state.FailedClosed = true
		state.Reason = "signed recipe digest does not match registry entry"
		return state
	}
	payload, err := metadata.SigningPayload()
	if err != nil {
		state.FailedClosed = true
		state.Reason = err.Error()
		return state
	}
	publicKey, ok := verifier.PublicKeys[entry.ID]
	if !ok || len(publicKey) != ed25519.PublicKeySize {
		state.FailedClosed = true
		state.Reason = "signed recipe public key is missing or invalid"
		return state
	}
	signature, err := base64.StdEncoding.DecodeString(metadata.SignatureBase64)
	if err != nil || len(signature) != ed25519.SignatureSize {
		state.FailedClosed = true
		state.Reason = "signed recipe signature encoding is invalid"
		return state
	}
	if !ed25519.Verify(publicKey, payload, signature) {
		state.FailedClosed = true
		state.Reason = "signed recipe signature is invalid"
		return state
	}

	state.ProductionTrusted = true
	state.Reason = "recipe digest and signature verified"
	return state
}

var _ Verifier = SignedVerifier{}

type SignedRecipeVerificationPreview struct {
	SchemaVersion            string   `json:"schema_version"`
	RequestType              string   `json:"request_type"`
	EvidenceType             string   `json:"evidence_type"`
	ApplicationID            string   `json:"application_id"`
	VerificationState        string   `json:"verification_state"`
	RecipeDigestVerified     bool     `json:"recipe_digest_verified"`
	SignatureVerified        bool     `json:"signature_verified"`
	VerifierBoundaryReady    bool     `json:"verifier_boundary_ready"`
	FixtureMode              bool     `json:"fixture_mode"`
	TestFixtureVerified      bool     `json:"test_fixture_verified"`
	ProductionKeyConfigured  bool     `json:"production_key_configured"`
	ProductionTrustReady     bool     `json:"production_trust_ready"`
	PrivateKeyLoaded         bool     `json:"private_key_loaded"`
	NetworkRequired          bool     `json:"network_required"`
	PackageManagerInvoked    bool     `json:"package_manager_invoked"`
	RecipeWritten            bool     `json:"recipe_written"`
	RegistryMigrated         bool     `json:"registry_migrated"`
	BackendLaunchEnabled     bool     `json:"backend_launch_enabled"`
	HostRootModified         bool     `json:"host_root_modified"`
	RecipePathExposed        bool     `json:"recipe_path_exposed"`
	PublicKeyPathExposed     bool     `json:"public_key_path_exposed"`
	SignatureMaterialExposed bool     `json:"signature_material_exposed"`
	BlockingReasons          []string `json:"blocking_reasons"`
	DesktopSafeSummary       string   `json:"desktop_safe_summary"`
}

func NewSignedRecipeVerificationPreview(entry Entry, recipeData []byte, metadata SignedMetadata, publicKey ed25519.PublicKey) SignedRecipeVerificationPreview {
	trust := (SignedVerifier{
		Metadata:   map[string]SignedMetadata{entry.ID: metadata},
		PublicKeys: map[string]ed25519.PublicKey{entry.ID: publicKey},
	}).Verify(entry, recipeData)
	state := "fixture-verified"
	blockingReasons := []string{"production key configuration is not enabled"}
	if !trust.ProductionTrusted {
		state = signedRecipeVerificationState(trust.Reason)
		blockingReasons = []string{trust.Reason, "production key configuration is not enabled"}
	}
	return SignedRecipeVerificationPreview{
		SchemaVersion: "xnix.runtime.signed_recipe_verification.v1", RequestType: "signed-recipe-verifier-preview",
		EvidenceType: "offline-signed-recipe-verifier-evidence", ApplicationID: entry.ID,
		VerificationState: state, RecipeDigestVerified: trust.DigestVerified,
		SignatureVerified: trust.ProductionTrusted, VerifierBoundaryReady: true, FixtureMode: true,
		TestFixtureVerified: trust.ProductionTrusted, BlockingReasons: blockingReasons,
		DesktopSafeSummary: "Offline recipe signature evidence is available while production key configuration and production trust remain disabled.",
	}
}

func signedRecipeVerificationState(reason string) string {
	switch {
	case strings.Contains(reason, "metadata is missing"):
		return "missing-signature-metadata"
	case strings.Contains(reason, "application id"):
		return "application-id-mismatch"
	case strings.Contains(reason, "digest"):
		return "digest-mismatch"
	case strings.Contains(reason, "schema"):
		return "unsupported-schema"
	case strings.Contains(reason, "algorithm"):
		return "unsupported-algorithm"
	case strings.Contains(reason, "public key"):
		return "invalid-public-key"
	case strings.Contains(reason, "signature"):
		return "invalid-signature"
	default:
		return "blocked"
	}
}

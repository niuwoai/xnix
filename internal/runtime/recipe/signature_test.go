package recipe

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"testing"
)

func signedRecipeFixture(t *testing.T) (Entry, []byte, SignedMetadata, ed25519.PublicKey) {
	t.Helper()
	recipeData := []byte(`{"id":"org.example.signed"}`)
	digest := sha256.Sum256(recipeData)
	entry := Entry{ID: "org.example.signed", SHA256: hex.EncodeToString(digest[:]), SignatureStatus: SignatureSigned}
	metadata := SignedMetadata{
		SchemaVersion: signedMetadataSchemaVersion, ApplicationID: entry.ID,
		RecipeSHA256: entry.SHA256, Algorithm: "ed25519",
	}
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	payload, err := metadata.SigningPayload()
	if err != nil {
		t.Fatalf("SigningPayload: %v", err)
	}
	metadata.SignatureBase64 = base64.StdEncoding.EncodeToString(ed25519.Sign(privateKey, payload))
	return entry, recipeData, metadata, publicKey
}

func TestSignedVerifierAcceptsValidFixture(t *testing.T) {
	entry, data, metadata, publicKey := signedRecipeFixture(t)
	state := (SignedVerifier{Metadata: map[string]SignedMetadata{entry.ID: metadata}, PublicKeys: map[string]ed25519.PublicKey{entry.ID: publicKey}}).Verify(entry, data)
	if !state.DigestVerified || !state.ProductionTrusted || state.FailedClosed || state.DevelopmentOnly {
		t.Fatalf("unexpected trust state: %+v", state)
	}
	preview := NewSignedRecipeVerificationPreview(entry, data, metadata, publicKey)
	if preview.VerificationState != "fixture-verified" || !preview.SignatureVerified || !preview.TestFixtureVerified || preview.ProductionTrustReady || preview.ProductionKeyConfigured || preview.PrivateKeyLoaded {
		t.Fatalf("unexpected preview: %+v", preview)
	}
}

func TestSignedVerifierFailsClosed(t *testing.T) {
	entry, data, metadata, publicKey := signedRecipeFixture(t)
	tests := []struct {
		name  string
		edit  func(*Entry, *[]byte, *SignedMetadata, *ed25519.PublicKey)
		state string
	}{
		{name: "digest mismatch", edit: func(_ *Entry, data *[]byte, _ *SignedMetadata, _ *ed25519.PublicKey) { *data = append(*data, ' ') }, state: "digest-mismatch"},
		{name: "application mismatch", edit: func(_ *Entry, _ *[]byte, metadata *SignedMetadata, _ *ed25519.PublicKey) {
			metadata.ApplicationID = "org.example.other"
		}, state: "application-id-mismatch"},
		{name: "invalid signature", edit: func(_ *Entry, _ *[]byte, metadata *SignedMetadata, _ *ed25519.PublicKey) {
			metadata.SignatureBase64 = base64.StdEncoding.EncodeToString(make([]byte, ed25519.SignatureSize))
		}, state: "invalid-signature"},
		{name: "invalid key", edit: func(_ *Entry, _ *[]byte, _ *SignedMetadata, publicKey *ed25519.PublicKey) {
			*publicKey = ed25519.PublicKey("short")
		}, state: "invalid-public-key"},
		{name: "unsupported schema", edit: func(_ *Entry, _ *[]byte, metadata *SignedMetadata, _ *ed25519.PublicKey) {
			metadata.SchemaVersion = "xnix.recipe.signed_metadata.v2"
		}, state: "unsupported-schema"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			candidateEntry := entry
			candidateData := append([]byte(nil), data...)
			candidateMetadata := metadata
			candidateKey := append(ed25519.PublicKey(nil), publicKey...)
			test.edit(&candidateEntry, &candidateData, &candidateMetadata, &candidateKey)
			preview := NewSignedRecipeVerificationPreview(candidateEntry, candidateData, candidateMetadata, candidateKey)
			if preview.VerificationState != test.state || preview.SignatureVerified || preview.ProductionTrustReady || len(preview.BlockingReasons) == 0 {
				t.Fatalf("unexpected blocked preview: %+v", preview)
			}
		})
	}
}

func TestSignedVerifierRequiresSignedStatusAndMetadata(t *testing.T) {
	entry, data, metadata, publicKey := signedRecipeFixture(t)
	entry.SignatureStatus = SignatureUnsigned
	state := (SignedVerifier{Metadata: map[string]SignedMetadata{entry.ID: metadata}, PublicKeys: map[string]ed25519.PublicKey{entry.ID: publicKey}}).Verify(entry, data)
	if !state.FailedClosed || state.ProductionTrusted {
		t.Fatalf("unsigned recipe must fail closed: %+v", state)
	}
	entry.SignatureStatus = SignatureSigned
	state = (SignedVerifier{PublicKeys: map[string]ed25519.PublicKey{entry.ID: publicKey}}).Verify(entry, data)
	if !state.FailedClosed || state.Reason != "signed recipe metadata is missing" {
		t.Fatalf("missing metadata must fail closed: %+v", state)
	}
}

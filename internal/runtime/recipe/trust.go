// Package recipe implements durable, read-only compatibility recipe storage and
// trust verification. A Store loads recipes from a local recipe root described
// by a signed registry, and a replaceable Verifier boundary classifies each
// recipe's trust state.
//
// Trust is fail-closed: a digest mismatch is never trusted. A recipe without a
// production signature is usable only as development-only and is never marked
// production-trusted. No signing keys, tokens, or secrets are part of this
// package.
package recipe

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
)

// SignatureStatus is the declared signing state of a registry entry.
type SignatureStatus string

const (
	// SignatureSigned means the registry declares a production signature.
	SignatureSigned SignatureStatus = "signed"
	// SignatureDevelopmentOnly means the entry is a development fixture.
	SignatureDevelopmentOnly SignatureStatus = "development-only"
	// SignatureUnsigned means the entry has no signature.
	SignatureUnsigned SignatureStatus = "unsigned"
)

func (s SignatureStatus) valid() bool {
	switch s {
	case SignatureSigned, SignatureDevelopmentOnly, SignatureUnsigned:
		return true
	default:
		return false
	}
}

// TrustState is the explicit, diagnostic trust classification of one recipe. It
// is the single source of truth that trust previews and owner readiness both
// consume.
type TrustState struct {
	ID                string          `json:"id"`
	DigestVerified    bool            `json:"digest_verified"`
	SignatureStatus   SignatureStatus `json:"signature_status"`
	ProductionTrusted bool            `json:"production_trusted"`
	DevelopmentOnly   bool            `json:"development_only"`
	FailedClosed      bool            `json:"failed_closed"`
	Reason            string          `json:"reason"`
}

// Usable reports whether the recipe may be loaded at all (in any environment).
// A fail-closed (digest mismatch / unknown status) recipe is never usable.
func (t TrustState) Usable() bool {
	return t.DigestVerified && !t.FailedClosed
}

// Verifier is the replaceable trust boundary. The default DigestVerifier checks
// content digests and classifies declared signature status; a real signature
// verifier can implement this interface later without changing the Store.
type Verifier interface {
	Verify(entry Entry, recipeData []byte) TrustState
}

// DigestVerifier verifies content digests and classifies the declared signature
// status. It performs no cryptographic signature checking itself: a declared
// "signed" status is treated as production-eligible only because a real
// signature Verifier is expected to replace this boundary in production.
type DigestVerifier struct {
	// TreatSignedAsProduction lets a deployment opt into trusting the declared
	// "signed" status. It defaults to false so this development verifier never
	// silently grants production trust.
	TreatSignedAsProduction bool
}

// Verify classifies the trust state of a registry entry against recipe bytes.
func (v DigestVerifier) Verify(entry Entry, recipeData []byte) TrustState {
	state := TrustState{ID: entry.ID, SignatureStatus: entry.SignatureStatus}

	if !entry.SignatureStatus.valid() {
		state.FailedClosed = true
		state.Reason = fmt.Sprintf("unsupported signature status %q", string(entry.SignatureStatus))
		return state
	}
	if err := verifyDigest(recipeData, entry.SHA256); err != nil {
		state.FailedClosed = true
		state.Reason = err.Error()
		return state
	}
	state.DigestVerified = true

	switch entry.SignatureStatus {
	case SignatureSigned:
		if v.TreatSignedAsProduction {
			state.ProductionTrusted = true
			state.Reason = "digest verified and production signature accepted"
		} else {
			// A declared signature is not production-trusted until a real
			// signature verifier confirms it.
			state.Reason = "digest verified; production signature not confirmed by this verifier"
		}
	case SignatureDevelopmentOnly, SignatureUnsigned:
		state.DevelopmentOnly = true
		state.Reason = "digest verified; development-only, never production trusted"
	}
	return state
}

func verifyDigest(data []byte, expected string) error {
	if len(expected) != 64 {
		return errors.New("registry sha256 must be 64 hex characters")
	}
	if _, err := hex.DecodeString(expected); err != nil {
		return fmt.Errorf("registry sha256 must be hex: %w", err)
	}
	sum := sha256.Sum256(data)
	actual := hex.EncodeToString(sum[:])
	if actual != expected {
		return fmt.Errorf("recipe digest mismatch: %s != %s", actual, expected)
	}
	return nil
}

// compile-time assertion that DigestVerifier satisfies Verifier.
var _ Verifier = DigestVerifier{}

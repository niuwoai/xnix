package artifact

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"xnix.local/xnix/internal/runtime/appid"
)

const stageReceiptSchemaVersion = "xnix.runtime.artifact_stage_receipt.v1"

// StageRequest describes a local-only artifact staging operation.
type StageRequest struct {
	Manifest   Manifest
	CacheRoot  string
	FixtureDir string
}

// StageReceipt is the persisted, desktop-safe artifact staging receipt.
type StageReceipt struct {
	SchemaVersion               string `json:"schema_version"`
	RecordType                  string `json:"record_type"`
	Source                      string `json:"source"`
	ApplicationID               string `json:"application_id"`
	RelativePath                string `json:"relative_path"`
	Plan                        Plan   `json:"plan"`
	SHA256                      string `json:"sha256"`
	RuntimeOwned                bool   `json:"runtime_owned"`
	GoRuntimeBacked             bool   `json:"go_runtime_backed"`
	KDEPolicyOwner              bool   `json:"kde_policy_owner"`
	CacheRootPathExposed        bool   `json:"cache_root_path_exposed"`
	FixtureRootPathExposed      bool   `json:"fixture_root_path_exposed"`
	NetworkRequired             bool   `json:"network_required"`
	NetworkFetchEnabled         bool   `json:"network_fetch_enabled"`
	PackageManagerInvoked       bool   `json:"package_manager_invoked"`
	HostRootModified            bool   `json:"host_root_modified"`
	PrivilegedContainerRequired bool   `json:"privileged_container_required"`
	BackendLaunchEnabled        bool   `json:"backend_launch_enabled"`
	BackendDetailsExposed       bool   `json:"backend_details_exposed"`
	Summary                     string `json:"summary"`
}

// ValidateStageReceipt returns blocking reasons for a persisted staging
// receipt. It verifies the schema, safety flags, relative receipt path,
// receipt digest, and required artifact staging signal without exposing cache
// or fixture roots.
func ValidateStageReceipt(receipt StageReceipt) []string {
	var reasons []string
	if receipt.SchemaVersion != stageReceiptSchemaVersion {
		reasons = append(reasons, "artifact stage receipt schema is not supported")
	}
	if receipt.RecordType != "compatibility-artifact-stage-receipt" {
		reasons = append(reasons, "artifact stage receipt record type is not supported")
	}
	if !appid.Valid(receipt.ApplicationID) {
		reasons = append(reasons, "artifact stage receipt application id is invalid")
	}
	if receipt.RelativePath == "" || filepath.IsAbs(receipt.RelativePath) ||
		strings.Contains(filepath.ToSlash(receipt.RelativePath), "../") ||
		strings.HasPrefix(filepath.ToSlash(receipt.RelativePath), "..") {
		reasons = append(reasons, "artifact stage receipt path must be relative and scoped")
	}
	if receipt.SHA256 == "" || !digestPattern.MatchString(receipt.SHA256) {
		reasons = append(reasons, "artifact stage receipt digest is invalid")
	} else if digest, err := stageReceiptDigest(receipt); err != nil {
		reasons = append(reasons, "artifact stage receipt digest could not be verified")
	} else if digest != receipt.SHA256 {
		reasons = append(reasons, "artifact stage receipt digest mismatch")
	}
	if receipt.Plan.ApplicationID != receipt.ApplicationID {
		reasons = append(reasons, "artifact stage receipt plan application id mismatch")
	}
	if !receipt.Plan.RequiredStaged() {
		reasons = append(reasons, "required artifacts are not staged")
	}
	if !receipt.RuntimeOwned || !receipt.GoRuntimeBacked || receipt.KDEPolicyOwner {
		reasons = append(reasons, "artifact stage receipt owner flags are invalid")
	}
	if receipt.CacheRootPathExposed || receipt.FixtureRootPathExposed {
		reasons = append(reasons, "artifact stage receipt exposes local roots")
	}
	if receipt.NetworkRequired || receipt.NetworkFetchEnabled || receipt.PackageManagerInvoked ||
		receipt.HostRootModified || receipt.PrivilegedContainerRequired ||
		receipt.BackendLaunchEnabled || receipt.BackendDetailsExposed {
		reasons = append(reasons, "artifact stage receipt contains unsafe side effects")
	}
	return reasons
}

// StageFromFixture verifies and stages manifest artifacts from a local fixture
// source into a controlled cache root, then persists a receipt under that root.
func StageFromFixture(req StageRequest) (StageReceipt, error) {
	if req.CacheRoot == "" {
		return StageReceipt{}, errors.New("artifact staging requires an explicit cache root")
	}
	if req.FixtureDir == "" {
		return StageReceipt{}, errors.New("artifact staging requires an explicit fixture root")
	}
	cache, err := NewCache(req.CacheRoot)
	if err != nil {
		return StageReceipt{}, err
	}
	source, err := NewLocalFixtureSource(req.FixtureDir)
	if err != nil {
		return StageReceipt{}, err
	}
	acquirer, err := NewAcquirer(cache, source)
	if err != nil {
		return StageReceipt{}, err
	}
	plan, err := acquirer.Acquire(req.Manifest)
	if err != nil {
		return StageReceipt{}, err
	}
	relativePath, path, err := receiptPath(cache.Root(), req.Manifest.ApplicationID)
	if err != nil {
		return StageReceipt{}, err
	}
	receipt := StageReceipt{
		SchemaVersion:               stageReceiptSchemaVersion,
		RecordType:                  "compatibility-artifact-stage-receipt",
		Source:                      "go-runtime-local-fixture-artifact-staging",
		ApplicationID:               req.Manifest.ApplicationID,
		RelativePath:                relativePath,
		Plan:                        plan,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		CacheRootPathExposed:        false,
		FixtureRootPathExposed:      false,
		NetworkRequired:             false,
		NetworkFetchEnabled:         false,
		PackageManagerInvoked:       false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendLaunchEnabled:        false,
		BackendDetailsExposed:       false,
		Summary:                     "Runtime staged compatibility artifacts from a local fixture source into the controlled cache root without network, package-manager, host-root, or backend side effects.",
	}
	data, digest, err := marshalStageReceipt(receipt)
	if err != nil {
		return StageReceipt{}, err
	}
	receipt.SHA256 = digest
	data, _, err = marshalStageReceipt(receipt)
	if err != nil {
		return StageReceipt{}, err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return StageReceipt{}, fmt.Errorf("prepare artifact stage receipt directory: %w", err)
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return StageReceipt{}, fmt.Errorf("write artifact stage receipt: %w", err)
	}
	return receipt, nil
}

func receiptPath(cacheRoot, applicationID string) (string, string, error) {
	if !appid.Valid(applicationID) {
		return "", "", errors.New("artifact stage receipt application id must be a reverse-DNS identifier")
	}
	relativePath := filepath.ToSlash(filepath.Join("artifact-ledger", "receipts", sanitizeNamespace(applicationID)+".json"))
	path := filepath.Join(cacheRoot, filepath.FromSlash(relativePath))
	rel, err := filepath.Rel(cacheRoot, path)
	if err != nil {
		return "", "", err
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
		return "", "", fmt.Errorf("artifact stage receipt path escapes cache root: %s", relativePath)
	}
	return relativePath, path, nil
}

func marshalStageReceipt(receipt StageReceipt) ([]byte, string, error) {
	data, err := json.MarshalIndent(receipt, "", "  ")
	if err != nil {
		return nil, "", err
	}
	data = append(data, '\n')
	sum := sha256.Sum256(data)
	return data, hex.EncodeToString(sum[:]), nil
}

func stageReceiptDigest(receipt StageReceipt) (string, error) {
	receipt.SHA256 = ""
	_, digest, err := marshalStageReceipt(receipt)
	return digest, err
}

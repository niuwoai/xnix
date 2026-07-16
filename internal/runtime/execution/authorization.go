package execution

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
	"xnix.local/xnix/internal/runtime/rootfs"
)

const (
	RestrictedTestMode                   = "test-only"
	RestrictedTestPreparationScope       = "restricted-test-preparation"
	RestrictedTestPreparationDirective   = "authorize-restricted-test-preparation"
	restrictedAuthorizationSchemaVersion = "xnix.runtime.restricted_launch_authorization.v1"
)

type RestrictedAuthorizationRequest struct {
	ApplicationID string
	RequestID     string
	Mode          string
	Scope         string
	Directive     string
}

// RestrictedAuthorizationReceipt records explicit permission to prepare a
// test-only launch boundary. It never authorizes launch or process start.
type RestrictedAuthorizationReceipt struct {
	SchemaVersion               string `json:"schema_version"`
	RecordType                  string `json:"record_type"`
	Source                      string `json:"source"`
	AuthorizationID             string `json:"authorization_id"`
	ApplicationID               string `json:"application_id"`
	RequestID                   string `json:"request_id"`
	Mode                        string `json:"mode"`
	Scope                       string `json:"scope"`
	Directive                   string `json:"directive"`
	State                       string `json:"state"`
	RelativePath                string `json:"relative_path"`
	SHA256                      string `json:"sha256"`
	PreparationAuthorized       bool   `json:"preparation_authorized"`
	LaunchAuthorized            bool   `json:"launch_authorized"`
	ProcessStartAuthorized      bool   `json:"process_start_authorized"`
	ExecutionApproved           bool   `json:"execution_approved"`
	ProductionTrustSatisfied    bool   `json:"production_trust_satisfied"`
	RuntimeWriteGateEnabled     bool   `json:"runtime_write_gate_enabled"`
	ArtifactAcquisitionEnabled  bool   `json:"artifact_acquisition_enabled"`
	BackendInstallEnabled       bool   `json:"backend_install_enabled"`
	BackendLaunchEnabled        bool   `json:"backend_launch_enabled"`
	BackendProcessStarted       bool   `json:"backend_process_started"`
	RealPortalCallEnabled       bool   `json:"real_portal_call_enabled"`
	StateRootPathExposed        bool   `json:"state_root_path_exposed"`
	RawCommandExposed           bool   `json:"raw_command_exposed"`
	BackendDetailsExposed       bool   `json:"backend_details_exposed"`
	NetworkRequired             bool   `json:"network_required"`
	PrivilegedContainerRequired bool   `json:"privileged_container_required"`
	HostRootModified            bool   `json:"host_root_modified"`
	Summary                     string `json:"summary"`
}

type RestrictedAuthorizationStore struct {
	root *rootfs.Root
}

func NewRestrictedAuthorizationStore(stateRoot string) (*RestrictedAuthorizationStore, error) {
	root, err := rootfs.Open(stateRoot)
	if err != nil {
		return nil, err
	}
	return &RestrictedAuthorizationStore{root: root}, nil
}

func (s *RestrictedAuthorizationStore) Record(request RestrictedAuthorizationRequest) (RestrictedAuthorizationReceipt, error) {
	if err := validateRestrictedAuthorizationRequest(request); err != nil {
		return RestrictedAuthorizationReceipt{}, err
	}
	authorizationID := restrictedAuthorizationID(request.ApplicationID, request.RequestID)
	relativePath, path, err := s.authorizationPath(authorizationID, true)
	if err != nil {
		return RestrictedAuthorizationReceipt{}, err
	}
	receipt := RestrictedAuthorizationReceipt{
		SchemaVersion:               restrictedAuthorizationSchemaVersion,
		RecordType:                  "restricted-launch-authorization-receipt",
		Source:                      "go-runtime-state-root-restricted-launch-authorization",
		AuthorizationID:             authorizationID,
		ApplicationID:               request.ApplicationID,
		RequestID:                   request.RequestID,
		Mode:                        request.Mode,
		Scope:                       request.Scope,
		Directive:                   request.Directive,
		State:                       "authorized-preparation-only",
		RelativePath:                relativePath,
		PreparationAuthorized:       true,
		LaunchAuthorized:            false,
		ProcessStartAuthorized:      false,
		ExecutionApproved:           false,
		ProductionTrustSatisfied:    false,
		RuntimeWriteGateEnabled:     false,
		ArtifactAcquisitionEnabled:  false,
		BackendInstallEnabled:       false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		RealPortalCallEnabled:       false,
		StateRootPathExposed:        false,
		RawCommandExposed:           false,
		BackendDetailsExposed:       false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		HostRootModified:            false,
		Summary:                     "Runtime recorded explicit permission to prepare a restricted test boundary without authorizing launch, process start, production trust, or Runtime writes.",
	}
	data, digest, err := marshalRestrictedAuthorization(receipt)
	if err != nil {
		return RestrictedAuthorizationReceipt{}, err
	}
	receipt.SHA256 = digest
	data, _, err = marshalRestrictedAuthorization(receipt)
	if err != nil {
		return RestrictedAuthorizationReceipt{}, err
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return RestrictedAuthorizationReceipt{}, fmt.Errorf("write restricted launch authorization receipt: %w", err)
	}
	return receipt, nil
}

func (s *RestrictedAuthorizationStore) Load(authorizationID string) (RestrictedAuthorizationReceipt, error) {
	relativePath, path, err := s.authorizationPath(authorizationID, false)
	if err != nil {
		return RestrictedAuthorizationReceipt{}, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return RestrictedAuthorizationReceipt{}, fmt.Errorf("read restricted launch authorization receipt: %w", err)
	}
	var receipt RestrictedAuthorizationReceipt
	if err := json.Unmarshal(data, &receipt); err != nil {
		return RestrictedAuthorizationReceipt{}, fmt.Errorf("parse restricted launch authorization receipt: %w", err)
	}
	if err := validateRestrictedAuthorizationReceipt(receipt, authorizationID, relativePath); err != nil {
		return RestrictedAuthorizationReceipt{}, err
	}
	return receipt, nil
}

func validateRestrictedAuthorizationRequest(request RestrictedAuthorizationRequest) error {
	if !appid.Valid(request.ApplicationID) {
		return errors.New("restricted launch authorization requires a reverse-DNS application id")
	}
	if err := validateLedgerID(request.RequestID); err != nil {
		return err
	}
	if request.Mode != RestrictedTestMode || request.Scope != RestrictedTestPreparationScope || request.Directive != RestrictedTestPreparationDirective {
		return errors.New("restricted launch authorization requires the exact test-only mode, preparation scope, and authorization directive")
	}
	return nil
}

func validateRestrictedAuthorizationReceipt(receipt RestrictedAuthorizationReceipt, authorizationID string, relativePath string) error {
	if receipt.SchemaVersion != restrictedAuthorizationSchemaVersion || receipt.RecordType != "restricted-launch-authorization-receipt" || receipt.Source != "go-runtime-state-root-restricted-launch-authorization" {
		return errors.New("restricted launch authorization receipt has unsupported schema")
	}
	if receipt.AuthorizationID != authorizationID || receipt.RelativePath != relativePath || !appid.Valid(receipt.ApplicationID) || receipt.RequestID == "" || receipt.Mode != RestrictedTestMode || receipt.Scope != RestrictedTestPreparationScope || receipt.Directive != RestrictedTestPreparationDirective || receipt.State != "authorized-preparation-only" {
		return errors.New("restricted launch authorization receipt identity, path, or scope mismatch")
	}
	storedDigest := receipt.SHA256
	receipt.SHA256 = ""
	_, expectedDigest, err := marshalRestrictedAuthorization(receipt)
	if err != nil {
		return err
	}
	if storedDigest == "" || storedDigest != expectedDigest {
		return errors.New("restricted launch authorization receipt digest mismatch")
	}
	if !receipt.PreparationAuthorized || restrictedAuthorizationUnsafe(receipt) {
		return errors.New("restricted launch authorization receipt has invalid or unsafe gates")
	}
	return nil
}

func restrictedAuthorizationUnsafe(receipt RestrictedAuthorizationReceipt) bool {
	return receipt.LaunchAuthorized || receipt.ProcessStartAuthorized || receipt.ExecutionApproved || receipt.ProductionTrustSatisfied || receipt.RuntimeWriteGateEnabled || receipt.ArtifactAcquisitionEnabled || receipt.BackendInstallEnabled || receipt.BackendLaunchEnabled || receipt.BackendProcessStarted || receipt.RealPortalCallEnabled || receipt.StateRootPathExposed || receipt.RawCommandExposed || receipt.BackendDetailsExposed || receipt.NetworkRequired || receipt.PrivilegedContainerRequired || receipt.HostRootModified
}

func (s *RestrictedAuthorizationStore) authorizationPath(authorizationID string, createDirectory bool) (string, string, error) {
	if strings.TrimSpace(authorizationID) == "" || len(authorizationID) > 240 || strings.Contains(authorizationID, "..") || strings.ContainsAny(authorizationID, `/\\`) {
		return "", "", errors.New("invalid restricted launch authorization id")
	}
	if err := ensureRestrictedAuthorizationDirectory(s.root.Path(), "execution-ledger", createDirectory); err != nil {
		return "", "", err
	}
	if err := ensureRestrictedAuthorizationDirectory(s.root.Path(), filepath.Join("execution-ledger", "authorizations"), createDirectory); err != nil {
		return "", "", err
	}
	relativePath := filepath.ToSlash(filepath.Join("execution-ledger", "authorizations", authorizationID+".json"))
	path, err := s.root.Resolve(relativePath)
	if err != nil {
		return "", "", err
	}
	return relativePath, path, nil
}

func ensureRestrictedAuthorizationDirectory(root string, relativePath string, create bool) error {
	path := filepath.Join(root, filepath.FromSlash(relativePath))
	info, err := os.Lstat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("inspect restricted authorization directory: %w", err)
		}
		if !create {
			return fmt.Errorf("restricted authorization directory does not exist: %s", filepath.ToSlash(relativePath))
		}
		if err := os.Mkdir(path, 0o700); err != nil {
			return fmt.Errorf("prepare restricted authorization directory: %w", err)
		}
		return nil
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return fmt.Errorf("restricted authorization managed path must be a real directory: %s", filepath.ToSlash(relativePath))
	}
	return nil
}

func restrictedAuthorizationID(applicationID string, requestID string) string {
	return "xnix-auth-" + sanitize(applicationID) + "-" + sanitize(requestID)
}

func marshalRestrictedAuthorization(receipt RestrictedAuthorizationReceipt) ([]byte, string, error) {
	data, err := json.MarshalIndent(receipt, "", "  ")
	if err != nil {
		return nil, "", err
	}
	data = append(data, '\n')
	sum := sha256.Sum256(data)
	return data, hex.EncodeToString(sum[:]), nil
}

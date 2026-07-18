package execution

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"xnix.local/xnix/internal/runtime/rootfs"
)

const restrictedMaterializationSchemaVersion = "xnix.runtime.restricted_launch_materialization.v1"

type RestrictedMaterializationRequest struct {
	Authorization RestrictedAuthorizationReceipt
	Preflight     RestrictedPreflightPacket
	Transaction   LedgerRecord
}

// RestrictedMaterializationPlan records a test-only launch materialization
// boundary. It materializes only a reviewable plan and receipt references, never
// a command line, executable path, backend selection, or process start.
type RestrictedMaterializationPlan struct {
	SchemaVersion               string   `json:"schema_version"`
	RecordType                  string   `json:"record_type"`
	Source                      string   `json:"source"`
	PlanID                      string   `json:"plan_id"`
	ApplicationID               string   `json:"application_id"`
	RequestID                   string   `json:"request_id"`
	Mode                        string   `json:"mode"`
	RelativePath                string   `json:"relative_path"`
	SHA256                      string   `json:"sha256"`
	AuthorizationRelativePath   string   `json:"authorization_relative_path"`
	AuthorizationSHA256         string   `json:"authorization_sha256"`
	PreflightRelativePath       string   `json:"preflight_relative_path"`
	PreflightSHA256             string   `json:"preflight_sha256"`
	TransactionRelativePath     string   `json:"transaction_relative_path"`
	TransactionSHA256           string   `json:"transaction_sha256"`
	Status                      string   `json:"status"`
	MaterializationScope        string   `json:"materialization_scope"`
	MaterializedArtifactIDs     []string `json:"materialized_artifact_ids"`
	BlockedByIDs                []string `json:"blocked_by_ids"`
	BlockedByCount              int      `json:"blocked_by_count"`
	SafeInputsReady             bool     `json:"safe_inputs_ready"`
	PreparationAuthorized       bool     `json:"preparation_authorized"`
	PreflightReadBack           bool     `json:"preflight_read_back"`
	PlanMaterialized            bool     `json:"plan_materialized"`
	TestOnly                    bool     `json:"test_only"`
	ProductionTrustSatisfied    bool     `json:"production_trust_satisfied"`
	RuntimeWriteGateEnabled     bool     `json:"runtime_write_gate_enabled"`
	LaunchPreflightPassed       bool     `json:"launch_preflight_passed"`
	LaunchAuthorized            bool     `json:"launch_authorized"`
	ExecutionApproved           bool     `json:"execution_approved"`
	ProcessStartAuthorized      bool     `json:"process_start_authorized"`
	CommandMaterialized         bool     `json:"command_materialized"`
	ExecutablePathResolved      bool     `json:"executable_path_resolved"`
	BackendSelectedForLaunch    bool     `json:"backend_selected_for_launch"`
	BackendLaunchEnabled        bool     `json:"backend_launch_enabled"`
	BackendProcessStarted       bool     `json:"backend_process_started"`
	StateRootPathExposed        bool     `json:"state_root_path_exposed"`
	RawCommandExposed           bool     `json:"raw_command_exposed"`
	RawExecutableExposed        bool     `json:"raw_executable_exposed"`
	BackendDetailsExposed       bool     `json:"backend_details_exposed"`
	NetworkRequired             bool     `json:"network_required"`
	PrivilegedContainerRequired bool     `json:"privileged_container_required"`
	HostRootModified            bool     `json:"host_root_modified"`
	Summary                     string   `json:"summary"`
}

type RestrictedMaterializationStore struct {
	root *rootfs.Root
}

func NewRestrictedMaterializationStore(stateRoot string) (*RestrictedMaterializationStore, error) {
	root, err := rootfs.Open(stateRoot)
	if err != nil {
		return nil, err
	}
	return &RestrictedMaterializationStore{root: root}, nil
}

func (s *RestrictedMaterializationStore) Record(request RestrictedMaterializationRequest) (RestrictedMaterializationPlan, error) {
	if err := validateRestrictedMaterializationRequest(request); err != nil {
		return RestrictedMaterializationPlan{}, err
	}
	planID := "xnix-materialization-" + sanitize(request.Transaction.RequestID)
	relativePath, path, err := s.planPath(planID, true)
	if err != nil {
		return RestrictedMaterializationPlan{}, err
	}
	blockers := append([]string{}, request.Preflight.BlockerIDs...)
	plan := RestrictedMaterializationPlan{
		SchemaVersion:               restrictedMaterializationSchemaVersion,
		RecordType:                  "restricted-launch-materialization-plan",
		Source:                      "go-runtime-state-root-restricted-launch-materialization",
		PlanID:                      planID,
		ApplicationID:               request.Transaction.ApplicationID,
		RequestID:                   request.Transaction.RequestID,
		Mode:                        RestrictedTestMode,
		RelativePath:                relativePath,
		AuthorizationRelativePath:   request.Authorization.RelativePath,
		AuthorizationSHA256:         request.Authorization.SHA256,
		PreflightRelativePath:       request.Preflight.RelativePath,
		PreflightSHA256:             request.Preflight.SHA256,
		TransactionRelativePath:     request.Transaction.RelativePath,
		TransactionSHA256:           request.Transaction.SHA256,
		Status:                      "blocked-plan-materialized",
		MaterializationScope:        "test-only-review-plan",
		MaterializedArtifactIDs:     []string{"launch-intent-reference", "authorization-reference", "preflight-reference", "blocked-transaction-reference"},
		BlockedByIDs:                blockers,
		BlockedByCount:              len(blockers),
		SafeInputsReady:             request.Preflight.SafeInputsReady,
		PreparationAuthorized:       request.Authorization.PreparationAuthorized,
		PreflightReadBack:           true,
		PlanMaterialized:            true,
		TestOnly:                    true,
		ProductionTrustSatisfied:    false,
		RuntimeWriteGateEnabled:     false,
		LaunchPreflightPassed:       false,
		LaunchAuthorized:            false,
		ExecutionApproved:           false,
		ProcessStartAuthorized:      false,
		CommandMaterialized:         false,
		ExecutablePathResolved:      false,
		BackendSelectedForLaunch:    false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		StateRootPathExposed:        false,
		RawCommandExposed:           false,
		RawExecutableExposed:        false,
		BackendDetailsExposed:       false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		HostRootModified:            false,
		Summary:                     "Runtime materialized only a test-only review plan from restricted launch preflight evidence; command, executable path, backend selection, launch, and process start remain blocked.",
	}
	data, digest, err := marshalRestrictedMaterialization(plan)
	if err != nil {
		return RestrictedMaterializationPlan{}, err
	}
	plan.SHA256 = digest
	data, _, err = marshalRestrictedMaterialization(plan)
	if err != nil {
		return RestrictedMaterializationPlan{}, err
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return RestrictedMaterializationPlan{}, fmt.Errorf("write restricted launch materialization plan: %w", err)
	}
	return plan, nil
}

func (s *RestrictedMaterializationStore) Load(planID string) (RestrictedMaterializationPlan, error) {
	relativePath, path, err := s.planPath(planID, false)
	if err != nil {
		return RestrictedMaterializationPlan{}, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return RestrictedMaterializationPlan{}, fmt.Errorf("read restricted launch materialization plan: %w", err)
	}
	var plan RestrictedMaterializationPlan
	if err := json.Unmarshal(data, &plan); err != nil {
		return RestrictedMaterializationPlan{}, fmt.Errorf("parse restricted launch materialization plan: %w", err)
	}
	if err := validateRestrictedMaterializationPlan(plan, planID, relativePath); err != nil {
		return RestrictedMaterializationPlan{}, err
	}
	return plan, nil
}

func validateRestrictedMaterializationRequest(request RestrictedMaterializationRequest) error {
	if request.Authorization.ApplicationID == "" || request.Authorization.ApplicationID != request.Preflight.ApplicationID || request.Authorization.ApplicationID != request.Transaction.ApplicationID {
		return errors.New("restricted launch materialization requires matching application ids")
	}
	if request.Authorization.RequestID != request.Preflight.RequestID || request.Authorization.RequestID != request.Transaction.RequestID {
		return errors.New("restricted launch materialization requires matching request ids")
	}
	if !request.Authorization.PreparationAuthorized || request.Preflight.Status != "blocked" || !request.Preflight.SafeInputsReady || !request.Preflight.PreparationAuthorized || !request.Preflight.ReadyForPacketAssembly {
		return errors.New("restricted launch materialization requires authorized and digest-verified preflight evidence")
	}
	if request.Transaction.Transaction.State != StateBlocked || request.Transaction.LaunchAllowed || request.Transaction.LaunchEnabled || request.Transaction.BackendStarted {
		return errors.New("restricted launch materialization requires a safe blocked transaction receipt")
	}
	if !sameStringSet(request.Preflight.BlockerIDs, []string{"recipe-trust", "runtime-write-gate"}) || restrictedPreflightUnsafe(request.Preflight) || restrictedAuthorizationUnsafe(request.Authorization) {
		return errors.New("restricted launch materialization requires fail-closed blockers and disabled unsafe gates")
	}
	return nil
}

func validateRestrictedMaterializationPlan(plan RestrictedMaterializationPlan, planID string, relativePath string) error {
	if plan.SchemaVersion != restrictedMaterializationSchemaVersion || plan.RecordType != "restricted-launch-materialization-plan" || plan.Source != "go-runtime-state-root-restricted-launch-materialization" {
		return errors.New("restricted launch materialization plan has unsupported schema")
	}
	if plan.PlanID != planID || plan.RelativePath != relativePath || plan.ApplicationID == "" || plan.RequestID == "" || plan.Mode != RestrictedTestMode || plan.Status != "blocked-plan-materialized" || plan.BlockedByCount != len(plan.BlockedByIDs) {
		return errors.New("restricted launch materialization plan identity, path, or count mismatch")
	}
	storedDigest := plan.SHA256
	plan.SHA256 = ""
	_, expectedDigest, err := marshalRestrictedMaterialization(plan)
	if err != nil {
		return err
	}
	if storedDigest == "" || storedDigest != expectedDigest {
		return errors.New("restricted launch materialization plan digest mismatch")
	}
	if !plan.SafeInputsReady || !plan.PreparationAuthorized || !plan.PreflightReadBack || !plan.PlanMaterialized || !plan.TestOnly || !sameStringSet(plan.BlockedByIDs, []string{"recipe-trust", "runtime-write-gate"}) || restrictedMaterializationUnsafe(plan) {
		return errors.New("restricted launch materialization plan has invalid blockers or unsafe gates")
	}
	return nil
}

func restrictedMaterializationUnsafe(plan RestrictedMaterializationPlan) bool {
	return plan.ProductionTrustSatisfied || plan.RuntimeWriteGateEnabled || plan.LaunchPreflightPassed || plan.LaunchAuthorized || plan.ExecutionApproved || plan.ProcessStartAuthorized || plan.CommandMaterialized || plan.ExecutablePathResolved || plan.BackendSelectedForLaunch || plan.BackendLaunchEnabled || plan.BackendProcessStarted || plan.StateRootPathExposed || plan.RawCommandExposed || plan.RawExecutableExposed || plan.BackendDetailsExposed || plan.NetworkRequired || plan.PrivilegedContainerRequired || plan.HostRootModified
}

func (s *RestrictedMaterializationStore) planPath(planID string, create bool) (string, string, error) {
	if planID == "" || len(planID) > 240 || filepath.Base(planID) != planID {
		return "", "", errors.New("invalid restricted launch materialization plan id")
	}
	if err := ensureRestrictedAuthorizationDirectory(s.root.Path(), "execution-ledger", create); err != nil {
		return "", "", err
	}
	if err := ensureRestrictedAuthorizationDirectory(s.root.Path(), filepath.Join("execution-ledger", "materialization-plans"), create); err != nil {
		return "", "", err
	}
	relativePath := filepath.ToSlash(filepath.Join("execution-ledger", "materialization-plans", planID+".json"))
	path, err := s.root.Resolve(relativePath)
	if err != nil {
		return "", "", err
	}
	return relativePath, path, nil
}

func marshalRestrictedMaterialization(plan RestrictedMaterializationPlan) ([]byte, string, error) {
	data, err := json.MarshalIndent(plan, "", "  ")
	if err != nil {
		return nil, "", err
	}
	data = append(data, '\n')
	sum := sha256.Sum256(data)
	return data, hex.EncodeToString(sum[:]), nil
}

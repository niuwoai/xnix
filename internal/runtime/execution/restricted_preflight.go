package execution

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"xnix.local/xnix/internal/runtime/rootfs"
)

const restrictedPreflightSchemaVersion = "xnix.runtime.restricted_launch_preflight.v1"

type RestrictedPreflightRequest struct {
	Authorization     RestrictedAuthorizationReceipt
	Transaction       LedgerRecord
	SafeInputsReady   bool
	ProductImageReady bool
}

// RestrictedPreflightPacket records a fail-closed launch gate evaluation. It
// carries receipt references and blocker ids, never commands or executable paths.
type RestrictedPreflightPacket struct {
	SchemaVersion               string   `json:"schema_version"`
	RecordType                  string   `json:"record_type"`
	Source                      string   `json:"source"`
	PacketID                    string   `json:"packet_id"`
	ApplicationID               string   `json:"application_id"`
	RequestID                   string   `json:"request_id"`
	Mode                        string   `json:"mode"`
	RelativePath                string   `json:"relative_path"`
	SHA256                      string   `json:"sha256"`
	AuthorizationRelativePath   string   `json:"authorization_relative_path"`
	AuthorizationSHA256         string   `json:"authorization_sha256"`
	TransactionRelativePath     string   `json:"transaction_relative_path"`
	TransactionSHA256           string   `json:"transaction_sha256"`
	Status                      string   `json:"status"`
	BlockerIDs                  []string `json:"blocker_ids"`
	BlockerCount                int      `json:"blocker_count"`
	SafeInputsReady             bool     `json:"safe_inputs_ready"`
	PreparationAuthorized       bool     `json:"preparation_authorized"`
	ReadyForPacketAssembly      bool     `json:"ready_for_packet_assembly"`
	ProductImageReady           bool     `json:"product_image_ready"`
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
	BackendDetailsExposed       bool     `json:"backend_details_exposed"`
	NetworkRequired             bool     `json:"network_required"`
	PrivilegedContainerRequired bool     `json:"privileged_container_required"`
	HostRootModified            bool     `json:"host_root_modified"`
	Summary                     string   `json:"summary"`
}

type RestrictedPreflightStore struct {
	root *rootfs.Root
}

func NewRestrictedPreflightStore(stateRoot string) (*RestrictedPreflightStore, error) {
	root, err := rootfs.Open(stateRoot)
	if err != nil {
		return nil, err
	}
	return &RestrictedPreflightStore{root: root}, nil
}

func (s *RestrictedPreflightStore) Record(request RestrictedPreflightRequest) (RestrictedPreflightPacket, error) {
	if err := validateRestrictedPreflightRequest(request); err != nil {
		return RestrictedPreflightPacket{}, err
	}
	packetID := "xnix-preflight-" + sanitize(request.Transaction.RequestID)
	relativePath, path, err := s.packetPath(packetID, true)
	if err != nil {
		return RestrictedPreflightPacket{}, err
	}
	blockerIDs := restrictedPreflightBlockerIDs(request.Transaction.Transaction.Gates)
	packet := RestrictedPreflightPacket{
		SchemaVersion:               restrictedPreflightSchemaVersion,
		RecordType:                  "restricted-launch-preflight-packet",
		Source:                      "go-runtime-state-root-restricted-launch-preflight",
		PacketID:                    packetID,
		ApplicationID:               request.Transaction.ApplicationID,
		RequestID:                   request.Transaction.RequestID,
		Mode:                        RestrictedTestMode,
		RelativePath:                relativePath,
		AuthorizationRelativePath:   request.Authorization.RelativePath,
		AuthorizationSHA256:         request.Authorization.SHA256,
		TransactionRelativePath:     request.Transaction.RelativePath,
		TransactionSHA256:           request.Transaction.SHA256,
		Status:                      "blocked",
		BlockerIDs:                  blockerIDs,
		BlockerCount:                len(blockerIDs),
		SafeInputsReady:             request.SafeInputsReady,
		PreparationAuthorized:       request.Authorization.PreparationAuthorized,
		ReadyForPacketAssembly:      request.SafeInputsReady && request.Authorization.PreparationAuthorized,
		ProductImageReady:           request.ProductImageReady,
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
		BackendDetailsExposed:       false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		HostRootModified:            false,
		Summary:                     "Restricted launch preflight remains blocked by production trust and Runtime write gates; no command, path, backend launch, or process start is authorized.",
	}
	data, digest, err := marshalRestrictedPreflight(packet)
	if err != nil {
		return RestrictedPreflightPacket{}, err
	}
	packet.SHA256 = digest
	data, _, err = marshalRestrictedPreflight(packet)
	if err != nil {
		return RestrictedPreflightPacket{}, err
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return RestrictedPreflightPacket{}, fmt.Errorf("write restricted launch preflight packet: %w", err)
	}
	return packet, nil
}

func (s *RestrictedPreflightStore) Load(packetID string) (RestrictedPreflightPacket, error) {
	relativePath, path, err := s.packetPath(packetID, false)
	if err != nil {
		return RestrictedPreflightPacket{}, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return RestrictedPreflightPacket{}, fmt.Errorf("read restricted launch preflight packet: %w", err)
	}
	var packet RestrictedPreflightPacket
	if err := json.Unmarshal(data, &packet); err != nil {
		return RestrictedPreflightPacket{}, fmt.Errorf("parse restricted launch preflight packet: %w", err)
	}
	if err := validateRestrictedPreflightPacket(packet, packetID, relativePath); err != nil {
		return RestrictedPreflightPacket{}, err
	}
	return packet, nil
}

func validateRestrictedPreflightRequest(request RestrictedPreflightRequest) error {
	if request.Authorization.ApplicationID == "" || request.Authorization.ApplicationID != request.Transaction.ApplicationID || request.Authorization.RequestID != request.Transaction.RequestID || !request.Authorization.PreparationAuthorized {
		return errors.New("restricted launch preflight requires matching preparation authorization and transaction receipts")
	}
	if request.Transaction.Transaction.State != StateBlocked || request.Transaction.LaunchAllowed || request.Transaction.LaunchEnabled || request.Transaction.BackendStarted {
		return errors.New("restricted launch preflight requires a safe blocked transaction receipt")
	}
	if !request.SafeInputsReady {
		return errors.New("restricted launch preflight requires converged safe input evidence")
	}
	return nil
}

func validateRestrictedPreflightPacket(packet RestrictedPreflightPacket, packetID string, relativePath string) error {
	if packet.SchemaVersion != restrictedPreflightSchemaVersion || packet.RecordType != "restricted-launch-preflight-packet" || packet.Source != "go-runtime-state-root-restricted-launch-preflight" {
		return errors.New("restricted launch preflight packet has unsupported schema")
	}
	if packet.PacketID != packetID || packet.RelativePath != relativePath || packet.ApplicationID == "" || packet.RequestID == "" || packet.Mode != RestrictedTestMode || packet.Status != "blocked" || packet.BlockerCount != len(packet.BlockerIDs) {
		return errors.New("restricted launch preflight packet identity, path, or count mismatch")
	}
	storedDigest := packet.SHA256
	packet.SHA256 = ""
	_, expectedDigest, err := marshalRestrictedPreflight(packet)
	if err != nil {
		return err
	}
	if storedDigest == "" || storedDigest != expectedDigest {
		return errors.New("restricted launch preflight packet digest mismatch")
	}
	if !packet.SafeInputsReady || !packet.PreparationAuthorized || !packet.ReadyForPacketAssembly || packet.ProductImageReady || !sameStringSet(packet.BlockerIDs, []string{"recipe-trust", "runtime-write-gate"}) || restrictedPreflightUnsafe(packet) {
		return errors.New("restricted launch preflight packet has invalid blockers or unsafe gates")
	}
	return nil
}

func restrictedPreflightUnsafe(packet RestrictedPreflightPacket) bool {
	return packet.ProductionTrustSatisfied || packet.RuntimeWriteGateEnabled || packet.LaunchPreflightPassed || packet.LaunchAuthorized || packet.ExecutionApproved || packet.ProcessStartAuthorized || packet.CommandMaterialized || packet.ExecutablePathResolved || packet.BackendSelectedForLaunch || packet.BackendLaunchEnabled || packet.BackendProcessStarted || packet.StateRootPathExposed || packet.RawCommandExposed || packet.BackendDetailsExposed || packet.NetworkRequired || packet.PrivilegedContainerRequired || packet.HostRootModified
}

func restrictedPreflightBlockerIDs(gates []Gate) []string {
	ids := make([]string, 0, len(gates))
	for _, gate := range gates {
		if gate.Status != GatePass {
			ids = append(ids, gate.ID)
		}
	}
	sort.Strings(ids)
	return ids
}

func (s *RestrictedPreflightStore) packetPath(packetID string, create bool) (string, string, error) {
	if packetID == "" || len(packetID) > 240 || filepath.Base(packetID) != packetID {
		return "", "", errors.New("invalid restricted launch preflight packet id")
	}
	if err := ensureRestrictedAuthorizationDirectory(s.root.Path(), "execution-ledger", create); err != nil {
		return "", "", err
	}
	if err := ensureRestrictedAuthorizationDirectory(s.root.Path(), filepath.Join("execution-ledger", "preflight-packets"), create); err != nil {
		return "", "", err
	}
	relativePath := filepath.ToSlash(filepath.Join("execution-ledger", "preflight-packets", packetID+".json"))
	path, err := s.root.Resolve(relativePath)
	if err != nil {
		return "", "", err
	}
	return relativePath, path, nil
}

func sameStringSet(left []string, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	l := append([]string{}, left...)
	r := append([]string{}, right...)
	sort.Strings(l)
	sort.Strings(r)
	for index := range l {
		if l[index] != r[index] {
			return false
		}
	}
	return true
}

func marshalRestrictedPreflight(packet RestrictedPreflightPacket) ([]byte, string, error) {
	data, err := json.MarshalIndent(packet, "", "  ")
	if err != nil {
		return nil, "", err
	}
	data = append(data, '\n')
	sum := sha256.Sum256(data)
	return data, hex.EncodeToString(sum[:]), nil
}

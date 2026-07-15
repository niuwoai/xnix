package appidentity

import (
	"os"
	"strings"
)

type RuntimeOwnerProcessPreview struct {
	Version                     string                     `json:"version"`
	SchemaVersion               string                     `json:"schema_version"`
	RequestType                 string                     `json:"request_type"`
	ProcessType                 string                     `json:"process_type"`
	Source                      string                     `json:"source"`
	RuntimeMethod               string                     `json:"runtime_method"`
	ReadMethod                  string                     `json:"read_method"`
	BusName                     string                     `json:"bus_name"`
	ObjectPath                  string                     `json:"object_path"`
	Interface                   string                     `json:"interface"`
	CurrentOwnerEntrypoint      string                     `json:"current_owner_entrypoint"`
	CurrentOwnerLanguage        string                     `json:"current_owner_language"`
	TargetOwnerLanguage         string                     `json:"target_owner_language"`
	Checks                      []RuntimeOwnerProcessCheck `json:"checks"`
	CheckIDs                    []string                   `json:"check_ids"`
	Counts                      RuntimeOwnerProcessCounts  `json:"counts"`
	ServiceActivationReady      bool                       `json:"service_activation_ready"`
	PackagedEntrypointReady     bool                       `json:"packaged_entrypoint_ready"`
	GoOwnerProcessReady         bool                       `json:"go_owner_process_ready"`
	ProductionOwnerProcessReady bool                       `json:"production_owner_process_ready"`
	RuntimeOwned                bool                       `json:"runtime_owned"`
	GoRuntimeBacked             bool                       `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                       `json:"kde_policy_owner"`
	KDEMayClaimRuntimeOwnership bool                       `json:"kde_may_claim_runtime_ownership"`
	SystemServiceStarted        bool                       `json:"system_service_started"`
	ProductionBusClaimed        bool                       `json:"production_bus_claimed"`
	NetworkRequired             bool                       `json:"network_required"`
	HostRootModified            bool                       `json:"host_root_modified"`
	PrivilegedContainerRequired bool                       `json:"privileged_container_required"`
	BackendDetailsExposed       bool                       `json:"backend_details_exposed"`
	BlockedActions              []string                   `json:"blocked_actions"`
	NextRequirements            []string                   `json:"next_requirements"`
	DesktopSafeSummary          string                     `json:"desktop_safe_summary"`
}

type RuntimeOwnerProcessCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type RuntimeOwnerProcessCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewRuntimeOwnerProcessPreview(root string) (RuntimeOwnerProcessPreview, error) {
	serviceBinding, err := NewRuntimeServiceBindingPreview(root)
	if err != nil {
		return RuntimeOwnerProcessPreview{}, err
	}

	wrapperContent := readRuntimeServiceBindingOptional(root, runtimeServiceBindingLibexecWrapper)
	checks := runtimeOwnerProcessChecks(root, serviceBinding.ActivationBindingReady, wrapperContent)
	counts := countRuntimeOwnerProcessChecks(checks)
	entrypointReady := runtimeOwnerProcessCheckStatus(checks, "packaged-entrypoint") == "pass"
	goOwnerReady := runtimeOwnerProcessCheckStatus(checks, "go-owner-target") == "pass"

	preview := RuntimeOwnerProcessPreview{
		Version:                     serviceBinding.Version,
		SchemaVersion:               "xnix.runtime.owner_process.v1",
		RequestType:                 "runtime-owner-process-preview",
		ProcessType:                 "runtime-owner-process",
		Source:                      "runtime-service-binding-preview+libexec-wrapper+go-owner-target",
		RuntimeMethod:               "GetRuntimeOwnerProcess",
		ReadMethod:                  "GetRuntimeOwnerProcessPreview",
		BusName:                     serviceBinding.BusName,
		ObjectPath:                  serviceBinding.ObjectPath,
		Interface:                   serviceBinding.Interface,
		CurrentOwnerEntrypoint:      runtimeServiceBindingLibexecWrapper,
		CurrentOwnerLanguage:        runtimeOwnerProcessCurrentLanguage(wrapperContent),
		TargetOwnerLanguage:         "go",
		Checks:                      checks,
		CheckIDs:                    runtimeOwnerProcessCheckIDs(checks),
		Counts:                      counts,
		ServiceActivationReady:      serviceBinding.ActivationBindingReady,
		PackagedEntrypointReady:     entrypointReady,
		GoOwnerProcessReady:         goOwnerReady,
		ProductionOwnerProcessReady: false,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		KDEMayClaimRuntimeOwnership: false,
		SystemServiceStarted:        false,
		ProductionBusClaimed:        false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		BlockedActions:              runtimeOwnerProcessBlockedActions(),
		NextRequirements:            runtimeOwnerProcessNextRequirements(),
		DesktopSafeSummary:          runtimeOwnerProcessSummary(serviceBinding.ActivationBindingReady, entrypointReady, goOwnerReady),
	}
	if err := validateNoBackendTerms(preview, "Runtime owner process preview"); err != nil {
		return RuntimeOwnerProcessPreview{}, err
	}
	return preview, nil
}

func runtimeOwnerProcessChecks(root string, activationReady bool, wrapperContent string) []RuntimeOwnerProcessCheck {
	return []RuntimeOwnerProcessCheck{
		runtimeOwnerProcessCheck("service-activation", runtimeOwnerProcessPassBlocked(activationReady), "D-Bus activation and systemd service files must point at the packaged Runtime entrypoint."),
		runtimeOwnerProcessCheck("packaged-entrypoint", runtimeOwnerProcessPassBlocked(runtimeOwnerProcessEntrypointReady(root, wrapperContent)), "The packaged Runtime entrypoint must exist, be executable, and delegate to the current Runtime daemon."),
		runtimeOwnerProcessCheck("go-owner-target", runtimeOwnerProcessPassPending(runtimeOwnerProcessGoCandidateReady(root)), "A Go Runtime owner candidate should exist before the packaged entrypoint can move away from the current wrapper."),
		runtimeOwnerProcessCheck("production-owner-loop", "pending", "The Go owner still needs an event loop that owns the stable D-Bus name and serves read-only Runtime methods."),
		runtimeOwnerProcessCheck("host-safety-boundary", "pass", "Owner process preview must not start services, claim bus names, require network, or mutate the host root."),
	}
}

func runtimeOwnerProcessCheck(id string, status string, summary string) RuntimeOwnerProcessCheck {
	return RuntimeOwnerProcessCheck{
		ID:      id,
		Status:  status,
		Summary: summary,
	}
}

func runtimeOwnerProcessPassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func runtimeOwnerProcessPassPending(passed bool) string {
	if passed {
		return "pass"
	}
	return "pending"
}

func runtimeOwnerProcessEntrypointReady(root string, wrapperContent string) bool {
	info, err := os.Stat(runtimeServicePath(root, runtimeServiceBindingLibexecWrapper))
	return err == nil && !info.IsDir() && info.Mode().Perm()&0o111 != 0 && strings.Contains(wrapperContent, "runtime_daemon")
}

func runtimeOwnerProcessGoCandidateReady(root string) bool {
	ownerCommand := readRuntimeServiceBindingOptional(root, "cmd/xnix-runtime-owner/main.go")
	ownerCandidate := readRuntimeServiceBindingOptional(root, "internal/runtime/owner/candidate.go")
	return strings.Contains(ownerCommand, "xnix-runtime-owner") &&
		strings.Contains(ownerCommand, "NewCandidate") &&
		strings.Contains(ownerCommand, "DisabledWriteResponse") &&
		strings.Contains(ownerCandidate, "runtime-owner-candidate") &&
		strings.Contains(ownerCandidate, "go-runtime-owner-candidate") &&
		strings.Contains(ownerCandidate, "ModeSmokeOwner")
}

func runtimeOwnerProcessCurrentLanguage(wrapperContent string) string {
	if strings.Contains(wrapperContent, "runtime_daemon") {
		return "ruby-wrapper"
	}
	return "unknown"
}

func runtimeOwnerProcessCheckIDs(checks []RuntimeOwnerProcessCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countRuntimeOwnerProcessChecks(checks []RuntimeOwnerProcessCheck) RuntimeOwnerProcessCounts {
	counts := RuntimeOwnerProcessCounts{Total: len(checks)}
	for _, check := range checks {
		switch check.Status {
		case "pass":
			counts.Passed++
		case "pending":
			counts.Pending++
		case "blocked":
			counts.Blocked++
		}
	}
	return counts
}

func runtimeOwnerProcessCheckStatus(checks []RuntimeOwnerProcessCheck, id string) string {
	for _, check := range checks {
		if check.ID == id {
			return check.Status
		}
	}
	return "blocked"
}

func runtimeOwnerProcessBlockedActions() []string {
	return []string{
		"start production Runtime owner from process preview",
		"claim production D-Bus name from process preview",
		"treat the current wrapper as the final Go owner",
		"let KDE claim Runtime ownership",
		"enable write methods before the Go owner process exists",
		"mutate host root during owner process preview",
	}
}

func runtimeOwnerProcessNextRequirements() []string {
	return []string{
		"Bind the Go owner candidate to a restricted session-bus smoke.",
		"Bind the packaged entrypoint to the Go owner after owner smoke parity passes.",
		"Prove stable D-Bus name ownership in a restricted production-owner smoke.",
		"Keep the current wrapper available only as a transition path until the production owner loop is ready.",
	}
}

func runtimeOwnerProcessSummary(activationReady bool, entrypointReady bool, goOwnerReady bool) string {
	if !activationReady || !entrypointReady {
		return "Runtime owner process activation is blocked by service or entrypoint defects."
	}
	if goOwnerReady {
		return "Runtime owner process has a Go owner candidate, but the production D-Bus event loop remains pending."
	}
	return "Runtime owner process activation is aligned, but the production Go owner loop remains pending."
}

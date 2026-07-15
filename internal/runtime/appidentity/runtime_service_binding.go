package appidentity

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	runtimeServiceBindingBusName          = "org.xnix.Compatibility1"
	runtimeServiceBindingObjectPath       = "/org/xnix/Compatibility1"
	runtimeServiceBindingInterface        = "org.xnix.Compatibility1"
	runtimeServiceBindingDBusServiceFile  = "runtime/dbus/org.xnix.Compatibility1.service"
	runtimeServiceBindingSystemdUnit      = "runtime/systemd/xnix-compatd.service"
	runtimeServiceBindingLibexecWrapper   = "libexec/xnix/compatd"
	runtimeServiceBindingDBusContract     = "runtime/dbus/org.xnix.Compatibility1.xml"
	runtimeServiceBindingSmokeAdapter     = "runtime/dbus/xnix_compatd_smoke.c"
	runtimeServiceBindingPackagedWrapper  = "/usr/libexec/xnix/compatd"
	runtimeServiceBindingProductionStatus = "pending-live-owner"
)

type RuntimeServiceBindingPreview struct {
	Version                     string                       `json:"version"`
	SchemaVersion               string                       `json:"schema_version"`
	RequestType                 string                       `json:"request_type"`
	BindingType                 string                       `json:"binding_type"`
	Source                      string                       `json:"source"`
	RuntimeMethod               string                       `json:"runtime_method"`
	ReadMethod                  string                       `json:"read_method"`
	ProductionStatus            string                       `json:"production_status"`
	BusName                     string                       `json:"bus_name"`
	ObjectPath                  string                       `json:"object_path"`
	Interface                   string                       `json:"interface"`
	Activation                  RuntimeServiceActivation     `json:"activation"`
	Checks                      []RuntimeServiceBindingCheck `json:"checks"`
	CheckIDs                    []string                     `json:"check_ids"`
	Counts                      RuntimeServiceBindingCounts  `json:"counts"`
	RuntimeOwned                bool                         `json:"runtime_owned"`
	GoRuntimeBacked             bool                         `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                         `json:"kde_policy_owner"`
	ActivationBindingReady      bool                         `json:"activation_binding_ready"`
	LiveDBusOwnerReady          bool                         `json:"live_dbus_owner_ready"`
	SmokeAdapterAvailable       bool                         `json:"smoke_adapter_available"`
	ServiceStarted              bool                         `json:"service_started"`
	ProductionBusClaimed        bool                         `json:"production_bus_claimed"`
	KDEMayClaimRuntimeOwnership bool                         `json:"kde_may_claim_runtime_ownership"`
	NetworkRequired             bool                         `json:"network_required"`
	HostRootModified            bool                         `json:"host_root_modified"`
	PrivilegedContainerRequired bool                         `json:"privileged_container_required"`
	BackendDetailsExposed       bool                         `json:"backend_details_exposed"`
	BlockedActions              []string                     `json:"blocked_actions"`
	DesktopSafeSummary          string                       `json:"desktop_safe_summary"`
}

type RuntimeServiceActivation struct {
	DBusServiceFile string `json:"dbus_service_file"`
	SystemdUnit     string `json:"systemd_unit"`
	LibexecWrapper  string `json:"libexec_wrapper"`
	DBusContract    string `json:"dbus_contract"`
	PackagedWrapper string `json:"packaged_wrapper"`
}

type RuntimeServiceBindingCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type RuntimeServiceBindingCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewRuntimeServiceBindingPreview(root string) (RuntimeServiceBindingPreview, error) {
	if root == "" {
		root = "."
	}
	version, err := readRuntimeServiceBindingVersion(root)
	if err != nil {
		return RuntimeServiceBindingPreview{}, err
	}

	checks := []RuntimeServiceBindingCheck{
		runtimeServiceDBusActivationCheck(root),
		runtimeServiceSystemdHardeningCheck(root),
		runtimeServiceLibexecWrapperCheck(root),
		runtimeServiceContractCheck(root),
		runtimeServiceLiveOwnerCheck(),
	}
	counts := countRuntimeServiceBindingChecks(checks)
	activationReady := runtimeServiceActivationReady(checks)

	preview := RuntimeServiceBindingPreview{
		Version:          version,
		SchemaVersion:    "xnix.runtime.service_binding.v1",
		RequestType:      "runtime-service-binding-preview",
		BindingType:      "runtime-service-binding",
		Source:           "activation-files+dbus-contract+runtime-owner-gate",
		RuntimeMethod:    "GetRuntimeServiceBinding",
		ReadMethod:       "GetRuntimeServiceBindingPreview",
		ProductionStatus: runtimeServiceBindingProductionStatus,
		BusName:          runtimeServiceBindingBusName,
		ObjectPath:       runtimeServiceBindingObjectPath,
		Interface:        runtimeServiceBindingInterface,
		Activation: RuntimeServiceActivation{
			DBusServiceFile: runtimeServiceBindingDBusServiceFile,
			SystemdUnit:     runtimeServiceBindingSystemdUnit,
			LibexecWrapper:  runtimeServiceBindingLibexecWrapper,
			DBusContract:    runtimeServiceBindingDBusContract,
			PackagedWrapper: runtimeServiceBindingPackagedWrapper,
		},
		Checks:                      checks,
		CheckIDs:                    runtimeServiceBindingCheckIDs(checks),
		Counts:                      counts,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		ActivationBindingReady:      activationReady,
		LiveDBusOwnerReady:          false,
		SmokeAdapterAvailable:       runtimeServicePathExists(root, runtimeServiceBindingSmokeAdapter),
		ServiceStarted:              false,
		ProductionBusClaimed:        false,
		KDEMayClaimRuntimeOwnership: false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"start Runtime service from preview",
			"claim production D-Bus owner from preview",
			"let KDE claim Runtime ownership",
			"mutate host root during service binding planning",
		},
		DesktopSafeSummary: runtimeServiceBindingSummary(activationReady),
	}
	if err := validateNoBackendTerms(preview, "Runtime service binding preview"); err != nil {
		return RuntimeServiceBindingPreview{}, err
	}
	return preview, nil
}

func runtimeServiceDBusActivationCheck(root string) RuntimeServiceBindingCheck {
	content := readRuntimeServiceBindingOptional(root, runtimeServiceBindingDBusServiceFile)
	passed := strings.Contains(content, "Name="+runtimeServiceBindingBusName) &&
		strings.Contains(content, "SystemdService=xnix-compatd.service") &&
		strings.Contains(content, "Exec="+runtimeServiceBindingPackagedWrapper)
	return runtimeServiceBindingCheck("dbus-service-activation", runtimeServiceCheckStatus(passed), "D-Bus activation points to the packaged Runtime wrapper.")
}

func runtimeServiceSystemdHardeningCheck(root string) RuntimeServiceBindingCheck {
	content := readRuntimeServiceBindingOptional(root, runtimeServiceBindingSystemdUnit)
	passed := strings.Contains(content, "Type=dbus") &&
		strings.Contains(content, "BusName="+runtimeServiceBindingBusName) &&
		strings.Contains(content, "ExecStart="+runtimeServiceBindingPackagedWrapper) &&
		strings.Contains(content, "NoNewPrivileges=yes") &&
		strings.Contains(content, "ProtectSystem=strict")
	return runtimeServiceBindingCheck("systemd-service-hardening", runtimeServiceCheckStatus(passed), "systemd activation uses the stable bus name and hardened service settings.")
}

func runtimeServiceLibexecWrapperCheck(root string) RuntimeServiceBindingCheck {
	content := readRuntimeServiceBindingOptional(root, runtimeServiceBindingLibexecWrapper)
	info, err := os.Stat(runtimeServicePath(root, runtimeServiceBindingLibexecWrapper))
	passed := err == nil && !info.IsDir() && info.Mode().Perm()&0o111 != 0 && strings.Contains(content, "runtime_daemon")
	return runtimeServiceBindingCheck("libexec-wrapper", runtimeServiceCheckStatus(passed), "Packaged Runtime wrapper delegates to the Runtime daemon entry point.")
}

func runtimeServiceContractCheck(root string) RuntimeServiceBindingCheck {
	content := readRuntimeServiceBindingOptional(root, runtimeServiceBindingDBusContract)
	passed := strings.Contains(content, runtimeServiceBindingInterface) &&
		strings.Contains(content, "ListApplications") &&
		strings.Contains(content, "GetRuntimeServiceBinding")
	return runtimeServiceBindingCheck("dbus-contract", runtimeServiceCheckStatus(passed), "D-Bus contract exposes read-only Runtime service binding status.")
}

func runtimeServiceLiveOwnerCheck() RuntimeServiceBindingCheck {
	return runtimeServiceBindingCheck("live-dbus-owner", "pending", "A long-running production D-Bus owner is still pending.")
}

func runtimeServiceBindingCheck(id string, status string, summary string) RuntimeServiceBindingCheck {
	return RuntimeServiceBindingCheck{
		ID:      id,
		Status:  status,
		Summary: summary,
	}
}

func runtimeServiceCheckStatus(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func countRuntimeServiceBindingChecks(checks []RuntimeServiceBindingCheck) RuntimeServiceBindingCounts {
	counts := RuntimeServiceBindingCounts{Total: len(checks)}
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

func runtimeServiceBindingCheckIDs(checks []RuntimeServiceBindingCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func runtimeServiceActivationReady(checks []RuntimeServiceBindingCheck) bool {
	for _, check := range checks {
		if check.ID == "live-dbus-owner" {
			continue
		}
		if check.Status != "pass" {
			return false
		}
	}
	return true
}

func runtimeServiceBindingSummary(activationReady bool) string {
	if activationReady {
		return "Runtime service activation files are aligned; live D-Bus ownership remains pending."
	}
	return "Runtime service activation files need repair before production D-Bus ownership can be enabled."
}

func readRuntimeServiceBindingVersion(root string) (string, error) {
	content, err := os.ReadFile(runtimeServicePath(root, "VERSION"))
	if err != nil {
		return "", fmt.Errorf("read VERSION: %w", err)
	}
	version := strings.TrimSpace(string(content))
	if version == "" {
		return "", fmt.Errorf("read VERSION: empty version")
	}
	return version, nil
}

func readRuntimeServiceBindingOptional(root string, relativePath string) string {
	content, err := os.ReadFile(runtimeServicePath(root, relativePath))
	if err != nil {
		return ""
	}
	return string(content)
}

func runtimeServicePathExists(root string, relativePath string) bool {
	info, err := os.Stat(runtimeServicePath(root, relativePath))
	return err == nil && !info.IsDir()
}

func runtimeServicePath(root string, relativePath string) string {
	return filepath.Join(root, filepath.FromSlash(relativePath))
}

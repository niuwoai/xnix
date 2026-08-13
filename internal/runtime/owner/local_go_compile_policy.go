package owner

import "strings"

type LocalGoCompilePolicyPreview struct {
	Version                            string                      `json:"version"`
	SchemaVersion                      string                      `json:"schema_version"`
	RequestType                        string                      `json:"request_type"`
	PolicyType                         string                      `json:"policy_type"`
	Source                             string                      `json:"source"`
	PolicyDecision                     string                      `json:"policy_decision"`
	DefaultRemoteHost                  string                      `json:"default_remote_host"`
	RemoteBuildCommand                 string                      `json:"remote_build_command"`
	RemoteTestCommand                  string                      `json:"remote_test_command"`
	LocalOverrideEnvironment           string                      `json:"local_override_environment"`
	LocalGoCompilationAllowedByDefault bool                        `json:"local_go_compilation_allowed_by_default"`
	LocalGoRunAllowedByDefault         bool                        `json:"local_go_run_allowed_by_default"`
	Q4CompilationRequiredByDefault     bool                        `json:"q4_compilation_required_by_default"`
	RemoteGoBuildRunnerPresent         bool                        `json:"remote_go_build_runner_present"`
	RemoteGoTestRunnerPresent          bool                        `json:"remote_go_test_runner_present"`
	RemoteGoCachesPinnedToQ4           bool                        `json:"remote_go_caches_pinned_to_q4"`
	TargetedSmallVersionTesting        bool                        `json:"targeted_small_version_testing"`
	TwentiethVersionFullGateRequired   bool                        `json:"twentieth_version_full_gate_required"`
	ProtectedClaudePackageExcluded     bool                        `json:"protected_claude_package_excluded"`
	HostCompilationAvoidedObservable   bool                        `json:"host_compilation_avoided_observable"`
	Q4HostRootModified                 bool                        `json:"q4_host_root_modified"`
	HostRootModified                   bool                        `json:"host_root_modified"`
	PrivilegedContainerRequired        bool                        `json:"privileged_container_required"`
	HostNetworkingRequired             bool                        `json:"host_networking_required"`
	DockerSocketMounted                bool                        `json:"docker_socket_mounted"`
	BroadHostMountRequired             bool                        `json:"broad_host_mount_required"`
	Checks                             []LocalGoCompilePolicyCheck `json:"checks"`
	CheckIDs                           []string                    `json:"check_ids"`
	Counts                             LocalGoCompilePolicyCounts  `json:"counts"`
	BlockedActions                     []string                    `json:"blocked_actions"`
	NextRequirements                   []string                    `json:"next_requirements"`
	DesktopSafeSummary                 string                      `json:"desktop_safe_summary"`
}

type LocalGoCompilePolicyCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type LocalGoCompilePolicyCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewLocalGoCompilePolicyPreview(root string) (LocalGoCompilePolicyPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return LocalGoCompilePolicyPreview{}, err
	}
	policySource := productionAuthorizationReadSources(root, []string{"docs/q4-remote-build-policy.md"})
	claudeSource := productionAuthorizationReadSources(root, []string{"CLAUDE.md"})
	remoteBuildSource := productionAuthorizationReadSources(root, []string{"scripts/remote_go_build.rb"})
	remoteTestSource := productionAuthorizationReadSources(root, []string{"scripts/remote_go_test.rb"})

	policyReady := allContain(policySource, []string{
		"Compile-heavy work runs on q4 by default",
		"Do not run local Go compilation",
		"XNIX_ALLOW_LOCAL_GO_COMPILE=1",
		"Every new small version is targeted-test only unless a twentieth-version full gate is explicitly authorized.",
	})
	claudeReady := allContain(claudeSource, []string{
		"Heavy build default",
		"Local compile guard",
		"Use q4 remote build/test commands",
		"XNIX_ALLOW_LOCAL_GO_COMPILE=1",
	})
	remoteBuildReady := allContain(remoteBuildSource, []string{
		"root@q4",
		"host_compilation_avoided",
		"remote_build_planned",
		"GOCACHE=",
		"GOMODCACHE=",
		"GOTMPDIR=",
		"docs/claude-code-implementation-packages.md",
	})
	remoteTestReady := allContain(remoteTestSource, []string{
		"root@q4",
		"host_compilation_avoided",
		"targeted_test_required",
		"GOCACHE=",
		"GOMODCACHE=",
		"GOTMPDIR=",
		"docs/claude-code-implementation-packages.md",
	})
	cacheReady := allContain(remoteBuildSource+remoteTestSource, []string{"GOCACHE=", "GOMODCACHE=", "GOTMPDIR="})
	claudeExcluded := strings.Contains(remoteBuildSource, "docs/claude-code-implementation-packages.md") && strings.Contains(remoteTestSource, "docs/claude-code-implementation-packages.md")
	fullGateReady := strings.Contains(policySource, "No full-smoke substitution with a narrow test at twentieth-version gates.") && strings.Contains(claudeSource, "every 20 code versions")
	ready := policyReady && claudeReady && remoteBuildReady && remoteTestReady && cacheReady && claudeExcluded && fullGateReady

	preview := LocalGoCompilePolicyPreview{
		Version:                            version,
		SchemaVersion:                      "xnix.runtime.local_go_compile_policy.v1",
		RequestType:                        "local-go-compile-policy-preview",
		PolicyType:                         "q4-first-go-compilation-policy",
		Source:                             "q4-remote-build-policy+claude-guide+remote-go-runners",
		PolicyDecision:                     "local-go-compilation-policy-blocked",
		DefaultRemoteHost:                  "root@q4",
		RemoteBuildCommand:                 "ruby scripts/remote_go_build.rb --execute",
		RemoteTestCommand:                  "ruby scripts/remote_go_test.rb --execute",
		LocalOverrideEnvironment:           "XNIX_ALLOW_LOCAL_GO_COMPILE=1",
		LocalGoCompilationAllowedByDefault: false,
		LocalGoRunAllowedByDefault:         false,
		Q4CompilationRequiredByDefault:     policyReady && claudeReady,
		RemoteGoBuildRunnerPresent:         remoteBuildReady,
		RemoteGoTestRunnerPresent:          remoteTestReady,
		RemoteGoCachesPinnedToQ4:           cacheReady,
		TargetedSmallVersionTesting:        strings.Contains(remoteTestSource, "targeted_test_required"),
		TwentiethVersionFullGateRequired:   fullGateReady,
		ProtectedClaudePackageExcluded:     claudeExcluded,
		HostCompilationAvoidedObservable:   strings.Contains(remoteBuildSource+remoteTestSource, "host_compilation_avoided"),
		Q4HostRootModified:                 false,
		HostRootModified:                   false,
		PrivilegedContainerRequired:        false,
		HostNetworkingRequired:             false,
		DockerSocketMounted:                false,
		BroadHostMountRequired:             false,
		BlockedActions: []string{
			"run go build, go test, go run, QEMU, Wine, or heavy smoke on the macOS host without a one-command human override",
			"sync docs/claude-code-implementation-packages.md, secrets, private keys, local caches, or build artifacts to q4",
			"replace a twentieth-version full gate with a narrow targeted test",
		},
		NextRequirements: []string{
			"Use q4 remote build and test runners for every compile-heavy Runtime validation by default.",
			"Keep Ruby wrappers as harnesses that report host_compilation_avoided instead of embedding Runtime business decisions.",
			"Only use XNIX_ALLOW_LOCAL_GO_COMPILE=1 for one explicitly approved local command.",
		},
		DesktopSafeSummary: "The local Go compile policy keeps the macOS host for editing and review while q4 owns compile-heavy Runtime validation; local Go compilation is disabled by default and only allowed with a one-command human override.",
	}
	checks := localGoCompilePolicyChecks(preview, policyReady, claudeReady, remoteBuildReady, remoteTestReady, cacheReady, claudeExcluded, fullGateReady)
	preview.Checks = checks
	preview.CheckIDs = localGoCompilePolicyCheckIDs(checks)
	preview.Counts = countLocalGoCompilePolicyChecks(checks)
	if ready {
		preview.PolicyDecision = "local-go-compilation-policy-ready-q4-first"
	}
	return preview, nil
}

func localGoCompilePolicyChecks(preview LocalGoCompilePolicyPreview, policyReady bool, claudeReady bool, remoteBuildReady bool, remoteTestReady bool, cacheReady bool, claudeExcluded bool, fullGateReady bool) []LocalGoCompilePolicyCheck {
	return []LocalGoCompilePolicyCheck{
		localGoCompilePolicyCheck("q4-policy-document-present", productionAuthorizationPassBlocked(policyReady), "The q4 remote build policy disables local Go compilation by default and names the q4 override workflow."),
		localGoCompilePolicyCheck("claude-guide-local-compile-guard-present", productionAuthorizationPassBlocked(claudeReady), "The contributor guide points compile-heavy work to q4 and documents the local override environment variable."),
		localGoCompilePolicyCheck("remote-go-build-runner-ready", productionAuthorizationPassBlocked(remoteBuildReady && preview.RemoteGoBuildRunnerPresent), "The remote Go build runner defaults to q4 and exposes host compilation avoidance."),
		localGoCompilePolicyCheck("remote-go-test-runner-ready", productionAuthorizationPassBlocked(remoteTestReady && preview.RemoteGoTestRunnerPresent), "The remote Go test runner defaults to q4, uses targeted tests, and exposes host compilation avoidance."),
		localGoCompilePolicyCheck("remote-go-cache-roots-pinned", productionAuthorizationPassBlocked(cacheReady && preview.RemoteGoCachesPinnedToQ4), "Go build, module, and temporary caches stay on the remote build host."),
		localGoCompilePolicyCheck("protected-claude-package-excluded", productionAuthorizationPassBlocked(claudeExcluded && preview.ProtectedClaudePackageExcluded), "The protected Claude implementation package document is excluded from remote source sync."),
		localGoCompilePolicyCheck("twentieth-version-full-gate-preserved", productionAuthorizationPassBlocked(fullGateReady && preview.TwentiethVersionFullGateRequired), "Targeted small-version tests do not replace twentieth-version full gates."),
		localGoCompilePolicyCheck("unsafe-host-gates-closed", productionAuthorizationPassBlocked(!preview.LocalGoCompilationAllowedByDefault && !preview.LocalGoRunAllowedByDefault && !preview.Q4HostRootModified && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.HostNetworkingRequired && !preview.DockerSocketMounted && !preview.BroadHostMountRequired), "Local compile/run, q4 host-root mutation, privileged containers, host networking, Docker socket mounts, and broad host mounts remain closed."),
	}
}

func localGoCompilePolicyCheck(id string, status string, summary string) LocalGoCompilePolicyCheck {
	return LocalGoCompilePolicyCheck{ID: id, Status: status, Summary: summary}
}

func localGoCompilePolicyCheckIDs(checks []LocalGoCompilePolicyCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countLocalGoCompilePolicyChecks(checks []LocalGoCompilePolicyCheck) LocalGoCompilePolicyCounts {
	counts := LocalGoCompilePolicyCounts{Total: len(checks)}
	for _, check := range checks {
		switch check.Status {
		case "pass":
			counts.Passed++
		case "pending":
			counts.Pending++
		default:
			counts.Blocked++
		}
	}
	return counts
}

func allContain(source string, tokens []string) bool {
	for _, token := range tokens {
		if !strings.Contains(source, token) {
			return false
		}
	}
	return true
}

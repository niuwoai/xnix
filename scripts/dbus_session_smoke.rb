#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"

BUS_NAME = "org.xnix.Compatibility1"
OBJECT_PATH = "/org/xnix/Compatibility1"
INTERFACE = "org.xnix.Compatibility1"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

unless ENV["DBUS_SESSION_BUS_ADDRESS"]
  stdout, stderr, status = Open3.capture3("dbus-run-session", "--", "ruby", __FILE__)
  print stdout
  warn stderr unless stderr.empty?
  exit status.exitstatus
end

server_log = "/tmp/xnix-dbus-smoke.log"
server_pid = spawn("xnix-dbus-smoke", out: server_log, err: [:child, :out])

begin
  _stdout, stderr, status = Open3.capture3("gdbus", "wait", "--session", "--timeout", "5", BUS_NAME)
  assert(status.success?, "runtime smoke adapter must own #{BUS_NAME}: #{stderr}")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "introspect",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH
  )
  assert(status.success?, "runtime smoke adapter must be introspectable: #{stderr}")
  assert(stdout.include?(INTERFACE), "runtime smoke adapter introspection must expose #{INTERFACE}")
  assert(stdout.include?("ListApplications"), "runtime smoke adapter introspection must expose ListApplications")
  assert(stdout.include?("GetRuntimeServiceBinding"), "runtime smoke adapter introspection must expose Runtime service binding")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.ListApplications"
  )
  assert(status.success?, "runtime smoke adapter must answer ListApplications: #{stderr}")
  assert(stdout.include?("org.xnix.sample.notepad"), "runtime smoke adapter must expose the sample recipe over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetDiagnostics",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetDiagnostics: #{stderr}")
  assert(stdout.include?("known"), "runtime smoke adapter diagnostics must identify known applications")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetEngineCatalog"
  )
  assert(status.success?, "runtime smoke adapter must answer GetEngineCatalog: #{stderr}")
  assert(stdout.include?("compatibility-engine"), "runtime smoke adapter must expose engine catalog over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetRunPlan",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetRunPlan: #{stderr}")
  assert(stdout.include?("compatibility-run"), "runtime smoke adapter must expose run plans over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetDesktopActivationManifest",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetDesktopActivationManifest: #{stderr}")
  assert(stdout.include?("desktop-activation"), "runtime smoke adapter must expose desktop activation manifests over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetDesktopEntryPlan",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetDesktopEntryPlan: #{stderr}")
  assert(stdout.include?("desktop-entry-plan"), "runtime smoke adapter must expose desktop entry plans over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetTaskManagerIdentityPlan",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetTaskManagerIdentityPlan: #{stderr}")
  assert(stdout.include?("task-manager-identity-plan"), "runtime smoke adapter must expose task manager identity plans over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetKDEIntegrationStatus"
  )
  assert(status.success?, "runtime smoke adapter must answer GetKDEIntegrationStatus: #{stderr}")
  assert(stdout.include?("kde-integration-status"), "runtime smoke adapter must expose KDE integration status over D-Bus")
  assert(stdout.include?("GetDesktopEntryPlan"), "runtime smoke adapter KDE status must expose Runtime method coverage")

  stdout, stderr, status = Open3.capture3(
    "ruby",
    "bin/xnix-kde-integration-status",
    "--source",
    "dbus"
  )
  assert(status.success?, "KDE integration status model must consume Runtime status over D-Bus: #{stderr}")
  kde_status = JSON.parse(stdout)
  assert(kde_status.fetch("source").fetch("kind") == "runtime-dbus-session", "KDE integration status model must report the D-Bus Runtime source")
  assert(kde_status.fetch("entry_points").length == 7, "KDE integration status model must expose seven D-Bus entry points")
  assert(kde_status.fetch("entry_points").all? { |entry| entry.fetch("dbus_read_available") }, "KDE integration status model must preserve D-Bus read availability")
  assert(!kde_status.fetch("backend_details_exposed"), "KDE integration status model must preserve backend detail gates from D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetKWinWindowRulePlan",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetKWinWindowRulePlan: #{stderr}")
  assert(stdout.include?("kwin-window-rule"), "runtime smoke adapter must expose KWin window rule plans over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "ruby",
    "bin/xnix-kwin-window-rule",
    "--source",
    "dbus",
    "--app",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "KWin window rule model must consume Runtime plans over D-Bus: #{stderr}")
  kwin_rule = JSON.parse(stdout)
  assert(kwin_rule.fetch("source").fetch("kind") == "runtime-dbus-session", "KWin window rule model must report the D-Bus Runtime source")
  assert(kwin_rule.fetch("request_type") == "kwin-window-rule", "KWin window rule model must preserve Runtime plan type")
  assert(kwin_rule.fetch("match").fetch("resource_name") == "org.xnix.sample.notepad", "KWin window rule model must expose D-Bus match identity")
  assert(!kwin_rule.fetch("safety").fetch("backend_details_exposed"), "KWin window rule model must preserve backend detail gates from D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetFileAssociationPlan",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetFileAssociationPlan: #{stderr}")
  assert(stdout.include?("file-association-plan"), "runtime smoke adapter must expose file association plans over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetNotificationPlan",
    "org.xnix.sample.notepad",
    "approval-required"
  )
  assert(status.success?, "runtime smoke adapter must answer GetNotificationPlan: #{stderr}")
  assert(stdout.include?("notification-plan"), "runtime smoke adapter must expose notification plans over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetTrayStatus"
  )
  assert(status.success?, "runtime smoke adapter must answer GetTrayStatus: #{stderr}")
  assert(stdout.include?("tray-status-plan"), "runtime smoke adapter must expose tray status over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetKRunnerQueryPlan",
    "notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetKRunnerQueryPlan: #{stderr}")
  assert(stdout.include?("krunner-query-plan"), "runtime smoke adapter must expose KRunner query plans over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "ruby",
    "bin/xnix-krunner-model",
    "--source",
    "dbus",
    "--query",
    "notepad"
  )
  assert(status.success?, "KRunner model must consume Runtime query plans over D-Bus: #{stderr}")
  krunner_model = JSON.parse(stdout)
  assert(krunner_model.fetch("source").fetch("kind") == "runtime-dbus-session", "KRunner model must report the D-Bus Runtime source")
  assert(krunner_model.fetch("query_type") == "krunner-query-plan", "KRunner model must preserve Runtime query plan type")
  assert(krunner_model.fetch("matches").first.fetch("application_id") == "org.xnix.sample.notepad", "KRunner model must expose D-Bus query plan matches")
  assert(!krunner_model.fetch("summary").fetch("backend_launch_enabled"), "KRunner model must preserve backend launch gates from D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetPortalRequestPlan",
    "org.xnix.sample.notepad",
    "file-open"
  )
  assert(status.success?, "runtime smoke adapter must answer GetPortalRequestPlan: #{stderr}")
  assert(stdout.include?("portal-request-plan"), "runtime smoke adapter must expose Portal request plans over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetApplicationStateRoot",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetApplicationStateRoot: #{stderr}")
  assert(stdout.include?("compatibility-application-state-root"), "runtime smoke adapter must expose application state roots over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetCompatibilityPackageSource",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetCompatibilityPackageSource: #{stderr}")
  assert(stdout.include?("compatibility-package-source"), "runtime smoke adapter must expose compatibility package sources over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetCompatibilityAcquisitionPreflight",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetCompatibilityAcquisitionPreflight: #{stderr}")
  assert(stdout.include?("compatibility-acquisition-preflight"), "runtime smoke adapter must expose compatibility acquisition preflight over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetCompatibilityArtifactManifest",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetCompatibilityArtifactManifest: #{stderr}")
  assert(stdout.include?("compatibility-artifact-manifest"), "runtime smoke adapter must expose compatibility artifact manifests over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetCompatibilityInstallPlan",
    "org.xnix.sample.notepad",
    "development"
  )
  assert(status.success?, "runtime smoke adapter must answer GetCompatibilityInstallPlan: #{stderr}")
  assert(stdout.include?("compatibility-install-plan"), "runtime smoke adapter must expose compatibility install plans over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetBackendBinding",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetBackendBinding: #{stderr}")
  assert(stdout.include?("compatibility-backend-binding"), "runtime smoke adapter must expose backend binding over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetBackendLifecycle",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetBackendLifecycle: #{stderr}")
  assert(stdout.include?("compatibility-backend-lifecycle"), "runtime smoke adapter must expose backend lifecycle over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetBackendEnvironmentPlan",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetBackendEnvironmentPlan: #{stderr}")
  assert(stdout.include?("compatibility-backend-environment-plan"), "runtime smoke adapter must expose backend environment plans over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetRepairPlan",
    "org.xnix.sample.notepad",
    "engine-binding-pending"
  )
  assert(status.success?, "runtime smoke adapter must answer GetRepairPlan: #{stderr}")
  assert(stdout.include?("compatibility-repair"), "runtime smoke adapter must expose repair plans over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetTestPlan",
    "org.xnix.sample.notepad",
    "preflight"
  )
  assert(status.success?, "runtime smoke adapter must answer GetTestPlan: #{stderr}")
  assert(stdout.include?("compatibility-test"), "runtime smoke adapter must expose test plans over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetTestResult",
    "org.xnix.sample.notepad",
    "preflight"
  )
  assert(status.success?, "runtime smoke adapter must answer GetTestResult: #{stderr}")
  assert(stdout.include?("compatibility-test-result"), "runtime smoke adapter must expose test results over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetExecutionReadiness",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetExecutionReadiness: #{stderr}")
  assert(stdout.include?("compatibility-execution-readiness"), "runtime smoke adapter must expose execution readiness over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetLaunchIntent",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetLaunchIntent: #{stderr}")
  assert(stdout.include?("runtime-launch-intent"), "runtime smoke adapter must expose launch intent over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetAIDiagnosticInput",
    "org.xnix.sample.notepad",
    "engine-binding-pending",
    "preflight"
  )
  assert(status.success?, "runtime smoke adapter must answer GetAIDiagnosticInput: #{stderr}")
  assert(stdout.include?("ai-diagnostic-input"), "runtime smoke adapter must expose AI diagnostic inputs over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetAIDiagnosticRecommendation",
    "org.xnix.sample.notepad",
    "engine-binding-pending",
    "preflight"
  )
  assert(status.success?, "runtime smoke adapter must answer GetAIDiagnosticRecommendation: #{stderr}")
  assert(stdout.include?("ai-diagnostic-recommendation"), "runtime smoke adapter must expose AI diagnostic recommendations over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetAIRepairApprovalGate",
    "org.xnix.sample.notepad",
    "engine-binding-pending",
    "preflight"
  )
  assert(status.success?, "runtime smoke adapter must answer GetAIRepairApprovalGate: #{stderr}")
  assert(stdout.include?("ai-repair-approval-gate"), "runtime smoke adapter must expose AI repair approval gates over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetSnapshotPlan",
    "org.xnix.sample.notepad",
    "before-repair"
  )
  assert(status.success?, "runtime smoke adapter must answer GetSnapshotPlan: #{stderr}")
  assert(stdout.include?("compatibility-snapshot"), "runtime smoke adapter must expose snapshot plans over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetPortalAccessPolicy",
    "org.xnix.sample.notepad",
    "file-open"
  )
  assert(status.success?, "runtime smoke adapter must answer GetPortalAccessPolicy: #{stderr}")
  assert(stdout.include?("portal-access"), "runtime smoke adapter must expose Portal policy over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetRuntimeServiceBinding"
  )
  assert(status.success?, "runtime smoke adapter must answer GetRuntimeServiceBinding: #{stderr}")
  assert(stdout.include?("runtime-service-binding"), "runtime smoke adapter must expose Runtime service binding over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetRuntimeLiveOwnerGate"
  )
  assert(status.success?, "runtime smoke adapter must answer GetRuntimeLiveOwnerGate: #{stderr}")
  assert(stdout.include?("runtime-live-owner-gate"), "runtime smoke adapter must expose Runtime live owner gate over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetRuntimeOwnerSmokePlan"
  )
  assert(status.success?, "runtime smoke adapter must answer GetRuntimeOwnerSmokePlan: #{stderr}")
  assert(stdout.include?("runtime-owner-smoke-plan"), "runtime smoke adapter must expose Runtime owner smoke plans over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetRuntimeMethodParityManifest"
  )
  assert(status.success?, "runtime smoke adapter must answer GetRuntimeMethodParityManifest: #{stderr}")
  assert(stdout.include?("runtime-method-parity-manifest"), "runtime smoke adapter must expose Runtime method parity manifests over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetRuntimeWriteGate",
    "Launch"
  )
  assert(status.success?, "runtime smoke adapter must answer GetRuntimeWriteGate: #{stderr}")
  assert(stdout.include?("runtime-write-gate"), "runtime smoke adapter must expose Runtime write gates over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetCompatibilitySettings",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetCompatibilitySettings: #{stderr}")
  assert(stdout.include?("settings-model"), "runtime smoke adapter must expose compatibility settings over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetCompatibilitySettingsChangePlan",
    "org.xnix.sample.notepad",
    "resource-access",
    "documents",
    "ask"
  )
  assert(status.success?, "runtime smoke adapter must answer GetCompatibilitySettingsChangePlan: #{stderr}")
  assert(stdout.include?("settings-change-plan"), "runtime smoke adapter must expose settings change plans over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetCompatibilityActionQueue",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetCompatibilityActionQueue: #{stderr}")
  assert(stdout.include?("compatibility-center-action-queue"), "runtime smoke adapter must expose Compatibility Center action queues over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetCompatibilityActionReviewReceipt",
    "org.xnix.sample.notepad",
    "review-ai-repair",
    "approved"
  )
  assert(status.success?, "runtime smoke adapter must answer GetCompatibilityActionReviewReceipt: #{stderr}")
  assert(stdout.include?("compatibility-center-action-review-receipt"), "runtime smoke adapter must expose action review receipts over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetCompatibilityCenterSummary",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetCompatibilityCenterSummary: #{stderr}")
  assert(stdout.include?("compatibility-center-summary"), "runtime smoke adapter must expose Compatibility Center summaries over D-Bus")

  puts "PASS: compatibility runtime D-Bus session smoke"
ensure
  begin
    Process.kill("TERM", server_pid)
    Process.wait(server_pid)
  rescue Errno::ESRCH, Errno::ECHILD
    nil
  end
end

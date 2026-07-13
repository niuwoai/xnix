# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"
require_relative "runtime_service_binding"

module Xnix
  module Compatibility
    class RuntimeMethodParityManifest
      PROJECT_ROOT = Pathname.new(__dir__).join("../../..").realpath
      VERSION = PROJECT_ROOT.join("VERSION").read.strip
      CONTRACT_FILE = PROJECT_ROOT.join("runtime/dbus/org.xnix.Compatibility1.xml")
      RUNTIME_DAEMON_FILE = PROJECT_ROOT.join("lib/xnix/compatibility/runtime_daemon.rb")
      DBUS_CLIENT_FILE = PROJECT_ROOT.join("lib/xnix/compatibility/dbus_runtime_client.rb")
      SMOKE_ADAPTER_FILE = PROJECT_ROOT.join("runtime/dbus/xnix_compatd_smoke.c")
      SESSION_SMOKE_FILE = PROJECT_ROOT.join("scripts/dbus_session_smoke.rb")

      READ_ONLY_METHODS = %w[
        ListApplications
        GetApplication
        GetDiagnostics
        GetEngineCatalog
        GetRunPlan
        GetDesktopActivationManifest
        GetDesktopEntryPlan
        GetTaskManagerIdentityPlan
        GetKDEIntegrationStatus
        GetKWinWindowRulePlan
        GetFileAssociationPlan
        GetNotificationPlan
        GetTrayStatus
        GetKRunnerQueryPlan
        GetPortalRequestPlan
        GetApplicationStateRoot
        GetCompatibilityPackageSource
        GetCompatibilityAcquisitionPreflight
        GetCompatibilityArtifactManifest
        GetCompatibilityInstallPlan
        GetBackendBinding
        GetRepairPlan
        GetTestPlan
        GetTestResult
        GetAIDiagnosticInput
        GetAIDiagnosticRecommendation
        GetAIRepairApprovalGate
        GetSnapshotPlan
        GetPortalAccessPolicy
        GetRuntimeServiceBinding
        GetRuntimeLiveOwnerGate
        GetRuntimeOwnerSmokePlan
        GetRuntimeMethodParityManifest
        GetRuntimeWriteGate
        GetCompatibilitySettings
        GetCompatibilitySettingsChangePlan
        GetCompatibilityActionQueue
        GetCompatibilityActionReviewReceipt
        GetCompatibilityCenterSummary
      ].freeze

      WRITE_METHODS = %w[
        InstallRecipe
        Launch
        CreateSnapshot
        RestoreSnapshot
      ].freeze

      CLIENT_METHODS = {
        "ListApplications" => "list_applications",
        "GetApplication" => "get_application",
        "GetDiagnostics" => "diagnostics",
        "GetEngineCatalog" => "engine_catalog",
        "GetRunPlan" => "run_plan",
        "GetDesktopActivationManifest" => "desktop_activation_manifest",
        "GetDesktopEntryPlan" => "desktop_entry_plan",
        "GetTaskManagerIdentityPlan" => "task_manager_identity_plan",
        "GetKDEIntegrationStatus" => "kde_integration_status",
        "GetKWinWindowRulePlan" => "kwin_window_rule_plan",
        "GetFileAssociationPlan" => "file_association_plan",
        "GetNotificationPlan" => "notification_plan",
        "GetTrayStatus" => "tray_status",
        "GetKRunnerQueryPlan" => "krunner_query_plan",
        "GetPortalRequestPlan" => "portal_request_plan",
        "GetApplicationStateRoot" => "state_root",
        "GetCompatibilityPackageSource" => "package_source",
        "GetCompatibilityAcquisitionPreflight" => "acquisition_preflight",
        "GetCompatibilityArtifactManifest" => "artifact_manifest",
        "GetCompatibilityInstallPlan" => "install_plan",
        "GetBackendBinding" => "backend_binding",
        "GetRepairPlan" => "repair_plan",
        "GetTestPlan" => "test_plan",
        "GetTestResult" => "test_result",
        "GetAIDiagnosticInput" => "ai_diagnostic_input",
        "GetAIDiagnosticRecommendation" => "ai_diagnostic_recommendation",
        "GetAIRepairApprovalGate" => "ai_repair_approval_gate",
        "GetSnapshotPlan" => "snapshot_plan",
        "GetPortalAccessPolicy" => "portal_access_policy",
        "GetRuntimeServiceBinding" => "runtime_service_binding",
        "GetRuntimeLiveOwnerGate" => "runtime_live_owner_gate",
        "GetRuntimeOwnerSmokePlan" => "runtime_owner_smoke_plan",
        "GetRuntimeMethodParityManifest" => "runtime_method_parity_manifest",
        "GetRuntimeWriteGate" => "runtime_write_gate",
        "GetCompatibilitySettings" => "settings",
        "GetCompatibilitySettingsChangePlan" => "settings_change_plan",
        "GetCompatibilityActionQueue" => "action_queue",
        "GetCompatibilityActionReviewReceipt" => "action_review_receipt",
        "GetCompatibilityCenterSummary" => "compatibility_center_summary"
      }.freeze

      def to_h
        checks = parity_checks

        {
          "version" => VERSION,
          "manifest_type" => "runtime-method-parity-manifest",
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "bus_name" => RuntimeServiceBinding::BUS_NAME,
          "object_path" => RuntimeServiceBinding::OBJECT_PATH,
          "interface" => RuntimeServiceBinding::INTERFACE,
          "read_only_methods" => READ_ONLY_METHODS,
          "method_count" => READ_ONLY_METHODS.length,
          "parity_checks" => checks,
          "counts" => counts(checks),
          "read_only_method_parity_ready" => read_only_method_parity_ready?(checks),
          "write_methods" => WRITE_METHODS,
          "write_methods_supported" => false,
          "write_method_dispatch_enabled" => false,
          "network_required" => false,
          "host_root_modified" => false,
          "privileged_container_required" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => desktop_safe_summary(checks)
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def parity_checks
        [
          source_check("dbus-contract", CONTRACT_FILE) { |source, method| source.include?("name=\"#{method}\"") },
          source_check("runtime-dispatch", RUNTIME_DAEMON_FILE) { |source, method| source.include?("\"#{method}\"") },
          source_check("dbus-client", DBUS_CLIENT_FILE) { |source, method| source.include?("def #{CLIENT_METHODS.fetch(method)}") },
          source_check("smoke-adapter", SMOKE_ADAPTER_FILE) { |source, method| source.include?(method) },
          source_check("session-smoke", SESSION_SMOKE_FILE) { |source, method| source.include?(method) }
        ]
      end

      def source_check(id, path)
        source = path.file? ? path.read : ""
        missing = READ_ONLY_METHODS.reject { |method| yield(source, method) }

        {
          "id" => id,
          "status" => missing.empty? ? "pass" : "blocked",
          "method_count" => READ_ONLY_METHODS.length - missing.length,
          "missing_methods" => missing,
          "summary" => missing.empty? ? "#{id} covers all Runtime read-only methods." : "#{id} is missing Runtime read-only methods."
        }
      end

      def counts(checks)
        statuses = checks.map { |item| item.fetch("status") }
        {
          "total" => statuses.length,
          "passed" => statuses.count("pass"),
          "blocked" => statuses.count("blocked"),
          "pending" => statuses.count("pending")
        }
      end

      def read_only_method_parity_ready?(checks)
        checks.all? { |item| item.fetch("status") == "pass" }
      end

      def desktop_safe_summary(checks)
        return "Runtime read-only D-Bus method parity is ready for owner smoke." if read_only_method_parity_ready?(checks)

        "Runtime read-only D-Bus method parity must be repaired before production owner smoke."
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
        end

        def run
          parser.parse!(@argv)
          puts RuntimeMethodParityManifest.new.to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-runtime-method-parity-manifest: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-runtime-method-parity-manifest"
          end
        end
      end
    end
  end
end

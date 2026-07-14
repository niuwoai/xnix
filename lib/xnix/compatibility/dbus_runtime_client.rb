# frozen_string_literal: true

require "open3"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class DBusRuntimeClient
      class Error < StandardError; end

      attr_reader :capture

      def self.available?
        return false unless ENV["DBUS_SESSION_BUS_ADDRESS"]

        _stdout, _stderr, status = Open3.capture3(
          "gdbus", "wait",
          "--session",
          "--timeout", "1",
          RuntimeDaemon::BUS_NAME
        )
        status.success?
      rescue SystemCallError
        false
      end

      def initialize(capture: Open3.method(:capture3))
        @capture = capture
      end

      def source_metadata
        {
          "kind" => "runtime-dbus-session",
          "bus_name" => RuntimeDaemon::BUS_NAME,
          "object_path" => RuntimeDaemon::OBJECT_PATH,
          "interface" => RuntimeDaemon::INTERFACE
        }
      end

      def list_applications
        parse_dictionaries(call("ListApplications"))
      end

      def get_application(application_id)
        parse_dictionary(call("GetApplication", application_id))
      end

      def diagnostics(application_id)
        parse_dictionary(call("GetDiagnostics", application_id))
      end

      def engine_catalog
        parse_dictionary(call("GetEngineCatalog"))
      end

      def run_plan(application_id)
        parse_dictionary(call("GetRunPlan", application_id))
      end

      def desktop_activation_manifest(application_id)
        parse_dictionary(call("GetDesktopActivationManifest", application_id))
      end

      def desktop_entry_plan(application_id)
        parse_dictionary(call("GetDesktopEntryPlan", application_id))
      end

      def task_manager_identity_plan(application_id)
        parse_dictionary(call("GetTaskManagerIdentityPlan", application_id))
      end

      def kde_integration_status
        parse_dictionary(call("GetKDEIntegrationStatus"))
      end

      def kwin_window_rule_plan(application_id)
        parse_dictionary(call("GetKWinWindowRulePlan", application_id))
      end

      def file_association_plan(application_id)
        parse_dictionary(call("GetFileAssociationPlan", application_id))
      end

      def notification_plan(application_id, event_type)
        parse_dictionary(call("GetNotificationPlan", application_id, event_type))
      end

      def tray_status
        parse_dictionary(call("GetTrayStatus"))
      end

      def krunner_query_plan(query)
        parse_dictionary(call("GetKRunnerQueryPlan", query))
      end

      def state_root(application_id)
        parse_dictionary(call("GetApplicationStateRoot", application_id))
      end

      def package_source(application_id)
        parse_dictionary(call("GetCompatibilityPackageSource", application_id))
      end

      def acquisition_preflight(application_id)
        parse_dictionary(call("GetCompatibilityAcquisitionPreflight", application_id))
      end

      def action_queue(application_id)
        parse_dictionary(call("GetCompatibilityActionQueue", application_id))
      end

      def action_review_receipt(application_id, action_id, decision)
        parse_dictionary(call("GetCompatibilityActionReviewReceipt", application_id, action_id, decision))
      end

      def compatibility_center_summary(application_id)
        parse_dictionary(call("GetCompatibilityCenterSummary", application_id))
      end

      def artifact_manifest(application_id)
        parse_dictionary(call("GetCompatibilityArtifactManifest", application_id))
      end

      def install_plan(application_id, environment = "development")
        parse_dictionary(call("GetCompatibilityInstallPlan", application_id, environment))
      end

      def backend_binding(application_id)
        parse_dictionary(call("GetBackendBinding", application_id))
      end

      def repair_plan(application_id, issue)
        parse_dictionary(call("GetRepairPlan", application_id, issue))
      end

      def test_plan(application_id, test_type = "preflight")
        parse_dictionary(call("GetTestPlan", application_id, test_type))
      end

      def test_result(application_id, test_type = "preflight")
        parse_dictionary(call("GetTestResult", application_id, test_type))
      end

      def execution_readiness(application_id)
        parse_dictionary(call("GetExecutionReadiness", application_id))
      end

      def launch_intent(application_id)
        parse_dictionary(call("GetLaunchIntent", application_id))
      end

      def ai_diagnostic_input(application_id, issue = "engine-binding-pending", test_type = "preflight")
        parse_dictionary(call("GetAIDiagnosticInput", application_id, issue, test_type))
      end

      def ai_diagnostic_recommendation(application_id, issue = "engine-binding-pending", test_type = "preflight")
        parse_dictionary(call("GetAIDiagnosticRecommendation", application_id, issue, test_type))
      end

      def ai_repair_approval_gate(application_id, issue = "engine-binding-pending", test_type = "preflight")
        parse_dictionary(call("GetAIRepairApprovalGate", application_id, issue, test_type))
      end

      def snapshot_plan(application_id, reason)
        parse_dictionary(call("GetSnapshotPlan", application_id, reason))
      end

      def portal_access_policy(application_id, operation)
        parse_dictionary(call("GetPortalAccessPolicy", application_id, operation))
      end

      def portal_request_plan(application_id, operation)
        parse_dictionary(call("GetPortalRequestPlan", application_id, operation))
      end

      def runtime_service_binding
        parse_dictionary(call("GetRuntimeServiceBinding"))
      end

      def runtime_live_owner_gate
        parse_dictionary(call("GetRuntimeLiveOwnerGate"))
      end

      def runtime_owner_smoke_plan
        parse_dictionary(call("GetRuntimeOwnerSmokePlan"))
      end

      def runtime_method_parity_manifest
        parse_dictionary(call("GetRuntimeMethodParityManifest"))
      end

      def runtime_write_gate(method_name)
        parse_dictionary(call("GetRuntimeWriteGate", method_name))
      end

      def settings(application_id)
        parse_dictionary(call("GetCompatibilitySettings", application_id))
      end

      def settings_change_plan(application_id, section_id, field_id, value)
        parse_dictionary(call("GetCompatibilitySettingsChangePlan", application_id, section_id, field_id, value))
      end

      private

      def call(method_name, *arguments)
        command = [
          "gdbus", "call",
          "--session",
          "--dest", RuntimeDaemon::BUS_NAME,
          "--object-path", RuntimeDaemon::OBJECT_PATH,
          "--method", "#{RuntimeDaemon::INTERFACE}.#{method_name}",
          *arguments
        ]
        stdout, stderr, status = capture.call(*command)
        raise Error, "D-Bus call #{method_name} failed: #{stderr}" unless status.success?

        stdout
      rescue SystemCallError => e
        raise Error, "D-Bus tooling is unavailable: #{e.message}"
      end

      def parse_dictionary(output)
        dictionaries = parse_dictionaries(output)
        raise Error, "D-Bus response did not contain a dictionary" if dictionaries.empty?

        dictionaries.first
      end

      def parse_dictionaries(output)
        output.scan(/\{([^{}]*)\}/).map do |match|
          parse_dictionary_body(match.first)
        end
      end

      def parse_dictionary_body(body)
        body.scan(/'([^']+)':\s*<([^>]*)>/).each_with_object({}) do |(key, raw_value), parsed|
          parsed[key] = parse_value(raw_value)
        end
      end

      def parse_value(raw_value)
        value = raw_value.strip
        string_match = value.match(/\A'(.*)'\z/m)
        return string_match[1] if string_match
        return true if value == "true"
        return false if value == "false"
        return value.to_i if value.match?(/\A-?\d+\z/)
        return value.scan(/'([^']*)'/).flatten if value.start_with?("[") && value.end_with?("]")

        value
      end
    end
  end
end

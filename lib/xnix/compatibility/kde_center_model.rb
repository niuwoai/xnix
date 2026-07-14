# frozen_string_literal: true

require "json"
require "optparse"
require_relative "dbus_runtime_client"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class KdeCenterModel
      MODE_LABELS = {
        "automatic" => "Automatic",
        "wine" => "Compatibility engine",
        "vm" => "Isolated environment"
      }.freeze

      STATUS_LABELS = {
        "known" => "Known",
        "unknown" => "Unknown"
      }.freeze

      attr_reader :runtime

      def initialize(runtime:)
        @runtime = runtime
      end

      def to_h
        applications = runtime.list_applications.map { |application| application_summary(application) }

        {
          "version" => RuntimeDaemon::VERSION,
          "title" => "Xnix Compatibility Center",
          "source" => source_metadata,
          "summary" => summary(applications),
          "applications" => applications
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def source_metadata
        return runtime.source_metadata if runtime.respond_to?(:source_metadata)

        {
          "kind" => "runtime-local-read-model",
          "bus_name" => RuntimeDaemon::BUS_NAME,
          "object_path" => RuntimeDaemon::OBJECT_PATH,
          "interface" => RuntimeDaemon::INTERFACE
        }
      end

      def application_summary(application)
        diagnostics = runtime.diagnostics(application.fetch("id"))
        pending_checks = diagnostics.fetch("checks", []).count { |check| check["status"] == "pending" }

        {
          "id" => application.fetch("id"),
          "name" => application.fetch("name"),
          "icon" => application.fetch("icon"),
          "mode_label" => MODE_LABELS.fetch(application.fetch("mode"), "Automatic"),
          "compatibility_status" => diagnostics.fetch("status"),
          "compatibility_label" => STATUS_LABELS.fetch(diagnostics.fetch("status"), "Needs review"),
          "pending_action_count" => pending_checks,
          "action_queue" => action_queue_summary(diagnostics.fetch("action_queue", nil)),
          "action_review_receipt" => action_review_receipt_summary(diagnostics.fetch("action_review_receipt", nil)),
          "compatibility_center_summary" => compatibility_center_summary_summary(diagnostics.fetch("compatibility_center_summary", nil)),
          "repair" => repair_summary(diagnostics.fetch("repair_plan", nil)),
          "test_plan" => test_plan_summary(diagnostics.fetch("test_plan", nil)),
          "test_result" => test_result_summary(diagnostics.fetch("test_result", nil)),
          "install_plan" => install_plan_summary(diagnostics.fetch("install_plan", nil)),
          "acquisition_preflight" => acquisition_preflight_summary(diagnostics.fetch("acquisition_preflight", nil)),
          "artifact_manifest" => artifact_manifest_summary(diagnostics.fetch("artifact_manifest", nil)),
          "package_source" => package_source_summary(diagnostics.fetch("package_source", nil)),
          "state_root" => state_root_summary(diagnostics.fetch("state_root", nil)),
          "backend_binding" => backend_binding_summary(diagnostics.fetch("backend_binding", nil)),
          "backend_lifecycle" => backend_lifecycle_summary(diagnostics.fetch("backend_lifecycle", nil)),
          "ai_diagnostic_input" => ai_diagnostic_input_summary(diagnostics.fetch("ai_diagnostic_input", nil)),
          "ai_diagnostic_recommendation" => ai_diagnostic_recommendation_summary(diagnostics.fetch("ai_diagnostic_recommendation", nil)),
          "ai_repair_approval_gate" => ai_repair_approval_gate_summary(diagnostics.fetch("ai_repair_approval_gate", nil)),
          "runtime_live_owner_gate" => runtime_live_owner_gate_summary(diagnostics.fetch("runtime_live_owner_gate", nil)),
          "runtime_method_parity_manifest" => runtime_method_parity_manifest_summary(diagnostics.fetch("runtime_method_parity_manifest", nil)),
          "runtime_owner_smoke_plan" => runtime_owner_smoke_plan_summary(diagnostics.fetch("runtime_owner_smoke_plan", nil)),
          "runtime_service_binding" => runtime_service_binding_summary(diagnostics.fetch("runtime_service_binding", nil)),
          "runtime_write_gate" => runtime_write_gate_summary(diagnostics.fetch("runtime_write_gate", nil)),
          "settings" => settings_summary(diagnostics.fetch("settings", nil)),
          "settings_change_plan" => settings_change_plan_summary(diagnostics.fetch("settings_change_plan", nil)),
          "supported_extensions" => application.fetch("supported_extensions", []),
          "summary" => application_status_summary(pending_checks)
        }
      end

      def application_status_summary(pending_checks)
        return "Ready for desktop integration" if pending_checks.zero?

        "#{pending_checks} compatibility task pending"
      end

      def repair_summary(repair_plan)
        return nil unless repair_plan

        {
          "issue" => repair_plan.fetch("issue"),
          "severity" => repair_plan.fetch("severity"),
          "user_approval_required" => repair_plan.fetch("user_approval_required"),
          "snapshot_required" => repair_plan.fetch("snapshot_required"),
          "snapshot" => repair_plan.fetch("snapshot_plan", nil),
          "notification_event" => repair_plan.fetch("notification_event"),
          "summary" => repair_plan.fetch("summary")
        }
      end

      def test_plan_summary(test_plan)
        return nil unless test_plan

        {
          "plan_type" => test_plan.fetch("plan_type"),
          "test_type" => test_plan.fetch("test_type"),
          "step_count" => test_plan.fetch("step_count"),
          "pending_step_count" => test_plan.fetch("pending_step_count"),
          "blocked" => test_plan.fetch("blocked"),
          "summary" => test_plan.fetch("summary")
        }
      end

      def test_result_summary(test_result)
        return nil unless test_result

        {
          "result_type" => test_result.fetch("result_type"),
          "test_type" => test_result.fetch("test_type"),
          "execution_state" => test_result.fetch("execution_state"),
          "overall_status" => test_result.fetch("overall_status"),
          "counts" => test_result.fetch("counts"),
          "summary" => test_result.fetch("summary")
        }
      end

      def action_queue_summary(queue)
        return nil unless queue

        {
          "queue_type" => queue.fetch("queue_type"),
          "action_count" => queue.fetch("action_count"),
          "pending_action_count" => queue.fetch("pending_action_count"),
          "user_review_required_count" => queue.fetch("user_review_required_count"),
          "execution_enabled" => queue.fetch("execution_enabled"),
          "repair_execution_enabled" => queue.fetch("repair_execution_enabled"),
          "settings_persistence_enabled" => queue.fetch("settings_persistence_enabled"),
          "backend_details_exposed" => queue.fetch("backend_details_exposed"),
          "summary" => queue.fetch("summary")
        }
      end

      def action_review_receipt_summary(receipt)
        return nil unless receipt

        {
          "receipt_type" => receipt.fetch("receipt_type"),
          "decision" => receipt.fetch("decision"),
          "decision_recorded" => receipt.fetch("decision_recorded"),
          "action_id" => receipt.fetch("action_id"),
          "execution_enabled" => receipt.fetch("execution_enabled"),
          "repair_execution_enabled" => receipt.fetch("repair_execution_enabled"),
          "settings_persistence_enabled" => receipt.fetch("settings_persistence_enabled"),
          "resource_grant_created" => receipt.fetch("resource_grant_created"),
          "backend_details_exposed" => receipt.fetch("backend_details_exposed"),
          "summary" => receipt.fetch("summary")
        }
      end

      def compatibility_center_summary_summary(summary)
        return nil unless summary

        {
          "summary_type" => summary.fetch("summary_type"),
          "compatibility_state" => summary.fetch("compatibility_state"),
          "runtime_mode" => summary.fetch("runtime_mode"),
          "known_issue_count" => summary.fetch("known_issue_count"),
          "repair_record_state" => summary.fetch("repair_record_state"),
          "repair_record_count" => summary.fetch("repair_record_count"),
          "last_repair_event" => summary.fetch("last_repair_event"),
          "action_count" => summary.fetch("action_count"),
          "action_execution_enabled" => summary.fetch("action_execution_enabled"),
          "repair_execution_enabled" => summary.fetch("repair_execution_enabled"),
          "backend_launch_enabled" => summary.fetch("backend_launch_enabled"),
          "backend_details_exposed" => summary.fetch("backend_details_exposed"),
          "summary" => summary.fetch("summary")
        }
      end

      def backend_binding_summary(binding)
        return nil unless binding

        {
          "binding_type" => binding.fetch("binding_type"),
          "selected_strategy" => binding.fetch("selected_strategy"),
          "managed_binding_ready" => binding.fetch("managed_binding_ready"),
          "launch_enabled" => binding.fetch("launch_enabled"),
          "preflight_count" => binding.fetch("preflight_count"),
          "summary" => binding.fetch("summary")
        }
      end

      def backend_lifecycle_summary(lifecycle)
        return nil unless lifecycle

        {
          "lifecycle_type" => lifecycle.fetch("lifecycle_type"),
          "lifecycle_state" => lifecycle.fetch("lifecycle_state"),
          "overall_status" => lifecycle.fetch("overall_status"),
          "stage_count" => lifecycle.fetch("stage_count"),
          "backend_process_started" => lifecycle.fetch("backend_process_started"),
          "launch_enabled" => lifecycle.fetch("launch_enabled"),
          "summary" => lifecycle.fetch("summary")
        }
      end

      def state_root_summary(state_root)
        return nil unless state_root

        {
          "root_type" => state_root.fetch("root_type"),
          "state_namespace" => state_root.fetch("state_namespace"),
          "allocation_state" => state_root.fetch("allocation_state"),
          "managed_scope_count" => state_root.fetch("managed_scope_count"),
          "snapshot_eligible" => state_root.fetch("snapshot_eligible"),
          "user_documents_included" => state_root.fetch("user_documents_included"),
          "summary" => state_root.fetch("summary")
        }
      end

      def package_source_summary(source)
        return nil unless source

        {
          "source_type" => source.fetch("source_type"),
          "selected_strategy" => source.fetch("selected_strategy"),
          "source_selection_state" => source.fetch("source_selection_state"),
          "package_source_ready" => source.fetch("package_source_ready"),
          "install_enabled" => source.fetch("install_enabled"),
          "channel_count" => source.fetch("channel_count"),
          "preflight_count" => source.fetch("preflight_count"),
          "summary" => source.fetch("summary")
        }
      end

      def acquisition_preflight_summary(preflight)
        return nil unless preflight

        {
          "preflight_type" => preflight.fetch("preflight_type"),
          "selected_strategy" => preflight.fetch("selected_strategy"),
          "preflight_state" => preflight.fetch("preflight_state"),
          "acquisition_ready" => preflight.fetch("acquisition_ready"),
          "download_enabled" => preflight.fetch("download_enabled"),
          "install_enabled" => preflight.fetch("install_enabled"),
          "check_count" => preflight.fetch("check_count"),
          "summary" => preflight.fetch("summary")
        }
      end

      def artifact_manifest_summary(manifest)
        return nil unless manifest

        {
          "manifest_type" => manifest.fetch("manifest_type"),
          "selected_strategy" => manifest.fetch("selected_strategy"),
          "manifest_state" => manifest.fetch("manifest_state"),
          "manifest_ready" => manifest.fetch("manifest_ready"),
          "signature_verified" => manifest.fetch("signature_verified"),
          "download_enabled" => manifest.fetch("download_enabled"),
          "artifact_group_count" => manifest.fetch("artifact_group_count"),
          "summary" => manifest.fetch("summary")
        }
      end

      def install_plan_summary(plan)
        return nil unless plan

        {
          "plan_type" => plan.fetch("plan_type"),
          "selected_strategy" => plan.fetch("selected_strategy"),
          "install_state" => plan.fetch("install_state"),
          "install_ready" => plan.fetch("install_ready"),
          "desktop_activation_ready" => plan.fetch("desktop_activation_ready"),
          "download_enabled" => plan.fetch("download_enabled"),
          "install_enabled" => plan.fetch("install_enabled"),
          "phase_count" => plan.fetch("phase_count"),
          "blocked_phase_count" => plan.fetch("blocked_phase_count"),
          "summary" => plan.fetch("summary")
        }
      end

      def ai_diagnostic_input_summary(ai_input)
        return nil unless ai_input

        {
          "input_type" => ai_input.fetch("input_type"),
          "section_count" => ai_input.fetch("section_count"),
          "signal_count" => ai_input.fetch("signal_count"),
          "ai_provider_called" => ai_input.fetch("ai_provider_called"),
          "network_required" => ai_input.fetch("network_required"),
          "safe_for_ai_diagnostics" => ai_input.fetch("safe_for_ai_diagnostics"),
          "summary" => ai_input.fetch("summary")
        }
      end

      def ai_diagnostic_recommendation_summary(recommendation)
        return nil unless recommendation

        {
          "recommendation_type" => recommendation.fetch("recommendation_type"),
          "recommendation_count" => recommendation.fetch("recommendation_count"),
          "approval_required_count" => recommendation.fetch("approval_required_count"),
          "ai_provider_called" => recommendation.fetch("ai_provider_called"),
          "network_required" => recommendation.fetch("network_required"),
          "safe_for_ai_diagnostics" => recommendation.fetch("safe_for_ai_diagnostics"),
          "summary" => recommendation.fetch("summary")
        }
      end

      def ai_repair_approval_gate_summary(gate)
        return nil unless gate

        {
          "gate_type" => gate.fetch("gate_type"),
          "gate_decision" => gate.fetch("gate_decision"),
          "required_gate_count" => gate.fetch("required_gate_count"),
          "approval_required_count" => gate.fetch("approval_required_count"),
          "auto_execution_allowed" => gate.fetch("auto_execution_allowed"),
          "repair_executed" => gate.fetch("repair_executed"),
          "summary" => gate.fetch("summary")
        }
      end

      def runtime_service_binding_summary(binding)
        return nil unless binding

        counts = binding.fetch("counts", {})
        {
          "binding_type" => binding.fetch("binding_type"),
          "activation_binding_ready" => binding.fetch("activation_binding_ready"),
          "live_dbus_owner_ready" => binding.fetch("live_dbus_owner_ready"),
          "pending_count" => counts.fetch("pending", 0),
          "blocked_count" => counts.fetch("blocked", 0),
          "summary" => binding.fetch("summary", "")
        }
      end

      def runtime_live_owner_gate_summary(gate)
        return nil unless gate

        {
          "gate_type" => gate.fetch("gate_type"),
          "activation_binding_ready" => gate.fetch("activation_binding_ready"),
          "live_dbus_owner_ready" => gate.fetch("live_dbus_owner_ready"),
          "production_owner_enabled" => gate.fetch("production_owner_enabled"),
          "owner_transition_ready" => gate.fetch("owner_transition_ready"),
          "smoke_adapter_is_production_owner" => gate.fetch("smoke_adapter_is_production_owner"),
          "pending_gate_count" => gate.fetch("pending_gate_count"),
          "blocked_reason_count" => gate.fetch("blocked_reason_count"),
          "summary" => gate.fetch("summary", "")
        }
      end

      def runtime_owner_smoke_plan_summary(plan)
        return nil unless plan

        {
          "plan_type" => plan.fetch("plan_type"),
          "smoke_state" => plan.fetch("smoke_state"),
          "smoke_environment" => plan.fetch("smoke_environment"),
          "activation_binding_ready" => plan.fetch("activation_binding_ready"),
          "live_dbus_owner_ready" => plan.fetch("live_dbus_owner_ready"),
          "production_owner_enabled" => plan.fetch("production_owner_enabled"),
          "owner_transition_ready" => plan.fetch("owner_transition_ready"),
          "pending_step_count" => plan.fetch("pending_step_count"),
          "system_service_started" => plan.fetch("system_service_started"),
          "production_bus_claimed" => plan.fetch("production_bus_claimed"),
          "summary" => plan.fetch("summary", "")
        }
      end

      def runtime_method_parity_manifest_summary(manifest)
        return nil unless manifest

        {
          "manifest_type" => manifest.fetch("manifest_type"),
          "method_count" => manifest.fetch("method_count"),
          "read_only_method_parity_ready" => manifest.fetch("read_only_method_parity_ready"),
          "passed_check_count" => manifest.fetch("passed_check_count"),
          "blocked_check_count" => manifest.fetch("blocked_check_count"),
          "write_methods_supported" => manifest.fetch("write_methods_supported"),
          "write_method_dispatch_enabled" => manifest.fetch("write_method_dispatch_enabled"),
          "summary" => manifest.fetch("summary", "")
        }
      end

      def runtime_write_gate_summary(gate)
        return nil unless gate

        {
          "gate_type" => gate.fetch("gate_type"),
          "method_name" => gate.fetch("method_name"),
          "gate_decision" => gate.fetch("gate_decision"),
          "write_method_enabled" => gate.fetch("write_method_enabled"),
          "dispatch_enabled" => gate.fetch("dispatch_enabled"),
          "request_object_created" => gate.fetch("request_object_created"),
          "required_gate_count" => gate.fetch("required_gate_count"),
          "denial_error_name" => gate.fetch("denial_error_name"),
          "backend_details_exposed" => gate.fetch("backend_details_exposed"),
          "summary" => gate.fetch("summary", "")
        }
      end

      def settings_summary(settings)
        return nil unless settings

        {
          "request_type" => settings.fetch("request_type"),
          "settings_state" => settings.fetch("settings_state"),
          "settings_persisted" => settings.fetch("settings_persisted"),
          "section_count" => settings.fetch("section_count"),
          "host_root_modified" => settings.fetch("host_root_modified"),
          "backend_details_exposed" => settings.fetch("backend_details_exposed"),
          "summary" => settings.fetch("summary", "")
        }
      end

      def settings_change_plan_summary(plan)
        return nil unless plan

        {
          "plan_type" => plan.fetch("plan_type"),
          "change_state" => plan.fetch("change_state"),
          "apply_enabled" => plan.fetch("apply_enabled"),
          "settings_persisted" => plan.fetch("settings_persisted"),
          "portal_policy_review_required" => plan.fetch("portal_policy_review_required"),
          "snapshot_recommended" => plan.fetch("snapshot_recommended"),
          "backend_details_exposed" => plan.fetch("backend_details_exposed"),
          "summary" => plan.fetch("summary", "")
        }
      end

      def summary(applications)
        {
          "application_count" => applications.length,
          "known_application_count" => applications.count { |application| application["compatibility_status"] == "known" },
          "pending_action_count" => applications.sum { |application| application["pending_action_count"] },
          "queued_compatibility_action_count" => applications.sum { |application| section(application, "action_queue").fetch("action_count", 0) },
          "queued_user_review_count" => applications.sum { |application| section(application, "action_queue").fetch("user_review_required_count", 0) },
          "recorded_action_review_count" => applications.count { |application| section(application, "action_review_receipt").fetch("decision_recorded", false) },
          "compatibility_center_summary_count" => applications.count { |application| section(application, "compatibility_center_summary").fetch("summary_type", nil) == "compatibility-center-summary" },
          "compatibility_center_known_issue_count" => applications.sum { |application| section(application, "compatibility_center_summary").fetch("known_issue_count", 0) },
          "pending_test_step_count" => applications.sum { |application| section(application, "test_plan").fetch("pending_step_count", 0) },
          "pending_test_result_count" => applications.count { |application| section(application, "test_result").fetch("overall_status", nil) == "pending" },
          "pending_install_plan_count" => applications.count { |application| !section(application, "install_plan").fetch("install_ready", false) },
          "pending_acquisition_preflight_count" => applications.count { |application| !section(application, "acquisition_preflight").fetch("acquisition_ready", false) },
          "pending_artifact_manifest_count" => applications.count { |application| !section(application, "artifact_manifest").fetch("manifest_ready", false) },
          "pending_package_source_count" => applications.count { |application| !section(application, "package_source").fetch("package_source_ready", false) },
          "planned_state_root_count" => applications.count { |application| section(application, "state_root").fetch("allocation_state", nil) == "planned" },
          "pending_backend_binding_count" => applications.count { |application| !section(application, "backend_binding").fetch("managed_binding_ready", false) },
          "ai_diagnostic_ready_count" => applications.count { |application| section(application, "ai_diagnostic_input").fetch("safe_for_ai_diagnostics", false) },
          "ai_recommendation_ready_count" => applications.count { |application| section(application, "ai_diagnostic_recommendation").fetch("safe_for_ai_diagnostics", false) },
          "blocked_ai_repair_gate_count" => applications.count { |application| section(application, "ai_repair_approval_gate").fetch("gate_decision", nil) == "blocked-until-approval" },
          "pending_runtime_live_owner_gate_count" => applications.count do |application|
            live_owner_gate = section(application, "runtime_live_owner_gate")
            !live_owner_gate.empty? && !live_owner_gate.fetch("owner_transition_ready", false)
          end,
          "pending_runtime_owner_smoke_plan_count" => applications.count do |application|
            owner_smoke_plan = section(application, "runtime_owner_smoke_plan")
            !owner_smoke_plan.empty? && owner_smoke_plan.fetch("smoke_state", nil) == "planned"
          end,
          "runtime_method_parity_ready_count" => applications.count do |application|
            section(application, "runtime_method_parity_manifest").fetch("read_only_method_parity_ready", false)
          end,
          "runtime_service_binding_ready_count" => applications.count { |application| section(application, "runtime_service_binding").fetch("activation_binding_ready", false) },
          "blocked_runtime_write_gate_count" => applications.count { |application| section(application, "runtime_write_gate").fetch("gate_decision", nil) == "blocked-until-production-backend" },
          "runtime_settings_model_count" => applications.count { |application| section(application, "settings").fetch("settings_state", nil) == "planned" },
          "pending_settings_change_plan_count" => applications.count { |application| !section(application, "settings_change_plan").fetch("apply_enabled", false) }
        }
      end

      def section(application, key)
        value = application.fetch(key, nil)
        return value if value.is_a?(Hash)

        {}
      end

      class CLI
        SOURCES = %w[auto local dbus].freeze

        def initialize(argv)
          @argv = argv.dup
          @recipe_dir = RuntimeDaemon::DEFAULT_RECIPE_DIR
          @source = "auto"
        end

        def run
          parser.parse!(@argv)
          runtime = runtime_source
          puts KdeCenterModel.new(runtime: runtime).to_json
          0
        rescue OptionParser::ParseError, KeyError, ArgumentError, DBusRuntimeClient::Error => e
          warn "xnix-kde-center-model: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-kde-center-model [--recipe-dir PATH]"
            options.on("--recipe-dir PATH", "Read application recipes from PATH") do |value|
              @recipe_dir = value
            end
            options.on("--source SOURCE", "Read from auto, local, or dbus") do |value|
              raise OptionParser::InvalidArgument, "source must be one of: #{SOURCES.join(", ")}" unless SOURCES.include?(value)

              @source = value
            end
          end
        end

        def runtime_source
          case @source
          when "local"
            local_runtime
          when "dbus"
            DBusRuntimeClient.new
          when "auto"
            DBusRuntimeClient.available? ? DBusRuntimeClient.new : local_runtime
          end
        end

        def local_runtime
          RuntimeDaemon.new(recipe_store: RecipeStore.new(path: @recipe_dir))
        end
      end
    end
  end
end

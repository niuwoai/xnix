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
          "repair" => repair_summary(diagnostics.fetch("repair_plan", nil)),
          "test_plan" => test_plan_summary(diagnostics.fetch("test_plan", nil)),
          "test_result" => test_result_summary(diagnostics.fetch("test_result", nil)),
          "ai_diagnostic_input" => ai_diagnostic_input_summary(diagnostics.fetch("ai_diagnostic_input", nil)),
          "ai_diagnostic_recommendation" => ai_diagnostic_recommendation_summary(diagnostics.fetch("ai_diagnostic_recommendation", nil)),
          "ai_repair_approval_gate" => ai_repair_approval_gate_summary(diagnostics.fetch("ai_repair_approval_gate", nil)),
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

      def summary(applications)
        {
          "application_count" => applications.length,
          "known_application_count" => applications.count { |application| application["compatibility_status"] == "known" },
          "pending_action_count" => applications.sum { |application| application["pending_action_count"] },
          "pending_test_step_count" => applications.sum { |application| application.fetch("test_plan", {}).fetch("pending_step_count", 0) },
          "pending_test_result_count" => applications.count { |application| application.fetch("test_result", {}).fetch("overall_status", nil) == "pending" },
          "ai_diagnostic_ready_count" => applications.count { |application| application.fetch("ai_diagnostic_input", {}).fetch("safe_for_ai_diagnostics", false) },
          "ai_recommendation_ready_count" => applications.count { |application| application.fetch("ai_diagnostic_recommendation", {}).fetch("safe_for_ai_diagnostics", false) },
          "blocked_ai_repair_gate_count" => applications.count { |application| application.fetch("ai_repair_approval_gate", {}).fetch("gate_decision", nil) == "blocked-until-approval" }
        }
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

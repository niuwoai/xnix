# frozen_string_literal: true

require "json"
require "optparse"
require_relative "compatibility_repair_plan"
require_relative "compatibility_run_plan"
require_relative "compatibility_test_result"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class AIDiagnosticInput
      DEFAULT_ISSUE = "engine-binding-pending"

      def initialize(recipe:, issue: DEFAULT_ISSUE, test_type: "preflight")
        @recipe = recipe
        @issue = issue
        @test_type = test_type
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "input_type" => "ai-diagnostic-input",
          "application" => application_context,
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "ai_provider_called" => false,
          "network_required" => false,
          "safe_for_ai_diagnostics" => true,
          "context_sections" => context_sections,
          "diagnostic_signals" => diagnostic_signals,
          "privacy_boundaries" => privacy_boundaries,
          "allowed_ai_tasks" => allowed_ai_tasks,
          "blocked_ai_tasks" => blocked_ai_tasks,
          "desktop_safe_summary" => desktop_safe_summary,
          "backend_details_exposed" => false
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      attr_reader :recipe

      def application_context
        {
          "id" => recipe.id,
          "name" => recipe.name,
          "icon" => recipe.icon,
          "mode" => recipe.mode,
          "supported_extensions" => recipe.supported_extensions,
          "mime_types" => recipe.mime_types
        }
      end

      def context_sections
        [
          recipe_section,
          run_section,
          test_section,
          repair_section
        ]
      end

      def recipe_section
        {
          "id" => "recipe",
          "summary" => "Application recipe metadata is available.",
          "facts" => {
            "application_id" => recipe.id,
            "display_name" => recipe.name,
            "requested_mode" => recipe.mode,
            "supported_extensions" => recipe.supported_extensions
          }
        }
      end

      def run_section
        execution = run_plan.fetch("execution")
        binding = execution.fetch("backend_binding")

        {
          "id" => "run-plan",
          "summary" => run_plan.fetch("desktop_safe_summary"),
          "facts" => {
            "strategy" => execution.fetch("strategy"),
            "backend_binding_ready" => binding.fetch("ready"),
            "launch_enabled" => binding.fetch("launch_enabled"),
            "backend_details_exposed" => execution.fetch("backend_details_exposed")
          }
        }
      end

      def test_section
        {
          "id" => "test-result",
          "summary" => test_result.fetch("desktop_safe_summary"),
          "facts" => {
            "test_type" => test_result.fetch("test_type"),
            "execution_state" => test_result.fetch("execution_state"),
            "overall_status" => test_result.fetch("overall_status"),
            "counts" => test_result.fetch("counts")
          }
        }
      end

      def repair_section
        {
          "id" => "repair-plan",
          "summary" => repair_plan.fetch("desktop_safe_summary"),
          "facts" => {
            "issue" => repair_plan.fetch("issue"),
            "severity" => repair_plan.fetch("severity"),
            "user_approval_required" => repair_plan.fetch("user_approval_required"),
            "snapshot_required" => repair_plan.fetch("snapshot_required"),
            "rollback_available" => repair_plan.fetch("rollback_available")
          }
        }
      end

      def diagnostic_signals
        [
          {
            "id" => "pending-runtime-launch-binding",
            "severity" => "medium",
            "source" => "run-plan",
            "summary" => "Managed compatibility launch binding is not ready yet.",
            "recommended_next_action" => "Keep automated smoke execution pending until a launch backend is bound."
          },
          {
            "id" => "pending-test-work",
            "severity" => "low",
            "source" => "test-result",
            "summary" => "#{test_result.fetch("counts").fetch("pending")} compatibility test steps are pending.",
            "recommended_next_action" => "Show pending Runtime-controlled preflight work in the Compatibility Center."
          },
          {
            "id" => "snapshot-before-risky-change",
            "severity" => "low",
            "source" => "repair-plan",
            "summary" => "A Runtime restore point is required before risky compatibility changes.",
            "recommended_next_action" => "Prepare a restore point before repair execution."
          }
        ]
      end

      def privacy_boundaries
        {
          "user_documents_included" => false,
          "host_paths_included" => false,
          "raw_backend_logs_included" => false,
          "secrets_included" => false,
          "network_calls_allowed" => false,
          "requires_user_approval_for_sensitive_actions" => true
        }
      end

      def allowed_ai_tasks
        [
          "summarize compatibility status",
          "explain pending Runtime work",
          "suggest safe next diagnostic steps",
          "prepare user-facing Compatibility Center text"
        ]
      end

      def blocked_ai_tasks
        [
          "start a compatibility backend",
          "read user documents",
          "change desktop permissions",
          "execute repair actions without Runtime approval"
        ]
      end

      def desktop_safe_summary
        "AI diagnostics can explain current compatibility status using Runtime-safe metadata only."
      end

      def run_plan
        @run_plan ||= CompatibilityRunPlan.new(recipe: recipe).to_h
      end

      def test_result
        @test_result ||= CompatibilityTestResult.new(recipe: recipe, test_type: @test_type).to_h
      end

      def repair_plan
        @repair_plan ||= CompatibilityRepairPlan.new(application_id: recipe.id, issue: @issue).to_h
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @recipe_dir = RuntimeDaemon::DEFAULT_RECIPE_DIR
          @issue = DEFAULT_ISSUE
          @test_type = "preflight"
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--app is required" unless @application_id

          recipe = RecipeStore.new(path: @recipe_dir).find(@application_id)
          raise ArgumentError, "unknown application: #{@application_id}" unless recipe

          puts AIDiagnosticInput.new(recipe: recipe, issue: @issue, test_type: @test_type).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-ai-diagnostic-input: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-ai-diagnostic-input --app APP_ID [--issue ISSUE] [--test-type TYPE] [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build an AI diagnostic input for APP_ID") do |value|
              @application_id = value
            end
            options.on("--issue ISSUE", "Use compatibility issue ISSUE") do |value|
              @issue = value
            end
            options.on("--test-type TYPE", CompatibilityTestPlan::TEST_TYPES, "Use a preflight, smoke, or repair-readiness test result") do |value|
              @test_type = value
            end
            options.on("--recipe-dir PATH", "Read application recipes from PATH") do |value|
              @recipe_dir = value
            end
          end
        end
      end
    end
  end
end

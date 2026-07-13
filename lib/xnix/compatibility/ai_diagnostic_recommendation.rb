# frozen_string_literal: true

require "json"
require "optparse"
require_relative "ai_diagnostic_input"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class AIDiagnosticRecommendation
      def initialize(recipe:, issue: AIDiagnosticInput::DEFAULT_ISSUE, test_type: "preflight")
        @recipe = recipe
        @issue = issue
        @test_type = test_type
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "recommendation_type" => "ai-diagnostic-recommendation",
          "input_type" => diagnostic_input.fetch("input_type"),
          "application" => diagnostic_input.fetch("application"),
          "runtime_owned" => true,
          "runtime_method" => "GetAIDiagnosticRecommendation",
          "c_runtime_backed" => true,
          "kde_policy_owner" => false,
          "ai_provider_called" => false,
          "network_required" => false,
          "safe_for_ai_diagnostics" => true,
          "auto_execution_allowed" => false,
          "recommendations" => recommendations,
          "approval_required_actions" => approval_required_actions,
          "blocked_actions" => diagnostic_input.fetch("blocked_ai_tasks"),
          "privacy_boundaries" => diagnostic_input.fetch("privacy_boundaries"),
          "desktop_safe_summary" => desktop_safe_summary,
          "backend_details_exposed" => false
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      attr_reader :recipe

      def diagnostic_input
        @diagnostic_input ||= AIDiagnosticInput.new(recipe: recipe, issue: @issue, test_type: @test_type).to_h
      end

      def recommendations
        [
          {
            "id" => "explain-pending-runtime-work",
            "priority" => "high",
            "source_signal" => "pending-runtime-launch-binding",
            "title" => "Explain pending Runtime launch binding",
            "summary" => "Tell the user that automated compatibility execution is waiting for a managed Runtime launch backend.",
            "user_visible" => true,
            "auto_execute" => false,
            "requires_user_approval" => false,
            "next_runtime_action" => "Keep smoke execution pending until backend binding is available."
          },
          {
            "id" => "prepare-safe-restore-point",
            "priority" => "medium",
            "source_signal" => "snapshot-before-risky-change",
            "title" => "Prepare a restore point before repair",
            "summary" => "Ask the Runtime to create an application-scoped restore point before risky compatibility changes.",
            "user_visible" => true,
            "auto_execute" => false,
            "requires_user_approval" => true,
            "next_runtime_action" => "Create a Runtime restore point only after approval."
          },
          {
            "id" => "surface-test-progress",
            "priority" => "medium",
            "source_signal" => "pending-test-work",
            "title" => "Show compatibility test progress",
            "summary" => "Display pending preflight work in the Compatibility Center without exposing backend details.",
            "user_visible" => true,
            "auto_execute" => false,
            "requires_user_approval" => false,
            "next_runtime_action" => "Update the Compatibility Center card with pending test status."
          }
        ]
      end

      def approval_required_actions
        recommendations.select { |recommendation| recommendation.fetch("requires_user_approval") }.map do |recommendation|
          {
            "id" => recommendation.fetch("id"),
            "title" => recommendation.fetch("title"),
            "reason" => recommendation.fetch("summary"),
            "approval_surface" => "Compatibility Center"
          }
        end
      end

      def desktop_safe_summary
        "AI diagnostic recommendations are ready for review without executing compatibility changes."
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @recipe_dir = RuntimeDaemon::DEFAULT_RECIPE_DIR
          @issue = AIDiagnosticInput::DEFAULT_ISSUE
          @test_type = "preflight"
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--app is required" unless @application_id

          recipe = RecipeStore.new(path: @recipe_dir).find(@application_id)
          raise ArgumentError, "unknown application: #{@application_id}" unless recipe

          puts AIDiagnosticRecommendation.new(recipe: recipe, issue: @issue, test_type: @test_type).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-ai-diagnostic-recommendation: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-ai-diagnostic-recommendation --app APP_ID [--issue ISSUE] [--test-type TYPE] [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build AI diagnostic recommendations for APP_ID") do |value|
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

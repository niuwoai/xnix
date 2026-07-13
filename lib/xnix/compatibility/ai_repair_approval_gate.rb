# frozen_string_literal: true

require "json"
require "optparse"
require_relative "ai_diagnostic_input"
require_relative "ai_diagnostic_recommendation"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class AIRepairApprovalGate
      def initialize(recipe:, issue: AIDiagnosticInput::DEFAULT_ISSUE, test_type: "preflight")
        @recipe = recipe
        @issue = issue
        @test_type = test_type
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "gate_type" => "ai-repair-approval-gate",
          "recommendation_type" => recommendation.fetch("recommendation_type"),
          "application" => recommendation.fetch("application"),
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "ai_provider_called" => false,
          "network_required" => false,
          "safe_for_ai_diagnostics" => true,
          "repair_execution_requested" => false,
          "repair_executed" => false,
          "auto_execution_allowed" => false,
          "gate_decision" => gate_decision,
          "approval_surface" => "Compatibility Center",
          "required_gates" => required_gates,
          "approval_required_actions" => recommendation.fetch("approval_required_actions"),
          "blocked_actions" => blocked_actions,
          "desktop_safe_summary" => desktop_safe_summary,
          "backend_details_exposed" => false
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      attr_reader :recipe

      def recommendation
        @recommendation ||= AIDiagnosticRecommendation.new(recipe: recipe, issue: @issue, test_type: @test_type).to_h
      end

      def gate_decision
        return "blocked-until-approval" unless recommendation.fetch("approval_required_actions").empty?

        "review-only"
      end

      def required_gates
        [
          {
            "id" => "compatibility-center-review",
            "status" => "required",
            "summary" => "A user-visible Compatibility Center review is required before any repair action."
          },
          {
            "id" => "runtime-approval-token",
            "status" => "missing",
            "summary" => "No Runtime approval token is present for repair execution."
          },
          {
            "id" => "restore-point-preflight",
            "status" => "required",
            "summary" => "A Runtime restore point must be prepared before risky compatibility repair."
          }
        ]
      end

      def blocked_actions
        (recommendation.fetch("blocked_actions") + [
          "execute repair without approval",
          "create restore point without approval",
          "change compatibility mode without Runtime gate"
        ]).uniq
      end

      def desktop_safe_summary
        "AI repair recommendations are blocked from execution until Runtime approval gates pass."
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

          puts AIRepairApprovalGate.new(recipe: recipe, issue: @issue, test_type: @test_type).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-ai-repair-approval-gate: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-ai-repair-approval-gate --app APP_ID [--issue ISSUE] [--test-type TYPE] [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build an AI repair approval gate for APP_ID") do |value|
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

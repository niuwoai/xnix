# frozen_string_literal: true

require "json"
require "optparse"
require_relative "compatibility_test_plan"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class CompatibilityTestResult
      RESULT_SOURCE = "runtime-model"

      def initialize(recipe:, test_type: "preflight")
        @recipe = recipe
        @test_type = test_type
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "result_type" => "compatibility-test-result",
          "plan_type" => plan.fetch("plan_type"),
          "test_type" => plan.fetch("test_type"),
          "application_id" => plan.fetch("application_id"),
          "name" => plan.fetch("name"),
          "runtime_method" => "GetTestResult",
          "runtime_owned" => true,
          "c_runtime_backed" => true,
          "kde_policy_owner" => false,
          "result_source" => RESULT_SOURCE,
          "execution_state" => execution_state,
          "overall_status" => overall_status,
          "counts" => counts,
          "step_results" => step_results,
          "artifacts" => artifacts,
          "safe_for_ai_diagnostics" => true,
          "test_executed" => false,
          "host_root_modified" => false,
          "desktop_safe_summary" => desktop_safe_summary,
          "backend_details_exposed" => false
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      attr_reader :recipe

      def plan
        @plan ||= CompatibilityTestPlan.new(recipe: recipe, test_type: @test_type).to_h
      end

      def step_results
        plan.fetch("steps").map do |step|
          {
            "id" => step.fetch("id"),
            "title" => step.fetch("title"),
            "status" => step.fetch("status"),
            "result" => result_for_status(step.fetch("status")),
            "evidence" => evidence_for_step(step),
            "next_action" => next_action_for_step(step)
          }
        end
      end

      def result_for_status(status)
        case status
        when "pass"
          "passed"
        when "blocked"
          "blocked"
        else
          "not-run"
        end
      end

      def evidence_for_step(step)
        case step.fetch("id")
        when "recipe-validation"
          ["Recipe metadata is loaded and valid."]
        when "portal-preflight"
          ["XDG Desktop Portal request model is available."]
        when "snapshot-preflight"
          ["Runtime snapshot plan is available and restore-capable."]
        when "runtime-launch-binding"
          ["Runtime launch backend binding is still pending."]
        else
          [step.fetch("summary")]
        end
      end

      def next_action_for_step(step)
        case step.fetch("status")
        when "pass"
          "No action required."
        when "blocked"
          "Request user approval before continuing."
        else
          pending_action_for_step(step.fetch("id"))
        end
      end

      def pending_action_for_step(step_id)
        case step_id
        when "portal-preflight"
          "Wait for a user-mediated Portal grant during execution."
        when "snapshot-preflight"
          "Create a Runtime restore point before risky compatibility changes."
        when "runtime-launch-binding"
          "Bind a managed compatibility launch backend before automated smoke execution."
        else
          "Run the pending compatibility test step."
        end
      end

      def counts
        statuses = step_results.map { |step| step.fetch("status") }
        {
          "total" => statuses.length,
          "passed" => statuses.count("pass"),
          "pending" => statuses.count("pending"),
          "blocked" => statuses.count("blocked")
        }
      end

      def execution_state
        return "blocked" if counts.fetch("blocked").positive?
        return "waiting-for-runtime" if counts.fetch("pending").positive?

        "complete"
      end

      def overall_status
        return "blocked" if counts.fetch("blocked").positive?
        return "pending" if counts.fetch("pending").positive?

        "pass"
      end

      def artifacts
        {
          "compatibility_center_card" => true,
          "diagnostics_record" => true,
          "notification_event" => overall_status == "blocked" ? "approval-required" : "mode-changed",
          "repair_plan_issue" => overall_status == "pass" ? nil : "engine-binding-pending"
        }
      end

      def desktop_safe_summary
        return "Compatibility test execution is blocked and needs user action." if overall_status == "blocked"
        return "Compatibility test execution is waiting for Runtime-controlled preflight work." if overall_status == "pending"

        "Compatibility test execution completed successfully."
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @recipe_dir = RuntimeDaemon::DEFAULT_RECIPE_DIR
          @test_type = "preflight"
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--app is required" unless @application_id

          recipe = RecipeStore.new(path: @recipe_dir).find(@application_id)
          raise ArgumentError, "unknown application: #{@application_id}" unless recipe

          puts CompatibilityTestResult.new(recipe: recipe, test_type: @test_type).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-test-result: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-test-result --app APP_ID [--test-type TYPE] [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build a compatibility test result for APP_ID") do |value|
              @application_id = value
            end
            options.on("--test-type TYPE", CompatibilityTestPlan::TEST_TYPES, "Build a preflight, smoke, or repair-readiness result") do |value|
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

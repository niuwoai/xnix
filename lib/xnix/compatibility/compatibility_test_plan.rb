# frozen_string_literal: true

require "json"
require "optparse"
require_relative "compatibility_run_plan"
require_relative "compatibility_snapshot_plan"
require_relative "portal_request_model"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class CompatibilityTestPlan
      TEST_TYPES = %w[preflight smoke repair-readiness].freeze

      def initialize(recipe:, test_type: "preflight")
        @recipe = recipe
        @test_type = test_type
        validate!
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "plan_type" => "compatibility-test",
          "test_type" => @test_type,
          "desktop" => "KDE Plasma",
          "application_id" => recipe.id,
          "name" => recipe.name,
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "steps" => steps,
          "blocked" => blocked?,
          "blocking_reasons" => blocking_reasons,
          "artifacts" => artifacts,
          "desktop_safe_summary" => desktop_safe_summary,
          "backend_details_exposed" => false
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      attr_reader :recipe

      def validate!
        return if TEST_TYPES.include?(@test_type)

        raise ArgumentError, "test type must be one of: #{TEST_TYPES.join(", ")}"
      end

      def steps
        [
          recipe_step,
          portal_step,
          snapshot_step,
          engine_step
        ]
      end

      def recipe_step
        {
          "id" => "recipe-validation",
          "status" => "pass",
          "title" => "Validate application recipe",
          "summary" => "Recipe metadata can produce desktop integration artifacts.",
          "required_for" => %w[launcher file-association diagnostics]
        }
      end

      def portal_step
        portal_request = PortalRequestModel.new(application_id: recipe.id, operation: "file-open").to_h

        {
          "id" => "portal-preflight",
          "status" => portal_request.fetch("request_allowed") ? "pending" : "blocked",
          "title" => "Prepare user-mediated file access",
          "summary" => "File access must use an XDG Desktop Portal request before selected documents are opened.",
          "portal_request" => {
            "destination" => portal_request.fetch("portal").fetch("destination"),
            "interface" => portal_request.fetch("portal").fetch("interface"),
            "method" => portal_request.fetch("portal").fetch("method"),
            "handle_token" => portal_request.fetch("request").fetch("handle_token")
          }
        }
      end

      def snapshot_step
        snapshot = CompatibilitySnapshotPlan.new(application_id: recipe.id, reason: "before-engine-change").to_h

        {
          "id" => "snapshot-preflight",
          "status" => "pending",
          "title" => "Prepare restore point",
          "summary" => "Create an application restore point before risky compatibility changes.",
          "snapshot" => {
            "reason" => snapshot.fetch("reason"),
            "enabled_by_default" => snapshot.fetch("enabled_by_default"),
            "restore_available" => snapshot.fetch("restore").fetch("available")
          }
        }
      end

      def engine_step
        run_plan = CompatibilityRunPlan.new(recipe: recipe).to_h
        execution = run_plan.fetch("execution")
        backend_binding = execution.fetch("backend_binding")

        {
          "id" => "runtime-launch-binding",
          "status" => backend_binding.fetch("ready") ? "pass" : "pending",
          "title" => "Bind managed launch backend",
          "summary" => "Runtime launch backend binding is required before automated compatibility execution.",
          "run_plan" => {
            "strategy" => execution.fetch("strategy"),
            "backend_ready" => backend_binding.fetch("ready"),
            "backend_details_exposed" => execution.fetch("backend_details_exposed")
          }
        }
      end

      def blocked?
        steps.any? { |step| step.fetch("status") == "blocked" }
      end

      def blocking_reasons
        steps.select { |step| step.fetch("status") == "blocked" }.map { |step| step.fetch("summary") }
      end

      def artifacts
        {
          "compatibility_center_card" => true,
          "notification_event" => blocked? ? "approval-required" : "mode-changed",
          "repair_plan_issue" => blocked? ? "portal-approval-required" : "engine-binding-pending"
        }
      end

      def desktop_safe_summary
        return "Compatibility test plan is blocked by user-mediated desktop access." if blocked?

        "Compatibility test plan is ready for Runtime-controlled preflight work."
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

          puts CompatibilityTestPlan.new(recipe: recipe, test_type: @test_type).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-test-plan: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-test-plan --app APP_ID [--test-type TYPE] [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build a compatibility test plan for APP_ID") do |value|
              @application_id = value
            end
            options.on("--test-type TYPE", TEST_TYPES, "Build a preflight, smoke, or repair-readiness plan") do |value|
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

# frozen_string_literal: true

require "json"
require "optparse"
require_relative "compatibility_run_plan"
require_relative "compatibility_test_result"
require_relative "recipe_store"
require_relative "runtime_write_gate"

module Xnix
  module Compatibility
    class CompatibilityExecutionReadiness
      def initialize(recipe:)
        @recipe = recipe
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "readiness_type" => "compatibility-execution-readiness",
          "runtime_method" => "GetExecutionReadiness",
          "application" => application,
          "engine" => engine,
          "execution_state" => execution_state,
          "overall_status" => overall_status,
          "recommended_action" => recommended_action,
          "runtime_owned" => true,
          "c_runtime_backed" => true,
          "kde_policy_owner" => false,
          "compatibility_center_card" => true,
          "safe_for_ai_diagnostics" => true,
          "desktop_entry_launch_visible" => true,
          "launch_allowed" => launch_allowed?,
          "launch_enabled" => launch_enabled?,
          "execution_request_created" => false,
          "backend_binding_ready" => backend_binding_ready?,
          "portal_policy_required" => run_plan.fetch("preflight").fetch("portal_policy_required"),
          "snapshot_required" => run_plan.fetch("preflight").fetch("snapshot_before_risky_change"),
          "user_action_required" => !launch_allowed?,
          "gates" => gates,
          "blocked_actions" => blocked_actions,
          "host_root_modified" => false,
          "network_required" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => desktop_safe_summary
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      attr_reader :recipe

      def application
        run_plan.fetch("application").merge(
          "runtime_mode" => recipe.mode,
          "launcher_command" => "xnix-compat-launch --app #{recipe.id} %U"
        )
      end

      def engine
        run_plan.fetch("execution").fetch("engine")
      end

      def run_plan
        @run_plan ||= CompatibilityRunPlan.new(recipe: recipe).to_h
      end

      def test_result
        @test_result ||= CompatibilityTestResult.new(recipe: recipe, test_type: "preflight").to_h
      end

      def launch_gate
        @launch_gate ||= RuntimeWriteGate.new(method_name: "Launch").to_h
      end

      def backend_binding_ready?
        run_plan.fetch("execution").fetch("backend_binding").fetch("ready")
      end

      def launch_enabled?
        run_plan.fetch("execution").fetch("backend_binding").fetch("launch_enabled") &&
          launch_gate.fetch("dispatch_enabled")
      end

      def launch_allowed?
        backend_binding_ready? &&
          launch_gate.fetch("dispatch_enabled") &&
          test_result.fetch("counts").fetch("blocked").zero?
      end

      def execution_state
        launch_allowed? ? "ready" : "blocked"
      end

      def overall_status
        launch_allowed? ? "ready" : "not-ready"
      end

      def recommended_action
        if launch_allowed?
          "Create a Runtime-owned execution request."
        else
          "Complete Runtime-owned backend, Portal, snapshot, and Launch gates before execution."
        end
      end

      def gates
        [
          gate("recipe-validation", "pass", "Recipe metadata is loaded and can produce a desktop-safe execution plan."),
          gate("portal-policy-review", "required", "Portal-mediated resource access must be reviewed before execution."),
          gate("snapshot-baseline", "required", "A Runtime-managed restore point must exist before compatibility execution."),
          gate("backend-binding", backend_binding_ready? ? "pass" : "pending", "A managed compatibility backend binding is required before launch."),
          gate("runtime-launch-write-gate", launch_gate.fetch("dispatch_enabled") ? "pass" : "blocked", "Runtime Launch remains disabled until production backend ownership is ready.")
        ]
      end

      def gate(id, status, summary)
        {
          "id" => id,
          "status" => status,
          "summary" => summary
        }
      end

      def blocked_actions
        return [] if launch_allowed?

        [
          "create execution request",
          "launch compatibility backend",
          "expose raw backend command to desktop shell",
          "grant desktop resources without Portal review",
          "mutate host root during execution readiness planning"
        ]
      end

      def desktop_safe_summary
        if launch_allowed?
          "Compatibility execution is ready for a Runtime-owned launch request."
        else
          "Compatibility execution is not ready; KDE may show the desktop entry but must route launch intent through Runtime gates."
        end
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @recipe_dir = RuntimeDaemon::DEFAULT_RECIPE_DIR
          @application_id = nil
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--app is required" unless @application_id

          recipe = RecipeStore.new(path: @recipe_dir).find(@application_id)
          raise ArgumentError, "unknown application: #{@application_id}" unless recipe

          puts CompatibilityExecutionReadiness.new(recipe: recipe).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-execution-readiness: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-execution-readiness --app APP_ID [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build a Runtime execution readiness model for APP_ID") do |value|
              @application_id = value
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

# frozen_string_literal: true

require "json"
require "optparse"
require_relative "compatibility_run_plan"
require_relative "recipe_store"
require_relative "runtime_daemon"
require_relative "runtime_write_gate"

module Xnix
  module Compatibility
    class CompatibilityBackendLifecycle
      def initialize(recipe:)
        @recipe = recipe
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "lifecycle_type" => "compatibility-backend-lifecycle",
          "runtime_method" => "GetBackendLifecycle",
          "application" => application,
          "selected_strategy" => execution.fetch("strategy"),
          "lifecycle_state" => "blocked",
          "overall_status" => "not-ready",
          "runtime_owned" => true,
          "c_runtime_backed" => true,
          "kde_policy_owner" => false,
          "backend_binding_ready" => execution.fetch("backend_binding").fetch("ready"),
          "launch_enabled" => execution.fetch("backend_binding").fetch("launch_enabled") && launch_gate.fetch("dispatch_enabled"),
          "execution_request_created" => false,
          "backend_process_started" => false,
          "local_backend_started" => false,
          "isolated_backend_started" => false,
          "state_root_ready" => false,
          "portal_review_required" => preflight.fetch("portal_policy_required"),
          "snapshot_required" => preflight.fetch("snapshot_before_risky_change"),
          "stages" => stages,
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
        {
          "id" => recipe.id,
          "name" => recipe.name,
          "requested_mode" => recipe.mode,
          "supported_extensions" => recipe.supported_extensions
        }
      end

      def run_plan
        @run_plan ||= CompatibilityRunPlan.new(recipe: recipe).to_h
      end

      def execution
        run_plan.fetch("execution")
      end

      def preflight
        run_plan.fetch("preflight")
      end

      def launch_gate
        @launch_gate ||= RuntimeWriteGate.new(method_name: "Launch").to_h
      end

      def stages
        [
          {
            "id" => "recipe-loaded",
            "status" => "pass",
            "summary" => "Recipe metadata is loaded and mapped to a Runtime-owned backend lifecycle."
          },
          {
            "id" => "state-root-ready",
            "status" => "pending",
            "summary" => "A Runtime-owned application state root must exist before backend lifecycle activation."
          },
          {
            "id" => "backend-binding-ready",
            "status" => execution.fetch("backend_binding").fetch("ready") ? "pass" : "pending",
            "summary" => "A managed backend binding must be ready before any lifecycle start request."
          },
          {
            "id" => "portal-and-snapshot-review",
            "status" => "required",
            "summary" => "Portal access and restore-point policy must be reviewed before backend lifecycle activation."
          },
          {
            "id" => "runtime-launch-write-gate",
            "status" => launch_gate.fetch("dispatch_enabled") ? "pass" : "blocked",
            "summary" => "Runtime Launch write dispatch remains disabled until production backend ownership is ready."
          }
        ]
      end

      def blocked_actions
        [
          "start local compatibility backend from KDE",
          "start isolated compatibility backend from KDE",
          "create backend process before Runtime lifecycle gates pass",
          "expose backend command to desktop shell",
          "mutate host root during lifecycle planning"
        ]
      end

      def desktop_safe_summary
        "Runtime backend lifecycle is modeled but blocked; KDE may display state and must not start compatibility backends."
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @recipe_dir = RuntimeDaemon::DEFAULT_RECIPE_DIR
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--app is required" unless @application_id

          recipe = RecipeStore.new(path: @recipe_dir).find(@application_id)
          raise ArgumentError, "unknown application: #{@application_id}" unless recipe

          puts CompatibilityBackendLifecycle.new(recipe: recipe).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-backend-lifecycle: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-backend-lifecycle --app APP_ID [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build a managed compatibility backend lifecycle model for APP_ID") do |value|
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

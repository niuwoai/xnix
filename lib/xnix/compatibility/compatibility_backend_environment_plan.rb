# frozen_string_literal: true

require "json"
require "optparse"
require_relative "compatibility_run_plan"
require_relative "recipe_store"
require_relative "runtime_daemon"
require_relative "runtime_write_gate"

module Xnix
  module Compatibility
    class CompatibilityBackendEnvironmentPlan
      def initialize(recipe:)
        @recipe = recipe
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "plan_type" => "compatibility-backend-environment-plan",
          "runtime_method" => "GetBackendEnvironmentPlan",
          "application" => application,
          "selected_strategy" => execution.fetch("strategy"),
          "environment_state" => "planned",
          "runtime_owned" => true,
          "c_runtime_backed" => true,
          "kde_policy_owner" => false,
          "local_environment_ready" => false,
          "isolated_environment_ready" => false,
          "environment_created" => false,
          "backend_process_started" => false,
          "host_storage_exposed" => false,
          "clipboard_bridge_enabled" => false,
          "print_bridge_enabled" => false,
          "portal_review_required" => preflight.fetch("portal_policy_required"),
          "snapshot_required" => preflight.fetch("snapshot_before_risky_change"),
          "launch_enabled" => execution.fetch("backend_binding").fetch("launch_enabled") && launch_gate.fetch("dispatch_enabled"),
          "profiles" => profiles,
          "required_reviews" => required_reviews,
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

      def profiles
        [
          profile(
            id: "local-compatibility-environment",
            kind: "local",
            summary: "Local compatibility environment preparation waits for Runtime package, state, Portal, snapshot, and Launch gates."
          ),
          profile(
            id: "isolated-compatibility-environment",
            kind: "isolated",
            summary: "Isolated compatibility environment preparation waits for Runtime package, state, Portal, snapshot, and Launch gates."
          )
        ]
      end

      def profile(id:, kind:, summary:)
        {
          "id" => id,
          "kind" => kind,
          "status" => "blocked",
          "selected" => selected_profile?(kind),
          "environment_created" => false,
          "process_started" => false,
          "host_storage_exposed" => false,
          "backend_details_exposed" => false,
          "summary" => summary
        }
      end

      def selected_profile?(kind)
        execution.fetch("strategy") == "#{kind}-compatibility-engine"
      end

      def required_reviews
        %w[
          package-source-review
          application-state-root-review
          portal-policy-review
          snapshot-baseline-review
        ]
      end

      def blocked_actions
        [
          "create local compatibility environment from KDE",
          "create isolated compatibility environment from KDE",
          "start backend process before environment gates pass",
          "bridge clipboard or print devices without Portal review",
          "expose backend environment commands to desktop shell"
        ]
      end

      def desktop_safe_summary
        "Backend environments are planned by the Runtime but not created; KDE may display readiness only."
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

          puts CompatibilityBackendEnvironmentPlan.new(recipe: recipe).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-backend-environment-plan: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-backend-environment-plan --app APP_ID [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build a managed compatibility backend environment plan for APP_ID") do |value|
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

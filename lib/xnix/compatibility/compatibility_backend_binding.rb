# frozen_string_literal: true

require "json"
require "optparse"
require_relative "compatibility_run_plan"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class CompatibilityBackendBinding
      def initialize(recipe:)
        @recipe = recipe
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "binding_type" => "compatibility-backend-binding",
          "application" => application,
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "selected_strategy" => run_plan.fetch("execution").fetch("strategy"),
          "managed_binding_ready" => false,
          "launch_enabled" => false,
          "execution_request_created" => false,
          "host_root_modified" => false,
          "network_required" => false,
          "privileged_container_required" => false,
          "required_preflight" => required_preflight,
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

      def required_preflight
        [
          {
            "id" => "engine-package-source",
            "status" => "pending",
            "summary" => "A managed compatibility engine package source must be selected by the Runtime."
          },
          {
            "id" => "application-state-root",
            "status" => "pending",
            "summary" => "A Runtime-owned application state root must be allocated before launch binding."
          },
          {
            "id" => "portal-policy-review",
            "status" => "required",
            "summary" => "Portal access policy must be reviewed before file and desktop resource bridging."
          },
          {
            "id" => "snapshot-baseline",
            "status" => "required",
            "summary" => "A baseline restore point must exist before managed compatibility execution."
          }
        ]
      end

      def blocked_actions
        [
          "launch compatibility engine without managed binding",
          "expose backend command to desktop shell",
          "write application state outside Runtime ownership",
          "grant desktop resources without Portal policy"
        ]
      end

      def desktop_safe_summary
        "Managed compatibility backend binding is pending Runtime preflight."
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

          puts CompatibilityBackendBinding.new(recipe: recipe).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-backend-binding: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-backend-binding --app APP_ID [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build a managed compatibility backend binding model for APP_ID") do |value|
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

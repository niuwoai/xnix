# frozen_string_literal: true

require "json"
require "optparse"
require_relative "application_recipe"
require_relative "compatibility_engine_catalog"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class CompatibilityRunPlan
      STRATEGIES = {
        "automatic" => {
          "strategy" => "automatic-managed",
          "summary" => "Xnix will choose the best available compatibility path."
        },
        "wine" => {
          "strategy" => "local-compatibility-engine",
          "summary" => "Xnix will use a local compatibility engine when backend binding is available."
        },
        "vm" => {
          "strategy" => "isolated-compatibility-engine",
          "summary" => "Xnix will use an isolated compatibility engine when backend binding is available."
        }
      }.freeze

      attr_reader :recipe

      def initialize(recipe:)
        @recipe = recipe
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "plan_type" => "compatibility-run",
          "application" => application,
          "execution" => execution,
          "preflight" => preflight,
          "desktop_safe_summary" => strategy.fetch("summary")
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def application
        {
          "id" => recipe.id,
          "name" => recipe.name,
          "icon" => recipe.icon
        }
      end

      def execution
        engine = CompatibilityEngineCatalog.new.select_for_mode(recipe.mode)

        {
          "strategy" => engine.fetch("engine_id"),
          "selection_source" => "recipe",
          "engine" => engine,
          "backend_details_exposed" => false,
          "backend_binding" => {
            "ready" => engine.fetch("ready"),
            "launch_enabled" => engine.fetch("launch_enabled"),
            "reason" => "Compatibility backend binding is pending."
          }
        }
      end

      def preflight
        {
          "portal_policy_required" => true,
          "snapshot_before_risky_change" => true,
          "diagnostics_required" => true
        }
      end

      def strategy
        STRATEGIES.fetch(recipe.mode)
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

          puts CompatibilityRunPlan.new(recipe: recipe).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-run-plan: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-run-plan --app APP_ID [--recipe-dir PATH]"
            options.on("--recipe-dir PATH", "Read application recipes from PATH") do |value|
              @recipe_dir = value
            end
            options.on("--app APP_ID", "Build a compatibility run plan for APP_ID") do |value|
              @application_id = value
            end
          end
        end
      end
    end
  end
end

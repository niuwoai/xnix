# frozen_string_literal: true

require "json"
require "optparse"
require "uri"
require_relative "compatibility_run_plan"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class LaunchRequest
      SOURCE = "desktop-launcher"

      attr_reader :recipe_store

      def initialize(recipe_store:)
        @recipe_store = recipe_store
      end

      def build(application_id:, file_uris: [])
        recipe = require_recipe(application_id)
        uris = normalize_file_uris(file_uris)

        {
          "request_type" => "launch-application",
          "source" => SOURCE,
          "application_id" => recipe.id,
          "application_name" => recipe.name,
          "runtime_method" => "Launch",
          "portal_required" => !uris.empty?,
          "run_plan" => run_plan_summary(recipe),
          "file_count" => uris.length,
          "file_uris" => uris
        }
      end

      private

      def normalize_file_uris(file_uris)
        file_uris.map { |file_uri| normalize_file_uri(file_uri) }
      end

      def normalize_file_uri(file_uri)
        uri = URI.parse(file_uri)
        raise ArgumentError, "only file URIs are accepted" unless uri.scheme == "file"
        raise ArgumentError, "file URI must include an absolute path" if uri.path.nil? || uri.path.empty? || !uri.path.start_with?("/")

        uri.to_s
      rescue URI::InvalidURIError
        raise ArgumentError, "invalid file URI: #{file_uri.inspect}"
      end

      def require_recipe(application_id)
        recipe = recipe_store.find(application_id)
        return recipe if recipe

        raise ArgumentError, "unknown application: #{application_id}"
      end

      def run_plan_summary(recipe)
        plan = CompatibilityRunPlan.new(recipe: recipe).to_h
        {
          "plan_type" => plan.fetch("plan_type"),
          "strategy" => plan.fetch("execution").fetch("strategy"),
          "backend_details_exposed" => plan.fetch("execution").fetch("backend_details_exposed"),
          "backend_ready" => plan.fetch("execution").fetch("backend_binding").fetch("ready"),
          "portal_policy_required" => plan.fetch("preflight").fetch("portal_policy_required"),
          "snapshot_before_risky_change" => plan.fetch("preflight").fetch("snapshot_before_risky_change")
        }
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

          request = LaunchRequest.new(recipe_store: RecipeStore.new(path: @recipe_dir))
                                 .build(application_id: @application_id, file_uris: @argv)
          puts JSON.pretty_generate(request)
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-launch: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-launch --app APP_ID [--recipe-dir PATH] [FILE_URI ...]"
            options.on("--recipe-dir PATH", "Read application recipes from PATH") do |value|
              @recipe_dir = value
            end
            options.on("--app APP_ID", "Launch a Runtime application") do |value|
              @application_id = value
            end
          end
        end
      end
    end
  end
end

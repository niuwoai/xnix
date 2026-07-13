# frozen_string_literal: true

require "json"
require "optparse"
require_relative "compatibility_package_source"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class CompatibilityAcquisitionPreflight
      def initialize(recipe:)
        @recipe = recipe
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "preflight_type" => "compatibility-acquisition-preflight",
          "application" => application,
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "selected_strategy" => package_source.fetch("selected_strategy"),
          "preflight_state" => "planned",
          "acquisition_ready" => false,
          "download_enabled" => false,
          "install_enabled" => false,
          "network_required_for_planning" => false,
          "network_request_created" => false,
          "artifacts_downloaded" => false,
          "host_root_modified" => false,
          "privileged_container_required" => false,
          "desktop_shell_command_exposed" => false,
          "package_source_ready" => package_source.fetch("package_source_ready"),
          "checks" => checks,
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
          "requested_mode" => recipe.mode
        }
      end

      def package_source
        @package_source ||= CompatibilityPackageSource.new(recipe: recipe).to_h
      end

      def checks
        [
          {
            "id" => "package-source-ready",
            "status" => "pending",
            "summary" => "Runtime package source selection must be ready before acquisition."
          },
          {
            "id" => "signed-artifact-manifest",
            "status" => "required",
            "summary" => "Runtime must verify a signed artifact manifest before acquisition."
          },
          {
            "id" => "runtime-cache-space",
            "status" => "pending",
            "summary" => "Runtime cache capacity must be checked before artifact acquisition."
          },
          {
            "id" => "network-policy-review",
            "status" => "required",
            "summary" => "Runtime must approve network policy before any acquisition request."
          },
          {
            "id" => "rollback-marker",
            "status" => "required",
            "summary" => "Runtime must define rollback markers before acquisition can change state."
          }
        ]
      end

      def blocked_actions
        [
          "download compatibility artifacts before acquisition preflight",
          "install compatibility artifacts before signed manifest verification",
          "invoke network access from KDE",
          "mutate host root during acquisition preflight"
        ]
      end

      def desktop_safe_summary
        "Compatibility acquisition preflight is planned and waiting for Runtime source readiness."
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

          puts CompatibilityAcquisitionPreflight.new(recipe: recipe).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-acquisition-preflight: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-acquisition-preflight --app APP_ID [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build a Runtime compatibility acquisition preflight model for APP_ID") do |value|
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

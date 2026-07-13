# frozen_string_literal: true

require "json"
require "optparse"
require_relative "compatibility_engine_catalog"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class CompatibilityPackageSource
      def initialize(recipe:)
        @recipe = recipe
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "source_type" => "compatibility-package-source",
          "application" => application,
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "selected_strategy" => selected_engine.fetch("engine_id"),
          "source_selection_state" => "planned",
          "package_source_ready" => false,
          "install_enabled" => false,
          "network_required_for_planning" => false,
          "host_root_modified" => false,
          "privileged_container_required" => false,
          "desktop_shell_command_exposed" => false,
          "source_policy" => source_policy,
          "source_channels" => source_channels,
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
          "requested_mode" => recipe.mode
        }
      end

      def selected_engine
        @selected_engine ||= CompatibilityEngineCatalog.new.select_for_mode(recipe.mode)
      end

      def source_policy
        {
          "signed_source_required" => true,
          "runtime_cache_required" => true,
          "direct_desktop_install_allowed" => false,
          "user_visible_backend_names" => false,
          "host_package_manager_invoked" => false
        }
      end

      def source_channels
        [
          {
            "id" => "os-managed-compatibility-packages",
            "kind" => "distribution-packages",
            "runtime_owned" => true,
            "selection_state" => "planned",
            "supported_strategies" => ["local-compatibility-engine"],
            "summary" => "Distribution-provided compatibility packages can be selected only through Runtime policy."
          },
          {
            "id" => "runtime-managed-toolcache",
            "kind" => "runtime-cache",
            "runtime_owned" => true,
            "selection_state" => "planned",
            "supported_strategies" => ["automatic-managed", "local-compatibility-engine", "isolated-compatibility-engine"],
            "summary" => "Runtime-managed tool cache keeps package selection outside desktop shell code."
          },
          {
            "id" => "isolated-environment-template-catalog",
            "kind" => "template-catalog",
            "runtime_owned" => true,
            "selection_state" => "planned",
            "supported_strategies" => ["isolated-compatibility-engine"],
            "summary" => "Isolated environment templates remain Runtime-owned and are not launched during planning."
          }
        ]
      end

      def required_preflight
        [
          {
            "id" => "signed-source-verification",
            "status" => "required",
            "summary" => "Runtime must verify a signed source before package installation is enabled."
          },
          {
            "id" => "source-policy-review",
            "status" => "pending",
            "summary" => "Runtime policy must select the package source before launch binding."
          },
          {
            "id" => "runtime-cache-quota",
            "status" => "pending",
            "summary" => "Runtime cache quota must be checked before package acquisition."
          },
          {
            "id" => "offline-fallback",
            "status" => "pending",
            "summary" => "Runtime must define the offline behavior before package acquisition."
          }
        ]
      end

      def blocked_actions
        [
          "install compatibility packages without Runtime source selection",
          "expose package manager commands to KDE",
          "use unsigned package sources",
          "mutate the host root during package-source planning"
        ]
      end

      def desktop_safe_summary
        "Compatibility package source selection is planned and Runtime-owned."
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

          puts CompatibilityPackageSource.new(recipe: recipe).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-package-source: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-package-source --app APP_ID [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build a Runtime compatibility package source model for APP_ID") do |value|
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

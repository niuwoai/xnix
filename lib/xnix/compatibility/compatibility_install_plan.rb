# frozen_string_literal: true

require "json"
require "optparse"
require_relative "application_state_root"
require_relative "compatibility_acquisition_preflight"
require_relative "compatibility_artifact_manifest"
require_relative "compatibility_package_source"
require_relative "recipe_install_gate"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class CompatibilityInstallPlan
      def initialize(recipe:, environment: "development")
        @recipe = recipe
        @environment = environment
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "plan_type" => "compatibility-install-plan",
          "application" => application,
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "environment" => environment,
          "install_state" => "planned",
          "install_ready" => false,
          "desktop_activation_ready" => false,
          "download_enabled" => false,
          "install_enabled" => false,
          "network_request_created" => false,
          "artifacts_downloaded" => false,
          "host_root_modified" => false,
          "privileged_container_required" => false,
          "desktop_shell_command_exposed" => false,
          "selected_strategy" => artifact_manifest.fetch("selected_strategy"),
          "readiness" => readiness,
          "phases" => phases,
          "blocked_actions" => blocked_actions,
          "desktop_safe_summary" => desktop_safe_summary,
          "backend_details_exposed" => false
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      attr_reader :recipe, :environment

      def application
        {
          "id" => recipe.id,
          "name" => recipe.name,
          "requested_mode" => recipe.mode,
          "supported_extensions" => recipe.supported_extensions,
          "mime_types" => recipe.mime_types
        }
      end

      def readiness
        {
          "artifact_manifest_ready" => artifact_manifest.fetch("manifest_ready"),
          "artifact_signature_verified" => artifact_manifest.fetch("signature_verified"),
          "acquisition_ready" => acquisition_preflight.fetch("acquisition_ready"),
          "package_source_ready" => package_source.fetch("package_source_ready"),
          "state_root_allocated" => state_root.fetch("allocation_state") == "allocated",
          "recipe_install_allowed" => install_gate.fetch("decision") == "allow",
          "recipe_install_decision" => install_gate.fetch("decision")
        }
      end

      def phases
        [
          phase("resolve-artifact-manifest", "blocked", "Wait for a signed compatibility artifact manifest."),
          phase("verify-artifact-digests", "blocked", "Wait for digest verification before artifact activation."),
          phase("prepare-package-source", "blocked", "Wait for Runtime-owned package source readiness."),
          phase("allocate-application-state", "blocked", "Wait for Runtime-owned application state allocation."),
          phase("stage-desktop-integration", "pending", "Stage launcher, file association, Dolphin, tray, notification, Compatibility Center, and settings artifacts only after install gates pass."),
          phase("enable-launch-binding", "blocked", "Backend launch binding stays disabled until install preflight completes.")
        ]
      end

      def phase(id, status, summary)
        {
          "id" => id,
          "status" => status,
          "summary" => summary
        }
      end

      def blocked_actions
        [
          "download artifacts before signed manifest verification",
          "install packages before Runtime install plan readiness",
          "stage desktop integration before recipe install gate approval",
          "launch backend before managed binding readiness",
          "expose backend commands or storage paths to KDE",
          "mutate the host root during install planning"
        ]
      end

      def desktop_safe_summary
        "Compatibility install is planned and waiting for Runtime-owned readiness gates."
      end

      def artifact_manifest
        @artifact_manifest ||= CompatibilityArtifactManifest.new(recipe: recipe).to_h
      end

      def acquisition_preflight
        @acquisition_preflight ||= CompatibilityAcquisitionPreflight.new(recipe: recipe).to_h
      end

      def package_source
        @package_source ||= CompatibilityPackageSource.new(recipe: recipe).to_h
      end

      def state_root
        @state_root ||= ApplicationStateRoot.new(recipe: recipe).to_h
      end

      def install_gate
        @install_gate ||= RecipeInstallGate.new(
          application_id: recipe.id,
          mode: environment,
          registry_report: development_registry_report
        ).to_h
      end

      def development_registry_report
        {
          "version" => RuntimeDaemon::VERSION,
          "schema_version" => RecipeRegistry::SCHEMA_VERSION,
          "registry_name" => "xnix-install-plan-development",
          "registry_path" => "runtime/recipes/registry.json",
          "recipe_count" => 1,
          "trust" => {
            "digest_verified" => true,
            "signed_recipe_validation" => false,
            "development_registry" => true
          },
          "recipes" => [
            {
              "id" => recipe.id,
              "path" => "#{recipe.id}.json",
              "sha256" => "0" * 64,
              "signature_status" => "development-only",
              "digest_verified" => true
            }
          ]
        }
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @recipe_dir = RuntimeDaemon::DEFAULT_RECIPE_DIR
          @environment = "development"
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--app is required" unless @application_id

          recipe = RecipeStore.new(path: @recipe_dir).find(@application_id)
          raise ArgumentError, "unknown application: #{@application_id}" unless recipe

          puts CompatibilityInstallPlan.new(recipe: recipe, environment: @environment).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-install-plan: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-install-plan --app APP_ID [--recipe-dir PATH] [--environment NAME]"
            options.on("--app APP_ID", "Build a Runtime-owned compatibility install plan for APP_ID") do |value|
              @application_id = value
            end
            options.on("--recipe-dir PATH", "Read application recipes from PATH") do |value|
              @recipe_dir = value
            end
            options.on("--environment NAME", "Evaluate install gates for development or production") do |value|
              @environment = value
            end
          end
        end
      end
    end
  end
end

# frozen_string_literal: true

require "json"
require "optparse"
require_relative "compatibility_acquisition_preflight"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class CompatibilityArtifactManifest
      def initialize(recipe:)
        @recipe = recipe
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "manifest_type" => "compatibility-artifact-manifest",
          "application" => application,
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "selected_strategy" => acquisition_preflight.fetch("selected_strategy"),
          "manifest_state" => "planned",
          "manifest_ready" => false,
          "signature_verified" => false,
          "acquisition_preflight_ready" => acquisition_preflight.fetch("acquisition_ready"),
          "download_enabled" => false,
          "install_enabled" => false,
          "network_request_created" => false,
          "artifacts_downloaded" => false,
          "host_root_modified" => false,
          "privileged_container_required" => false,
          "desktop_shell_command_exposed" => false,
          "artifact_groups" => artifact_groups,
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

      def acquisition_preflight
        @acquisition_preflight ||= CompatibilityAcquisitionPreflight.new(recipe: recipe).to_h
      end

      def artifact_groups
        [
          {
            "id" => "runtime-launch-metadata",
            "kind" => "metadata",
            "required" => true,
            "resolved" => false,
            "downloaded" => false,
            "cache_namespace" => cache_namespace("launch-metadata"),
            "summary" => "Runtime launch metadata must be resolved before launch binding."
          },
          {
            "id" => "local-execution-artifacts",
            "kind" => "execution-artifacts",
            "required" => true,
            "resolved" => false,
            "downloaded" => false,
            "cache_namespace" => cache_namespace("local-execution"),
            "summary" => "Local execution artifacts remain planned until signed manifest verification passes."
          },
          {
            "id" => "isolated-environment-artifacts",
            "kind" => "environment-artifacts",
            "required" => false,
            "resolved" => false,
            "downloaded" => false,
            "cache_namespace" => cache_namespace("isolated-environment"),
            "summary" => "Isolated environment artifacts remain optional until Runtime policy selects them."
          }
        ]
      end

      def required_preflight
        [
          {
            "id" => "acquisition-preflight-ready",
            "status" => "pending",
            "summary" => "Runtime acquisition preflight must be ready before artifact manifest resolution."
          },
          {
            "id" => "manifest-signature-verification",
            "status" => "required",
            "summary" => "Runtime must verify the signed artifact manifest before artifact use."
          },
          {
            "id" => "artifact-digest-verification",
            "status" => "required",
            "summary" => "Runtime must verify artifact digests before cache activation."
          },
          {
            "id" => "cache-namespace-allocation",
            "status" => "pending",
            "summary" => "Runtime must allocate cache namespaces before artifact acquisition."
          },
          {
            "id" => "rollback-reference",
            "status" => "required",
            "summary" => "Runtime must record rollback references before artifacts can affect state."
          }
        ]
      end

      def blocked_actions
        [
          "download artifacts before signed manifest verification",
          "activate artifacts before digest verification",
          "expose artifact cache paths to KDE",
          "mutate host root during artifact manifest planning"
        ]
      end

      def cache_namespace(suffix)
        "#{recipe.id.gsub(/[^a-zA-Z0-9.-]/, "-")}.#{suffix}"
      end

      def desktop_safe_summary
        "Compatibility artifact manifest is planned and waiting for acquisition preflight."
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

          puts CompatibilityArtifactManifest.new(recipe: recipe).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-artifact-manifest: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-artifact-manifest --app APP_ID [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build a Runtime compatibility artifact manifest model for APP_ID") do |value|
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

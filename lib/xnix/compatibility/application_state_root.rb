# frozen_string_literal: true

require "json"
require "optparse"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class ApplicationStateRoot
      def initialize(recipe:)
        @recipe = recipe
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "root_type" => "compatibility-application-state-root",
          "application" => application,
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "state_namespace" => state_namespace,
          "storage_scope" => "per-application",
          "allocation_state" => "planned",
          "directories_created" => false,
          "host_root_modified" => false,
          "user_documents_included" => false,
          "portal_required_for_user_files" => true,
          "snapshot_eligible" => true,
          "restore_requires_confirmation" => true,
          "retention_policy" => retention_policy,
          "managed_scopes" => managed_scopes,
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

      def state_namespace
        recipe.id.gsub(/[^a-zA-Z0-9.-]/, "-")
      end

      def retention_policy
        {
          "automatic_restore_points" => 5,
          "manual_restore_points" => 10,
          "user_documents_excluded" => true
        }
      end

      def managed_scopes
        [
          {
            "id" => "application-data",
            "runtime_owned" => true,
            "snapshot_included" => true,
            "summary" => "Application-managed state is isolated under Runtime ownership."
          },
          {
            "id" => "runtime-metadata",
            "runtime_owned" => true,
            "snapshot_included" => true,
            "summary" => "Runtime metadata tracks compatibility state without exposing host paths."
          },
          {
            "id" => "diagnostic-cache",
            "runtime_owned" => true,
            "snapshot_included" => false,
            "summary" => "Diagnostic cache is Runtime-owned and can be regenerated."
          },
          {
            "id" => "desktop-activation-receipts",
            "runtime_owned" => true,
            "snapshot_included" => true,
            "summary" => "Desktop activation receipts remain part of rollback planning."
          }
        ]
      end

      def blocked_actions
        [
          "write application state outside Runtime ownership",
          "include user documents in state snapshots",
          "expose host storage paths to KDE",
          "restore state without user confirmation"
        ]
      end

      def desktop_safe_summary
        "Runtime application state root is planned and isolated from user documents."
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

          puts ApplicationStateRoot.new(recipe: recipe).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-state-root: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-state-root --app APP_ID [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build a Runtime application state root model for APP_ID") do |value|
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

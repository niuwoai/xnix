# frozen_string_literal: true

require "json"
require "optparse"
require_relative "desktop_entry"
require_relative "dolphin_service_menu"
require_relative "kde_integration_status"
require_relative "portal_access_policy"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class DesktopIntegrationManifest
      MANIFEST_TYPE = "desktop-integration"
      DESKTOP = "KDE Plasma"
      LAUNCHER_COMMAND = "xnix-compat-launch"
      WINDOW_IDENTITY_COMMAND = "xnix-compat-window-identity"
      KWIN_WINDOW_RULE_COMMAND = "xnix-kwin-window-rule"
      FILE_OPEN_COMMAND = "xnix-compat-open"
      TRAY_STATUS_COMMAND = "xnix-compat-tray-status"
      NOTIFICATION_COMMAND = "xnix-compat-notify"
      CENTER_MODEL_COMMAND = "xnix-kde-center-model"
      SETTINGS_COMMAND = "xnix-compat-settings"
      PORTAL_POLICY_COMMAND = "xnix-portal-access-policy"

      def initialize(recipe:)
        @recipe = recipe
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "manifest_type" => MANIFEST_TYPE,
          "desktop" => DESKTOP,
          "official_desktop_only" => true,
          "application" => application,
          "entry_points" => entry_points,
          "artifacts" => artifacts,
          "safety" => safety
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
          "icon" => recipe.icon,
          "supported_extensions" => recipe.supported_extensions,
          "mime_types" => recipe.mime_types
        }
      end

      def entry_points
        KdeIntegrationStatus::ENTRY_POINTS.map { |entry| entry.fetch("id") }
      end

      def artifacts
        [
          launcher_artifact,
          task_manager_artifact,
          file_manager_artifact,
          system_tray_artifact,
          notifications_artifact,
          compatibility_center_artifact,
          settings_artifact
        ]
      end

      def launcher_artifact
        desktop_entry = DesktopEntry.new(recipe)

        {
          "entry_point" => "launcher",
          "kind" => "desktop-entry",
          "path" => File.join("applications", desktop_entry.file_name),
          "argv" => [LAUNCHER_COMMAND, "--app", recipe.id, "%U"],
          "desktop_file" => desktop_entry.file_name
        }
      end

      def task_manager_artifact
        {
          "entry_point" => "task-manager",
          "kind" => "window-identity",
          "argv" => [
            WINDOW_IDENTITY_COMMAND,
            "--app",
            recipe.id,
            "--name",
            recipe.name
          ],
          "kwin_rule" => {
            "argv" => [
              KWIN_WINDOW_RULE_COMMAND,
              "--app",
              recipe.id,
              "--name",
              recipe.name
            ],
            "script_role" => "identity-and-layout"
          },
          "desktop_file" => DesktopEntry.new(recipe).file_name
        }
      end

      def file_manager_artifact
        service_menu = DolphinServiceMenu.new

        {
          "entry_point" => "file-manager",
          "kind" => "dolphin-service-menu",
          "path" => File.join("servicemenus", service_menu.file_name),
          "argv" => [FILE_OPEN_COMMAND, "%U"],
          "portal_required" => true
        }
      end

      def system_tray_artifact
        {
          "entry_point" => "system-tray",
          "kind" => "tray-status-model",
          "argv" => [TRAY_STATUS_COMMAND],
          "runtime_attention" => true
        }
      end

      def notifications_artifact
        {
          "entry_point" => "notifications",
          "kind" => "notification-request-model",
          "argv" => [NOTIFICATION_COMMAND, "--app", recipe.id],
          "event_types" => %w[install-failed repair-applied mode-changed approval-required]
        }
      end

      def compatibility_center_artifact
        {
          "entry_point" => "compatibility-center",
          "kind" => "read-model",
          "argv" => [CENTER_MODEL_COMMAND],
          "application_reference" => recipe.id
        }
      end

      def settings_artifact
        {
          "entry_point" => "settings",
          "kind" => "settings-model",
          "argv" => [SETTINGS_COMMAND, recipe.id],
          "portal_policy" => {
            "argv" => [PORTAL_POLICY_COMMAND, "--app", recipe.id, "--operation", "file-open"],
            "operations" => PortalAccessPolicy::OPERATIONS.keys
          },
          "sections" => %w[run-mode resource-access devices network snapshots]
        }
      end

      def safety
        {
          "backend_commands_exposed" => false,
          "portal_required_for_file_access" => true,
          "host_privilege_required" => false,
          "stable_desktop_contract" => true
        }
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

          puts DesktopIntegrationManifest.new(recipe: recipe).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-desktop-integration-manifest: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-desktop-integration-manifest --app APP_ID [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build a desktop integration manifest for APP_ID") do |value|
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

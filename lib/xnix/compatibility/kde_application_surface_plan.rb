# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"
require_relative "recipe_store"

module Xnix
  module Compatibility
    class KDEApplicationSurfacePlan
      PROJECT_ROOT = Pathname.new(__dir__).join("../../..").realpath
      VERSION = PROJECT_ROOT.join("VERSION").read.strip
      DEFAULT_RECIPE_DIR = PROJECT_ROOT.join("runtime/recipes").to_s

      ENTRY_POINTS = [
        ["launcher", "Launcher", "GetDesktopEntryPlan"],
        ["task-manager", "Task Manager", "GetTaskManagerIdentityPlan"],
        ["file-manager", "File Manager", "GetFileAssociationPlan"],
        ["system-tray", "System Tray", "GetTrayStatus"],
        ["notifications", "Notifications", "GetNotificationPlan"],
        ["compatibility-center", "Compatibility Center", "GetCompatibilityCenterSummary"],
        ["settings", "Settings", "GetCompatibilitySettings"]
      ].freeze

      REQUIRED_RUNTIME_GATES = %w[
        recipe-install-gate
        portal-policy-review
        snapshot-baseline
        backend-environment-plan
        backend-lifecycle-plan
        runtime-write-gate
      ].freeze

      def initialize(recipe:)
        @recipe = recipe
      end

      def to_h
        {
          "version" => VERSION,
          "plan_type" => "kde-application-surface-plan",
          "runtime_method" => "GetKDEApplicationSurfacePlan",
          "application" => application,
          "surface_state" => "planned",
          "desktop_shell" => "KDE Plasma",
          "entry_points" => entry_points,
          "entry_point_count" => ENTRY_POINTS.length,
          "required_runtime_gates" => REQUIRED_RUNTIME_GATES,
          "runtime_owned" => true,
          "c_runtime_backed" => true,
          "kde_policy_owner" => false,
          "official_desktop_only" => true,
          "normal_linux_application_surface" => true,
          "standard_launcher_visible" => true,
          "task_manager_identity_ready" => true,
          "file_associations_planned" => true,
          "dolphin_action_planned" => true,
          "krunner_query_planned" => true,
          "tray_status_planned" => true,
          "notification_route_planned" => true,
          "settings_surface_planned" => true,
          "portal_review_required" => true,
          "execution_ready" => false,
          "launch_enabled" => false,
          "backend_process_started" => false,
          "desktop_files_written" => false,
          "mimeapps_written" => false,
          "host_root_modified" => false,
          "backend_command_exposed" => false,
          "raw_windows_executable_exposed" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => "KDE can present this compatibility application as a normal Linux application while Runtime gates still block execution."
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
          "requested_mode" => recipe.mode,
          "supported_extensions" => recipe.supported_extensions
        }
      end

      def entry_points
        ENTRY_POINTS.map do |id, name, runtime_method|
          {
            "id" => id,
            "name" => name,
            "runtime_method" => runtime_method,
            "state" => "planned",
            "runtime_backed" => true,
            "c_runtime_backed" => true,
            "kde_writes_policy" => false,
            "backend_details_exposed" => false
          }
        end
      end
    end

    class KDEApplicationSurfacePlanCLI
      def initialize(argv)
        @argv = argv
        @recipe_dir = KDEApplicationSurfacePlan::DEFAULT_RECIPE_DIR
      end

      def run
        app_id = nil

        OptionParser.new do |options|
          options.banner = "Usage: xnix-kde-application-surface-plan --app APP_ID [--recipe-dir PATH]"
          options.on("--app APP_ID", "Application id") { |value| app_id = value }
          options.on("--recipe-dir PATH", "Recipe directory") { |value| @recipe_dir = value }
        end.parse!(@argv)

        raise ArgumentError, "--app is required" unless app_id

        store = RecipeStore.new(path: @recipe_dir)
        recipe = store.find(app_id)
        raise ArgumentError, "unknown application: #{app_id}" unless recipe

        puts KDEApplicationSurfacePlan.new(recipe: recipe).to_json
      rescue ArgumentError => e
        warn "xnix-kde-application-surface-plan: #{e.message}"
        64
      else
        0
      end
    end
  end
end

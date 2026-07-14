# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"
require_relative "portal_access_policy"
require_relative "recipe_store"

module Xnix
  module Compatibility
    class CompatibilityPermissionReviewPlan
      PROJECT_ROOT = Pathname.new(__dir__).join("../../..").realpath
      VERSION = PROJECT_ROOT.join("VERSION").read.strip
      DEFAULT_RECIPE_DIR = PROJECT_ROOT.join("runtime/recipes").to_s

      PERMISSIONS = [
        ["documents", "Documents", "file-open", "ask", "org.freedesktop.portal.FileChooser", true],
        ["downloads", "Downloads", "file-open", "ask", "org.freedesktop.portal.FileChooser", true],
        ["camera", "Camera", "camera", "deny", "org.freedesktop.portal.Camera", true],
        ["network", "Network", "network", "allow", "none", false],
        ["clipboard", "Clipboard", "clipboard", "ask", "org.freedesktop.portal.Clipboard", true],
        ["print", "Print", "print", "ask", "org.freedesktop.portal.Print", true],
        ["screenshot", "Screenshot", "screenshot", "ask", "org.freedesktop.portal.Screenshot", true]
      ].freeze

      REQUIRED_RUNTIME_GATES = %w[
        user-review
        portal-policy-review
        runtime-write-gate
        settings-persistence
        audit-log
      ].freeze

      def initialize(recipe:)
        @recipe = recipe
      end

      def to_h
        permission_rows = permissions

        {
          "version" => VERSION,
          "plan_type" => "compatibility-permission-review-plan",
          "runtime_method" => "GetCompatibilityPermissionReviewPlan",
          "application" => application,
          "review_state" => "planned",
          "permissions" => permission_rows,
          "permission_count" => permission_rows.length,
          "allow_count" => permission_rows.count { |row| row.fetch("decision") == "allow" },
          "ask_count" => permission_rows.count { |row| row.fetch("decision") == "ask" },
          "deny_count" => permission_rows.count { |row| row.fetch("decision") == "deny" },
          "required_runtime_gates" => REQUIRED_RUNTIME_GATES,
          "runtime_owned" => true,
          "c_runtime_backed" => true,
          "kde_policy_owner" => false,
          "user_review_required" => true,
          "portal_review_required" => true,
          "permission_changes_applied" => false,
          "request_objects_created" => false,
          "permissions_granted" => false,
          "settings_persisted" => false,
          "host_permission_changed" => false,
          "host_root_modified" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => "Compatibility permissions are grouped for KDE review and remain unchanged until Runtime gates pass."
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
          "icon" => recipe.icon
        }
      end

      def permissions
        PERMISSIONS.map do |id, label, operation, decision, portal_interface, portal_required|
          {
            "id" => id,
            "label" => label,
            "operation" => operation,
            "decision" => decision,
            "portal_interface" => portal_interface,
            "portal_required" => portal_required,
            "user_mediation_required" => true,
            "current_value" => decision,
            "requested_value" => decision,
            "change_pending" => false,
            "request_object_created" => false,
            "permission_granted" => false,
            "direct_access_allowed" => false,
            "backend_details_exposed" => false
          }
        end
      end
    end

    class CompatibilityPermissionReviewPlanCLI
      def initialize(argv)
        @argv = argv.dup
        @recipe_dir = CompatibilityPermissionReviewPlan::DEFAULT_RECIPE_DIR
      end

      def run
        app_id = nil

        OptionParser.new do |options|
          options.banner = "Usage: xnix-compat-permission-review-plan --app APP_ID [--recipe-dir PATH]"
          options.on("--app APP_ID", "Application id") { |value| app_id = value }
          options.on("--recipe-dir PATH", "Recipe directory") { |value| @recipe_dir = value }
        end.parse!(@argv)

        raise ArgumentError, "--app is required" unless app_id

        store = RecipeStore.new(path: @recipe_dir)
        recipe = store.find(app_id)
        raise ArgumentError, "unknown application: #{app_id}" unless recipe

        puts CompatibilityPermissionReviewPlan.new(recipe: recipe).to_json
      rescue ArgumentError => e
        warn "xnix-compat-permission-review-plan: #{e.message}"
        64
      else
        0
      end
    end
  end
end

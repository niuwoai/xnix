# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"
require_relative "recipe_store"

module Xnix
  module Compatibility
    class DesktopResourceBridgePlan
      PROJECT_ROOT = Pathname.new(__dir__).join("../../..").realpath
      VERSION = PROJECT_ROOT.join("VERSION").read.strip
      DEFAULT_RECIPE_DIR = PROJECT_ROOT.join("runtime/recipes").to_s

      RESOURCES = [
        ["file-open", "File Open", "file-open", "org.freedesktop.portal.FileChooser"],
        ["uri-open", "URI Open", "uri-open", "org.freedesktop.portal.OpenURI"],
        ["print", "Print", "print", "org.freedesktop.portal.Print"],
        ["clipboard", "Clipboard", "clipboard", "org.freedesktop.portal.Clipboard"],
        ["screenshot", "Screenshot", "screenshot", "org.freedesktop.portal.Screenshot"]
      ].freeze

      REQUIRED_RUNTIME_GATES = %w[
        portal-policy-review
        portal-request-plan
        snapshot-baseline
        backend-environment-plan
        runtime-write-gate
      ].freeze

      def initialize(recipe:)
        @recipe = recipe
      end

      def to_h
        {
          "version" => VERSION,
          "plan_type" => "desktop-resource-bridge-plan",
          "runtime_method" => "GetDesktopResourceBridgePlan",
          "application" => application,
          "bridge_state" => "planned",
          "resources" => resources,
          "resource_count" => RESOURCES.length,
          "required_runtime_gates" => REQUIRED_RUNTIME_GATES,
          "runtime_owned" => true,
          "c_runtime_backed" => true,
          "kde_policy_owner" => false,
          "portal_mediated" => true,
          "file_bridge_planned" => true,
          "uri_bridge_planned" => true,
          "print_bridge_planned" => true,
          "clipboard_bridge_planned" => true,
          "screenshot_bridge_planned" => true,
          "bridges_enabled" => false,
          "requests_created" => false,
          "backend_process_started" => false,
          "direct_host_file_access" => false,
          "direct_clipboard_access" => false,
          "direct_print_access" => false,
          "host_root_modified" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => "Desktop resource bridges are planned through Runtime-owned XDG Desktop Portal requests and remain disabled until review gates pass."
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

      def resources
        RESOURCES.map do |id, name, operation, portal_interface|
          {
            "id" => id,
            "name" => name,
            "operation" => operation,
            "portal_interface" => portal_interface,
            "runtime_method" => "GetPortalRequestPlan",
            "state" => "planned",
            "portal_required" => true,
            "user_approval_required" => true,
            "bridge_enabled" => false,
            "request_created" => false,
            "direct_backend_access_allowed" => false,
            "backend_details_exposed" => false
          }
        end
      end
    end

    class DesktopResourceBridgePlanCLI
      def initialize(argv)
        @argv = argv
        @recipe_dir = DesktopResourceBridgePlan::DEFAULT_RECIPE_DIR
      end

      def run
        app_id = nil

        OptionParser.new do |options|
          options.banner = "Usage: xnix-desktop-resource-bridge-plan --app APP_ID [--recipe-dir PATH]"
          options.on("--app APP_ID", "Application id") { |value| app_id = value }
          options.on("--recipe-dir PATH", "Recipe directory") { |value| @recipe_dir = value }
        end.parse!(@argv)

        raise ArgumentError, "--app is required" unless app_id

        store = RecipeStore.new(path: @recipe_dir)
        recipe = store.find(app_id)
        raise ArgumentError, "unknown application: #{app_id}" unless recipe

        puts DesktopResourceBridgePlan.new(recipe: recipe).to_json
      rescue ArgumentError => e
        warn "xnix-desktop-resource-bridge-plan: #{e.message}"
        64
      else
        0
      end
    end
  end
end

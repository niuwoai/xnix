# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"
require_relative "recipe_store"

module Xnix
  module Compatibility
    class CompatibilityModeSwitchPlan
      PROJECT_ROOT = Pathname.new(__dir__).join("../../..").realpath
      VERSION = PROJECT_ROOT.join("VERSION").read.strip
      DEFAULT_RECIPE_DIR = PROJECT_ROOT.join("runtime/recipes").to_s

      MODES = [
        {
          "id" => "automatic",
          "label" => "Automatic",
          "intent" => "Let Xnix choose the safest available compatibility path."
        },
        {
          "id" => "prefer-performance",
          "label" => "Prefer performance",
          "intent" => "Prefer a low-overhead compatibility path when policy gates allow it."
        },
        {
          "id" => "prefer-compatibility",
          "label" => "Prefer compatibility",
          "intent" => "Prefer a more conservative compatibility path for difficult applications."
        },
        {
          "id" => "isolated-execution",
          "label" => "Isolated execution",
          "intent" => "Prefer a stronger isolation boundary for higher-risk applications."
        }
      ].freeze

      REQUIRED_RUNTIME_GATES = %w[
        settings-review
        portal-policy-review
        snapshot-baseline
        backend-environment-plan
        runtime-write-gate
      ].freeze

      def initialize(recipe:, requested_mode:)
        @recipe = recipe
        @requested_mode = requested_mode
        validate!
      end

      def to_h
        {
          "version" => VERSION,
          "plan_type" => "compatibility-mode-switch-plan",
          "runtime_method" => "GetCompatibilityModeSwitchPlan",
          "application" => application,
          "current_mode" => recipe.mode,
          "requested_mode" => requested_mode,
          "mode_state" => "planned",
          "modes" => modes,
          "mode_count" => MODES.length,
          "required_runtime_gates" => REQUIRED_RUNTIME_GATES,
          "runtime_owned" => true,
          "c_runtime_backed" => true,
          "kde_policy_owner" => false,
          "valid_mode" => true,
          "requires_user_confirmation" => true,
          "portal_review_required" => true,
          "snapshot_required" => true,
          "settings_persistence_enabled" => false,
          "backend_reconfiguration_enabled" => false,
          "backend_process_started" => false,
          "launch_enabled" => false,
          "host_root_modified" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => "Compatibility mode switch is planned for user review and cannot change Runtime state in this version."
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      attr_reader :recipe, :requested_mode

      def validate!
        return if MODES.any? { |mode| mode.fetch("id") == requested_mode }

        raise ArgumentError, "unsupported compatibility mode: #{requested_mode}"
      end

      def application
        {
          "id" => recipe.id,
          "name" => recipe.name,
          "icon" => recipe.icon
        }
      end

      def modes
        MODES.map do |mode|
          mode.merge(
            "selected" => mode.fetch("id") == recipe.mode,
            "requested" => mode.fetch("id") == requested_mode,
            "backend_details_exposed" => false
          )
        end
      end
    end

    class CompatibilityModeSwitchPlanCLI
      def initialize(argv)
        @argv = argv.dup
        @recipe_dir = CompatibilityModeSwitchPlan::DEFAULT_RECIPE_DIR
      end

      def run
        app_id = nil
        requested_mode = nil

        OptionParser.new do |options|
          options.banner = "Usage: xnix-compat-mode-switch-plan --app APP_ID --mode MODE [--recipe-dir PATH]"
          options.on("--app APP_ID", "Application id") { |value| app_id = value }
          options.on("--mode MODE", "Requested compatibility mode") { |value| requested_mode = value }
          options.on("--recipe-dir PATH", "Recipe directory") { |value| @recipe_dir = value }
        end.parse!(@argv)

        raise ArgumentError, "--app is required" unless app_id
        raise ArgumentError, "--mode is required" unless requested_mode

        store = RecipeStore.new(path: @recipe_dir)
        recipe = store.find(app_id)
        raise ArgumentError, "unknown application: #{app_id}" unless recipe

        puts CompatibilityModeSwitchPlan.new(recipe: recipe, requested_mode: requested_mode).to_json
      rescue ArgumentError => e
        warn "xnix-compat-mode-switch-plan: #{e.message}"
        64
      else
        0
      end
    end
  end
end

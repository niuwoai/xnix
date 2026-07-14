# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"
require_relative "portal_request_model"
require_relative "recipe_store"
require_relative "settings_change_plan"

module Xnix
  module Compatibility
    class CompatibilityReviewFlowPlan
      PROJECT_ROOT = Pathname.new(__dir__).join("../../..").realpath
      VERSION = PROJECT_ROOT.join("VERSION").read.strip
      DEFAULT_RECIPE_DIR = PROJECT_ROOT.join("runtime/recipes").to_s
      DEFAULT_SECTION_ID = "resource-access"
      DEFAULT_FIELD_ID = "documents"
      DEFAULT_VALUE = "ask"
      DEFAULT_OPERATION = "file-open"

      FLOW_STEPS = [
        [
          "settings-change-review",
          "Review settings change",
          "GetCompatibilitySettingsChangePlan",
          "required",
          "KDE presents the requested user-facing settings change before Runtime persistence."
        ],
        [
          "permission-review",
          "Review permissions",
          "GetCompatibilityPermissionReviewPlan",
          "required",
          "KDE shows the grouped Runtime permission review before any permission grant."
        ],
        [
          "portal-request-review",
          "Review Portal request",
          "GetPortalRequestPlan",
          "required",
          "Runtime maps sensitive desktop access to an XDG Desktop Portal request plan."
        ],
        [
          "runtime-write-gate",
          "Check Runtime write gate",
          "GetRuntimeWriteGate",
          "blocked",
          "Runtime write gates still block persistence, launch, snapshots, and restore."
        ],
        [
          "review-receipt",
          "Record review intent",
          "GetCompatibilityActionReviewReceipt",
          "pending",
          "Review intent can be recorded, but it cannot approve execution by itself."
        ]
      ].freeze

      def initialize(recipe:, section_id: DEFAULT_SECTION_ID, field_id: DEFAULT_FIELD_ID, value: DEFAULT_VALUE, operation: DEFAULT_OPERATION)
        @recipe = recipe
        @section_id = section_id
        @field_id = field_id
        @value = value
        @operation = operation
        @settings_change_plan = SettingsChangePlan.new(
          application_id: recipe.id,
          section_id: section_id,
          field_id: field_id,
          value: value
        ).to_h
        @portal_request_plan = PortalRequestModel.new(application_id: recipe.id, operation: operation).to_h
      end

      def to_h
        flow_steps = steps

        {
          "version" => VERSION,
          "plan_type" => "compatibility-review-flow-plan",
          "runtime_method" => "GetCompatibilityReviewFlowPlan",
          "application" => application,
          "review_state" => "planned",
          "section_id" => section_id,
          "field_id" => field_id,
          "requested_value" => value,
          "operation" => operation,
          "steps" => flow_steps,
          "step_count" => flow_steps.length,
          "required_review_count" => flow_steps.count { |step| step.fetch("status") == "required" },
          "blocked_step_count" => flow_steps.count { |step| step.fetch("status") == "blocked" },
          "pending_step_count" => flow_steps.count { |step| step.fetch("status") == "pending" },
          "settings_change_plan" => settings_change_summary,
          "portal_request_plan" => portal_request_summary,
          "runtime_owned" => true,
          "c_runtime_backed" => true,
          "kde_policy_owner" => false,
          "user_confirmation_required" => true,
          "portal_policy_review_required" => portal_request_plan.fetch("safety").fetch("portal_required"),
          "settings_change_planned" => true,
          "permission_review_planned" => true,
          "portal_request_planned" => true,
          "review_receipt_required" => true,
          "apply_enabled" => false,
          "request_object_created" => false,
          "permission_granted" => false,
          "settings_persisted" => false,
          "execution_started" => false,
          "host_root_modified" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => "KDE can display the full Runtime-owned review flow, but no permission, setting, or execution change is applied."
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      attr_reader :recipe, :section_id, :field_id, :value, :operation, :settings_change_plan, :portal_request_plan

      def application
        {
          "id" => recipe.id,
          "name" => recipe.name,
          "icon" => recipe.icon
        }
      end

      def steps
        FLOW_STEPS.map do |id, label, runtime_method, status, summary|
          {
            "id" => id,
            "label" => label,
            "runtime_method" => runtime_method,
            "status" => status,
            "user_visible" => true,
            "user_action_required" => %w[required pending].include?(status),
            "runtime_gate_required" => true,
            "blocks_apply" => status != "pass",
            "summary" => summary
          }
        end
      end

      def settings_change_summary
        {
          "plan_type" => settings_change_plan.fetch("plan_type"),
          "change_state" => settings_change_plan.fetch("change_state"),
          "section_id" => settings_change_plan.fetch("section_id"),
          "field_id" => settings_change_plan.fetch("field_id"),
          "requested_value" => settings_change_plan.fetch("requested_value"),
          "apply_enabled" => settings_change_plan.fetch("apply_enabled"),
          "settings_persisted" => settings_change_plan.fetch("settings_persisted")
        }
      end

      def portal_request_summary
        {
          "request_type" => portal_request_plan.fetch("request_type"),
          "operation" => portal_request_plan.fetch("operation"),
          "decision" => portal_request_plan.fetch("decision"),
          "request_allowed" => portal_request_plan.fetch("request_allowed"),
          "portal_required" => portal_request_plan.fetch("safety").fetch("portal_required"),
          "host_permission_changed" => portal_request_plan.fetch("safety").fetch("host_permission_changed")
        }
      end
    end

    class CompatibilityReviewFlowPlanCLI
      def initialize(argv)
        @argv = argv.dup
        @recipe_dir = CompatibilityReviewFlowPlan::DEFAULT_RECIPE_DIR
        @section_id = CompatibilityReviewFlowPlan::DEFAULT_SECTION_ID
        @field_id = CompatibilityReviewFlowPlan::DEFAULT_FIELD_ID
        @value = CompatibilityReviewFlowPlan::DEFAULT_VALUE
        @operation = CompatibilityReviewFlowPlan::DEFAULT_OPERATION
      end

      def run
        app_id = nil

        OptionParser.new do |options|
          options.banner = "Usage: xnix-compat-review-flow-plan --app APP_ID [--section SECTION] [--field FIELD] [--value VALUE] [--operation OPERATION] [--recipe-dir PATH]"
          options.on("--app APP_ID", "Application id") { |value| app_id = value }
          options.on("--section SECTION", "Settings section identifier") { |value| @section_id = value }
          options.on("--field FIELD", "Settings field identifier") { |value| @field_id = value }
          options.on("--value VALUE", "Requested settings value") { |value| @value = value }
          options.on("--operation OPERATION", PortalAccessPolicy::OPERATIONS.keys, "Sensitive desktop operation") { |value| @operation = value }
          options.on("--recipe-dir PATH", "Recipe directory") { |value| @recipe_dir = value }
        end.parse!(@argv)

        raise ArgumentError, "--app is required" unless app_id

        store = RecipeStore.new(path: @recipe_dir)
        recipe = store.find(app_id)
        raise ArgumentError, "unknown application: #{app_id}" unless recipe

        puts CompatibilityReviewFlowPlan.new(
          recipe: recipe,
          section_id: @section_id,
          field_id: @field_id,
          value: @value,
          operation: @operation
        ).to_json
      rescue OptionParser::ParseError, ArgumentError => e
        warn "xnix-compat-review-flow-plan: #{e.message}"
        64
      else
        0
      end
    end
  end
end

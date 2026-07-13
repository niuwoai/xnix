# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"
require_relative "ai_repair_approval_gate"
require_relative "compatibility_install_plan"
require_relative "portal_access_policy"
require_relative "recipe_store"
require_relative "runtime_service_binding"
require_relative "settings_change_plan"

module Xnix
  module Compatibility
    class CompatibilityActionQueue
      PROJECT_ROOT = Pathname.new(__dir__).join("../../..").realpath
      VERSION = PROJECT_ROOT.join("VERSION").read.strip
      DEFAULT_RECIPE_DIR = PROJECT_ROOT.join("runtime/recipes").to_s

      def initialize(recipe:)
        @recipe = recipe
      end

      def to_h
        queue_actions = actions

        {
          "version" => VERSION,
          "queue_type" => "compatibility-center-action-queue",
          "application" => application,
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "surface" => "Compatibility Center",
          "action_count" => queue_actions.length,
          "pending_action_count" => queue_actions.count { |action| action.fetch("status") != "pass" },
          "user_review_required_count" => queue_actions.count { |action| action.fetch("user_review_required") },
          "execution_enabled" => false,
          "repair_execution_enabled" => false,
          "settings_persistence_enabled" => false,
          "host_root_modified" => false,
          "network_required" => false,
          "backend_details_exposed" => false,
          "actions" => queue_actions,
          "blocked_actions" => blocked_actions,
          "desktop_safe_summary" => desktop_safe_summary
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
          "requested_mode" => recipe.mode
        }
      end

      def actions
        [
          install_readiness_action,
          settings_change_action,
          ai_repair_review_action,
          runtime_service_action,
          portal_policy_action
        ]
      end

      def install_readiness_action
        plan = install_plan
        action(
          id: "review-install-readiness",
          source_type: plan.fetch("plan_type"),
          status: plan.fetch("install_ready") ? "pass" : "blocked",
          priority: "high",
          title: "Review install readiness",
          summary: plan.fetch("desktop_safe_summary"),
          user_review_required: false,
          runtime_gate: "install-plan-readiness",
          next_step: "Wait for signed artifacts, package source readiness, state allocation, and recipe install gate approval."
        )
      end

      def settings_change_action
        plan = settings_change_plan
        action(
          id: "review-settings-change",
          source_type: plan.fetch("plan_type"),
          status: plan.fetch("apply_enabled") ? "pass" : "review-required",
          priority: "medium",
          title: "Review compatibility settings change",
          summary: plan.fetch("desktop_safe_summary"),
          user_review_required: plan.fetch("user_confirmation_required"),
          runtime_gate: "runtime-settings-persistence",
          next_step: "Keep the change pending until Runtime persistence support and Portal policy review are available."
        )
      end

      def ai_repair_review_action
        gate = ai_repair_gate
        action(
          id: "review-ai-repair",
          source_type: gate.fetch("gate_type"),
          status: gate.fetch("gate_decision") == "blocked-until-approval" ? "approval-required" : "review-only",
          priority: "high",
          title: "Review AI repair recommendation",
          summary: gate.fetch("desktop_safe_summary"),
          user_review_required: gate.fetch("gate_decision") == "blocked-until-approval",
          runtime_gate: "ai-repair-approval",
          next_step: "Require Compatibility Center review, Runtime approval, and restore-point preflight before repair execution."
        )
      end

      def runtime_service_action
        binding = runtime_service_binding
        status = binding.fetch("live_dbus_owner_ready") ? "pass" : "pending"
        action(
          id: "verify-runtime-service",
          source_type: binding.fetch("binding_type"),
          status: status,
          priority: "medium",
          title: "Verify Runtime service ownership",
          summary: binding.fetch("desktop_safe_summary"),
          user_review_required: false,
          runtime_gate: "live-dbus-owner",
          next_step: "Keep using the session-bus smoke adapter until production D-Bus ownership is enabled."
        )
      end

      def portal_policy_action
        policy = portal_policy
        action(
          id: "review-portal-policy",
          source_type: policy.fetch("policy_type"),
          status: policy.fetch("portal_required") ? "review-required" : "pass",
          priority: "medium",
          title: "Review desktop resource policy",
          summary: policy.fetch("desktop_safe_summary"),
          user_review_required: policy.fetch("user_mediation_required"),
          runtime_gate: "portal-policy-review",
          next_step: "Use XDG Desktop Portal approval before granting file access to compatibility applications."
        )
      end

      def action(id:, source_type:, status:, priority:, title:, summary:, user_review_required:, runtime_gate:, next_step:)
        {
          "id" => id,
          "source_type" => source_type,
          "status" => status,
          "priority" => priority,
          "title" => title,
          "summary" => summary,
          "user_review_required" => user_review_required,
          "runtime_gate" => runtime_gate,
          "next_step" => next_step,
          "execution_enabled" => false,
          "backend_details_exposed" => false
        }
      end

      def blocked_actions
        [
          "execute queued actions from KDE without Runtime approval",
          "persist settings while action queue execution is disabled",
          "start compatibility backends from the Compatibility Center",
          "grant desktop resources without XDG Desktop Portal review",
          "mutate the host root from Compatibility Center actions",
          "expose backend implementation details in action cards"
        ]
      end

      def desktop_safe_summary
        "Compatibility Center actions are queued for review and cannot execute until Runtime gates are implemented."
      end

      def install_plan
        @install_plan ||= CompatibilityInstallPlan.new(recipe: recipe).to_h
      end

      def settings_change_plan
        @settings_change_plan ||= SettingsChangePlan.new(
          application_id: recipe.id,
          section_id: "resource-access",
          field_id: "documents",
          value: "ask"
        ).to_h
      end

      def ai_repair_gate
        @ai_repair_gate ||= AIRepairApprovalGate.new(recipe: recipe).to_h
      end

      def runtime_service_binding
        @runtime_service_binding ||= RuntimeServiceBinding.new.to_h
      end

      def portal_policy
        @portal_policy ||= PortalAccessPolicy.new(application_id: recipe.id, operation: "file-open").to_h
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @recipe_dir = DEFAULT_RECIPE_DIR
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--app is required" unless @application_id

          recipe = RecipeStore.new(path: @recipe_dir).find(@application_id)
          raise ArgumentError, "unknown application: #{@application_id}" unless recipe

          puts CompatibilityActionQueue.new(recipe: recipe).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-action-queue: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-action-queue --app APP_ID [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build a Runtime-owned Compatibility Center action queue for APP_ID") do |value|
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

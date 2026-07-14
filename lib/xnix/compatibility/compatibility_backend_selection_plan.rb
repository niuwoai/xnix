# frozen_string_literal: true

require "json"
require "optparse"
require_relative "compatibility_backend_binding"
require_relative "compatibility_backend_capability_matrix"
require_relative "compatibility_run_plan"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class CompatibilityBackendSelectionPlan
      def initialize(recipe:)
        @recipe = recipe
      end

      def to_h
        candidate_rows = candidates

        {
          "version" => RuntimeDaemon::VERSION,
          "plan_type" => "compatibility-backend-selection-plan",
          "runtime_method" => "GetBackendSelectionPlan",
          "application" => application,
          "selected_strategy" => run_plan.fetch("execution").fetch("strategy"),
          "recommended_profile_id" => recommended_profile_id,
          "candidate_profiles" => candidate_rows,
          "candidate_count" => candidate_rows.length,
          "ready_candidate_count" => candidate_rows.count { |candidate| candidate.fetch("ready") },
          "blocked_candidate_count" => candidate_rows.count { |candidate| candidate.fetch("blocked") },
          "required_reviews" => required_reviews,
          "blocked_actions" => blocked_actions,
          "runtime_owned" => true,
          "c_runtime_backed" => true,
          "kde_policy_owner" => false,
          "selection_committed" => false,
          "selection_change_enabled" => false,
          "backend_launch_enabled" => false,
          "capability_activation_enabled" => false,
          "environment_created" => false,
          "request_object_created" => false,
          "state_root_created" => false,
          "snapshot_created" => false,
          "host_root_modified" => false,
          "privileged_container_required" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => "Runtime can explain the recommended compatibility profile, but no backend selection is committed yet."
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
          "requested_mode" => recipe.mode,
          "supported_extensions" => recipe.supported_extensions
        }
      end

      def run_plan
        @run_plan ||= CompatibilityRunPlan.new(recipe: recipe).to_h
      end

      def binding
        @binding ||= CompatibilityBackendBinding.new(recipe: recipe).to_h
      end

      def capability_matrix
        @capability_matrix ||= CompatibilityBackendCapabilityMatrix.new.to_h
      end

      def recommended_profile_id
        case run_plan.fetch("execution").fetch("strategy")
        when "isolated-compatibility-engine"
          "isolated-compatibility"
        else
          "local-compatibility"
        end
      end

      def candidates
        capability_matrix.fetch("profiles").map do |profile|
          selected = profile.fetch("id") == recommended_profile_id

          {
            "id" => profile.fetch("id"),
            "label" => profile.fetch("label"),
            "kind" => profile.fetch("kind"),
            "selection_state" => selected ? "recommended" : "available-after-review",
            "recommended" => selected,
            "ready" => false,
            "blocked" => true,
            "ready_capability_count" => profile.fetch("ready_count"),
            "pending_capability_count" => profile.fetch("pending_count"),
            "required_preflight" => binding.fetch("required_preflight").map { |step| step.fetch("id") },
            "selection_committed" => false,
            "environment_created" => false,
            "backend_process_started" => false,
            "backend_details_exposed" => false,
            "summary" => "#{profile.fetch("label")} is #{selected ? "recommended" : "available"} only after Runtime preflight passes."
          }
        end
      end

      def required_reviews
        [
          "backend-capability-review",
          "backend-binding-review",
          "application-state-root-review",
          "portal-policy-review",
          "snapshot-baseline-review"
        ]
      end

      def blocked_actions
        [
          "commit backend selection from KDE",
          "start compatibility backend after selecting a profile",
          "create profile environment before Runtime preflight",
          "expose backend implementation details to the desktop shell"
        ]
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

          puts CompatibilityBackendSelectionPlan.new(recipe: recipe).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-backend-selection-plan: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-backend-selection-plan --app APP_ID [--recipe-dir PATH]"
            options.on("--app APP_ID", "Build a Runtime-owned backend selection plan for APP_ID") do |value|
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

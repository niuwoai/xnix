# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"
require_relative "compatibility_action_queue"
require_relative "recipe_store"

module Xnix
  module Compatibility
    class ActionReviewReceipt
      PROJECT_ROOT = Pathname.new(__dir__).join("../../..").realpath
      VERSION = PROJECT_ROOT.join("VERSION").read.strip
      DEFAULT_RECIPE_DIR = PROJECT_ROOT.join("runtime/recipes").to_s
      DECISIONS = %w[reviewed approved deferred rejected].freeze

      def initialize(recipe:, action_id:, decision:)
        @recipe = recipe
        @action_id = action_id
        @decision = decision
        validate!
      end

      def to_h
        {
          "version" => VERSION,
          "receipt_type" => "compatibility-center-action-review-receipt",
          "receipt_id" => receipt_id,
          "queue_type" => queue.fetch("queue_type"),
          "application" => application,
          "action" => action_summary,
          "decision" => decision,
          "decision_recorded" => true,
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "surface" => "Compatibility Center",
          "execution_enabled" => false,
          "repair_execution_enabled" => false,
          "settings_persistence_enabled" => false,
          "resource_grant_created" => false,
          "host_root_modified" => false,
          "network_required" => false,
          "backend_details_exposed" => false,
          "required_runtime_gate" => action.fetch("runtime_gate"),
          "next_step" => next_step,
          "blocked_actions" => blocked_actions,
          "desktop_safe_summary" => desktop_safe_summary
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      attr_reader :recipe, :action_id, :decision

      def validate!
        raise ArgumentError, "decision must be one of: #{DECISIONS.join(", ")}" unless DECISIONS.include?(decision)
        return if action

        raise ArgumentError, "unknown action: #{action_id}"
      end

      def queue
        @queue ||= CompatibilityActionQueue.new(recipe: recipe).to_h
      end

      def action
        @action ||= queue.fetch("actions").find { |item| item.fetch("id") == action_id }
      end

      def application
        {
          "id" => recipe.id,
          "name" => recipe.name,
          "icon" => recipe.icon
        }
      end

      def action_summary
        {
          "id" => action.fetch("id"),
          "source_type" => action.fetch("source_type"),
          "status" => action.fetch("status"),
          "title" => action.fetch("title"),
          "user_review_required" => action.fetch("user_review_required"),
          "execution_enabled" => action.fetch("execution_enabled"),
          "backend_details_exposed" => action.fetch("backend_details_exposed")
        }
      end

      def receipt_id
        ["compat-review", recipe.id, action_id, decision].join("-").gsub(/[^a-zA-Z0-9.-]+/, "-")
      end

      def next_step
        return "Runtime records the rejection and keeps the queued action non-executing." if decision == "rejected"
        return "Runtime keeps the queued action deferred until the user reviews it again." if decision == "deferred"

        "Runtime records the review intent, then waits for #{action.fetch("runtime_gate")} before any execution."
      end

      def blocked_actions
        [
          "execute queued action from a review receipt",
          "treat KDE review intent as Runtime execution approval",
          "persist compatibility settings from a review receipt",
          "grant desktop resources from a review receipt",
          "mutate the host root from a review receipt",
          "expose backend implementation details in review receipts"
        ]
      end

      def desktop_safe_summary
        "Compatibility Center review intent is recorded, but Runtime execution gates still block action execution."
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @action_id = nil
          @decision = nil
          @recipe_dir = DEFAULT_RECIPE_DIR
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--app is required" unless @application_id
          raise ArgumentError, "--action is required" unless @action_id
          raise ArgumentError, "--decision is required" unless @decision

          recipe = RecipeStore.new(path: @recipe_dir).find(@application_id)
          raise ArgumentError, "unknown application: #{@application_id}" unless recipe

          puts ActionReviewReceipt.new(recipe: recipe, action_id: @action_id, decision: @decision).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-action-review: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-action-review --app APP_ID --action ACTION_ID --decision DECISION [--recipe-dir PATH]"
            options.on("--app APP_ID", "Record Compatibility Center review intent for APP_ID") do |value|
              @application_id = value
            end
            options.on("--action ACTION_ID", "Queued Compatibility Center action id") do |value|
              @action_id = value
            end
            options.on("--decision DECISION", "Review decision: #{DECISIONS.join(", ")}") do |value|
              @decision = value
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

# frozen_string_literal: true

require "json"
require "optparse"
require_relative "application_recipe"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class CompatibilityRepairPlan
      ISSUES = {
        "engine-binding-pending" => {
          "severity" => "warning",
          "automatic_allowed" => false,
          "user_approval_required" => true,
          "summary" => "Compatibility engine setup is not ready yet.",
          "actions" => [
            {
              "id" => "open-diagnostics",
              "kind" => "user-visible",
              "label" => "Open diagnostics"
            },
            {
              "id" => "prepare-engine-binding",
              "kind" => "runtime-task",
              "label" => "Prepare compatibility engine binding"
            }
          ]
        },
        "portal-approval-required" => {
          "severity" => "info",
          "automatic_allowed" => false,
          "user_approval_required" => true,
          "summary" => "A desktop permission request needs user approval.",
          "actions" => [
            {
              "id" => "request-portal-grant",
              "kind" => "portal-request",
              "label" => "Request desktop permission"
            }
          ]
        },
        "recipe-trust-blocked" => {
          "severity" => "critical",
          "automatic_allowed" => false,
          "user_approval_required" => true,
          "summary" => "Recipe trust requirements are not satisfied.",
          "actions" => [
            {
              "id" => "review-recipe-source",
              "kind" => "user-visible",
              "label" => "Review recipe source"
            }
          ]
        },
        "runtime-repair-applied" => {
          "severity" => "info",
          "automatic_allowed" => true,
          "user_approval_required" => false,
          "summary" => "A safe compatibility repair was applied.",
          "actions" => [
            {
              "id" => "show-repair-record",
              "kind" => "user-visible",
              "label" => "Show repair record"
            }
          ]
        }
      }.freeze

      attr_reader :application_id, :issue

      def initialize(application_id:, issue:)
        @application_id = application_id
        @issue = issue
        validate!
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "plan_type" => "compatibility-repair",
          "application_id" => application_id,
          "issue" => issue,
          "severity" => rule.fetch("severity"),
          "automatic_allowed" => rule.fetch("automatic_allowed"),
          "user_approval_required" => rule.fetch("user_approval_required"),
          "snapshot_required" => snapshot_required?,
          "rollback_available" => true,
          "actions" => rule.fetch("actions"),
          "notification_event" => notification_event,
          "desktop_safe_summary" => rule.fetch("summary")
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def validate!
        unless application_id.is_a?(String) && application_id.match?(ApplicationRecipe::ID_PATTERN)
          raise ArgumentError, "application id must be a reverse-DNS identifier"
        end

        return if ISSUES.key?(issue)

        raise ArgumentError, "issue must be one of: #{ISSUES.keys.join(", ")}"
      end

      def rule
        ISSUES.fetch(issue)
      end

      def snapshot_required?
        %w[engine-binding-pending runtime-repair-applied].include?(issue)
      end

      def notification_event
        return "repair-applied" if issue == "runtime-repair-applied"
        return "approval-required" if rule.fetch("user_approval_required")

        "install-failed"
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @issue = nil
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--app is required" unless @application_id
          raise ArgumentError, "--issue is required" unless @issue
          raise ArgumentError, "unexpected arguments: #{@argv.join(" ")}" unless @argv.empty?

          puts CompatibilityRepairPlan.new(application_id: @application_id, issue: @issue).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-repair-plan: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-repair-plan --app APP_ID --issue ISSUE"
            options.on("--app APP_ID", "Build a repair plan for APP_ID") do |value|
              @application_id = value
            end
            options.on("--issue ISSUE", ISSUES.keys, "Repair issue type") do |value|
              @issue = value
            end
          end
        end
      end
    end
  end
end

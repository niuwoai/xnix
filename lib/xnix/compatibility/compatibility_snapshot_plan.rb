# frozen_string_literal: true

require "json"
require "optparse"
require_relative "application_recipe"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class CompatibilitySnapshotPlan
      REASONS = {
        "before-repair" => "Create a restore point before a compatibility repair.",
        "before-engine-change" => "Create a restore point before changing the compatibility engine.",
        "manual" => "Create a user-requested restore point."
      }.freeze

      attr_reader :application_id, :reason

      def initialize(application_id:, reason:)
        @application_id = application_id
        @reason = reason
        validate!
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "plan_type" => "compatibility-snapshot",
          "application_id" => application_id,
          "reason" => reason,
          "enabled_by_default" => true,
          "snapshot_scope" => snapshot_scope,
          "restore" => restore,
          "retention" => retention,
          "desktop_safe_summary" => REASONS.fetch(reason)
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

        return if REASONS.key?(reason)

        raise ArgumentError, "reason must be one of: #{REASONS.keys.join(", ")}"
      end

      def snapshot_scope
        {
          "application_state" => true,
          "runtime_metadata" => true,
          "desktop_activation_receipts" => true,
          "user_documents" => false,
          "host_system" => false
        }
      end

      def restore
        {
          "available" => true,
          "method" => "Runtime.RestoreSnapshot",
          "requires_user_confirmation" => true,
          "preserve_user_documents" => true
        }
      end

      def retention
        {
          "policy" => "bounded",
          "keep_latest" => 5,
          "prune_automatically" => true
        }
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @reason = "manual"
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--app is required" unless @application_id
          raise ArgumentError, "unexpected arguments: #{@argv.join(" ")}" unless @argv.empty?

          puts CompatibilitySnapshotPlan.new(application_id: @application_id, reason: @reason).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-snapshot-plan: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-snapshot-plan --app APP_ID [--reason REASON]"
            options.on("--app APP_ID", "Build a snapshot plan for APP_ID") do |value|
              @application_id = value
            end
            options.on("--reason REASON", REASONS.keys, "Snapshot reason") do |value|
              @reason = value
            end
          end
        end
      end
    end
  end
end

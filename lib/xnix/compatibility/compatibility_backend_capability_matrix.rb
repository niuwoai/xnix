# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"

module Xnix
  module Compatibility
    class CompatibilityBackendCapabilityMatrix
      PROJECT_ROOT = Pathname.new(__dir__).join("../../..").realpath
      VERSION = PROJECT_ROOT.join("VERSION").read.strip

      CAPABILITIES = [
        ["application-launch", "Application launch", "pending", "Runtime launch binding is still gated."],
        ["package-management", "Managed packages", "pending", "Runtime package sources are planned before backend availability."],
        ["file-bridge", "File bridge", "pending", "File access must pass Portal policy and state-root preflight."],
        ["clipboard-bridge", "Clipboard bridge", "pending", "Clipboard access must pass Portal policy review."],
        ["print-bridge", "Print bridge", "pending", "Print access must pass Portal policy review."],
        ["snapshot-restore", "Snapshot and restore", "pending", "Restore points must exist before risky compatibility changes."],
        ["diagnostics", "Diagnostics", "ready", "Runtime can describe diagnostics without starting a backend."]
      ].freeze

      PROFILES = [
        ["local-compatibility", "Local compatibility profile", "local"],
        ["isolated-compatibility", "Isolated compatibility profile", "isolated"]
      ].freeze

      def to_h
        profile_rows = profiles

        {
          "version" => VERSION,
          "matrix_type" => "compatibility-backend-capability-matrix",
          "runtime_method" => "GetBackendCapabilityMatrix",
          "profiles" => profile_rows,
          "profile_count" => profile_rows.length,
          "capability_count" => CAPABILITIES.length,
          "ready_capability_count" => profile_rows.sum { |profile| profile.fetch("ready_count") },
          "pending_capability_count" => profile_rows.sum { |profile| profile.fetch("pending_count") },
          "blocked_capability_count" => profile_rows.sum { |profile| profile.fetch("blocked_count") },
          "runtime_owned" => true,
          "c_runtime_backed" => true,
          "kde_policy_owner" => false,
          "selection_enabled" => false,
          "backend_launch_enabled" => false,
          "capability_activation_enabled" => false,
          "request_objects_created" => false,
          "state_root_created" => false,
          "snapshots_created" => false,
          "host_root_modified" => false,
          "privileged_container_required" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => "Runtime backend capability planning is visible to KDE, while backend selection and launch remain disabled."
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def profiles
        PROFILES.map do |id, label, kind|
          capability_rows = capabilities

          {
            "id" => id,
            "label" => label,
            "kind" => kind,
            "selection_state" => "planned",
            "capabilities" => capability_rows,
            "capability_count" => capability_rows.length,
            "ready_count" => capability_rows.count { |capability| capability.fetch("status") == "ready" },
            "pending_count" => capability_rows.count { |capability| capability.fetch("status") == "pending" },
            "blocked_count" => capability_rows.count { |capability| capability.fetch("status") == "blocked" },
            "runtime_owned" => true,
            "kde_policy_owner" => false,
            "profile_ready" => false,
            "selection_enabled" => false,
            "backend_process_started" => false,
            "network_required_for_planning" => false,
            "host_root_modified" => false,
            "privileged_container_required" => false,
            "backend_details_exposed" => false,
            "summary" => "#{label} is planned but not selectable until Runtime preflight passes."
          }
        end
      end

      def capabilities
        CAPABILITIES.map do |id, label, status, summary|
          {
            "id" => id,
            "label" => label,
            "status" => status,
            "summary" => summary,
            "user_visible" => true,
            "requires_portal_review" => %w[file-bridge clipboard-bridge print-bridge].include?(id),
            "requires_state_root" => %w[application-launch file-bridge snapshot-restore].include?(id),
            "requires_snapshot" => id == "snapshot-restore",
            "activation_enabled" => false,
            "request_object_created" => false,
            "backend_process_started" => false,
            "backend_details_exposed" => false
          }
        end
      end
    end

    class CompatibilityBackendCapabilityMatrixCLI
      def initialize(argv)
        @argv = argv.dup
      end

      def run
        OptionParser.new do |options|
          options.banner = "Usage: xnix-compat-backend-capability-matrix"
        end.parse!(@argv)
        raise ArgumentError, "unexpected arguments: #{@argv.join(" ")}" unless @argv.empty?

        puts CompatibilityBackendCapabilityMatrix.new.to_json
      rescue OptionParser::ParseError, ArgumentError => e
        warn "xnix-compat-backend-capability-matrix: #{e.message}"
        64
      else
        0
      end
    end
  end
end

# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"

module Xnix
  module Compatibility
    class RuntimeWriteGate
      PROJECT_ROOT = Pathname.new(__dir__).join("../../..").realpath
      VERSION = PROJECT_ROOT.join("VERSION").read.strip
      WRITE_METHODS = %w[
        InstallRecipe
        Launch
        CreateSnapshot
        RestoreSnapshot
      ].freeze
      ERROR_NAME = "org.xnix.Compatibility1.Error.WriteMethodDisabled"

      attr_reader :method_name

      def initialize(method_name:)
        @method_name = method_name
        validate_method_name
      end

      def to_h
        {
          "version" => VERSION,
          "gate_type" => "runtime-write-gate",
          "method_name" => method_name,
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "gate_decision" => "blocked-until-production-backend",
          "write_method_enabled" => false,
          "dispatch_enabled" => false,
          "request_object_created" => false,
          "execution_started" => false,
          "supported_write_methods" => WRITE_METHODS,
          "required_gates" => required_gates,
          "denial_error_name" => ERROR_NAME,
          "network_required" => false,
          "host_root_modified" => false,
          "privileged_container_required" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => desktop_safe_summary
        }
      end

      def failure_message
        "#{ERROR_NAME}: #{method_name} is blocked until Runtime production backend gates pass"
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def validate_method_name
        return if WRITE_METHODS.include?(method_name)

        raise ArgumentError, "method must be one of: #{WRITE_METHODS.join(", ")}"
      end

      def required_gates
        [
          gate("production-runtime-owner", "pending", "The packaged Runtime service must own the stable D-Bus name."),
          gate("backend-binding-ready", "pending", "A managed compatibility backend must be selected and verified."),
          gate("recipe-trust-production", "pending", "The application recipe must satisfy production trust policy."),
          gate("user-action-review", "pending", "The Compatibility Center must record user approval for the operation."),
          gate("portal-approval-if-sensitive", "pending", "Sensitive desktop resources must be mediated through XDG Desktop Portal."),
          gate("snapshot-preflight-for-risky-change", "pending", "Risky state changes must have a Runtime restore point plan.")
        ]
      end

      def gate(id, status, summary)
        {
          "id" => id,
          "status" => status,
          "summary" => summary
        }
      end

      def desktop_safe_summary
        "#{method_name} is visible to desktop integrations but remains blocked until production Runtime gates pass."
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @method_name = nil
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--method is required" unless @method_name

          puts RuntimeWriteGate.new(method_name: @method_name).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-runtime-write-gate: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-runtime-write-gate --method METHOD"
            options.on("--method METHOD", WRITE_METHODS, "Evaluate a gated Runtime write method") do |value|
              @method_name = value
            end
          end
        end
      end
    end
  end
end

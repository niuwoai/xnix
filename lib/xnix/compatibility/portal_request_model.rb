# frozen_string_literal: true

require "json"
require "optparse"
require_relative "portal_access_policy"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class PortalRequestModel
      PORTAL_DESTINATION = "org.freedesktop.portal.Desktop"
      OPERATION_METHODS = {
        "file-open" => "OpenFile",
        "uri-open" => "OpenURI",
        "print" => "Print",
        "screenshot" => "Screenshot",
        "clipboard" => "RequestClipboard",
        "camera" => "AccessCamera",
        "remote-desktop" => "CreateSession"
      }.freeze

      def initialize(application_id:, operation:, reason: nil)
        @application_id = application_id
        @operation = operation
        @reason = reason
        @policy = PortalAccessPolicy.new(application_id: application_id, operation: operation)
      end

      def to_h
        policy = @policy.to_h

        {
          "version" => RuntimeDaemon::VERSION,
          "request_type" => "portal-request",
          "desktop" => "KDE Plasma",
          "application_id" => @application_id,
          "operation" => @operation,
          "reason" => reason,
          "decision" => policy.fetch("decision"),
          "request_allowed" => policy.fetch("decision") != "deny",
          "portal" => portal(policy),
          "request" => request(policy),
          "completion" => completion,
          "denied" => denied(policy),
          "safety" => safety(policy)
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def reason
        return @reason if @reason && !@reason.empty?

        "Compatibility application requested #{@operation} access."
      end

      def portal(policy)
        {
          "destination" => PORTAL_DESTINATION,
          "interface" => policy.fetch("portal_interface"),
          "method" => OPERATION_METHODS.fetch(@operation),
          "object_path" => "/org/freedesktop/portal/desktop",
          "dbus_api" => "XDG Desktop Portal"
        }
      end

      def request(policy)
        {
          "object_path_required" => policy.fetch("request_flow").fetch("request_object_required"),
          "handle_token" => handle_token,
          "user_mediation_required" => policy.fetch("user_mediation_required"),
          "resources" => policy.fetch("resources"),
          "runtime_policy_owner" => policy.fetch("request_flow").fetch("runtime_policy_owner"),
          "desktop_shell_policy_owner" => policy.fetch("request_flow").fetch("desktop_shell_policy_owner")
        }
      end

      def completion
        {
          "signal" => "Response",
          "response_field" => "response",
          "success_code" => 0,
          "cancelled_code" => 1,
          "denied_code" => 2,
          "result_owner" => "Runtime"
        }
      end

      def denied(policy)
        return nil unless policy.fetch("decision") == "deny"

        {
          "reason" => policy.fetch("desktop_safe_summary"),
          "next_action" => "open-compatibility-settings",
          "notification_event" => "approval-required"
        }
      end

      def safety(policy)
        {
          "direct_access_allowed" => policy.fetch("direct_access_allowed"),
          "portal_required" => policy.fetch("portal_required"),
          "backend_details_exposed" => false,
          "host_permission_changed" => false
        }
      end

      def handle_token
        "xnix_#{@application_id.gsub(/[^a-zA-Z0-9]/, "_")}_#{@operation.gsub(/[^a-zA-Z0-9]/, "_")}"
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @operation = nil
          @reason = nil
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--app is required" unless @application_id
          raise ArgumentError, "--operation is required" unless @operation
          raise ArgumentError, "unexpected arguments: #{@argv.join(" ")}" unless @argv.empty?

          puts PortalRequestModel.new(
            application_id: @application_id,
            operation: @operation,
            reason: @reason
          ).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-portal-request-model: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-portal-request-model --app APP_ID --operation OPERATION [--reason TEXT]"
            options.on("--app APP_ID", "Build a Portal request model for APP_ID") do |value|
              @application_id = value
            end
            options.on("--operation OPERATION", PortalAccessPolicy::OPERATIONS.keys, "Sensitive desktop operation") do |value|
              @operation = value
            end
            options.on("--reason TEXT", "User-facing request reason") do |value|
              @reason = value
            end
          end
        end
      end
    end
  end
end

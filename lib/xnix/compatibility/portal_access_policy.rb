# frozen_string_literal: true

require "json"
require "optparse"
require_relative "application_recipe"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class PortalAccessPolicy
      OPERATIONS = {
        "file-open" => {
          "portal_interface" => "org.freedesktop.portal.FileChooser",
          "decision" => "ask",
          "resources" => %w[documents downloads selected-files],
          "summary" => "File access requires a user-approved desktop portal request."
        },
        "uri-open" => {
          "portal_interface" => "org.freedesktop.portal.OpenURI",
          "decision" => "ask",
          "resources" => %w[external-uri],
          "summary" => "URI handling requires a user-approved desktop portal request."
        },
        "print" => {
          "portal_interface" => "org.freedesktop.portal.Print",
          "decision" => "ask",
          "resources" => %w[printer],
          "summary" => "Printing requires a user-approved desktop portal request."
        },
        "screenshot" => {
          "portal_interface" => "org.freedesktop.portal.Screenshot",
          "decision" => "ask",
          "resources" => %w[screen],
          "summary" => "Screenshots require a user-approved desktop portal request."
        },
        "clipboard" => {
          "portal_interface" => "org.freedesktop.portal.Clipboard",
          "decision" => "ask",
          "resources" => %w[clipboard],
          "summary" => "Clipboard access requires a user-approved desktop portal request."
        },
        "camera" => {
          "portal_interface" => "org.freedesktop.portal.Camera",
          "decision" => "deny",
          "resources" => %w[camera],
          "summary" => "Camera access is denied until the user changes the application policy."
        },
        "remote-desktop" => {
          "portal_interface" => "org.freedesktop.portal.RemoteDesktop",
          "decision" => "deny",
          "resources" => %w[screen input-devices],
          "summary" => "Remote desktop access is denied until the user changes the application policy."
        }
      }.freeze

      attr_reader :application_id, :operation

      def initialize(application_id:, operation:)
        @application_id = application_id
        @operation = operation
        validate!
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "policy_type" => "portal-access",
          "desktop" => "KDE Plasma",
          "application_id" => application_id,
          "operation" => operation,
          "decision" => rule.fetch("decision"),
          "portal_required" => true,
          "portal_interface" => rule.fetch("portal_interface"),
          "user_mediation_required" => true,
          "direct_access_allowed" => false,
          "resources" => rule.fetch("resources"),
          "request_flow" => request_flow,
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

        return if OPERATIONS.key?(operation)

        raise ArgumentError, "operation must be one of: #{OPERATIONS.keys.join(", ")}"
      end

      def rule
        OPERATIONS.fetch(operation)
      end

      def request_flow
        {
          "dbus_api" => "XDG Desktop Portal",
          "request_object_required" => true,
          "completion_signal" => "portal-response",
          "runtime_policy_owner" => true,
          "desktop_shell_policy_owner" => false
        }
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @operation = nil
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--app is required" unless @application_id
          raise ArgumentError, "--operation is required" unless @operation
          raise ArgumentError, "unexpected arguments: #{@argv.join(" ")}" unless @argv.empty?

          puts PortalAccessPolicy.new(application_id: @application_id, operation: @operation).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-portal-access-policy: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-portal-access-policy --app APP_ID --operation OPERATION"
            options.on("--app APP_ID", "Evaluate portal access policy for APP_ID") do |value|
              @application_id = value
            end
            options.on("--operation OPERATION", OPERATIONS.keys, "Evaluate a sensitive desktop operation") do |value|
              @operation = value
            end
          end
        end
      end
    end
  end
end

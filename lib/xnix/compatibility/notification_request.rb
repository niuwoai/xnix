# frozen_string_literal: true

require "json"
require "optparse"

module Xnix
  module Compatibility
    class NotificationRequest
      EVENT_TYPES = {
        "install-failed" => {
          "urgency" => "critical",
          "title" => "Installation needs attention",
          "default_body" => "A compatibility installation failed and needs review."
        },
        "repair-applied" => {
          "urgency" => "normal",
          "title" => "Compatibility repair applied",
          "default_body" => "Xnix applied an automatic compatibility repair."
        },
        "mode-changed" => {
          "urgency" => "low",
          "title" => "Compatibility mode changed",
          "default_body" => "The Runtime changed an application's compatibility mode."
        },
        "approval-required" => {
          "urgency" => "critical",
          "title" => "Approval required",
          "default_body" => "A compatibility action needs your approval."
        }
      }.freeze

      attr_reader :application_id, :event_type, :body

      def initialize(application_id:, event_type:, body: nil)
        @application_id = application_id
        @event_type = event_type
        @body = body
        validate!
      end

      def to_h
        event = EVENT_TYPES.fetch(event_type)

        {
          "request_type" => "desktop-notification",
          "source" => "runtime-event",
          "desktop" => "KDE Plasma",
          "application_id" => application_id,
          "event_type" => event_type,
          "urgency" => event.fetch("urgency"),
          "title" => event.fetch("title"),
          "body" => body || event.fetch("default_body"),
          "actions" => actions,
          "backend_required" => "kde-notification"
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def validate!
        unless application_id.is_a?(String) && application_id.match?(ApplicationId::PATTERN)
          raise ArgumentError, "application id must be a reverse-DNS identifier"
        end
        raise ArgumentError, "event type must be one of: #{EVENT_TYPES.keys.join(", ")}" unless EVENT_TYPES.key?(event_type)
        raise ArgumentError, "body must be a single-line string" if body && (!body.is_a?(String) || body.empty? || body.match?(/[\r\n]/))
      end

      def actions
        return ["open-compatibility-center", "review-request"] if event_type == "approval-required"
        return ["open-compatibility-center", "show-diagnostics"] if event_type == "install-failed"

        ["open-compatibility-center"]
      end

      module ApplicationId
        PATTERN = /\A[a-z][a-z0-9-]*(?:\.[a-z0-9-]+)+\z/
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @event_type = nil
          @body = nil
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--app is required" unless @application_id
          raise ArgumentError, "--event is required" unless @event_type

          puts NotificationRequest.new(
            application_id: @application_id,
            event_type: @event_type,
            body: @body
          ).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-notify: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-notify --app APP_ID --event EVENT_TYPE [--body TEXT]"
            options.on("--app APP_ID", "Notify for a Runtime application") do |value|
              @application_id = value
            end
            options.on("--event EVENT_TYPE", "Runtime event type") do |value|
              @event_type = value
            end
            options.on("--body TEXT", "Override notification body") do |value|
              @body = value
            end
          end
        end
      end
    end
  end
end

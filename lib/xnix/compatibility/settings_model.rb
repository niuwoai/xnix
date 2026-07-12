# frozen_string_literal: true

require "json"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class SettingsModel
      MODE_OPTIONS = %w[automatic performance compatibility].freeze
      ACCESS_OPTIONS = %w[allow ask deny].freeze
      DEVICE_OPTIONS = %w[allow ask deny].freeze
      NETWORK_OPTIONS = %w[allow ask deny].freeze

      def initialize(application_id:)
        @application_id = application_id
        validate_application_id!
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "request_type" => "settings-model",
          "desktop" => "KDE Plasma",
          "application_id" => @application_id,
          "sections" => sections
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def sections
        [
          run_mode_section,
          resource_access_section,
          device_section,
          network_section,
          snapshot_section
        ]
      end

      def run_mode_section
        {
          "id" => "run-mode",
          "title" => "Run mode",
          "description" => "Choose how Xnix balances speed and compatibility.",
          "fields" => [
            field("mode", "Run mode", "automatic", MODE_OPTIONS),
            field("preference", "Priority", "compatibility", %w[performance compatibility])
          ]
        }
      end

      def resource_access_section
        {
          "id" => "resource-access",
          "title" => "File access",
          "description" => "Control which user folders this application may request.",
          "fields" => [
            field("documents", "Documents", "ask", ACCESS_OPTIONS),
            field("downloads", "Downloads", "ask", ACCESS_OPTIONS)
          ]
        }
      end

      def device_section
        {
          "id" => "devices",
          "title" => "Devices",
          "description" => "Control sensitive device access.",
          "fields" => [
            field("camera", "Camera", "deny", DEVICE_OPTIONS)
          ]
        }
      end

      def network_section
        {
          "id" => "network",
          "title" => "Network",
          "description" => "Control network access for compatibility actions.",
          "fields" => [
            field("network", "Network", "allow", NETWORK_OPTIONS)
          ]
        }
      end

      def snapshot_section
        {
          "id" => "snapshots",
          "title" => "Snapshots",
          "description" => "Keep restore points before risky compatibility changes.",
          "fields" => [
            field("snapshots", "Environment snapshots", "enabled", %w[enabled disabled])
          ]
        }
      end

      def field(id, label, value, options)
        {
          "id" => id,
          "label" => label,
          "value" => value,
          "options" => options
        }
      end

      def validate_application_id!
        return if @application_id.is_a?(String) && @application_id.match?(/\A[a-z][a-z0-9-]*(?:\.[a-z0-9-]+)+\z/)

        raise ArgumentError, "application id must be a reverse-DNS identifier"
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
        end

        def run
          application_id = @argv.shift
          raise ArgumentError, "application id is required" unless application_id
          raise ArgumentError, "unexpected arguments: #{@argv.join(" ")}" unless @argv.empty?

          puts SettingsModel.new(application_id: application_id).to_json
          0
        rescue ArgumentError => e
          warn "xnix-compat-settings: #{e.message}"
          64
        end
      end
    end
  end
end

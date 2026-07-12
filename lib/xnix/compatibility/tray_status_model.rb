# frozen_string_literal: true

require "json"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class TrayStatusModel
      def initialize(active_count:, attention_count:, bridged_tray_count:)
        @active_count = integer_count(active_count, "active count")
        @attention_count = integer_count(attention_count, "attention count")
        @bridged_tray_count = integer_count(bridged_tray_count, "bridged tray count")
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "request_type" => "tray-status-model",
          "desktop" => "KDE Plasma",
          "runtime_activity" => runtime_activity,
          "compatibility_status" => compatibility_status,
          "tray_bridge" => tray_bridge,
          "actions" => actions
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def runtime_activity
        {
          "active_application_count" => @active_count,
          "attention_required_count" => @attention_count,
          "summary" => runtime_summary
        }
      end

      def compatibility_status
        {
          "state" => @attention_count.positive? ? "attention-required" : "ready",
          "label" => @attention_count.positive? ? "Attention required" : "Ready"
        }
      end

      def tray_bridge
        {
          "bridged_tray_application_count" => @bridged_tray_count,
          "state" => @bridged_tray_count.positive? ? "active" : "idle",
          "label" => @bridged_tray_count.positive? ? "Tray bridge active" : "No bridged tray applications"
        }
      end

      def actions
        ["open-compatibility-center", "open-settings"]
      end

      def runtime_summary
        return "No compatibility applications are active" if @active_count.zero?
        return "#{@active_count} compatibility application active" if @active_count == 1

        "#{@active_count} compatibility applications active"
      end

      def integer_count(value, label)
        integer = Integer(value)
        raise ArgumentError, "#{label} must not be negative" if integer.negative?

        integer
      rescue ArgumentError, TypeError
        raise ArgumentError, "#{label} must be a non-negative integer"
      end

      class CLI
        DEFAULTS = {
          "--active" => "0",
          "--attention" => "0",
          "--bridged-tray" => "0"
        }.freeze

        def initialize(argv)
          @argv = argv.dup
        end

        def run
          values = parse_values
          puts TrayStatusModel.new(
            active_count: values.fetch("--active"),
            attention_count: values.fetch("--attention"),
            bridged_tray_count: values.fetch("--bridged-tray")
          ).to_json
          0
        rescue ArgumentError => e
          warn "xnix-compat-tray-status: #{e.message}"
          64
        end

        private

        def parse_values
          values = DEFAULTS.dup
          until @argv.empty?
            key = @argv.shift
            raise ArgumentError, "unknown option: #{key}" unless values.key?(key)
            raise ArgumentError, "#{key} requires a value" if @argv.empty?

            values[key] = @argv.shift
          end
          values
        end
      end
    end
  end
end

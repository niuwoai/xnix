# frozen_string_literal: true

require "open3"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class DBusRuntimeClient
      class Error < StandardError; end

      attr_reader :capture

      def self.available?
        return false unless ENV["DBUS_SESSION_BUS_ADDRESS"]

        _stdout, _stderr, status = Open3.capture3(
          "gdbus", "wait",
          "--session",
          "--timeout", "1",
          RuntimeDaemon::BUS_NAME
        )
        status.success?
      rescue SystemCallError
        false
      end

      def initialize(capture: Open3.method(:capture3))
        @capture = capture
      end

      def source_metadata
        {
          "kind" => "runtime-dbus-session",
          "bus_name" => RuntimeDaemon::BUS_NAME,
          "object_path" => RuntimeDaemon::OBJECT_PATH,
          "interface" => RuntimeDaemon::INTERFACE
        }
      end

      def list_applications
        parse_dictionaries(call("ListApplications"))
      end

      def get_application(application_id)
        parse_dictionary(call("GetApplication", application_id))
      end

      def diagnostics(application_id)
        parse_dictionary(call("GetDiagnostics", application_id))
      end

      def engine_catalog
        parse_dictionary(call("GetEngineCatalog"))
      end

      def run_plan(application_id)
        parse_dictionary(call("GetRunPlan", application_id))
      end

      def repair_plan(application_id, issue)
        parse_dictionary(call("GetRepairPlan", application_id, issue))
      end

      def test_plan(application_id, test_type = "preflight")
        parse_dictionary(call("GetTestPlan", application_id, test_type))
      end

      def test_result(application_id, test_type = "preflight")
        parse_dictionary(call("GetTestResult", application_id, test_type))
      end

      def snapshot_plan(application_id, reason)
        parse_dictionary(call("GetSnapshotPlan", application_id, reason))
      end

      def portal_access_policy(application_id, operation)
        parse_dictionary(call("GetPortalAccessPolicy", application_id, operation))
      end

      private

      def call(method_name, *arguments)
        command = [
          "gdbus", "call",
          "--session",
          "--dest", RuntimeDaemon::BUS_NAME,
          "--object-path", RuntimeDaemon::OBJECT_PATH,
          "--method", "#{RuntimeDaemon::INTERFACE}.#{method_name}",
          *arguments
        ]
        stdout, stderr, status = capture.call(*command)
        raise Error, "D-Bus call #{method_name} failed: #{stderr}" unless status.success?

        stdout
      rescue SystemCallError => e
        raise Error, "D-Bus tooling is unavailable: #{e.message}"
      end

      def parse_dictionary(output)
        dictionaries = parse_dictionaries(output)
        raise Error, "D-Bus response did not contain a dictionary" if dictionaries.empty?

        dictionaries.first
      end

      def parse_dictionaries(output)
        output.scan(/\{([^{}]*)\}/).map do |match|
          parse_dictionary_body(match.first)
        end
      end

      def parse_dictionary_body(body)
        body.scan(/'([^']+)':\s*<([^>]*)>/).each_with_object({}) do |(key, raw_value), parsed|
          parsed[key] = parse_value(raw_value)
        end
      end

      def parse_value(raw_value)
        value = raw_value.strip
        string_match = value.match(/\A'(.*)'\z/m)
        return string_match[1] if string_match
        return true if value == "true"
        return false if value == "false"
        return value.scan(/'([^']*)'/).flatten if value.start_with?("[") && value.end_with?("]")

        value
      end
    end
  end
end

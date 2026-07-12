# frozen_string_literal: true

require "json"
require "optparse"
require_relative "dbus_runtime_client"
require_relative "recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class KdeCenterModel
      MODE_LABELS = {
        "automatic" => "Automatic",
        "wine" => "Compatibility engine",
        "vm" => "Isolated environment"
      }.freeze

      STATUS_LABELS = {
        "known" => "Known",
        "unknown" => "Unknown"
      }.freeze

      attr_reader :runtime

      def initialize(runtime:)
        @runtime = runtime
      end

      def to_h
        applications = runtime.list_applications.map { |application| application_summary(application) }

        {
          "version" => RuntimeDaemon::VERSION,
          "title" => "Xnix Compatibility Center",
          "source" => source_metadata,
          "summary" => summary(applications),
          "applications" => applications
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def source_metadata
        return runtime.source_metadata if runtime.respond_to?(:source_metadata)

        {
          "kind" => "runtime-local-read-model",
          "bus_name" => RuntimeDaemon::BUS_NAME,
          "object_path" => RuntimeDaemon::OBJECT_PATH,
          "interface" => RuntimeDaemon::INTERFACE
        }
      end

      def application_summary(application)
        diagnostics = runtime.diagnostics(application.fetch("id"))
        pending_checks = diagnostics.fetch("checks", []).count { |check| check["status"] == "pending" }

        {
          "id" => application.fetch("id"),
          "name" => application.fetch("name"),
          "icon" => application.fetch("icon"),
          "mode_label" => MODE_LABELS.fetch(application.fetch("mode"), "Automatic"),
          "compatibility_status" => diagnostics.fetch("status"),
          "compatibility_label" => STATUS_LABELS.fetch(diagnostics.fetch("status"), "Needs review"),
          "pending_action_count" => pending_checks,
          "supported_extensions" => application.fetch("supported_extensions", []),
          "summary" => application_status_summary(pending_checks)
        }
      end

      def application_status_summary(pending_checks)
        return "Ready for desktop integration" if pending_checks.zero?

        "#{pending_checks} compatibility task pending"
      end

      def summary(applications)
        {
          "application_count" => applications.length,
          "known_application_count" => applications.count { |application| application["compatibility_status"] == "known" },
          "pending_action_count" => applications.sum { |application| application["pending_action_count"] }
        }
      end

      class CLI
        SOURCES = %w[auto local dbus].freeze

        def initialize(argv)
          @argv = argv.dup
          @recipe_dir = RuntimeDaemon::DEFAULT_RECIPE_DIR
          @source = "auto"
        end

        def run
          parser.parse!(@argv)
          runtime = runtime_source
          puts KdeCenterModel.new(runtime: runtime).to_json
          0
        rescue OptionParser::ParseError, KeyError, ArgumentError, DBusRuntimeClient::Error => e
          warn "xnix-kde-center-model: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-kde-center-model [--recipe-dir PATH]"
            options.on("--recipe-dir PATH", "Read application recipes from PATH") do |value|
              @recipe_dir = value
            end
            options.on("--source SOURCE", "Read from auto, local, or dbus") do |value|
              raise OptionParser::InvalidArgument, "source must be one of: #{SOURCES.join(", ")}" unless SOURCES.include?(value)

              @source = value
            end
          end
        end

        def runtime_source
          case @source
          when "local"
            local_runtime
          when "dbus"
            DBusRuntimeClient.new
          when "auto"
            DBusRuntimeClient.available? ? DBusRuntimeClient.new : local_runtime
          end
        end

        def local_runtime
          RuntimeDaemon.new(recipe_store: RecipeStore.new(path: @recipe_dir))
        end
      end
    end
  end
end

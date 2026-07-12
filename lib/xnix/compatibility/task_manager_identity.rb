# frozen_string_literal: true

require "json"
require "optparse"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class TaskManagerIdentity
      WINDOW_KINDS = %w[compatibility-application isolated-application].freeze

      def initialize(application_id:, name:, window_kind: "compatibility-application")
        @application_id = application_id
        @name = name
        @window_kind = window_kind
        validate!
      end

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "request_type" => "task-manager-identity",
          "desktop" => "KDE Plasma",
          "application_id" => @application_id,
          "name" => @name,
          "desktop_file" => "xnix-#{@application_id}.desktop",
          "window" => window,
          "task_manager" => task_manager,
          "kwin" => kwin
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def window
        {
          "kind" => @window_kind,
          "class_group" => "xnix-compatibility",
          "resource_name" => @application_id,
          "title_hint" => @name
        }
      end

      def task_manager
        {
          "pinning_allowed" => true,
          "restore_allowed" => true,
          "grouping_key" => @application_id,
          "launcher_url" => "applications:xnix-#{@application_id}.desktop"
        }
      end

      def kwin
        {
          "script_role" => "identity-only",
          "match" => {
            "resource_name" => @application_id,
            "class_group" => "xnix-compatibility"
          },
          "set" => {
            "desktop_file" => "xnix-#{@application_id}.desktop",
            "application_id" => @application_id
          }
        }
      end

      def validate!
        unless @application_id.is_a?(String) && @application_id.match?(/\A[a-z][a-z0-9-]*(?:\.[a-z0-9-]+)+\z/)
          raise ArgumentError, "application id must be a reverse-DNS identifier"
        end
        raise ArgumentError, "name must be a non-empty single-line string" unless @name.is_a?(String) && !@name.empty? && !@name.match?(/[\r\n]/)
        raise ArgumentError, "window kind must be one of: #{WINDOW_KINDS.join(", ")}" unless WINDOW_KINDS.include?(@window_kind)
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @name = nil
          @window_kind = "compatibility-application"
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--app is required" unless @application_id
          raise ArgumentError, "--name is required" unless @name
          raise ArgumentError, "unexpected arguments: #{@argv.join(" ")}" unless @argv.empty?

          puts TaskManagerIdentity.new(
            application_id: @application_id,
            name: @name,
            window_kind: @window_kind
          ).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-compat-window-identity: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-compat-window-identity --app APP_ID --name NAME [--window-kind KIND]"
            options.on("--app APP_ID", "Runtime application id") do |value|
              @application_id = value
            end
            options.on("--name NAME", "User-facing application name") do |value|
              @name = value
            end
            options.on("--window-kind KIND", "Window kind") do |value|
              @window_kind = value
            end
          end
        end
      end
    end
  end
end

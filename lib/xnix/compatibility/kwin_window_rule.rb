# frozen_string_literal: true

require "json"
require "optparse"
require_relative "runtime_daemon"
require_relative "task_manager_identity"

module Xnix
  module Compatibility
    class KWinWindowRule
      def initialize(application_id:, name:, window_kind: "compatibility-application")
        @identity = TaskManagerIdentity.new(
          application_id: application_id,
          name: name,
          window_kind: window_kind
        )
      end

      def to_h
        identity = @identity.to_h

        {
          "version" => RuntimeDaemon::VERSION,
          "request_type" => "kwin-window-rule",
          "desktop" => "KDE Plasma",
          "application_id" => identity.fetch("application_id"),
          "name" => identity.fetch("name"),
          "script_role" => "identity-and-layout",
          "match" => match(identity),
          "set" => rule_set(identity),
          "restore" => restore(identity),
          "safety" => safety
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def match(identity)
        {
          "resource_name" => identity.fetch("window").fetch("resource_name"),
          "class_group" => identity.fetch("window").fetch("class_group"),
          "title_hint" => identity.fetch("window").fetch("title_hint")
        }
      end

      def rule_set(identity)
        {
          "desktop_file" => identity.fetch("desktop_file"),
          "application_id" => identity.fetch("application_id"),
          "task_manager_grouping_key" => identity.fetch("task_manager").fetch("grouping_key"),
          "launcher_url" => identity.fetch("task_manager").fetch("launcher_url"),
          "skip_taskbar" => false,
          "show_in_switcher" => true,
          "placement" => "normal-window"
        }
      end

      def restore(identity)
        {
          "pinning_allowed" => identity.fetch("task_manager").fetch("pinning_allowed"),
          "restore_allowed" => identity.fetch("task_manager").fetch("restore_allowed"),
          "restore_key" => identity.fetch("application_id"),
          "prefer_existing_window" => true
        }
      end

      def safety
        {
          "window_manager_policy_only" => true,
          "runtime_owns_backend_policy" => true,
          "backend_details_exposed" => false
        }
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

          puts KWinWindowRule.new(
            application_id: @application_id,
            name: @name,
            window_kind: @window_kind
          ).to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-kwin-window-rule: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-kwin-window-rule --app APP_ID --name NAME [--window-kind KIND]"
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

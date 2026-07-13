# frozen_string_literal: true

require "json"
require "optparse"
require_relative "dbus_runtime_client"
require_relative "recipe_store"
require_relative "runtime_daemon"
require_relative "task_manager_identity"

module Xnix
  module Compatibility
    class KWinWindowRule
      SOURCES = %w[auto local dbus].freeze

      def initialize(application_id:, name: nil, window_kind: "compatibility-application", runtime: nil)
        @application_id = application_id
        @name = name
        @window_kind = window_kind
        @runtime = runtime
      end

      def to_h
        plan = runtime_plan

        {
          "version" => RuntimeDaemon::VERSION,
          "request_type" => plan.fetch("request_type", "kwin-window-rule"),
          "desktop" => plan.fetch("desktop", "KDE Plasma"),
          "source" => source_metadata,
          "runtime_owned" => plan.fetch("runtime_owned", true),
          "kde_policy_owner" => plan.fetch("kde_policy_owner", false),
          "application_id" => plan.fetch("application_id"),
          "name" => plan.fetch("name"),
          "script_role" => plan.fetch("script_role", "identity-and-layout"),
          "match" => normalized_match(plan),
          "set" => normalized_set(plan),
          "restore" => normalized_restore(plan),
          "safety" => normalized_safety(plan),
          "host_root_modified" => plan.fetch("host_root_modified", false),
          "backend_details_exposed" => plan.fetch("backend_details_exposed", false),
          "desktop_safe_summary" => plan.fetch(
            "desktop_safe_summary",
            "KWin window rules are Runtime-owned identity and layout hints for normal desktop behavior."
          )
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      attr_reader :application_id, :name, :window_kind, :runtime

      def runtime_plan
        return runtime.kwin_window_rule_plan(application_id) if runtime&.respond_to?(:kwin_window_rule_plan)

        fallback_plan
      end

      def source_metadata
        return runtime.source_metadata if runtime&.respond_to?(:source_metadata)

        {
          "kind" => "runtime-local-read-model",
          "bus_name" => RuntimeDaemon::BUS_NAME,
          "object_path" => RuntimeDaemon::OBJECT_PATH,
          "interface" => RuntimeDaemon::INTERFACE
        }
      end

      def fallback_plan
        resolved_name = name || application_id.split(".").last.capitalize
        identity = TaskManagerIdentity.new(
          application_id: application_id,
          name: resolved_name,
          window_kind: window_kind
        ).to_h

        {
          "request_type" => "kwin-window-rule",
          "desktop" => "KDE Plasma",
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "application_id" => identity.fetch("application_id"),
          "name" => identity.fetch("name"),
          "script_role" => "identity-and-layout",
          "match" => fallback_match(identity),
          "set" => fallback_rule_set(identity),
          "restore" => fallback_restore(identity),
          "safety" => fallback_safety,
          "host_root_modified" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => "KWin window rules are Runtime-owned identity and layout hints for normal desktop behavior."
        }
      end

      def fallback_match(identity)
        {
          "resource_name" => identity.fetch("window").fetch("resource_name"),
          "class_group" => identity.fetch("window").fetch("class_group"),
          "title_hint" => identity.fetch("window").fetch("title_hint")
        }
      end

      def fallback_rule_set(identity)
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

      def fallback_restore(identity)
        {
          "pinning_allowed" => identity.fetch("task_manager").fetch("pinning_allowed"),
          "restore_allowed" => identity.fetch("task_manager").fetch("restore_allowed"),
          "restore_key" => identity.fetch("application_id"),
          "prefer_existing_window" => true
        }
      end

      def fallback_safety
        {
          "window_manager_policy_only" => true,
          "runtime_owns_backend_policy" => true,
          "host_root_modified" => false,
          "backend_details_exposed" => false
        }
      end

      def normalized_match(plan)
        return plan.fetch("match") if plan.fetch("match", nil).is_a?(Hash)

        {
          "resource_name" => plan.fetch("resource_name"),
          "class_group" => plan.fetch("class_group"),
          "title_hint" => plan.fetch("title_hint")
        }
      end

      def normalized_set(plan)
        return plan.fetch("set") if plan.fetch("set", nil).is_a?(Hash)

        {
          "desktop_file" => plan.fetch("desktop_file"),
          "application_id" => plan.fetch("application_id"),
          "task_manager_grouping_key" => plan.fetch("task_manager_grouping_key"),
          "launcher_url" => plan.fetch("launcher_url"),
          "skip_taskbar" => plan.fetch("skip_taskbar", false),
          "show_in_switcher" => plan.fetch("show_in_switcher", true),
          "placement" => plan.fetch("placement", "normal-window")
        }
      end

      def normalized_restore(plan)
        return plan.fetch("restore") if plan.fetch("restore", nil).is_a?(Hash)

        {
          "pinning_allowed" => plan.fetch("pinning_allowed", true),
          "restore_allowed" => plan.fetch("restore_allowed", true),
          "restore_key" => plan.fetch("restore_key", plan.fetch("application_id")),
          "prefer_existing_window" => plan.fetch("prefer_existing_window", true)
        }
      end

      def normalized_safety(plan)
        return plan.fetch("safety") if plan.fetch("safety", nil).is_a?(Hash)

        {
          "window_manager_policy_only" => plan.fetch("window_manager_policy_only", true),
          "runtime_owns_backend_policy" => plan.fetch("runtime_owns_backend_policy", true),
          "host_root_modified" => plan.fetch("host_root_modified", false),
          "backend_details_exposed" => plan.fetch("backend_details_exposed", false)
        }
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @application_id = nil
          @name = nil
          @window_kind = "compatibility-application"
          @recipe_dir = RuntimeDaemon::DEFAULT_RECIPE_DIR
          @source = "auto"
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "--app is required" unless @application_id
          validate_application_id(@application_id)
          raise ArgumentError, "unexpected arguments: #{@argv.join(" ")}" unless @argv.empty?

          puts KWinWindowRule.new(
            application_id: @application_id,
            name: @name,
            window_kind: @window_kind,
            runtime: runtime_source
          ).to_json
          0
        rescue OptionParser::ParseError, ArgumentError, DBusRuntimeClient::Error => e
          warn "xnix-kwin-window-rule: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-kwin-window-rule --app APP_ID [--name NAME] [--window-kind KIND] [--source SOURCE] [--recipe-dir PATH]"
            options.on("--app APP_ID", "Runtime application id") do |value|
              @application_id = value
            end
            options.on("--name NAME", "User-facing application name") do |value|
              @name = value
            end
            options.on("--window-kind KIND", "Window kind") do |value|
              @window_kind = value
            end
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

        def validate_application_id(application_id)
          return if application_id.match?(/\A[a-z0-9]+(\.[a-z0-9-]+)+\z/)

          raise ArgumentError, "application id must be reverse-DNS"
        end
      end
    end
  end
end

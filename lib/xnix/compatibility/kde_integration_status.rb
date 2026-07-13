# frozen_string_literal: true

require "json"
require "optparse"
require_relative "dbus_runtime_client"
require_relative "registry_backed_recipe_store"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class KdeIntegrationStatus
      SOURCES = %w[auto local dbus].freeze

      def initialize(runtime: nil)
        @runtime = runtime
      end

      def to_h
        plan = runtime_plan
        entry_points = normalized_entry_points(plan)

        {
          "version" => RuntimeDaemon::VERSION,
          "status_type" => plan.fetch("status_type", "kde-integration-status"),
          "desktop" => plan.fetch("desktop", "KDE Plasma"),
          "source" => source_metadata,
          "runtime_owned" => plan.fetch("runtime_owned", true),
          "kde_policy_owner" => plan.fetch("kde_policy_owner", false),
          "official_desktop_only" => plan.fetch("official_desktop_only", true),
          "stable_desktop_contract" => plan.fetch("stable_desktop_contract", true),
          "entry_points" => entry_points,
          "summary" => summary(plan, entry_points),
          "safety" => {
            "official_desktop_only" => plan.fetch("official_desktop_only", true),
            "stable_desktop_contract" => plan.fetch("stable_desktop_contract", true),
            "host_root_modified" => plan.fetch("host_root_modified", false),
            "backend_details_exposed" => plan.fetch("backend_details_exposed", false)
          },
          "host_root_modified" => plan.fetch("host_root_modified", false),
          "backend_details_exposed" => plan.fetch("backend_details_exposed", false),
          "desktop_safe_summary" => plan.fetch(
            "desktop_safe_summary",
            "KDE Plasma is the only official first-release shell, and all seven entry points are Runtime-backed."
          )
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      def entry_point_ids
        to_h.fetch("entry_points").map { |entry| entry.fetch("id") }
      end

      private

      attr_reader :runtime

      def runtime_plan
        return runtime.kde_integration_status if runtime&.respond_to?(:kde_integration_status)

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

      def summary(plan, entry_points)
        states = entry_points.group_by { |entry| entry.fetch("state") }
        {
          "total" => plan.fetch("entry_point_count", entry_points.length),
          "initial" => plan.fetch("initial_count", states.fetch("initial", []).length),
          "planned" => plan.fetch("planned_count", states.fetch("planned", []).length),
          "complete" => plan.fetch("complete_count", states.fetch("complete", []).length)
        }
      end

      def normalized_entry_points(plan)
        entries = plan.fetch("entry_points", nil)
        return entries.map { |entry| normalized_entry(entry) } if entries.is_a?(Array)

        ids = plan.fetch("entry_point_ids")
        names = plan.fetch("entry_point_names")
        states = plan.fetch("entry_point_states")
        methods = plan.fetch("runtime_methods")

        ids.each_with_index.map do |id, index|
          normalized_entry(
            "id" => id,
            "name" => names.fetch(index),
            "state" => states.fetch(index),
            "runtime_method" => methods.fetch(index),
            "adapter_role" => adapter_role_for(id),
            "runtime_backed" => true,
            "c_runtime_backed" => true,
            "dbus_read_available" => true,
            "kde_policy_owner" => false,
            "summary" => summary_for(id),
            "next_step" => next_step_for(id)
          )
        end
      end

      def normalized_entry(entry)
        {
          "id" => entry.fetch("id"),
          "name" => entry.fetch("name"),
          "state" => entry.fetch("state"),
          "runtime_method" => entry.fetch("runtime_method"),
          "adapter_role" => entry.fetch("adapter_role", adapter_role_for(entry.fetch("id"))),
          "runtime_backed" => entry.fetch("runtime_backed", true),
          "c_runtime_backed" => entry.fetch("c_runtime_backed", true),
          "dbus_read_available" => entry.fetch("dbus_read_available", true),
          "kde_policy_owner" => entry.fetch("kde_policy_owner", false),
          "evidence" => evidence_for(entry),
          "summary" => entry.fetch("summary", summary_for(entry.fetch("id"))),
          "next_step" => entry.fetch("next_step", next_step_for(entry.fetch("id")))
        }
      end

      def evidence_for(entry)
        [
          "Runtime method #{entry.fetch("runtime_method")} exposes this KDE entry point.",
          "Adapter role #{entry.fetch("adapter_role", adapter_role_for(entry.fetch("id")))} keeps KDE presentation separate from Runtime policy."
        ]
      end

      def fallback_plan
        RuntimeDaemon.new(
          recipe_store: RegistryBackedRecipeStore.for_path(RuntimeDaemon::DEFAULT_RECIPE_DIR)
        ).kde_integration_status
      end

      def adapter_role_for(id)
        {
          "launcher" => "standard-desktop-entry",
          "task-manager" => "window-identity-and-restore",
          "file-manager" => "dolphin-service-menu-and-mime",
          "system-tray" => "runtime-status-surface",
          "notifications" => "runtime-event-notification",
          "compatibility-center" => "plasma-read-model",
          "settings" => "user-facing-policy-controls"
        }.fetch(id)
      end

      def summary_for(id)
        {
          "launcher" => "Windows applications appear in the KDE launcher through generated desktop entries.",
          "task-manager" => "Compatibility windows expose grouping, pinning, switcher, and restore identity.",
          "file-manager" => "Dolphin opens selected files through portal-mediated Runtime file-open planning.",
          "system-tray" => "The tray can show Runtime activity, attention state, and bridge readiness.",
          "notifications" => "Runtime events map to KDE notification payloads for install, repair, mode, and approval states.",
          "compatibility-center" => "The Compatibility Center can show Runtime-owned application state and safe action cards.",
          "settings" => "Settings expose user-facing Runtime policy without backend terminology."
        }.fetch(id)
      end

      def next_step_for(id)
        {
          "launcher" => "Connect activation receipts to a production recipe installer.",
          "task-manager" => "Connect Runtime identity plans to a production KWin script and task manager bridge.",
          "file-manager" => "Connect file-open requests to production Runtime launch requests.",
          "system-tray" => "Connect tray status plans to a production Plasma tray surface.",
          "notifications" => "Connect notification plans to the production KDE notification path.",
          "compatibility-center" => "Render live Runtime applications and diagnostics in the Plasmoid.",
          "settings" => "Connect settings plans to a production KDE settings module and persisted Runtime policy."
        }.fetch(id)
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
          @recipe_dir = RuntimeDaemon::DEFAULT_RECIPE_DIR
          @source = "auto"
        end

        def run
          parser.parse!(@argv)
          raise ArgumentError, "unexpected arguments: #{@argv.join(" ")}" unless @argv.empty?

          puts KdeIntegrationStatus.new(runtime: runtime_source).to_json
          0
        rescue OptionParser::ParseError, ArgumentError, DBusRuntimeClient::Error => e
          warn "xnix-kde-integration-status: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-kde-integration-status [--source SOURCE] [--recipe-dir PATH]"
            options.on("--source SOURCE", "Read from auto, local, or dbus") do |value|
              raise OptionParser::InvalidArgument, "source must be one of: #{SOURCES.join(", ")}" unless SOURCES.include?(value)

              @source = value
            end
            options.on("--recipe-dir PATH", "Read application recipes from PATH") do |value|
              @recipe_dir = value
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
          RuntimeDaemon.new(recipe_store: RegistryBackedRecipeStore.for_path(@recipe_dir))
        end
      end
    end
  end
end

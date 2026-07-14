# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"

module Xnix
  module Compatibility
    class KDEShellIntegrationPlan
      PROJECT_ROOT = Pathname.new(__dir__).join("../../..").realpath
      VERSION = PROJECT_ROOT.join("VERSION").read.strip

      COMPONENTS = [
        ["start-menu", "Start menu", "initial", "Generated desktop entries make compatibility applications visible beside native Linux applications.", "GetDesktopEntryPlan"],
        ["task-manager", "Task manager", "initial", "Window identity plans let compatibility windows group and restore like normal applications.", "GetTaskManagerIdentityPlan"],
        ["file-manager", "Dolphin file manager", "initial", "Dolphin actions and MIME associations route file opens through Runtime planning.", "GetFileAssociationPlan"],
        ["system-tray", "System tray", "planned", "Tray status is modeled while live tray bridging remains disabled.", "GetTrayStatus"],
        ["notification-center", "Notification center", "initial", "Runtime notification plans describe install, repair, approval, and mode-change events.", "GetNotificationPlan"],
        ["compatibility-center", "Compatibility Center", "initial", "Compatibility Center reads Runtime diagnostics, review tasks, and repair state.", "GetCompatibilityCenterSummary"],
        ["unified-settings", "Unified settings", "initial", "Settings expose user-facing Runtime policy without backend terminology.", "GetCompatibilitySettings"],
        ["krunner-search", "KRunner search", "initial", "KRunner query plans map natural-language queries to Runtime application identities.", "GetKRunnerQueryPlan"],
        ["kwin-window-management", "KWin window management", "initial", "KWin rule plans provide identity and layout hints without KDE owning backend policy.", "GetKWinWindowRulePlan"]
      ].freeze

      def to_h
        component_rows = components

        {
          "version" => VERSION,
          "plan_type" => "kde-shell-integration-plan",
          "runtime_method" => "GetKDEShellIntegrationPlan",
          "desktop_shell" => "KDE Plasma",
          "official_desktop_only" => true,
          "fallback_desktops_supported" => false,
          "runtime_owned" => true,
          "c_runtime_backed" => true,
          "kde_policy_owner" => false,
          "plasma_fork_required" => false,
          "plasma_source_modified" => false,
          "shell_configuration_written" => false,
          "component_activation_enabled" => false,
          "backend_launch_enabled" => false,
          "backend_details_exposed" => false,
          "host_root_modified" => false,
          "privileged_container_required" => false,
          "component_count" => component_rows.length,
          "initial_component_count" => component_rows.count { |component| component.fetch("state") == "initial" },
          "planned_component_count" => component_rows.count { |component| component.fetch("state") == "planned" },
          "components" => component_rows,
          "desktop_safe_summary" => "KDE Plasma is the first supported shell; Runtime owns policy and KDE components remain replaceable adapters."
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def components
        COMPONENTS.map do |id, label, state, summary, runtime_method|
          {
            "id" => id,
            "label" => label,
            "state" => state,
            "runtime_method" => runtime_method,
            "runtime_owned" => true,
            "kde_policy_owner" => false,
            "shell_writes_enabled" => false,
            "backend_launch_enabled" => false,
            "backend_details_exposed" => false,
            "summary" => summary
          }
        end
      end
    end

    class KDEShellIntegrationPlanCLI
      def initialize(argv)
        @argv = argv.dup
      end

      def run
        OptionParser.new do |options|
          options.banner = "Usage: xnix-kde-shell-integration-plan"
        end.parse!(@argv)
        raise ArgumentError, "unexpected arguments: #{@argv.join(" ")}" unless @argv.empty?

        puts KDEShellIntegrationPlan.new.to_json
      rescue OptionParser::ParseError, ArgumentError => e
        warn "xnix-kde-shell-integration-plan: #{e.message}"
        64
      else
        0
      end
    end
  end
end

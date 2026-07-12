# frozen_string_literal: true

require "json"
require_relative "runtime_daemon"

module Xnix
  module Compatibility
    class KdeIntegrationStatus
      ENTRY_POINTS = [
        {
          "id" => "launcher",
          "name" => "Launcher",
          "state" => "initial",
          "evidence" => [
            "Generated desktop entries call xnix-compat-launch.",
            "Launch requests target the Runtime Launch method without backend commands."
          ],
          "next_step" => "Install generated desktop entries into the KDE application menu during recipe activation."
        },
        {
          "id" => "task-manager",
          "name" => "Task Manager",
          "state" => "planned",
          "evidence" => [],
          "next_step" => "Add KWin window identity rules for Runtime applications."
        },
        {
          "id" => "file-manager",
          "name" => "File Manager",
          "state" => "initial",
          "evidence" => [
            "Dolphin service menu delegates file URIs to xnix-compat-open.",
            "File-open requests require portal-mediated access."
          ],
          "next_step" => "Connect file-open requests to the production Runtime D-Bus Launch method."
        },
        {
          "id" => "system-tray",
          "name" => "System Tray",
          "state" => "planned",
          "evidence" => [],
          "next_step" => "Expose Runtime activity and compatible tray status through a Plasma tray surface."
        },
        {
          "id" => "notifications",
          "name" => "Notifications",
          "state" => "initial",
          "evidence" => [
            "Runtime events can be modeled as KDE notification payloads through xnix-compat-notify.",
            "Notification requests cover install failure, repair, mode-change, and approval events."
          ],
          "next_step" => "Connect notification request models to the production KDE notification D-Bus path."
        },
        {
          "id" => "compatibility-center",
          "name" => "Compatibility Center",
          "state" => "initial",
          "evidence" => [
            "Plasma package exists.",
            "KDE-safe read model can consume Runtime data over the session bus."
          ],
          "next_step" => "Render live Runtime applications and diagnostics in the Plasmoid."
        },
        {
          "id" => "settings",
          "name" => "Settings",
          "state" => "planned",
          "evidence" => [],
          "next_step" => "Add a user-facing settings model for mode, resources, network, devices, and snapshots."
        }
      ].freeze

      def to_h
        {
          "version" => RuntimeDaemon::VERSION,
          "desktop" => "KDE Plasma",
          "official_desktop_only" => true,
          "entry_points" => ENTRY_POINTS,
          "summary" => summary
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def summary
        states = ENTRY_POINTS.group_by { |entry| entry.fetch("state") }

        {
          "total" => ENTRY_POINTS.length,
          "initial" => states.fetch("initial", []).length,
          "planned" => states.fetch("planned", []).length,
          "complete" => states.fetch("complete", []).length
        }
      end

      class CLI
        def run
          puts KdeIntegrationStatus.new.to_json
          0
        end
      end
    end
  end
end

# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"
require_relative "runtime_live_owner_gate"
require_relative "runtime_service_binding"

module Xnix
  module Compatibility
    class RuntimeOwnerSmokePlan
      PROJECT_ROOT = Pathname.new(__dir__).join("../../..").realpath
      VERSION = PROJECT_ROOT.join("VERSION").read.strip

      def initialize(service_binding: RuntimeServiceBinding.new, live_owner_gate: RuntimeLiveOwnerGate.new)
        @service_binding = service_binding
        @live_owner_gate = live_owner_gate
      end

      def to_h
        {
          "version" => VERSION,
          "plan_type" => "runtime-owner-smoke-plan",
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "bus_name" => RuntimeServiceBinding::BUS_NAME,
          "object_path" => RuntimeServiceBinding::OBJECT_PATH,
          "interface" => RuntimeServiceBinding::INTERFACE,
          "activation_binding_ready" => service_binding_model.fetch("activation_binding_ready"),
          "live_dbus_owner_ready" => live_owner_gate_model.fetch("live_dbus_owner_ready"),
          "production_owner_enabled" => live_owner_gate_model.fetch("production_owner_enabled"),
          "owner_transition_ready" => live_owner_gate_model.fetch("owner_transition_ready"),
          "smoke_state" => "planned",
          "smoke_environment" => "restricted-session",
          "steps" => steps,
          "counts" => counts,
          "blocked_actions" => blocked_actions,
          "network_required" => false,
          "host_root_modified" => false,
          "privileged_container_required" => false,
          "system_service_started" => false,
          "production_bus_claimed" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => desktop_safe_summary
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      attr_reader :service_binding, :live_owner_gate

      def service_binding_model
        @service_binding_model ||= service_binding.to_h
      end

      def live_owner_gate_model
        @live_owner_gate_model ||= live_owner_gate.to_h
      end

      def steps
        [
          step("validate-activation-files", "pass",
               "Verify D-Bus service activation, systemd hardening, libexec wrapper, and contract alignment."),
          step("start-packaged-runtime-owner", "pending",
               "Start the packaged Runtime owner in an isolated session without modifying the host root."),
          step("assert-stable-bus-name", "pending",
               "Prove the packaged Runtime owner owns org.xnix.Compatibility1 on the test bus."),
          step("check-read-only-method-parity", "pending",
               "Call the read-only planning methods required by KDE and compare them with the contract."),
          step("reject-write-methods", "pending",
               "Confirm launch, install, snapshot, restore, repair, and settings persistence remain gated."),
          step("verify-non-production-smoke-adapter-boundary", "pending",
               "Confirm the smoke adapter is never accepted as a production Runtime owner."),
          step("report-kde-safe-summary", "pending",
               "Return a Compatibility Center summary without backend details or host paths.")
        ]
      end

      def counts
        statuses = steps.map { |item| item.fetch("status") }
        {
          "total" => statuses.length,
          "passed" => statuses.count("pass"),
          "pending" => statuses.count("pending"),
          "blocked" => statuses.count("blocked")
        }
      end

      def blocked_actions
        [
          "Do not start a host system service from the smoke plan.",
          "Do not claim the production Runtime bus name from the smoke adapter.",
          "Do not enable backend launch, install, repair, restore, or settings persistence.",
          "Do not mutate the host root while planning owner smoke.",
          "Do not expose backend implementation details or host paths to KDE."
        ]
      end

      def step(id, status, summary)
        {
          "id" => id,
          "status" => status,
          "summary" => summary
        }
      end

      def desktop_safe_summary
        return "Runtime activation files must be repaired before owner smoke can run." unless service_binding_model.fetch("activation_binding_ready")

        "Runtime owner smoke is planned; production bus ownership remains disabled until all gates pass."
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
        end

        def run
          parser.parse!(@argv)
          puts RuntimeOwnerSmokePlan.new.to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-runtime-owner-smoke-plan: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-runtime-owner-smoke-plan"
          end
        end
      end
    end
  end
end

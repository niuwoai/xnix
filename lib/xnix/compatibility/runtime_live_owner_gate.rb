# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"
require_relative "runtime_service_binding"

module Xnix
  module Compatibility
    class RuntimeLiveOwnerGate
      PROJECT_ROOT = Pathname.new(__dir__).join("../../..").realpath
      VERSION = PROJECT_ROOT.join("VERSION").read.strip

      def initialize(service_binding: RuntimeServiceBinding.new)
        @service_binding = service_binding
      end

      def to_h
        {
          "version" => VERSION,
          "gate_type" => "runtime-live-owner-gate",
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "bus_name" => RuntimeServiceBinding::BUS_NAME,
          "object_path" => RuntimeServiceBinding::OBJECT_PATH,
          "interface" => RuntimeServiceBinding::INTERFACE,
          "activation_binding_ready" => binding.fetch("activation_binding_ready"),
          "live_dbus_owner_ready" => false,
          "production_owner_enabled" => false,
          "owner_transition_ready" => false,
          "smoke_adapter_available" => binding.fetch("smoke_adapter_available"),
          "smoke_adapter_is_production_owner" => false,
          "kde_may_claim_runtime_ownership" => false,
          "required_gates" => required_gates,
          "blocked_reasons" => blocked_reasons,
          "network_required" => false,
          "host_root_modified" => false,
          "privileged_container_required" => false,
          "backend_details_exposed" => false,
          "desktop_safe_summary" => desktop_safe_summary
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      attr_reader :service_binding

      def binding
        @binding ||= service_binding.to_h
      end

      def required_gates
        [
          gate("activation-binding", binding.fetch("activation_binding_ready") ? "pass" : "blocked",
               "D-Bus activation files, systemd unit, libexec wrapper, and contract must stay aligned."),
          gate("long-running-runtime-owner", "pending",
               "The Runtime needs a packaged long-running process that owns the stable bus name."),
          gate("bus-name-acquisition", "pending",
               "Production smoke must prove the packaged Runtime owns org.xnix.Compatibility1."),
          gate("read-only-method-parity", "pending",
               "The live owner must answer the same read-only planning methods as the smoke adapter."),
          gate("production-recipe-trust", "pending",
               "Production ownership must be gated by signed recipe validation instead of development registry trust.")
        ]
      end

      def blocked_reasons
        [
          "Do not treat the D-Bus smoke adapter as the production Runtime owner.",
          "Do not let KDE own Runtime policy or bus-name readiness decisions.",
          "Do not enable launch, install, repair, restore, or settings persistence from this gate.",
          "Do not mutate the host root while evaluating live-owner readiness.",
          "Do not expose compatibility backend implementation details in live-owner readiness."
        ]
      end

      def gate(id, status, summary)
        {
          "id" => id,
          "status" => status,
          "summary" => summary
        }
      end

      def desktop_safe_summary
        return "Runtime activation files are blocked; production D-Bus ownership cannot be enabled." unless binding.fetch("activation_binding_ready")

        "Runtime activation files are aligned, but production D-Bus ownership remains gated."
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
        end

        def run
          parser.parse!(@argv)
          puts RuntimeLiveOwnerGate.new.to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-runtime-live-owner-gate: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-runtime-live-owner-gate"
          end
        end
      end
    end
  end
end

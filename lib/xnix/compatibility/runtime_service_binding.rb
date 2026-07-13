# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"

module Xnix
  module Compatibility
    class RuntimeServiceBinding
      PROJECT_ROOT = Pathname.new(__dir__).join("../../..").realpath
      VERSION = PROJECT_ROOT.join("VERSION").read.strip
      BUS_NAME = "org.xnix.Compatibility1"
      OBJECT_PATH = "/org/xnix/Compatibility1"
      INTERFACE = "org.xnix.Compatibility1"
      SERVICE_FILE = PROJECT_ROOT.join("runtime/dbus/org.xnix.Compatibility1.service")
      SYSTEMD_UNIT = PROJECT_ROOT.join("runtime/systemd/xnix-compatd.service")
      CONTRACT_FILE = PROJECT_ROOT.join("runtime/dbus/org.xnix.Compatibility1.xml")
      LIBEXEC_WRAPPER = PROJECT_ROOT.join("libexec/xnix/compatd")
      SMOKE_ADAPTER = PROJECT_ROOT.join("runtime/dbus/xnix_compatd_smoke.c")
      PACKAGED_WRAPPER = "/usr/libexec/xnix/compatd"

      def to_h
        {
          "version" => VERSION,
          "binding_type" => "runtime-service-binding",
          "runtime_owned" => true,
          "kde_policy_owner" => false,
          "bus_name" => BUS_NAME,
          "object_path" => OBJECT_PATH,
          "interface" => INTERFACE,
          "activation" => activation,
          "checks" => checks,
          "counts" => counts,
          "activation_binding_ready" => activation_binding_ready?,
          "live_dbus_owner_ready" => false,
          "smoke_adapter_available" => SMOKE_ADAPTER.file?,
          "network_required" => false,
          "host_root_modified" => false,
          "privileged_container_required" => false,
          "desktop_safe_summary" => desktop_safe_summary,
          "backend_details_exposed" => false
        }
      end

      def to_json(*args)
        JSON.pretty_generate(to_h, *args)
      end

      private

      def activation
        {
          "dbus_service_file" => relative(SERVICE_FILE),
          "systemd_unit" => relative(SYSTEMD_UNIT),
          "libexec_wrapper" => relative(LIBEXEC_WRAPPER),
          "dbus_contract" => relative(CONTRACT_FILE),
          "packaged_wrapper" => PACKAGED_WRAPPER
        }
      end

      def checks
        [
          dbus_service_check,
          systemd_unit_check,
          libexec_wrapper_check,
          contract_check,
          live_owner_check
        ]
      end

      def dbus_service_check
        content = read_optional(SERVICE_FILE)
        pass = content.include?("Name=#{BUS_NAME}") &&
               content.include?("SystemdService=xnix-compatd.service") &&
               content.include?("Exec=#{PACKAGED_WRAPPER}")
        check("dbus-service-activation", pass ? "pass" : "blocked", "D-Bus activation points to the packaged Runtime wrapper.")
      end

      def systemd_unit_check
        content = read_optional(SYSTEMD_UNIT)
        pass = content.include?("Type=dbus") &&
               content.include?("BusName=#{BUS_NAME}") &&
               content.include?("ExecStart=#{PACKAGED_WRAPPER}") &&
               content.include?("NoNewPrivileges=yes") &&
               content.include?("ProtectSystem=strict")
        check("systemd-service-hardening", pass ? "pass" : "blocked", "systemd activation uses the stable bus name and hardened service settings.")
      end

      def libexec_wrapper_check
        content = read_optional(LIBEXEC_WRAPPER)
        pass = LIBEXEC_WRAPPER.file? &&
               LIBEXEC_WRAPPER.executable? &&
               content.include?("runtime_daemon")
        check("libexec-wrapper", pass ? "pass" : "blocked", "Packaged Runtime wrapper delegates to the Runtime daemon entry point.")
      end

      def contract_check
        content = read_optional(CONTRACT_FILE)
        pass = content.include?(INTERFACE) &&
               content.include?("ListApplications") &&
               content.include?("GetRuntimeServiceBinding")
        check("dbus-contract", pass ? "pass" : "blocked", "D-Bus contract exposes read-only Runtime service binding status.")
      end

      def live_owner_check
        check("live-dbus-owner", "pending", "A long-running production D-Bus owner is still pending.")
      end

      def check(id, status, summary)
        {
          "id" => id,
          "status" => status,
          "summary" => summary
        }
      end

      def counts
        statuses = checks.map { |item| item.fetch("status") }
        {
          "total" => statuses.length,
          "passed" => statuses.count("pass"),
          "pending" => statuses.count("pending"),
          "blocked" => statuses.count("blocked")
        }
      end

      def activation_binding_ready?
        checks.reject { |item| item.fetch("id") == "live-dbus-owner" }.all? { |item| item.fetch("status") == "pass" }
      end

      def desktop_safe_summary
        return "Runtime service activation files are aligned; live D-Bus ownership remains pending." if activation_binding_ready?

        "Runtime service activation files need repair before production D-Bus ownership can be enabled."
      end

      def read_optional(path)
        return "" unless path.file?

        path.read
      end

      def relative(path)
        path.relative_path_from(PROJECT_ROOT).to_s
      end

      class CLI
        def initialize(argv)
          @argv = argv.dup
        end

        def run
          parser.parse!(@argv)
          puts RuntimeServiceBinding.new.to_json
          0
        rescue OptionParser::ParseError, ArgumentError => e
          warn "xnix-runtime-service-binding: #{e.message}"
          64
        end

        private

        def parser
          OptionParser.new do |options|
            options.banner = "Usage: xnix-runtime-service-binding"
          end
        end
      end
    end
  end
end

#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require "tmpdir"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
INSTALLER = PROJECT_ROOT.join("scripts/install_runtime_activation.rb")

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

def run_json(*command)
  stdout, stderr, status = Open3.capture3(*command)
  assert(status.success?, "#{command.join(" ")} must exit successfully: #{stderr}")
  JSON.parse(stdout)
end

Dir.mktmpdir("xnix-runtime-activation-smoke") do |root|
  stdout, stderr, status = Open3.capture3("ruby", INSTALLER.to_s, "--root", root)
  assert(status.success?, "Runtime activation installer must succeed: #{stderr}")
  assert(stdout.include?(root), "Runtime activation installer must report the staging root")

  root_path = Pathname.new(root)
  wrapper = root_path.join("usr/libexec/xnix/compatd")
  runtime_lib = root_path.join("usr/lib/xnix/compatibility/runtime_daemon.rb")

  assert(wrapper.file?, "Runtime activation smoke requires a staged libexec wrapper")
  assert((wrapper.stat.mode & 0o111) != 0, "Runtime activation smoke requires executable mode bits on the libexec wrapper")
  assert(runtime_lib.file?, "Runtime activation smoke requires installed Runtime Ruby libraries")

  probe = run_json("ruby", wrapper.to_s, "probe")
  assert(probe["version"] == PROJECT_ROOT.join("VERSION").read.strip, "staged Runtime wrapper must report the current version")
  assert(probe["bus_name"] == "org.xnix.Compatibility1", "staged Runtime wrapper must preserve the stable bus name")
  assert(probe["capabilities"]["runtime_write_gates"], "staged Runtime wrapper must expose Runtime write gates")
  assert(!probe["capabilities"]["dbus_binding"], "staged Runtime wrapper must not claim live D-Bus binding")
  assert(!probe["capabilities"]["wine_backend"], "staged Runtime wrapper must not claim Wine backend readiness")
  assert(!probe["capabilities"]["vm_backend"], "staged Runtime wrapper must not claim VM backend readiness")

  parity = run_json("ruby", wrapper.to_s, "dispatch", "GetRuntimeMethodParityManifest")
  assert(parity["read_only_method_parity_ready"], "staged Runtime wrapper must dispatch method parity checks")

  write_gate = run_json("ruby", wrapper.to_s, "dispatch", "GetRuntimeWriteGate", JSON.generate(["Launch"]))
  assert(write_gate["gate_decision"] == "blocked-until-production-backend", "staged Runtime wrapper must dispatch write gates")
  assert(!write_gate["write_method_enabled"], "staged Runtime wrapper must keep write methods disabled")

  _stdout, stderr, status = Open3.capture3("ruby", wrapper.to_s, "dispatch", "Launch", JSON.generate(["org.xnix.sample.notepad", {}]))
  assert(!status.success?, "staged Runtime wrapper must reject gated write dispatch")
  assert(stderr.include?("WriteMethodDisabled"), "staged Runtime wrapper must explain gated write dispatch")
end

puts "PASS: Runtime activation smoke"

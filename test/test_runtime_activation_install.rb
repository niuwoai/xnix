#!/usr/bin/env ruby
# frozen_string_literal: true

require "open3"
require "pathname"
require "tmpdir"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
installer = project_root.join("scripts/install_runtime_activation.rb")

Dir.mktmpdir("xnix-runtime-root") do |root|
  stdout, stderr, status = Open3.capture3("ruby", installer.to_s, "--root", root)
  assert(status.success?, "activation installer must succeed: #{stderr}")
  assert(stdout.include?(root), "activation installer must report the target root")

  root_path = Pathname.new(root)
  libexec = root_path.join("usr/libexec/xnix/compatd")
  runtime_lib = root_path.join("usr/lib/xnix/compatibility/runtime_daemon.rb")
  dbus_service = root_path.join("usr/share/dbus-1/system-services/org.xnix.Compatibility1.service")
  systemd_unit = root_path.join("usr/lib/systemd/system/xnix-compatd.service")
  dbus_smoke = root_path.join("usr/runtime/dbus/xnix_compatd_smoke.c")
  dbus_introspection = root_path.join("usr/runtime/dbus/xnix_compatd_introspection.inc")
  dbus_kde_center = root_path.join("usr/runtime/dbus/xnix_compatd_kde_center.inc")

  assert(libexec.file?, "activation installer must install libexec wrapper")
  assert((libexec.stat.mode & 0o111) != 0, "activation installer must keep libexec wrapper executable mode bits")
  assert(runtime_lib.file?, "activation installer must install Runtime Ruby libraries")
  assert(dbus_service.file?, "activation installer must install D-Bus system service")
  assert(systemd_unit.file?, "activation installer must install systemd unit")
  assert(dbus_smoke.file?, "activation installer must install the D-Bus smoke adapter source")
  assert(dbus_introspection.file?, "activation installer must install the D-Bus introspection include")
  assert(dbus_kde_center.file?, "activation installer must install the D-Bus KDE Center include")
  assert(dbus_smoke.read.include?("xnix_compatd_kde_center.inc"), "installed D-Bus smoke adapter must include KDE Center read models")
  assert(dbus_service.read.include?("SystemdService=xnix-compatd.service"), "D-Bus service must delegate to systemd")
  assert(systemd_unit.read.include?("BusName=org.xnix.Compatibility1"), "systemd unit must own the runtime bus name")

  wrapper_stdout, wrapper_stderr, wrapper_status = Open3.capture3("ruby", libexec.to_s, "probe")
  assert(wrapper_status.success?, "installed libexec wrapper must run against installed Runtime libraries: #{wrapper_stderr}")
  assert(wrapper_stdout.include?("\"bus_name\": \"org.xnix.Compatibility1\""), "installed libexec wrapper must expose the bus name")
end

puts "PASS: compatibility runtime activation install unit tests"

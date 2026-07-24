#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "optparse"
require "pathname"
require_relative "../lib/xnix/ssh_test_key"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip
DEFAULT_STATE_ROOT = PROJECT_ROOT.join(".cache", "xnix", "wine-guest-gui-smoke")
DEFAULT_KERNEL_IMAGE = PROJECT_ROOT.join(".cache", "xnix-wine-i386-output", "images", "bzImage")
DEFAULT_SSH_KEY = Xnix::SshTestKey::PRIVATE_KEY_PATH
DEFAULT_QEMU_BINARY = "qemu-system-i386"
DEFAULT_DISPLAY_NUMBER = 100
DEFAULT_SSH_PORT = "2223"
DEFAULT_GUEST_DISPLAY_HOST = "10.0.2.2"
DEFAULT_GUI_APP = "/usr/lib/wine/i386-windows/winemine.exe"
DEFAULT_WAIT_SECONDS = 10
DEFAULT_BOOT_TIMEOUT_SECONDS = 180

options = {
  execute: false,
  format: "text",
  report_output: "",
  state_root: ENV.fetch("XNIX_WINE_GUI_STATE_ROOT", DEFAULT_STATE_ROOT.to_s),
  qemu_binary: ENV.fetch("XNIX_WINE_GUI_QEMU", DEFAULT_QEMU_BINARY),
  kernel_image: ENV.fetch("XNIX_WINE_GUI_KERNEL", DEFAULT_KERNEL_IMAGE.to_s),
  ssh_key: ENV.fetch("XNIX_WINE_GUI_KEY", DEFAULT_SSH_KEY),
  ssh_port: ENV.fetch("XNIX_WINE_GUI_SSH_PORT", DEFAULT_SSH_PORT),
  display_number: Integer(ENV.fetch("XNIX_WINE_GUI_DISPLAY", DEFAULT_DISPLAY_NUMBER.to_s), 10),
  guest_display_host: ENV.fetch("XNIX_WINE_GUI_GUEST_DISPLAY_HOST", DEFAULT_GUEST_DISPLAY_HOST),
  gui_app: ENV.fetch("XNIX_WINE_GUI_APP", DEFAULT_GUI_APP),
  wait_seconds: Integer(ENV.fetch("XNIX_WINE_GUI_WAIT_SECONDS", DEFAULT_WAIT_SECONDS.to_s), 10),
  boot_timeout_seconds: Integer(ENV.fetch("XNIX_WINE_GUI_BOOT_TIMEOUT_SECONDS", DEFAULT_BOOT_TIMEOUT_SECONDS.to_s), 10)
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/wine_guest_gui_smoke.rb [--execute] [--format text|json]"
  parser.on("--execute", "Start Xvfb, QEMU, and a Wine GUI app smoke.") { options[:execute] = true }
  parser.on("--format FORMAT", "Output format: text or json.") { |value| options[:format] = value }
  parser.on("--report-output PATH", "Write the JSON report to PATH.") { |value| options[:report_output] = value }
  parser.on("--state-root PATH", "State root for smoke evidence.") { |value| options[:state_root] = value }
  parser.on("--qemu-binary PATH", "QEMU binary.") { |value| options[:qemu_binary] = value }
  parser.on("--kernel-image PATH", "Wine guest kernel image.") { |value| options[:kernel_image] = value }
  parser.on("--ssh-key PATH", "Guest SSH private key.") { |value| options[:ssh_key] = value }
  parser.on("--ssh-port PORT", "Loopback SSH port forwarded to the guest.") { |value| options[:ssh_port] = value }
  parser.on("--display-number NUMBER", Integer, "Host Xvfb display number.") { |value| options[:display_number] = value }
  parser.on("--guest-display-host HOST", "Guest-visible host display address.") { |value| options[:guest_display_host] = value }
  parser.on("--gui-app PATH", "Windows GUI app path inside the Wine guest.") { |value| options[:gui_app] = value }
  parser.on("--wait-seconds SECONDS", Integer, "Seconds to wait for the GUI window.") { |value| options[:wait_seconds] = value }
  parser.on("--boot-timeout-seconds SECONDS", Integer, "Seconds to wait for guest SSH.") { |value| options[:boot_timeout_seconds] = value }
  parser.on("--plan-only", "Emit the non-executing plan.") { options[:execute] = false }
end.parse!

abort "wine guest GUI smoke does not accept positional arguments" unless ARGV.empty?
abort "unsupported output format #{options.fetch(:format)}" unless %w[text json].include?(options.fetch(:format))

def command_available?(name)
  _stdout, _stderr, status = Open3.capture3("sh", "-c", "command -v #{name} >/dev/null 2>&1")
  status.success?
end

def run_command(env, *argv)
  stdout, stderr, status = Open3.capture3(env, *argv)
  [stdout, stderr, status.exitstatus]
end

def ssh_args(options)
  [
    "ssh",
    "-p", options.fetch(:ssh_port),
    "-i", options.fetch(:ssh_key),
    "-o", "BatchMode=yes",
    "-o", "ConnectTimeout=5",
    "-o", "IdentitiesOnly=yes",
    "-o", "StrictHostKeyChecking=no",
    "-o", "UserKnownHostsFile=/dev/null",
    "root@127.0.0.1"
  ]
end

def guest_ssh(options, command)
  run_command({}, *ssh_args(options), command)
end

def wait_for_guest_ssh(options)
  deadline = Process.clock_gettime(Process::CLOCK_MONOTONIC) + options.fetch(:boot_timeout_seconds)
  loop do
    _stdout, _stderr, status = guest_ssh(options, "true")
    return true if status.zero?
    return false if Process.clock_gettime(Process::CLOCK_MONOTONIC) >= deadline

    sleep 1
  end
end

def start_xvfb(options, state_root)
  display = ":#{options.fetch(:display_number)}"
  log_path = state_root.join("xvfb.log")
  process = Process.spawn(
    "Xvfb", display,
    "-screen", "0", "1024x768x24",
    "-ac",
    "-listen", "tcp",
    out: log_path.to_s,
    err: log_path.to_s
  )
  sleep 1
  stdout, stderr, status = run_command({ "DISPLAY" => display }, "xdpyinfo")
  return [process, ""] if status.zero?

  Process.kill("TERM", process)
  [nil, "Xvfb display did not become ready: #{stdout} #{stderr}".strip]
end

def start_qemu(options, state_root)
  serial_log = state_root.join("qemu-serial.log")
  args = [
    options.fetch(:qemu_binary),
    "-machine", "q35,accel=tcg",
    "-cpu", "qemu32",
    "-m", "1024M",
    "-smp", "2",
    "-nographic",
    "-serial", "mon:stdio",
    "-no-reboot",
    "-kernel", options.fetch(:kernel_image),
    "-append", "console=ttyS0,115200 panic=-1",
    "-netdev", "user,id=net0,hostfwd=tcp:127.0.0.1:#{options.fetch(:ssh_port)}-:22",
    "-device", "e1000,netdev=net0"
  ]
  process = Process.spawn(*args, out: serial_log.to_s, err: serial_log.to_s)
  [process, serial_log]
end

def stop_process(pid)
  return if pid.nil?

  Process.kill("TERM", pid)
  _, status = Process.wait2(pid)
  status.success?
rescue Errno::ESRCH, Errno::ECHILD
  true
end

def base_report(options)
  state_root = Pathname.new(options.fetch(:state_root)).cleanpath
  {
    "version" => VERSION,
    "schema_version" => "xnix.scripts.wine_guest_gui_smoke.v1",
    "request_type" => "wine-guest-gui-smoke",
    "status" => options.fetch(:execute) ? "running" : "planned",
    "execute" => options.fetch(:execute),
    "backend" => "qemu-guest-wine-x11",
    "gui_app_name" => File.basename(options.fetch(:gui_app)),
    "state_root" => state_root.to_s,
    "qemu_binary" => options.fetch(:qemu_binary),
    "kernel_image" => options.fetch(:kernel_image),
    "ssh_port" => options.fetch(:ssh_port),
    "display_number" => options.fetch(:display_number),
    "guest_display" => "#{options.fetch(:guest_display_host)}:#{options.fetch(:display_number)}",
    "xvfb_required" => true,
    "qemu_required" => true,
    "wine_required" => true,
    "qemu_user_network_restrict_disabled_for_display" => true,
    "loopback_ssh_forwarding_only" => true,
    "privileged_container_required" => false,
    "host_networking_required" => false,
    "docker_socket_mounted" => false,
    "broad_host_mount_required" => false,
    "host_root_modified" => false,
    "wineboot_invoked" => false,
    "x_window_observed" => false,
    "x_window_child_count" => 0,
    "failure_reason" => "",
    "skip_reason" => ""
  }
end

def emit(report, options)
  if !options.fetch(:report_output).strip.empty?
    path = Pathname.new(options.fetch(:report_output)).expand_path
    FileUtils.mkdir_p(path.dirname)
    path.write(JSON.pretty_generate(report) + "\n")
  end

  if options.fetch(:format) == "json"
    puts JSON.pretty_generate(report)
  elsif report.fetch("status") == "passed"
    puts "PASS: Wine guest GUI smoke"
  elsif report.fetch("status") == "skipped"
    puts "SKIP: Wine guest GUI smoke (#{report.fetch("skip_reason")})"
  elsif report.fetch("status") == "planned"
    puts "PLAN: Wine guest GUI smoke"
  else
    puts "FAIL: Wine guest GUI smoke (#{report.fetch("failure_reason")})"
  end
end

report = base_report(options)

unless options.fetch(:execute)
  emit(report, options)
  exit 0
end

state_root = Pathname.new(options.fetch(:state_root)).expand_path(PROJECT_ROOT)
FileUtils.mkdir_p(state_root)
report["state_root"] = state_root.to_s

missing_tools = %w[Xvfb xdpyinfo xwininfo].reject { |tool| command_available?(tool) }
unless missing_tools.empty?
  report["status"] = "skipped"
  report["skip_reason"] = "host X11 smoke tools unavailable: #{missing_tools.join(", ")}"
  emit(report, options)
  exit 0
end

unless File.file?(options.fetch(:kernel_image))
  report["status"] = "skipped"
  report["skip_reason"] = "Wine guest kernel unavailable"
  emit(report, options)
  exit 0
end

unless File.file?(options.fetch(:ssh_key))
  report["status"] = "skipped"
  report["skip_reason"] = "Wine guest SSH key unavailable"
  emit(report, options)
  exit 0
end

xvfb_pid = nil
qemu_pid = nil
begin
  xvfb_pid, xvfb_error = start_xvfb(options, state_root)
  unless xvfb_pid
    report["status"] = "failed"
    report["failure_reason"] = xvfb_error
    emit(report, options)
    exit 1
  end
  report["xvfb_started"] = true

  qemu_pid, _serial_log = start_qemu(options, state_root)
  report["qemu_started"] = true
  unless wait_for_guest_ssh(options)
    report["status"] = "failed"
    report["failure_reason"] = "QEMU guest timed out waiting for SSH"
    emit(report, options)
    exit 1
  end
  report["guest_ssh_ready"] = true

  remote_prefix = "/tmp/xnix-wine-guest-gui-smoke"
  guest_display = "#{options.fetch(:guest_display_host)}:#{options.fetch(:display_number)}"
  wineboot_command = "mkdir -p #{remote_prefix}; " \
                     "DISPLAY=#{guest_display} WINEPREFIX=#{remote_prefix}/wineprefix WINEDEBUG=-all " \
                     "wineboot --init >#{remote_prefix}/wineboot-stdout.txt 2>#{remote_prefix}/wineboot-stderr.txt || true"
  _wineboot_stdout, wineboot_stderr, wineboot_status = guest_ssh(options, wineboot_command)
  unless wineboot_status.zero?
    report["status"] = "failed"
    report["failure_reason"] = "guest Wine boot command failed: #{wineboot_stderr.strip}"
    emit(report, options)
    exit 1
  end
  report["wineboot_invoked"] = true

  start_command = "mkdir -p #{remote_prefix}; " \
                  "DISPLAY=#{guest_display} WINEPREFIX=#{remote_prefix}/wineprefix WINEDEBUG=-all " \
                  "wine #{options.fetch(:gui_app)} >#{remote_prefix}/stdout.txt 2>#{remote_prefix}/stderr.txt " \
                  "& printf '%s\\n' \"$!\" >#{remote_prefix}/pid"
  _stdout, stderr, status = guest_ssh(options, start_command)
  unless status.zero?
    report["status"] = "failed"
    report["failure_reason"] = "guest Wine GUI launch command failed: #{stderr.strip}"
    emit(report, options)
    exit 1
  end

  sleep options.fetch(:wait_seconds)
  display = ":#{options.fetch(:display_number)}"
  xwininfo, xwininfo_stderr, _xwininfo_status = run_command({ "DISPLAY" => display }, "xwininfo", "-root", "-tree")
  state_root.join("xwininfo.txt").write(xwininfo)
  wineboot_stderr_text, _wineboot_ssh_stderr, _wineboot_status = guest_ssh(options, "cat #{remote_prefix}/wineboot-stderr.txt 2>/dev/null || true")
  state_root.join("wineboot-stderr.txt").write(wineboot_stderr_text)
  guest_stderr, _guest_ssh_stderr, _guest_status = guest_ssh(options, "cat #{remote_prefix}/stderr.txt 2>/dev/null || true")
  state_root.join("guest-stderr.txt").write(guest_stderr)

  child_lines = xwininfo.lines.select { |line| line.match?(/^\s+0x[0-9a-f]+/i) }
  report["x_window_child_count"] = child_lines.length
  report["x_window_observed"] = !child_lines.empty?
  report["xwininfo_bytes"] = xwininfo.bytesize
  report["wineboot_stderr_bytes"] = wineboot_stderr_text.bytesize
  report["guest_stderr_bytes"] = guest_stderr.bytesize
  report["guest_graphics_driver_error_observed"] = [wineboot_stderr_text, guest_stderr].any? { |text| text.match?(/graphics driver is missing|no driver could be loaded/i) }
  report["xwininfo_error"] = xwininfo_stderr.strip

  if report.fetch("x_window_observed")
    report["status"] = "passed"
    emit(report, options)
    exit 0
  end

  report["status"] = "failed"
  report["failure_reason"] = "Wine GUI app did not create an X window"
  emit(report, options)
  exit 1
ensure
  guest_ssh(options, "wineserver -k 2>/dev/null || true") if report["guest_ssh_ready"]
  stop_process(qemu_pid)
  stop_process(xvfb_pid)
end

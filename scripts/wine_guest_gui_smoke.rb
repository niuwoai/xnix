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
DEFAULT_GUI_EXECUTABLE = ENV.fetch("XNIX_WINE_GUI_EXECUTABLE", "")
DEFAULT_WAIT_SECONDS = 10
DEFAULT_BOOT_TIMEOUT_SECONDS = 180
DEFAULT_RUNTIME_BIN = ENV.fetch("XNIX_RUNTIME_GO_BIN", "go")
DEFAULT_LAUNCH_MODE = ENV.fetch("XNIX_WINE_GUI_LAUNCH_MODE", "direct")
DEFAULT_OWNER_BIN = ENV.fetch("XNIX_RUNTIME_OWNER_BIN", "go")
DEFAULT_LAUNCHER_BIN = ENV.fetch("XNIX_COMPAT_LAUNCH_BIN", "")
DEFAULT_KNOWN_APP_CACHE_ROOT = ENV.fetch("XNIX_KNOWN_APP_CACHE_ROOT", "")

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
  executable: DEFAULT_GUI_EXECUTABLE,
  runtime_bin: DEFAULT_RUNTIME_BIN,
  launch_mode: DEFAULT_LAUNCH_MODE,
  owner_bin: DEFAULT_OWNER_BIN,
  launcher_bin: DEFAULT_LAUNCHER_BIN,
  known_app_cache_root: DEFAULT_KNOWN_APP_CACHE_ROOT,
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
  parser.on("--executable PATH", "Local Windows GUI .exe copied into the Wine guest before launch.") { |value| options[:executable] = value }
  parser.on("--runtime-bin PATH", "Runtime binary; use `go` to run ./cmd/xnix-runtime-go from source.") { |value| options[:runtime_bin] = value }
  parser.on("--launch-mode MODE", "Launch mode: direct or owner-controlled-launch.") { |value| options[:launch_mode] = value }
  parser.on("--owner-bin PATH", "Runtime owner binary; use `go` to run ./cmd/xnix-runtime-owner from source.") { |value| options[:owner_bin] = value }
  parser.on("--launcher-bin PATH", "Managed xnix-compat-launch binary for owner-controlled launch mode.") { |value| options[:launcher_bin] = value }
  parser.on("--known-app-cache-root PATH", "Known Windows app cache root supplied to the Runtime owner.") { |value| options[:known_app_cache_root] = value }
  parser.on("--wait-seconds SECONDS", Integer, "Seconds to wait for the GUI window.") { |value| options[:wait_seconds] = value }
  parser.on("--boot-timeout-seconds SECONDS", Integer, "Seconds to wait for guest SSH.") { |value| options[:boot_timeout_seconds] = value }
  parser.on("--plan-only", "Emit the non-executing plan.") { options[:execute] = false }
end.parse!

abort "wine guest GUI smoke does not accept positional arguments" unless ARGV.empty?
abort "unsupported output format #{options.fetch(:format)}" unless %w[text json].include?(options.fetch(:format))
abort "unsupported launch mode #{options.fetch(:launch_mode)}" unless %w[direct owner-controlled-launch].include?(options.fetch(:launch_mode))

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

def runtime_base_command(options)
  base = if options.fetch(:runtime_bin).strip == "go"
           ["go", "run", "./cmd/xnix-runtime-go"]
         else
           [options.fetch(:runtime_bin)]
         end
  base
end

def owner_base_command(options)
  if options.fetch(:owner_bin).strip == "go"
    ["go", "run", "./cmd/xnix-runtime-owner"]
  else
    [options.fetch(:owner_bin)]
  end
end

def runtime_command(options)
  base = runtime_base_command(options)
  display_number = options.fetch(:display_number)
  command = [
    *base,
    "windows-app-guest-wine-gui-smoke",
    "--gui-app", options.fetch(:gui_app),
    "--host", "127.0.0.1",
    "--port", options.fetch(:ssh_port),
    "--user", "root",
    "--key", options.fetch(:ssh_key),
    "--remote-dir", "/tmp/xnix-wine-guest-gui-smoke",
    "--guest-display", "#{options.fetch(:guest_display_host)}:#{display_number}",
    "--host-display", ":#{display_number}",
    "--wait", "#{options.fetch(:wait_seconds)}s",
    "--timeout", "#{options.fetch(:boot_timeout_seconds)}s"
  ]
  unless options.fetch(:executable).strip.empty?
    command.push("--executable", options.fetch(:executable))
  end
  command
end

def runtime_gui_evidence_command(options, report_path, evidence_path)
  [
    *runtime_base_command(options),
    "gui-smoke-evidence-preview",
    "--gui-smoke-report", report_path.to_s,
    "--app-id", "org.xnix.apps.mines",
    "--display-name", "Mines",
    "--app-version", VERSION,
    "--output", evidence_path.to_s
  ]
end

def runtime_owner_fixture_command(options, evidence_path, owner_state_root, cache_root)
  [
    *runtime_base_command(options),
    "known-app-runtime-status-launch-owner-fixture-record",
    "--state-root", owner_state_root.to_s,
    "--cache-root", cache_root.to_s,
    "--gui-smoke-evidence-file", evidence_path.to_s
  ]
end

def owner_controlled_launch_command(options, evidence_relative_path)
  [
    *owner_base_command(options),
    "--root", PROJECT_ROOT.to_s,
    "--mode", "smoke-owner",
    "--service-call", "ShowRuntimeControlledLaunch",
    "evidence-relative-path", evidence_relative_path
  ]
end

def owner_guest_gui_app_path(options)
  return "" if options.fetch(:executable).strip.empty?

  "/tmp/xnix-wine-guest-gui-smoke/#{File.basename(options.fetch(:executable))}"
end

def owner_gui_executable_path(options)
  return "" if options.fetch(:executable).strip.empty?

  options.fetch(:executable)
end

def owner_controlled_launch_env(options, owner_state_root, cache_root)
  display_number = options.fetch(:display_number)
  {
    "XNIX_RUNTIME_OWNER_STATE_ROOT" => owner_state_root.to_s,
    "XNIX_RUNTIME_OWNER_KNOWN_APP_CACHE_ROOT" => cache_root.to_s,
    "XNIX_RUNTIME_OWNER_MANAGED_LAUNCHER" => options.fetch(:launcher_bin),
    "XNIX_RUNTIME_OWNER_TIMEOUT" => "#{options.fetch(:boot_timeout_seconds)}s",
    "XNIX_RUNTIME_OWNER_GUEST_TIMEOUT" => "#{options.fetch(:boot_timeout_seconds)}s",
    "XNIX_RUNTIME_OWNER_GUEST_HOST" => "127.0.0.1",
    "XNIX_RUNTIME_OWNER_GUEST_PORT" => options.fetch(:ssh_port),
    "XNIX_RUNTIME_OWNER_GUEST_USER" => "root",
    "XNIX_RUNTIME_OWNER_GUEST_KEY" => options.fetch(:ssh_key),
    "XNIX_RUNTIME_OWNER_GUEST_REMOTE_DIR" => "/tmp/xnix-owner-controlled-wine-gui-smoke",
    "XNIX_RUNTIME_OWNER_GUEST_SSH" => "ssh",
    "XNIX_RUNTIME_OWNER_GUEST_SCP" => "scp",
    "XNIX_RUNTIME_OWNER_GUEST_XWININFO" => "xwininfo",
    "XNIX_RUNTIME_OWNER_GUI_EXECUTABLE" => owner_gui_executable_path(options),
    "XNIX_RUNTIME_OWNER_GUEST_GUI_APP" => "",
    "XNIX_RUNTIME_OWNER_GUEST_DISPLAY" => "#{options.fetch(:guest_display_host)}:#{display_number}",
    "XNIX_RUNTIME_OWNER_HOST_DISPLAY" => ":#{display_number}",
    "XNIX_RUNTIME_OWNER_GUI_WAIT" => "#{options.fetch(:wait_seconds)}s"
  }
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
    "gui_app_name" => File.basename(options.fetch(:executable).strip.empty? ? options.fetch(:gui_app) : options.fetch(:executable)),
    "local_gui_executable_configured" => !options.fetch(:executable).strip.empty?,
    "runtime_go_owned_gui_smoke" => true,
    "launch_mode" => options.fetch(:launch_mode),
    "owner_controlled_launch_requested" => options.fetch(:launch_mode) == "owner-controlled-launch",
    "owner_service_call_planned" => options.fetch(:launch_mode) == "owner-controlled-launch",
    "runtime_owner_bin_configured" => !options.fetch(:owner_bin).strip.empty?,
    "managed_launcher_bin_configured" => !options.fetch(:launcher_bin).strip.empty?,
    "owner_seed_gui_smoke_planned" => options.fetch(:launch_mode) == "owner-controlled-launch",
    "owner_external_gui_app_requested" => options.fetch(:launch_mode) == "owner-controlled-launch" && !options.fetch(:executable).strip.empty?,
    "owner_external_gui_app_path_exposed" => false,
    "owner_external_gui_app_delivery" => options.fetch(:executable).strip.empty? ? "" : "owner-managed-copy",
    "runtime_bin_configured" => !options.fetch(:runtime_bin).strip.empty?,
    "state_root" => state_root.to_s,
    "qemu_binary" => options.fetch(:qemu_binary),
    "kernel_image" => options.fetch(:kernel_image),
    "ssh_port" => options.fetch(:ssh_port),
    "display_number" => options.fetch(:display_number),
    "guest_display" => "#{options.fetch(:guest_display_host)}:#{options.fetch(:display_number)}",
    "xvfb_required" => true,
    "qemu_required" => true,
    "wine_required" => true,
    "guest_x11_driver_available" => false,
    "qemu_user_network_restrict_disabled_for_display" => true,
    "loopback_ssh_forwarding_only" => true,
    "privileged_container_required" => false,
    "host_networking_required" => false,
    "docker_socket_mounted" => false,
    "broad_host_mount_required" => false,
    "host_root_modified" => false,
    "wineboot_invoked" => false,
    "owner_seed_gui_smoke_passed" => false,
    "owner_fixture_ready" => false,
    "owner_service_call_ready" => false,
    "owner_managed_launcher_invoked" => false,
    "owner_delegated_smoke_passed" => false,
    "owner_delegated_evidence_source" => "",
    "x_window_observed" => false,
    "x_window_child_count" => 0,
    "failure_reason" => "",
    "skip_reason" => ""
  }
end

def apply_runtime_payload(report, runtime_payload)
  report["runtime_payload_schema_version"] = runtime_payload.fetch("schema_version")
  report["gui_app_name"] = runtime_payload.fetch("gui_app_name", report.fetch("gui_app_name"))
  report["executable_copied"] = runtime_payload.fetch("executable_copied", false)
  report["guest_x11_driver_available"] = runtime_payload.fetch("guest_x11_driver_available", false)
  report["wineboot_invoked"] = runtime_payload.fetch("wineboot_invoked", false)
  report["x_window_child_count"] = runtime_payload.fetch("x_window_child_count", 0)
  report["x_window_observed"] = runtime_payload.fetch("x_window_observed", false)
  report["x_window_observation_attempts"] = runtime_payload.fetch("x_window_observation_attempts", 0)
  report["xwininfo_bytes"] = runtime_payload.fetch("xwininfo_bytes", 0)
  report["wineboot_stderr_bytes"] = runtime_payload.fetch("wineboot_stderr_bytes", 0)
  report["guest_stderr_bytes"] = runtime_payload.fetch("guest_stderr_bytes", 0)
  report["guest_graphics_driver_error_observed"] = runtime_payload.fetch("guest_graphics_driver_error_observed", false)
  report["kde_safe_output_summary"] = runtime_payload.fetch("kde_safe_output_summary", "")
end

def seed_script_report(base_report, runtime_payload)
  seeded = base_report.merge(
    "status" => runtime_payload.fetch("status"),
    "execute" => true
  )
  apply_runtime_payload(seeded, runtime_payload)
  seeded
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

if options.fetch(:launch_mode) == "owner-controlled-launch" && options.fetch(:launcher_bin).strip.empty?
  report["status"] = "failed"
  report["failure_reason"] = "owner-controlled launch mode requires --launcher-bin"
  emit(report, options)
  exit 1
end

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

  runtime_stdout, runtime_stderr, runtime_status = run_command({}, *runtime_command(options))
  unless runtime_status.zero?
    report["status"] = "failed"
    report["failure_reason"] = "Go Runtime Wine GUI smoke command failed"
    report["runtime_stderr_bytes"] = runtime_stderr.bytesize
    emit(report, options)
    exit 1
  end

  runtime_payload = JSON.parse(runtime_stdout)
  apply_runtime_payload(report, runtime_payload)

  if options.fetch(:launch_mode) == "owner-controlled-launch"
    unless report.fetch("x_window_observed")
      report["status"] = "failed"
      report["failure_reason"] = runtime_payload.fetch("failure_reason", "owner seed GUI smoke did not create an X window")
      emit(report, options)
      exit 1
    end

    report["owner_seed_gui_smoke_passed"] = true
    owner_state_root = state_root.join("owner-controlled-launch-state")
    cache_root = if options.fetch(:known_app_cache_root).strip.empty?
                   state_root.join("known-winapps")
                 else
                   Pathname.new(options.fetch(:known_app_cache_root)).expand_path(PROJECT_ROOT)
                 end
    seed_report_path = state_root.join("owner-seed-gui-smoke.json")
    seed_evidence_path = state_root.join("owner-seed-gui-evidence.json")
    FileUtils.mkdir_p(owner_state_root)
    FileUtils.mkdir_p(cache_root)
    seed_report_path.write(JSON.pretty_generate(seed_script_report(report, runtime_payload)) + "\n")

    evidence_stdout, evidence_stderr, evidence_status = run_command({}, *runtime_gui_evidence_command(options, seed_report_path, seed_evidence_path))
    unless evidence_status.zero?
      report["status"] = "failed"
      report["failure_reason"] = "Go Runtime GUI evidence projection failed"
      report["runtime_stderr_bytes"] = evidence_stderr.bytesize
      emit(report, options)
      exit 1
    end
    evidence_payload = JSON.parse(evidence_stdout)
    smoke_evidence = evidence_payload.fetch("known_app_smoke_evidence", {})
    report["owner_seed_evidence_projected"] = smoke_evidence.fetch("runtime_dispatch_verified", false)

    fixture_stdout, fixture_stderr, fixture_status = run_command({}, *runtime_owner_fixture_command(options, seed_evidence_path, owner_state_root, cache_root))
    unless fixture_status.zero?
      report["status"] = "failed"
      report["failure_reason"] = "Runtime owner launch fixture record failed"
      report["runtime_stderr_bytes"] = fixture_stderr.bytesize
      emit(report, options)
      exit 1
    end
    fixture_payload = JSON.parse(fixture_stdout)
    report["owner_fixture_ready"] = fixture_payload.fetch("fixture_ready", false)
    report["owner_service_call_ready"] = fixture_payload.fetch("owner_service_call_ready", false)
    evidence_relative_path = fixture_payload.fetch("evidence_relative_path", "")
    unless report["owner_fixture_ready"] && report["owner_service_call_ready"] && !evidence_relative_path.empty?
      report["status"] = "failed"
      report["failure_reason"] = "Runtime owner launch fixture was not ready"
      emit(report, options)
      exit 1
    end

    owner_stdout, owner_stderr, owner_status = run_command(
      owner_controlled_launch_env(options, owner_state_root, cache_root),
      *owner_controlled_launch_command(options, evidence_relative_path)
    )
    unless owner_status.zero?
      report["status"] = "failed"
      report["failure_reason"] = "Runtime owner ShowRuntimeControlledLaunch service call failed"
      report["runtime_stderr_bytes"] = owner_stderr.bytesize
      emit(report, options)
      exit 1
    end

    owner_payload = JSON.parse(owner_stdout)
    owner_action = owner_payload.fetch("payload", owner_payload)
    report["owner_payload_schema_version"] = owner_action.fetch("owner_schema_version", "")
    report["owner_managed_launcher_invoked"] = owner_action.fetch("managed_launcher_invoked", false)
    report["owner_existing_managed_launcher_invoked"] = owner_action.fetch("existing_managed_launcher_invoked", false)
    report["owner_delegated_evidence_source"] = owner_action.fetch("delegated_evidence_source", "")
    report["owner_delegated_status"] = owner_action.fetch("delegated_status", "")
    report["owner_delegated_smoke_passed"] = owner_action.fetch("delegated_smoke_passed", false)
    report["owner_delegated_execution_started"] = owner_action.fetch("delegated_execution_started", false)
    report["owner_delegated_controlled_session_window_observed"] = owner_action.fetch("delegated_controlled_session_window_observed", false)
    report["owner_delegated_host_root_modified"] = owner_action.fetch("delegated_host_root_modified", false)
    report["owner_delegated_docker_socket_mounted"] = owner_action.fetch("delegated_docker_socket_mounted", false)
    report["owner_delegated_broad_host_mount_required"] = owner_action.fetch("delegated_broad_host_mount_required", false)
    report["owner_delegated_raw_command_exposed"] = owner_action.fetch("delegated_raw_command_exposed", false)
    report["owner_delegated_backend_details_exposed"] = owner_action.fetch("delegated_backend_details_exposed", false)

    if report["owner_managed_launcher_invoked"] &&
       report["owner_delegated_smoke_passed"] &&
       report["owner_delegated_evidence_source"] == "wine-guest-gui-smoke" &&
       report["owner_delegated_controlled_session_window_observed"]
      report["status"] = "passed"
      emit(report, options)
      exit 0
    end

    report["status"] = "failed"
    report["failure_reason"] = owner_action.fetch("delegated_failure_reason", "Runtime owner controlled GUI launch did not pass")
    emit(report, options)
    exit 1
  end

  if report.fetch("x_window_observed")
    report["status"] = "passed"
    emit(report, options)
    exit 0
  end

  report["status"] = "failed"
  report["failure_reason"] = runtime_payload.fetch("failure_reason", "Wine GUI app did not create an X window")
  emit(report, options)
  exit 1
ensure
  guest_ssh(options, "wineserver -k 2>/dev/null || true") if report["guest_ssh_ready"]
  stop_process(qemu_pid)
  stop_process(xvfb_pid)
end

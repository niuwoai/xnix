#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "optparse"
require "pathname"
require "shellwords"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip

DEFAULT_REMOTE_HOST = ENV.fetch("XNIX_REMOTE_HOST", "root@q4")
DEFAULT_REMOTE_SOURCE_ROOT = ENV.fetch("XNIX_REMOTE_SOURCE_ROOT", "/home/xnix-build/xnix-runtime-source")
DEFAULT_REMOTE_BUILD_ROOT = ENV.fetch("XNIX_REMOTE_BUILD_ROOT", "/home/xnix-build-cache")
DEFAULT_REMOTE_MATERIALS_ROOT = ENV.fetch("XNIX_REMOTE_MATERIALS_ROOT", "/home/xnix-run-materials")
DEFAULT_REMOTE_GO = ENV.fetch("XNIX_REMOTE_GO", "/home/xnix-toolchains/go1.24.4-linux-amd64/bin/go")
DEFAULT_APP_ID = ENV.fetch("XNIX_KNOWN_WINAPP_ID", "7zr")
DEFAULT_LOCAL_SHELL = ENV.fetch("XNIX_LOCAL_SHELL", "/bin/zsh")

options = {
  execute: false,
  sync_source: true,
  local_shell: DEFAULT_LOCAL_SHELL,
  remote_host: DEFAULT_REMOTE_HOST,
  remote_source_root: DEFAULT_REMOTE_SOURCE_ROOT,
  remote_build_root: DEFAULT_REMOTE_BUILD_ROOT,
  remote_materials_root: DEFAULT_REMOTE_MATERIALS_ROOT,
  remote_go: DEFAULT_REMOTE_GO,
  app_id: DEFAULT_APP_ID,
  timeout: ENV.fetch("XNIX_KNOWN_WINAPP_GUEST_TIMEOUT", "90s"),
  boot_timeout: ENV.fetch("XNIX_KNOWN_WINAPP_QEMU_BOOT_TIMEOUT", "180s")
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/remote_known_winapp_guest_wine_smoke.rb [--execute]"
  parser.on("--execute", "Run the remote q4 build and QEMU/Wine known app smoke.") { options[:execute] = true }
  parser.on("--no-sync-source", "Use the existing remote source tree without syncing this checkout.") { options[:sync_source] = false }
  parser.on("--local-shell PATH", "Local login shell used for ssh/rsync alias resolution.") { |value| options[:local_shell] = value }
  parser.on("--remote HOST", "Remote SSH target, default: #{DEFAULT_REMOTE_HOST}") { |value| options[:remote_host] = value }
  parser.on("--remote-source-root PATH", "Remote source root under /home/xnix*.") { |value| options[:remote_source_root] = value }
  parser.on("--remote-build-root PATH", "Remote build cache root under /home/xnix*.") { |value| options[:remote_build_root] = value }
  parser.on("--remote-materials-root PATH", "Remote run materials root under /home/xnix*.") { |value| options[:remote_materials_root] = value }
  parser.on("--remote-go PATH", "Remote Go binary path.") { |value| options[:remote_go] = value }
  parser.on("--app APP_ID", "Known Windows app id, default: #{DEFAULT_APP_ID}") { |value| options[:app_id] = value }
  parser.on("--timeout DURATION", "Guest Wine execution timeout.") { |value| options[:timeout] = value }
  parser.on("--qemu-boot-timeout DURATION", "QEMU SSH boot timeout.") { |value| options[:boot_timeout] = value }
end.parse!

abort "remote known Windows app smoke does not accept positional arguments" unless ARGV.empty?

def ensure_remote_xnix_path!(label, path)
  clean = Pathname.new(path).cleanpath.to_s
  return clean if clean.start_with?("/home/xnix-")

  abort "#{label} must stay under /home/xnix-* on the remote build host"
end

def run_shell(shell, command)
  stdout, stderr, status = Open3.capture3(shell, "-lc", command, chdir: PROJECT_ROOT.to_s)
  [stdout, stderr, status.exitstatus]
end

def shell_join(argv)
  Shellwords.join(argv)
end

remote_host = options.fetch(:remote_host)
remote_source_root = ensure_remote_xnix_path!("remote source root", options.fetch(:remote_source_root))
remote_build_root = ensure_remote_xnix_path!("remote build root", options.fetch(:remote_build_root))
remote_materials_root = ensure_remote_xnix_path!("remote materials root", options.fetch(:remote_materials_root))
remote_bin = "#{remote_build_root}/bin/xnix-runtime-go"
remote_report = "#{remote_materials_root}/state/known-run-#{VERSION}.json"
remote_serial_log = "#{remote_materials_root}/state/qemu-serial-#{VERSION}.log"
remote_key = "#{remote_materials_root}/ssh/id_ed25519"
remote_kernel = "#{remote_materials_root}/wine-guest/bzImage"
remote_cache_root = "#{remote_materials_root}/known-winapps"

runtime_args = [
  remote_bin,
  "windows-known-app-run",
  "--backend", "guest-wine",
  "--start-qemu",
  "--port", "auto",
  "--redact-output",
  "--app", options.fetch(:app_id),
  "--cache-root", remote_cache_root,
  "--key", remote_key,
  "--qemu-binary", "qemu-system-i386",
  "--qemu-kernel", remote_kernel,
  "--qemu-memory", "1024M",
  "--qemu-cpus", "2",
  "--qemu-cpu", "qemu32",
  "--qemu-serial-log", remote_serial_log,
  "--qemu-boot-timeout", options.fetch(:boot_timeout),
  "--timeout", options.fetch(:timeout),
  "--report-output", remote_report
]

remote_go = options.fetch(:remote_go)
remote_go_dir = Pathname.new(remote_go).dirname.to_s
remote_commands = [
  "set -eu",
  shell_join(["mkdir", "-p", "#{remote_build_root}/bin", "#{remote_build_root}/go-build", "#{remote_build_root}/go-mod", "#{remote_build_root}/tmp", "#{remote_materials_root}/state"]),
  "cd #{Shellwords.escape(remote_source_root)}",
  "PATH=#{Shellwords.escape(remote_go_dir)}:$PATH GOCACHE=#{Shellwords.escape("#{remote_build_root}/go-build")} GOMODCACHE=#{Shellwords.escape("#{remote_build_root}/go-mod")} GOTMPDIR=#{Shellwords.escape("#{remote_build_root}/tmp")} #{shell_join([remote_go, "build", "-o", remote_bin, "./cmd/xnix-runtime-go"])}",
  shell_join(runtime_args)
]
remote_command = remote_commands.join("\n")

plan = {
  "schema_version" => "xnix.scripts.remote_known_windows_app_guest_wine_smoke.v1",
  "request_type" => "remote-known-winapp-guest-wine-smoke",
  "status" => options.fetch(:execute) ? "running" : "planned",
  "execute" => options.fetch(:execute),
  "local_shell" => options.fetch(:local_shell),
  "remote_host" => remote_host,
  "source_sync_planned" => options.fetch(:sync_source),
  "remote_source_root" => remote_source_root,
  "remote_build_root" => remote_build_root,
  "remote_materials_root" => remote_materials_root,
  "app_id" => options.fetch(:app_id),
  "backend" => "guest-wine",
  "start_qemu" => true,
  "guest_port" => "auto",
  "redact_output" => true,
  "report_output" => remote_report,
  "serial_log_output" => remote_serial_log,
  "host_root_modified" => false,
  "privileged_container_required" => false,
  "host_networking_required" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false
}

unless options.fetch(:execute)
  puts JSON.pretty_generate(plan)
  exit 0
end

if options.fetch(:sync_source)
  rsync_args = [
    "rsync",
    "-az",
    "--exclude", ".git",
    "--exclude", ".cache",
    "--exclude", "buildroot/output",
    "--exclude", "docs/claude-code-implementation-packages.md",
    "#{PROJECT_ROOT}/",
    "#{remote_host}:#{remote_source_root}/"
  ]
  rsync_stdout, rsync_stderr, rsync_status = run_shell(options.fetch(:local_shell), shell_join(rsync_args))
  unless rsync_status.zero?
    warn rsync_stdout unless rsync_stdout.empty?
    warn rsync_stderr unless rsync_stderr.empty?
    warn "FAIL: remote known Windows app source sync failed"
    exit 1
  end
end

stdout, stderr, status = run_shell(options.fetch(:local_shell), shell_join(["ssh", remote_host, remote_command]))
unless status.zero?
  warn stdout unless stdout.empty?
  warn stderr unless stderr.empty?
  warn "FAIL: remote known Windows app QEMU guest Wine smoke command failed"
  exit 1
end

payload = JSON.parse(stdout)
case payload.fetch("status")
when "passed"
  abort "remote known Windows app smoke used unexpected backend" unless payload.fetch("backend") == "guest-wine"
  abort "remote known Windows app smoke did not use Go QEMU start mode" unless payload.fetch("guest_start_mode") == "go-qemu"
  abort "remote known Windows app smoke did not start QEMU from Go" unless payload.fetch("guest_started") == true
  abort "remote known Windows app smoke did not redact output" unless payload.fetch("raw_output_redacted") == true

  puts "PASS: remote known Windows app QEMU guest Wine smoke (#{payload.fetch("app_id")} #{payload.fetch("app_version")})"
when "skipped"
  puts "SKIP: remote known Windows app QEMU guest Wine smoke (#{payload.fetch("skip_reason")})"
else
  warn stdout
  warn "FAIL: remote known Windows app QEMU guest Wine smoke"
  exit 1
end

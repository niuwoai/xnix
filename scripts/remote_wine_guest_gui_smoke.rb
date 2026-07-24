#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "optparse"
require "pathname"
require "shellwords"
require "timeout"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip

DEFAULT_REMOTE_HOST = ENV.fetch("XNIX_REMOTE_HOST", "root@q4")
DEFAULT_REMOTE_SOURCE_ROOT = ENV.fetch("XNIX_REMOTE_SOURCE_ROOT", "/home/xnix-build/xnix-runtime-source-gui-#{VERSION}")
DEFAULT_REMOTE_MATERIALS_ROOT = ENV.fetch("XNIX_REMOTE_MATERIALS_ROOT", "/home/xnix-run-materials")
DEFAULT_LOCAL_SHELL = ENV.fetch("XNIX_LOCAL_SHELL", "/bin/zsh")

options = {
  execute: false,
  sync_source: true,
  local_shell: DEFAULT_LOCAL_SHELL,
  remote_host: DEFAULT_REMOTE_HOST,
  remote_source_root: DEFAULT_REMOTE_SOURCE_ROOT,
  remote_materials_root: DEFAULT_REMOTE_MATERIALS_ROOT,
  remote_kernel: ENV.fetch("XNIX_WINE_GUI_REMOTE_KERNEL", "#{DEFAULT_REMOTE_MATERIALS_ROOT}/wine-guest/bzImage"),
  remote_ssh_key: ENV.fetch("XNIX_WINE_GUI_REMOTE_SSH_KEY", "#{DEFAULT_REMOTE_MATERIALS_ROOT}/ssh/id_ed25519"),
  remote_build_root: ENV.fetch("XNIX_REMOTE_BUILD_ROOT", "/home/xnix-build-cache"),
  remote_go: ENV.fetch("XNIX_REMOTE_GO", "/home/xnix-toolchains/go1.24.4-linux-amd64/bin/go"),
  report_output: ENV.fetch("XNIX_WINE_GUI_REMOTE_REPORT", "#{DEFAULT_REMOTE_MATERIALS_ROOT}/state/wine-gui-smoke-#{VERSION}.json"),
  state_root: ENV.fetch("XNIX_WINE_GUI_REMOTE_STATE_ROOT", "#{DEFAULT_REMOTE_MATERIALS_ROOT}/state/wine-gui-smoke-#{VERSION}"),
  ssh_port: ENV.fetch("XNIX_WINE_GUI_REMOTE_SSH_PORT", "40229"),
  display_number: Integer(ENV.fetch("XNIX_WINE_GUI_REMOTE_DISPLAY", "106"), 10),
  wait_seconds: Integer(ENV.fetch("XNIX_WINE_GUI_REMOTE_WAIT_SECONDS", "15"), 10),
  remote_timeout_seconds: Integer(ENV.fetch("XNIX_WINE_GUI_REMOTE_TIMEOUT_SECONDS", "300"), 10)
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/remote_wine_guest_gui_smoke.rb [--execute]"
  parser.on("--execute", "Sync and run the q4 Wine guest GUI smoke.") { options[:execute] = true }
  parser.on("--no-sync-source", "Use the existing remote source tree.") { options[:sync_source] = false }
  parser.on("--local-shell PATH", "Local shell used for ssh/rsync alias resolution.") { |value| options[:local_shell] = value }
  parser.on("--remote HOST", "Remote SSH target.") { |value| options[:remote_host] = value }
  parser.on("--remote-source-root PATH", "Remote source root under /home/xnix*.") { |value| options[:remote_source_root] = value }
  parser.on("--remote-materials-root PATH", "Remote materials root under /home/xnix*.") { |value| options[:remote_materials_root] = value }
  parser.on("--remote-kernel PATH", "Remote Wine guest kernel image.") { |value| options[:remote_kernel] = value }
  parser.on("--remote-ssh-key PATH", "Remote Wine guest SSH key.") { |value| options[:remote_ssh_key] = value }
  parser.on("--remote-build-root PATH", "Remote build cache root under /home/xnix*.") { |value| options[:remote_build_root] = value }
  parser.on("--remote-go PATH", "Remote Go binary used to build xnix-runtime-go.") { |value| options[:remote_go] = value }
  parser.on("--report-output PATH", "Remote JSON report output path under /home/xnix*.") { |value| options[:report_output] = value }
  parser.on("--state-root PATH", "Remote smoke state root under /home/xnix*.") { |value| options[:state_root] = value }
  parser.on("--ssh-port PORT", "Loopback SSH port for the temporary QEMU guest.") { |value| options[:ssh_port] = value }
  parser.on("--display-number NUMBER", Integer, "Remote Xvfb display number.") { |value| options[:display_number] = value }
  parser.on("--wait-seconds SECONDS", Integer, "Seconds to wait for the GUI window.") { |value| options[:wait_seconds] = value }
  parser.on("--remote-timeout-seconds SECONDS", Integer, "Timeout for each remote shell operation.") { |value| options[:remote_timeout_seconds] = value }
end.parse!

abort "remote Wine guest GUI smoke does not accept positional arguments" unless ARGV.empty?

def ensure_remote_xnix_path!(label, path)
  clean = Pathname.new(path).cleanpath.to_s
  return clean if clean.start_with?("/home/xnix-")

  abort "#{label} must stay under /home/xnix-* on the remote build host"
end

def run_shell(shell, command, timeout_seconds:)
  stdout = +""
  stderr = +""
  status = nil
  timed_out = false
  Open3.popen3(shell, "-lc", command, chdir: PROJECT_ROOT.to_s, pgroup: true) do |_stdin, out, err, wait_thread|
    out_reader = Thread.new { stdout = out.read }
    err_reader = Thread.new { stderr = err.read }
    begin
      Timeout.timeout(timeout_seconds) { status = wait_thread.value }
    rescue Timeout::Error
      timed_out = true
      begin
        Process.kill("TERM", -wait_thread.pid)
      rescue Errno::ESRCH
        nil
      end
      sleep 2
      begin
        Process.kill("KILL", -wait_thread.pid)
      rescue Errno::ESRCH
        nil
      end
      status = wait_thread.value
    ensure
      out_reader.join
      err_reader.join
    end
  end
  stderr = [stderr, "remote shell operation timed out after #{timeout_seconds}s"].reject(&:empty?).join("\n") if timed_out
  [stdout, stderr, timed_out ? 124 : status.exitstatus]
end

def shell_join(argv)
  Shellwords.join(argv)
end

def ssh_command(remote_host, remote_command)
  [
    "ssh",
    "-o", "BatchMode=yes",
    "-o", "ConnectTimeout=15",
    "-o", "ServerAliveInterval=15",
    "-o", "ServerAliveCountMax=4",
    remote_host,
    remote_command
  ]
end

remote_host = options.fetch(:remote_host)
remote_source_root = ensure_remote_xnix_path!("remote source root", options.fetch(:remote_source_root))
remote_materials_root = ensure_remote_xnix_path!("remote materials root", options.fetch(:remote_materials_root))
remote_kernel = ensure_remote_xnix_path!("remote kernel", options.fetch(:remote_kernel))
remote_ssh_key = ensure_remote_xnix_path!("remote SSH key", options.fetch(:remote_ssh_key))
remote_build_root = ensure_remote_xnix_path!("remote build root", options.fetch(:remote_build_root))
report_output = ensure_remote_xnix_path!("report output", options.fetch(:report_output))
state_root = ensure_remote_xnix_path!("state root", options.fetch(:state_root))
remote_runtime_bin = "#{remote_build_root}/bin/xnix-runtime-go"
remote_go_dir = Pathname.new(options.fetch(:remote_go)).dirname.to_s

plan = {
  "schema_version" => "xnix.scripts.remote_wine_guest_gui_smoke.v1",
  "request_type" => "remote-wine-guest-gui-smoke",
  "version" => VERSION,
  "status" => options.fetch(:execute) ? "running" : "planned",
  "execute" => options.fetch(:execute),
  "remote_host" => remote_host,
  "source_sync_planned" => options.fetch(:sync_source),
  "remote_source_root" => remote_source_root,
  "remote_materials_root" => remote_materials_root,
  "remote_kernel" => remote_kernel,
  "remote_ssh_key" => remote_ssh_key,
  "remote_build_root" => remote_build_root,
  "remote_runtime_bin" => remote_runtime_bin,
  "runtime_build_planned" => true,
  "report_output" => report_output,
  "state_root" => state_root,
  "ssh_port" => options.fetch(:ssh_port),
  "display_number" => options.fetch(:display_number),
  "wait_seconds" => options.fetch(:wait_seconds),
  "remote_timeout_seconds" => options.fetch(:remote_timeout_seconds),
  "remote_command" => "ruby scripts/wine_guest_gui_smoke.rb --execute",
  "backend" => "qemu-guest-wine-x11",
  "gui_app_name" => "winemine.exe",
  "requires_rebuilt_wine_guest_with_x11" => true,
  "privileged_container_required" => false,
  "host_networking_required" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false,
  "host_root_modified" => false
}

unless options.fetch(:execute)
  puts JSON.pretty_generate(plan)
  exit 0
end

if options.fetch(:sync_source)
  mkdir_stdout, mkdir_stderr, mkdir_status = run_shell(options.fetch(:local_shell), shell_join(ssh_command(remote_host, shell_join(["mkdir", "-p", remote_source_root]))), timeout_seconds: options.fetch(:remote_timeout_seconds))
  unless mkdir_status.zero?
    warn mkdir_stdout unless mkdir_stdout.empty?
    warn mkdir_stderr unless mkdir_stderr.empty?
    warn "FAIL: remote Wine guest GUI source root preparation failed"
    exit 1
  end

  rsync_args = [
    "rsync",
    "-az",
    "--timeout", "60",
    "--contimeout", "15",
    "--exclude", ".git",
    "--exclude", ".cache",
    "--exclude", "buildroot/output",
    "--exclude", "docs/claude-code-implementation-packages.md",
    "#{PROJECT_ROOT}/",
    "#{remote_host}:#{remote_source_root}/"
  ]
  rsync_stdout, rsync_stderr, rsync_status = run_shell(options.fetch(:local_shell), shell_join(rsync_args), timeout_seconds: options.fetch(:remote_timeout_seconds))
  unless rsync_status.zero?
    warn rsync_stdout unless rsync_stdout.empty?
    warn rsync_stderr unless rsync_stderr.empty?
    warn "FAIL: remote Wine guest GUI source sync failed"
    exit 1
  end
end

build_command = [
  "set -eu",
  shell_join(["mkdir", "-p", "#{remote_build_root}/bin", "#{remote_build_root}/go-build", "#{remote_build_root}/go-mod", "#{remote_build_root}/tmp"]),
  "cd #{Shellwords.escape(remote_source_root)}",
  "PATH=#{Shellwords.escape(remote_go_dir)}:$PATH GOCACHE=#{Shellwords.escape("#{remote_build_root}/go-build")} GOMODCACHE=#{Shellwords.escape("#{remote_build_root}/go-mod")} GOTMPDIR=#{Shellwords.escape("#{remote_build_root}/tmp")} #{shell_join([options.fetch(:remote_go), "build", "-o", remote_runtime_bin, "./cmd/xnix-runtime-go"])}"
].join("\n")
build_stdout, build_stderr, build_status = run_shell(options.fetch(:local_shell), shell_join(ssh_command(remote_host, build_command)), timeout_seconds: options.fetch(:remote_timeout_seconds))
unless build_status.zero?
  warn build_stdout unless build_stdout.empty?
  warn build_stderr unless build_stderr.empty?
  warn "FAIL: remote Wine guest GUI Runtime build failed"
  exit 1
end

remote_args = [
  "ruby", "scripts/wine_guest_gui_smoke.rb",
  "--execute",
  "--format", "json",
  "--state-root", state_root,
  "--kernel-image", remote_kernel,
  "--ssh-key", remote_ssh_key,
  "--ssh-port", options.fetch(:ssh_port),
  "--display-number", options.fetch(:display_number).to_s,
  "--wait-seconds", options.fetch(:wait_seconds).to_s,
  "--runtime-bin", remote_runtime_bin,
  "--report-output", report_output
]
remote_command = [
  "set -eu",
  "cd #{Shellwords.escape(remote_source_root)}",
  shell_join(remote_args)
].join("\n")

stdout, stderr, status = run_shell(options.fetch(:local_shell), shell_join(ssh_command(remote_host, remote_command)), timeout_seconds: options.fetch(:remote_timeout_seconds))
warn stderr unless stderr.empty?
puts stdout unless stdout.empty?
exit status.zero? ? 0 : 1

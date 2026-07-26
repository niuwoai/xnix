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
DEFAULT_SOURCE_SYNC_MODE = ENV.fetch("XNIX_SOURCE_SYNC_MODE", "runtime")
DEFAULT_REMOTE_SOURCE_ROOT = ENV.fetch("XNIX_REMOTE_SOURCE_ROOT", "/home/xnix-build/xnix-remote-go-build-#{DEFAULT_SOURCE_SYNC_MODE}-#{VERSION}")
DEFAULT_REMOTE_BUILD_ROOT = ENV.fetch("XNIX_REMOTE_BUILD_ROOT", "/home/xnix-build-cache")
DEFAULT_REMOTE_GO = ENV.fetch("XNIX_REMOTE_GO", "/home/xnix-toolchains/go1.24.4-linux-amd64/bin/go")
DEFAULT_LOCAL_SHELL = ENV.fetch("XNIX_LOCAL_SHELL", "/bin/zsh")
DEFAULT_PACKAGES = %w[
  ./cmd/xnix-runtime-go
  ./cmd/xnix-runtime-owner
  ./cmd/xnix-compat-launch
  ./cmd/xnix-compat-open
].freeze

options = {
  execute: false,
  sync_source: true,
  source_sync_mode: DEFAULT_SOURCE_SYNC_MODE,
  local_shell: DEFAULT_LOCAL_SHELL,
  remote_host: DEFAULT_REMOTE_HOST,
  remote_source_root: DEFAULT_REMOTE_SOURCE_ROOT,
  remote_build_root: DEFAULT_REMOTE_BUILD_ROOT,
  remote_go: DEFAULT_REMOTE_GO,
  packages: [],
  goos: ENV.fetch("XNIX_REMOTE_GOOS", "linux"),
  goarch: ENV.fetch("XNIX_REMOTE_GOARCH", "amd64"),
  remote_timeout_seconds: Integer(ENV.fetch("XNIX_REMOTE_GO_BUILD_TIMEOUT_SECONDS", "300"), 10)
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/remote_go_build.rb [--execute] [--package ./cmd/name]"
  parser.on("--execute", "Sync source and build selected Go command packages on q4.") { options[:execute] = true }
  parser.on("--no-sync-source", "Use the existing remote source tree without syncing this checkout.") { options[:sync_source] = false }
  parser.on("--source-sync-mode MODE", "Source sync mode: runtime or full.") { |value| options[:source_sync_mode] = value }
  parser.on("--local-shell PATH", "Local login shell used for ssh/rsync alias resolution.") { |value| options[:local_shell] = value }
  parser.on("--remote HOST", "Remote SSH target, default: #{DEFAULT_REMOTE_HOST}") { |value| options[:remote_host] = value }
  parser.on("--remote-source-root PATH", "Remote source root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_source_root] = value }
  parser.on("--remote-build-root PATH", "Remote build cache root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_build_root] = value }
  parser.on("--remote-go PATH", "Remote Go binary path.") { |value| options[:remote_go] = value }
  parser.on("--package PACKAGE", "Go command package to build; may be repeated. Defaults to core Runtime command packages.") { |value| options[:packages] << value }
  parser.on("--goos GOOS", "Target GOOS, default: linux.") { |value| options[:goos] = value }
  parser.on("--goarch GOARCH", "Target GOARCH, default: amd64.") { |value| options[:goarch] = value }
  parser.on("--remote-timeout-seconds SECONDS", Integer, "Timeout for each remote shell operation.") { |value| options[:remote_timeout_seconds] = value }
end.parse!

abort "remote Go build does not accept positional arguments" unless ARGV.empty?

if !ENV.key?("XNIX_REMOTE_SOURCE_ROOT") && options.fetch(:remote_source_root) == DEFAULT_REMOTE_SOURCE_ROOT
  options[:remote_source_root] = "/home/xnix-build/xnix-remote-go-build-#{options.fetch(:source_sync_mode)}-#{VERSION}"
end

def ensure_remote_xnix_path!(label, path)
  clean = Pathname.new(path).cleanpath.to_s
  return clean if clean.start_with?("/home/xnix-") || clean.start_with?("/tmp/xnix-")

  abort "#{label} must stay under /home/xnix-* or /tmp/xnix-* on the remote build host"
end

def ensure_go_target!(target)
  clean = target.to_s.strip
  abort "Go package target must be non-empty" if clean.empty?
  abort "Go package target must be a repository-relative command package" unless clean.start_with?("./cmd/")
  abort "Go package target must not contain parent traversal" if clean.split("/").include?("..")

  clean
end

def validate_go_pair!(goos, goarch)
  abort "GOOS must be linux or windows for remote q4 builds" unless %w[linux windows].include?(goos)
  abort "GOARCH must be amd64, arm64, or 386 for remote q4 builds" unless %w[amd64 arm64 386].include?(goarch)
end

def source_sync_entries(mode)
  case mode
  when "runtime"
    %w[VERSION go.mod cmd internal runtime scripts lib]
  when "full"
    ["."]
  else
    abort "source sync mode must be runtime or full"
  end
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

def package_binary_name(package)
  File.basename(package)
end

validate_go_pair!(options.fetch(:goos), options.fetch(:goarch))

remote_host = options.fetch(:remote_host)
source_sync_mode = options.fetch(:source_sync_mode)
source_entries = source_sync_entries(source_sync_mode)
remote_source_root = ensure_remote_xnix_path!("remote source root", options.fetch(:remote_source_root))
remote_build_root = ensure_remote_xnix_path!("remote build root", options.fetch(:remote_build_root))
remote_go = options.fetch(:remote_go)
remote_go_dir = Pathname.new(remote_go).dirname.to_s
packages = options.fetch(:packages).empty? ? DEFAULT_PACKAGES : options.fetch(:packages)
packages = packages.map { |target| ensure_go_target!(target) }
remote_output_root = "#{remote_build_root}/bin/#{options.fetch(:goos)}-#{options.fetch(:goarch)}"
remote_cache_root = "#{remote_build_root}/go-build"
remote_mod_root = "#{remote_build_root}/go-mod"
remote_tmp_root = "#{remote_build_root}/tmp"

build_targets = packages.map do |package|
  {
    "package" => package,
    "binary_name" => package_binary_name(package),
    "remote_output" => "#{remote_output_root}/#{package_binary_name(package)}"
  }
end

build_env = [
  "PATH=#{Shellwords.escape(remote_go_dir)}:$PATH",
  "GOOS=#{Shellwords.escape(options.fetch(:goos))}",
  "GOARCH=#{Shellwords.escape(options.fetch(:goarch))}",
  "GOCACHE=#{Shellwords.escape(remote_cache_root)}",
  "GOMODCACHE=#{Shellwords.escape(remote_mod_root)}",
  "GOTMPDIR=#{Shellwords.escape(remote_tmp_root)}"
].join(" ")

build_commands = [
  "set -eu",
  shell_join(["mkdir", "-p", remote_output_root, remote_cache_root, remote_mod_root, remote_tmp_root]),
  "cd #{Shellwords.escape(remote_source_root)}",
  *build_targets.map { |target| "#{build_env} #{shell_join([remote_go, "build", "-o", target.fetch("remote_output"), target.fetch("package")])}" }
]

plan = {
  "schema_version" => "xnix.scripts.remote_go_build.v1",
  "request_type" => "remote-go-build",
  "version" => VERSION,
  "status" => options.fetch(:execute) ? "running" : "planned",
  "execute" => options.fetch(:execute),
  "local_shell" => options.fetch(:local_shell),
  "remote_host" => remote_host,
  "source_sync_planned" => options.fetch(:sync_source),
  "source_sync_mode" => source_sync_mode,
  "source_sync_entry_count" => source_entries.length,
  "source_sync_entries" => source_entries,
  "remote_source_root" => remote_source_root,
  "remote_build_root" => remote_build_root,
  "remote_output_root" => remote_output_root,
  "remote_go" => remote_go,
  "goos" => options.fetch(:goos),
  "goarch" => options.fetch(:goarch),
  "package_count" => build_targets.length,
  "build_targets" => build_targets,
  "remote_timeout_seconds" => options.fetch(:remote_timeout_seconds),
  "remote_build_planned" => true,
  "host_compilation_avoided" => true,
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
  mkdir_stdout, mkdir_stderr, mkdir_status = run_shell(
    options.fetch(:local_shell),
    shell_join(ssh_command(remote_host, shell_join(["mkdir", "-p", remote_source_root]))),
    timeout_seconds: options.fetch(:remote_timeout_seconds)
  )
  unless mkdir_status.zero?
    warn mkdir_stdout unless mkdir_stdout.empty?
    warn mkdir_stderr unless mkdir_stderr.empty?
    warn "FAIL: remote Go build source root preparation failed"
    exit 1
  end

  rsync_sources = if source_sync_mode == "full"
                    ["#{PROJECT_ROOT}/"]
                  else
                    source_entries.map { |entry| "#{PROJECT_ROOT}/#{entry}" }
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
    *rsync_sources,
    "#{remote_host}:#{remote_source_root}/"
  ]
  rsync_stdout, rsync_stderr, rsync_status = run_shell(
    options.fetch(:local_shell),
    shell_join(rsync_args),
    timeout_seconds: options.fetch(:remote_timeout_seconds)
  )
  unless rsync_status.zero?
    warn rsync_stdout unless rsync_stdout.empty?
    warn rsync_stderr unless rsync_stderr.empty?
    warn "FAIL: remote Go build source sync failed"
    exit 1
  end
end

build_stdout, build_stderr, build_status = run_shell(
  options.fetch(:local_shell),
  shell_join(ssh_command(remote_host, build_commands.join("\n"))),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless build_status.zero?
  warn build_stdout unless build_stdout.empty?
  warn build_stderr unless build_stderr.empty?
  warn "FAIL: remote Go build failed"
  exit 1
end

finished = plan.merge(
  "status" => "passed",
  "remote_build_completed" => true,
  "built_binary_count" => build_targets.length
)
puts JSON.pretty_generate(finished)

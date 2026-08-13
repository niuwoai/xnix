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
DEFAULT_REMOTE_SOURCE_ROOT = ENV.fetch("XNIX_REMOTE_SOURCE_ROOT", "/home/xnix-build/xnix-remote-go-test-#{DEFAULT_SOURCE_SYNC_MODE}-#{VERSION}")
DEFAULT_REMOTE_BUILD_ROOT = ENV.fetch("XNIX_REMOTE_BUILD_ROOT", "/home/xnix-build-cache")
DEFAULT_REMOTE_GO = ENV.fetch("XNIX_REMOTE_GO", "/home/xnix-toolchains/go1.24.4-linux-amd64/bin/go")
DEFAULT_LOCAL_SHELL = ENV.fetch("XNIX_LOCAL_SHELL", "/bin/zsh")
DEFAULT_PACKAGES = %w[
  ./internal/runtime/owner
  ./cmd/xnix-runtime-owner
  ./cmd/xnix-compat-open
  ./cmd/xnix-compat-launch
  ./internal/runtime/appidentity
  ./cmd/xnix-runtime-go
].freeze
DEFAULT_RUN_REGEX = "TestServiceCallDispatchesShowRuntimeControlledLaunch|TestRuntimeOwnerCommandRendersShowRuntimeControlledLaunch|TestCompatOpen|TestCompatLaunchPassesFileArgumentToGuestGUINotepad|TestPreviewKDEControlledLaunchAction|TestKDEControlledLaunchActionPreviewCommand|TestPreviewRealWinAppRunReceiptSummary|TestKnownAppSmokeEvidenceFromRealWinAppRunReceiptSummary|TestRealWinAppRunReceiptSummaryPreviewCommand|TestRealWinAppRunReceiptSummaryFeedsCompatibilityAndKDECenterPreviews|TestPreviewRealWinAppRunAcceptance|TestRealWinAppRunAcceptancePreviewCommand|TestPreviewQ4SampleNotepadAcceptance|TestQ4SampleNotepadAcceptancePreviewCommand|TestPreviewQ4WinAppAcceptance|TestQ4WinAppAcceptancePreviewCommand|TestPreviewKnownExistingWinAppAcceptance|TestKnownExistingWinAppAcceptancePreviewCommand|TestPreviewKnownAppMatrixEvidence|TestKnownAppMatrixEvidencePreviewCommand|TestPreviewKnownAppVerifiedCatalog|TestKnownAppVerifiedCatalogPreviewCommand|TestPreviewKnownAppVerifiedCatalogRunPlan|TestRunKnownAppVerifiedCatalogRunPlanExecution|TestRunKnownAppVerifiedCatalogAppExecution|TestKnownAppVerifiedCatalogRunPlanPreviewCommand|TestKnownAppVerifiedCatalogRunPlanExecutionCommand|TestKnownAppVerifiedCatalogAppExecutionCommand|TestPreviewKnownAppVerifiedCatalogRunAcceptance|TestKnownAppVerifiedCatalogRunAcceptancePreviewCommand|TestKnownAppSmokeEvidenceFromKnownAppVerifiedCatalogRunAcceptance|TestRecordKnownAppRuntimeStatusLaunchOwnerFixtureConsumesVerifiedCatalogAppExecutionEvidenceForMessageBox|TestKnownAppRuntimeStatusLaunchOwnerFixtureRecordCommandConsumesVerifiedCatalogAppExecutionEvidence|TestRecordKnownAppVerifiedCatalogLaunchHandoff|TestKnownAppVerifiedCatalogLaunchHandoffRecordCommand|TestRecordKnownAppVerifiedCatalogLaunchMaterialization|TestKnownAppVerifiedCatalogLaunchMaterializationRecordCommand|TestRecordKnownAppVerifiedCatalogDispatchRequest|TestKnownAppVerifiedCatalogDispatchRequestRecordCommand|TestKnownAppVerifiedCatalogDispatchRunnerExecutionCommand|TestCompatibilityCenterPreviewCommandConsumesKnownAppVerifiedCatalogRunAcceptance|TestKDECenterPagePreviewCommandConsumesKnownAppVerifiedCatalogRunAcceptance|TestCompatibilityCenterPreviewCommandConsumesKnownAppVerifiedCatalog|TestCompatibilityCenterPreviewCommandRejectsUnsafeKnownAppVerifiedCatalog|TestKDECenterPagePreviewCommandConsumesKnownAppVerifiedCatalog|TestKDECenterPagePreviewCommandRejectsUnsafeKnownAppVerifiedCatalog|TestKnownPortableCatalogContainsPinnedPuttyStandaloneGUIExecutable|TestLocalGoCompilePolicyPreviewRequiresQ4ByDefault|TestLocalGoCompilePolicyPreviewCommand|TestPreviewQ4KnownPortableWinAppRunPlanConsumesRuntimeCatalog|TestPreviewQ4KnownPortableWinAppRunPlanRejectsKnownNonBundleCatalogApp|TestQ4KnownPortableWinAppRunPlanPreviewCommand"

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
  run_regex: ENV.fetch("XNIX_REMOTE_GO_TEST_RUN", DEFAULT_RUN_REGEX),
  count: Integer(ENV.fetch("XNIX_REMOTE_GO_TEST_COUNT", "1"), 10),
  go_test_timeout: ENV.fetch("XNIX_REMOTE_GO_TEST_TIMEOUT", "5m"),
  remote_timeout_seconds: Integer(ENV.fetch("XNIX_REMOTE_GO_TEST_TIMEOUT_SECONDS", "300"), 10)
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/remote_go_test.rb [--execute] [--package ./internal/runtime/owner] [--run REGEX]"
  parser.on("--execute", "Sync source and run selected Go tests on q4.") { options[:execute] = true }
  parser.on("--no-sync-source", "Use the existing remote source tree without syncing this checkout.") { options[:sync_source] = false }
  parser.on("--source-sync-mode MODE", "Source sync mode: runtime or full.") { |value| options[:source_sync_mode] = value }
  parser.on("--local-shell PATH", "Local login shell used for ssh/rsync alias resolution.") { |value| options[:local_shell] = value }
  parser.on("--remote HOST", "Remote SSH target, default: #{DEFAULT_REMOTE_HOST}") { |value| options[:remote_host] = value }
  parser.on("--remote-source-root PATH", "Remote source root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_source_root] = value }
  parser.on("--remote-build-root PATH", "Remote build cache root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_build_root] = value }
  parser.on("--remote-go PATH", "Remote Go binary path.") { |value| options[:remote_go] = value }
  parser.on("--package PACKAGE", "Go package to test; may be repeated. Defaults to the current Runtime targeted package set.") { |value| options[:packages] << value }
  parser.on("--run REGEX", "Go test -run regex for targeted small-version checks.") { |value| options[:run_regex] = value }
  parser.on("--count COUNT", Integer, "Go test -count value, default: 1.") { |value| options[:count] = value }
  parser.on("--go-test-timeout TEXT", "Go test -timeout value, default: 5m.") { |value| options[:go_test_timeout] = value }
  parser.on("--remote-timeout-seconds SECONDS", Integer, "Timeout for each remote shell operation.") { |value| options[:remote_timeout_seconds] = value }
end.parse!

abort "remote Go test does not accept positional arguments" unless ARGV.empty?

if !ENV.key?("XNIX_REMOTE_SOURCE_ROOT") && options.fetch(:remote_source_root) == DEFAULT_REMOTE_SOURCE_ROOT
  options[:remote_source_root] = "/home/xnix-build/xnix-remote-go-test-#{options.fetch(:source_sync_mode)}-#{VERSION}"
end

def ensure_remote_xnix_path!(label, path)
  clean = Pathname.new(path).cleanpath.to_s
  return clean if clean.start_with?("/home/xnix-") || clean.start_with?("/tmp/xnix-")

  abort "#{label} must stay under /home/xnix-* or /tmp/xnix-* on the remote build host"
end

def ensure_go_package!(target)
  clean = target.to_s.strip
  abort "Go package target must be non-empty" if clean.empty?
  abort "Go package target must be repository-relative" unless clean.start_with?("./cmd/") || clean.start_with?("./internal/")
  abort "Go package target must not contain parent traversal" if clean.split("/").include?("..")

  clean
end

def ensure_go_test_regex!(regex)
  clean = regex.to_s.strip
  abort "Go test regex must be non-empty" if clean.empty?
  abort "Go test regex must be single-line" if clean.include?("\n") || clean.include?("\r")

  clean
end

def ensure_positive_count!(count)
  abort "Go test count must be positive" unless count.positive?

  count
end

def ensure_go_test_timeout!(timeout_text)
  clean = timeout_text.to_s.strip
  abort "Go test timeout must be non-empty" if clean.empty?
  abort "Go test timeout must be single-line" if clean.include?("\n") || clean.include?("\r")

  clean
end

def source_sync_entries(mode)
  case mode
  when "runtime"
    %w[.dockerignore Dockerfile VERSION go.mod CLAUDE.md cmd internal runtime scripts lib kde docs]
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

source_sync_mode = options.fetch(:source_sync_mode)
source_entries = source_sync_entries(source_sync_mode)
remote_host = options.fetch(:remote_host)
remote_source_root = ensure_remote_xnix_path!("remote source root", options.fetch(:remote_source_root))
remote_build_root = ensure_remote_xnix_path!("remote build root", options.fetch(:remote_build_root))
remote_go = options.fetch(:remote_go)
remote_go_dir = Pathname.new(remote_go).dirname.to_s
packages = options.fetch(:packages).empty? ? DEFAULT_PACKAGES : options.fetch(:packages)
packages = packages.map { |target| ensure_go_package!(target) }
run_regex = ensure_go_test_regex!(options.fetch(:run_regex))
count = ensure_positive_count!(options.fetch(:count))
go_test_timeout = ensure_go_test_timeout!(options.fetch(:go_test_timeout))
remote_cache_root = "#{remote_build_root}/go-build"
remote_mod_root = "#{remote_build_root}/go-mod"
remote_tmp_root = "#{remote_build_root}/tmp"

test_env = [
  "PATH=#{Shellwords.escape(remote_go_dir)}:$PATH",
  "GOCACHE=#{Shellwords.escape(remote_cache_root)}",
  "GOMODCACHE=#{Shellwords.escape(remote_mod_root)}",
  "GOTMPDIR=#{Shellwords.escape(remote_tmp_root)}"
].join(" ")
test_command = "#{test_env} #{shell_join([remote_go, "test", *packages, "-run", run_regex, "-count", count.to_s, "-timeout", go_test_timeout])}"
remote_commands = [
  "set -eu",
  shell_join(["mkdir", "-p", remote_cache_root, remote_mod_root, remote_tmp_root]),
  "cd #{Shellwords.escape(remote_source_root)}",
  test_command
]

plan = {
  "schema_version" => "xnix.scripts.remote_go_test.v1",
  "request_type" => "remote-go-test",
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
  "remote_go" => remote_go,
  "package_count" => packages.length,
  "packages" => packages,
  "run_regex" => run_regex,
  "count" => count,
  "go_test_timeout" => go_test_timeout,
  "remote_timeout_seconds" => options.fetch(:remote_timeout_seconds),
  "remote_test_planned" => true,
  "targeted_test_required" => true,
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
    warn "FAIL: remote Go test source root preparation failed"
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
    warn "FAIL: remote Go test source sync failed"
    exit 1
  end
end

test_stdout, test_stderr, test_status = run_shell(
  options.fetch(:local_shell),
  shell_join(ssh_command(remote_host, remote_commands.join("\n"))),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless test_status.zero?
  warn test_stdout unless test_stdout.empty?
  warn test_stderr unless test_stderr.empty?
  warn "FAIL: remote Go test failed"
  exit 1
end

finished = plan.merge(
  "status" => "passed",
  "remote_test_completed" => true,
  "tested_package_count" => packages.length
)
puts JSON.pretty_generate(finished)

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
Q4_WINAPP_SMOKE = PROJECT_ROOT.join("scripts/q4_winapp_smoke.rb")

DEFAULT_REMOTE_HOST = ENV.fetch("XNIX_REMOTE_HOST", "root@q4")
DEFAULT_REMOTE_TIMEOUT_SECONDS = Integer(ENV.fetch("XNIX_Q4_MESSAGEBOX_TIMEOUT_SECONDS", "420"), 10)
DEFAULT_LOCAL_SHELL = ENV.fetch("XNIX_LOCAL_SHELL", "/bin/zsh")
DEFAULT_REMOTE_SOURCE_ROOT = ENV.fetch("XNIX_Q4_MESSAGEBOX_SOURCE_ROOT", "/home/xnix-build/xnix-messagebox-fixture-#{VERSION}")
DEFAULT_REMOTE_BUILD_ROOT = ENV.fetch("XNIX_REMOTE_BUILD_ROOT", "/home/xnix-build-cache")
DEFAULT_REMOTE_MATERIALS_ROOT = ENV.fetch("XNIX_REMOTE_MATERIALS_ROOT", "/home/xnix-run-materials")
DEFAULT_REMOTE_GO = ENV.fetch("XNIX_REMOTE_GO", "/home/xnix-toolchains/go1.24.4-linux-amd64/bin/go")
DEFAULT_APP_ID = "org.xnix.apps.messagebox"
DEFAULT_DISPLAY_NAME = "Xnix MessageBox"
DEFAULT_WINDOW_MATCH = "Xnix Windows GUI Smoke"

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
  display_name: DEFAULT_DISPLAY_NAME,
  window_match: DEFAULT_WINDOW_MATCH,
  remote_timeout_seconds: DEFAULT_REMOTE_TIMEOUT_SECONDS,
  output: ENV.fetch("XNIX_Q4_MESSAGEBOX_OUTPUT", "")
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/q4_messagebox_smoke.rb [--execute]"
  parser.on("--execute", "Build the MessageBox Windows GUI fixture on q4 and run it through the generic q4 Windows app smoke.") { options[:execute] = true }
  parser.on("--no-sync-source", "Use the existing q4 fixture source root before building the Windows executable.") { options[:sync_source] = false }
  parser.on("--local-shell PATH", "Local shell used for ssh/rsync alias resolution.") { |value| options[:local_shell] = value }
  parser.on("--remote HOST", "Remote SSH target, default: #{DEFAULT_REMOTE_HOST}") { |value| options[:remote_host] = value }
  parser.on("--remote-source-root PATH", "Remote fixture source root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_source_root] = value }
  parser.on("--remote-build-root PATH", "Remote build cache root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_build_root] = value }
  parser.on("--remote-materials-root PATH", "Remote run materials root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_materials_root] = value }
  parser.on("--remote-go PATH", "Remote Go binary used to cross-compile the Windows fixture.") { |value| options[:remote_go] = value }
  parser.on("--app-id ID", "Evidence application id, default: #{DEFAULT_APP_ID}") { |value| options[:app_id] = value }
  parser.on("--display-name NAME", "Evidence display name, default: #{DEFAULT_DISPLAY_NAME}") { |value| options[:display_name] = value }
  parser.on("--window-match TEXT", "Window title/text required for GUI observation, default: #{DEFAULT_WINDOW_MATCH}") { |value| options[:window_match] = value }
  parser.on("--remote-timeout-seconds SECONDS", Integer, "Timeout for remote q4 operations.") { |value| options[:remote_timeout_seconds] = value }
  parser.on("--output PATH", "Write the plan or passed result JSON under this checkout or /tmp/xnix-*.") { |value| options[:output] = value }
end.parse!

abort "q4 MessageBox smoke does not accept positional arguments" unless ARGV.empty?

def ensure_local_output_path!(path)
  return nil if path.to_s.strip.empty?

  clean = Pathname.new(path).expand_path(PROJECT_ROOT).cleanpath
  return clean if clean.to_s.start_with?(PROJECT_ROOT.to_s)
  return clean if clean.to_s.start_with?("/tmp/xnix-")

  abort "output path must stay under this checkout or /tmp/xnix-*"
end

def ensure_remote_xnix_path!(label, path)
  clean = Pathname.new(path).cleanpath.to_s
  return clean if clean.start_with?("/home/xnix-") || clean.start_with?("/tmp/xnix-")

  abort "#{label} must stay under /home/xnix-* or /tmp/xnix-* on q4"
end

def emit_json(payload, output_path)
  text = JSON.pretty_generate(payload) + "\n"
  File.write(output_path, text) if output_path
  puts text
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

def bool(payload, key)
  payload[key] == true
end

remote_host = options.fetch(:remote_host)
remote_source_root = ensure_remote_xnix_path!("remote source root", options.fetch(:remote_source_root))
remote_build_root = ensure_remote_xnix_path!("remote build root", options.fetch(:remote_build_root))
remote_materials_root = ensure_remote_xnix_path!("remote materials root", options.fetch(:remote_materials_root))
remote_fixture_root = "#{remote_materials_root}/fixtures"
remote_executable = "#{remote_fixture_root}/xnix-messagebox-smoke-#{VERSION}.exe"
remote_go_dir = Pathname.new(options.fetch(:remote_go)).dirname.to_s
output_path = ensure_local_output_path!(options.fetch(:output))
delegated_output = "/tmp/xnix-q4-messagebox-winapp-#{VERSION}.json"

delegated_command = [
  "ruby",
  "scripts/q4_winapp_smoke.rb",
  "--execute",
  "--remote-executable", remote_executable,
  "--app-id", options.fetch(:app_id),
  "--display-name", options.fetch(:display_name),
  "--window-match", options.fetch(:window_match),
  "--remote", remote_host,
  "--remote-timeout-seconds", options.fetch(:remote_timeout_seconds).to_s,
  "--output", delegated_output
]

plan = {
  "schema_version" => "xnix.scripts.q4_messagebox_smoke.v1",
  "request_type" => "q4-messagebox-smoke",
  "version" => VERSION,
  "status" => options.fetch(:execute) ? "running" : "planned",
  "execute" => options.fetch(:execute),
  "remote_host" => remote_host,
  "source_sync_planned" => options.fetch(:sync_source),
  "remote_source_root" => remote_source_root,
  "remote_build_root" => remote_build_root,
  "remote_materials_root" => remote_materials_root,
  "remote_go" => options.fetch(:remote_go),
  "fixture_source" => "test/fixtures/winapp/messagebox",
  "fixture_goos" => "windows",
  "fixture_goarch" => "386",
  "remote_windows_fixture_build_planned" => true,
  "remote_windows_executable_built" => false,
  "remote_executable_configured" => true,
  "remote_executable_path_exposed" => false,
  "app_id" => options.fetch(:app_id),
  "display_name" => options.fetch(:display_name),
  "window_match" => options.fetch(:window_match),
  "delegated_script" => "scripts/q4_winapp_smoke.rb",
  "delegated_command" => delegated_command,
  "delegated_output" => delegated_output,
  "output_path" => output_path ? output_path.to_s : "",
  "q4_compile_required" => true,
  "host_compilation_avoided" => true,
  "host_root_modified" => false,
  "privileged_container_required" => false,
  "host_networking_required" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false
}

unless options.fetch(:execute)
  emit_json(plan, output_path)
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
    warn "FAIL: q4 MessageBox fixture source root preparation failed"
    exit 1
  end

  rsync_args = [
    "rsync",
    "-az",
    "--timeout", "60",
    "--contimeout", "15",
    "#{PROJECT_ROOT}/go.mod",
    "#{PROJECT_ROOT}/test/fixtures/winapp/messagebox",
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
    warn "FAIL: q4 MessageBox fixture source sync failed"
    exit 1
  end
end

build_command = [
  "set -eu",
  shell_join(["mkdir", "-p", remote_fixture_root, "#{remote_build_root}/go-build", "#{remote_build_root}/go-mod", "#{remote_build_root}/tmp"]),
  "cd #{Shellwords.escape(remote_source_root)}",
  "PATH=#{Shellwords.escape(remote_go_dir)}:$PATH GOOS=windows GOARCH=386 GOCACHE=#{Shellwords.escape("#{remote_build_root}/go-build")} GOMODCACHE=#{Shellwords.escape("#{remote_build_root}/go-mod")} GOTMPDIR=#{Shellwords.escape("#{remote_build_root}/tmp")} #{shell_join([options.fetch(:remote_go), "build", "-o", remote_executable, "./messagebox"])}"
].join("\n")
build_stdout, build_stderr, build_status = run_shell(
  options.fetch(:local_shell),
  shell_join(ssh_command(remote_host, build_command)),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless build_status.zero?
  warn build_stdout unless build_stdout.empty?
  warn build_stderr unless build_stderr.empty?
  warn "FAIL: q4 MessageBox Windows fixture build failed"
  exit 1
end

delegated_args = delegated_command.dup
delegated_args[1] = Q4_WINAPP_SMOKE.to_s
stdout, stderr, status = Open3.capture3(*delegated_args, chdir: PROJECT_ROOT.to_s)
unless status.success?
  warn stdout unless stdout.empty?
  warn stderr unless stderr.empty?
  warn "FAIL: q4 MessageBox delegated Windows app smoke failed"
  exit 1
end

delegated = JSON.parse(stdout)
unless delegated.fetch("status") == "passed" &&
       bool(delegated, "window_observed") &&
       bool(delegated, "window_match_observed") &&
       bool(delegated, "go_owned_q4_winapp_acceptance_ready")
  warn stdout
  warn "FAIL: q4 MessageBox delegated Windows app smoke did not satisfy Go-owned acceptance"
  exit 1
end

result = plan.merge(
  "status" => "passed",
  "remote_windows_executable_built" => true,
  "delegated_q4_winapp_smoke_schema" => delegated.fetch("schema_version"),
  "delegated_q4_winapp_smoke_request_type" => delegated.fetch("request_type"),
  "delegated_q4_winapp_smoke_status" => delegated.fetch("status"),
  "delegated_execute_result_schema" => delegated.fetch("delegated_execute_result_schema"),
  "delegated_execute_result_output_written" => bool(delegated, "delegated_execute_result_output_written"),
  "remote_build_completed" => bool(delegated, "remote_build_completed"),
  "evidence_output_written" => bool(delegated, "evidence_output_written"),
  "kde_page_output_written" => bool(delegated, "kde_page_output_written"),
  "window_observed" => bool(delegated, "window_observed"),
  "window_match_observed" => bool(delegated, "window_match_observed"),
  "go_owned_q4_winapp_acceptance_schema" => delegated.fetch("go_owned_q4_winapp_acceptance_schema"),
  "go_owned_q4_winapp_acceptance_request_type" => delegated.fetch("go_owned_q4_winapp_acceptance_request_type"),
  "go_owned_q4_winapp_acceptance_ready" => delegated.fetch("go_owned_q4_winapp_acceptance_ready"),
  "go_owned_q4_winapp_acceptance_consumed" => delegated.fetch("go_owned_q4_winapp_acceptance_consumed"),
  "go_owned_q4_winapp_acceptance_path_exposed" => delegated.fetch("go_owned_q4_winapp_acceptance_path_exposed"),
  "go_owned_q4_winapp_acceptance_remote_host_exposed" => delegated.fetch("go_owned_q4_winapp_acceptance_remote_host_exposed"),
  "go_owned_q4_winapp_acceptance_delegated_command_exposed" => delegated.fetch("go_owned_q4_winapp_acceptance_delegated_command_exposed"),
  "go_owned_q4_winapp_acceptance_remote_executable_path_exposed" => delegated.fetch("go_owned_q4_winapp_acceptance_remote_executable_path_exposed"),
  "go_owned_q4_winapp_acceptance_remote_file_argument_path_exposed" => delegated.fetch("go_owned_q4_winapp_acceptance_remote_file_argument_path_exposed"),
  "host_root_modified" => bool(delegated, "host_root_modified"),
  "privileged_container_required" => bool(delegated, "privileged_container_required"),
  "host_networking_required" => bool(delegated, "host_networking_required"),
  "docker_socket_mounted" => bool(delegated, "docker_socket_mounted"),
  "broad_host_mount_required" => bool(delegated, "broad_host_mount_required")
)

emit_json(result, output_path)

#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "optparse"
require "pathname"
require "securerandom"
require "shellwords"
require "timeout"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip
REMOTE_GO_BUILD = PROJECT_ROOT.join("scripts/remote_go_build.rb")

DEFAULT_REMOTE_HOST = ENV.fetch("XNIX_REMOTE_HOST", "root@q4")
DEFAULT_LOCAL_SHELL = ENV.fetch("XNIX_LOCAL_SHELL", "/bin/zsh")
DEFAULT_REMOTE_SOURCE_ROOT = ENV.fetch("XNIX_Q4_DBUS_REMOTE_SOURCE_ROOT", "/home/xnix-build/xnix-remote-go-build-runtime-#{VERSION}")
DEFAULT_REMOTE_BUILD_ROOT = ENV.fetch("XNIX_REMOTE_BUILD_ROOT", "/home/xnix-build-cache")
DEFAULT_REMOTE_MATERIALS_ROOT = ENV.fetch("XNIX_Q4_DBUS_MATERIALS_ROOT", "/home/xnix-run-materials/xnix-q4-dbus-controlled-launch-owner-fixture-#{VERSION}")
DEFAULT_CONTAINER_GOARCH = ENV.fetch("XNIX_Q4_DBUS_CONTAINER_GOARCH", "amd64")
DEFAULT_TOOLS_IMAGE = ENV.fetch("XNIX_Q4_DBUS_TOOLS_IMAGE", "xnix-builder-tools:#{VERSION}-dbus-#{DEFAULT_CONTAINER_GOARCH}")
DEFAULT_TOOLS_TARGET = ENV.fetch("XNIX_Q4_DBUS_TOOLS_TARGET", "dbus-tools")
DEFAULT_TOOLS_BASE_IMAGE = ENV.fetch("XNIX_Q4_DBUS_TOOLS_BASE_IMAGE", "python:3.12-slim")
DEFAULT_TIMEOUT_SECONDS = Integer(ENV.fetch("XNIX_Q4_DBUS_TIMEOUT_SECONDS", "420"), 10)
MESSAGEBOX_APP_ID = "org.xnix.apps.messagebox"
SCRATCH_ROOT = "/workspace/.xnix-dbus-controlled-launch-scratch"

options = {
  execute: false,
  sync_source: true,
  local_shell: DEFAULT_LOCAL_SHELL,
  remote_host: DEFAULT_REMOTE_HOST,
  remote_source_root: DEFAULT_REMOTE_SOURCE_ROOT,
  remote_build_root: DEFAULT_REMOTE_BUILD_ROOT,
  remote_materials_root: DEFAULT_REMOTE_MATERIALS_ROOT,
  tools_image: DEFAULT_TOOLS_IMAGE,
  tools_target: DEFAULT_TOOLS_TARGET,
  tools_base_image: DEFAULT_TOOLS_BASE_IMAGE,
  container_goarch: DEFAULT_CONTAINER_GOARCH,
  app_id: MESSAGEBOX_APP_ID,
  app_execution_file: "",
  timeout_seconds: DEFAULT_TIMEOUT_SECONDS,
  output: ENV.fetch("XNIX_Q4_DBUS_OUTPUT", "")
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/q4_dbus_controlled_launch_owner_fixture_smoke.rb [--execute --app-execution-file FILE]"
  parser.on("--execute", "Build q4 Linux Runtime binaries and run the restricted D-Bus owner fixture in a q4 container.") { options[:execute] = true }
  parser.on("--no-sync-source", "Use the existing q4 source tree instead of syncing this checkout before building Linux Runtime binaries.") { options[:sync_source] = false }
  parser.on("--local-shell PATH", "Local shell used for ssh/rsync alias resolution.") { |value| options[:local_shell] = value }
  parser.on("--remote HOST", "Remote SSH target, default: #{DEFAULT_REMOTE_HOST}.") { |value| options[:remote_host] = value }
  parser.on("--remote-source-root PATH", "Remote source root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_source_root] = value }
  parser.on("--remote-build-root PATH", "Remote build cache root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_build_root] = value }
  parser.on("--remote-materials-root PATH", "Remote materials root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_materials_root] = value }
  parser.on("--tools-image IMAGE", "q4 Docker tools image tag with ruby, gcc, dbus, and gio-2.0 headers.") { |value| options[:tools_image] = value }
  parser.on("--tools-target TARGET", "Dockerfile target used for the q4 tools image, default: #{DEFAULT_TOOLS_TARGET}.") { |value| options[:tools_target] = value }
  parser.on("--tools-base-image IMAGE", "q4-local Docker base image used when building tools, default: #{DEFAULT_TOOLS_BASE_IMAGE}.") { |value| options[:tools_base_image] = value }
  parser.on("--container-goarch GOARCH", "Runtime GOARCH to match the q4 tools image, default: #{DEFAULT_CONTAINER_GOARCH}.") { |value| options[:container_goarch] = value }
  parser.on("--app APP_ID", "Known Windows app id, default: #{MESSAGEBOX_APP_ID}.") { |value| options[:app_id] = value }
  parser.on("--app-execution-file PATH", "Local verified-catalog app-execution evidence JSON under this checkout or /tmp/xnix-*.") { |value| options[:app_execution_file] = value }
  parser.on("--timeout-seconds SECONDS", Integer, "Timeout for each q4 operation.") { |value| options[:timeout_seconds] = value }
  parser.on("--output PATH", "Write the planned or passed summary JSON under this checkout or /tmp/xnix-*.") { |value| options[:output] = value }
end.parse!

abort "q4 D-Bus controlled launch owner fixture smoke does not accept positional arguments" unless ARGV.empty?

def ensure_remote_xnix_path!(label, path)
  clean = Pathname.new(path).cleanpath.to_s
  return clean if clean.start_with?("/home/xnix-") || clean.start_with?("/tmp/xnix-")

  abort "#{label} must stay under /home/xnix-* or /tmp/xnix-* on q4"
end

def ensure_local_xnix_path!(label, path)
  clean = Pathname.new(path).expand_path(PROJECT_ROOT).cleanpath
  return clean if clean.to_s.start_with?(PROJECT_ROOT.to_s)
  return clean if clean.to_s.start_with?("/tmp/xnix-")

  abort "#{label} must stay under this checkout or /tmp/xnix-*"
end

def ensure_single_line!(label, value)
  clean = value.to_s.strip
  abort "#{label} must be non-empty" if clean.empty?
  abort "#{label} must be single-line" if clean.include?("\n") || clean.include?("\r")

  clean
end

def ensure_tools_image!(image)
  clean = ensure_single_line!("tools image", image)
  abort "tools image must not request latest" if clean.end_with?(":latest")
  abort "tools image must be an xnix builder tools image" unless clean.start_with?("xnix-builder-tools:")

  clean
end

def ensure_tools_target!(target)
  clean = ensure_single_line!("tools target", target)
  abort "tools target must be dbus-tools or tools" unless %w[dbus-tools tools].include?(clean)

  clean
end

def ensure_tools_base_image!(image)
  clean = ensure_single_line!("tools base image", image)
  abort "tools base image must not request latest" if clean.end_with?(":latest")

  clean
end

def ensure_goarch!(goarch)
  clean = ensure_single_line!("container goarch", goarch)
  abort "container goarch must be amd64, arm64, or 386" unless %w[amd64 arm64 386].include?(clean)

  clean
end

def emit_json(payload, output_path)
  text = JSON.pretty_generate(payload) + "\n"
  File.write(output_path, text) if output_path
  puts text
end

def run_command(argv, timeout_seconds:, chdir: PROJECT_ROOT)
  stdout = +""
  stderr = +""
  status = nil
  timed_out = false
  Open3.popen3(*argv, chdir: chdir.to_s, pgroup: true) do |_stdin, out, err, wait_thread|
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
  stderr = [stderr, "command timed out after #{timeout_seconds}s"].reject(&:empty?).join("\n") if timed_out
  [stdout, stderr, timed_out ? 124 : status.exitstatus]
end

def run_json_command(argv, timeout_seconds:)
  stdout, stderr, status = run_command(argv, timeout_seconds: timeout_seconds)
  unless status.zero?
    warn stdout unless stdout.empty?
    warn stderr unless stderr.empty?
    abort "command failed: #{argv.first}"
  end
  JSON.parse(stdout)
end

def run_shell(shell, command, timeout_seconds:)
  run_command([shell, "-lc", command], timeout_seconds: timeout_seconds)
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

def docker_exec(container, *argv, user: nil, env: {}, workdir: nil)
  command = ["docker", "exec"]
  command += ["--user", user] if user
  command += ["--workdir", workdir] if workdir
  env.each { |key, value| command += ["--env", "#{key}=#{value}"] }
  command << container
  command.concat(argv)
  shell_join(command)
end

def require_passed_remote_fixture!(stdout)
  unless stdout.include?("PASS: D-Bus controlled launch owner fixture smoke")
    abort "q4 D-Bus fixture did not pass"
  end
  forbidden = [
    "docker.sock",
    "--privileged",
    "--network host",
    "type=bind"
  ]
  forbidden.each do |term|
    abort "q4 D-Bus fixture output exposed #{term}" if stdout.include?(term)
  end
end

remote_host = options.fetch(:remote_host)
remote_source_root = ensure_remote_xnix_path!("remote source root", options.fetch(:remote_source_root))
remote_build_root = ensure_remote_xnix_path!("remote build root", options.fetch(:remote_build_root))
remote_materials_root = ensure_remote_xnix_path!("remote materials root", options.fetch(:remote_materials_root))
tools_image = ensure_tools_image!(options.fetch(:tools_image))
tools_target = ensure_tools_target!(options.fetch(:tools_target))
tools_base_image = ensure_tools_base_image!(options.fetch(:tools_base_image))
container_goarch = ensure_goarch!(options.fetch(:container_goarch))
app_id = ensure_single_line!("app id", options.fetch(:app_id))
output_path = options.fetch(:output).empty? ? nil : ensure_local_xnix_path!("output path", options.fetch(:output))
app_execution_path = options.fetch(:app_execution_file).to_s.strip.empty? ? nil : ensure_local_xnix_path!("app execution file", options.fetch(:app_execution_file))
run_id = "#{Time.now.utc.strftime("%Y%m%d%H%M%S")}-#{Process.pid}-#{SecureRandom.hex(4)}"
remote_run_root = "#{remote_materials_root}/#{run_id}"
remote_evidence = "#{remote_run_root}/known-app-verified-catalog-app-execution.json"
setup_container_name = "xnix-q4-dbus-fixture-setup-#{VERSION.gsub(/[^A-Za-z0-9_.-]/, "-")}-#{run_id.gsub(/[^A-Za-z0-9_.-]/, "-")}"
run_container_name = "xnix-q4-dbus-fixture-run-#{VERSION.gsub(/[^A-Za-z0-9_.-]/, "-")}-#{run_id.gsub(/[^A-Za-z0-9_.-]/, "-")}"
prepared_image = "xnix-q4-dbus-fixture-runtime:#{VERSION.gsub(/[^A-Za-z0-9_.-]/, "-")}-#{run_id.gsub(/[^A-Za-z0-9_.-]/, "-")}"
remote_linux_bin_root = "#{remote_build_root}/bin/linux-#{container_goarch}"

plan = {
  "schema_version" => "xnix.scripts.q4_dbus_controlled_launch_owner_fixture_smoke.v1",
  "request_type" => "q4-dbus-controlled-launch-owner-fixture-smoke",
  "version" => VERSION,
  "status" => options.fetch(:execute) ? "running" : "planned",
  "execute" => options.fetch(:execute),
  "remote_host" => remote_host,
  "remote_source_root" => remote_source_root,
  "remote_build_root" => remote_build_root,
  "remote_materials_root" => remote_materials_root,
  "tools_image" => tools_image,
  "tools_target" => tools_target,
  "tools_base_image" => tools_base_image,
  "container_goarch" => container_goarch,
  "tools_image_build_planned" => true,
  "tools_image_built_on_q4" => false,
  "app_id" => app_id,
  "app_execution_evidence_required" => true,
  "app_execution_evidence_copied_to_q4" => false,
  "app_execution_evidence_path_exposed" => false,
  "linux_runtime_build_planned" => true,
  "linux_runtime_built_on_q4" => false,
  "dbus_adapter_compiled_in_q4_container" => false,
  "dbus_fixture_executed" => false,
  "dbus_fixture_passed" => false,
  "container_runtime" => "docker",
  "container_network_none" => true,
  "container_read_only" => true,
  "container_cap_drop_all" => true,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false,
  "host_networking_required" => false,
  "privileged_container_required" => false,
  "host_compilation_avoided" => true,
  "host_root_modified" => false
}

unless options.fetch(:execute)
  emit_json(plan, output_path)
  exit 0
end
abort "q4 D-Bus fixture execute requires --app-execution-file" unless app_execution_path
abort "app execution file must exist" unless app_execution_path.file?

build = run_json_command(
  [
    "ruby", REMOTE_GO_BUILD.to_s,
    "--execute",
    *(options.fetch(:sync_source) ? [] : ["--no-sync-source"]),
    "--package", "./cmd/xnix-runtime-go",
    "--package", "./cmd/xnix-runtime-owner",
    "--goos", "linux",
    "--goarch", container_goarch,
    "--remote", remote_host,
    "--remote-source-root", remote_source_root,
    "--remote-build-root", remote_build_root,
    "--remote-timeout-seconds", options.fetch(:timeout_seconds).to_s
  ],
  timeout_seconds: options.fetch(:timeout_seconds)
)
abort "q4 Linux Runtime build did not pass" unless build.fetch("status") == "passed" && build.fetch("host_compilation_avoided") == true

mkdir_stdout, mkdir_stderr, mkdir_status = run_shell(
  options.fetch(:local_shell),
  shell_join(ssh_command(remote_host, shell_join(["mkdir", "-p", remote_run_root]))),
  timeout_seconds: options.fetch(:timeout_seconds)
)
unless mkdir_status.zero?
  warn mkdir_stdout unless mkdir_stdout.empty?
  warn mkdir_stderr unless mkdir_stderr.empty?
  abort "q4 D-Bus fixture remote materials root preparation failed"
end

rsync_stdout, rsync_stderr, rsync_status = run_command(
  ["rsync", "-az", "--timeout", "60", "--contimeout", "15", app_execution_path.to_s, "#{remote_host}:#{remote_evidence}"],
  timeout_seconds: options.fetch(:timeout_seconds)
)
unless rsync_status.zero?
  warn rsync_stdout unless rsync_stdout.empty?
  warn rsync_stderr unless rsync_stderr.empty?
  abort "q4 D-Bus fixture evidence copy failed"
end

remote_commands = [
  "set -eu",
  "tools_image_built=0",
  "tools_image_arch=$(docker image inspect --format '{{.Architecture}}' #{Shellwords.escape(tools_image)} 2>/dev/null || true)",
  "if [ \"$tools_image_arch\" != #{Shellwords.escape(container_goarch)} ]; then docker build --pull=false --platform #{Shellwords.escape("linux/#{container_goarch}")} --build-arg #{Shellwords.escape("XNIX_TOOLS_BASE_IMAGE=#{tools_base_image}")} --target #{Shellwords.escape(tools_target)} --tag #{Shellwords.escape(tools_image)} --file #{Shellwords.escape("#{remote_source_root}/Dockerfile")} #{Shellwords.escape(remote_source_root)}; tools_image_built=1; fi",
  "docker run --rm --platform #{Shellwords.escape("linux/#{container_goarch}")} --network none --cap-drop ALL --security-opt no-new-privileges #{Shellwords.escape(tools_image)} /bin/sh -lc #{Shellwords.escape("command -v ruby >/dev/null && command -v gcc >/dev/null && pkg-config --cflags --libs gio-2.0 >/dev/null")}",
  "echo XNIX_Q4_TOOLS_IMAGE_BUILT=$tools_image_built",
  "docker rm -f #{Shellwords.escape(setup_container_name)} #{Shellwords.escape(run_container_name)} >/dev/null 2>&1 || true",
  "docker rmi #{Shellwords.escape(prepared_image)} >/dev/null 2>&1 || true",
  "docker create --platform #{Shellwords.escape("linux/#{container_goarch}")} --name #{Shellwords.escape(setup_container_name)} --network none --cap-drop ALL --security-opt no-new-privileges #{Shellwords.escape(tools_image)} sleep 600 >/dev/null",
  "docker start #{Shellwords.escape(setup_container_name)} >/dev/null",
  "docker cp #{Shellwords.escape("#{remote_source_root}/.")} #{Shellwords.escape(setup_container_name)}:/workspace",
  docker_exec(setup_container_name, "mkdir", "-p", "/workspace/q4-dbus-input", "/usr/local/bin", user: "root"),
  "docker cp #{Shellwords.escape("#{remote_linux_bin_root}/xnix-runtime-go")} #{Shellwords.escape(setup_container_name)}:/usr/local/bin/xnix-runtime-go",
  "docker cp #{Shellwords.escape("#{remote_linux_bin_root}/xnix-runtime-owner")} #{Shellwords.escape(setup_container_name)}:/usr/local/bin/xnix-runtime-owner",
  "docker cp #{Shellwords.escape(remote_evidence)} #{Shellwords.escape(setup_container_name)}:/workspace/q4-dbus-input/known-app-verified-catalog-app-execution.json",
  docker_exec(setup_container_name, "chmod", "755", "/usr/local/bin/xnix-runtime-go", "/usr/local/bin/xnix-runtime-owner", user: "root"),
  docker_exec(setup_container_name, "/bin/sh", "-lc", "gcc /workspace/runtime/dbus/xnix_compatd_smoke.c -o /usr/local/bin/xnix-dbus-smoke $(pkg-config --cflags --libs gio-2.0) && chmod 755 /usr/local/bin/xnix-dbus-smoke", user: "root", workdir: "/workspace"),
  "docker commit #{Shellwords.escape(setup_container_name)} #{Shellwords.escape(prepared_image)} >/dev/null",
  "docker rm -f #{Shellwords.escape(setup_container_name)} >/dev/null",
  "docker run --rm --platform #{Shellwords.escape("linux/#{container_goarch}")} --name #{Shellwords.escape(run_container_name)} --network none --read-only --cap-drop ALL --security-opt no-new-privileges --tmpfs /tmp:rw,noexec,nosuid,size=64m --tmpfs #{Shellwords.escape("#{SCRATCH_ROOT}:rw,exec,nosuid,size=67108864,mode=1777")} --workdir /workspace --env PATH=/usr/local/bin:/usr/bin:/bin #{Shellwords.escape(prepared_image)} " +
    shell_join([
    "ruby", "scripts/dbus_controlled_launch_owner_fixture_smoke.rb",
    "--app", app_id,
    "--state-root", "#{SCRATCH_ROOT}/state",
    "--cache-root", "#{SCRATCH_ROOT}/known-winapps",
    "--gui-smoke-evidence-file", "/workspace/q4-dbus-input/known-app-verified-catalog-app-execution.json",
    "--runtime-command", "/usr/local/bin/xnix-runtime-go"
  ])
]
remote_script = [
  "trap #{Shellwords.escape("docker rm -f #{setup_container_name} #{run_container_name} >/dev/null 2>&1 || true; docker rmi #{prepared_image} >/dev/null 2>&1 || true")} EXIT",
  *remote_commands
].join("\n")

stdout, stderr, status = run_shell(
  options.fetch(:local_shell),
  shell_join(ssh_command(remote_host, remote_script)),
  timeout_seconds: options.fetch(:timeout_seconds)
)
unless status.zero?
  warn stdout unless stdout.empty?
  warn stderr unless stderr.empty?
  abort "q4 D-Bus controlled launch owner fixture smoke failed"
end
require_passed_remote_fixture!(stdout)

summary = plan.merge(
  "status" => "passed",
  "tools_image_built_on_q4" => stdout.include?("XNIX_Q4_TOOLS_IMAGE_BUILT=1"),
  "app_execution_evidence_copied_to_q4" => true,
  "linux_runtime_built_on_q4" => true,
  "dbus_adapter_compiled_in_q4_container" => true,
  "dbus_fixture_executed" => true,
  "dbus_fixture_passed" => true,
  "host_compilation_avoided" => build.fetch("host_compilation_avoided"),
  "host_root_modified" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false,
  "host_networking_required" => false,
  "privileged_container_required" => false
)
emit_json(summary, output_path)

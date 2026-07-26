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
DEFAULT_REMOTE_SOURCE_ROOT = ENV.fetch("XNIX_REMOTE_SOURCE_ROOT", "/home/xnix-build/xnix-runtime-source-gui-#{DEFAULT_SOURCE_SYNC_MODE}-#{VERSION}")
DEFAULT_REMOTE_MATERIALS_ROOT = ENV.fetch("XNIX_REMOTE_MATERIALS_ROOT", "/home/xnix-run-materials")
DEFAULT_LOCAL_SHELL = ENV.fetch("XNIX_LOCAL_SHELL", "/bin/zsh")

options = {
  execute: false,
  sync_source: true,
  source_sync_mode: DEFAULT_SOURCE_SYNC_MODE,
  local_shell: DEFAULT_LOCAL_SHELL,
  remote_host: DEFAULT_REMOTE_HOST,
  remote_source_root: DEFAULT_REMOTE_SOURCE_ROOT,
  remote_materials_root: DEFAULT_REMOTE_MATERIALS_ROOT,
  remote_kernel: ENV.fetch("XNIX_WINE_GUI_REMOTE_KERNEL", "#{DEFAULT_REMOTE_MATERIALS_ROOT}/wine-guest/bzImage"),
  remote_ssh_key: ENV.fetch("XNIX_WINE_GUI_REMOTE_SSH_KEY", "#{DEFAULT_REMOTE_MATERIALS_ROOT}/ssh/id_ed25519"),
  remote_executable: ENV.fetch("XNIX_WINE_GUI_REMOTE_EXECUTABLE", ""),
  remote_file_argument: ENV.fetch("XNIX_WINE_GUI_REMOTE_FILE_ARGUMENT", ""),
  sample_file_argument: ENV.fetch("XNIX_WINE_GUI_REMOTE_SAMPLE_FILE_ARGUMENT", ""),
  remote_build_root: ENV.fetch("XNIX_REMOTE_BUILD_ROOT", "/home/xnix-build-cache"),
  remote_go: ENV.fetch("XNIX_REMOTE_GO", "/home/xnix-toolchains/go1.24.4-linux-amd64/bin/go"),
  launch_mode: ENV.fetch("XNIX_WINE_GUI_REMOTE_LAUNCH_MODE", "direct"),
  window_match: ENV.fetch("XNIX_WINE_GUI_REMOTE_WINDOW_MATCH", ""),
  report_output: ENV.fetch("XNIX_WINE_GUI_REMOTE_REPORT", "#{DEFAULT_REMOTE_MATERIALS_ROOT}/state/wine-gui-smoke-#{VERSION}.json"),
  evidence_output: ENV.fetch("XNIX_WINE_GUI_REMOTE_EVIDENCE", "#{DEFAULT_REMOTE_MATERIALS_ROOT}/state/wine-gui-evidence-#{VERSION}.json"),
  kde_page_output: ENV.fetch("XNIX_WINE_GUI_REMOTE_KDE_PAGE", "#{DEFAULT_REMOTE_MATERIALS_ROOT}/state/wine-gui-kde-page-#{VERSION}.json"),
  kde_action_output: ENV.fetch("XNIX_WINE_GUI_REMOTE_KDE_ACTION", "#{DEFAULT_REMOTE_MATERIALS_ROOT}/state/wine-gui-kde-action-#{VERSION}.json"),
  known_app_id: ENV.fetch("XNIX_WINE_GUI_REMOTE_KNOWN_APP_ID", ""),
  evidence_app_id: ENV.fetch("XNIX_WINE_GUI_REMOTE_EVIDENCE_APP_ID", "org.xnix.apps.mines"),
  evidence_display_name: ENV.fetch("XNIX_WINE_GUI_REMOTE_EVIDENCE_DISPLAY_NAME", "Mines"),
  evidence_app_version: ENV.fetch("XNIX_WINE_GUI_REMOTE_EVIDENCE_APP_VERSION", VERSION),
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
  parser.on("--source-sync-mode MODE", "Source sync mode: runtime or full.") { |value| options[:source_sync_mode] = value }
  parser.on("--local-shell PATH", "Local shell used for ssh/rsync alias resolution.") { |value| options[:local_shell] = value }
  parser.on("--remote HOST", "Remote SSH target.") { |value| options[:remote_host] = value }
  parser.on("--remote-source-root PATH", "Remote source root under /home/xnix*.") { |value| options[:remote_source_root] = value }
  parser.on("--remote-materials-root PATH", "Remote materials root under /home/xnix*.") { |value| options[:remote_materials_root] = value }
  parser.on("--remote-kernel PATH", "Remote Wine guest kernel image.") { |value| options[:remote_kernel] = value }
  parser.on("--remote-ssh-key PATH", "Remote Wine guest SSH key.") { |value| options[:remote_ssh_key] = value }
  parser.on("--remote-executable PATH", "Remote Windows GUI .exe under /home/xnix*.") { |value| options[:remote_executable] = value }
  parser.on("--remote-file-argument PATH", "Remote file under /home/xnix* copied into the Wine guest and passed to the Windows GUI app.") { |value| options[:remote_file_argument] = value }
  parser.on("--sample-file-argument NAME", "Create a remote sample file under the smoke state root and pass it to the Windows GUI app.") { |value| options[:sample_file_argument] = value }
  parser.on("--remote-build-root PATH", "Remote build cache root under /home/xnix*.") { |value| options[:remote_build_root] = value }
  parser.on("--remote-go PATH", "Remote Go binary used to build xnix-runtime-go.") { |value| options[:remote_go] = value }
  parser.on("--launch-mode MODE", "Launch mode: direct or owner-controlled-launch.") { |value| options[:launch_mode] = value }
  parser.on("--window-match TEXT", "Case-insensitive X window title/text required for GUI observation.") { |value| options[:window_match] = value }
  parser.on("--report-output PATH", "Remote JSON report output path under /home/xnix*.") { |value| options[:report_output] = value }
  parser.on("--evidence-output PATH", "Remote Runtime GUI evidence output path under /home/xnix*.") { |value| options[:evidence_output] = value }
  parser.on("--kde-page-output PATH", "Remote KDE center page JSON output path under /home/xnix*.") { |value| options[:kde_page_output] = value }
  parser.on("--kde-action-output PATH", "Remote KDE controlled-launch action JSON output path under /home/xnix*.") { |value| options[:kde_action_output] = value }
  parser.on("--known-app-id ID", "Known Windows GUI app id resolved by the remote Go Runtime.") { |value| options[:known_app_id] = value }
  parser.on("--evidence-app-id ID", "Application id for the Runtime GUI evidence projection.") { |value| options[:evidence_app_id] = value }
  parser.on("--evidence-display-name NAME", "Display name for the Runtime GUI evidence projection.") { |value| options[:evidence_display_name] = value }
  parser.on("--evidence-app-version VERSION", "Application version for the Runtime GUI evidence projection.") { |value| options[:evidence_app_version] = value }
  parser.on("--state-root PATH", "Remote smoke state root under /home/xnix*.") { |value| options[:state_root] = value }
  parser.on("--ssh-port PORT", "Loopback SSH port for the temporary QEMU guest.") { |value| options[:ssh_port] = value }
  parser.on("--display-number NUMBER", Integer, "Remote Xvfb display number.") { |value| options[:display_number] = value }
  parser.on("--wait-seconds SECONDS", Integer, "Seconds to wait for the GUI window.") { |value| options[:wait_seconds] = value }
  parser.on("--remote-timeout-seconds SECONDS", Integer, "Timeout for each remote shell operation.") { |value| options[:remote_timeout_seconds] = value }
end.parse!

abort "remote Wine guest GUI smoke does not accept positional arguments" unless ARGV.empty?
abort "launch mode must be direct or owner-controlled-launch" unless %w[direct owner-controlled-launch].include?(options.fetch(:launch_mode))

if !ENV.key?("XNIX_REMOTE_SOURCE_ROOT") && options.fetch(:remote_source_root) == DEFAULT_REMOTE_SOURCE_ROOT
  options[:remote_source_root] = "/home/xnix-build/xnix-runtime-source-gui-#{options.fetch(:source_sync_mode)}-#{VERSION}"
end

def ensure_remote_xnix_path!(label, path)
  clean = Pathname.new(path).cleanpath.to_s
  return clean if clean.start_with?("/home/xnix-") || clean.start_with?("/tmp/xnix-")

  abort "#{label} must stay under /home/xnix-* or /tmp/xnix-* on the remote build host"
end

def ensure_optional_remote_xnix_path!(label, path)
  clean = path.to_s.strip
  return "" if clean.empty?

  ensure_remote_xnix_path!(label, clean)
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

def remote_json_command(remote_source_root, output_path, argv)
  writer = "stdout, stderr, status = Open3.capture3(*ARGV[1..]); warn stderr unless stderr.empty?; exit 1 unless status.success?; payload = JSON.parse(stdout); File.write(ARGV[0], JSON.pretty_generate(payload) + \"\\n\")"
  [
    "set -eu",
    "cd #{Shellwords.escape(remote_source_root)}",
    shell_join(["ruby", "-rjson", "-ropen3", "-e", writer, output_path, *argv])
  ].join("\n")
end

remote_host = options.fetch(:remote_host)
source_sync_mode = options.fetch(:source_sync_mode)
source_entries = source_sync_entries(source_sync_mode)
remote_source_root = ensure_remote_xnix_path!("remote source root", options.fetch(:remote_source_root))
remote_materials_root = ensure_remote_xnix_path!("remote materials root", options.fetch(:remote_materials_root))
remote_kernel = ensure_remote_xnix_path!("remote kernel", options.fetch(:remote_kernel))
remote_ssh_key = ensure_remote_xnix_path!("remote SSH key", options.fetch(:remote_ssh_key))
remote_executable = ensure_optional_remote_xnix_path!("remote executable", options.fetch(:remote_executable))
remote_build_root = ensure_remote_xnix_path!("remote build root", options.fetch(:remote_build_root))
launch_mode = options.fetch(:launch_mode)
report_output = ensure_remote_xnix_path!("report output", options.fetch(:report_output))
evidence_output = ensure_remote_xnix_path!("evidence output", options.fetch(:evidence_output))
kde_page_output = ensure_remote_xnix_path!("KDE page output", options.fetch(:kde_page_output))
kde_action_output = ensure_remote_xnix_path!("KDE action output", options.fetch(:kde_action_output))
state_root = ensure_remote_xnix_path!("state root", options.fetch(:state_root))
sample_file_argument_value = options.fetch(:sample_file_argument).strip
abort "use either --remote-file-argument or --sample-file-argument, not both" if !options.fetch(:remote_file_argument).strip.empty? && !sample_file_argument_value.empty?
sample_file_argument_name = File.basename(sample_file_argument_value)
sample_file_argument_requested = !sample_file_argument_value.empty?
remote_file_argument = if sample_file_argument_requested
                         ensure_remote_xnix_path!("remote sample file argument", "#{state_root}/#{sample_file_argument_name}")
                       else
                         ensure_optional_remote_xnix_path!("remote file argument", options.fetch(:remote_file_argument))
                       end
remote_runtime_bin = "#{remote_build_root}/bin/xnix-runtime-go"
remote_owner_bin = "#{remote_build_root}/bin/xnix-runtime-owner"
remote_launcher_bin = "#{remote_build_root}/bin/xnix-compat-launch"
remote_known_app_cache_root = "#{remote_build_root}/known-winapps"
remote_go_dir = Pathname.new(options.fetch(:remote_go)).dirname.to_s

plan = {
  "schema_version" => "xnix.scripts.remote_wine_guest_gui_smoke.v1",
  "request_type" => "remote-wine-guest-gui-smoke",
  "version" => VERSION,
  "status" => options.fetch(:execute) ? "running" : "planned",
  "execute" => options.fetch(:execute),
  "remote_host" => remote_host,
  "source_sync_planned" => options.fetch(:sync_source),
  "source_sync_mode" => source_sync_mode,
  "source_sync_entry_count" => source_entries.length,
  "source_sync_entries" => source_entries,
  "remote_source_root" => remote_source_root,
  "remote_materials_root" => remote_materials_root,
  "remote_kernel" => remote_kernel,
  "remote_ssh_key" => remote_ssh_key,
  "remote_executable" => remote_executable,
  "remote_file_argument" => remote_file_argument,
  "sample_file_argument_requested" => sample_file_argument_requested,
  "file_argument_count" => remote_file_argument.empty? ? 0 : 1,
  "file_argument_delivery" => remote_file_argument.empty? ? "" : "remote-file-to-guest-copy-and-winepath",
  "window_match" => options.fetch(:window_match),
  "remote_build_root" => remote_build_root,
  "remote_runtime_bin" => remote_runtime_bin,
  "remote_owner_bin" => remote_owner_bin,
  "remote_launcher_bin" => remote_launcher_bin,
  "remote_known_app_cache_root" => remote_known_app_cache_root,
  "runtime_build_planned" => true,
  "owner_build_planned" => launch_mode == "owner-controlled-launch",
  "launcher_build_planned" => launch_mode == "owner-controlled-launch",
  "report_output" => report_output,
  "evidence_output" => evidence_output,
  "evidence_preview_planned" => true,
  "kde_page_output" => kde_page_output,
  "kde_action_output" => kde_action_output,
  "kde_center_page_preview_planned" => true,
  "kde_controlled_launch_action_preview_planned" => launch_mode == "owner-controlled-launch",
  "kde_action_state_root" => "#{state_root}/owner-controlled-launch-state",
  "known_app_id" => options.fetch(:known_app_id),
  "known_app_selection_planned" => !options.fetch(:known_app_id).strip.empty?,
  "evidence_app_id" => options.fetch(:evidence_app_id),
  "evidence_display_name" => options.fetch(:evidence_display_name),
  "evidence_app_version" => options.fetch(:evidence_app_version),
  "state_root" => state_root,
  "ssh_port" => options.fetch(:ssh_port),
  "display_number" => options.fetch(:display_number),
  "wait_seconds" => options.fetch(:wait_seconds),
  "remote_timeout_seconds" => options.fetch(:remote_timeout_seconds),
  "launch_mode" => launch_mode,
  "owner_controlled_launch_requested" => launch_mode == "owner-controlled-launch",
  "remote_command" => "ruby scripts/wine_guest_gui_smoke.rb --execute --launch-mode #{launch_mode}",
  "backend" => "qemu-guest-wine-x11",
  "gui_app_name" => remote_executable.empty? ? "winemine.exe" : File.basename(remote_executable),
  "remote_gui_executable_configured" => !remote_executable.empty?,
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
  shell_join(["mkdir", "-p", "#{remote_build_root}/bin", "#{remote_build_root}/go-build", "#{remote_build_root}/go-mod", "#{remote_build_root}/tmp", remote_known_app_cache_root]),
  "cd #{Shellwords.escape(remote_source_root)}",
  "PATH=#{Shellwords.escape(remote_go_dir)}:$PATH GOCACHE=#{Shellwords.escape("#{remote_build_root}/go-build")} GOMODCACHE=#{Shellwords.escape("#{remote_build_root}/go-mod")} GOTMPDIR=#{Shellwords.escape("#{remote_build_root}/tmp")} #{shell_join([options.fetch(:remote_go), "build", "-o", remote_runtime_bin, "./cmd/xnix-runtime-go"])}"
]
if launch_mode == "owner-controlled-launch"
  build_env = "PATH=#{Shellwords.escape(remote_go_dir)}:$PATH GOCACHE=#{Shellwords.escape("#{remote_build_root}/go-build")} GOMODCACHE=#{Shellwords.escape("#{remote_build_root}/go-mod")} GOTMPDIR=#{Shellwords.escape("#{remote_build_root}/tmp")}"
  build_command.push("#{build_env} #{shell_join([options.fetch(:remote_go), "build", "-o", remote_owner_bin, "./cmd/xnix-runtime-owner"])}")
  build_command.push("#{build_env} #{shell_join([options.fetch(:remote_go), "build", "-o", remote_launcher_bin, "./cmd/xnix-compat-launch"])}")
end
build_command = build_command.join("\n")
build_stdout, build_stderr, build_status = run_shell(options.fetch(:local_shell), shell_join(ssh_command(remote_host, build_command)), timeout_seconds: options.fetch(:remote_timeout_seconds))
unless build_status.zero?
  warn build_stdout unless build_stdout.empty?
  warn build_stderr unless build_stderr.empty?
  warn "FAIL: remote Wine guest GUI Runtime build failed"
  exit 1
end

if sample_file_argument_requested
  sample_prepare = [
    "set -eu",
    shell_join(["mkdir", "-p", Pathname.new(remote_file_argument).dirname.to_s]),
    shell_join(["ruby", "-e", "File.write(ARGV.fetch(0), \"Xnix remote Wine GUI file-open smoke\\n\")", remote_file_argument])
  ].join("\n")
  sample_stdout, sample_stderr, sample_status = run_shell(options.fetch(:local_shell), shell_join(ssh_command(remote_host, sample_prepare)), timeout_seconds: options.fetch(:remote_timeout_seconds))
  warn sample_stdout unless sample_stdout.empty?
  warn sample_stderr unless sample_stderr.empty?
  unless sample_status.zero?
    warn "FAIL: remote Wine guest GUI sample file preparation failed"
    exit 1
  end
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
  "--launch-mode", launch_mode,
  "--runtime-bin", remote_runtime_bin,
  "--evidence-app-id", options.fetch(:evidence_app_id),
  "--evidence-display-name", options.fetch(:evidence_display_name),
  "--evidence-app-version", options.fetch(:evidence_app_version),
  "--report-output", report_output
]
remote_args.push("--known-app-id", options.fetch(:known_app_id)) unless options.fetch(:known_app_id).strip.empty?
remote_args.push("--file-argument", remote_file_argument) unless remote_file_argument.empty?
remote_args.push("--window-match", options.fetch(:window_match)) unless options.fetch(:window_match).strip.empty?
if launch_mode == "owner-controlled-launch"
  remote_args.push("--owner-bin", remote_owner_bin)
  remote_args.push("--launcher-bin", remote_launcher_bin)
  remote_args.push("--known-app-cache-root", remote_known_app_cache_root)
end
remote_args.push("--executable", remote_executable) unless remote_executable.empty?
remote_command = [
  "set -eu",
  "cd #{Shellwords.escape(remote_source_root)}",
  shell_join(remote_args)
].join("\n")

stdout, stderr, status = run_shell(options.fetch(:local_shell), shell_join(ssh_command(remote_host, remote_command)), timeout_seconds: options.fetch(:remote_timeout_seconds))
warn stderr unless stderr.empty?
unless status.zero?
  puts stdout unless stdout.empty?
  exit 1
end

smoke_payload = begin
  JSON.parse(stdout)
rescue JSON::ParserError
  {}
end
smoke_status = smoke_payload.fetch("status", "")
if smoke_status != "passed"
  puts stdout unless stdout.empty?
  exit 0
end

evidence_args = [
  remote_runtime_bin,
  "gui-smoke-evidence-preview",
  "--gui-smoke-report", report_output,
  "--app-id", options.fetch(:evidence_app_id),
  "--display-name", options.fetch(:evidence_display_name),
  "--app-version", options.fetch(:evidence_app_version),
  "--output", evidence_output
]
evidence_command = [
  "set -eu",
  "cd #{Shellwords.escape(remote_source_root)}",
  shell_join(evidence_args)
].join("\n")
_evidence_stdout, evidence_stderr, evidence_status = run_shell(options.fetch(:local_shell), shell_join(ssh_command(remote_host, evidence_command)), timeout_seconds: options.fetch(:remote_timeout_seconds))
warn evidence_stderr unless evidence_stderr.empty?
exit 1 unless evidence_status.zero?

kde_page_args = [
  remote_runtime_bin,
  "kde-center-page-preview",
  "--registry", "#{remote_source_root}/runtime/recipes/registry.json",
  "--app", options.fetch(:evidence_app_id),
  "--decision", "approved",
  "--known-app-evidence-file", evidence_output
]
kde_page_stdout, kde_page_stderr, kde_page_status = run_shell(
  options.fetch(:local_shell),
  shell_join(ssh_command(remote_host, remote_json_command(remote_source_root, kde_page_output, kde_page_args))),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
warn kde_page_stdout unless kde_page_stdout.empty?
warn kde_page_stderr unless kde_page_stderr.empty?
exit 1 unless kde_page_status.zero?

kde_action_output_written = false
if launch_mode == "owner-controlled-launch"
  kde_action_args = [
    remote_runtime_bin,
    "kde-controlled-launch-action-preview",
    "--state-root", "#{state_root}/owner-controlled-launch-state",
    "--kde-center-page-file", kde_page_output,
    "--app", options.fetch(:evidence_app_id)
  ]
  kde_action_stdout, kde_action_stderr, kde_action_status = run_shell(
    options.fetch(:local_shell),
    shell_join(ssh_command(remote_host, remote_json_command(remote_source_root, kde_action_output, kde_action_args))),
    timeout_seconds: options.fetch(:remote_timeout_seconds)
  )
  warn kde_action_stdout unless kde_action_stdout.empty?
  warn kde_action_stderr unless kde_action_stderr.empty?
  exit 1 unless kde_action_status.zero?

  kde_action_output_written = true
end

summary_reader = <<~RUBY
  smoke = JSON.parse(File.read(ARGV.fetch(0)))
  evidence = JSON.parse(File.read(ARGV.fetch(1)))
  kde = JSON.parse(File.read(ARGV.fetch(2)))
  action_output = ARGV.fetch(3)
  action_written = ARGV.fetch(4) == "true" && File.exist?(action_output)
  summary = {
    "schema_version" => "xnix.scripts.remote_wine_guest_gui_smoke.execute_result.v1",
    "request_type" => "remote-wine-guest-gui-smoke",
    "version" => smoke.fetch("version"),
    "status" => "passed",
    "execute" => true,
    "remote_host" => ARGV.fetch(5),
    "remote_build_completed" => true,
    "remote_runtime_bin" => ARGV.fetch(6),
    "remote_owner_bin" => ARGV.fetch(7),
    "remote_launcher_bin" => ARGV.fetch(8),
    "report_output" => ARGV.fetch(0),
    "evidence_output" => ARGV.fetch(1),
    "kde_page_output" => ARGV.fetch(2),
    "kde_action_output" => action_output,
    "evidence_output_written" => File.exist?(ARGV.fetch(1)),
    "kde_page_output_written" => File.exist?(ARGV.fetch(2)),
    "kde_action_output_written" => action_written,
    "smoke_status" => smoke.fetch("status"),
    "backend" => smoke.fetch("backend"),
    "launch_mode" => smoke.fetch("launch_mode"),
    "known_app_id" => smoke.fetch("known_app_id"),
    "known_app_name" => smoke.fetch("known_app_name", ""),
    "known_app_version" => smoke.fetch("known_app_version", ""),
    "gui_app_name" => smoke.fetch("gui_app_name"),
    "runtime_go_owned_gui_smoke" => smoke.fetch("runtime_go_owned_gui_smoke"),
    "qemu_started" => smoke.fetch("qemu_started"),
    "guest_ssh_ready" => smoke.fetch("guest_ssh_ready"),
    "wineboot_invoked" => smoke.fetch("wineboot_invoked"),
    "guest_x11_driver_available" => smoke.fetch("guest_x11_driver_available"),
    "file_argument_count" => smoke.fetch("file_argument_count", 0),
    "file_argument_copied_count" => smoke.fetch("file_argument_copied_count", 0),
    "file_arguments_passed" => smoke.fetch("file_arguments_passed", false),
    "file_argument_winepath_translated" => smoke.fetch("file_argument_winepath_translated", false),
    "file_argument_winepath_translated_count" => smoke.fetch("file_argument_winepath_translated_count", 0),
    "raw_file_argument_path_exposed" => smoke.fetch("raw_file_argument_path_exposed", false),
    "window_match" => smoke.fetch("window_match", ""),
    "window_match_observed" => smoke.fetch("window_match_observed", false),
    "window_evidence_summary" => smoke.fetch("window_evidence_summary", ""),
    "x_window_observed" => smoke.fetch("x_window_observed"),
    "x_window_child_count" => smoke.fetch("x_window_child_count"),
    "x_window_observation_attempts" => smoke.fetch("x_window_observation_attempts"),
    "runtime_evidence_report_consumed" => evidence.fetch("report_consumed"),
    "runtime_evidence_app_id" => evidence.fetch("app_id"),
    "runtime_evidence_window_observed" => evidence.fetch("x_window_observed"),
    "kde_page_app_id" => kde.fetch("application_id"),
    "kde_page_known_app_gui_evidence_count" => kde.fetch("known_app_gui_evidence_count"),
    "kde_page_backend_details_exposed" => kde.fetch("backend_details_exposed"),
    "owner_controlled_launch_requested" => smoke.fetch("owner_controlled_launch_requested"),
    "owner_evidence_handoff_ready" => smoke.fetch("owner_evidence_handoff_ready"),
    "host_root_modified" => smoke.fetch("host_root_modified"),
    "privileged_container_required" => smoke.fetch("privileged_container_required"),
    "host_networking_required" => smoke.fetch("host_networking_required"),
    "docker_socket_mounted" => smoke.fetch("docker_socket_mounted"),
    "broad_host_mount_required" => smoke.fetch("broad_host_mount_required")
  }
  puts JSON.pretty_generate(summary)
RUBY
summary_args = [
  "ruby",
  "-rjson",
  "-e",
  summary_reader,
  report_output,
  evidence_output,
  kde_page_output,
  kde_action_output,
  kde_action_output_written.to_s,
  remote_host,
  remote_runtime_bin,
  remote_owner_bin,
  remote_launcher_bin
]
summary_stdout, summary_stderr, summary_status = run_shell(
  options.fetch(:local_shell),
  shell_join(ssh_command(remote_host, ["set -eu", shell_join(summary_args)].join("\n"))),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
warn summary_stderr unless summary_stderr.empty?
exit 1 unless summary_status.zero?

puts summary_stdout

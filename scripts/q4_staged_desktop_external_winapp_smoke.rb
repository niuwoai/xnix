#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "optparse"
require "pathname"
require "shellwords"
require "timeout"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip
REMOTE_GO_BUILD = PROJECT_ROOT.join("scripts/remote_go_build.rb")

SCHEMA_VERSION = "xnix.scripts.q4_staged_desktop_external_winapp_smoke.v1"
REQUEST_TYPE = "q4-staged-desktop-external-winapp-smoke"

DEFAULT_REMOTE_HOST = ENV.fetch("XNIX_REMOTE_HOST", "root@q4")
DEFAULT_REMOTE_SOURCE_ROOT = ENV.fetch("XNIX_Q4_STAGED_EXTERNAL_SOURCE_ROOT", "/home/xnix-build/xnix-q4-staged-external-winapp-#{VERSION}")
DEFAULT_REMOTE_BUILD_ROOT = ENV.fetch("XNIX_REMOTE_BUILD_ROOT", "/home/xnix-build-cache")
DEFAULT_REMOTE_RUN_ROOT = ENV.fetch("XNIX_Q4_STAGED_EXTERNAL_RUN_ROOT", "/tmp/xnix-q4-staged-external-winapp-#{VERSION}")
DEFAULT_LOCAL_OUTPUT_ROOT = PROJECT_ROOT.join("output", "q4-staged-desktop-external-winapp-smoke-#{VERSION}")
DEFAULT_REMOTE_TIMEOUT_SECONDS = Integer(ENV.fetch("XNIX_Q4_STAGED_EXTERNAL_TIMEOUT_SECONDS", "900"), 10)
DEFAULT_IMAGE = ENV.fetch("XNIX_WINE_IMAGE", "xnix-wine-smoke:local")

options = {
  execute: false,
  local_shell: ENV.fetch("XNIX_LOCAL_SHELL", "/bin/zsh"),
  remote_host: DEFAULT_REMOTE_HOST,
  remote_source_root: DEFAULT_REMOTE_SOURCE_ROOT,
  remote_build_root: DEFAULT_REMOTE_BUILD_ROOT,
  remote_run_root: DEFAULT_REMOTE_RUN_ROOT,
  output: ENV.fetch("XNIX_Q4_STAGED_EXTERNAL_OUTPUT", DEFAULT_LOCAL_OUTPUT_ROOT.join("q4-staged-desktop-external-winapp-smoke.json").to_s),
  markdown_output: ENV.fetch("XNIX_Q4_STAGED_EXTERNAL_MARKDOWN_OUTPUT", DEFAULT_LOCAL_OUTPUT_ROOT.join("q4-staged-desktop-external-winapp-smoke.md").to_s),
  remote_timeout_seconds: DEFAULT_REMOTE_TIMEOUT_SECONDS,
  fixture: ENV.fetch("XNIX_Q4_STAGED_EXTERNAL_FIXTURE", "notepad-file-argument"),
  image: DEFAULT_IMAGE,
  timeout: ENV.fetch("XNIX_Q4_STAGED_EXTERNAL_GUI_TIMEOUT", "120s")
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/q4_staged_desktop_external_winapp_smoke.rb [--execute]"
  parser.on("--execute", "Build Runtime binaries on q4 and run the staged external Windows app desktop smoke on q4.") { options[:execute] = true }
  parser.on("--local-shell PATH", "Local shell used for ssh/rsync alias resolution.") { |value| options[:local_shell] = value }
  parser.on("--remote HOST", "Remote SSH target, default: #{DEFAULT_REMOTE_HOST}.") { |value| options[:remote_host] = value }
  parser.on("--remote-source-root PATH", "Remote source root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_source_root] = value }
  parser.on("--remote-build-root PATH", "Remote build/cache root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_build_root] = value }
  parser.on("--remote-run-root PATH", "Remote smoke run root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_run_root] = value }
  parser.on("--fixture NAME", "Delegated fixture, default: notepad-file-argument.") { |value| options[:fixture] = value }
  parser.on("--image IMAGE", "q4-local Wine GUI smoke image, default: #{DEFAULT_IMAGE}.") { |value| options[:image] = value }
  parser.on("--timeout DURATION", "Delegated GUI timeout, default: #{options.fetch(:timeout)}.") { |value| options[:timeout] = value }
  parser.on("--remote-timeout-seconds SECONDS", Integer, "Timeout for q4 build, remote smoke, and report fetch.") { |value| options[:remote_timeout_seconds] = value }
  parser.on("--output PATH", "Local JSON report output under this checkout or /tmp/xnix-*.") { |value| options[:output] = value }
  parser.on("--markdown-output PATH", "Local Markdown report output under this checkout or /tmp/xnix-*.") { |value| options[:markdown_output] = value }
end.parse!

abort "q4 staged desktop external Windows app smoke does not accept positional arguments" unless ARGV.empty?

def ensure_remote_xnix_path!(label, path)
  clean = Pathname.new(path).cleanpath.to_s
  return clean if clean.start_with?("/home/xnix-") || clean.start_with?("/tmp/xnix-")

  abort "#{label} must stay under /home/xnix-* or /tmp/xnix-* on q4"
end

def ensure_local_output_path!(label, path)
  clean = Pathname.new(path).expand_path(PROJECT_ROOT).cleanpath
  return clean if clean.to_s.start_with?(PROJECT_ROOT.to_s)
  return clean if clean.to_s.start_with?("/tmp/xnix-")

  abort "#{label} must stay under this checkout or /tmp/xnix-*"
end

def emit_json(payload, output_path)
  text = JSON.pretty_generate(payload) + "\n"
  FileUtils.mkdir_p(output_path.dirname) if output_path
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

def fetch_remote_text(remote_host, remote_path, timeout_seconds:)
  reader = shell_join(["ruby", "-e", "print File.read(ARGV.fetch(0))", remote_path])
  stdout, stderr, status = run_command(ssh_command(remote_host, reader), timeout_seconds: timeout_seconds)
  unless status.zero?
    warn stderr unless stderr.empty?
    abort "failed to fetch remote report #{remote_path}"
  end
  stdout
end

def fetch_remote_artifact(remote_host, remote_path, local_path, timeout_seconds:)
  text = fetch_remote_text(remote_host, remote_path, timeout_seconds: timeout_seconds)
  JSON.parse(text)
  FileUtils.mkdir_p(local_path.dirname)
  File.write(local_path, text.end_with?("\n") ? text : "#{text}\n")
  true
end

def bool(payload, key)
  payload[key] == true
end

remote_host = options.fetch(:remote_host)
remote_source_root = ensure_remote_xnix_path!("remote source root", options.fetch(:remote_source_root))
remote_build_root = ensure_remote_xnix_path!("remote build root", options.fetch(:remote_build_root))
remote_run_root = ensure_remote_xnix_path!("remote run root", options.fetch(:remote_run_root))
output_path = ensure_local_output_path!("output path", options.fetch(:output))
markdown_output_path = ensure_local_output_path!("markdown output path", options.fetch(:markdown_output))
artifact_output_root = output_path.dirname.join("artifacts")
remote_binary_root = "#{remote_build_root}/bin/linux-amd64"
remote_runtime_bin = "#{remote_binary_root}/xnix-runtime-go"
remote_launcher_bin = "#{remote_binary_root}/xnix-compat-launch"
remote_report = "#{remote_run_root}/q4-staged-desktop-external-winapp-smoke.json"
remote_markdown = "#{remote_run_root}/q4-staged-desktop-external-winapp-smoke.md"
remote_compatibility_bundle = "#{remote_run_root}/compatibility-evidence-bundle.json"
remote_application_detail = "#{remote_run_root}/external-winapp-application-detail.json"
remote_kde_page_from_detail = "#{remote_run_root}/kde-page-from-application-detail.json"
artifact_specs = [
  ["delegated_launcher_payload", "delegated_launcher_payload_path", "delegated-launcher-payload.json"],
  ["activation_status", "activation_status_path", "activation-status.json"],
  ["desktop_launch_packet", "desktop_launch_packet_path", "desktop-launch-packet.json"],
  ["one_shot", "one_shot_path", "one-shot-import-stage-launch.json"],
  ["one_shot_launch_packet", "one_shot_launch_packet_path", "one-shot-desktop-launch-packet.json"],
  ["runtime_packet", "runtime_packet_path", "runtime-gui-evidence-packet.json"],
  ["kde_page", "kde_page_path", "kde-external-app-page.json"],
  ["compatibility_evidence_bundle", "compatibility_evidence_bundle_path", "compatibility-evidence-bundle.json"],
  ["application_detail", "application_detail_path", "external-winapp-application-detail.json"],
  ["kde_page_from_application_detail", "kde_page_from_application_detail_path", "kde-page-from-application-detail.json"]
]
planned_artifact_outputs = artifact_specs.to_h do |name, _remote_key, local_name|
  ["#{name}_artifact_output_path", artifact_output_root.join(local_name).to_s]
end

delegated_command = [
  "ruby",
  "scripts/staged_desktop_external_winapp_smoke.rb",
  "--fixture", options.fetch(:fixture),
  "--run-root", "#{remote_run_root}/run",
  "--report-output", remote_report,
  "--markdown-output", remote_markdown,
  "--image", options.fetch(:image),
  "--timeout", options.fetch(:timeout)
]

plan = {
  "schema_version" => SCHEMA_VERSION,
  "request_type" => REQUEST_TYPE,
  "version" => VERSION,
  "status" => options.fetch(:execute) ? "running" : "planned",
  "execute" => options.fetch(:execute),
  "remote_host" => remote_host,
  "remote_source_root" => remote_source_root,
  "remote_build_root" => remote_build_root,
  "remote_run_root" => remote_run_root,
  "remote_runtime_binary" => remote_runtime_bin,
  "remote_launcher_binary" => remote_launcher_bin,
  "remote_paths_exposed_only_for_operator" => true,
  "delegated_script" => "scripts/staged_desktop_external_winapp_smoke.rb",
  "delegated_command" => delegated_command,
  "fixture" => options.fetch(:fixture),
  "image" => options.fetch(:image),
  "output_path" => output_path.to_s,
  "markdown_output_path" => markdown_output_path.to_s,
  "artifact_output_root" => artifact_output_root.to_s,
  "artifact_output_count" => artifact_specs.length,
  "artifact_outputs" => planned_artifact_outputs,
  "q4_compile_required" => true,
  "host_compilation_avoided" => true,
  "targeted_smoke_required" => true,
  "full_smoke_required" => false,
  "desktop_shell" => "KDE Plasma",
  "runtime_owned" => true,
  "go_runtime_backed" => true,
  "kde_policy_owner" => false,
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

build = run_json_command(
  [
    "ruby",
    REMOTE_GO_BUILD.to_s,
    "--execute",
    "--package", "./cmd/xnix-runtime-go",
    "--package", "./cmd/xnix-compat-launch",
    "--remote", remote_host,
    "--remote-source-root", remote_source_root,
    "--remote-build-root", remote_build_root,
    "--remote-timeout-seconds", options.fetch(:remote_timeout_seconds).to_s
  ],
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless build.fetch("status") == "passed" && bool(build, "remote_build_completed") && bool(build, "host_compilation_avoided")
  abort "q4 remote Go build did not pass"
end

cleanup_stdout, cleanup_stderr, cleanup_status = run_command(
  ssh_command(remote_host, shell_join(["rm", "-rf", remote_run_root]) + " && " + shell_join(["mkdir", "-p", remote_run_root])),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless cleanup_status.zero?
  warn cleanup_stdout unless cleanup_stdout.empty?
  warn cleanup_stderr unless cleanup_stderr.empty?
  abort "q4 staged desktop external Windows app smoke could not clean its remote run root"
end

remote_env = {
  "XNIX_RUNTIME_GO_BIN" => remote_runtime_bin,
  "XNIX_COMPAT_LAUNCH_BIN" => remote_launcher_bin,
  "XNIX_ALLOW_LOCAL_GO_COMPILE" => "0",
  "XNIX_WINE_IMAGE" => options.fetch(:image)
}
remote_command = [
  "cd", remote_source_root,
  "&&",
  *remote_env.map { |key, value| "#{key}=#{Shellwords.escape(value)}" },
  *delegated_command.map { |value| Shellwords.escape(value) }
].join(" ")

stdout, stderr, status = run_command(
  ssh_command(remote_host, remote_command),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless status.zero?
  warn stdout unless stdout.empty?
  warn stderr unless stderr.empty?
  abort "q4 staged desktop external Windows app smoke failed"
end

remote_payload_text = fetch_remote_text(remote_host, remote_report, timeout_seconds: options.fetch(:remote_timeout_seconds))
remote_markdown_text = fetch_remote_text(remote_host, remote_markdown, timeout_seconds: options.fetch(:remote_timeout_seconds))
delegated = JSON.parse(remote_payload_text)
unless delegated.fetch("status") == "passed" &&
       bool(delegated, "window_observed") &&
       bool(delegated, "desktop_exec_uses_external_app_handle") &&
       bool(delegated, "external_app_handle_consumed") &&
       bool(delegated, "external_file_bridge_ready") &&
       delegated.fetch("one_shot_status") == "passed" &&
       bool(delegated, "one_shot_staged_launcher_invoked") &&
       bool(delegated, "one_shot_runtime_launch_executed") &&
       bool(delegated, "one_shot_window_observed") &&
       bool(delegated, "desktop_launch_packet_safe_for_kde")
  warn remote_payload_text
  abort "q4 staged desktop external Windows app smoke did not prove the staged desktop launch path"
end

FileUtils.mkdir_p(markdown_output_path.dirname)
File.write(markdown_output_path, remote_markdown_text)

bundle_command = [
  remote_runtime_bin,
  "external-winapp-compatibility-evidence-bundle-preview",
  "--one-shot-result", delegated.fetch("one_shot_path"),
  "--desktop-launch-packet", delegated.fetch("desktop_launch_packet_path"),
  "--runtime-gui-evidence-packet", delegated.fetch("runtime_packet_path"),
  "--kde-page", delegated.fetch("kde_page_path"),
  "--output", remote_compatibility_bundle
]
bundle_stdout, bundle_stderr, bundle_status = run_command(
  ssh_command(remote_host, shell_join(bundle_command)),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless bundle_status.zero?
  warn bundle_stdout unless bundle_stdout.empty?
  warn bundle_stderr unless bundle_stderr.empty?
  abort "q4 external Windows app compatibility evidence bundle generation failed"
end
compatibility_bundle = JSON.parse(bundle_stdout)
unless compatibility_bundle.fetch("request_type") == "external-winapp-compatibility-evidence-bundle-preview" &&
       compatibility_bundle.fetch("safe_for_kde") == true &&
       compatibility_bundle.fetch("real_windows_app_run_verified") == true &&
       compatibility_bundle.fetch("kde_external_app_page_verified") == true
  warn bundle_stdout
  abort "q4 external Windows app compatibility evidence bundle did not prove the real app path"
end
delegated["compatibility_evidence_bundle_path"] = remote_compatibility_bundle

application_detail_command = [
  remote_runtime_bin,
  "external-winapp-application-detail-preview",
  "--compatibility-evidence-bundle", remote_compatibility_bundle,
  "--output", remote_application_detail
]
application_detail_stdout, application_detail_stderr, application_detail_status = run_command(
  ssh_command(remote_host, shell_join(application_detail_command)),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless application_detail_status.zero?
  warn application_detail_stdout unless application_detail_stdout.empty?
  warn application_detail_stderr unless application_detail_stderr.empty?
  abort "q4 external Windows app application detail generation failed"
end
application_detail = JSON.parse(application_detail_stdout)
unless application_detail.fetch("request_type") == "external-winapp-application-detail-preview" &&
       application_detail.fetch("safe_for_kde") == true &&
       application_detail.fetch("safe_for_ai_diagnostics") == true &&
       application_detail.fetch("real_windows_app_run_verified") == true &&
       application_detail.fetch("file_open_verified") == true &&
       application_detail.fetch("runtime_owned") == true &&
       application_detail.fetch("go_runtime_backed") == true &&
       application_detail.fetch("kde_policy_owner") == false
  warn application_detail_stdout
  abort "q4 external Windows app application detail did not prove the real app detail path"
end
delegated["application_detail_path"] = remote_application_detail

kde_page_from_detail_command = [
  remote_runtime_bin,
  "kde-center-page-preview",
  "--external-app-application-detail", remote_application_detail,
  "--decision", "approved",
  "--output", remote_kde_page_from_detail
]
kde_page_from_detail_stdout, kde_page_from_detail_stderr, kde_page_from_detail_status = run_command(
  ssh_command(remote_host, "cd #{Shellwords.escape(remote_source_root)} && #{shell_join(kde_page_from_detail_command)}"),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless kde_page_from_detail_status.zero?
  warn kde_page_from_detail_stdout unless kde_page_from_detail_stdout.empty?
  warn kde_page_from_detail_stderr unless kde_page_from_detail_stderr.empty?
  abort "q4 KDE page from external Windows app application detail generation failed"
end
kde_page_from_detail = JSON.parse(kde_page_from_detail_stdout)
kde_page_from_detail_header = kde_page_from_detail.fetch("header")
kde_page_from_detail_application_summary = kde_page_from_detail.fetch("application_summary")
unless kde_page_from_detail.fetch("request_type") == "kde-center-page-preview" &&
       kde_page_from_detail.fetch("external_winapp_application_detail_consumed") == true &&
       kde_page_from_detail.fetch("external_winapp_application_detail_count") == 1 &&
       kde_page_from_detail.fetch("safe_for_ai_diagnostics") == true &&
       kde_page_from_detail.fetch("runtime_owned") == true &&
       kde_page_from_detail.fetch("go_runtime_backed") == true &&
       kde_page_from_detail.fetch("kde_policy_owner") == false &&
       kde_page_from_detail.fetch("launch_enabled") == false &&
       kde_page_from_detail.fetch("backend_details_exposed") == false &&
       kde_page_from_detail_header.fetch("badge") == application_detail.fetch("compatibility_label") &&
       kde_page_from_detail_header.fetch("badge_tone") == "success" &&
       kde_page_from_detail_header.fetch("backend_details_exposed") == false &&
       kde_page_from_detail_application_summary.fetch("compatibility_state") == application_detail.fetch("compatibility_state") &&
       kde_page_from_detail_application_summary.fetch("compatibility_label") == application_detail.fetch("compatibility_label") &&
       kde_page_from_detail_application_summary.fetch("diagnostics_state") == "real-app-run-verified" &&
       kde_page_from_detail_application_summary.fetch("backend_launch_enabled") == false &&
       kde_page_from_detail_application_summary.fetch("action_execution_enabled") == false &&
       kde_page_from_detail_application_summary.fetch("settings_persistence_enabled") == false &&
       kde_page_from_detail_application_summary.fetch("backend_details_exposed") == false &&
       kde_page_from_detail_application_summary.fetch("host_root_modified") == false
  warn kde_page_from_detail_stdout
  abort "q4 KDE page from external Windows app application detail did not prove desktop consumption"
end
delegated["kde_page_from_application_detail_path"] = remote_kde_page_from_detail

fetched_artifacts = {}
artifact_specs.each do |name, remote_key, local_name|
  remote_path = delegated.fetch(remote_key)
  local_path = artifact_output_root.join(local_name)
  fetched_artifacts["#{name}_artifact_fetched"] = fetch_remote_artifact(
    remote_host,
    remote_path,
    local_path,
    timeout_seconds: options.fetch(:remote_timeout_seconds)
  )
  fetched_artifacts["#{name}_artifact_output_path"] = local_path.to_s
end

result = plan.merge(
  "status" => "passed",
  "remote_go_build_status" => build.fetch("status"),
  "remote_build_completed" => bool(build, "remote_build_completed"),
  "remote_run_root_cleaned" => true,
  "built_binary_count" => build.fetch("built_binary_count"),
  "delegated_schema_version" => delegated.fetch("schema_version"),
  "delegated_status" => delegated.fetch("status"),
  "app_id" => delegated.fetch("app_id"),
  "display_name" => delegated.fetch("display_name"),
  "desktop_exec_uses_external_app_handle" => delegated.fetch("desktop_exec_uses_external_app_handle"),
  "external_app_desktop_handle_ready" => delegated.fetch("external_app_desktop_handle_ready"),
  "activation_receipt_external_app_desktop_handle_ready" => delegated.fetch("activation_receipt_external_app_desktop_handle_ready"),
  "desktop_exec_invocation_exact" => delegated.fetch("desktop_exec_invocation_exact"),
  "launcher_context_from_environment" => delegated.fetch("launcher_context_from_environment"),
  "launcher_extra_arguments_appended" => delegated.fetch("launcher_extra_arguments_appended"),
  "external_desktop_argument_count" => delegated.fetch("external_desktop_argument_count"),
  "external_file_uri_arguments_accepted" => delegated.fetch("external_file_uri_arguments_accepted"),
  "external_file_open_requested" => delegated.fetch("external_file_open_requested"),
  "external_file_bridge_copy_enabled" => delegated.fetch("external_file_bridge_copy_enabled"),
  "external_file_bridge_copied_count" => delegated.fetch("external_file_bridge_copied_count"),
  "external_file_bridge_arguments_passed" => delegated.fetch("external_file_bridge_arguments_passed"),
  "external_file_bridge_argument_observed_count" => delegated.fetch("external_file_bridge_argument_observed_count"),
  "external_file_bridge_winepath_translated" => delegated.fetch("external_file_bridge_winepath_translated"),
  "external_file_bridge_winepath_translated_count" => delegated.fetch("external_file_bridge_winepath_translated_count"),
  "external_file_bridge_ready" => delegated.fetch("external_file_bridge_ready"),
  "windows_process_file_argument_window_observed" => delegated.fetch("windows_process_file_argument_window_observed"),
  "external_file_bridge_mount_enabled" => delegated.fetch("external_file_bridge_mount_enabled"),
  "raw_file_uri_arguments_exposed" => delegated.fetch("raw_file_uri_arguments_exposed"),
  "desktop_launch_packet_ready" => delegated.fetch("desktop_launch_packet_ready"),
  "desktop_launch_packet_safe_for_kde" => delegated.fetch("desktop_launch_packet_safe_for_kde"),
  "desktop_launch_packet_external_app_handle_consumed" => delegated.fetch("desktop_launch_packet_external_app_handle_consumed"),
  "desktop_launch_packet_window_observed" => delegated.fetch("desktop_launch_packet_window_observed"),
  "desktop_launch_packet_x_window_observed" => delegated.fetch("desktop_launch_packet_x_window_observed"),
  "one_shot_output_written" => delegated.fetch("one_shot_output_written"),
  "one_shot_status" => delegated.fetch("one_shot_status"),
  "one_shot_import_recorded" => delegated.fetch("one_shot_import_recorded"),
  "one_shot_desktop_activation_staged" => delegated.fetch("one_shot_desktop_activation_staged"),
  "one_shot_staged_launcher_invoked" => delegated.fetch("one_shot_staged_launcher_invoked"),
  "one_shot_staged_launcher_from_activation_root" => delegated.fetch("one_shot_staged_launcher_from_activation_root"),
  "one_shot_managed_launcher_executable_staged" => delegated.fetch("one_shot_managed_launcher_executable_staged"),
  "one_shot_desktop_exec_uses_external_app_handle" => delegated.fetch("one_shot_desktop_exec_uses_external_app_handle"),
  "one_shot_external_app_desktop_handle_ready" => delegated.fetch("one_shot_external_app_desktop_handle_ready"),
  "one_shot_desktop_launch_packet_written" => delegated.fetch("one_shot_desktop_launch_packet_written"),
  "one_shot_external_app_handle_consumed" => delegated.fetch("one_shot_external_app_handle_consumed"),
  "one_shot_external_file_bridge_ready" => delegated.fetch("one_shot_external_file_bridge_ready"),
  "one_shot_imported_artifact_digest_verified" => delegated.fetch("one_shot_imported_artifact_digest_verified"),
  "one_shot_runtime_launch_executed" => delegated.fetch("one_shot_runtime_launch_executed"),
  "one_shot_window_observed" => delegated.fetch("one_shot_window_observed"),
  "one_shot_x_window_observed" => delegated.fetch("one_shot_x_window_observed"),
  "one_shot_host_root_modified" => delegated.fetch("one_shot_host_root_modified"),
  "one_shot_docker_socket_mounted" => delegated.fetch("one_shot_docker_socket_mounted"),
  "one_shot_raw_paths_exposed" => delegated.fetch("one_shot_raw_paths_exposed"),
  "external_app_import_record_consumed" => delegated.fetch("external_app_import_record_consumed"),
  "external_app_handle_consumed" => delegated.fetch("external_app_handle_consumed"),
  "imported_artifact_digest_verified" => delegated.fetch("imported_artifact_digest_verified"),
  "window_observed" => delegated.fetch("window_observed"),
  "x_window_observed" => delegated.fetch("x_window_observed"),
  "container_network_mode" => delegated.fetch("container_network_mode"),
  "container_host_mount_count" => delegated.fetch("container_host_mount_count"),
  "runtime_packet_consumed_report" => delegated.fetch("runtime_packet_consumed_report"),
  "runtime_packet_external_app_run_record_consumed" => delegated.fetch("runtime_packet_external_app_run_record_consumed"),
  "runtime_packet_external_app_handle_consumed" => delegated.fetch("runtime_packet_external_app_handle_consumed"),
  "runtime_packet_known_app_gui_evidence_verified_count" => delegated.fetch("runtime_packet_known_app_gui_evidence_verified_count"),
  "kde_page_known_app_gui_evidence_count" => delegated.fetch("kde_page_known_app_gui_evidence_count"),
  "kde_page_card_external_app_run_record_consumed" => delegated.fetch("kde_page_card_external_app_run_record_consumed"),
  "kde_page_card_external_app_handle_consumed" => delegated.fetch("kde_page_card_external_app_handle_consumed"),
  "kde_page_card_window_observed" => delegated.fetch("kde_page_card_window_observed"),
  "kde_page_card_x_window_observed" => delegated.fetch("kde_page_card_x_window_observed"),
  "remote_report_fetched" => true,
  "remote_markdown_fetched" => true,
  "compatibility_evidence_bundle_generated" => true,
  "compatibility_evidence_bundle_request_type" => compatibility_bundle.fetch("request_type"),
  "compatibility_evidence_bundle_safe_for_kde" => compatibility_bundle.fetch("safe_for_kde"),
  "compatibility_evidence_bundle_real_windows_app_run_verified" => compatibility_bundle.fetch("real_windows_app_run_verified"),
  "compatibility_evidence_bundle_runtime_owned" => compatibility_bundle.fetch("runtime_owned"),
  "compatibility_evidence_bundle_go_runtime_backed" => compatibility_bundle.fetch("go_runtime_backed"),
  "compatibility_evidence_bundle_kde_policy_owner" => compatibility_bundle.fetch("kde_policy_owner"),
  "application_detail_generated" => true,
  "application_detail_request_type" => application_detail.fetch("request_type"),
  "application_detail_safe_for_kde" => application_detail.fetch("safe_for_kde"),
  "application_detail_safe_for_ai_diagnostics" => application_detail.fetch("safe_for_ai_diagnostics"),
  "application_detail_real_windows_app_run_verified" => application_detail.fetch("real_windows_app_run_verified"),
  "application_detail_file_open_verified" => application_detail.fetch("file_open_verified"),
  "application_detail_runtime_owned" => application_detail.fetch("runtime_owned"),
  "application_detail_go_runtime_backed" => application_detail.fetch("go_runtime_backed"),
  "application_detail_kde_policy_owner" => application_detail.fetch("kde_policy_owner"),
  "kde_page_from_application_detail_generated" => true,
  "kde_page_from_application_detail_request_type" => kde_page_from_detail.fetch("request_type"),
  "kde_page_from_application_detail_consumed" => kde_page_from_detail.fetch("external_winapp_application_detail_consumed"),
  "kde_page_from_application_detail_count" => kde_page_from_detail.fetch("external_winapp_application_detail_count"),
  "kde_page_from_application_detail_safe_for_ai_diagnostics" => kde_page_from_detail.fetch("safe_for_ai_diagnostics"),
  "kde_page_from_application_detail_runtime_owned" => kde_page_from_detail.fetch("runtime_owned"),
  "kde_page_from_application_detail_go_runtime_backed" => kde_page_from_detail.fetch("go_runtime_backed"),
  "kde_page_from_application_detail_kde_policy_owner" => kde_page_from_detail.fetch("kde_policy_owner"),
  "kde_page_from_application_detail_header_badge" => kde_page_from_detail_header.fetch("badge"),
  "kde_page_from_application_detail_header_badge_tone" => kde_page_from_detail_header.fetch("badge_tone"),
  "kde_page_from_application_detail_header_backend_details_exposed" => kde_page_from_detail_header.fetch("backend_details_exposed"),
  "kde_page_from_application_detail_summary_compatibility_state" => kde_page_from_detail_application_summary.fetch("compatibility_state"),
  "kde_page_from_application_detail_summary_compatibility_label" => kde_page_from_detail_application_summary.fetch("compatibility_label"),
  "kde_page_from_application_detail_summary_diagnostics_state" => kde_page_from_detail_application_summary.fetch("diagnostics_state"),
  "kde_page_from_application_detail_summary_backend_launch_enabled" => kde_page_from_detail_application_summary.fetch("backend_launch_enabled"),
  "kde_page_from_application_detail_summary_action_execution_enabled" => kde_page_from_detail_application_summary.fetch("action_execution_enabled"),
  "kde_page_from_application_detail_summary_settings_persistence_enabled" => kde_page_from_detail_application_summary.fetch("settings_persistence_enabled"),
  "kde_page_from_application_detail_summary_backend_details_exposed" => kde_page_from_detail_application_summary.fetch("backend_details_exposed"),
  "kde_page_from_application_detail_summary_host_root_modified" => kde_page_from_detail_application_summary.fetch("host_root_modified"),
  "artifact_fetch_count" => fetched_artifacts.count { |key, value| key.end_with?("_artifact_fetched") && value == true },
  "artifacts_fetched" => fetched_artifacts,
  "markdown_output_written" => markdown_output_path.file?,
  "host_compilation_avoided" => bool(build, "host_compilation_avoided"),
  "host_root_modified" => delegated.fetch("host_root_modified"),
  "privileged_container_required" => false,
  "host_networking_required" => false,
  "docker_socket_mounted" => delegated.fetch("docker_socket_mounted"),
  "broad_host_mount_required" => delegated.fetch("broad_host_mount_required")
)

emit_json(result, output_path)

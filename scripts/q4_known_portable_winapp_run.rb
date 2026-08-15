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
NOTEPADPP_SMOKE = PROJECT_ROOT.join("scripts/q4_notepadpp_portable_winapp_smoke.rb")
PUTTY_SMOKE = PROJECT_ROOT.join("scripts/q4_putty_external_winapp_smoke.rb")

SCHEMA_VERSION = "xnix.scripts.q4_known_portable_winapp_run.v1"
REQUEST_TYPE = "q4-known-portable-winapp-run"
SUPPORTED_NOTEPADPP_APP_ID = "org.xnix.external.notepadplusplus"
SUPPORTED_PUTTY_APP_ID = "org.xnix.external.putty"
SUPPORTED_APP_IDS = [SUPPORTED_NOTEPADPP_APP_ID, SUPPORTED_PUTTY_APP_ID].freeze
SUPPORTED_DISPLAY_NAMES = {
  SUPPORTED_NOTEPADPP_APP_ID => "Notepad++ Portable",
  SUPPORTED_PUTTY_APP_ID => "PuTTY"
}.freeze

DEFAULT_REMOTE_HOST = ENV.fetch("XNIX_REMOTE_HOST", "root@q4")
DEFAULT_LOCAL_SHELL = ENV.fetch("XNIX_LOCAL_SHELL", "/bin/zsh")
DEFAULT_REMOTE_MATERIALS_ROOT = ENV.fetch("XNIX_REMOTE_MATERIALS_ROOT", "/home/xnix-run-materials")
DEFAULT_REMOTE_SOURCE_ROOT = ENV.fetch("XNIX_Q4_KNOWN_PORTABLE_SOURCE_ROOT", "/home/xnix-build/xnix-q4-known-portable-winapp-run-#{VERSION}")
DEFAULT_REMOTE_BUILD_ROOT = ENV.fetch("XNIX_REMOTE_BUILD_ROOT", "/home/xnix-build-cache")
DEFAULT_REMOTE_TIMEOUT_SECONDS = Integer(ENV.fetch("XNIX_Q4_KNOWN_PORTABLE_WINAPP_RUN_TIMEOUT_SECONDS", "1800"), 10)
DEFAULT_IMAGE = ENV.fetch("XNIX_WINE_IMAGE", "xnix-wine-smoke:local")
DEFAULT_OUTPUT = PROJECT_ROOT.join("output", "q4-known-portable-winapp-run-#{VERSION}.json").to_s
DEFAULT_MARKDOWN_OUTPUT = PROJECT_ROOT.join("output", "q4-known-portable-winapp-run-#{VERSION}.md").to_s

options = {
  execute: false,
  app: ENV.fetch("XNIX_Q4_KNOWN_PORTABLE_WINAPP_APP", SUPPORTED_NOTEPADPP_APP_ID),
  local_shell: DEFAULT_LOCAL_SHELL,
  remote_host: DEFAULT_REMOTE_HOST,
  remote_materials_root: DEFAULT_REMOTE_MATERIALS_ROOT,
  remote_source_root: DEFAULT_REMOTE_SOURCE_ROOT,
  remote_build_root: DEFAULT_REMOTE_BUILD_ROOT,
  remote_timeout_seconds: DEFAULT_REMOTE_TIMEOUT_SECONDS,
  image: DEFAULT_IMAGE,
  output: ENV.fetch("XNIX_Q4_KNOWN_PORTABLE_WINAPP_RUN_OUTPUT", DEFAULT_OUTPUT),
  markdown_output: ENV.fetch("XNIX_Q4_KNOWN_PORTABLE_WINAPP_RUN_MARKDOWN_OUTPUT", DEFAULT_MARKDOWN_OUTPUT)
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/q4_known_portable_winapp_run.rb --app APP_ID [--execute]"
  parser.on("--execute", "Run a catalog-backed known portable Windows app through the q4 Runtime/KDE lane.") { options[:execute] = true }
  parser.on("--app APP_ID", "Known portable app id, currently: #{SUPPORTED_APP_IDS.join(", ")}.") { |value| options[:app] = value }
  parser.on("--local-shell PATH", "Local shell used for SSH alias resolution.") { |value| options[:local_shell] = value }
  parser.on("--remote HOST", "Remote SSH target, default: #{DEFAULT_REMOTE_HOST}.") { |value| options[:remote_host] = value }
  parser.on("--remote-materials-root PATH", "Remote materials root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_materials_root] = value }
  parser.on("--remote-source-root PATH", "Remote Runtime source root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_source_root] = value }
  parser.on("--remote-build-root PATH", "Remote Runtime build/cache root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_build_root] = value }
  parser.on("--image IMAGE", "q4-local Wine GUI smoke image, default: #{DEFAULT_IMAGE}.") { |value| options[:image] = value }
  parser.on("--remote-timeout-seconds SECONDS", Integer, "Timeout for q4 download, extraction, compile, and execution.") { |value| options[:remote_timeout_seconds] = value }
  parser.on("--output PATH", "Write JSON result under this checkout or /tmp/xnix-*.") { |value| options[:output] = value }
  parser.on("--markdown-output PATH", "Write Markdown result under this checkout or /tmp/xnix-*.") { |value| options[:markdown_output] = value }
end.parse!

abort "q4 known portable Windows app run does not accept positional arguments" unless ARGV.empty?

def ensure_local_output_path!(label, path)
  clean = Pathname.new(path).expand_path(PROJECT_ROOT).cleanpath
  return clean if clean.to_s.start_with?(PROJECT_ROOT.to_s)
  return clean if clean.to_s.start_with?("/tmp/xnix-")

  abort "#{label} must stay under this checkout or /tmp/xnix-*"
end

def ensure_remote_xnix_path!(label, path)
  clean = Pathname.new(path).cleanpath.to_s
  return clean if clean.start_with?("/home/xnix-") || clean.start_with?("/tmp/xnix-")

  abort "#{label} must stay under /home/xnix-* or /tmp/xnix-* on q4"
end

def ensure_supported_app!(app_id)
  clean = app_id.to_s.strip
  abort "known portable app id must be non-empty" if clean.empty?
  abort "unsupported known portable app #{clean}; currently supported: #{SUPPORTED_APP_IDS.join(", ")}" unless SUPPORTED_APP_IDS.include?(clean)

  clean
end

def emit_json(payload, output_path)
  text = JSON.pretty_generate(payload) + "\n"
  FileUtils.mkdir_p(output_path.dirname)
  File.write(output_path, text)
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
  stderr = [stderr, "q4 known portable Windows app run operation timed out after #{timeout_seconds}s"].reject(&:empty?).join("\n") if timed_out
  [stdout, stderr, timed_out ? 124 : status.exitstatus]
end

def push_remote_artifact(shell, remote_host, local_path, remote_path, timeout_seconds:)
  _stdout, stderr, status = run_shell(
    shell,
    shell_join(["scp", "-q", local_path.to_s, "#{remote_host}:#{remote_path}"]),
    timeout_seconds: timeout_seconds
  )
  abort "failed to push q4 known portable run artifact #{remote_path}: #{stderr}" unless status.zero?

  true
end

def fetch_remote_artifact(shell, remote_host, remote_path, local_path, timeout_seconds:)
  FileUtils.mkdir_p(local_path.dirname)
  _stdout, stderr, status = run_shell(
    shell,
    shell_join(["scp", "-q", "#{remote_host}:#{remote_path}", local_path.to_s]),
    timeout_seconds: timeout_seconds
  )
  abort "failed to fetch q4 known portable run artifact #{remote_path}: #{stderr}" unless status.zero?

  true
end

def remote_project_command(remote_source_root, argv)
  [
    "cd #{Shellwords.escape(remote_source_root)}",
    shell_join(argv)
  ].join("\n")
end

def write_markdown(payload, markdown_output_path)
  lines = [
    "# q4 Known Portable Windows App Run",
    "",
    "- Status: `#{payload.fetch("status")}`",
    "- App: `#{payload.fetch("display_name")}` (`#{payload.fetch("app_id")}`)",
    "- Version: `#{payload.fetch("version")}`",
    "- Operator run ready: `#{payload.fetch("operator_run_ready", false)}`",
    "- Operator acceptance ready: `#{payload.fetch("operator_run_acceptance_ready", false)}`",
    "- Accepted application detail state: `#{payload.fetch("accepted_application_detail_state", "planned")}`",
    "- KDE accepted page state: `#{payload.fetch("kde_accepted_page_state", "planned")}`",
    "- q4 download required: `#{payload.fetch("q4_download_required")}`",
    "- q4 compile required: `#{payload.fetch("q4_compile_required")}`",
    "- Host compilation avoided: `#{payload.fetch("host_compilation_avoided")}`",
    "- Host download avoided: `#{payload.fetch("host_download_avoided")}`",
    "- Unsafe gates: privileged=`#{payload.fetch("privileged_container_required")}`, host_network=`#{payload.fetch("host_networking_required")}`, docker_socket=`#{payload.fetch("docker_socket_mounted")}`, broad_mount=`#{payload.fetch("broad_host_mount_required")}`, host_root_modified=`#{payload.fetch("host_root_modified")}`",
    ""
  ]
  FileUtils.mkdir_p(markdown_output_path.dirname)
  File.write(markdown_output_path, lines.join("\n"))
end

def bool(payload, key)
  payload[key] == true
end

execute = options.fetch(:execute)
app_id = ensure_supported_app!(options.fetch(:app))
display_name = SUPPORTED_DISPLAY_NAMES.fetch(app_id)
notepadpp_lane = app_id == SUPPORTED_NOTEPADPP_APP_ID
putty_lane = app_id == SUPPORTED_PUTTY_APP_ID
output_path = ensure_local_output_path!("output path", options.fetch(:output))
markdown_output_path = ensure_local_output_path!("markdown output path", options.fetch(:markdown_output))
remote_materials_root = ensure_remote_xnix_path!("remote materials root", options.fetch(:remote_materials_root))
remote_source_root = ensure_remote_xnix_path!("remote source root", options.fetch(:remote_source_root))
remote_build_root = ensure_remote_xnix_path!("remote build root", options.fetch(:remote_build_root))
artifact_root = PROJECT_ROOT.join("output", "artifacts")
delegated_slug = notepadpp_lane ? "notepadpp" : "putty"
delegated_output_path = artifact_root.join("q4-known-portable-winapp-run-#{delegated_slug}-smoke-#{VERSION}.json")
delegated_markdown_output_path = artifact_root.join("q4-known-portable-winapp-run-#{delegated_slug}-smoke-#{VERSION}.md")
remote_acceptance_root = if notepadpp_lane
                           "#{remote_materials_root}/operator-run/notepad-plus-plus/8.9.7"
                         else
                           "#{remote_materials_root}/operator-run/putty/0.84"
                         end
remote_acceptance_input = "#{remote_acceptance_root}/q4-known-portable-winapp-run-acceptance-input.json"
remote_acceptance_output = "#{remote_acceptance_root}/q4-known-portable-winapp-run-acceptance.json"
remote_operator_detail = "#{remote_acceptance_root}/q4-known-portable-winapp-run-application-detail.json"
remote_operator_kde_page = "#{remote_acceptance_root}/q4-known-portable-winapp-run-kde-page.json"
acceptance_local_path = artifact_root.join("q4-known-portable-winapp-run-acceptance-#{VERSION}.json")
operator_detail_local_path = artifact_root.join("q4-known-portable-winapp-run-application-detail-#{VERSION}.json")
operator_kde_page_local_path = artifact_root.join("q4-known-portable-winapp-run-kde-page-#{VERSION}.json")

delegated_script = notepadpp_lane ? "scripts/q4_notepadpp_portable_winapp_smoke.rb" : "scripts/q4_putty_external_winapp_smoke.rb"
delegated_request_type = notepadpp_lane ? "q4-notepadpp-portable-winapp-smoke" : "q4-putty-external-winapp-smoke"
operator_acceptance_type = notepadpp_lane ? "q4-known-portable-bundle-winapp-acceptance-preview" : "q4-staged-external-winapp-acceptance-preview"
catalog_artifact_kind = notepadpp_lane ? "portable-zip-bundle" : "single-executable"
download_artifact_name = notepadpp_lane ? "npp.8.9.7.portable.zip" : "putty.exe"
executable_relative_path = notepadpp_lane ? "notepad++.exe" : "putty.exe"

delegated_command = if notepadpp_lane
                      [
  "ruby",
  "scripts/q4_notepadpp_portable_winapp_smoke.rb",
  "--remote", options.fetch(:remote_host),
  "--remote-materials-root", remote_materials_root,
  "--remote-source-root", remote_source_root,
  "--remote-build-root", remote_build_root,
  "--image", options.fetch(:image),
  "--remote-timeout-seconds", options.fetch(:remote_timeout_seconds).to_s,
  "--output", delegated_output_path.to_s,
  "--markdown-output", delegated_markdown_output_path.to_s
                      ]
                    else
                      [
                        "ruby",
                        "scripts/q4_putty_external_winapp_smoke.rb",
                        "--remote", options.fetch(:remote_host),
                        "--remote-materials-root", remote_materials_root,
                        "--remote-timeout-seconds", options.fetch(:remote_timeout_seconds).to_s,
                        "--output", delegated_output_path.to_s,
                        "--markdown-output", delegated_markdown_output_path.to_s
                      ]
                    end
delegated_command << "--execute" if execute

plan = {
  "schema_version" => SCHEMA_VERSION,
  "request_type" => REQUEST_TYPE,
  "version" => VERSION,
  "status" => execute ? "running" : "planned",
  "execute" => execute,
  "app_id" => app_id,
  "display_name" => display_name,
  "known_catalog_app" => true,
  "portable_directory_external_app" => notepadpp_lane,
  "single_file_external_app" => putty_lane,
  "q4_execute_supported" => true,
  "argumentless_desktop_launch_lane_required" => false,
  "argumentless_desktop_launch_lane_ready" => putty_lane,
  "argumentless_desktop_launch_lane_reason" => putty_lane ? "PuTTY uses the argumentless desktop launch lane because it is a GUI SSH client rather than a file-open editor." : "",
  "official_download_required" => true,
  "pinned_checksum_required" => true,
  "remote_host" => options.fetch(:remote_host),
  "remote_materials_root" => remote_materials_root,
  "remote_source_root" => remote_source_root,
  "remote_build_root" => remote_build_root,
  "remote_paths_exposed_only_for_operator" => true,
  "go_runtime_plan_required" => true,
  "go_runtime_plan_schema_version" => "xnix.runtime.q4_known_portable_winapp_run_plan.v1",
  "go_runtime_plan_request_type" => "q4-known-portable-winapp-run-plan-preview",
  "go_runtime_plan_status" => "planned",
  "go_runtime_plan_catalog_consumed" => false,
  "go_runtime_plan_command" => [
    "xnix-runtime-go",
    "q4-known-portable-winapp-run-plan-preview",
    "--app", app_id,
    "--remote", options.fetch(:remote_host),
    "--remote-materials-root", remote_materials_root,
    "--remote-source-root", remote_source_root,
    "--remote-build-root", remote_build_root,
    "--output", output_path.to_s,
    "--markdown-output", markdown_output_path.to_s
  ],
  "delegated_script" => delegated_script,
  "delegated_request_type" => delegated_request_type,
  "delegated_command" => delegated_command,
  "delegated_output_path" => delegated_output_path.to_s,
  "delegated_markdown_output_path" => delegated_markdown_output_path.to_s,
  "operator_run_acceptance_planned" => true,
  "operator_run_acceptance_request_type" => "q4-known-portable-winapp-run-acceptance-preview",
  "operator_run_underlying_acceptance_request_type" => operator_acceptance_type,
  "operator_run_acceptance_report" => remote_acceptance_output,
  "operator_run_acceptance_status" => "planned",
  "operator_run_application_detail_planned" => true,
  "operator_run_application_detail_report" => remote_operator_detail,
  "operator_run_application_detail_status" => "planned",
  "operator_run_kde_page_planned" => true,
  "operator_run_kde_page_report" => remote_operator_kde_page,
  "operator_run_kde_page_status" => "planned",
  "output_path" => output_path.to_s,
  "markdown_output_path" => markdown_output_path.to_s,
  "runtime_owned" => true,
  "go_runtime_backed" => true,
  "kde_policy_owner" => false,
  "q4_download_required" => true,
  "q4_extract_required" => notepadpp_lane,
  "q4_compile_required" => true,
  "q4_execution_required" => true,
  "host_compilation_avoided" => true,
  "host_download_avoided" => true,
  "full_smoke_required" => false,
  "host_root_modified" => false,
  "privileged_container_required" => false,
  "host_networking_required" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false
}

unless execute
  write_markdown(plan, markdown_output_path)
  emit_json(plan, output_path)
  exit 0
end

delegated_args = delegated_command.dup
delegated_args[1] = (notepadpp_lane ? NOTEPADPP_SMOKE : PUTTY_SMOKE).to_s
stdout, stderr, status = Open3.capture3(*delegated_args, chdir: PROJECT_ROOT.to_s)
unless status.success?
  warn stdout unless stdout.empty?
  warn stderr unless stderr.empty?
  abort "q4 known portable Windows app delegated run failed"
end

delegated = JSON.parse(stdout)
notepadpp_passed = notepadpp_lane &&
                   delegated.fetch("status") == "passed" &&
                   delegated.fetch("app_id") == app_id &&
                   delegated.fetch("display_name") == display_name &&
                   bool(delegated, "notepadpp_sha256_verified") &&
                   bool(delegated, "known_portable_bundle_archive_verified") &&
                   bool(delegated, "known_portable_bundle_stage_launch_existing_import_record_consumed") &&
                   bool(delegated, "known_portable_bundle_stage_launch_windows_process_file_argument_window_observed") &&
                   bool(delegated, "known_portable_bundle_acceptance_ready") &&
                   bool(delegated, "known_portable_bundle_acceptance_runtime_accepted_chain_verified") &&
                   delegated.fetch("known_portable_bundle_acceptance_launch_source_request_type") == "windows-known-app-bundle-stage-and-launch" &&
                   delegated.fetch("accepted_application_detail_state") == "runtime-accepted-real-app-run" &&
                   delegated.fetch("kde_accepted_page_state") == "runtime-accepted-real-app-run"
putty_passed = putty_lane &&
               delegated.fetch("status") == "passed" &&
               delegated.fetch("app_id") == app_id &&
               delegated.fetch("display_name") == display_name &&
               bool(delegated, "putty_sha256_verified") &&
               bool(delegated, "single_file_windows_app") &&
               delegated.fetch("desktop_argument_mode", "") == "none" &&
               !bool(delegated, "external_file_open_requested") &&
               bool(delegated, "window_observed") &&
               bool(delegated, "x_window_observed") &&
               bool(delegated, "staged_external_winapp_acceptance_ready") &&
               delegated.fetch("staged_external_winapp_acceptance_request_type") == "q4-staged-external-winapp-acceptance-preview" &&
               delegated.fetch("accepted_application_detail_state") == "runtime-accepted-real-app-run" &&
               delegated.fetch("kde_accepted_page_state") == "runtime-accepted-real-app-run"
common_delegated_passed = bool(delegated, "host_compilation_avoided") &&
                          bool(delegated, "host_download_avoided") &&
                          !bool(delegated, "host_root_modified") &&
                          !bool(delegated, "privileged_container_required") &&
                          !bool(delegated, "host_networking_required") &&
                          !bool(delegated, "docker_socket_mounted") &&
                          !bool(delegated, "broad_host_mount_required")
unless (notepadpp_passed || putty_passed) && common_delegated_passed
  warn stdout
  abort "q4 known portable Windows app run did not reach accepted state"
end

lane_result = if notepadpp_lane
                {
                  "official_archive_checksum_verified" => bool(delegated, "notepadpp_sha256_verified"),
                  "official_executable_checksum_verified" => false,
                  "known_portable_bundle_imported" => bool(delegated, "known_portable_bundle_import_recorded"),
                  "known_portable_bundle_stage_launch_status" => delegated.fetch("known_portable_bundle_stage_launch_status"),
                  "known_portable_bundle_acceptance_request_type" => delegated.fetch("known_portable_bundle_acceptance_request_type"),
                  "known_portable_bundle_acceptance_ready" => bool(delegated, "known_portable_bundle_acceptance_ready"),
                  "staged_external_winapp_acceptance_request_type" => "",
                  "staged_external_winapp_acceptance_ready" => false,
                  "external_file_bridge_ready" => false,
                  "windows_process_file_argument_window_observed" => bool(delegated, "known_portable_bundle_stage_launch_windows_process_file_argument_window_observed"),
                  "runtime_accepted_chain_verified" => bool(delegated, "known_portable_bundle_acceptance_runtime_accepted_chain_verified"),
                  "compatibility_evidence_bundle_report" => delegated.fetch("known_portable_bundle_compatibility_bundle_report")
                }
              else
                {
                  "official_archive_checksum_verified" => false,
                  "official_executable_checksum_verified" => bool(delegated, "putty_sha256_verified"),
                  "known_portable_bundle_imported" => false,
                  "known_portable_bundle_stage_launch_status" => "",
                  "known_portable_bundle_acceptance_request_type" => "",
                  "known_portable_bundle_acceptance_ready" => false,
                  "staged_external_winapp_acceptance_request_type" => delegated.fetch("staged_external_winapp_acceptance_request_type"),
                  "staged_external_winapp_acceptance_ready" => bool(delegated, "staged_external_winapp_acceptance_ready"),
                  "external_file_bridge_ready" => bool(delegated, "external_file_bridge_ready"),
                  "windows_process_file_argument_window_observed" => bool(delegated, "windows_process_file_argument_window_observed"),
                  "external_file_open_requested" => bool(delegated, "external_file_open_requested"),
                  "window_observed" => bool(delegated, "window_observed"),
                  "x_window_observed" => bool(delegated, "x_window_observed"),
                  "runtime_accepted_chain_verified" => bool(delegated, "staged_external_winapp_acceptance_ready"),
                  "compatibility_evidence_bundle_report" => delegated.fetch("compatibility_evidence_bundle_report")
                }
              end

result = plan.merge(
  "status" => "passed",
  "delegated_status" => delegated.fetch("status"),
  "delegated_artifact_fetch_count" => delegated.fetch("artifact_fetch_count"),
  "catalog_artifact_kind" => catalog_artifact_kind,
  "download_artifact_name" => download_artifact_name,
  "executable_relative_path" => executable_relative_path,
  "accepted_application_detail_state" => delegated.fetch("accepted_application_detail_state"),
  "kde_accepted_page_state" => delegated.fetch("kde_accepted_page_state"),
  "operator_run_ready" => true
).merge(lane_result)

abort "q4 known portable Windows app run missing compatibility evidence bundle report" if result.fetch("compatibility_evidence_bundle_report").to_s.empty?

remote_runtime_binary = delegated.fetch("remote_runtime_binary")
remote_runtime_source_root = delegated.fetch("remote_runtime_source_root", remote_source_root)
runtime_plan_command = [
  remote_runtime_binary,
  "q4-known-portable-winapp-run-plan-preview",
  "--app", app_id,
  "--remote", options.fetch(:remote_host),
  "--remote-materials-root", remote_materials_root,
  "--remote-source-root", remote_source_root,
  "--remote-build-root", remote_build_root
]
runtime_plan_stdout, runtime_plan_stderr, runtime_plan_status = run_shell(
  options.fetch(:local_shell),
  shell_join(ssh_command(options.fetch(:remote_host), remote_project_command(remote_runtime_source_root, runtime_plan_command))),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless runtime_plan_status.zero?
  warn runtime_plan_stdout unless runtime_plan_stdout.empty?
  warn runtime_plan_stderr unless runtime_plan_stderr.empty?
  abort "q4 known portable Windows app run Go Runtime plan failed"
end
runtime_plan = JSON.parse(runtime_plan_stdout)
unless runtime_plan.fetch("request_type") == "q4-known-portable-winapp-run-plan-preview" &&
       runtime_plan.fetch("app_id") == app_id &&
       runtime_plan.fetch("display_name") == display_name &&
       runtime_plan.fetch("catalog_artifact_kind") == catalog_artifact_kind &&
       runtime_plan.fetch("download_artifact_name") == download_artifact_name &&
       runtime_plan.fetch("executable_relative_path") == executable_relative_path &&
       runtime_plan.fetch("supported_known_portable_app_ids").include?(app_id) &&
       bool(runtime_plan, "known_catalog_app") &&
       bool(runtime_plan, "portable_directory_external_app") == notepadpp_lane &&
       bool(runtime_plan, "single_file_external_app") == putty_lane &&
       bool(runtime_plan, "q4_compile_required") &&
       bool(runtime_plan, "q4_extract_required") == notepadpp_lane &&
       bool(runtime_plan, "host_compilation_avoided") &&
       !bool(runtime_plan, "remote_paths_exposed_to_desktop") &&
       !bool(runtime_plan, "host_root_modified")
  warn runtime_plan_stdout
  abort "q4 known portable Windows app run Go Runtime plan did not consume catalog metadata"
end
result = result.merge(
  "go_runtime_plan_status" => "passed",
  "go_runtime_plan_catalog_consumed" => true,
  "go_runtime_plan_supported_known_portable_app_ids" => runtime_plan.fetch("supported_known_portable_app_ids"),
  "catalog_artifact_kind" => runtime_plan.fetch("catalog_artifact_kind"),
  "download_artifact_name" => runtime_plan.fetch("download_artifact_name"),
  "executable_relative_path" => runtime_plan.fetch("executable_relative_path"),
  "app_version" => runtime_plan.fetch("app_version")
)
acceptance_input_path = artifact_root.join("q4-known-portable-winapp-run-acceptance-input-#{VERSION}.json")
FileUtils.mkdir_p(acceptance_input_path.dirname)
File.write(acceptance_input_path, JSON.pretty_generate(result) + "\n")

mkdir_stdout, mkdir_stderr, mkdir_status = run_shell(
  options.fetch(:local_shell),
  shell_join(ssh_command(options.fetch(:remote_host), shell_join(["mkdir", "-p", remote_acceptance_root]))),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless mkdir_status.zero?
  warn mkdir_stdout unless mkdir_stdout.empty?
  warn mkdir_stderr unless mkdir_stderr.empty?
  abort "q4 known portable Windows app run acceptance directory preparation failed"
end
push_remote_artifact(
  options.fetch(:local_shell),
  options.fetch(:remote_host),
  acceptance_input_path,
  remote_acceptance_input,
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
acceptance_command = [
  remote_runtime_binary,
  "q4-known-portable-winapp-run-acceptance-preview",
  "--q4-known-portable-winapp-run", remote_acceptance_input,
  "--output", remote_acceptance_output
]
acceptance_stdout, acceptance_stderr, acceptance_status = run_shell(
  options.fetch(:local_shell),
  shell_join(ssh_command(options.fetch(:remote_host), remote_project_command(remote_runtime_source_root, acceptance_command))),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless acceptance_status.zero?
  warn acceptance_stdout unless acceptance_stdout.empty?
  warn acceptance_stderr unless acceptance_stderr.empty?
  abort "q4 known portable Windows app run Go acceptance failed"
end
operator_acceptance = JSON.parse(acceptance_stdout)
unless operator_acceptance.fetch("request_type") == "q4-known-portable-winapp-run-acceptance-preview" &&
       operator_acceptance.fetch("app_id") == app_id &&
       operator_acceptance.fetch("display_name") == display_name &&
       bool(operator_acceptance, "operator_run_ready") &&
       bool(operator_acceptance, "runtime_accepted_chain_verified") &&
       bool(operator_acceptance, "acceptance_ready") &&
       !bool(operator_acceptance, "remote_host_exposed") &&
       !bool(operator_acceptance, "raw_path_exposed")
  warn acceptance_stdout
  abort "q4 known portable Windows app run Go acceptance did not accept the operator run"
end
acceptance_fetched = fetch_remote_artifact(
  options.fetch(:local_shell),
  options.fetch(:remote_host),
  remote_acceptance_output,
  acceptance_local_path,
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)

operator_detail_command = [
  remote_runtime_binary,
  "external-winapp-application-detail-preview",
  "--compatibility-evidence-bundle", result.fetch("compatibility_evidence_bundle_report"),
  "--q4-known-portable-winapp-run-acceptance", remote_acceptance_output,
  "--output", remote_operator_detail
]
operator_detail_stdout, operator_detail_stderr, operator_detail_status = run_shell(
  options.fetch(:local_shell),
  shell_join(ssh_command(options.fetch(:remote_host), remote_project_command(remote_runtime_source_root, operator_detail_command))),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless operator_detail_status.zero?
  warn operator_detail_stdout unless operator_detail_stdout.empty?
  warn operator_detail_stderr unless operator_detail_stderr.empty?
  abort "q4 known portable Windows app run application detail generation failed"
end
operator_detail = JSON.parse(operator_detail_stdout)
unless operator_detail.fetch("request_type") == "external-winapp-application-detail-preview" &&
       operator_detail.fetch("application_id") == app_id &&
       operator_detail.fetch("display_name") == display_name &&
       operator_detail.fetch("compatibility_state") == "runtime-accepted-real-app-run" &&
       bool(operator_detail, "q4_known_portable_winapp_run_acceptance_consumed") &&
       bool(operator_detail, "q4_known_portable_winapp_run_acceptance_ready") &&
       bool(operator_detail, "go_owned_known_portable_winapp_run_acceptance_verified") &&
       !bool(operator_detail, "backend_details_exposed") &&
       !bool(operator_detail, "raw_paths_exposed") &&
       !bool(operator_detail, "host_root_modified")
  warn operator_detail_stdout
  abort "q4 known portable Windows app run application detail did not consume operator acceptance"
end
operator_detail_fetched = fetch_remote_artifact(
  options.fetch(:local_shell),
  options.fetch(:remote_host),
  remote_operator_detail,
  operator_detail_local_path,
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)

operator_kde_page_command = [
  remote_runtime_binary,
  "kde-center-page-preview",
  "--external-app-application-detail", remote_operator_detail,
  "--decision", "approved",
  "--output", remote_operator_kde_page
]
operator_kde_stdout, operator_kde_stderr, operator_kde_status = run_shell(
  options.fetch(:local_shell),
  shell_join(ssh_command(options.fetch(:remote_host), remote_project_command(remote_runtime_source_root, operator_kde_page_command))),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless operator_kde_status.zero?
  warn operator_kde_stdout unless operator_kde_stdout.empty?
  warn operator_kde_stderr unless operator_kde_stderr.empty?
  abort "q4 known portable Windows app run KDE page generation failed"
end
operator_kde_page = JSON.parse(operator_kde_stdout)
operator_kde_cards = operator_kde_page.fetch("external_winapp_application_detail_cards")
operator_kde_card = operator_kde_cards.first
unless operator_kde_page.fetch("request_type") == "kde-center-page-preview" &&
       operator_kde_page.fetch("application_id") == app_id &&
       operator_kde_card.fetch("compatibility_state") == "runtime-accepted-real-app-run" &&
       bool(operator_kde_card, "q4_known_portable_winapp_run_acceptance_consumed") &&
       bool(operator_kde_card, "q4_known_portable_winapp_run_acceptance_ready") &&
       bool(operator_kde_card, "go_owned_known_portable_winapp_run_acceptance_verified") &&
       !bool(operator_kde_card, "backend_details_exposed") &&
       !bool(operator_kde_card, "raw_paths_exposed") &&
       !bool(operator_kde_card, "host_root_modified")
  warn operator_kde_stdout
  abort "q4 known portable Windows app run KDE page did not consume operator accepted detail"
end
operator_kde_page_fetched = fetch_remote_artifact(
  options.fetch(:local_shell),
  options.fetch(:remote_host),
  remote_operator_kde_page,
  operator_kde_page_local_path,
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)

result = result.merge(
  "operator_run_acceptance_status" => "passed",
  "operator_run_acceptance_request_type" => operator_acceptance.fetch("request_type"),
  "operator_run_acceptance_ready" => bool(operator_acceptance, "acceptance_ready"),
  "operator_run_acceptance_runtime_accepted_chain_verified" => bool(operator_acceptance, "runtime_accepted_chain_verified"),
  "operator_run_acceptance_artifact_fetched" => acceptance_fetched,
  "operator_run_acceptance_artifact_output_path" => acceptance_local_path.to_s,
  "operator_run_application_detail_status" => "passed",
  "operator_run_application_detail_request_type" => operator_detail.fetch("request_type"),
  "operator_run_application_detail_acceptance_consumed" => bool(operator_detail, "q4_known_portable_winapp_run_acceptance_consumed"),
  "operator_run_application_detail_acceptance_ready" => bool(operator_detail, "q4_known_portable_winapp_run_acceptance_ready"),
  "operator_run_application_detail_artifact_fetched" => operator_detail_fetched,
  "operator_run_application_detail_artifact_output_path" => operator_detail_local_path.to_s,
  "operator_run_kde_page_status" => "passed",
  "operator_run_kde_page_request_type" => operator_kde_page.fetch("request_type"),
  "operator_run_kde_page_acceptance_consumed" => bool(operator_kde_card, "q4_known_portable_winapp_run_acceptance_consumed"),
  "operator_run_kde_page_acceptance_ready" => bool(operator_kde_card, "q4_known_portable_winapp_run_acceptance_ready"),
  "operator_run_kde_page_artifact_fetched" => operator_kde_page_fetched,
  "operator_run_kde_page_artifact_output_path" => operator_kde_page_local_path.to_s
)
write_markdown(result, markdown_output_path)
emit_json(result, output_path)

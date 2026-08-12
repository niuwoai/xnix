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

SCHEMA_VERSION = "xnix.scripts.q4_notepadpp_portable_winapp_smoke.v1"
REQUEST_TYPE = "q4-notepadpp-portable-winapp-smoke"
NOTEPADPP_VERSION = "8.9.7"
NOTEPADPP_ZIP = "npp.8.9.7.portable.zip"
NOTEPADPP_URL = "https://github.com/notepad-plus-plus/notepad-plus-plus/releases/download/v8.9.7/npp.8.9.7.portable.zip"
NOTEPADPP_SHA256 = "ce0690fac91c1fc5d61dcdf5b09733ff0d143a61d0a27c6cb9f4003ea92765bb"
NOTEPADPP_EXE_RELATIVE_PATH = "notepad++.exe"
NOTEPADPP_WINDOW_MATCH = "sample-document.txt"

DEFAULT_REMOTE_HOST = ENV.fetch("XNIX_REMOTE_HOST", "root@q4")
DEFAULT_LOCAL_SHELL = ENV.fetch("XNIX_LOCAL_SHELL", "/bin/zsh")
DEFAULT_REMOTE_MATERIALS_ROOT = ENV.fetch("XNIX_REMOTE_MATERIALS_ROOT", "/home/xnix-run-materials")
DEFAULT_REMOTE_SOURCE_ROOT = ENV.fetch("XNIX_Q4_NOTEPADPP_SOURCE_ROOT", "/home/xnix-build/xnix-q4-notepadpp-portable-#{VERSION}")
DEFAULT_REMOTE_BUILD_ROOT = ENV.fetch("XNIX_REMOTE_BUILD_ROOT", "/home/xnix-build-cache")
DEFAULT_REMOTE_TIMEOUT_SECONDS = Integer(ENV.fetch("XNIX_Q4_NOTEPADPP_TIMEOUT_SECONDS", "1800"), 10)
DEFAULT_OUTPUT = PROJECT_ROOT.join("output", "q4-notepadpp-portable-winapp-smoke-#{VERSION}.json").to_s
DEFAULT_MARKDOWN_OUTPUT = PROJECT_ROOT.join("output", "q4-notepadpp-portable-winapp-smoke-#{VERSION}.md").to_s
DEFAULT_IMAGE = ENV.fetch("XNIX_WINE_IMAGE", "xnix-wine-smoke:local")

options = {
  execute: false,
  local_shell: DEFAULT_LOCAL_SHELL,
  remote_host: DEFAULT_REMOTE_HOST,
  remote_materials_root: DEFAULT_REMOTE_MATERIALS_ROOT,
  remote_source_root: DEFAULT_REMOTE_SOURCE_ROOT,
  remote_build_root: DEFAULT_REMOTE_BUILD_ROOT,
  remote_timeout_seconds: DEFAULT_REMOTE_TIMEOUT_SECONDS,
  image: DEFAULT_IMAGE,
  output: ENV.fetch("XNIX_Q4_NOTEPADPP_OUTPUT", DEFAULT_OUTPUT),
  markdown_output: ENV.fetch("XNIX_Q4_NOTEPADPP_MARKDOWN_OUTPUT", DEFAULT_MARKDOWN_OUTPUT)
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/q4_notepadpp_portable_winapp_smoke.rb [--execute]"
  parser.on("--execute", "Download, verify, extract, and run the real Notepad++ Portable Windows GUI app on q4.") { options[:execute] = true }
  parser.on("--local-shell PATH", "Local shell used for SSH alias resolution.") { |value| options[:local_shell] = value }
  parser.on("--remote HOST", "Remote SSH target, default: #{DEFAULT_REMOTE_HOST}.") { |value| options[:remote_host] = value }
  parser.on("--remote-materials-root PATH", "Remote materials root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_materials_root] = value }
  parser.on("--remote-source-root PATH", "Remote Runtime source root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_source_root] = value }
  parser.on("--remote-build-root PATH", "Remote Runtime build/cache root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_build_root] = value }
  parser.on("--image IMAGE", "q4-local Wine GUI smoke image, default: #{DEFAULT_IMAGE}.") { |value| options[:image] = value }
  parser.on("--remote-timeout-seconds SECONDS", Integer, "Timeout for q4 download, extraction, compile, and staged execution.") { |value| options[:remote_timeout_seconds] = value }
  parser.on("--output PATH", "Write JSON result under this checkout or /tmp/xnix-*.") { |value| options[:output] = value }
  parser.on("--markdown-output PATH", "Write Markdown result under this checkout or /tmp/xnix-*.") { |value| options[:markdown_output] = value }
end.parse!

abort "q4 Notepad++ Portable Windows app smoke does not accept positional arguments" unless ARGV.empty?

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
  stderr = [stderr, "q4 Notepad++ Portable Windows app smoke operation timed out after #{timeout_seconds}s"].reject(&:empty?).join("\n") if timed_out
  [stdout, stderr, timed_out ? 124 : status.exitstatus]
end

def emit_json(payload, output_path)
  text = JSON.pretty_generate(payload) + "\n"
  FileUtils.mkdir_p(output_path.dirname)
  File.write(output_path, text)
  puts text
end

def bool(payload, key)
  payload[key] == true
end

def remote_step!(shell, remote_host, remote_command, timeout_seconds, failure)
  stdout, stderr, status = run_shell(
    shell,
    shell_join(ssh_command(remote_host, remote_command)),
    timeout_seconds: timeout_seconds
  )
  unless status.zero?
    warn stdout unless stdout.empty?
    warn stderr unless stderr.empty?
    abort failure
  end
  stdout
end

def fetch_remote_artifact(shell, remote_host, remote_path, local_path, timeout_seconds:)
  FileUtils.mkdir_p(local_path.dirname)
  _stdout, stderr, status = run_shell(
    shell,
    shell_join(["scp", "-q", "#{remote_host}:#{remote_path}", local_path.to_s]),
    timeout_seconds: timeout_seconds
  )
  abort "failed to fetch q4 artifact #{remote_path}: #{stderr}" unless status.zero?

  true
end

def run_json_command(shell, argv, timeout_seconds:)
  stdout, stderr, status = run_shell(
    shell,
    shell_join(argv),
    timeout_seconds: timeout_seconds
  )
  unless status.zero?
    warn stdout unless stdout.empty?
    warn stderr unless stderr.empty?
    abort "command failed: #{argv.first}"
  end
  JSON.parse(stdout)
end

remote_host = options.fetch(:remote_host)
remote_materials_root = ensure_remote_xnix_path!("remote materials root", options.fetch(:remote_materials_root))
remote_source_root = ensure_remote_xnix_path!("remote source root", options.fetch(:remote_source_root))
remote_build_root = ensure_remote_xnix_path!("remote build root", options.fetch(:remote_build_root))
remote_app_root = "#{remote_materials_root}/third-party/notepad-plus-plus/#{NOTEPADPP_VERSION}"
remote_cache_root = "#{remote_app_root}/cache"
remote_state_root = "#{remote_app_root}/runtime-state"
remote_staging_root = "#{remote_app_root}/runtime-staging"
remote_document = "#{remote_app_root}/sample-document.txt"
remote_known_bundle_import_report = "#{remote_app_root}/known-portable-bundle-import-record.json"
remote_known_bundle_stage_launch_report = "#{remote_app_root}/known-portable-bundle-stage-launch.json"
remote_known_bundle_gui_evidence_packet = "#{remote_app_root}/known-portable-bundle-gui-evidence-packet.json"
remote_binary_root = "#{remote_build_root}/bin/linux-amd64"
remote_runtime_bin = "#{remote_binary_root}/xnix-runtime-go"
remote_launcher_bin = "#{remote_binary_root}/xnix-compat-launch"
output_path = ensure_local_output_path!("output path", options.fetch(:output))
markdown_output_path = ensure_local_output_path!("markdown output path", options.fetch(:markdown_output))
artifact_output_root = output_path.dirname.join("artifacts")

known_bundle_import_command = [
  remote_runtime_bin,
  "windows-known-app-bundle-import-record",
  "--app", "org.xnix.external.notepadplusplus",
  "--cache-root", remote_cache_root,
  "--state-root", remote_state_root,
  "--allow-download",
  "--include-fetch",
  "--timeout", "#{options.fetch(:remote_timeout_seconds)}s",
  "--report-output", remote_known_bundle_import_report
]

known_bundle_stage_launch_command = [
  remote_runtime_bin,
  "windows-known-app-bundle-stage-and-launch",
  "--app", "org.xnix.external.notepadplusplus",
  "--cache-root", remote_cache_root,
  "--state-root", remote_state_root,
  "--staging-root", remote_staging_root,
  "--managed-launcher-bin", remote_launcher_bin,
  "--allow-download",
  "--include-fetch",
  "--image", options.fetch(:image),
  "--timeout", "#{options.fetch(:remote_timeout_seconds)}s",
  "--report-output", remote_known_bundle_stage_launch_report,
  "file://#{remote_document}"
]

known_bundle_gui_evidence_packet_command = [
  remote_runtime_bin,
  "real-winapp-gui-evidence-packet-preview",
  "--gui-smoke-report", remote_known_bundle_stage_launch_report,
  "--output", remote_known_bundle_gui_evidence_packet
]

def staged_command(remote_host, remote_import_record, output_path, markdown_output_path, timeout_seconds)
  [
    "ruby",
    "scripts/q4_staged_desktop_external_winapp_smoke.rb",
    "--execute",
    "--fixture", "external",
    "--remote-import-record", remote_import_record,
    "--app-id", "org.xnix.external.notepadplusplus",
    "--display-name", "Notepad++ Portable",
    "--window-match", NOTEPADPP_WINDOW_MATCH,
    "--remote", remote_host,
    "--remote-timeout-seconds", timeout_seconds.to_s,
    "--output", output_path.to_s,
    "--markdown-output", markdown_output_path.to_s
  ]
end

planned_staged_command = [
  "ruby",
  "scripts/q4_staged_desktop_external_winapp_smoke.rb",
  "--execute",
  "--fixture", "external",
  "--remote-import-record", "#{remote_state_root}/external-apps/org.xnix.external.notepadplusplus/import-record.json",
  "--app-id", "org.xnix.external.notepadplusplus",
  "--display-name", "Notepad++ Portable",
  "--window-match", NOTEPADPP_WINDOW_MATCH,
  "--remote", remote_host,
  "--remote-timeout-seconds", options.fetch(:remote_timeout_seconds).to_s,
  "--output", output_path.to_s,
  "--markdown-output", markdown_output_path.to_s
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
  "remote_runtime_binary" => remote_runtime_bin,
  "source_kind" => "official-notepad-plus-plus-github-release",
  "notepadpp_version" => NOTEPADPP_VERSION,
  "notepadpp_zip" => NOTEPADPP_ZIP,
  "notepadpp_source_page" => "https://github.com/notepad-plus-plus/notepad-plus-plus/releases/tag/v#{NOTEPADPP_VERSION}",
  "notepadpp_download_url" => NOTEPADPP_URL,
  "notepadpp_sha256" => NOTEPADPP_SHA256,
  "notepadpp_download_planned" => true,
  "notepadpp_downloaded" => false,
  "notepadpp_sha256_verified" => false,
  "notepadpp_extracted" => false,
  "notepadpp_executable_configured" => true,
  "notepadpp_executable_relative_path" => NOTEPADPP_EXE_RELATIVE_PATH,
  "known_portable_bundle_import_planned" => true,
  "known_portable_bundle_import_command" => known_bundle_import_command,
  "known_portable_bundle_import_report" => remote_known_bundle_import_report,
  "known_portable_bundle_import_recorded" => false,
  "known_portable_bundle_import_status" => "planned",
  "go_known_portable_bundle_import_backed" => true,
  "known_portable_bundle_stage_launch_planned" => true,
  "known_portable_bundle_stage_launch_command" => known_bundle_stage_launch_command,
  "known_portable_bundle_stage_launch_report" => remote_known_bundle_stage_launch_report,
  "known_portable_bundle_stage_launch_status" => "planned",
  "go_known_portable_bundle_stage_launch_backed" => true,
  "known_portable_bundle_gui_evidence_packet_planned" => true,
  "known_portable_bundle_gui_evidence_packet_command" => known_bundle_gui_evidence_packet_command,
  "known_portable_bundle_gui_evidence_packet_report" => remote_known_bundle_gui_evidence_packet,
  "known_portable_bundle_gui_evidence_packet_status" => "planned",
  "record_first_launch_path" => true,
  "remote_cache_root" => remote_cache_root,
  "remote_state_root" => remote_state_root,
  "remote_staging_root" => remote_staging_root,
  "remote_launcher_binary" => remote_launcher_bin,
  "portable_directory_external_app" => true,
  "portable_directory_bundle_import_required" => true,
  "remote_materials_root" => remote_materials_root,
  "remote_app_root" => remote_app_root,
  "remote_zip" => "#{remote_cache_root}/org.xnix.external.notepadplusplus/#{NOTEPADPP_ZIP}",
  "remote_bundle_root" => "",
  "remote_executable" => "",
  "remote_import_record" => "",
  "remote_paths_exposed_only_for_operator" => true,
  "app_id" => "org.xnix.external.notepadplusplus",
  "display_name" => "Notepad++ Portable",
  "window_match" => NOTEPADPP_WINDOW_MATCH,
  "real_third_party_windows_app" => true,
  "single_file_windows_app" => false,
  "delegated_script" => "scripts/q4_staged_desktop_external_winapp_smoke.rb",
  "delegated_command" => planned_staged_command,
  "output_path" => output_path.to_s,
  "markdown_output_path" => markdown_output_path.to_s,
  "artifact_output_root" => artifact_output_root.to_s,
  "runtime_owned" => true,
  "go_runtime_backed" => true,
  "kde_policy_owner" => false,
  "q4_download_required" => true,
  "q4_extract_required" => true,
  "q4_compile_required" => true,
  "host_compilation_avoided" => true,
  "host_download_avoided" => true,
  "full_smoke_required" => false,
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

remote_step!(
  options.fetch(:local_shell),
  remote_host,
  shell_join(["mkdir", "-p", remote_app_root]),
  options.fetch(:remote_timeout_seconds),
  "q4 Notepad++ remote directory preparation failed"
)

build = run_json_command(
  options.fetch(:local_shell),
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
abort "q4 Runtime build for known portable bundle stage launch did not pass" unless build.fetch("status") == "passed" && bool(build, "remote_build_completed") && bool(build, "host_compilation_avoided")

remote_step!(
  options.fetch(:local_shell),
  remote_host,
  shell_join(["ruby", "-rfileutils", "-e", "abort 'unsafe staging root' unless ARGV[1].start_with?('/home/xnix-', '/tmp/xnix-'); FileUtils.rm_rf(ARGV[1]); FileUtils.mkdir_p(File.dirname(ARGV[0])); File.write(ARGV[0], ARGV[2])", remote_document, remote_staging_root, "Xnix Notepad++ Portable record-first staged launch smoke document"]),
  options.fetch(:remote_timeout_seconds),
  "q4 Notepad++ sample document and staging preparation failed"
)

known_stage_launch_stdout = remote_step!(
  options.fetch(:local_shell),
  remote_host,
  "cd #{Shellwords.escape(remote_source_root)} && #{shell_join(known_bundle_stage_launch_command)}",
  options.fetch(:remote_timeout_seconds),
  "q4 Notepad++ Go Runtime known portable bundle staged launch failed"
)
known_stage_launch = JSON.parse(known_stage_launch_stdout)
known_import = known_stage_launch.fetch("known_portable_bundle_import")
unless known_import.fetch("status") == "passed" &&
       known_import.fetch("request_type") == "windows-known-app-bundle-import-record" &&
       bool(known_import, "checksum_verified") &&
       bool(known_import, "archive_verified") &&
       bool(known_import, "extracted") &&
       bool(known_import, "import_recorded") &&
       known_import.fetch("import_record_request_type") == "external-winapp-bundle-import-record" &&
       known_import.fetch("imported_artifact_kind") == "portable-directory" &&
       bool(known_import, "bundle_manifest_sha256_present")
  warn JSON.pretty_generate(known_import)
  abort "q4 Notepad++ Go Runtime known portable bundle import did not pass"
end
unless known_stage_launch.fetch("status") == "passed" &&
       known_stage_launch.fetch("request_type") == "windows-known-app-bundle-stage-and-launch" &&
       bool(known_stage_launch, "existing_import_record_consumed") &&
       bool(known_stage_launch, "record_first_launch_path") &&
       known_stage_launch.fetch("artifact_kind", "") == "portable-directory" &&
       bool(known_stage_launch, "application_workspace_copied") &&
       known_stage_launch.fetch("application_workspace_mode", "") == "portable-directory" &&
       bool(known_stage_launch, "external_file_bridge_ready") &&
       bool(known_stage_launch, "windows_process_file_argument_window_observed")
  warn JSON.pretty_generate(known_stage_launch)
  abort "q4 Notepad++ Go Runtime known portable bundle staged launch did not pass"
end
known_bundle_gui_packet_stdout = remote_step!(
  options.fetch(:local_shell),
  remote_host,
  "cd #{Shellwords.escape(remote_source_root)} && #{shell_join(known_bundle_gui_evidence_packet_command)}",
  options.fetch(:remote_timeout_seconds),
  "q4 Notepad++ Go Runtime known portable bundle GUI evidence packet failed"
)
known_bundle_gui_packet = JSON.parse(known_bundle_gui_packet_stdout)
unless known_bundle_gui_packet.fetch("report_status") == "passed" &&
       known_bundle_gui_packet.fetch("request_type") == "real-winapp-gui-evidence-packet-preview" &&
       known_bundle_gui_packet.fetch("app_id") == "org.xnix.external.notepadplusplus" &&
       bool(known_bundle_gui_packet, "external_app_run_record_consumed") &&
       bool(known_bundle_gui_packet, "external_app_handle_consumed") &&
       bool(known_bundle_gui_packet, "external_app_import_record_consumed") &&
       bool(known_bundle_gui_packet, "imported_artifact_digest_verified") &&
       bool(known_bundle_gui_packet, "x_window_observed") &&
       bool(known_bundle_gui_packet, "window_observed")
  warn JSON.pretty_generate(known_bundle_gui_packet)
  abort "q4 Notepad++ Go Runtime known portable bundle GUI evidence packet did not pass"
end
known_bundle_gui_packet_local_path = artifact_output_root.join("known-portable-bundle-gui-evidence-packet.json")
known_bundle_gui_packet_fetched = fetch_remote_artifact(
  options.fetch(:local_shell),
  remote_host,
  remote_known_bundle_gui_evidence_packet,
  known_bundle_gui_packet_local_path,
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
remote_bundle_root = "#{remote_state_root}/#{known_stage_launch.fetch("known_portable_bundle_bundle_relative_path")}"
remote_executable = "#{remote_bundle_root}/#{NOTEPADPP_EXE_RELATIVE_PATH}"
remote_import_record = "#{remote_state_root}/#{known_stage_launch.fetch("known_portable_bundle_record_relative_path")}"
plan["notepadpp_downloaded"] = bool(known_import, "downloaded") || known_import.fetch("fetch_cache_status", "") == "verified"
plan["notepadpp_sha256_verified"] = bool(known_import, "checksum_verified")
plan["notepadpp_extracted"] = bool(known_import, "extracted")
plan["known_portable_bundle_import_recorded"] = bool(known_import, "import_recorded")
plan["known_portable_bundle_import_status"] = known_import.fetch("status")
plan["known_portable_bundle_stage_launch_status"] = known_stage_launch.fetch("status")
plan["known_portable_bundle_gui_evidence_packet_status"] = known_bundle_gui_packet.fetch("report_status")
plan["remote_bundle_root"] = remote_bundle_root
plan["remote_executable"] = remote_executable
plan["remote_import_record"] = remote_import_record

staged_command = staged_command(
  remote_host,
  remote_import_record,
  output_path,
  markdown_output_path,
  options.fetch(:remote_timeout_seconds)
)

staged_stdout, staged_stderr, staged_status = run_shell(
  options.fetch(:local_shell),
  shell_join(staged_command),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless staged_status.zero?
  warn staged_stdout unless staged_stdout.empty?
  warn staged_stderr unless staged_stderr.empty?
  abort "q4 Notepad++ Portable staged external run failed"
end

delegated = JSON.parse(staged_stdout)
passed = delegated.fetch("status") == "passed" &&
         bool(delegated, "remote_build_completed") &&
         bool(delegated, "portable_directory_external_app") &&
         bool(delegated, "portable_directory_bundle_import_recorded") &&
         bool(delegated, "existing_import_record_consumed") &&
         bool(delegated, "record_first_launch_path") &&
         delegated.fetch("artifact_kind", "") == "portable-directory" &&
         bool(delegated, "bundle_manifest_sha256_present") &&
         bool(delegated, "application_workspace_copied") &&
         delegated.fetch("application_workspace_mode", "") == "portable-directory" &&
         bool(delegated, "external_file_bridge_ready") &&
         bool(delegated, "windows_process_file_argument_window_observed") &&
         bool(delegated, "go_owned_q4_staged_external_winapp_acceptance_ready") &&
         delegated.fetch("accepted_application_detail_compatibility_state", "") == "runtime-accepted-real-app-run" &&
         delegated.fetch("kde_page_from_accepted_application_detail_compatibility_state", "") == "runtime-accepted-real-app-run"
abort "q4 Notepad++ Portable staged external run did not reach accepted state" unless passed

result = plan.merge(
  "status" => "passed",
  "remote_go_build_status" => build.fetch("status"),
  "remote_build_completed_for_known_import" => bool(build, "remote_build_completed"),
  "remote_build_completed_for_known_stage_launch" => bool(build, "remote_build_completed"),
  "known_portable_bundle_stage_launch_status" => known_stage_launch.fetch("status"),
  "known_portable_bundle_stage_launch_request_type" => known_stage_launch.fetch("request_type"),
  "known_portable_bundle_stage_launch_record_first" => bool(known_stage_launch, "record_first_launch_path"),
  "known_portable_bundle_stage_launch_existing_import_record_consumed" => bool(known_stage_launch, "existing_import_record_consumed"),
  "known_portable_bundle_stage_launch_application_workspace_copied" => bool(known_stage_launch, "application_workspace_copied"),
  "known_portable_bundle_stage_launch_windows_process_file_argument_window_observed" => bool(known_stage_launch, "windows_process_file_argument_window_observed"),
  "known_portable_bundle_gui_evidence_packet_status" => known_bundle_gui_packet.fetch("report_status"),
  "known_portable_bundle_gui_evidence_packet_request_type" => known_bundle_gui_packet.fetch("request_type"),
  "known_portable_bundle_gui_evidence_packet_external_app_run_record_consumed" => bool(known_bundle_gui_packet, "external_app_run_record_consumed"),
  "known_portable_bundle_gui_evidence_packet_external_app_import_record_consumed" => bool(known_bundle_gui_packet, "external_app_import_record_consumed"),
  "known_portable_bundle_gui_evidence_packet_imported_artifact_digest_verified" => bool(known_bundle_gui_packet, "imported_artifact_digest_verified"),
  "known_portable_bundle_gui_evidence_packet_window_observed" => bool(known_bundle_gui_packet, "window_observed"),
  "known_portable_bundle_gui_evidence_packet_artifact_fetched" => known_bundle_gui_packet_fetched,
  "known_portable_bundle_gui_evidence_packet_artifact_output_path" => known_bundle_gui_packet_local_path.to_s,
  "known_portable_bundle_import_status" => known_import.fetch("status"),
  "known_portable_bundle_import_recorded" => bool(known_import, "import_recorded"),
  "known_portable_bundle_import_request_type" => known_import.fetch("request_type"),
  "known_portable_bundle_import_record_request_type" => known_import.fetch("import_record_request_type"),
  "known_portable_bundle_archive_verified" => bool(known_import, "archive_verified"),
  "known_portable_bundle_checksum_verified" => bool(known_import, "checksum_verified"),
  "known_portable_bundle_extracted" => bool(known_import, "extracted"),
  "known_portable_bundle_extracted_file_count" => known_import.fetch("extracted_file_count"),
  "known_portable_bundle_manifest_sha256_present" => bool(known_import, "bundle_manifest_sha256_present"),
  "known_portable_bundle_imported_artifact_kind" => known_import.fetch("imported_artifact_kind"),
  "delegated_status" => delegated.fetch("status"),
  "remote_build_completed" => bool(delegated, "remote_build_completed"),
  "portable_directory_external_app" => bool(delegated, "portable_directory_external_app"),
  "portable_directory_bundle_import_recorded" => bool(delegated, "portable_directory_bundle_import_recorded"),
  "existing_import_record_consumed" => bool(delegated, "existing_import_record_consumed"),
  "record_first_launch_path" => bool(delegated, "record_first_launch_path"),
  "artifact_kind" => delegated.fetch("artifact_kind", ""),
  "bundle_manifest_sha256_present" => bool(delegated, "bundle_manifest_sha256_present"),
  "application_workspace_copied" => bool(delegated, "application_workspace_copied"),
  "application_workspace_mode" => delegated.fetch("application_workspace_mode", ""),
  "external_file_bridge_ready" => bool(delegated, "external_file_bridge_ready"),
  "windows_process_file_argument_window_observed" => bool(delegated, "windows_process_file_argument_window_observed"),
  "acceptance_ready" => bool(delegated, "go_owned_q4_staged_external_winapp_acceptance_ready"),
  "accepted_application_detail_state" => delegated.fetch("accepted_application_detail_compatibility_state", ""),
  "kde_accepted_page_state" => delegated.fetch("kde_page_from_accepted_application_detail_compatibility_state", ""),
  "artifact_fetch_count" => delegated.fetch("artifact_fetch_count", 0)
)
emit_json(result, output_path)

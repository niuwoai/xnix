#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "optparse"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip
NOTEPADPP_SMOKE = PROJECT_ROOT.join("scripts/q4_notepadpp_portable_winapp_smoke.rb")

SCHEMA_VERSION = "xnix.scripts.q4_known_portable_winapp_run.v1"
REQUEST_TYPE = "q4-known-portable-winapp-run"
SUPPORTED_APP_ID = "org.xnix.external.notepadplusplus"
SUPPORTED_DISPLAY_NAME = "Notepad++ Portable"

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
  app: ENV.fetch("XNIX_Q4_KNOWN_PORTABLE_WINAPP_APP", SUPPORTED_APP_ID),
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
  parser.banner = "Usage: ruby scripts/q4_known_portable_winapp_run.rb --app org.xnix.external.notepadplusplus [--execute]"
  parser.on("--execute", "Run a catalog-backed known portable Windows app through the q4 Runtime/KDE lane.") { options[:execute] = true }
  parser.on("--app APP_ID", "Known portable app id, currently: #{SUPPORTED_APP_ID}.") { |value| options[:app] = value }
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
  abort "unsupported known portable app #{clean}; currently supported: #{SUPPORTED_APP_ID}" unless clean == SUPPORTED_APP_ID

  clean
end

def emit_json(payload, output_path)
  text = JSON.pretty_generate(payload) + "\n"
  FileUtils.mkdir_p(output_path.dirname)
  File.write(output_path, text)
  puts text
end

def write_markdown(payload, markdown_output_path)
  lines = [
    "# q4 Known Portable Windows App Run",
    "",
    "- Status: `#{payload.fetch("status")}`",
    "- App: `#{payload.fetch("display_name")}` (`#{payload.fetch("app_id")}`)",
    "- Version: `#{payload.fetch("version")}`",
    "- Operator run ready: `#{payload.fetch("operator_run_ready", false)}`",
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
output_path = ensure_local_output_path!("output path", options.fetch(:output))
markdown_output_path = ensure_local_output_path!("markdown output path", options.fetch(:markdown_output))
remote_materials_root = ensure_remote_xnix_path!("remote materials root", options.fetch(:remote_materials_root))
remote_source_root = ensure_remote_xnix_path!("remote source root", options.fetch(:remote_source_root))
remote_build_root = ensure_remote_xnix_path!("remote build root", options.fetch(:remote_build_root))
artifact_root = PROJECT_ROOT.join("output", "artifacts")
delegated_output_path = artifact_root.join("q4-known-portable-winapp-run-notepadpp-smoke-#{VERSION}.json")
delegated_markdown_output_path = artifact_root.join("q4-known-portable-winapp-run-notepadpp-smoke-#{VERSION}.md")

delegated_command = [
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
delegated_command << "--execute" if execute

plan = {
  "schema_version" => SCHEMA_VERSION,
  "request_type" => REQUEST_TYPE,
  "version" => VERSION,
  "status" => execute ? "running" : "planned",
  "execute" => execute,
  "app_id" => app_id,
  "display_name" => SUPPORTED_DISPLAY_NAME,
  "known_catalog_app" => true,
  "portable_directory_external_app" => true,
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
  "delegated_script" => "scripts/q4_notepadpp_portable_winapp_smoke.rb",
  "delegated_request_type" => "q4-notepadpp-portable-winapp-smoke",
  "delegated_command" => delegated_command,
  "delegated_output_path" => delegated_output_path.to_s,
  "delegated_markdown_output_path" => delegated_markdown_output_path.to_s,
  "output_path" => output_path.to_s,
  "markdown_output_path" => markdown_output_path.to_s,
  "runtime_owned" => true,
  "go_runtime_backed" => true,
  "kde_policy_owner" => false,
  "q4_download_required" => true,
  "q4_extract_required" => true,
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
delegated_args[1] = NOTEPADPP_SMOKE.to_s
stdout, stderr, status = Open3.capture3(*delegated_args, chdir: PROJECT_ROOT.to_s)
unless status.success?
  warn stdout unless stdout.empty?
  warn stderr unless stderr.empty?
  abort "q4 known portable Windows app delegated run failed"
end

delegated = JSON.parse(stdout)
unless delegated.fetch("status") == "passed" &&
       delegated.fetch("app_id") == app_id &&
       delegated.fetch("display_name") == SUPPORTED_DISPLAY_NAME &&
       bool(delegated, "notepadpp_sha256_verified") &&
       bool(delegated, "known_portable_bundle_archive_verified") &&
       bool(delegated, "known_portable_bundle_stage_launch_existing_import_record_consumed") &&
       bool(delegated, "known_portable_bundle_stage_launch_windows_process_file_argument_window_observed") &&
       bool(delegated, "known_portable_bundle_acceptance_ready") &&
       bool(delegated, "known_portable_bundle_acceptance_runtime_accepted_chain_verified") &&
       delegated.fetch("known_portable_bundle_acceptance_launch_source_request_type") == "windows-known-app-bundle-stage-and-launch" &&
       delegated.fetch("accepted_application_detail_state") == "runtime-accepted-real-app-run" &&
       delegated.fetch("kde_accepted_page_state") == "runtime-accepted-real-app-run" &&
       bool(delegated, "host_compilation_avoided") &&
       bool(delegated, "host_download_avoided") &&
       !bool(delegated, "host_root_modified") &&
       !bool(delegated, "privileged_container_required") &&
       !bool(delegated, "host_networking_required") &&
       !bool(delegated, "docker_socket_mounted") &&
       !bool(delegated, "broad_host_mount_required")
  warn stdout
  abort "q4 known portable Windows app run did not reach accepted state"
end

result = plan.merge(
  "status" => "passed",
  "delegated_status" => delegated.fetch("status"),
  "delegated_artifact_fetch_count" => delegated.fetch("artifact_fetch_count"),
  "official_archive_checksum_verified" => bool(delegated, "notepadpp_sha256_verified"),
  "known_portable_bundle_imported" => bool(delegated, "known_portable_bundle_import_recorded"),
  "known_portable_bundle_stage_launch_status" => delegated.fetch("known_portable_bundle_stage_launch_status"),
  "known_portable_bundle_acceptance_request_type" => delegated.fetch("known_portable_bundle_acceptance_request_type"),
  "known_portable_bundle_acceptance_ready" => bool(delegated, "known_portable_bundle_acceptance_ready"),
  "runtime_accepted_chain_verified" => bool(delegated, "known_portable_bundle_acceptance_runtime_accepted_chain_verified"),
  "accepted_application_detail_state" => delegated.fetch("accepted_application_detail_state"),
  "kde_accepted_page_state" => delegated.fetch("kde_accepted_page_state"),
  "operator_run_ready" => true
)
write_markdown(result, markdown_output_path)
emit_json(result, output_path)

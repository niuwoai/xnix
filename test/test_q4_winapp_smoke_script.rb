#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/q4_winapp_smoke.rb")
source = script.read

assert(source.include?("xnix.scripts.q4_winapp_smoke.v1"), "q4 Windows app smoke must expose a stable schema")
assert(source.include?("q4-winapp-smoke"), "q4 Windows app smoke must expose a stable request type")
assert(source.include?("scripts/remote_wine_guest_gui_smoke.rb"), "q4 Windows app smoke must delegate to the maintained q4 GUI smoke")
assert(source.include?("--known-app-id"), "q4 Windows app smoke must support known app ids")
assert(source.include?("--remote-executable"), "q4 Windows app smoke must support remote Windows executables")
assert(source.include?("--owner-file-open"), "q4 Windows app smoke must support owner-controlled file-open runs")
assert(source.include?("--require-real-run-acceptance"), "q4 Windows app smoke must optionally require real-run acceptance")
assert(source.include?("owner_file_open_entrypoint_invoked"), "q4 Windows app smoke must preserve owner file-open entrypoint evidence")
assert(source.include?("document_content_marker_observed"), "q4 Windows app smoke must preserve document content marker evidence")
assert(source.include?("go_owned_q4_winapp_acceptance_document_content_marker_observed"), "q4 Windows app smoke must expose Go-owned document marker acceptance evidence")
assert(!source.include?("--owner-file-open requires --known-app-id"), "q4 Windows app smoke must allow remote executables through owner file-open")
assert(source.include?("q4-winapp-acceptance-preview"), "q4 Windows app smoke must call the Go-owned generic q4 acceptance")
assert(source.include?("go_owned_q4_winapp_acceptance_ready"), "q4 Windows app smoke must expose Go-owned q4 acceptance readiness")
assert(source.include?("go_owned_q4_winapp_acceptance_schema"), "q4 Windows app smoke must expose Go-owned q4 acceptance schema")
assert(source.include?("host_compilation_avoided"), "q4 Windows app smoke must document host compilation avoidance")
assert(!source.include?("go build"), "q4 Windows app smoke wrapper must not compile locally")
assert(!source.include?("go test"), "q4 Windows app smoke wrapper must not test locally")

known_stdout, known_stderr, known_status = Open3.capture3(
  "ruby",
  script.to_s,
  "--known-app-id", "org.xnix.apps.mines",
  "--display-name", "Mines",
  "--window-match", "Mines",
  "--owner-file-open",
  "--sample-file-argument", "mines-save.txt",
  "--require-real-run-acceptance",
  "--output", "/tmp/xnix-q4-winapp-known-plan.json"
)
assert(known_status.success?, "q4 Windows app known-app plan must exit successfully: #{known_stderr}")
known = JSON.parse(known_stdout)
assert(known.fetch("schema_version") == "xnix.scripts.q4_winapp_smoke.v1", "known-app plan must expose schema")
assert(known.fetch("request_type") == "q4-winapp-smoke", "known-app plan must expose request type")
assert(known.fetch("status") == "planned", "known-app plan must not execute by default")
assert(known.fetch("execute") == false, "known-app plan must keep execution disabled by default")
assert(known.fetch("known_app_id") == "org.xnix.apps.mines", "known-app plan must preserve app id")
assert(known.fetch("remote_executable_configured") == false, "known-app plan must not require a remote executable")
assert(known.fetch("launch_mode") == "owner-controlled-launch", "known-app plan must use owner-controlled launch when requested")
assert(known.fetch("file_open_entrypoint_requested") == true, "known-app plan must request file-open entrypoint")
assert(known.fetch("real_run_acceptance_required") == true, "known-app plan must require real-run acceptance when requested")
assert(known.fetch("go_owned_q4_winapp_acceptance_planned") == true, "known-app plan must plan Go-owned q4 acceptance")
assert(known.fetch("go_owned_q4_winapp_acceptance_ready") == false, "known-app plan must not claim Go-owned acceptance before execute")
assert(known.fetch("sample_file_argument") == "mines-save.txt", "known-app plan must preserve sample file argument")
assert(known.fetch("host_compilation_avoided") == true, "known-app plan must avoid host compilation")
assert(known.fetch("host_root_modified") == false, "known-app plan must not mutate the host root")
assert(known.fetch("delegated_command").include?("--known-app-id"), "known-app plan must delegate known app id")
assert(known.fetch("delegated_command").include?("--file-open-entrypoint"), "known-app plan must delegate file-open entrypoint")

remote_stdout, remote_stderr, remote_status = Open3.capture3(
  "ruby",
  script.to_s,
  "--remote-executable", "/home/xnix-run-materials/apps/demo-tool.exe",
  "--app-id", "org.example.demo-tool",
  "--display-name", "Demo Tool",
  "--window-match", "Demo Tool",
  "--remote-file-argument", "/home/xnix-run-materials/docs/demo.txt",
  "--no-sync-source"
)
assert(remote_status.success?, "q4 Windows app remote-executable plan must exit successfully: #{remote_stderr}")
remote = JSON.parse(remote_stdout)
assert(remote.fetch("status") == "planned", "remote-executable plan must not execute by default")
assert(remote.fetch("known_app_id") == "", "remote-executable plan must not pretend to use known app ids")
assert(remote.fetch("remote_executable_configured") == true, "remote-executable plan must record executable configuration")
assert(remote.fetch("remote_executable_path_exposed") == false, "remote-executable plan must hide the executable path from safe fields")
assert(remote.fetch("remote_file_argument_configured") == true, "remote-executable plan must record file argument configuration")
assert(remote.fetch("remote_file_argument_path_exposed") == false, "remote-executable plan must hide raw file argument paths from safe fields")
assert(remote.fetch("source_sync_planned") == false, "remote-executable plan must forward no-sync-source")
assert(remote.fetch("launch_mode") == "direct", "remote-executable plan must default to direct launch")
assert(remote.fetch("go_owned_q4_winapp_acceptance_planned") == true, "remote-executable plan must plan Go-owned q4 acceptance")
assert(remote.fetch("go_owned_q4_winapp_acceptance_ready") == false, "remote-executable plan must not claim Go-owned acceptance before execute")
assert(remote.fetch("host_compilation_avoided") == true, "remote-executable plan must avoid host compilation")
assert(remote.fetch("delegated_command").include?("--remote-executable"), "remote-executable plan must delegate remote executable")
assert(remote.fetch("delegated_command").include?("--remote-file-argument"), "remote-executable plan must delegate remote file argument")
assert(!remote.fetch("delegated_command").include?("--execute"), "remote-executable plan must not delegate execute by default")

remote_owner_stdout, remote_owner_stderr, remote_owner_status = Open3.capture3(
  "ruby",
  script.to_s,
  "--remote-executable", "/home/xnix-run-materials/apps/demo-tool.exe",
  "--app-id", "org.example.demo-tool",
  "--display-name", "Demo Tool",
  "--window-match", "Demo Tool",
  "--sample-file-argument", "demo-document.txt",
  "--owner-file-open",
  "--require-real-run-acceptance"
)
assert(remote_owner_status.success?, "q4 Windows app remote-executable owner file-open plan must exit successfully: #{remote_owner_stderr}")
remote_owner = JSON.parse(remote_owner_stdout)
assert(remote_owner.fetch("known_app_id") == "", "remote-executable owner plan must not require known app ids")
assert(remote_owner.fetch("remote_executable_configured") == true, "remote-executable owner plan must keep the remote executable")
assert(remote_owner.fetch("launch_mode") == "owner-controlled-launch", "remote-executable owner plan must use owner-controlled launch")
assert(remote_owner.fetch("file_open_entrypoint_requested") == true, "remote-executable owner plan must request the file-open entrypoint")
assert(remote_owner.fetch("real_run_acceptance_required") == true, "remote-executable owner plan must require real-run acceptance")
assert(remote_owner.fetch("sample_file_argument") == "demo-document.txt", "remote-executable owner plan must preserve the sample file name")
assert(remote_owner.fetch("delegated_command").include?("--remote-executable"), "remote-executable owner plan must delegate the executable")
assert(remote_owner.fetch("delegated_command").include?("--file-open-entrypoint"), "remote-executable owner plan must delegate file-open entrypoint")
assert(remote_owner.fetch("delegated_command").include?("--sample-file-argument"), "remote-executable owner plan must delegate sample file")

_stdout, stderr, status = Open3.capture3("ruby", script.to_s, "--window-match", "Demo")
assert(!status.success?, "q4 Windows app smoke must reject missing app selectors")
assert(stderr.include?("use exactly one of --known-app-id or --remote-executable"), "q4 Windows app smoke must explain missing selector")

_stdout, stderr, status = Open3.capture3(
  "ruby",
  script.to_s,
  "--known-app-id", "org.xnix.apps.mines",
  "--remote-executable", "/home/xnix-run-materials/apps/demo-tool.exe",
  "--window-match", "Demo"
)
assert(!status.success?, "q4 Windows app smoke must reject conflicting app selectors")
assert(stderr.include?("use exactly one of --known-app-id or --remote-executable"), "q4 Windows app smoke must explain conflicting selector")

puts "PASS: q4 Windows app smoke script is generic and q4-first"

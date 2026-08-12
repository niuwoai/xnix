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
script = project_root.join("scripts/q4_external_winapp_run.rb")
source = script.read

assert(source.include?("xnix.scripts.q4_external_winapp_run.v1"), "q4 external app run must expose a stable schema")
assert(source.include?("q4-external-winapp-run"), "q4 external app run must expose a stable request type")
assert(source.include?("scripts/q4_staged_desktop_external_winapp_smoke.rb"), "q4 external app run must delegate to staged external Runtime/KDE path")
assert(source.include?("local executable must have an MZ header"), "q4 external app run must verify the MZ header before upload")
assert(source.include?("Digest::SHA256.file"), "q4 external app run must hash the uploaded executable")
assert(source.include?("scp"), "q4 external app run must upload the local executable to q4")
assert(source.include?("/home/xnix-run-materials"), "q4 external app run must default to scoped q4 materials")
assert(source.include?("/tmp/xnix-"), "q4 external app run must allow scoped temporary local and remote paths")
assert(source.include?("remote_upload_completed"), "q4 external app run must report upload completion")
assert(source.include?("windows_process_file_argument_window_observed"), "q4 external app run must require file-open window evidence")
assert(source.include?("runtime-accepted-real-app-run"), "q4 external app run must require accepted Runtime/KDE state")
assert(source.include?("host_compilation_avoided"), "q4 external app run must avoid host compilation")
assert(source.include?("q4-external-winapp-run-plan-preview"), "q4 external app run must advertise its Go Runtime run-plan entrypoint")

stdout, stderr, status = Open3.capture3(
  "ruby",
  script.to_s,
  "--executable", "/tmp/xnix-upload-example/example.exe",
  "--window-match", "Example Document",
  "--output", "/tmp/xnix-q4-external-run-plan.json",
  "--markdown-output", "/tmp/xnix-q4-external-run-plan.md"
)
assert(status.success?, "q4 external app run plan must exit successfully: #{stderr}")
payload = JSON.parse(stdout)
assert(payload.fetch("schema_version") == "xnix.scripts.q4_external_winapp_run.v1", "plan must expose schema")
assert(payload.fetch("request_type") == "q4-external-winapp-run", "plan must expose request type")
assert(payload.fetch("status") == "planned", "plan must not execute by default")
assert(payload.fetch("execute") == false, "plan must keep execute disabled by default")
assert(payload.fetch("remote_host") == "root@q4", "plan must default to q4")
assert(payload.fetch("local_executable_configured") == true, "plan must record local executable configuration")
assert(payload.fetch("local_executable_path_exposed") == false, "plan must hide local executable path from safe fields")
assert(payload.fetch("local_executable_sha256") == "", "plan must not hash a missing executable before execute")
assert(payload.fetch("local_executable_mz_header_checked") == false, "plan must not claim MZ validation before execute")
assert(payload.fetch("remote_upload_planned") == true, "plan must upload the executable")
assert(payload.fetch("remote_upload_completed") == false, "plan must not claim upload before execute")
assert(payload.fetch("remote_executable_configured") == true, "plan must configure a remote executable")
assert(payload.fetch("remote_executable_path_exposed_only_for_operator") == true, "plan must keep remote path operator-only")
assert(payload.fetch("app_id") == "org.xnix.external.uploaded", "plan must use default app id")
assert(payload.fetch("display_name") == "Uploaded Windows App", "plan must use default display name")
assert(payload.fetch("window_match") == "Example Document", "plan must preserve required window match")
assert(payload.fetch("delegated_command").include?("scripts/q4_staged_desktop_external_winapp_smoke.rb"), "plan must delegate to staged external smoke")
assert(payload.fetch("delegated_command").include?("--fixture"), "plan must pass fixture")
assert(payload.fetch("delegated_command").include?("external"), "plan must select external fixture")
assert(payload.fetch("delegated_command").include?("--remote-executable"), "plan must pass remote executable")
assert(payload.fetch("delegated_command").include?("--window-match"), "plan must pass required window match")
assert(payload.fetch("go_runtime_plan_required") == true, "plan must require Go Runtime planning")
assert(payload.fetch("go_runtime_plan_schema_version") == "xnix.runtime.q4_external_winapp_run_plan.v1", "plan must advertise Go Runtime run-plan schema")
assert(payload.fetch("go_runtime_plan_request_type") == "q4-external-winapp-run-plan-preview", "plan must advertise Go Runtime run-plan request")
assert(payload.fetch("go_runtime_plan_command").include?("q4-external-winapp-run-plan-preview"), "plan must expose the Go Runtime plan command")
assert(payload.fetch("go_runtime_plan_command").include?("--window-match"), "Go Runtime plan command must require window-match")
assert(payload.fetch("runtime_owned") == true, "plan must keep Runtime ownership")
assert(payload.fetch("go_runtime_backed") == true, "plan must keep Go Runtime backing")
assert(payload.fetch("kde_policy_owner") == false, "plan must not make KDE the policy owner")
assert(payload.fetch("q4_compile_required") == true, "plan must require q4 build for Runtime binaries")
assert(payload.fetch("host_compilation_avoided") == true, "plan must avoid host compilation")
assert(payload.fetch("full_smoke_required") == false, "plan must remain targeted")
assert(payload.fetch("host_root_modified") == false, "plan must not mutate host root")
assert(payload.fetch("privileged_container_required") == false, "plan must not require privileged containers")
assert(payload.fetch("host_networking_required") == false, "plan must not require host networking")
assert(payload.fetch("docker_socket_mounted") == false, "plan must not mount Docker socket")
assert(payload.fetch("broad_host_mount_required") == false, "plan must not require broad host mounts")

stdout, stderr, status = Open3.capture3(
  "ruby",
  script.to_s,
  "--executable", "/Users/rocky/Downloads/app.exe",
  "--window-match", "Example"
)
assert(!status.success?, "q4 external app run must reject broad local paths")
assert(stderr.include?("local executable must stay under this checkout or /tmp/xnix-*"), "q4 external app run must explain broad local path rejection")

stdout, stderr, status = Open3.capture3(
  "ruby",
  script.to_s,
  "--executable", "/tmp/xnix-upload-example/example.exe"
)
assert(!status.success?, "q4 external app run must require window-match")
assert(stderr.include?("window match must be non-empty"), "q4 external app run must explain missing window-match")

puts "PASS: q4 external Windows app run script is upload-and-staged-runtime backed"

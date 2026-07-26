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
script = project_root.join("scripts/q4_runtime_run_plan_execution_smoke.rb")
source = script.read

assert(source.include?("xnix.scripts.q4_runtime_run_plan_execution_smoke.v1"), "q4 Runtime run-plan execution smoke must expose a stable schema")
assert(source.include?("q4-runtime-run-plan-execution-smoke"), "q4 Runtime run-plan execution smoke must expose a request type")
assert(source.include?("scripts/remote_go_build.rb"), "q4 Runtime run-plan execution smoke must build the Runtime binary on q4")
assert(source.include?("--goos"), "q4 Runtime run-plan execution smoke must pass GOOS to q4 build")
assert(source.include?("darwin"), "q4 Runtime run-plan execution smoke must support q4 cross-compiled macOS Runtime binaries")
assert(source.include?("arm64"), "q4 Runtime run-plan execution smoke must default to Apple Silicon Runtime binaries")
assert(source.include?("known-app-verified-catalog-preview"), "q4 Runtime run-plan execution smoke must generate the verified catalog through Go Runtime")
assert(source.include?("known-app-verified-catalog-app-execution"), "q4 Runtime run-plan execution smoke must execute the app from the verified catalog through Go Runtime")
assert(source.include?("verified_catalog_to_app_execution_planned"), "q4 Runtime run-plan execution smoke must plan verified-catalog-to-app execution")
assert(source.include?("run_plan_generated_by_runtime"), "q4 Runtime run-plan execution smoke must report Runtime-generated run plans")
assert(source.include?("direct_run_plan_input"), "q4 Runtime run-plan execution smoke must prove no caller-supplied run plan is needed")
assert(source.include?("--execute"), "q4 Runtime run-plan execution smoke must support explicit execution")
assert(source.include?("q4_messagebox_smoke_planned"), "q4 Runtime run-plan execution smoke must target the real q4 MessageBox smoke")
assert(source.include?("actual_windows_app_run_observed"), "q4 Runtime run-plan execution smoke must require actual Windows app evidence")
assert(source.include?("owner_file_open_entrypoint_invoked"), "q4 Runtime run-plan execution smoke must require owner file-open evidence")
assert(source.include?("document_content_marker_observed"), "q4 Runtime run-plan execution smoke must require document marker evidence")
assert(source.include?("go_owned_q4_winapp_acceptance_ready"), "q4 Runtime run-plan execution smoke must require Go-owned q4 acceptance")
assert(source.include?("runtime_binary_path_exposed"), "q4 Runtime run-plan execution smoke must expose the binary path gate")
assert(source.include?("runtime_smoke_command_arguments_exposed"), "q4 Runtime run-plan execution smoke must expose the command argument gate")
assert(source.include?("host_compilation_avoided"), "q4 Runtime run-plan execution smoke must report host compilation avoidance")
assert(source.include?("docs/claude-code-implementation-packages.md") == false, "q4 Runtime run-plan execution smoke must not touch Claude's package document")

stdout, stderr, status = Open3.capture3(
  "ruby",
  script.to_s,
  "--output", "/tmp/xnix-q4-runtime-run-plan-execution-plan.json"
)
assert(status.success?, "q4 Runtime run-plan execution plan must exit successfully: #{stderr}")
payload = JSON.parse(stdout)
assert(payload.fetch("schema_version") == "xnix.scripts.q4_runtime_run_plan_execution_smoke.v1", "plan must expose schema")
assert(payload.fetch("request_type") == "q4-runtime-run-plan-execution-smoke", "plan must expose request type")
assert(payload.fetch("status") == "planned", "plan must not execute by default")
assert(payload.fetch("execute") == false, "plan must keep execute disabled by default")
assert(payload.fetch("remote_host") == "root@q4", "plan must default to q4")
assert(payload.fetch("remote_runtime_build_planned") == true, "plan must build Runtime on q4")
assert(payload.fetch("runtime_binary_goos") == "darwin", "plan must default Runtime host GOOS to darwin")
assert(payload.fetch("runtime_binary_goarch") == "arm64", "plan must default Runtime host GOARCH to arm64")
assert(payload.fetch("runtime_binary_built_on_q4") == false, "plan must not claim a built binary before execute")
assert(payload.fetch("runtime_binary_fetched") == false, "plan must not fetch a binary before execute")
assert(payload.fetch("runtime_binary_path_exposed") == false, "plan must not expose fetched binary paths")
assert(payload.fetch("runtime_command") == "known-app-verified-catalog-app-execution", "plan must use the Runtime app execution command")
assert(payload.fetch("runtime_command_execute_planned") == true, "plan must plan Runtime command execution")
assert(payload.fetch("verified_catalog_to_app_execution_planned") == true, "plan must execute from verified catalog and app id")
assert(payload.fetch("run_plan_generated_by_runtime") == false, "plan must not claim Runtime-generated run plans before execute")
assert(payload.fetch("direct_run_plan_input") == false, "plan must not rely on a caller-supplied run plan")
assert(payload.fetch("go_runtime_entrypoint_invoked") == false, "plan must not invoke Runtime before execute")
assert(payload.fetch("app_id") == "org.xnix.apps.messagebox", "plan must target MessageBox")
assert(payload.fetch("q4_messagebox_smoke_planned") == true, "plan must target q4 MessageBox smoke")
assert(payload.fetch("q4_messagebox_smoke_passed") == false, "plan must not claim q4 smoke before execute")
assert(payload.fetch("actual_windows_app_run_observed") == false, "plan must not claim actual run before execute")
assert(payload.fetch("host_compilation_avoided") == true, "plan must avoid host compilation")
assert(payload.fetch("host_root_modified") == false, "plan must not mutate host root")
assert(payload.fetch("privileged_container_required") == false, "plan must not require privileged containers")
assert(payload.fetch("host_networking_required") == false, "plan must not require host networking")
assert(payload.fetch("docker_socket_mounted") == false, "plan must not mount Docker socket")
assert(payload.fetch("broad_host_mount_required") == false, "plan must not require broad host mounts")

bad_stdout, bad_stderr, bad_status = Open3.capture3(
  "ruby",
  script.to_s,
  "--fetch-root", "/Users/rocky/not-xnix"
)
assert(!bad_status.success?, "q4 Runtime run-plan execution smoke must reject unsafe fetch roots")
assert((bad_stdout + bad_stderr).include?("fetch root must stay under /tmp/xnix-*"), "q4 Runtime run-plan execution smoke must explain unsafe fetch roots")

puts "PASS: q4 Runtime run-plan execution smoke script is q4-built and Go-entrypoint-backed"

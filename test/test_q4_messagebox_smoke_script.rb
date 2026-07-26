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
script = project_root.join("scripts/q4_messagebox_smoke.rb")
source = script.read

assert(source.include?("xnix.scripts.q4_messagebox_smoke.v1"), "q4 MessageBox smoke must expose a stable schema")
assert(source.include?("q4-messagebox-smoke"), "q4 MessageBox smoke must expose a stable request type")
assert(source.include?("test/fixtures/winapp/messagebox"), "q4 MessageBox smoke must build the real MessageBox Windows fixture")
assert(source.include?("GOOS=windows"), "q4 MessageBox smoke must cross-compile a Windows executable")
assert(source.include?("GOARCH=386"), "q4 MessageBox smoke must build the 32-bit fixture for the Wine guest")
assert(source.include?("scripts/q4_winapp_smoke.rb"), "q4 MessageBox smoke must delegate execution through the generic q4 app smoke")
assert(source.include?("go_owned_q4_winapp_acceptance_ready"), "q4 MessageBox smoke must require Go-owned q4 acceptance")
assert(source.include?("host_compilation_avoided"), "q4 MessageBox smoke must document host compilation avoidance")

stdout, stderr, status = Open3.capture3(
  "ruby",
  script.to_s,
  "--output", "/tmp/xnix-q4-messagebox-plan.json"
)
assert(status.success?, "q4 MessageBox plan must exit successfully: #{stderr}")
payload = JSON.parse(stdout)
assert(payload.fetch("schema_version") == "xnix.scripts.q4_messagebox_smoke.v1", "q4 MessageBox plan must expose schema")
assert(payload.fetch("request_type") == "q4-messagebox-smoke", "q4 MessageBox plan must expose request type")
assert(payload.fetch("status") == "planned", "q4 MessageBox plan must not execute by default")
assert(payload.fetch("execute") == false, "q4 MessageBox plan must keep execute disabled by default")
assert(payload.fetch("remote_host") == "root@q4", "q4 MessageBox plan must default to q4")
assert(payload.fetch("fixture_source") == "test/fixtures/winapp/messagebox", "q4 MessageBox plan must identify the fixture source")
assert(payload.fetch("fixture_goos") == "windows", "q4 MessageBox plan must target Windows")
assert(payload.fetch("fixture_goarch") == "386", "q4 MessageBox plan must target 32-bit Windows")
assert(payload.fetch("remote_windows_fixture_build_planned") == true, "q4 MessageBox plan must build the fixture on q4")
assert(payload.fetch("remote_windows_executable_built") == false, "q4 MessageBox plan must not claim a built executable before execute")
assert(payload.fetch("remote_executable_configured") == true, "q4 MessageBox plan must configure a remote executable")
assert(payload.fetch("remote_executable_path_exposed") == false, "q4 MessageBox plan must hide the remote executable path from safe fields")
assert(payload.fetch("app_id") == "org.xnix.apps.messagebox", "q4 MessageBox plan must use the MessageBox app id")
assert(payload.fetch("display_name") == "Xnix MessageBox", "q4 MessageBox plan must use the MessageBox display name")
assert(payload.fetch("window_match") == "Xnix Windows GUI Smoke", "q4 MessageBox plan must match the MessageBox window")
assert(payload.fetch("q4_compile_required") == true, "q4 MessageBox plan must require q4 compilation")
assert(payload.fetch("host_compilation_avoided") == true, "q4 MessageBox plan must avoid host compilation")
assert(payload.fetch("host_root_modified") == false, "q4 MessageBox plan must not mutate the host root")
assert(payload.fetch("privileged_container_required") == false, "q4 MessageBox plan must not require privileged containers")
assert(payload.fetch("host_networking_required") == false, "q4 MessageBox plan must not require host networking")
assert(payload.fetch("docker_socket_mounted") == false, "q4 MessageBox plan must not mount Docker socket")
assert(payload.fetch("broad_host_mount_required") == false, "q4 MessageBox plan must not require broad host mounts")
assert(payload.fetch("delegated_command").include?("scripts/q4_winapp_smoke.rb"), "q4 MessageBox plan must delegate through q4_winapp_smoke")
assert(payload.fetch("delegated_command").include?("--remote-executable"), "q4 MessageBox plan must pass a remote executable")
assert(payload.fetch("delegated_command").include?("org.xnix.apps.messagebox"), "q4 MessageBox plan must pass MessageBox app identity")
assert(!payload.fetch("delegated_command").include?("--known-app-id"), "q4 MessageBox plan must prove the external executable lane")

stdout, stderr, status = Open3.capture3(
  "ruby",
  script.to_s,
  "--remote-source-root", "/Users/rocky/not-remote"
)
assert(!status.success?, "q4 MessageBox plan must reject unsafe remote source roots")
assert(stderr.include?("remote source root must stay under /home/xnix-* or /tmp/xnix-* on q4"), "q4 MessageBox plan must explain unsafe remote source roots")

puts "PASS: q4 MessageBox smoke script is q4-built and generic-acceptance-backed"

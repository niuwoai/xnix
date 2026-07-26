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
script = project_root.join("scripts/q4_dbus_controlled_launch_owner_fixture_smoke.rb")
source = script.read

assert(source.include?("xnix.scripts.q4_dbus_controlled_launch_owner_fixture_smoke.v1"), "q4 D-Bus fixture smoke must expose a stable schema")
assert(source.include?("q4-dbus-controlled-launch-owner-fixture-smoke"), "q4 D-Bus fixture smoke must expose a request type")
assert(source.include?("scripts/remote_go_build.rb"), "q4 D-Bus fixture smoke must build Runtime binaries on q4")
assert(source.include?("--goos"), "q4 D-Bus fixture smoke must build Linux Runtime binaries")
assert(source.include?("linux"), "q4 D-Bus fixture smoke must target Linux for the remote fixture")
assert(source.include?("--container-goarch"), "q4 D-Bus fixture smoke must expose the container GOARCH selector")
assert(source.include?("amd64"), "q4 D-Bus fixture smoke must default to the q4 host-runnable tools image architecture")
assert(source.include?("xnix-builder-tools:"), "q4 D-Bus fixture smoke must use an existing q4 tools image")
assert(source.include?("docker build"), "q4 D-Bus fixture smoke must build a q4 tools image when the existing image is missing or wrong-arch")
assert(source.include?("dbus-tools"), "q4 D-Bus fixture smoke must default to the lightweight D-Bus tools Dockerfile target")
assert(source.include?("python:3.12-slim"), "q4 D-Bus fixture smoke must default to a q4-local base image")
assert(source.include?("XNIX_TOOLS_BASE_IMAGE"), "q4 D-Bus fixture smoke must pass the q4 tools base image as a build arg")
assert(source.include?("--platform"), "q4 D-Bus fixture smoke must pin the tools image platform")
assert(source.include?("docker create"), "q4 D-Bus fixture smoke must create a restricted q4 container")
assert(source.include?("--network none"), "q4 D-Bus fixture smoke must disable container networking")
assert(source.include?("--read-only"), "q4 D-Bus fixture smoke must keep the container root read-only")
assert(source.include?("--cap-drop ALL"), "q4 D-Bus fixture smoke must drop Linux capabilities")
assert(source.include?("no-new-privileges"), "q4 D-Bus fixture smoke must disable privilege escalation")
assert(source.include?("docker cp"), "q4 D-Bus fixture smoke must copy materials instead of bind mounting host directories")
assert(source.include?("docker commit"), "q4 D-Bus fixture smoke must prepare a copied-image layer before read-only execution")
assert(source.include?("docker rmi"), "q4 D-Bus fixture smoke must clean up the temporary prepared image")
assert(!source.include?("--mount"), "q4 D-Bus fixture smoke must not mount host or Docker paths")
assert(source.include?("pkg-config --cflags --libs gio-2.0"), "q4 D-Bus fixture smoke must compile the C D-Bus adapter inside the q4 tools container")
assert(source.include?("scripts/dbus_controlled_launch_owner_fixture_smoke.rb"), "q4 D-Bus fixture smoke must run the existing D-Bus owner fixture")
assert(source.include?("--gui-smoke-evidence-file"), "q4 D-Bus fixture smoke must pass app-execution evidence into the D-Bus fixture")
assert(source.include?("--runtime-command"), "q4 D-Bus fixture smoke must pass the q4-built Runtime command into the D-Bus fixture")
assert(source.include?("--desktop-entry-file"), "q4 D-Bus fixture smoke must pass KDE desktop action metadata into the D-Bus fixture")
assert(source.include?("desktop_action_metadata_consumed"), "q4 D-Bus fixture smoke must report desktop action metadata consumption")
assert(source.include?("desktop_action_metadata_path_exposed"), "q4 D-Bus fixture smoke must keep desktop action paths hidden")
assert(source.include?("PASS: D-Bus controlled launch owner fixture smoke"), "q4 D-Bus fixture smoke must require the D-Bus fixture PASS marker")
assert(source.include?("docker_socket_mounted"), "q4 D-Bus fixture smoke must report the Docker socket gate")
assert(source.include?("broad_host_mount_required"), "q4 D-Bus fixture smoke must report the broad mount gate")
assert(source.include?("host_compilation_avoided"), "q4 D-Bus fixture smoke must report host compilation avoidance")
assert(source.include?("docs/claude-code-implementation-packages.md") == false, "q4 D-Bus fixture smoke must not touch Claude's package document")

stdout, stderr, status = Open3.capture3(
  "ruby",
  script.to_s,
  "--output", "/tmp/xnix-q4-dbus-controlled-launch-owner-fixture-plan.json"
)
assert(status.success?, "q4 D-Bus fixture smoke plan must exit successfully: #{stderr}")
payload = JSON.parse(stdout)
assert(payload.fetch("schema_version") == "xnix.scripts.q4_dbus_controlled_launch_owner_fixture_smoke.v1", "plan must expose schema")
assert(payload.fetch("request_type") == "q4-dbus-controlled-launch-owner-fixture-smoke", "plan must expose request type")
assert(payload.fetch("status") == "planned", "plan must not execute by default")
assert(payload.fetch("execute") == false, "plan must keep execute disabled by default")
assert(payload.fetch("remote_host") == "root@q4", "plan must default to q4")
assert(payload.fetch("container_goarch") == "amd64", "plan must default to the q4 host-runnable tools image architecture")
assert(payload.fetch("tools_target") == "dbus-tools", "plan must default to the lightweight D-Bus tools target")
assert(payload.fetch("tools_base_image") == "python:3.12-slim", "plan must default to the q4-local base image")
assert(payload.fetch("tools_image_build_planned") == true, "plan must build or verify the q4 tools image")
assert(payload.fetch("tools_image_built_on_q4") == false, "plan must not claim q4 tools image build before execute")
assert(payload.fetch("desktop_action_metadata_consumed") == false, "plan must not claim desktop action consumption before execute")
assert(payload.fetch("desktop_action_metadata_path_exposed") == false, "plan must not expose desktop action metadata paths")
assert(payload.fetch("app_id") == "org.xnix.apps.messagebox", "plan must target MessageBox by default")
assert(payload.fetch("app_execution_evidence_required") == true, "plan must require app-execution evidence")
assert(payload.fetch("app_execution_evidence_copied_to_q4") == false, "plan must not copy evidence before execute")
assert(payload.fetch("app_execution_evidence_path_exposed") == false, "plan must not expose evidence paths")
assert(payload.fetch("linux_runtime_build_planned") == true, "plan must build Linux Runtime binaries on q4")
assert(payload.fetch("linux_runtime_built_on_q4") == false, "plan must not claim q4 Linux binaries before execute")
assert(payload.fetch("dbus_adapter_compiled_in_q4_container") == false, "plan must not claim C adapter compilation before execute")
assert(payload.fetch("dbus_fixture_executed") == false, "plan must not execute before --execute")
assert(payload.fetch("dbus_fixture_passed") == false, "plan must not pass before --execute")
assert(payload.fetch("container_network_none") == true, "plan must require no container networking")
assert(payload.fetch("container_read_only") == true, "plan must require a read-only container root")
assert(payload.fetch("container_cap_drop_all") == true, "plan must drop all container capabilities")
assert(payload.fetch("docker_socket_mounted") == false, "plan must not mount Docker socket")
assert(payload.fetch("broad_host_mount_required") == false, "plan must not require broad host mounts")
assert(payload.fetch("host_networking_required") == false, "plan must not require host networking")
assert(payload.fetch("privileged_container_required") == false, "plan must not require privileged containers")
assert(payload.fetch("host_compilation_avoided") == true, "plan must avoid host compilation")
assert(payload.fetch("host_root_modified") == false, "plan must not mutate host root")

bad_stdout, bad_stderr, bad_status = Open3.capture3(
  "ruby",
  script.to_s,
  "--remote-materials-root", "/Users/rocky/not-q4"
)
assert(!bad_status.success?, "q4 D-Bus fixture smoke must reject unsafe remote materials roots")
assert((bad_stdout + bad_stderr).include?("remote materials root must stay under /home/xnix-* or /tmp/xnix-*"), "q4 D-Bus fixture smoke must explain unsafe remote materials roots")

puts "PASS: q4 D-Bus controlled launch owner fixture smoke script is q4-first and restricted"

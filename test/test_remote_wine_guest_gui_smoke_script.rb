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
script = project_root.join("scripts/remote_wine_guest_gui_smoke.rb")

stdout, stderr, status = Open3.capture3("ruby", script.to_s, chdir: project_root.to_s)
assert(status.success?, "remote Wine guest GUI smoke plan must succeed: #{stderr}")

payload = JSON.parse(stdout)
assert(payload["schema_version"] == "xnix.scripts.remote_wine_guest_gui_smoke.v1", "remote GUI smoke must expose a stable schema")
assert(payload["status"] == "planned", "remote GUI smoke must be planned by default")
assert(payload["execute"] == false, "remote GUI smoke must not execute without --execute")
assert(payload["remote_host"] == "root@q4", "remote GUI smoke must default to q4")
assert(payload["backend"] == "qemu-guest-wine-x11", "remote GUI smoke must target QEMU guest Wine X11")
assert(payload["gui_app_name"] == "winemine.exe", "remote GUI smoke must use a real GUI Windows app")
assert(payload["requires_rebuilt_wine_guest_with_x11"] == true, "remote GUI smoke must document the rebuilt Wine guest requirement")
assert(payload["remote_timeout_seconds"] == 300, "remote GUI smoke must expose a bounded remote timeout")
assert(payload["remote_source_root"].start_with?("/home/xnix-"), "remote source root must stay under /home/xnix-*")
assert(payload["report_output"].start_with?("/home/xnix-"), "remote report output must stay under /home/xnix-*")
assert(payload["state_root"].start_with?("/home/xnix-"), "remote state root must stay under /home/xnix-*")
assert(payload["privileged_container_required"] == false, "remote GUI smoke must not require privileged containers")
assert(payload["host_networking_required"] == false, "remote GUI smoke must not require host networking")
assert(payload["docker_socket_mounted"] == false, "remote GUI smoke must not mount the Docker socket")
assert(payload["broad_host_mount_required"] == false, "remote GUI smoke must not require broad host mounts")
assert(payload["host_root_modified"] == false, "remote GUI smoke must not mutate the host root")

bad_stdout, bad_stderr, bad_status = Open3.capture3(
  "ruby", script.to_s,
  "--remote-source-root", "/tmp/not-xnix",
  chdir: project_root.to_s
)
assert(!bad_status.success?, "remote GUI smoke must reject source roots outside /home/xnix-*")
assert((bad_stdout + bad_stderr).include?("remote source root must stay under /home/xnix-*"), "remote GUI smoke must explain unsafe source roots")

puts "PASS: remote Wine guest GUI smoke script plan"

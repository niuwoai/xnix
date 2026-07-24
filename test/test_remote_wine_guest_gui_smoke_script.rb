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
assert(payload["remote_gui_executable_configured"] == false, "remote GUI smoke must not configure an executable by default")
assert(payload["requires_rebuilt_wine_guest_with_x11"] == true, "remote GUI smoke must document the rebuilt Wine guest requirement")
assert(payload["remote_timeout_seconds"] == 300, "remote GUI smoke must expose a bounded remote timeout")
assert(payload["runtime_build_planned"] == true, "remote GUI smoke must build the Go Runtime before execution")
assert(payload["launch_mode"] == "direct", "remote GUI smoke must default to direct launch mode")
assert(payload["owner_controlled_launch_requested"] == false, "remote GUI smoke direct mode must not request owner-controlled launch")
assert(payload["owner_build_planned"] == false, "remote GUI smoke direct mode must not build the Runtime owner")
assert(payload["launcher_build_planned"] == false, "remote GUI smoke direct mode must not build the managed launcher")
assert(payload["source_sync_mode"] == "runtime", "remote GUI smoke must default to Runtime-only source sync")
assert(payload["source_sync_entry_count"] == 7, "remote GUI smoke Runtime-only sync must include the minimal source entries")
%w[VERSION go.mod cmd internal runtime scripts lib].each do |entry|
  assert(payload["source_sync_entries"].include?(entry), "remote GUI smoke Runtime-only sync must include #{entry}")
end
assert(payload["remote_runtime_bin"].end_with?("/bin/xnix-runtime-go"), "remote GUI smoke must expose the managed Runtime binary location")
assert(payload["remote_owner_bin"].end_with?("/bin/xnix-runtime-owner"), "remote GUI smoke must expose the managed Runtime owner binary location")
assert(payload["remote_launcher_bin"].end_with?("/bin/xnix-compat-launch"), "remote GUI smoke must expose the managed launcher binary location")
assert(payload["remote_known_app_cache_root"].start_with?("/home/xnix-"), "remote known-app cache root must stay under /home/xnix-*")
assert(payload["remote_source_root"].start_with?("/home/xnix-"), "remote source root must stay under /home/xnix-*")
assert(payload["remote_source_root"].include?("runtime"), "remote source root must reflect the Runtime-only sync mode")
assert(payload["report_output"].start_with?("/home/xnix-"), "remote report output must stay under /home/xnix-*")
assert(payload["evidence_output"].start_with?("/home/xnix-"), "remote evidence output must stay under /home/xnix-*")
assert(payload["evidence_preview_planned"] == true, "remote GUI smoke must project Runtime GUI evidence after a pass")
assert(payload["evidence_app_id"] == "org.xnix.apps.mines", "remote GUI smoke must expose the default evidence app id")
assert(payload["evidence_display_name"] == "Mines", "remote GUI smoke must expose the default evidence display name")
assert(payload["state_root"].start_with?("/home/xnix-"), "remote state root must stay under /home/xnix-*")
assert(payload["privileged_container_required"] == false, "remote GUI smoke must not require privileged containers")
assert(payload["host_networking_required"] == false, "remote GUI smoke must not require host networking")
assert(payload["docker_socket_mounted"] == false, "remote GUI smoke must not mount the Docker socket")
assert(payload["broad_host_mount_required"] == false, "remote GUI smoke must not require broad host mounts")
assert(payload["host_root_modified"] == false, "remote GUI smoke must not mutate the host root")

owner_stdout, owner_stderr, owner_status = Open3.capture3(
  "ruby", script.to_s,
  "--launch-mode", "owner-controlled-launch",
  chdir: project_root.to_s
)
assert(owner_status.success?, "remote GUI smoke owner-controlled plan must succeed: #{owner_stderr}")
owner_payload = JSON.parse(owner_stdout)
assert(owner_payload["launch_mode"] == "owner-controlled-launch", "remote GUI smoke must expose owner-controlled launch mode")
assert(owner_payload["owner_controlled_launch_requested"] == true, "remote GUI smoke owner mode must request controlled launch")
assert(owner_payload["owner_build_planned"] == true, "remote GUI smoke owner mode must build the Runtime owner")
assert(owner_payload["launcher_build_planned"] == true, "remote GUI smoke owner mode must build the managed launcher")
assert(owner_payload["remote_command"].include?("--launch-mode owner-controlled-launch"), "remote GUI smoke must forward owner launch mode")

bad_launch_mode_stdout, bad_launch_mode_stderr, bad_launch_mode_status = Open3.capture3(
  "ruby", script.to_s,
  "--launch-mode", "unsafe",
  chdir: project_root.to_s
)
assert(!bad_launch_mode_status.success?, "remote GUI smoke must reject unsupported launch modes")
assert((bad_launch_mode_stdout + bad_launch_mode_stderr).include?("launch mode must be direct or owner-controlled-launch"), "remote GUI smoke must explain unsupported launch modes")

bad_stdout, bad_stderr, bad_status = Open3.capture3(
  "ruby", script.to_s,
  "--remote-source-root", "/tmp/not-xnix",
  chdir: project_root.to_s
)
assert(!bad_status.success?, "remote GUI smoke must reject source roots outside /home/xnix-*")
assert((bad_stdout + bad_stderr).include?("remote source root must stay under /home/xnix-* or /tmp/xnix-*"), "remote GUI smoke must explain unsafe source roots")

tmp_stdout, tmp_stderr, tmp_status = Open3.capture3(
  "ruby", script.to_s,
  "--remote-source-root", "/tmp/xnix-build/runtime",
  "--remote-build-root", "/tmp/xnix-build-cache",
  "--report-output", "/tmp/xnix-run-materials/state/report.json",
  "--evidence-output", "/tmp/xnix-run-materials/state/evidence.json",
  "--state-root", "/tmp/xnix-run-materials/state/gui",
  chdir: project_root.to_s
)
assert(tmp_status.success?, "remote GUI smoke must accept constrained /tmp/xnix-* scratch paths: #{tmp_stderr}")
tmp_payload = JSON.parse(tmp_stdout)
assert(tmp_payload["remote_source_root"].start_with?("/tmp/xnix-"), "remote GUI smoke must expose constrained /tmp source roots")
assert(tmp_payload["remote_build_root"].start_with?("/tmp/xnix-"), "remote GUI smoke must expose constrained /tmp build roots")

exe_stdout, exe_stderr, exe_status = Open3.capture3(
  "ruby", script.to_s,
  "--remote-executable", "/home/xnix-run-materials/fixtures/xnix-messagebox-smoke.exe",
  chdir: project_root.to_s
)
assert(exe_status.success?, "remote GUI smoke executable plan must succeed: #{exe_stderr}")
exe_payload = JSON.parse(exe_stdout)
assert(exe_payload["remote_executable"] == "/home/xnix-run-materials/fixtures/xnix-messagebox-smoke.exe", "remote GUI smoke must expose the managed remote executable")
assert(exe_payload["remote_gui_executable_configured"] == true, "remote GUI smoke must mark remote executable configuration")
assert(exe_payload["gui_app_name"] == "xnix-messagebox-smoke.exe", "remote GUI smoke must show the executable basename")

bad_exe_stdout, bad_exe_stderr, bad_exe_status = Open3.capture3(
  "ruby", script.to_s,
  "--remote-executable", "/tmp/not-xnix/xnix-messagebox-smoke.exe",
  chdir: project_root.to_s
)
assert(!bad_exe_status.success?, "remote GUI smoke must reject executables outside /home/xnix-*")
assert((bad_exe_stdout + bad_exe_stderr).include?("remote executable must stay under /home/xnix-* or /tmp/xnix-*"), "remote GUI smoke must explain unsafe executable paths")

bad_evidence_stdout, bad_evidence_stderr, bad_evidence_status = Open3.capture3(
  "ruby", script.to_s,
  "--evidence-output", "/tmp/wine-gui-evidence.json",
  chdir: project_root.to_s
)
assert(!bad_evidence_status.success?, "remote GUI smoke must reject evidence output outside /home/xnix-*")
assert((bad_evidence_stdout + bad_evidence_stderr).include?("evidence output must stay under /home/xnix-* or /tmp/xnix-*"), "remote GUI smoke must explain unsafe evidence output paths")

full_stdout, full_stderr, full_status = Open3.capture3(
  "ruby", script.to_s,
  "--source-sync-mode", "full",
  chdir: project_root.to_s
)
assert(full_status.success?, "remote GUI smoke full source sync plan must succeed: #{full_stderr}")
full_payload = JSON.parse(full_stdout)
assert(full_payload["source_sync_mode"] == "full", "remote GUI smoke must support explicit full source sync")
assert(full_payload["source_sync_entries"] == ["."], "remote GUI smoke full source sync must include the whole checkout")
assert(full_payload["remote_source_root"].include?("full"), "remote source root must reflect full source sync mode")

bad_mode_stdout, bad_mode_stderr, bad_mode_status = Open3.capture3(
  "ruby", script.to_s,
  "--source-sync-mode", "everything",
  chdir: project_root.to_s
)
assert(!bad_mode_status.success?, "remote GUI smoke must reject unsupported source sync modes")
assert((bad_mode_stdout + bad_mode_stderr).include?("source sync mode must be runtime or full"), "remote GUI smoke must explain unsupported source sync modes")

puts "PASS: remote Wine guest GUI smoke script plan"

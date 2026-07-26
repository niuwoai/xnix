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
assert(payload["remote_file_argument"] == "", "remote GUI smoke must not configure file arguments by default")
assert(payload["file_argument_count"] == 0, "remote GUI smoke must default to no file arguments")
assert(payload["sample_file_argument_requested"] == false, "remote GUI smoke must not create sample files by default")
assert(payload["window_match"] == "", "remote GUI smoke must not require a window match by default")
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
assert(payload["remote_file_open_bin"].end_with?("/bin/xnix-compat-open"), "remote GUI smoke must expose the managed file-open binary location")
assert(payload["remote_known_app_cache_root"].start_with?("/home/xnix-"), "remote known-app cache root must stay under /home/xnix-*")
assert(payload["remote_source_root"].start_with?("/home/xnix-"), "remote source root must stay under /home/xnix-*")
assert(payload["remote_source_root"].include?("runtime"), "remote source root must reflect the Runtime-only sync mode")
assert(payload["report_output"].start_with?("/home/xnix-"), "remote report output must stay under /home/xnix-*")
assert(payload["evidence_output"].start_with?("/home/xnix-"), "remote evidence output must stay under /home/xnix-*")
assert(payload["evidence_preview_planned"] == true, "remote GUI smoke must project Runtime GUI evidence after a pass")
assert(payload["kde_page_output"].start_with?("/home/xnix-"), "remote KDE page output must stay under /home/xnix-*")
assert(payload["kde_action_output"].start_with?("/home/xnix-"), "remote KDE action output must stay under /home/xnix-*")
assert(payload["kde_center_page_preview_planned"] == true, "remote GUI smoke must plan KDE center page evidence after a pass")
assert(payload["kde_controlled_launch_action_preview_planned"] == false, "remote GUI smoke direct mode must not plan a controlled-launch action preview")
assert(script.read.include?("xnix.scripts.remote_wine_guest_gui_smoke.execute_result.v1"), "remote GUI smoke must expose an execute-result summary schema")
assert(script.read.include?("remote_build_completed"), "remote GUI smoke execute result must expose remote build completion")
assert(script.read.include?("evidence_output_written"), "remote GUI smoke execute result must expose Runtime evidence output")
assert(script.read.include?("kde_page_output_written"), "remote GUI smoke execute result must expose KDE page output")
assert(script.read.include?("kde_page_known_app_gui_evidence_count"), "remote GUI smoke execute result must expose KDE GUI evidence consumption")
assert(script.read.include?("runtime_evidence_report_consumed"), "remote GUI smoke execute result must expose Runtime report consumption")
assert(script.read.include?("runtime_evidence_owner_file_open_verified"), "remote GUI smoke execute result must expose Runtime owner file-open evidence")
assert(script.read.include?("runtime_evidence_owner_file_open_entrypoint_invoked"), "remote GUI smoke execute result must expose Runtime owner file-open entrypoint evidence")
assert(script.read.include?("kde_page_known_app_owner_file_open_verified_count"), "remote GUI smoke execute result must expose KDE owner file-open evidence consumption")
assert(script.read.include?("kde_page_known_app_owner_file_open_entrypoint_count"), "remote GUI smoke execute result must expose KDE owner file-open entrypoint evidence consumption")
assert(script.read.include?("kde_action_owner_file_open_verified"), "remote GUI smoke execute result must expose KDE action owner file-open evidence")
assert(script.read.include?("kde_action_owner_file_open_entrypoint_invoked"), "remote GUI smoke execute result must expose KDE action owner file-open entrypoint evidence")
assert(script.read.include?("kde_action_owner_delegated_window_match"), "remote GUI smoke execute result must expose KDE action safe window-match evidence")
assert(script.read.include?("file_arguments_passed"), "remote GUI smoke execute result must expose file argument handoff evidence")
assert(script.read.include?("window_match_observed"), "remote GUI smoke execute result must expose window-match evidence")
assert(script.read.include?("owner_delegated_file_arguments_passed"), "remote GUI smoke execute result must expose owner delegated file argument handoff evidence")
assert(script.read.include?("owner_delegated_file_argument_winepath_translated"), "remote GUI smoke execute result must expose owner delegated Wine path translation evidence")
assert(script.read.include?("owner_delegated_window_match_observed"), "remote GUI smoke execute result must expose owner delegated window-match evidence")
assert(script.read.include?("file_open_entrypoint_requested"), "remote GUI smoke execute result must expose file-open entrypoint selection")
assert(script.read.include?("owner_file_open_entrypoint_invoked"), "remote GUI smoke execute result must expose owner file-open entrypoint invocation")
assert(payload["known_app_id"] == "", "remote GUI smoke must not force known app selection by default")
assert(payload["known_app_selection_planned"] == false, "remote GUI smoke must keep known app selection explicit")
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
assert(owner_payload["file_open_entrypoint_requested"] == false, "remote GUI smoke owner mode must not use the file-open entrypoint unless requested")
assert(owner_payload["remote_command"].include?("--launch-mode owner-controlled-launch"), "remote GUI smoke must forward owner launch mode")
assert(owner_payload["evidence_app_id"] == "org.xnix.apps.mines", "remote GUI smoke owner mode must default evidence identity to Mines")
assert(owner_payload["kde_controlled_launch_action_preview_planned"] == true, "remote GUI smoke owner mode must plan KDE controlled-launch action evidence")
assert(owner_payload["kde_action_state_root"].end_with?("/owner-controlled-launch-state"), "remote GUI smoke owner mode must expose the owner action state root")

file_open_owner_stdout, file_open_owner_stderr, file_open_owner_status = Open3.capture3(
  "ruby", script.to_s,
  "--launch-mode", "owner-controlled-launch",
  "--file-open-entrypoint",
  chdir: project_root.to_s
)
assert(file_open_owner_status.success?, "remote GUI smoke file-open owner plan must succeed: #{file_open_owner_stderr}")
file_open_owner_payload = JSON.parse(file_open_owner_stdout)
assert(file_open_owner_payload["file_open_entrypoint_requested"] == true, "remote GUI smoke file-open owner plan must request the file-open entrypoint")
assert(file_open_owner_payload["file_open_build_planned"] == true, "remote GUI smoke file-open owner plan must build xnix-compat-open")
assert(file_open_owner_payload["remote_command"].include?("--file-open-bin"), "remote GUI smoke must forward the file-open entrypoint binary")
assert(file_open_owner_payload["remote_command"].include?("--file-open-registry"), "remote GUI smoke must forward the file-open registry")

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
  "--kde-page-output", "/tmp/xnix-run-materials/state/kde-page.json",
  "--kde-action-output", "/tmp/xnix-run-materials/state/kde-action.json",
  "--state-root", "/tmp/xnix-run-materials/state/gui",
  chdir: project_root.to_s
)
assert(tmp_status.success?, "remote GUI smoke must accept constrained /tmp/xnix-* scratch paths: #{tmp_stderr}")
tmp_payload = JSON.parse(tmp_stdout)
assert(tmp_payload["remote_source_root"].start_with?("/tmp/xnix-"), "remote GUI smoke must expose constrained /tmp source roots")
assert(tmp_payload["remote_build_root"].start_with?("/tmp/xnix-"), "remote GUI smoke must expose constrained /tmp build roots")

known_stdout, known_stderr, known_status = Open3.capture3(
  "ruby", script.to_s,
  "--known-app-id", "org.xnix.apps.mines",
  chdir: project_root.to_s
)
assert(known_status.success?, "remote GUI smoke known-app plan must succeed: #{known_stderr}")
known_payload = JSON.parse(known_stdout)
assert(known_payload["known_app_id"] == "org.xnix.apps.mines", "remote GUI smoke known-app plan must expose the selected app id")
assert(known_payload["known_app_selection_planned"] == true, "remote GUI smoke known-app plan must route app selection through the remote Go Runtime")

file_arg_stdout, file_arg_stderr, file_arg_status = Open3.capture3(
  "ruby", script.to_s,
  "--known-app-id", "org.xnix.sample.notepad",
  "--sample-file-argument", "sample-document.txt",
  "--window-match", "sample-document.txt",
  "--evidence-app-id", "org.xnix.sample.notepad",
  "--evidence-display-name", "Sample Notepad",
  chdir: project_root.to_s
)
assert(file_arg_status.success?, "remote GUI smoke file argument plan must succeed: #{file_arg_stderr}")
file_arg_payload = JSON.parse(file_arg_stdout)
assert(file_arg_payload["known_app_id"] == "org.xnix.sample.notepad", "remote GUI smoke file argument plan must expose Sample Notepad")
assert(file_arg_payload["remote_file_argument"].end_with?("/sample-document.txt"), "remote GUI smoke file argument plan must place the sample document under the smoke state root")
assert(file_arg_payload["sample_file_argument_requested"] == true, "remote GUI smoke file argument plan must mark sample creation")
assert(file_arg_payload["file_argument_count"] == 1, "remote GUI smoke file argument plan must count the sample document")
assert(file_arg_payload["file_argument_delivery"] == "remote-file-to-guest-copy-and-winepath", "remote GUI smoke file argument plan must describe the guest handoff")
assert(file_arg_payload["window_match"] == "sample-document.txt", "remote GUI smoke file argument plan must expose the required X window match")
assert(file_arg_payload["remote_command"].include?("--launch-mode direct"), "remote GUI smoke file argument plan must keep direct launch mode")

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

messagebox_stdout, messagebox_stderr, messagebox_status = Open3.capture3(
  "ruby", script.to_s,
  "--remote-executable", "/home/xnix-run-materials/fixtures/xnix-messagebox-smoke.exe",
  "--launch-mode", "owner-controlled-launch",
  "--evidence-app-id", "org.xnix.apps.messagebox",
  "--evidence-display-name", "Xnix MessageBox",
  chdir: project_root.to_s
)
assert(messagebox_status.success?, "remote GUI smoke MessageBox identity plan must succeed: #{messagebox_stderr}")
messagebox_payload = JSON.parse(messagebox_stdout)
assert(messagebox_payload["remote_gui_executable_configured"] == true, "remote GUI smoke MessageBox identity plan must configure the executable")
assert(messagebox_payload["gui_app_name"] == "xnix-messagebox-smoke.exe", "remote GUI smoke MessageBox identity plan must expose the executable basename")
assert(messagebox_payload["evidence_app_id"] == "org.xnix.apps.messagebox", "remote GUI smoke MessageBox identity plan must expose MessageBox evidence app id")
assert(messagebox_payload["evidence_display_name"] == "Xnix MessageBox", "remote GUI smoke MessageBox identity plan must expose MessageBox evidence display name")
assert(messagebox_payload["remote_command"].include?("--launch-mode owner-controlled-launch"), "remote GUI smoke MessageBox identity plan must keep owner mode")
assert(messagebox_payload["kde_center_page_preview_planned"] == true, "remote GUI smoke MessageBox identity plan must plan KDE page evidence")
assert(messagebox_payload["kde_controlled_launch_action_preview_planned"] == true, "remote GUI smoke MessageBox identity plan must plan KDE action evidence")

bad_exe_stdout, bad_exe_stderr, bad_exe_status = Open3.capture3(
  "ruby", script.to_s,
  "--remote-executable", "/tmp/not-xnix/xnix-messagebox-smoke.exe",
  chdir: project_root.to_s
)
assert(!bad_exe_status.success?, "remote GUI smoke must reject executables outside /home/xnix-*")
assert((bad_exe_stdout + bad_exe_stderr).include?("remote executable must stay under /home/xnix-* or /tmp/xnix-*"), "remote GUI smoke must explain unsafe executable paths")

bad_file_arg_stdout, bad_file_arg_stderr, bad_file_arg_status = Open3.capture3(
  "ruby", script.to_s,
  "--remote-file-argument", "/tmp/not-xnix/sample-document.txt",
  chdir: project_root.to_s
)
assert(!bad_file_arg_status.success?, "remote GUI smoke must reject file arguments outside /home/xnix-*")
assert((bad_file_arg_stdout + bad_file_arg_stderr).include?("remote file argument must stay under /home/xnix-* or /tmp/xnix-*"), "remote GUI smoke must explain unsafe file argument paths")

mixed_file_arg_stdout, mixed_file_arg_stderr, mixed_file_arg_status = Open3.capture3(
  "ruby", script.to_s,
  "--remote-file-argument", "/home/xnix-run-materials/state/sample-document.txt",
  "--sample-file-argument", "sample-document.txt",
  chdir: project_root.to_s
)
assert(!mixed_file_arg_status.success?, "remote GUI smoke must reject mixed file argument sources")
assert((mixed_file_arg_stdout + mixed_file_arg_stderr).include?("use either --remote-file-argument or --sample-file-argument, not both"), "remote GUI smoke must explain mixed file argument sources")

bad_evidence_stdout, bad_evidence_stderr, bad_evidence_status = Open3.capture3(
  "ruby", script.to_s,
  "--evidence-output", "/tmp/wine-gui-evidence.json",
  chdir: project_root.to_s
)
assert(!bad_evidence_status.success?, "remote GUI smoke must reject evidence output outside /home/xnix-*")
assert((bad_evidence_stdout + bad_evidence_stderr).include?("evidence output must stay under /home/xnix-* or /tmp/xnix-*"), "remote GUI smoke must explain unsafe evidence output paths")

bad_kde_page_stdout, bad_kde_page_stderr, bad_kde_page_status = Open3.capture3(
  "ruby", script.to_s,
  "--kde-page-output", "/tmp/wine-gui-kde-page.json",
  chdir: project_root.to_s
)
assert(!bad_kde_page_status.success?, "remote GUI smoke must reject KDE page output outside /home/xnix-*")
assert((bad_kde_page_stdout + bad_kde_page_stderr).include?("KDE page output must stay under /home/xnix-* or /tmp/xnix-*"), "remote GUI smoke must explain unsafe KDE page output paths")

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

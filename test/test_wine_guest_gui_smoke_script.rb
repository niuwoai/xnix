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
script = project_root.join("scripts/wine_guest_gui_smoke.rb")
script_source = script.read
defconfig = project_root.join("buildroot/configs/xnix_wine_i386_defconfig").read
sshd_config = project_root.join("buildroot/board/xnix/rootfs-overlay/etc/ssh/sshd_config").read

required_symbols = %w[
  BR2_PACKAGE_XORG7=y
  BR2_PACKAGE_XLIB_LIBX11=y
  BR2_PACKAGE_XLIB_LIBXCOMPOSITE=y
  BR2_PACKAGE_XLIB_LIBXCURSOR=y
  BR2_PACKAGE_XLIB_LIBXEXT=y
  BR2_PACKAGE_XLIB_LIBXI=y
  BR2_PACKAGE_XLIB_LIBXINERAMA=y
  BR2_PACKAGE_XLIB_LIBXRANDR=y
  BR2_PACKAGE_XLIB_LIBXRENDER=y
  BR2_PACKAGE_XLIB_LIBXXF86VM=y
  BR2_PACKAGE_FONTCONFIG=y
  BR2_PACKAGE_FREETYPE=y
  BR2_PACKAGE_LIBPNG=y
  BR2_PACKAGE_JPEG=y
  BR2_PACKAGE_ZLIB=y
]
required_symbols.each do |symbol|
  assert(defconfig.include?(symbol), "Wine guest defconfig must include GUI dependency #{symbol}")
end

assert(sshd_config.include?("X11Forwarding no"), "Wine guest SSH must keep X11 forwarding disabled by default")
assert(sshd_config.include?("AllowTcpForwarding no"), "Wine guest SSH must keep TCP forwarding disabled by default")
go_gui_smoke_source = project_root.join("internal/runtime/winapp/guest_gui_smoke.go").read
assert(go_gui_smoke_source.include?("\"&\"") && go_gui_smoke_source.include?("printf"), "Go GUI smoke must record the background Wine process id without invalid shell separators")
assert(!script_source.include?("&;"), "GUI smoke must not emit an invalid background shell separator")
assert(script_source.include?("windows-app-guest-wine-gui-smoke"), "GUI smoke must delegate Wine GUI execution to the Go Runtime")
assert(script_source.include?("--executable"), "GUI smoke must expose a local Windows GUI executable delivery path")
assert(script_source.include?("--file-argument"), "GUI smoke must expose file-argument delivery to the Go Runtime")
assert(script_source.include?("--window-match"), "GUI smoke must expose an X window title match gate")
assert(script_source.include?("XNIX_RUNTIME_OWNER_GUI_FILE_ARGUMENTS_JSON"), "GUI smoke owner path must pass file arguments through the Runtime owner boundary")
assert(script_source.include?("XNIX_RUNTIME_OWNER_WINDOW_MATCH"), "GUI smoke owner path must pass window matching through the Runtime owner boundary")
assert(script_source.include?("\"evidence-relative-path\", evidence_relative_path"), "GUI smoke owner path must call xnix-runtime-owner with positional evidence handoff")
assert(go_gui_smoke_source.include?("\"wineboot\"") && go_gui_smoke_source.include?("\"--init\""), "Go Runtime must initialize the Wine prefix before launching the GUI app")

stdout, stderr, status = Open3.capture3("ruby", script.to_s, "--plan-only", "--format", "json", chdir: project_root.to_s)
assert(status.success?, "Wine guest GUI smoke plan must succeed: #{stderr}")

payload = JSON.parse(stdout)
assert(payload["schema_version"] == "xnix.scripts.wine_guest_gui_smoke.v1", "GUI smoke must expose a stable schema")
assert(payload["status"] == "planned", "GUI smoke plan must not execute by default")
assert(payload["execute"] == false, "GUI smoke plan must keep execution disabled")
assert(payload["backend"] == "qemu-guest-wine-x11", "GUI smoke must target the QEMU guest Wine X11 backend")
assert(payload["gui_app_name"] == "winemine.exe", "GUI smoke must use a real Wine GUI Windows app by default")
assert(payload["local_gui_executable_configured"] == false, "GUI smoke plan must default to the in-guest app path")
assert(payload["known_app_id"] == "", "GUI smoke plan must not force known app selection by default")
assert(payload["known_app_selection_planned"] == false, "GUI smoke plan must keep known app selection explicit")
assert(payload["file_argument_count"] == 0, "GUI smoke plan must default to no file arguments")
assert(payload["file_argument_delivery"] == "", "GUI smoke plan must not plan file delivery by default")
assert(payload["window_match"] == "", "GUI smoke plan must not require a window match by default")
assert(payload["window_match_observed"] == false, "GUI smoke plan must not report observed window matches")
assert(payload["raw_file_argument_path_exposed"] == false, "GUI smoke plan must not expose raw file argument paths")
assert(payload["launch_mode"] == "direct", "GUI smoke plan must default to direct launch mode")
assert(payload["owner_controlled_launch_requested"] == false, "GUI smoke direct plan must not request owner-controlled launch")
assert(payload["owner_service_call_planned"] == false, "GUI smoke direct plan must not plan an owner service call")
assert(payload["qemu_user_network_restrict_disabled_for_display"] == true, "GUI smoke must disclose the temporary display networking exception")
assert(payload["loopback_ssh_forwarding_only"] == true, "GUI smoke must keep SSH forwarding loopback-bound")
assert(payload["privileged_container_required"] == false, "GUI smoke must not require privileged containers")
assert(payload["host_networking_required"] == false, "GUI smoke must not require host networking")
assert(payload["docker_socket_mounted"] == false, "GUI smoke must not mount the Docker socket")
assert(payload["broad_host_mount_required"] == false, "GUI smoke must not require broad host mounts")
assert(payload["host_root_modified"] == false, "GUI smoke must not mutate the host root")
assert(payload["wineboot_invoked"] == false, "GUI smoke plan must not invoke wineboot")
assert(payload["runtime_go_owned_gui_smoke"] == true, "GUI smoke must report Go-owned Runtime GUI execution")

known_stdout, known_stderr, known_status = Open3.capture3(
  "ruby", script.to_s,
  "--plan-only",
  "--format", "json",
  "--known-app-id", "org.xnix.apps.mines",
  chdir: project_root.to_s
)
assert(known_status.success?, "Wine guest GUI smoke known-app plan must succeed: #{known_stderr}")
known_payload = JSON.parse(known_stdout)
assert(known_payload["known_app_id"] == "org.xnix.apps.mines", "GUI smoke known-app plan must expose the selected app id")
assert(known_payload["known_app_selection_planned"] == true, "GUI smoke known-app plan must route app selection through the Go Runtime")
assert(known_payload["gui_app_name"] == "winemine.exe", "GUI smoke known-app plan must preserve the Mines GUI app name")

owner_stdout, owner_stderr, owner_status = Open3.capture3(
  "ruby", script.to_s,
  "--plan-only",
  "--format", "json",
  "--launch-mode", "owner-controlled-launch",
  "--launcher-bin", "/home/xnix-build-cache/bin/xnix-compat-launch",
  chdir: project_root.to_s
)
assert(owner_status.success?, "Wine guest GUI smoke owner-controlled plan must succeed: #{owner_stderr}")
owner_payload = JSON.parse(owner_stdout)
assert(owner_payload["launch_mode"] == "owner-controlled-launch", "GUI smoke must expose owner-controlled launch mode")
assert(owner_payload["owner_controlled_launch_requested"] == true, "GUI smoke owner mode must request controlled launch")
assert(owner_payload["owner_service_call_planned"] == true, "GUI smoke owner mode must plan the owner service call")
assert(owner_payload["owner_seed_gui_smoke_planned"] == true, "GUI smoke owner mode must plan seed GUI evidence")
assert(owner_payload["runtime_owner_bin_configured"] == true, "GUI smoke owner mode must configure the Runtime owner binary")
assert(owner_payload["managed_launcher_bin_configured"] == true, "GUI smoke owner mode must configure the managed launcher")
assert(owner_payload["evidence_app_id"] == "org.xnix.apps.mines", "GUI smoke owner mode must default evidence identity to Mines")
assert(owner_payload["evidence_display_name"] == "Mines", "GUI smoke owner mode must default evidence display name to Mines")
assert(owner_payload["owner_external_gui_app_requested"] == false, "GUI smoke owner mode must default to the built-in GUI app")
assert(owner_payload["owner_external_gui_app_path_exposed"] == false, "GUI smoke owner mode must not expose owner guest GUI paths")
assert(owner_payload["owner_external_gui_app_delivery"] == "", "GUI smoke owner mode must not plan external GUI app delivery by default")
assert(owner_payload["owner_delegated_file_argument_count"] == 0, "GUI smoke owner mode must default delegated file argument count to zero")
assert(owner_payload["owner_delegated_file_arguments_passed"] == false, "GUI smoke owner mode must default delegated file handoff evidence to false")
assert(owner_payload["owner_delegated_raw_file_argument_path_exposed"] == false, "GUI smoke owner mode must keep delegated raw file paths hidden")
assert(owner_payload["owner_delegated_window_match"] == "", "GUI smoke owner mode must default delegated window match to empty")
assert(owner_payload["owner_delegated_window_match_observed"] == false, "GUI smoke owner mode must default delegated window match evidence to false")

bad_mode_stdout, bad_mode_stderr, bad_mode_status = Open3.capture3(
  "ruby", script.to_s,
  "--plan-only",
  "--launch-mode", "free-for-all",
  chdir: project_root.to_s
)
assert(!bad_mode_status.success?, "Wine guest GUI smoke must reject unsupported launch modes")
assert((bad_mode_stdout + bad_mode_stderr).include?("unsupported launch mode free-for-all"), "Wine guest GUI smoke must explain unsupported launch modes")

fixture_executable = project_root.join("test/fixtures/winapp/messagebox/xnix-messagebox-smoke.exe")
stdout, stderr, status = Open3.capture3("ruby", script.to_s, "--plan-only", "--format", "json", "--executable", fixture_executable.to_s, chdir: project_root.to_s)
assert(status.success?, "Wine guest GUI smoke executable plan must succeed: #{stderr}")
payload = JSON.parse(stdout)
assert(payload["gui_app_name"] == "xnix-messagebox-smoke.exe", "GUI smoke executable plan must surface the executable basename")
assert(payload["local_gui_executable_configured"] == true, "GUI smoke executable plan must record local executable delivery mode")

fixture_document = project_root.join("test/fixtures/winapp/sample-document.txt")
file_arg_stdout, file_arg_stderr, file_arg_status = Open3.capture3(
  "ruby", script.to_s,
  "--plan-only",
  "--format", "json",
  "--known-app-id", "org.xnix.sample.notepad",
  "--file-argument", fixture_document.to_s,
  "--window-match", "sample-document.txt",
  chdir: project_root.to_s
)
assert(file_arg_status.success?, "Wine guest GUI smoke file argument plan must succeed: #{file_arg_stderr}")
file_arg_payload = JSON.parse(file_arg_stdout)
assert(file_arg_payload["known_app_id"] == "org.xnix.sample.notepad", "GUI smoke file argument plan must expose the selected Notepad app")
assert(file_arg_payload["file_argument_count"] == 1, "GUI smoke file argument plan must count the requested document")
assert(file_arg_payload["file_argument_delivery"] == "guest-copy-and-winepath", "GUI smoke file argument plan must describe guest copy and Wine path translation")
assert(file_arg_payload["window_match"] == "sample-document.txt", "GUI smoke file argument plan must expose the required window match")
assert(file_arg_payload["raw_file_argument_path_exposed"] == false, "GUI smoke file argument plan must keep raw paths out of report fields")

owner_file_arg_stdout, owner_file_arg_stderr, owner_file_arg_status = Open3.capture3(
  "ruby", script.to_s,
  "--plan-only",
  "--format", "json",
  "--launch-mode", "owner-controlled-launch",
  "--launcher-bin", "/home/xnix-build-cache/bin/xnix-compat-launch",
  "--known-app-id", "org.xnix.sample.notepad",
  "--file-argument", fixture_document.to_s,
  "--window-match", "sample-document.txt",
  chdir: project_root.to_s
)
assert(owner_file_arg_status.success?, "Wine guest GUI smoke owner file argument plan must succeed: #{owner_file_arg_stderr}")
owner_file_arg_payload = JSON.parse(owner_file_arg_stdout)
assert(owner_file_arg_payload["owner_controlled_launch_requested"] == true, "GUI smoke owner file argument plan must request owner-controlled launch")
assert(owner_file_arg_payload["file_argument_count"] == 1, "GUI smoke owner file argument plan must count the requested document")
assert(owner_file_arg_payload["file_argument_delivery"] == "guest-copy-and-winepath", "GUI smoke owner file argument plan must describe owner delegated guest copy")
assert(owner_file_arg_payload["window_match"] == "sample-document.txt", "GUI smoke owner file argument plan must preserve the window match")
assert(owner_file_arg_payload["raw_file_argument_path_exposed"] == false, "GUI smoke owner file argument plan must keep raw paths out of report fields")

owner_exe_stdout, owner_exe_stderr, owner_exe_status = Open3.capture3(
  "ruby", script.to_s,
  "--plan-only",
  "--format", "json",
  "--launch-mode", "owner-controlled-launch",
  "--launcher-bin", "/home/xnix-build-cache/bin/xnix-compat-launch",
  "--executable", fixture_executable.to_s,
  chdir: project_root.to_s
)
assert(owner_exe_status.success?, "Wine guest GUI smoke owner executable plan must succeed: #{owner_exe_stderr}")
owner_exe_payload = JSON.parse(owner_exe_stdout)
assert(owner_exe_payload["gui_app_name"] == "xnix-messagebox-smoke.exe", "GUI smoke owner executable plan must expose the executable basename")
assert(owner_exe_payload["owner_external_gui_app_requested"] == true, "GUI smoke owner executable plan must request owner external GUI app launch")
assert(owner_exe_payload["owner_external_gui_app_path_exposed"] == false, "GUI smoke owner executable plan must not expose owner guest GUI paths")
assert(owner_exe_payload["owner_external_gui_app_delivery"] == "owner-managed-copy", "GUI smoke owner executable plan must use owner-managed copy delivery")

messagebox_stdout, messagebox_stderr, messagebox_status = Open3.capture3(
  "ruby", script.to_s,
  "--plan-only",
  "--format", "json",
  "--launch-mode", "owner-controlled-launch",
  "--launcher-bin", "/home/xnix-build-cache/bin/xnix-compat-launch",
  "--executable", fixture_executable.to_s,
  "--evidence-app-id", "org.xnix.apps.messagebox",
  "--evidence-display-name", "Xnix MessageBox",
  chdir: project_root.to_s
)
assert(messagebox_status.success?, "Wine guest GUI smoke MessageBox identity plan must succeed: #{messagebox_stderr}")
messagebox_payload = JSON.parse(messagebox_stdout)
assert(messagebox_payload["gui_app_name"] == "xnix-messagebox-smoke.exe", "GUI smoke MessageBox identity plan must preserve the executable basename")
assert(messagebox_payload["evidence_app_id"] == "org.xnix.apps.messagebox", "GUI smoke MessageBox identity plan must expose MessageBox evidence app id")
assert(messagebox_payload["evidence_display_name"] == "Xnix MessageBox", "GUI smoke MessageBox identity plan must expose MessageBox evidence display name")
assert(messagebox_payload["evidence_app_version"] == project_root.join("VERSION").read.strip, "GUI smoke MessageBox identity plan must default evidence version to the project version")

puts "PASS: Wine guest GUI smoke script plan"

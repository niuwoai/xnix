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
assert(owner_payload["owner_external_gui_app_requested"] == false, "GUI smoke owner mode must default to the built-in GUI app")
assert(owner_payload["owner_external_gui_app_path_exposed"] == false, "GUI smoke owner mode must not expose owner guest GUI paths")
assert(owner_payload["owner_external_gui_app_delivery"] == "", "GUI smoke owner mode must not plan external GUI app delivery by default")

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

puts "PASS: Wine guest GUI smoke script plan"

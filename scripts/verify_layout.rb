#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
EXPECTED_VERSION = "0.2.9"
REQUIRED_FILES = %w[
  Dockerfile
  VERSION
  buildroot/Config.in
  buildroot/external.desc
  buildroot/external.mk
  buildroot/sources.lock
  buildroot/configs/xnix_x86_64_defconfig
  buildroot/board/xnix/rootfs-overlay/etc/ssh/sshd_config
  buildroot/board/xnix/post-build.sh
  lib/xnix/container.rb
  lib/xnix/buildroot.rb
  lib/xnix/qemu.rb
  lib/xnix/serial_log.rb
  lib/xnix/sshd.rb
  lib/xnix/ssh_probe.rb
  lib/xnix/ssh_test_key.rb
  lib/xnix/compatibility/application_recipe.rb
  lib/xnix/compatibility/dbus_runtime_client.rb
  lib/xnix/compatibility/desktop_entry.rb
  lib/xnix/compatibility/dolphin_service_menu.rb
  lib/xnix/compatibility/file_open_request.rb
  lib/xnix/compatibility/kde_center_model.rb
  lib/xnix/compatibility/launch_request.rb
  lib/xnix/compatibility/recipe_store.rb
  lib/xnix/compatibility/runtime_daemon.rb
  lib/xnix/milestone.rb
  bin/xnix-compatd
  bin/xnix-compat-launch
  bin/xnix-compat-open
  bin/xnix-kde-center-model
  libexec/xnix/compatd
  scripts/container.rb
  scripts/dbus_session_smoke.rb
  scripts/kde_center_dbus_smoke.rb
  scripts/fetch_buildroot.rb
  scripts/full_smoke.rb
  scripts/install_runtime_activation.rb
  scripts/prepare_ssh_test_key.rb
  scripts/ssh_smoke.rb
  runtime/dbus/org.xnix.Compatibility1.xml
  runtime/dbus/org.xnix.Compatibility1.service
  runtime/dbus/xnix_compatd_smoke.c
  runtime/recipes/org.xnix.sample.notepad.json
  runtime/systemd/xnix-compatd.service
  kde/plasmoids/org.xnix.compatibilitycenter/metadata.json
  kde/plasmoids/org.xnix.compatibilitycenter/contents/ui/main.qml
  kde/dolphin/servicemenus/xnix-open-with-compatibility.desktop
  docs/compatibility-runtime.md
  test/test_container.rb
  test/test_buildroot.rb
  test/test_qemu.rb
  test/test_serial_log.rb
  test/test_sshd.rb
  test/test_ssh_probe.rb
  test/test_milestone.rb
  test/test_application_recipe.rb
  test/test_recipe_store.rb
  test/test_desktop_entry.rb
  test/test_dbus_runtime_client.rb
  test/test_dolphin_service_menu.rb
  test/test_file_open_request.rb
  test/test_launch_request.rb
  test/test_runtime_contract.rb
  test/test_runtime_daemon.rb
  test/test_runtime_dispatch.rb
  test/test_runtime_activation.rb
  test/test_runtime_activation_install.rb
  test/test_runtime_dbus_smoke_script.rb
  test/test_kde_center_model.rb
].freeze
FORBIDDEN_CONTAINER_TOKENS = ["--privileged", "--network host", "docker.sock"].freeze
REQUIRED_CONFIG_LINES = [
  "BR2_x86_64=y",
  "BR2_SYSTEM_DHCP=\"eth0\"",
  "BR2_TARGET_GENERIC_GETTY_PORT=\"ttyS0\"",
  "BR2_LINUX_KERNEL=y",
  "BR2_TARGET_ROOTFS_INITRAMFS=y",
  "BR2_PACKAGE_OPENSSH_SERVER=y"
].freeze
REQUIRED_VERSION_MARKERS = {
  "CLAUDE.md" => "Version: `#{EXPECTED_VERSION}`",
  "PRODUCT_OVERVIEW.md" => "Current version: v#{EXPECTED_VERSION}",
  "CHANGELOG.md" => "## [#{EXPECTED_VERSION}]",
  "README.md" => "`v#{EXPECTED_VERSION}`"
}.freeze

def read_project_file(relative_path)
  PROJECT_ROOT.join(relative_path).read
end

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

REQUIRED_FILES.each do |relative_path|
  assert(PROJECT_ROOT.join(relative_path).file?, "missing required file: #{relative_path}")
end

assert(read_project_file("VERSION").strip == EXPECTED_VERSION, "VERSION must be #{EXPECTED_VERSION}")

REQUIRED_VERSION_MARKERS.each do |relative_path, marker|
  assert(read_project_file(relative_path).include?(marker), "#{relative_path} must reference #{EXPECTED_VERSION}")
end

source_lock = read_project_file("buildroot/sources.lock")
assert(source_lock.include?("buildroot.version=2025.02.15"), "source lock must pin Buildroot 2025.02.15")
assert(source_lock.match?(/^buildroot.sha256=[0-9a-f]{64}$/), "source lock must contain a SHA-256")

config = read_project_file("buildroot/configs/xnix_x86_64_defconfig")
REQUIRED_CONFIG_LINES.each do |line|
  assert(config.include?(line), "defconfig must include #{line}")
end

dockerfile = read_project_file("Dockerfile")
assert(dockerfile.include?("qemu-system-x86"), "Dockerfile must install QEMU system emulation")
assert(dockerfile.include?("ruby"), "Dockerfile must install Ruby for build utilities")
assert(dockerfile.include?("dbus"), "Dockerfile must install D-Bus tooling for runtime smoke tests")
assert(dockerfile.include?("libglib2.0-bin"), "Dockerfile must install gdbus for runtime smoke tests")
assert(dockerfile.include?("libglib2.0-dev"), "Dockerfile must install GIO headers for runtime smoke tests")
assert(dockerfile.include?("pkg-config"), "Dockerfile must install pkg-config for runtime smoke tests")
assert(dockerfile.include?("USER xnix"), "Dockerfile must run as the non-root xnix user")
FORBIDDEN_CONTAINER_TOKENS.each do |token|
  assert(!dockerfile.include?(token), "Dockerfile must not contain #{token}")
end

plasmoid_metadata = read_project_file("kde/plasmoids/org.xnix.compatibilitycenter/metadata.json")
assert(plasmoid_metadata.include?("\"Version\": \"#{EXPECTED_VERSION}\""), "Plasma metadata must reference #{EXPECTED_VERSION}")
assert(plasmoid_metadata.include?("xnix-kde-center-model"), "Plasma metadata must declare the Runtime read-model command")

dolphin_service_menu = read_project_file("kde/dolphin/servicemenus/xnix-open-with-compatibility.desktop")
assert(dolphin_service_menu.include?("X-KDE-ServiceTypes=KonqPopupMenu/Plugin"), "Dolphin service menu must target Dolphin popup integration")
assert(dolphin_service_menu.include?("Exec=xnix-compat-open %U"), "Dolphin service menu must delegate to the Runtime file-open command")

desktop_entry_source = read_project_file("lib/xnix/compatibility/desktop_entry.rb")
assert(desktop_entry_source.include?("xnix-compat-launch"), "Desktop entries must delegate to the Runtime launcher command")

puts "PASS: Xnix #{EXPECTED_VERSION} scaffold is consistent"

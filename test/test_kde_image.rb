#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "pathname"
require "tempfile"
require_relative "../lib/xnix/image/kde_image"
require_relative "../lib/xnix/image/boot_smoke"
require_relative "../lib/xnix/image/serial_probe"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
image = Xnix::Image::KdeImage.new(project_root: project_root.to_s)

# --- Manifest identity ------------------------------------------------
report = image.to_h
assert(report["version"] == project_root.join("VERSION").read.strip,
       "image report must read the version from the VERSION file, not a hardcoded string")
assert(report["artifact_type"] == "kde-plasma-atomic-image", "image must identify its artifact type")
assert(report["image_name"] == "xnix-kinoite", "image name must be xnix-kinoite")
assert(report["base_image"].start_with?("quay.io/fedora/fedora-kinoite"),
       "image must build on a Fedora Kinoite base")
assert(report["architecture"] == "x86_64", "image must target x86_64")
assert(report["desktop"]["environment"] == "kde-plasma-6", "flagship desktop must be KDE Plasma 6")
assert(report["desktop"]["display_manager"] == "plasmalogin",
       "image must use Fedora 44's Plasma Login Manager")

# --- Provenance / policy ownership ------------------------------------
assert(report["runtime_owned"], "the Compatibility Runtime must own image policy")
assert(!report["kde_policy_owner"], "KDE must not own compatibility policy in the image")
assert(report["privileged_build_required"], "manifest must declare the privileged build requirement")

# --- Validity + referenced files exist --------------------------------
assert(image.valid?, "checked-in manifest must be internally consistent: #{image.problems.join('; ')}")
assert(report["problems"].empty?, "a valid image must report no problems")
image.referenced_sources.each do |rel|
  assert(project_root.join(rel).exist?, "referenced source must exist on disk: #{rel}")
end
assert(!image.referenced_sources.include?("runtime/systemd/xnix-compatd.service"),
       "image must not install the development CLI wrapper as a production daemon")
assert(!image.referenced_sources.include?("runtime/dbus/org.xnix.Compatibility1.service"),
       "image must gate D-Bus activation until a production bus owner exists")

# --- KDE entry points -------------------------------------------------
%w[compatibility-center krunner kwin dolphin system-tray notifications system-settings].each do |point|
  assert(report["kde_entry_points"].include?(point), "image must cover KDE entry point: #{point}")
end

# --- Containerfile rendering ------------------------------------------
containerfile = image.render_containerfile
assert(containerfile.include?("FROM quay.io/fedora/fedora-kinoite:44"),
       "rendered Containerfile must pin the Kinoite base image")
assert(containerfile.include?("rpm-ostree install"), "rendered Containerfile must layer packages via rpm-ostree")
assert(containerfile.include?("plasma-desktop"), "rendered Containerfile must install the Plasma desktop")
assert(containerfile.include?("plasma-login-manager"),
       "rendered Containerfile must install Fedora 44's Plasma Login Manager")
assert(!containerfile.match?(/^\s+sddm(?:-kcm)?\s*\\$/),
       "rendered Containerfile must not layer obsolete SDDM packages")
assert(!containerfile.include?("        plasma-workspace-wayland "),
       "rendered Containerfile must not install the obsolete plasma-workspace-wayland subpackage")
assert(!containerfile.include?("        kwin-wayland "),
       "rendered Containerfile must not install the obsolete kwin-wayland subpackage")
assert(containerfile.include?("xdg-desktop-portal-kde"), "rendered Containerfile must install the KDE portal backend")
assert(containerfile.include?("systemctl preset-all"), "rendered Containerfile must apply unit presets")
assert(!containerfile.include?("COPY runtime/systemd/xnix-compatd.service"),
       "rendered Containerfile must not copy the development Runtime unit")

# --- Container build command -----------------------------------------
build_cmd = image.build_command(
  builder: "/usr/bin/podman",
  tag: "localhost/xnix-kinoite:test",
  storage_root: "/home/xnix-build/containers/storage",
  runroot: "/home/xnix-build/containers/runroot",
  network: "slirp4netns",
  add_host: "updates.example.test:192.0.2.10"
)
build_joined = build_cmd.join(" ")
assert(build_joined.include?("--root /home/xnix-build/containers/storage"),
       "container build must support an isolated Podman graphroot")
assert(build_joined.include?("--network slirp4netns"),
       "container build must support isolated user-mode networking")
assert(build_joined.include?("--add-host updates.example.test:192.0.2.10"),
       "container build must support a pinned repository host")
begin
  image.build_command(builder: "podman", tag: "test", network: "host")
  assert(false, "container build must reject host networking")
rescue Xnix::Image::KdeImage::ManifestError
  # expected
end

# --- Drift guard: checked-in snapshot matches the render --------------
on_disk_containerfile = project_root.join("image/kinoite/Containerfile")
assert(on_disk_containerfile.file?, "image/kinoite/Containerfile snapshot must exist")
assert(on_disk_containerfile.read.b == containerfile.b,
       "checked-in Containerfile must match the manifest render (run: ruby -Ilib lib/xnix/image/kde_image.rb containerfile)")

# --- Config overlays reference real files with expected content -------
preset = project_root.join("image/kinoite/config/systemd-preset/80-xnix.preset").read
assert(preset.include?("enable plasmalogin.service"), "unit preset must enable Plasma Login Manager")
assert(!preset.include?("enable xnix-compatd.service"),
       "unit preset must gate the development Runtime wrapper")
portal = project_root.join("image/kinoite/config/portal/xnix-portals.conf").read
assert(portal.include?("default=kde"), "portal config must prefer the KDE backend")

# --- Invalid manifest is detected -------------------------------------
Tempfile.create(["broken-manifest", ".json"]) do |file|
  broken = JSON.parse(project_root.join("image/kinoite/manifest.json").read)
  broken["package_groups"]["plasma_desktop"] = []
  broken["layered_artifacts"] << { "id" => "ghost", "source" => "does/not/exist.file", "dest" => "/x", "role" => "x" }
  file.write(JSON.generate(broken))
  file.flush
  broken_image = Xnix::Image::KdeImage.new(project_root: project_root.to_s, manifest_path: file.path)
  assert(!broken_image.valid?, "a manifest with an empty package group + missing source must be invalid")
  assert(broken_image.problems.any? { |p| p.include?("plasma_desktop") },
         "validation must flag the empty plasma_desktop group")
  assert(broken_image.problems.any? { |p| p.include?("does/not/exist.file") },
         "validation must flag the missing source file")
end

# --- Boot smoke marker logic ------------------------------------------
smoke = Xnix::Image::BootSmoke.new(image)
assert(!smoke.expected_markers.empty?, "boot smoke must have expected markers from the manifest")
good_serial = smoke.expected_markers.join("\n... boot log ...\n")
assert(smoke.booted?(good_serial), "boot smoke must pass when all markers are present")
colored_serial = smoke.expected_markers.map { |marker| "\e[0;32m#{marker}\e[0m\r" }.join("\n")
assert(smoke.booted?(colored_serial), "boot smoke must ignore ANSI control sequences in serial output")
fedora_shell_serial = "\e]3008;start=command;cwd=/var/home/xnix\e\\XNIX_BOOT_PROBE_PASS\r\n" \
                      "\e]3008;end=command;exit=success\e\\"
assert(smoke.booted?(fedora_shell_serial),
       "boot smoke must preserve markers adjacent to Fedora shell OSC sequences")
assert(!smoke.booted?("nothing useful here"), "boot smoke must fail when markers are absent")
boot_cmd = smoke.boot_command(disk_path: "/tmp/x.qcow2", firmware_path: "/tmp/OVMF.fd")
assert(boot_cmd.include?("qemu-system-x86_64"), "boot command must invoke QEMU")
assert(boot_cmd.join(" ").include?("restrict=on"), "boot command must keep networking restricted per safety constraints")
assert(boot_cmd.include?("-snapshot"), "boot smoke must preserve the original disk image")
kvm_cmd = smoke.boot_command(
  disk_path: "/tmp/x.qcow2",
  firmware_path: "/tmp/OVMF.fd",
  acceleration: "kvm"
)
assert(kvm_cmd.include?("q35,accel=kvm"), "KVM boot command must select hardware acceleration")
assert(kvm_cmd.include?("host"), "KVM boot command must expose the host CPU")
assert(smoke.active_units == %w[graphical.target plasmalogin.service],
       "boot smoke must actively probe the graphical target and login manager")
assert(smoke.probe_command.include?("systemctl is-active --quiet graphical.target"),
       "boot probe must query graphical.target")
assert(!smoke.probe_command.include?(smoke.expected_markers.first),
       "the echoed probe command must not contain the complete pass marker")

# --- Serial login state machine --------------------------------------
probe = Xnix::Image::SerialProbe.new(username: "test-user", password: "test-password", command: "probe")
assert(probe.next_input("fedora login:") == "test-user\n", "serial probe must answer the login prompt")
assert(probe.next_input("Password:") == "test-password\n", "serial probe must answer the password prompt")
assert(probe.next_input("Welcome to Fedora Linux") == "probe\n", "serial probe must run the health command")
assert(probe.next_input("Welcome to Fedora Linux") == nil, "serial probe must run each input only once")

puts "PASS: KDE Plasma atomic image pipeline is consistent"

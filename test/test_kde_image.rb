#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "pathname"
require "tempfile"
require_relative "../lib/xnix/image/kde_image"
require_relative "../lib/xnix/image/boot_smoke"

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
assert(report["desktop"]["display_manager"] == "sddm", "image must use the SDDM greeter")

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
assert(image.referenced_sources.include?("runtime/systemd/xnix-compatd.service"),
       "image must layer the Compatibility Runtime systemd unit")
assert(image.referenced_sources.include?("runtime/dbus/org.xnix.Compatibility1.service"),
       "image must layer the Compatibility Runtime D-Bus activation service")

# --- KDE entry points -------------------------------------------------
%w[compatibility-center krunner kwin dolphin system-tray notifications system-settings].each do |point|
  assert(report["kde_entry_points"].include?(point), "image must cover KDE entry point: #{point}")
end

# --- Containerfile rendering ------------------------------------------
containerfile = image.render_containerfile
assert(containerfile.include?("FROM quay.io/fedora/fedora-kinoite:41"),
       "rendered Containerfile must pin the Kinoite base image")
assert(containerfile.include?("rpm-ostree install"), "rendered Containerfile must layer packages via rpm-ostree")
assert(containerfile.include?("plasma-desktop"), "rendered Containerfile must install the Plasma desktop")
assert(!containerfile.include?("        plasma-workspace-wayland "),
       "rendered Containerfile must not install the obsolete plasma-workspace-wayland subpackage")
assert(!containerfile.include?("        kwin-wayland "),
       "rendered Containerfile must not install the obsolete kwin-wayland subpackage")
assert(containerfile.include?("xdg-desktop-portal-kde"), "rendered Containerfile must install the KDE portal backend")
assert(containerfile.include?("systemctl preset-all"), "rendered Containerfile must apply unit presets")
assert(containerfile.include?("COPY runtime/systemd/xnix-compatd.service"),
       "rendered Containerfile must copy the runtime unit")

# --- Drift guard: checked-in snapshot matches the render --------------
on_disk_containerfile = project_root.join("image/kinoite/Containerfile")
assert(on_disk_containerfile.file?, "image/kinoite/Containerfile snapshot must exist")
assert(on_disk_containerfile.read.b == containerfile.b,
       "checked-in Containerfile must match the manifest render (run: ruby -Ilib lib/xnix/image/kde_image.rb containerfile)")

# --- Config overlays reference real files with expected content -------
sddm = project_root.join("image/kinoite/config/sddm.conf.d/10-xnix.conf").read
assert(sddm.include?("DisplayServer=wayland"), "SDDM config must select the Wayland display server")
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
assert(!smoke.booted?("nothing useful here"), "boot smoke must fail when markers are absent")
boot_cmd = smoke.boot_command(disk_path: "/tmp/x.qcow2", firmware_path: "/tmp/OVMF.fd")
assert(boot_cmd.include?("qemu-system-x86_64"), "boot command must invoke QEMU")
assert(boot_cmd.join(" ").include?("restrict=on"), "boot command must keep networking restricted per safety constraints")

puts "PASS: KDE Plasma atomic image pipeline is consistent"

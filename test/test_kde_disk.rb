#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "pathname"
require "tempfile"
require_relative "../lib/xnix/image/disk_build"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
disk = Xnix::Image::DiskBuild.new(project_root: project_root.to_s)

# --- Identity + validity ---------------------------------------------
report = disk.to_h
assert(report["version"] == project_root.join("VERSION").read.strip,
       "disk build must read the version from the VERSION file")
assert(report["artifact_type"] == "kde-plasma-disk-image", "disk build must identify its artifact type")
assert(report["source_image"] == "xnix-kinoite", "disk build source image must be xnix-kinoite")
assert(report["source_reference"] == "xnix-kinoite:#{report['version']}",
       "source reference must combine image name and version")
assert(report["output_types"].include?("qcow2"), "disk build must produce a qcow2 by default")
assert(report["privileged_build_required"], "disk build must declare the privileged build requirement")
assert(disk.valid?, "checked-in disk config must be valid: #{disk.problems.join('; ')}")

# --- Cross-consistency with the image manifest ------------------------
# source_image must match the KDE image manifest's image_name.
kde_image = Xnix::Image::KdeImage.new(project_root: project_root.to_s)
assert(disk.config.fetch("source_image") == kde_image.manifest.fetch("image_name"),
       "disk source_image must match the image manifest name")

# --- Builder command rendering ----------------------------------------
command = disk.builder_command(type: "qcow2", output_dir: "/out", blueprint_path: "/tmp/bp.json")
joined = command.join(" ")
assert(command.first == "podman", "builder command must invoke podman")
assert(joined.include?("--privileged"), "bootc-image-builder must run privileged")
assert(joined.include?("--type qcow2"), "builder command must request the qcow2 type")
assert(joined.include?("--local"), "builder command must use the local container image")
assert(joined.include?("/out:/output"), "builder command must mount the output directory")
assert(joined.include?("/tmp/bp.json:/config.json:ro"), "builder command must mount the blueprint read-only")
assert(joined.include?(disk.source_reference), "builder command must reference the source image tag")
assert(joined.include?("quay.io/centos-bootc/bootc-image-builder"), "builder command must use bootc-image-builder")

# --- Output path layout -----------------------------------------------
assert(disk.output_path("qcow2", "/out") == "/out/qcow2/disk.qcow2", "qcow2 output path must follow bib layout")
assert(disk.output_path("raw", "/out") == "/out/image/disk.raw", "raw output path must follow bib layout")
assert(disk.output_path("iso", "/out") == "/out/bootiso/install.iso", "iso output path must follow bib layout")

# --- Blueprint carries the serial console so boot smoke markers appear -
kernel_append = disk.blueprint.dig("customizations", "kernel", "append").to_s
assert(kernel_append.include?("ttyS0"), "blueprint kernel args must enable the serial console for boot smoke")

# --- Unsupported type is rejected -------------------------------------
begin
  disk.builder_command(type: "vhdx", output_dir: "/out", blueprint_path: "/tmp/bp.json")
  assert(false, "unsupported output type must raise")
rescue Xnix::Image::DiskBuild::ConfigError
  # expected
end

# --- Invalid config is detected ---------------------------------------
Tempfile.create(["broken-disk", ".json"]) do |file|
  broken = JSON.parse(project_root.join("image/kinoite/disk-config.json").read)
  broken["source_image"] = "not-the-image"
  broken["output_types"] = ["vhdx"]
  file.write(JSON.generate(broken))
  file.flush
  broken_disk = Xnix::Image::DiskBuild.new(project_root: project_root.to_s, config_path: file.path)
  assert(!broken_disk.valid?, "a config with a mismatched image + unsupported type must be invalid")
  assert(broken_disk.problems.any? { |p| p.include?("must match image manifest name") },
         "validation must flag the source/manifest mismatch")
  assert(broken_disk.problems.any? { |p| p.include?("not in supported_output_types") },
         "validation must flag the unsupported output type")
end

puts "PASS: KDE Plasma disk-build pipeline is consistent"

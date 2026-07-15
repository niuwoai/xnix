#!/usr/bin/env ruby
# frozen_string_literal: true

# Disk-build driver: turns the xnix-kinoite bootc container image into a
# bootable disk image with bootc-image-builder. It validates the disk
# config, writes the blueprint bib consumes, detects the privileged podman
# toolchain, and either runs the build or reports exactly what is missing.
# It never claims to have produced a disk it did not build.
#
# Usage:
#   ruby scripts/build_kde_disk.rb                 # build the default output type(s)
#   ruby scripts/build_kde_disk.rb --check         # validate + toolchain detect only
#   ruby scripts/build_kde_disk.rb --output DIR    # output directory (default: output/)
#   ruby scripts/build_kde_disk.rb --type qcow2    # override output type

require "fileutils"
require "json"
require "pathname"
require_relative "../lib/xnix/image/disk_build"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath

def which(binary)
  ENV.fetch("PATH", "").split(File::PATH_SEPARATOR).each do |dir|
    candidate = File.join(dir, binary)
    return candidate if File.executable?(candidate) && !File.directory?(candidate)
  end
  nil
end

def arg_value(flag)
  index = ARGV.index(flag)
  index ? ARGV[index + 1] : nil
end

check_only = ARGV.delete("--check")
output_dir = Pathname.new(arg_value("--output") || PROJECT_ROOT.join("output")).expand_path
type_override = arg_value("--type")

disk = Xnix::Image::DiskBuild.new(project_root: PROJECT_ROOT.to_s)

# 1. Validate config + cross-check against the image manifest.
problems = disk.problems
unless problems.empty?
  problems.each { |problem| warn "FAIL: #{problem}" }
  abort "Disk build config is inconsistent; refusing to build."
end
puts "OK: disk config validated (#{disk.source_reference})"

types = type_override ? [type_override] : disk.output_types

# 2. Toolchain detection — bib runs as a privileged podman container.
podman = which("podman")
unless podman
  warn "NOTE: podman not found. bootc-image-builder runs as a privileged"
  warn "      podman container and reads the local container store, so the"
  warn "      disk build needs a privileged podman host. Example:"
  types.each do |type|
    warn "        podman run --rm --privileged \\"
    warn "          --security-opt label=type:unconfined_t \\"
    warn "          -v /var/lib/containers/storage:/var/lib/containers/storage \\"
    warn "          -v #{output_dir}:/output -v <blueprint>.json:/config.json:ro \\"
    warn "          #{disk.config.fetch('builder_image')} \\"
    warn "          --type #{type} --local --config /config.json #{disk.source_reference}"
  end
  abort "Toolchain unavailable — disk not built (expected in the constrained container)."
end
puts "OK: podman found (#{podman})"

if check_only
  puts "CHECK PASSED: disk config consistent, toolchain present. Skipping build (--check)."
  exit 0
end

# 3. Materialize the blueprint bib consumes, then build each type.
FileUtils.mkdir_p(output_dir)
blueprint_path = output_dir.join("xnix-kinoite-blueprint.json")
blueprint_path.write(JSON.pretty_generate(disk.blueprint) + "\n")
puts "OK: wrote blueprint #{blueprint_path}"

types.each do |type|
  command = disk.builder_command(type: type, output_dir: output_dir.to_s, blueprint_path: blueprint_path.to_s)
  puts "RUN: #{command.join(' ')}"
  abort "Disk build failed for type #{type}." unless system(*command)
  puts "DONE: #{type} -> #{disk.output_path(type, output_dir.to_s)}"
end

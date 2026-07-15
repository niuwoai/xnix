#!/usr/bin/env ruby
# frozen_string_literal: true

# Build driver for the Xnix KDE Plasma atomic image (Delivery Sequence
# step 1). It validates the manifest, guards against Containerfile drift,
# detects the container build toolchain, and either runs the build or
# reports exactly what is missing. It never pretends the image was built.
#
# The heavy compose (a multi-GB Fedora Kinoite ostree image) needs a
# privileged podman/bootc build host. In the constrained CI container the
# driver runs validation + drift + toolchain detection and exits with a
# clear "toolchain unavailable" message rather than a false success.
#
# Usage:
#   ruby scripts/build_kde_image.rb            # detect + build if possible
#   ruby scripts/build_kde_image.rb --check    # validate + drift + detect only
#   ruby scripts/build_kde_image.rb --tag NAME # override the output image tag

require "pathname"
require_relative "../lib/xnix/image/kde_image"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
CONTAINERFILE_PATH = PROJECT_ROOT.join("image/kinoite/Containerfile")

def which(binary)
  ENV.fetch("PATH", "").split(File::PATH_SEPARATOR).each do |dir|
    candidate = File.join(dir, binary)
    return candidate if File.executable?(candidate) && !File.directory?(candidate)
  end
  nil
end

check_only = ARGV.delete("--check")
tag_index = ARGV.index("--tag")
tag = tag_index ? ARGV[tag_index + 1] : nil

image = Xnix::Image::KdeImage.new(project_root: PROJECT_ROOT.to_s)
tag ||= "#{image.manifest.fetch('image_name')}:#{image.version}"

# 1. Manifest validation.
problems = image.problems
unless problems.empty?
  problems.each { |problem| warn "FAIL: #{problem}" }
  abort "Image definition is inconsistent; refusing to build."
end
puts "OK: manifest validated (#{image.manifest.fetch('image_name')})"

# 2. Drift guard — the checked-in Containerfile must match the render.
rendered = image.render_containerfile
on_disk = CONTAINERFILE_PATH.file? ? CONTAINERFILE_PATH.read : nil
if on_disk.nil? || on_disk.b != rendered.b
  warn "FAIL: image/kinoite/Containerfile is out of sync with manifest.json."
  warn "      Regenerate it: ruby -Ilib lib/xnix/image/kde_image.rb containerfile > image/kinoite/Containerfile"
  abort "Refusing to build from a drifted Containerfile."
end
puts "OK: Containerfile matches manifest (no drift)"

# 3. Toolchain detection.
podman = which("podman")
buildah = which("buildah")
builder = podman || buildah
unless builder
  warn "NOTE: no container build toolchain found (podman/buildah)."
  warn "      The Fedora Kinoite compose is multi-GB and needs a privileged"
  warn "      build host. Run this on a host with podman, e.g.:"
  warn "        podman build --file image/kinoite/Containerfile --tag #{tag} ."
  abort "Toolchain unavailable — image not built (this is expected in the constrained container)."
end
puts "OK: build toolchain found (#{builder})"

if check_only
  puts "CHECK PASSED: definition consistent, no drift, toolchain present. Skipping build (--check)."
  exit 0
end

# 4. Build.
command = [builder, "build", "--file", CONTAINERFILE_PATH.to_s, "--tag", tag, PROJECT_ROOT.to_s]
puts "RUN: #{command.join(' ')}"
success = system(*command)
abort "Image build failed (#{builder})." unless success
puts "DONE: built #{tag}"

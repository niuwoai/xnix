#!/usr/bin/env ruby
# frozen_string_literal: true

require_relative "../lib/xnix/buildroot"
require_relative "../lib/xnix/container"
require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

buildroot = Xnix::Buildroot.new
configure_command = buildroot.configure_command
build_command = buildroot.build_command
source_command = buildroot.source_command
container = Xnix::Container.new(project_root: "/workspace", version: "0.1.5")
offline_build = container.cache_run_command(build_command)

assert(configure_command.first == "make", "configuration must invoke make")
assert(configure_command.include?("-C"), "configuration must set the Buildroot source directory")
assert(configure_command.include?(Xnix::Buildroot::SOURCE_DIRECTORY), "configuration must use the pinned Buildroot source")
assert(configure_command.include?("O=#{Xnix::Buildroot::OUTPUT_DIRECTORY}"), "configuration must use the managed output directory")
assert(configure_command.include?("BR2_EXTERNAL=#{Xnix::Buildroot::EXTERNAL_DIRECTORY}"), "configuration must use the Xnix external tree")
assert(configure_command.include?(Xnix::Buildroot::DEFCONFIG), "configuration must select the Xnix defconfig")
assert(build_command == ["make", "-C", Xnix::Buildroot::SOURCE_DIRECTORY, "O=#{Xnix::Buildroot::OUTPUT_DIRECTORY}"], "build command must reuse the configured output")
assert(source_command.last == "source", "dependency download must use Buildroot's source target")
assert(offline_build.fetch(offline_build.index("--network") + 1) == "none", "full build must run without network access")
assert(offline_build.include?("--mount"), "full build must mount the managed cache")
assert(offline_build.none? { |argument| argument.include?("type=bind") }, "full build must not mount a host directory")
networked_download = container.networked_cache_run_command(source_command)
assert(networked_download.fetch(networked_download.index("--network") + 1) == "bridge", "dependency download requires only bridge networking")

project_root = Pathname.new(__dir__).join("..").realpath
defconfig = project_root.join("buildroot/configs/xnix_x86_64_defconfig").read
assert(defconfig.include?("BR2_JLEVEL=1"), "Buildroot must use a single job within the memory limit")

puts "PASS: Buildroot command unit tests"

#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/full_smoke.rb").read

%w[build fetch-sources configure-system download-system build-system boot-system].each do |step|
  assert(script.include?("\"#{step}\""), "full smoke script must include #{step}")
end

build_index = script.index("\"build\"")
fetch_index = script.index("\"fetch-sources\"")
configure_index = script.index("\"configure-system\"")
download_index = script.index("\"download-system\"")
build_system_index = script.index("\"build-system\"")
boot_index = script.index("\"boot-system\"")

assert(build_index < fetch_index, "full smoke must build the container image before fetching sources")
assert(fetch_index < configure_index, "full smoke must fetch Buildroot before configuring")
assert(configure_index < download_index, "full smoke must configure before downloading package sources")
assert(download_index < build_system_index, "full smoke must download package sources before building")
assert(build_system_index < boot_index, "full smoke must build the system before booting QEMU")
assert(script.include?("Xnix::Milestone.full_build_required?"), "full smoke must remain gated to tenth versions")
assert(project_root.join("scripts/container.rb").read.include?("\"90s\""), "container boot-system must allow enough time for QEMU login markers")

puts "PASS: full smoke script unit tests"

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

%w[
  build-tools
  build
  fetch-sources
  configure-system
  download-system
  build-system
  prepare-ssh-test-key
  configure-wine-guest
  download-wine-guest
  build-ssh-wine-guest
  fetch-known-winapp
  boot-system
  staged-launcher-dispatch-smoke
  winapp-guest-wine-smoke
].each do |step|
  assert(script.include?("\"#{step}\""), "full smoke script must include #{step}")
end

build_tools_index = script.index("\"build-tools\"")
build_index = script.index("\"build\"")
fetch_index = script.index("\"fetch-sources\"")
configure_index = script.index("\"configure-system\"")
download_index = script.index("\"download-system\"")
build_system_index = script.index("\"build-system\"")
prepare_key_index = script.index("\"prepare-ssh-test-key\"")
configure_wine_index = script.index("\"configure-wine-guest\"")
download_wine_index = script.index("\"download-wine-guest\"")
build_wine_index = script.index("\"build-ssh-wine-guest\"")
fetch_known_app_index = script.index("\"fetch-known-winapp\"")
boot_index = script.index("\"boot-system\"")
known_app_smoke_index = script.index("\"staged-launcher-dispatch-smoke\"")
fixture_app_smoke_index = script.index("\"winapp-guest-wine-smoke\"")

assert(build_tools_index < build_index, "full smoke must build the tools image before the tested runtime image")
assert(build_index < fetch_index, "full smoke must build the container image before fetching sources")
assert(fetch_index < configure_index, "full smoke must fetch Buildroot before configuring")
assert(configure_index < download_index, "full smoke must configure before downloading package sources")
assert(download_index < build_system_index, "full smoke must download package sources before building")
assert(build_system_index < prepare_key_index, "full smoke must build the base system before preparing SSH test keys")
assert(prepare_key_index < configure_wine_index, "full smoke must prepare SSH test keys before configuring the Wine guest")
assert(configure_wine_index < download_wine_index, "full smoke must configure the Wine guest before downloading Wine guest sources")
assert(download_wine_index < build_wine_index, "full smoke must download Wine guest sources before building the SSH Wine guest")
assert(build_wine_index < fetch_known_app_index, "full smoke must build the SSH Wine guest before fetching known Windows apps")
assert(fetch_known_app_index < boot_index, "full smoke must fetch known Windows apps before the final base boot smoke")
assert(build_system_index < boot_index, "full smoke must build the system before booting QEMU")
assert(boot_index < known_app_smoke_index, "full smoke must prove the base QEMU boot before the known Windows app smoke")
assert(known_app_smoke_index < fixture_app_smoke_index, "full smoke must prove known app smoke before fixture regression smoke")
assert(script.include?("require_pass_step"), "full smoke must reject skipped real-app smoke lanes")
assert(script.include?("PASS: staged managed launcher dispatch smoke"), "full smoke must require staged launcher known Windows app PASS evidence")
assert(script.include?("PASS: QEMU guest real Windows app Wine smoke"), "full smoke must require fixture Windows app PASS evidence")
assert(script.include?("Xnix::Milestone.full_build_required?"), "full smoke must remain gated to twentieth versions")
assert(script.include?("Xnix::FullSmokeReport"), "full smoke must emit a structured report")
assert(script.include?("full-smoke-report.json"), "full smoke must write a JSON report")
assert(script.include?("full-smoke-report.md"), "full smoke must write a Markdown report")
assert(script.include?("JSON.pretty_generate"), "full smoke JSON report must be pretty generated")
assert(project_root.join("scripts/container.rb").read.include?("\"180s\""), "container boot-system must allow enough time for QEMU login markers")

puts "PASS: full smoke script unit tests"

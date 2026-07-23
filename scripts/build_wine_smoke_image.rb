#!/usr/bin/env ruby
# frozen_string_literal: true

require "open3"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
IMAGE = ENV.fetch("XNIX_WINE_IMAGE", "xnix-wine-smoke:local")
DOCKERFILE_RELATIVE = "containers/wine-smoke.Dockerfile"
DOCKERFILE = PROJECT_ROOT.join(DOCKERFILE_RELATIVE)

argv = [
  "docker", "build",
  "--pull",
  "-f", DOCKERFILE.to_s,
  "-t", IMAGE,
  PROJECT_ROOT.to_s
]

stdout, stderr, status = Open3.capture3(*argv, chdir: PROJECT_ROOT.to_s)
puts stdout unless stdout.empty?
warn stderr unless stderr.empty?

if status.success?
  puts "PASS: built local Wine smoke image #{IMAGE}"
  exit 0
end

warn "FAIL: local Wine smoke image build failed"
exit 1

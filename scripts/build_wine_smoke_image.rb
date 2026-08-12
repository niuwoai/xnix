#!/usr/bin/env ruby
# frozen_string_literal: true

require "open3"
require "pathname"
require "securerandom"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
IMAGE = ENV.fetch("XNIX_WINE_IMAGE", "xnix-wine-smoke:local")
BASE_IMAGE = ENV.fetch("XNIX_WINE_BASE_IMAGE", "debian:bookworm-slim")
PLATFORM = ENV.fetch("XNIX_WINE_PLATFORM", "linux/amd64")
BASE_PLATFORM = ENV.fetch("XNIX_WINE_BASE_PLATFORM", PLATFORM)
PULL_BASE = ENV.fetch("XNIX_WINE_PULL", "0") == "1"
DOCKERFILE_RELATIVE = "containers/wine-smoke.Dockerfile"
DOCKERFILE = PROJECT_ROOT.join(DOCKERFILE_RELATIVE)

argv = ["docker", "build", "--platform", PLATFORM, "--build-arg", "XNIX_WINE_BASE_IMAGE=#{BASE_IMAGE}", "--build-arg", "XNIX_WINE_BASE_PLATFORM=#{BASE_PLATFORM}"]
argv << "--pull" if PULL_BASE
argv.concat([
  "-f", DOCKERFILE.to_s,
  "-t", IMAGE,
  PROJECT_ROOT.to_s
])

def run_command(*argv)
  stdout, stderr, status = Open3.capture3(*argv, chdir: PROJECT_ROOT.to_s)
  puts stdout unless stdout.empty?
  warn stderr unless stderr.empty?
  status.success?
end

def image_platform(image)
  stdout, stderr, status = Open3.capture3(
    "docker", "image", "inspect", image, "--format", "{{.Os}}/{{.Architecture}}",
    chdir: PROJECT_ROOT.to_s
  )
  warn stderr unless stderr.empty?
  return nil unless status.success?

  stdout.strip
end

def commit_fallback_image(image, platform)
  container_name = "xnix-wine-smoke-build-#{SecureRandom.hex(6)}"
  install_command = [
    "if [ \"$(dpkg --print-architecture)\" = \"amd64\" ]; then dpkg --add-architecture i386; fi",
    "apt-get update",
    "if [ \"$(dpkg --print-architecture)\" = \"amd64\" ]; then wine_packages=\"wine wine32 wine64\"; else wine_packages=\"wine wine64\"; fi",
    "apt-get install -y --no-install-recommends ca-certificates procps ${wine_packages} x11-utils xvfb",
    "rm -rf /var/lib/apt/lists/*"
  ].join(" && ")

  begin
    ok = run_command(
      "docker", "run",
      "--name", container_name,
      "--platform", platform,
      "--cpus", "2",
      "--memory", "2g",
      "--pids-limit", "512",
      "--env", "DEBIAN_FRONTEND=noninteractive",
      "debian:bookworm-slim",
      "sh", "-lc", install_command
    )
    return false unless ok

    run_command(
      "docker", "commit",
      "--change", "ENV DEBIAN_FRONTEND=noninteractive",
      "--change", "ENV WINEDEBUG=-all",
      "--change", "WORKDIR /work",
      container_name,
      image
    )
  ensure
    Open3.capture3("docker", "rm", "-f", container_name, chdir: PROJECT_ROOT.to_s)
  end
end

def buildx_available?
  _stdout, _stderr, status = Open3.capture3("docker", "buildx", "version", chdir: PROJECT_ROOT.to_s)
  status.success?
end

unless buildx_available?
  warn "WARN: docker buildx unavailable; trying run-and-commit fallback"
  unless commit_fallback_image(IMAGE, PLATFORM)
    warn "FAIL: local Wine smoke image fallback build failed"
    exit 1
  end
  actual_platform = image_platform(IMAGE)
  unless actual_platform == PLATFORM
    warn "FAIL: local Wine smoke image platform is #{actual_platform}, expected #{PLATFORM}"
    exit 1
  end
  puts "PASS: built local Wine smoke image #{IMAGE} for #{PLATFORM}"
  exit 0
end

status_success = run_command(*argv)

unless status_success
  warn "FAIL: local Wine smoke image build failed"
  exit 1
end

actual_platform = image_platform(IMAGE)
unless actual_platform == PLATFORM
  warn "WARN: docker build produced #{actual_platform}, expected #{PLATFORM}; trying run-and-commit fallback"
  unless commit_fallback_image(IMAGE, PLATFORM)
    warn "FAIL: local Wine smoke image fallback build failed"
    exit 1
  end
  actual_platform = image_platform(IMAGE)
  unless actual_platform == PLATFORM
    warn "FAIL: local Wine smoke image platform is #{actual_platform}, expected #{PLATFORM}"
    exit 1
  end
end

puts "PASS: built local Wine smoke image #{IMAGE} for #{PLATFORM}"

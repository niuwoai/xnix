#!/usr/bin/env ruby
# frozen_string_literal: true

require "open3"
require "pathname"
require_relative "../lib/xnix/container"
require_relative "../lib/xnix/buildroot"
require_relative "../lib/xnix/qemu"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip

def colima_context?
  output, status = Open3.capture2("docker", "context", "show")
  status.success? && output.strip == "colima"
end

abort "Xnix commands require the Colima Docker context" unless colima_context?

container = Xnix::Container.new(project_root: PROJECT_ROOT.to_s, version: VERSION)
buildroot = Xnix::Buildroot.new
qemu = Xnix::Qemu.new

case ARGV.shift
when "build"
  exec(*container.build_command)
when "offline-run"
  abort "Usage: ruby scripts/container.rb offline-run COMMAND [ARGUMENT ...]" if ARGV.empty?

  exec(*container.offline_run_command(ARGV))
when "fetch-sources"
  abort "Usage: ruby scripts/container.rb fetch-sources" unless ARGV.empty?

  exec(*container.source_retrieval_command(["ruby", "scripts/fetch_buildroot.rb"]))
when "configure-system"
  abort "Usage: ruby scripts/container.rb configure-system" unless ARGV.empty?

  exec(*container.cache_run_command(buildroot.configure_command))
when "build-system"
  abort "Usage: ruby scripts/container.rb build-system" unless ARGV.empty?

  exec(*container.cache_run_command(buildroot.build_command))
when "download-system"
  abort "Usage: ruby scripts/container.rb download-system" unless ARGV.empty?

  exec(*container.networked_cache_run_command(buildroot.source_command))
when "boot-system"
  abort "Usage: ruby scripts/container.rb boot-system" unless ARGV.empty?

  exec(*container.cache_run_command(["timeout", "45s", *qemu.boot_command]))
else
  abort "Usage: ruby scripts/container.rb {boot-system|build|build-system|configure-system|download-system|fetch-sources|offline-run COMMAND [ARGUMENT ...]}"
end

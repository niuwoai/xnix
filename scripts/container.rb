#!/usr/bin/env ruby
# frozen_string_literal: true

require "open3"
require "pathname"
require_relative "../lib/xnix/container"
require_relative "../lib/xnix/buildroot"
require_relative "../lib/xnix/qemu"
require_relative "../lib/xnix/ssh_test_key"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip
DOCKER_BIN = Xnix::Container.docker_bin

def docker_available?
  system(DOCKER_BIN, "--version", out: File::NULL, err: File::NULL)
rescue SystemCallError
  false
end

def colima_context?
  output, status = Open3.capture2(DOCKER_BIN, "context", "show")
  status.success? && output.strip == "colima"
rescue SystemCallError
  false
end

abort "Xnix commands require a Docker CLI. Install Docker or set XNIX_DOCKER_BIN=/absolute/path/to/docker." unless docker_available?
abort "Xnix commands require the Colima Docker context. Run `docker context use colima` or set XNIX_DOCKER_BIN to a Docker-compatible CLI that uses the Colima context." unless colima_context?

container = Xnix::Container.new(project_root: PROJECT_ROOT.to_s, version: VERSION, docker_bin: DOCKER_BIN)
buildroot = Xnix::Buildroot.new
wine_buildroot = Xnix::Buildroot.new(profile: :wine_i386)
qemu = Xnix::Qemu.new

case ARGV.shift
when "build"
  exec(*container.build_command)
when "build-tools"
  abort "Usage: ruby scripts/container.rb build-tools" unless ARGV.empty?

  exec(*container.build_tools_command)
when "offline-run"
  abort "Usage: ruby scripts/container.rb offline-run COMMAND [ARGUMENT ...]" if ARGV.empty?

  exec(*container.offline_run_command(ARGV))
when "runtime-activation-smoke"
  abort "Usage: ruby scripts/container.rb runtime-activation-smoke" unless ARGV.empty?

  exec(*container.runtime_activation_smoke_command)
when "runtime-dbus-smoke"
  abort "Usage: ruby scripts/container.rb runtime-dbus-smoke" unless ARGV.empty?

  exec(*container.runtime_dbus_smoke_command)
when "runtime-owner-candidate-smoke"
  abort "Usage: ruby scripts/container.rb runtime-owner-candidate-smoke" unless ARGV.empty?

  exec(*container.runtime_owner_candidate_smoke_command)
when "kde-center-dbus-smoke"
  abort "Usage: ruby scripts/container.rb kde-center-dbus-smoke" unless ARGV.empty?

  exec(*container.kde_center_dbus_smoke_command)
when "fetch-sources"
  abort "Usage: ruby scripts/container.rb fetch-sources" unless ARGV.empty?

  exec(*container.source_retrieval_command(["ruby", "scripts/fetch_buildroot.rb"]))
when "configure-system"
  abort "Usage: ruby scripts/container.rb configure-system" unless ARGV.empty?

  exec(*container.tools_cache_run_command(buildroot.configure_command))
when "configure-wine-guest"
  abort "Usage: ruby scripts/container.rb configure-wine-guest" unless ARGV.empty?

  exec(*container.tools_cache_run_command(wine_buildroot.configure_command))
when "build-system"
  abort "Usage: ruby scripts/container.rb build-system" unless ARGV.empty?

  exec(*container.tools_cache_run_command(buildroot.build_command))
when "build-wine-guest"
  abort "Usage: ruby scripts/container.rb build-wine-guest" unless ARGV.empty?

  exec(*container.tools_cache_run_command(wine_buildroot.build_command))
when "prepare-ssh-test-key"
  abort "Usage: ruby scripts/container.rb prepare-ssh-test-key" unless ARGV.empty?

  exec(*container.tools_cache_run_command(["ruby", "scripts/prepare_ssh_test_key.rb"]))
when "build-ssh-test-system"
  abort "Usage: ruby scripts/container.rb build-ssh-test-system" unless ARGV.empty?

  exec(*container.tools_cache_run_command(buildroot.build_command(ssh_test_key: true)))
when "build-ssh-wine-guest"
  abort "Usage: ruby scripts/container.rb build-ssh-wine-guest" unless ARGV.empty?

  exec(*container.tools_cache_run_command(wine_buildroot.build_command(ssh_test_key: true)))
when "start-build-ssh-wine-guest"
  abort "Usage: ruby scripts/container.rb start-build-ssh-wine-guest" unless ARGV.empty?

  name = "xnix-wine-guest-build-#{VERSION.gsub(/[^a-zA-Z0-9]+/, "-")}"
  exec(*container.observed_cache_run_command(name: name, command: wine_buildroot.build_command(ssh_test_key: true)))
when "ssh-smoke"
  abort "Usage: ruby scripts/container.rb ssh-smoke" unless ARGV.empty?

  exec(*container.tools_cache_run_command(["ruby", "scripts/ssh_smoke.rb"]))
when "winapp-guest-wine-smoke"
  abort "Usage: ruby scripts/container.rb winapp-guest-wine-smoke" unless ARGV.empty?

  exec(*container.tools_cache_run_command(["ruby", "scripts/winapp_guest_wine_smoke.rb"]))
when "fetch-known-winapp"
  abort "Usage: ruby scripts/container.rb fetch-known-winapp" unless ARGV.empty?

  exec(*container.known_winapp_fetch_command)
when "known-winapp-guest-wine-smoke"
  abort "Usage: ruby scripts/container.rb known-winapp-guest-wine-smoke" unless ARGV.empty?

  exec(*container.known_winapp_guest_wine_smoke_command)
when "staged-launcher-dispatch-smoke"
  abort "Usage: ruby scripts/container.rb staged-launcher-dispatch-smoke" unless ARGV.empty?

  exec(*container.staged_launcher_dispatch_smoke_command)
when "runtime-status-owner-service-session-bus-smoke"
  abort "Usage: ruby scripts/container.rb runtime-status-owner-service-session-bus-smoke" unless ARGV.empty?

  exec(*container.runtime_status_owner_service_session_bus_smoke_command)
when "kde-controlled-launch-action-smoke"
  abort "Usage: ruby scripts/container.rb kde-controlled-launch-action-smoke" unless ARGV.empty?

  exec(*container.kde_controlled_launch_action_smoke_command)
when "desktop-trigger-request-preflight-smoke"
  abort "Usage: ruby scripts/container.rb desktop-trigger-request-preflight-smoke" unless ARGV.empty?

  exec(*container.desktop_trigger_request_preflight_smoke_command)
when "kde-controlled-launch-action-dbus-fixture-smoke"
  abort "Usage: ruby scripts/container.rb kde-controlled-launch-action-dbus-fixture-smoke" unless ARGV.empty?

  exec(*container.kde_controlled_launch_action_dbus_fixture_smoke_command)
when "dbus-controlled-launch-owner-fixture-smoke"
  abort "Usage: ruby scripts/container.rb dbus-controlled-launch-owner-fixture-smoke" unless ARGV.empty?

  exec(*container.dbus_controlled_launch_owner_fixture_smoke_command)
when "start-build-system"
  abort "Usage: ruby scripts/container.rb start-build-system" unless ARGV.empty?

  name = "xnix-full-build-#{VERSION.gsub(/[^a-zA-Z0-9]+/, "-")}"
  exec(*container.observed_cache_run_command(name: name, command: buildroot.build_command))
when "download-system"
  abort "Usage: ruby scripts/container.rb download-system" unless ARGV.empty?

  exec(*container.networked_cache_run_command(buildroot.source_command))
when "download-wine-guest"
  abort "Usage: ruby scripts/container.rb download-wine-guest" unless ARGV.empty?

  exec(*container.networked_cache_run_command(wine_buildroot.source_command))
when "boot-system"
  abort "Usage: ruby scripts/container.rb boot-system" unless ARGV.empty?

  exec(*container.tools_cache_run_command(["timeout", "180s", *qemu.boot_command]))
else
  abort "Usage: ruby scripts/container.rb {boot-system|build|build-ssh-test-system|build-ssh-wine-guest|build-system|build-tools|build-wine-guest|configure-system|configure-wine-guest|dbus-controlled-launch-owner-fixture-smoke|desktop-trigger-request-preflight-smoke|download-system|download-wine-guest|fetch-known-winapp|fetch-sources|kde-center-dbus-smoke|kde-controlled-launch-action-dbus-fixture-smoke|kde-controlled-launch-action-smoke|known-winapp-guest-wine-smoke|offline-run COMMAND [ARGUMENT ...]|prepare-ssh-test-key|runtime-activation-smoke|runtime-dbus-smoke|runtime-owner-candidate-smoke|runtime-status-owner-service-session-bus-smoke|ssh-smoke|staged-launcher-dispatch-smoke|start-build-ssh-wine-guest|start-build-system|winapp-guest-wine-smoke}"
end

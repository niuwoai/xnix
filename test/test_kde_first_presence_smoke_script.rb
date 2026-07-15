#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
script_path = project_root.join("scripts/kde_first_presence_smoke.rb")
script = script_path.read

%w[
  desktop-identity-plan
  desktop-entry-preview
  mimeapps-preview
  desktop-activation-bundle-preview
  desktop-activation-staging-preview
  desktop-activation-transaction-preview
  desktop-activation-status-preview
  kde-entrypoints-preview
  kde-action-card-deck-preview
  kde-center-page-preview
  kde-center-page-sections-preview
  kde-center-page-section-detail-preview
  file-open-preview
  dolphin-drop-preview
  dolphin-ai-analysis-preview
  window-identity-preview
  tray-status-preview
  notification-preview
  settings-preview
  runtime-owner-route-manifest-preview
  runtime-method-parity-manifest-preview
  runtime-write-gate-preview
  ai-diagnostic-input-preview
  ai-diagnostic-recommendation-preview
  ai-repair-approval-gate-preview
].each do |preview|
  assert(script.include?(preview), "KDE-first presence smoke must inspect #{preview}")
end

%w[
  launcher
  task-manager
  file-manager
  system-tray
  notifications
  compatibility-center
  settings
].each do |entrypoint|
  assert(script.include?(entrypoint), "KDE-first presence smoke must assert #{entrypoint}")
end

%w[
  host_root_modified
  network_required
  privileged_container_required
  backend_details_exposed
  raw_command_exposed
  raw_windows_executable_exposed
  execution_started
  backend_launch_enabled
  launch_enabled
  request_object_created
  permission_granted
  file_content_read
  file_paths_exposed
  ai_provider_call_enabled
].each do |flag|
  assert(script.include?(flag), "KDE-first presence smoke must enforce #{flag}")
end

%w[
  desktop-entry
  dolphin-service-menu
  mimeapps-list
  desktop-integration-manifest
  desktop-activation-receipt
].each do |file_id|
  assert(script.include?(file_id), "KDE-first presence smoke must require staged #{file_id}")
end

assert(script.include?("GOCACHE"), "KDE-first presence smoke must keep Go cache inside the project")
assert(script.include?(".cache"), "KDE-first presence smoke must use the project cache directory")
assert(!script.include?("scripts/container.rb"), "KDE-first presence smoke must not run Docker container smokes")
assert(!script.include?("boot-system"), "KDE-first presence smoke must not boot QEMU")
assert(!script.include?("docs/claude-code-implementation-packages.md"), "KDE-first presence smoke must not touch Claude Code implementation package work")

puts "PASS: KDE-first presence smoke script unit tests"

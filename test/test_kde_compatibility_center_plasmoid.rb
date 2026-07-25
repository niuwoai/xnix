#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
version = project_root.join("VERSION").read.strip
metadata = JSON.parse(project_root.join("kde/plasmoids/org.xnix.compatibilitycenter/metadata.json").read)
qml = project_root.join("kde/plasmoids/org.xnix.compatibilitycenter/contents/ui/main.qml").read

assert(metadata.fetch("KPlugin").fetch("Id") == "org.xnix.compatibilitycenter", "Plasmoid metadata must keep the Compatibility Center id")
assert(metadata.fetch("KPlugin").fetch("Version") == version, "Plasmoid metadata must match VERSION")
assert(metadata.fetch("KPlugin").fetch("X-Xnix-RuntimeModelCommand") == "xnix-kde-center-model", "Plasmoid metadata must declare the Runtime model command")

%w[
  runtimeModelCommand
  centerPreviewCommand
  kde-center-page-preview
  known_app_gui_evidence_count
  known_app_owner_controlled_gui_evidence_count
  known_app_owner_managed_copy_verified_count
  known_app_gui_evidence_cards
  guiEvidenceCardFields
  display_name
  recipe_backed
  recipe_app_id
  smoke_status
  compatibility_state
  center_card_state
  execution_evidence_recorded
  staged_launcher_verified
  owner_controlled_runtime_launch_verified
  owner_managed_copy_verified
  owner_service_call_ready
  owner_evidence_handoff_ready
  owner_evidence_relative_path
  runtime_dispatch_verified
  primary_action_id
  primary_action_kind
  primary_action_label
  desktop_callable_route
  desktop_callable_runtime_method
  desktop_callable_execution_type
  desktop_dbus_method
  desktop_evidence_handle_forwarded
  kde_forwarded_argument_kind
  kde_forwarded_arguments
  owner_service_args_exposed_to_kde
  wine-guest-gui-smoke
  winapp-smoke-container-x-gui
  known-application-gui-smoke
  evidence-relative-path
  Real\ Windows\ GUI\ evidence
].each do |token|
  assert(qml.include?(token.gsub("\\ ", " ")), "Plasmoid QML must include #{token}")
end

%w[
  direct_launch_enabled
  backend_launch_enabled
  host_root_modified
  backend_details_exposed
].each do |unsafe_field|
  assert(!qml.include?("#{unsafe_field}: true"), "Plasmoid QML must not enable #{unsafe_field}")
end

%w[
  qemu-system
  wineboot
  WINEPREFIX
  /home/
  /Users/
  /tmp/
].each do |forbidden|
  assert(!qml.include?(forbidden), "Plasmoid QML must not expose backend detail #{forbidden}")
end

puts "PASS: KDE Compatibility Center plasmoid GUI evidence card"

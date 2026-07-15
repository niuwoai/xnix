#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "optparse"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
DEFAULT_GO_BIN = Pathname.new("/opt/homebrew/bin/go").executable? ? "/opt/homebrew/bin/go" : "go"
DEFAULT_REGISTRY = "runtime/recipes/registry.json"
DEFAULT_APP = "org.xnix.sample.notepad"
DEFAULT_DECISION = "approved"
DEFAULT_MODE = "development"
DEFAULT_FILE_URI = "file:///home/xnix/Documents/sample.txt"
EXPECTED_ENTRYPOINTS = %w[
  launcher
  task-manager
  file-manager
  system-tray
  notifications
  compatibility-center
  settings
].freeze
EXPECTED_STAGED_FILES = %w[
  desktop-entry
  dolphin-service-menu
  mimeapps-list
  desktop-integration-manifest
  desktop-activation-receipt
].freeze
SAFETY_FALSE_KEYS = %w[
  backend_details_exposed
  backend_launch_enabled
  backend_process_started
  compatibility_storage_exposed
  direct_host_file_access
  direct_file_read_enabled
  execution_started
  file_content_read
  file_paths_exposed
  file_writes_performed
  host_permission_changed
  host_root_modified
  ai_provider_call_enabled
  launch_allowed
  launch_enabled
  live_tray_bridge_enabled
  network_required
  notifications_sent
  permission_granted
  privileged_container_required
  raw_command_exposed
  raw_executable_exposed
  raw_windows_executable_exposed
  request_object_created
  request_objects_created
  settings_persisted
  task_manager_entry_active
].freeze
FORBIDDEN_TEXT_PATTERNS = [
  /docker\.sock/i,
  %r{/Users/},
  %r{/private/},
  %r{/var/},
  /\.wine/i,
  /Program Files/i,
  /\bwine\b/i,
  /\bproton\b/i,
  /qemu-system/i
].freeze

class SmokeFailure < StandardError; end

def parse_options(argv)
  options = {
    go_bin: DEFAULT_GO_BIN,
    registry: DEFAULT_REGISTRY,
    app: DEFAULT_APP,
    decision: DEFAULT_DECISION,
    mode: DEFAULT_MODE
  }

  parser = OptionParser.new do |opts|
    opts.banner = "Usage: ruby scripts/kde_first_presence_smoke.rb [options] [file-uri]"
    opts.on("--go-bin PATH", "Go binary to use") { |value| options[:go_bin] = value }
    opts.on("--registry PATH", "Recipe registry path") { |value| options[:registry] = value }
    opts.on("--app ID", "Application id") { |value| options[:app] = value }
    opts.on("--decision DECISION", "Review decision") { |value| options[:decision] = value }
    opts.on("--mode MODE", "Activation mode") { |value| options[:mode] = value }
  end
  parser.parse!(argv)
  options[:file_uri] = argv.first || DEFAULT_FILE_URI
  options
end

def go_preview(options, command, *arguments)
  stdout = go_preview_output(options, command, *arguments)
  JSON.parse(stdout)
rescue JSON::ParserError => e
  raise SmokeFailure, "#{command} did not return JSON: #{e.message}"
end

def go_preview_text(options, command, *arguments)
  go_preview_output(options, command, *arguments)
end

def go_preview_output(options, command, *arguments)
  env = {
    "GOCACHE" => PROJECT_ROOT.join(".cache", "go-build").to_s
  }
  command_line = [
    options.fetch(:go_bin),
    "run",
    "./cmd/xnix-runtime-go",
    command,
    *arguments
  ]
  stdout, stderr, status = Open3.capture3(env, *command_line, chdir: PROJECT_ROOT.to_s)
  raise SmokeFailure, "#{command} failed: #{stderr.strip}" unless status.success?

  stdout
end

def base_args(options)
  ["--registry", options.fetch(:registry), "--app", options.fetch(:app)]
end

def decision_args(options)
  [*base_args(options), "--decision", options.fetch(:decision), options.fetch(:file_uri)]
end

def mode_args(options)
  [*base_args(options), "--mode", options.fetch(:mode)]
end

def collect_payloads(options)
  {
    "desktop-identity-plan" => go_preview(options, "desktop-identity-plan", *base_args(options)),
    "desktop-entry-preview" => go_preview_text(options, "desktop-entry-preview", *base_args(options)),
    "mimeapps-preview" => go_preview_text(options, "mimeapps-preview", *base_args(options)),
    "desktop-activation-bundle-preview" => go_preview(options, "desktop-activation-bundle-preview", *base_args(options)),
    "desktop-activation-staging-preview" => go_preview(options, "desktop-activation-staging-preview", *mode_args(options)),
    "desktop-activation-transaction-preview" => go_preview(options, "desktop-activation-transaction-preview", *mode_args(options)),
    "desktop-activation-status-preview" => go_preview(options, "desktop-activation-status-preview", *mode_args(options)),
    "kde-entrypoints-preview" => go_preview(options, "kde-entrypoints-preview", *decision_args(options)),
    "kde-action-card-deck-preview" => go_preview(options, "kde-action-card-deck-preview", *decision_args(options)),
    "kde-center-page-preview" => go_preview(options, "kde-center-page-preview", *decision_args(options)),
    "kde-center-page-sections-preview" => go_preview(options, "kde-center-page-sections-preview", *decision_args(options)),
    "kde-center-page-section-detail-preview" => go_preview(options, "kde-center-page-section-detail-preview", *base_args(options), "--section", "diagnostics", "--decision", options.fetch(:decision), options.fetch(:file_uri)),
    "file-open-preview" => go_preview(options, "file-open-preview", "--registry", options.fetch(:registry), "--app", options.fetch(:app), options.fetch(:file_uri)),
    "dolphin-drop-preview" => go_preview(options, "dolphin-drop-preview", "--registry", options.fetch(:registry), "--app", options.fetch(:app), options.fetch(:file_uri)),
    "dolphin-ai-analysis-preview" => go_preview(options, "dolphin-ai-analysis-preview", "--registry", options.fetch(:registry), "--app", options.fetch(:app), options.fetch(:file_uri)),
    "window-identity-preview" => go_preview(options, "window-identity-preview", *base_args(options)),
    "tray-status-preview" => go_preview(options, "tray-status-preview", *base_args(options)),
    "notification-preview" => go_preview(options, "notification-preview", *base_args(options), "--event", "approval-required"),
    "settings-preview" => go_preview(options, "settings-preview", *base_args(options)),
    "runtime-owner-route-manifest-preview" => go_preview(options, "runtime-owner-route-manifest-preview", "--root", "."),
    "runtime-method-parity-manifest-preview" => go_preview(options, "runtime-method-parity-manifest-preview", "--root", ".")
  }
end

def assert(condition, check_id, message)
  raise SmokeFailure, "#{check_id}: #{message}" unless condition
end

def assert_text_artifacts(desktop_entry, mimeapps, desktop_file)
  assert(desktop_entry.include?("[Desktop Entry]"), "start_menu.desktop_entry", "desktop entry text must be rendered")
  assert(desktop_entry.include?("Exec=xnix-compat-launch --app"), "start_menu.launcher", "desktop entry must route launch through the Runtime")
  assert(desktop_entry.include?("X-Xnix-RuntimeOwned=true"), "start_menu.runtime_owned", "desktop entry must identify Runtime ownership")
  assert(mimeapps.include?(desktop_file), "file_manager.mimeapps", "MIME apps output must target the generated desktop file")
end

def deep_each(value, path = [], &block)
  yield(path, value)
  case value
  when Hash
    value.each { |key, child| deep_each(child, [*path, key], &block) }
  when Array
    value.each_with_index { |child, index| deep_each(child, [*path, index], &block) }
  end
end

def safe_payload?(payload, allowed_file_uri)
  deep_each(payload) do |path, value|
    key = path.last.to_s
    if SAFETY_FALSE_KEYS.include?(key) && value == true
      raise SmokeFailure, "global_safety.#{key}: expected false, got true"
    end
    next unless value.is_a?(String)
    next if value == allowed_file_uri
    next if value.start_with?("applications:")
    next if value.start_with?("usr/share/")

    FORBIDDEN_TEXT_PATTERNS.each do |pattern|
      raise SmokeFailure, "global_safety.text: forbidden value at #{path.join(".")}: #{value}" if value.match?(pattern)
    end
  end
end

def entry_point_ids(payload)
  payload.fetch("entry_point_ids") do
    payload.fetch("entry_points", []).map { |entry| entry.fetch("id") }
  end
end

def assert_entrypoints(payload)
  ids = entry_point_ids(payload)
  missing = EXPECTED_ENTRYPOINTS - ids
  assert(missing.empty?, "kde_entrypoints.complete", "missing entrypoints: #{missing.join(", ")}")
  assert(payload.fetch("surface_type") == "kde-first-release-entrypoints", "kde_entrypoints.surface", "unexpected surface type")
  assert(payload.fetch("normal_application_surface") == true, "kde_entrypoints.normal_surface", "normal application surface must be true")
  assert(payload.fetch("kde_policy_owner") == false, "kde_entrypoints.runtime_boundary", "KDE must not own policy")
end

def assert_staging(payload)
  ids = payload.fetch("planned_file_ids")
  missing = EXPECTED_STAGED_FILES - ids
  assert(missing.empty?, "activation_staging.files", "missing staged files: #{missing.join(", ")}")
  assert(payload.fetch("file_writes_performed") == false, "activation_staging.non_executing", "staging preview must not write files")
  assert(payload.fetch("rollback_receipt_planned") == true, "activation_staging.rollback", "rollback receipt must be planned")
end

def assert_transaction(payload)
  assert(payload.fetch("transaction_committed") == false, "activation_transaction.commit", "activation commit must remain disabled")
  assert(payload.fetch("write_gate").fetch("dispatch_enabled") == false, "activation_transaction.write_gate", "activation write dispatch must remain disabled")
  rollback_steps = payload.fetch("rollback_steps", [])
  assert(!rollback_steps.empty?, "activation_transaction.rollback", "rollback steps must be present")
end

def assert_file_manager(file_open, drop, ai_analysis)
  assert(file_open.fetch("source") == "dolphin-service-menu", "file_manager.source", "file-open source must be Dolphin service menu")
  assert(file_open.fetch("portal_required") == true, "file_manager.portal", "file-open must require Portal review")
  assert(drop.fetch("source") == "dolphin-drag-and-drop", "file_manager.drop", "drop source must be Dolphin drag-and-drop")
  assert(ai_analysis.fetch("source") == "dolphin-ai-action", "file_manager.ai_source", "AI analysis source must be Dolphin AI action")
  assert(ai_analysis.fetch("file_content_read") == false, "file_manager.ai_privacy", "AI analysis must not read file contents")
  assert(ai_analysis.fetch("file_paths_exposed") == false, "file_manager.ai_paths", "AI analysis must not expose file paths")
end

def assert_center(page, sections, diagnostics)
  assert(page.fetch("request_type") == "kde-center-page-preview", "center.page", "center page preview required")
  assert(page.fetch("action_deck").fetch("request_type") == "kde-action-card-deck-preview", "center.deck", "center page must include action deck")
  section_ids = sections.fetch("sections").map { |section| section.fetch("id") }
  %w[overview backend activation execution launch window files tray notifications actions settings diagnostics].each do |section|
    assert(section_ids.include?(section), "center.section.#{section}", "center sections must include #{section}")
  end
  assert(diagnostics.fetch("section_id") == "diagnostics", "center.diagnostics", "diagnostics detail required")
end

def assert_route_baseline(payload)
  counts = payload.fetch("route_counts")
  checks = payload.fetch("checks").to_h { |check| [check.fetch("id"), check.fetch("status")] }
  assert(counts.fetch("ruby_legacy") == 0, "routes.ruby_legacy", "Ruby legacy routes must be zero")
  assert(checks.fetch("method-parity") == "pass", "routes.method_parity", "method parity must pass")
  assert(checks.fetch("write-route-gate") == "pass", "routes.write_gate", "write route gate must pass")
  assert(checks.fetch("host-safety-boundary") == "pass", "routes.host_safety", "host safety boundary must pass")
end

def assert_method_parity(payload)
  assert(payload.fetch("read_only_method_parity_ready") == true, "method_parity.ready", "read-only method parity must be ready")
  assert(payload.fetch("write_method_dispatch_enabled") == false, "method_parity.write_dispatch", "write dispatch must remain disabled")
end

def run_smoke(options)
  payloads = collect_payloads(options)
  payloads.each_value { |payload| safe_payload?(payload, options.fetch(:file_uri)) }

  desktop_file = payloads.fetch("desktop-identity-plan").fetch("desktop_file")
  assert_text_artifacts(payloads.fetch("desktop-entry-preview"), payloads.fetch("mimeapps-preview"), desktop_file)
  assert_entrypoints(payloads.fetch("kde-entrypoints-preview"))
  assert_staging(payloads.fetch("desktop-activation-staging-preview"))
  assert_transaction(payloads.fetch("desktop-activation-transaction-preview"))
  assert_file_manager(
    payloads.fetch("file-open-preview"),
    payloads.fetch("dolphin-drop-preview"),
    payloads.fetch("dolphin-ai-analysis-preview")
  )
  assert_center(
    payloads.fetch("kde-center-page-preview"),
    payloads.fetch("kde-center-page-sections-preview"),
    payloads.fetch("kde-center-page-section-detail-preview")
  )
  assert_route_baseline(payloads.fetch("runtime-owner-route-manifest-preview"))
  assert_method_parity(payloads.fetch("runtime-method-parity-manifest-preview"))

  route_counts = payloads.fetch("runtime-owner-route-manifest-preview").fetch("route_counts")
  puts "PASS: KDE-first presence smoke"
  puts "application: #{options.fetch(:app)}"
  puts "entrypoints: #{EXPECTED_ENTRYPOINTS.join(", ")}"
  puts "route-baseline: #{route_counts.fetch("total")} total, #{route_counts.fetch("go_routed")} Go-routed, #{route_counts.fetch("c_core_backed")} C-backed, #{route_counts.fetch("ruby_legacy")} Ruby legacy"
  puts "execution: disabled by Runtime gates"
  puts "host-root: unchanged"
end

begin
  run_smoke(parse_options(ARGV))
rescue SmokeFailure => e
  warn "FAIL: KDE-first presence smoke"
  warn "reason: #{e.message}"
  exit 1
end
